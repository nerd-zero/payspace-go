package payspace

import (
	"context"
	"math"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// buildTransport creates the HTTP transport middleware chain.
// Order (outermost first): retry → rateLimit → logging → auth → base
func buildTransport(cfg *clientConfig, auth *tokenManager, logger *zap.Logger, rl *rateLimitTracker) http.RoundTripper {
	var base http.RoundTripper
	if cfg.httpClient != nil && cfg.httpClient.Transport != nil {
		base = cfg.httpClient.Transport
	} else {
		base = http.DefaultTransport
	}

	// auth is innermost: injects Bearer token
	transport := &authTransport{base: base, auth: auth, logger: logger}

	// logging wraps auth
	logged := &loggingTransport{base: transport, logger: logger}

	// rate limit wraps logging
	rateLimited := &rateLimitTransport{base: logged, logger: logger, tracker: rl, wait: cfg.rateLimitWait}

	// retry is outermost
	retried := &retryTransport{
		base:       rateLimited,
		logger:     logger,
		maxRetries: cfg.maxRetries,
		backoff:    cfg.retryBackoff,
		auth:       auth,
	}

	return retried
}

// authTransport injects the Bearer token into requests.
type authTransport struct {
	base   http.RoundTripper
	auth   *tokenManager
	logger *zap.Logger
}

func (t *authTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.auth.Token(req.Context())
	if err != nil {
		t.logger.Error("failed to obtain token for request", zap.Error(err))
		return nil, err
	}
	req2 := req.Clone(req.Context())
	req2.Header.Set("Authorization", "Bearer "+token)
	return t.base.RoundTrip(req2)
}

// loggingTransport logs HTTP requests and responses.
type loggingTransport struct {
	base   http.RoundTripper
	logger *zap.Logger
}

func (t *loggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	t.logger.Debug("http request",
		zap.String("method", req.Method),
		zap.String("url", req.URL.String()),
	)

	resp, err := t.base.RoundTrip(req)
	duration := time.Since(start)

	if err != nil {
		t.logger.Error("http request failed",
			zap.String("method", req.Method),
			zap.String("url", req.URL.String()),
			zap.Duration("duration", duration),
			zap.Error(err),
		)
		return nil, err
	}

	logLevel := zap.InfoLevel
	if resp.StatusCode >= 400 {
		logLevel = zap.WarnLevel
	}
	if resp.StatusCode >= 500 {
		logLevel = zap.ErrorLevel
	}

	if ce := t.logger.Check(logLevel, "http response"); ce != nil {
		ce.Write(
			zap.String("method", req.Method),
			zap.String("url", req.URL.String()),
			zap.Int("status", resp.StatusCode),
			zap.Duration("duration", duration),
			zap.String("rate_limit_remaining", resp.Header.Get("X-RateLimit-Remaining")),
		)
	}

	return resp, nil
}

// rateLimitTransport optionally waits when a 429 is received.
type rateLimitTransport struct {
	base    http.RoundTripper
	logger  *zap.Logger
	tracker *rateLimitTracker
	wait    bool
}

func (t *rateLimitTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Check pre-emptive rate limit
	current := t.tracker.Current()
	if current.Limit > 0 && current.Remaining <= 0 && time.Now().Before(current.Reset) && t.wait {
		waitDur := time.Until(current.Reset)
		t.logger.Warn("pre-emptive rate limit wait",
			zap.Duration("wait", waitDur),
			zap.Time("reset", current.Reset),
		)
		select {
		case <-time.After(waitDur):
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
	}

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusTooManyRequests && t.wait {
		rl := parseRateLimitHeaders(resp.Header)
		waitDur := time.Until(rl.Reset)
		if waitDur <= 0 {
			waitDur = 5 * time.Second
		}
		t.logger.Warn("rate limited (429), waiting for reset",
			zap.Duration("wait", waitDur),
			zap.Int("limit", rl.Limit),
		)
		resp.Body.Close()
		select {
		case <-time.After(waitDur):
		case <-req.Context().Done():
			return nil, req.Context().Err()
		}
		return t.base.RoundTrip(req)
	}

	return resp, nil
}

// retryTransport retries on transient errors (5xx, network errors).
// It also handles 401 by forcing a token refresh.
type retryTransport struct {
	base       http.RoundTripper
	logger     *zap.Logger
	maxRetries int
	backoff    time.Duration
	auth       *tokenManager
}

func (t *retryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.maxRetries; attempt++ {
		if attempt > 0 {
			wait := t.backoff * time.Duration(math.Pow(2, float64(attempt-1)))
			t.logger.Warn("retrying request",
				zap.Int("attempt", attempt),
				zap.Duration("backoff", wait),
				zap.String("method", req.Method),
				zap.String("url", req.URL.String()),
			)
			select {
			case <-time.After(wait):
			case <-req.Context().Done():
				return nil, req.Context().Err()
			}
		}

		resp, err = t.base.RoundTrip(req)
		if err != nil {
			if isRetryableError(err) {
				continue
			}
			return nil, err
		}

		// 401: force token refresh and retry once
		if resp.StatusCode == http.StatusUnauthorized && attempt == 0 {
			resp.Body.Close()
			t.logger.Warn("received 401, forcing token refresh")
			if _, refreshErr := t.auth.ForceRefresh(context.Background()); refreshErr != nil {
				t.logger.Error("token refresh failed after 401", zap.Error(refreshErr))
				return resp, nil
			}
			continue
		}

		// 5xx: retry
		if resp.StatusCode >= 500 {
			resp.Body.Close()
			continue
		}

		return resp, nil
	}

	return resp, err
}

func isRetryableError(err error) bool {
	// Context cancellations and deadlines are not retryable.
	if err == context.Canceled || err == context.DeadlineExceeded {
		return false
	}
	return true
}

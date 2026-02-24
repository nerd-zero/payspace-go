package payspace

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
)

const tokenRefreshBuffer = 5 * time.Minute

type tokenManagerConfig struct {
	clientID     string
	clientSecret string
	scope        Scope
	identityURL  string
	httpClient   *http.Client
	logger       *zap.Logger
}

type tokenManager struct {
	mu           sync.RWMutex
	token        string
	expiresAt    time.Time
	companies    []GroupCompany
	clientID     string
	clientSecret string
	scope        Scope
	identityURL  string
	httpClient   *http.Client
	logger       *zap.Logger
}

type tokenResponse struct {
	AccessToken    string         `json:"access_token"`
	ExpiresIn      int            `json:"expires_in"`
	TokenType      string         `json:"token_type"`
	GroupCompanies []GroupCompany `json:"group_companies"`
}

func newTokenManager(cfg tokenManagerConfig) *tokenManager {
	hc := cfg.httpClient
	if hc == nil {
		hc = &http.Client{Timeout: 30 * time.Second}
	}
	return &tokenManager{
		clientID:     cfg.clientID,
		clientSecret: cfg.clientSecret,
		scope:        cfg.scope,
		identityURL:  cfg.identityURL,
		httpClient:   hc,
		logger:       cfg.logger,
	}
}

// Token returns a valid access token, refreshing if necessary.
func (tm *tokenManager) Token(ctx context.Context) (string, error) {
	tm.mu.RLock()
	if tm.token != "" && time.Now().Before(tm.expiresAt.Add(-tokenRefreshBuffer)) {
		token := tm.token
		tm.mu.RUnlock()
		tm.logger.Debug("using cached token", zap.Time("expires_at", tm.expiresAt))
		return token, nil
	}
	tm.mu.RUnlock()

	return tm.refresh(ctx)
}

// ForceRefresh forces a token refresh regardless of expiry.
func (tm *tokenManager) ForceRefresh(ctx context.Context) (string, error) {
	tm.logger.Info("forcing token refresh")
	return tm.refresh(ctx)
}

func (tm *tokenManager) refresh(ctx context.Context) (string, error) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	// Double-check: another goroutine may have refreshed while we waited for the lock.
	if tm.token != "" && time.Now().Before(tm.expiresAt.Add(-tokenRefreshBuffer)) {
		tm.logger.Debug("token already refreshed by another goroutine")
		return tm.token, nil
	}

	tm.logger.Info("refreshing access token", zap.String("identity_url", tm.identityURL))

	form := url.Values{}
	form.Set("client_id", tm.clientID)
	form.Set("client_secret", tm.clientSecret)
	form.Set("scope", string(tm.scope))

	endpoint := tm.identityURL + "/connect/token"
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("payspace: failed to create token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", defaultUserAgent)

	resp, err := tm.httpClient.Do(req)
	if err != nil {
		tm.logger.Error("token request failed", zap.Error(err))
		return "", fmt.Errorf("payspace: token request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("payspace: failed to read token response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		tm.logger.Error("token request returned non-200",
			zap.Int("status_code", resp.StatusCode),
			zap.String("body", string(body)),
		)
		return "", fmt.Errorf("payspace: authentication failed with status %d: %s", resp.StatusCode, string(body))
	}

	var tokenResp tokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", fmt.Errorf("payspace: failed to decode token response: %w", err)
	}

	tm.token = tokenResp.AccessToken
	tm.expiresAt = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)
	tm.companies = tokenResp.GroupCompanies

	tm.logger.Info("access token refreshed",
		zap.Time("expires_at", tm.expiresAt),
		zap.Int("company_groups", len(tm.companies)),
	)

	return tm.token, nil
}

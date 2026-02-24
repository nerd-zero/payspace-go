package payspace

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestLoggingTransport(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Remaining", "99")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	logger := testNopLogger()
	lt := &loggingTransport{
		base:   http.DefaultTransport,
		logger: logger,
	}

	client := &http.Client{Transport: lt}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if got := resp.Header.Get("X-RateLimit-Remaining"); got != "99" {
		t.Errorf("expected X-RateLimit-Remaining='99', got %q", got)
	}
}

func TestRetryTransport_ServerError(t *testing.T) {
	var attempts int32
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n < 3 {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	logger := testNopLogger()

	// Create a minimal token manager with a pre-seeded token.
	auth := &tokenManager{
		token:     "test-token",
		expiresAt: time.Now().Add(1 * time.Hour),
		logger:    logger,
	}

	rt := &retryTransport{
		base:       http.DefaultTransport,
		logger:     logger,
		maxRetries: 3,
		backoff:    1 * time.Millisecond, // fast retries for testing
		auth:       auth,
	}

	client := &http.Client{Transport: rt}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 after retries, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&attempts); got != 3 {
		t.Errorf("expected 3 attempts, got %d", got)
	}
}

func TestRetryTransport_401RefreshesToken(t *testing.T) {
	var attempts int32
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&attempts, 1)
		if n == 1 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"Message":"Unauthorized"}`))
			return
		}
		// After token refresh, the second attempt should succeed.
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	})

	// Token endpoint for the forced refresh.
	mux.HandleFunc("POST /connect/token", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"access_token": "refreshed-token",
			"expires_in":   3600,
			"token_type":   "Bearer",
		})
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	logger := testNopLogger()

	auth := newTokenManager(tokenManagerConfig{
		clientID:     "test-id",
		clientSecret: "test-secret",
		scope:        ScopeFullAccess,
		identityURL:  srv.URL,
		logger:       logger,
	})
	// Pre-seed a token that is within the refresh buffer window so
	// ForceRefresh actually performs a refresh instead of returning the cached token.
	auth.token = "old-token"
	auth.expiresAt = time.Now().Add(2 * time.Minute) // within 5-minute tokenRefreshBuffer

	rt := &retryTransport{
		base:       http.DefaultTransport,
		logger:     logger,
		maxRetries: 2,
		backoff:    1 * time.Millisecond,
		auth:       auth,
	}

	client := &http.Client{Transport: rt}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200 after token refresh, got %d", resp.StatusCode)
	}
	if got := atomic.LoadInt32(&attempts); got != 2 {
		t.Errorf("expected 2 attempts (initial + retry after refresh), got %d", got)
	}
	// Verify the token was refreshed.
	if auth.token != "refreshed-token" {
		t.Errorf("expected token to be 'refreshed-token', got %q", auth.token)
	}
}

func TestAuthTransport_InjectsToken(t *testing.T) {
	var receivedAuth string
	mux := http.NewServeMux()
	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		receivedAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	logger := testNopLogger()

	auth := &tokenManager{
		token:     "my-bearer-token",
		expiresAt: time.Now().Add(1 * time.Hour),
		logger:    logger,
	}

	at := &authTransport{
		base:   http.DefaultTransport,
		auth:   auth,
		logger: logger,
	}

	client := &http.Client{Transport: at}
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL+"/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	expected := "Bearer my-bearer-token"
	if receivedAuth != expected {
		t.Errorf("expected Authorization=%q, got %q", expected, receivedAuth)
	}
}

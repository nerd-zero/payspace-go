package payspace

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestTokenManager_Token(t *testing.T) {
	var called atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called.Add(1)
		if r.URL.Path != "/connect/token" {
			t.Errorf("expected path /connect/token, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if err := r.ParseForm(); err != nil {
			t.Fatalf("failed to parse form: %v", err)
		}
		if r.FormValue("client_id") != "test-id" {
			t.Errorf("expected client_id=test-id, got %s", r.FormValue("client_id"))
		}
		if r.FormValue("client_secret") != "test-secret" {
			t.Errorf("expected client_secret=test-secret, got %s", r.FormValue("client_secret"))
		}
		if r.FormValue("scope") != string(ScopeFullAccess) {
			t.Errorf("expected scope=%s, got %s", ScopeFullAccess, r.FormValue("scope"))
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResponse{
			AccessToken: "fresh-token-123",
			ExpiresIn:   3600,
			TokenType:   "Bearer",
			GroupCompanies: []GroupCompany{
				{GroupID: 1, GroupDescription: "Test Group"},
			},
		})
	}))
	defer srv.Close()

	tm := newTokenManager(tokenManagerConfig{
		clientID:     "test-id",
		clientSecret: "test-secret",
		scope:        ScopeFullAccess,
		identityURL:  srv.URL,
		logger:       testNopLogger(),
	})

	token, err := tm.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "fresh-token-123" {
		t.Errorf("expected token 'fresh-token-123', got %q", token)
	}
	if called.Load() != 1 {
		t.Errorf("expected 1 call to identity server, got %d", called.Load())
	}
}

func TestTokenManager_TokenCaching(t *testing.T) {
	var called atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called.Add(1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResponse{
			AccessToken: "cached-token",
			ExpiresIn:   3600,
			TokenType:   "Bearer",
		})
	}))
	defer srv.Close()

	tm := newTokenManager(tokenManagerConfig{
		clientID:     "test-id",
		clientSecret: "test-secret",
		scope:        ScopeFullAccess,
		identityURL:  srv.URL,
		logger:       testNopLogger(),
	})

	// First call fetches a new token.
	token1, err := tm.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error on first call: %v", err)
	}

	// Second call should return cached token.
	token2, err := tm.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error on second call: %v", err)
	}

	if token1 != token2 {
		t.Errorf("expected same cached token, got %q and %q", token1, token2)
	}
	if called.Load() != 1 {
		t.Errorf("expected 1 call to identity server (caching), got %d", called.Load())
	}
}

func TestTokenManager_TokenRefresh(t *testing.T) {
	var called atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called.Add(1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResponse{
			AccessToken: "refreshed-token",
			ExpiresIn:   3600,
			TokenType:   "Bearer",
		})
	}))
	defer srv.Close()

	tm := newTokenManager(tokenManagerConfig{
		clientID:     "test-id",
		clientSecret: "test-secret",
		scope:        ScopeFullAccess,
		identityURL:  srv.URL,
		logger:       testNopLogger(),
	})

	// Pre-seed a token that is about to expire (within the 5-minute buffer).
	tm.token = "old-token"
	tm.expiresAt = time.Now().Add(2 * time.Minute) // less than 5-minute buffer

	token, err := tm.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if token != "refreshed-token" {
		t.Errorf("expected refreshed-token, got %q", token)
	}
	if called.Load() != 1 {
		t.Errorf("expected 1 refresh call, got %d", called.Load())
	}
}

func TestTokenManager_ForceRefresh(t *testing.T) {
	var called atomic.Int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called.Add(1)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenResponse{
			AccessToken: "force-refreshed-token",
			ExpiresIn:   3600,
			TokenType:   "Bearer",
		})
	}))
	defer srv.Close()

	tm := newTokenManager(tokenManagerConfig{
		clientID:     "test-id",
		clientSecret: "test-secret",
		scope:        ScopeFullAccess,
		identityURL:  srv.URL,
		logger:       testNopLogger(),
	})

	// First, get a token (this caches it).
	token1, err := tm.Token(context.Background())
	if err != nil {
		t.Fatalf("unexpected error on first call: %v", err)
	}
	if called.Load() != 1 {
		t.Errorf("expected 1 call after Token(), got %d", called.Load())
	}

	// Invalidate the token so ForceRefresh's inner refresh() will actually fetch.
	tm.mu.Lock()
	tm.token = ""
	tm.mu.Unlock()

	// ForceRefresh should request a new token.
	token2, err := tm.ForceRefresh(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if called.Load() != 2 {
		t.Errorf("expected 2 calls after ForceRefresh, got %d", called.Load())
	}
	// Both calls return the same access token from the server, but verify ForceRefresh returned something.
	_ = token1
	if token2 != "force-refreshed-token" {
		t.Errorf("expected force-refreshed-token, got %q", token2)
	}
}

func TestTokenManager_AuthFailure(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"invalid_client"}`))
	}))
	defer srv.Close()

	tm := newTokenManager(tokenManagerConfig{
		clientID:     "bad-id",
		clientSecret: "bad-secret",
		scope:        ScopeFullAccess,
		identityURL:  srv.URL,
		logger:       testNopLogger(),
	})

	_, err := tm.Token(context.Background())
	if err == nil {
		t.Fatal("expected error for failed authentication")
	}
}

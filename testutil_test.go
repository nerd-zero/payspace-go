package payspace

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

// testServerAndClient creates a test HTTP server and a Client configured to use it.
func testServerAndClient(t *testing.T, handler http.Handler) (*httptest.Server, *Client) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	logger := zaptest.NewLogger(t, zaptest.Level(zapcore.DebugLevel))

	client, err := NewClient("test-client-id", "test-client-secret",
		WithAPIBaseURL(srv.URL),
		WithIdentityURL(srv.URL),
		WithLogger(logger),
		WithRetry(0, 0),
		WithRateLimitWait(false),
	)
	if err != nil {
		t.Fatalf("failed to create test client: %v", err)
	}
	// Pre-seed a token so tests don't need a real auth server.
	client.auth.token = "test-token"
	client.auth.expiresAt = time.Now().Add(24 * time.Hour)

	return srv, client
}

// writeJSON writes a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// writeODataList writes an OData collection response.
func writeODataList(w http.ResponseWriter, items any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"@odata.context": "https://test.payspace.com/odata/v2.0/1/$metadata",
		"value":          items,
	})
}

// testNopLogger returns a no-op zap logger.
func testNopLogger() *zap.Logger {
	return zap.NewNop()
}

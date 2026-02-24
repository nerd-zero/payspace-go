package payspace

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("my-client-id", "my-secret",
		WithLogger(testNopLogger()),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if client == nil {
		t.Fatal("expected non-nil client")
	}
	if client.Employees == nil {
		t.Error("expected Employees service to be initialized")
	}
	if client.EmployeePayroll == nil {
		t.Error("expected EmployeePayroll service to be initialized")
	}
	if client.EmployeePayslips == nil {
		t.Error("expected EmployeePayslips service to be initialized")
	}
	if client.EmployeeLeave == nil {
		t.Error("expected EmployeeLeave service to be initialized")
	}
	if client.EmployeeSuspensions == nil {
		t.Error("expected EmployeeSuspensions service to be initialized")
	}
	if client.Companies == nil {
		t.Error("expected Companies service to be initialized")
	}
}

func TestNewClient_MissingCredentials(t *testing.T) {
	_, err := NewClient("", "secret", WithLogger(testNopLogger()))
	if err == nil {
		t.Fatal("expected error for empty clientID")
	}

	_, err = NewClient("client-id", "", WithLogger(testNopLogger()))
	if err == nil {
		t.Fatal("expected error for empty clientSecret")
	}

	_, err = NewClient("", "", WithLogger(testNopLogger()))
	if err == nil {
		t.Fatal("expected error for empty clientID and clientSecret")
	}
}

func TestNewClient_DefaultConfig(t *testing.T) {
	client, err := NewClient("id", "secret", WithLogger(testNopLogger()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Default environment is Staging.
	if client.apiBaseURL != "https://apistaging.payspace.com" {
		t.Errorf("expected staging API URL, got %s", client.apiBaseURL)
	}
	if client.identityURL != "https://staging-identity.yourhcm.com" {
		t.Errorf("expected staging identity URL, got %s", client.identityURL)
	}
	if client.apiVersion != "v2.0" {
		t.Errorf("expected api version v2.0, got %s", client.apiVersion)
	}
	if client.userAgent != "payspace.com" {
		t.Errorf("expected user agent payspace.com, got %s", client.userAgent)
	}
}

func TestNewClient_WithOptions(t *testing.T) {
	client, err := NewClient("id", "secret",
		WithEnvironment(Production),
		WithTimeout(60*time.Second),
		WithRetry(5, 1*time.Second),
		WithScope(ScopeReadOnly),
		WithUserAgent("test-agent"),
		WithAPIVersion("v3.0"),
		WithLogger(testNopLogger()),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if client.apiBaseURL != "https://api.payspace.com" {
		t.Errorf("expected production API URL, got %s", client.apiBaseURL)
	}
	if client.identityURL != "https://identity.yourhcm.com" {
		t.Errorf("expected production identity URL, got %s", client.identityURL)
	}
	if client.apiVersion != "v3.0" {
		t.Errorf("expected api version v3.0, got %s", client.apiVersion)
	}
	if client.userAgent != "test-agent" {
		t.Errorf("expected user agent test-agent, got %s", client.userAgent)
	}
}

func TestClient_ODataURL(t *testing.T) {
	mux := http.NewServeMux()
	_, client := testServerAndClient(t, mux)

	url := client.odataURL(42, "Employee")
	expected := client.apiBaseURL + "/odata/v2.0/42/Employee"
	if url != expected {
		t.Errorf("expected %s, got %s", expected, url)
	}

	url = client.odataURL(1, "EmployeePayRate")
	expected = client.apiBaseURL + "/odata/v2.0/1/EmployeePayRate"
	if url != expected {
		t.Errorf("expected %s, got %s", expected, url)
	}
}

func TestClient_APIError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"Message": "Invalid request",
			"Details": []map[string]string{
				{"Message": "EmployeeNumber is required"},
			},
		})
	})
	_, client := testServerAndClient(t, mux)

	_, _, err := client.Employees.List(context.Background(), 1, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "Invalid request" {
		t.Errorf("expected message 'Invalid request', got %q", apiErr.Message)
	}
}

func TestClient_NotFound(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee(999)", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusNotFound, map[string]any{
			"Message": "Employee not found",
		})
	})
	_, client := testServerAndClient(t, mux)

	_, _, err := client.Employees.Get(context.Background(), 1, 999)
	if err == nil {
		t.Fatal("expected error")
	}
	if !IsNotFound(err) {
		t.Errorf("expected IsNotFound to be true, got false for error: %v", err)
	}
}

func TestClient_RateLimited(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-RateLimit-Limit", "100")
		w.Header().Set("X-RateLimit-Remaining", "0")
		w.Header().Set("X-RateLimit-Reset", "1700000000")
		writeJSON(w, http.StatusTooManyRequests, map[string]any{
			"Message": "Rate limit exceeded",
		})
	})
	_, client := testServerAndClient(t, mux)

	_, _, err := client.Employees.List(context.Background(), 1, nil)
	if err == nil {
		t.Fatal("expected error")
	}

	rlErr, ok := err.(*RateLimitError)
	if !ok {
		t.Fatalf("expected *RateLimitError, got %T: %v", err, err)
	}
	if rlErr.Limit != 100 {
		t.Errorf("expected Limit=100, got %d", rlErr.Limit)
	}
	if rlErr.Remaining != 0 {
		t.Errorf("expected Remaining=0, got %d", rlErr.Remaining)
	}
	if !IsRateLimited(err) {
		t.Error("expected IsRateLimited to be true")
	}
}

func TestClient_TokenInfo(t *testing.T) {
	mux := http.NewServeMux()
	_, client := testServerAndClient(t, mux)

	expiresAt, hasToken := client.TokenInfo()
	if !hasToken {
		t.Error("expected hasToken to be true")
	}
	if expiresAt.Before(time.Now()) {
		t.Error("expected expiresAt to be in the future")
	}
}

func TestClient_ODataURLv1(t *testing.T) {
	mux := http.NewServeMux()
	_, client := testServerAndClient(t, mux)

	url := client.odataURLv1(42, "EmployeePayslip/comment")
	expected := client.apiBaseURL + "/odata/v1.0/42/EmployeePayslip/comment"
	if url != expected {
		t.Errorf("expected %s, got %s", expected, url)
	}
}

func TestClient_APIErrorStringFormatNoMessage(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		// Empty body - no JSON message.
		json.NewEncoder(w).Encode(map[string]any{})
	})
	_, client := testServerAndClient(t, mux)

	_, _, err := client.Employees.List(context.Background(), 1, nil)
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	expected := "payspace: HTTP 500"
	if apiErr.Error() != expected {
		t.Errorf("expected error string %q, got %q", expected, apiErr.Error())
	}
}

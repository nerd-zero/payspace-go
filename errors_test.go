package payspace

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"
)

func TestIsNotFound(t *testing.T) {
	err := &APIError{StatusCode: http.StatusNotFound, Message: "not found"}
	if !IsNotFound(err) {
		t.Error("expected IsNotFound to return true for 404")
	}

	err2 := &APIError{StatusCode: http.StatusBadRequest, Message: "bad request"}
	if IsNotFound(err2) {
		t.Error("expected IsNotFound to return false for 400")
	}

	// Non-APIError should return false.
	if IsNotFound(io.EOF) {
		t.Error("expected IsNotFound to return false for non-APIError")
	}
}

func TestIsUnauthorized(t *testing.T) {
	err := &APIError{StatusCode: http.StatusUnauthorized, Message: "unauthorized"}
	if !IsUnauthorized(err) {
		t.Error("expected IsUnauthorized to return true for 401")
	}

	err2 := &APIError{StatusCode: http.StatusForbidden, Message: "forbidden"}
	if IsUnauthorized(err2) {
		t.Error("expected IsUnauthorized to return false for 403")
	}

	if IsUnauthorized(io.EOF) {
		t.Error("expected IsUnauthorized to return false for non-APIError")
	}
}

func TestIsRateLimited(t *testing.T) {
	// RateLimitError should be detected.
	rlErr := &RateLimitError{
		APIError: &APIError{StatusCode: http.StatusTooManyRequests},
		Limit:    100,
	}
	if !IsRateLimited(rlErr) {
		t.Error("expected IsRateLimited to return true for RateLimitError")
	}

	// APIError with 429 should also be detected.
	apiErr := &APIError{StatusCode: http.StatusTooManyRequests}
	if !IsRateLimited(apiErr) {
		t.Error("expected IsRateLimited to return true for APIError with 429")
	}

	// Non-429 APIError should be false.
	apiErr2 := &APIError{StatusCode: http.StatusBadRequest}
	if IsRateLimited(apiErr2) {
		t.Error("expected IsRateLimited to return false for 400")
	}

	// Non-APIError should be false.
	if IsRateLimited(io.EOF) {
		t.Error("expected IsRateLimited to return false for non-APIError")
	}
}

func TestIsValidationError(t *testing.T) {
	err := &APIError{
		StatusCode: http.StatusBadRequest,
		Message:    "Validation failed",
		Details: []ErrorDetail{
			{Message: "EmployeeNumber is required"},
		},
	}
	if !IsValidationError(err) {
		t.Error("expected IsValidationError to return true")
	}

	// 400 without details should be false.
	err2 := &APIError{StatusCode: http.StatusBadRequest, Message: "Bad request"}
	if IsValidationError(err2) {
		t.Error("expected IsValidationError to return false without details")
	}

	// Non-400 with details should be false.
	err3 := &APIError{
		StatusCode: http.StatusInternalServerError,
		Details:    []ErrorDetail{{Message: "unexpected"}},
	}
	if IsValidationError(err3) {
		t.Error("expected IsValidationError to return false for 500")
	}
}

func TestAPIError_Error(t *testing.T) {
	err := &APIError{StatusCode: 400, Message: "Bad request"}
	expected := "payspace: HTTP 400: Bad request"
	if err.Error() != expected {
		t.Errorf("expected %q, got %q", expected, err.Error())
	}

	err2 := &APIError{StatusCode: 500}
	expected2 := "payspace: HTTP 500"
	if err2.Error() != expected2 {
		t.Errorf("expected %q, got %q", expected2, err2.Error())
	}
}

func TestRateLimitError_Error(t *testing.T) {
	resetTime := time.Date(2024, 1, 15, 12, 0, 0, 0, time.UTC)
	err := &RateLimitError{
		APIError:  &APIError{StatusCode: 429},
		Limit:     100,
		Remaining: 0,
		Reset:     resetTime,
	}
	got := err.Error()
	if got == "" {
		t.Error("expected non-empty error string")
	}
	// Should contain key information.
	if !containsSubstring(got, "rate limit exceeded") {
		t.Errorf("expected error to contain 'rate limit exceeded', got %q", got)
	}
	if !containsSubstring(got, "limit=100") {
		t.Errorf("expected error to contain 'limit=100', got %q", got)
	}
}

func TestCheckResponse_Success(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(nil)),
	}
	if err := checkResponse(resp); err != nil {
		t.Fatalf("expected nil error for 200, got %v", err)
	}

	// 201 should also be success.
	resp2 := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(bytes.NewReader(nil)),
	}
	if err := checkResponse(resp2); err != nil {
		t.Fatalf("expected nil error for 201, got %v", err)
	}

	// 204 should also be success.
	resp3 := &http.Response{
		StatusCode: http.StatusNoContent,
		Body:       io.NopCloser(bytes.NewReader(nil)),
	}
	if err := checkResponse(resp3); err != nil {
		t.Fatalf("expected nil error for 204, got %v", err)
	}
}

func TestCheckResponse_400(t *testing.T) {
	body := map[string]any{
		"Message": "Validation error",
		"Details": []map[string]string{
			{"Message": "field required"},
		},
	}
	jsonBody, _ := json.Marshal(body)

	resp := &http.Response{
		StatusCode: http.StatusBadRequest,
		Header:     http.Header{},
		Body:       io.NopCloser(bytes.NewReader(jsonBody)),
	}
	err := checkResponse(resp)
	if err == nil {
		t.Fatal("expected error for 400")
	}

	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 400 {
		t.Errorf("expected status 400, got %d", apiErr.StatusCode)
	}
	if apiErr.Message != "Validation error" {
		t.Errorf("expected message 'Validation error', got %q", apiErr.Message)
	}
}

func TestCheckResponse_429(t *testing.T) {
	body := map[string]any{
		"Message": "Rate limited",
	}
	jsonBody, _ := json.Marshal(body)

	resp := &http.Response{
		StatusCode: http.StatusTooManyRequests,
		Header: http.Header{
			"X-Ratelimit-Limit":     []string{"100"},
			"X-Ratelimit-Remaining": []string{"0"},
			"X-Ratelimit-Reset":     []string{"1700000000"},
		},
		Body: io.NopCloser(bytes.NewReader(jsonBody)),
	}
	err := checkResponse(resp)
	if err == nil {
		t.Fatal("expected error for 429")
	}

	rlErr, ok := err.(*RateLimitError)
	if !ok {
		t.Fatalf("expected *RateLimitError, got %T", err)
	}
	if rlErr.Limit != 100 {
		t.Errorf("expected Limit=100, got %d", rlErr.Limit)
	}
	if rlErr.Remaining != 0 {
		t.Errorf("expected Remaining=0, got %d", rlErr.Remaining)
	}
}

// containsSubstring is a helper for checking if a string contains a substring.
func containsSubstring(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > 0 && contains(s, sub))
}

func contains(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

package payspace

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// APIError represents an error response from the PaySpace API.
type APIError struct {
	StatusCode int           `json:"-"`
	ODataCtx   string        `json:"@odata.context,omitempty"`
	Message    string        `json:"Message,omitempty"`
	Details    []ErrorDetail `json:"Details,omitempty"`
	Success    *bool         `json:"Success,omitempty"`
}

// ErrorDetail holds an individual validation error message.
type ErrorDetail struct {
	Message string `json:"Message"`
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("payspace: HTTP %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("payspace: HTTP %d", e.StatusCode)
}

// IsNotFound returns true if the error is a 404.
func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusNotFound
	}
	return false
}

// IsUnauthorized returns true if the error is a 401.
func IsUnauthorized(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusUnauthorized
	}
	return false
}

// IsRateLimited returns true if the error is a 429.
func IsRateLimited(err error) bool {
	if apiErr, ok := err.(*RateLimitError); ok {
		return apiErr.StatusCode == http.StatusTooManyRequests
	}
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusTooManyRequests
	}
	return false
}

// IsValidationError returns true if the error is a 400 with validation details.
func IsValidationError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == http.StatusBadRequest && len(apiErr.Details) > 0
	}
	return false
}

// RateLimitError is returned when the API rate limit is exceeded.
type RateLimitError struct {
	*APIError
	Limit     int
	Remaining int
	Reset     time.Time
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("payspace: rate limit exceeded (limit=%d, remaining=%d, reset=%s)",
		e.Limit, e.Remaining, e.Reset.Format(time.RFC3339))
}

// checkResponse inspects the HTTP response status and returns an appropriate error.
func checkResponse(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}

	body, _ := io.ReadAll(r.Body)

	apiErr := &APIError{StatusCode: r.StatusCode}
	if len(body) > 0 {
		_ = json.Unmarshal(body, apiErr)
	}

	if r.StatusCode == http.StatusTooManyRequests {
		rl := parseRateLimitHeaders(r.Header)
		return &RateLimitError{
			APIError:  apiErr,
			Limit:     rl.Limit,
			Remaining: rl.Remaining,
			Reset:     rl.Reset,
		}
	}

	return apiErr
}

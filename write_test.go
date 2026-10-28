package payspace

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// countingTransport records every round trip and answers with fn.
type countingTransport struct {
	calls  int32
	bodies []string
	fn     func(call int32) (*http.Response, error)
}

func (t *countingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	n := atomic.AddInt32(&t.calls, 1)
	body := ""
	if req.Body != nil {
		data, _ := io.ReadAll(req.Body)
		body = string(data)
	}
	t.bodies = append(t.bodies, body)
	return t.fn(n)
}

// closableBody fails reads once closed, like a real response body.
type closableBody struct {
	*strings.Reader
	closed bool
}

func (b *closableBody) Read(p []byte) (int, error) {
	if b.closed {
		return 0, errors.New("read on closed body")
	}
	return b.Reader.Read(p)
}

func (b *closableBody) Close() error {
	b.closed = true
	return nil
}

func statusResponse(status int, body string) *http.Response {
	return &http.Response{StatusCode: status, Header: http.Header{}, Body: &closableBody{Reader: strings.NewReader(body)}}
}

func newRetryTransport(base http.RoundTripper) *retryTransport {
	return &retryTransport{
		base:       base,
		logger:     testNopLogger(),
		maxRetries: 3,
		backoff:    time.Millisecond,
		auth:       &tokenManager{token: "t", expiresAt: time.Now().Add(time.Hour), logger: testNopLogger()},
	}
}

func newJSONRequest(t *testing.T, method, body string) *http.Request {
	t.Helper()
	req, err := http.NewRequestWithContext(context.Background(), method, "http://payspace.test/x", strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func TestRetryTransport_WriteServerError_NotRetried(t *testing.T) {
	for _, method := range []string{http.MethodPost, http.MethodPatch} {
		t.Run(method, func(t *testing.T) {
			base := &countingTransport{fn: func(int32) (*http.Response, error) {
				return statusResponse(http.StatusInternalServerError, `{"Message":"boom"}`), nil
			}}

			resp, err := newRetryTransport(base).RoundTrip(newJSONRequest(t, method, `{"a":1}`))

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			body, _ := io.ReadAll(resp.Body)
			if base.calls != 1 || resp.StatusCode != http.StatusInternalServerError || string(body) != `{"Message":"boom"}` {
				t.Errorf("calls=%d status=%d body=%q, want one call returning the 500 body", base.calls, resp.StatusCode, body)
			}
		})
	}
}

func TestRetryTransport_WriteNetworkError_NotRetried(t *testing.T) {
	networkErr := errors.New("connection reset")
	base := &countingTransport{fn: func(int32) (*http.Response, error) { return nil, networkErr }}

	_, err := newRetryTransport(base).RoundTrip(newJSONRequest(t, http.MethodPost, `{"a":1}`))

	if !errors.Is(err, networkErr) || base.calls != 1 {
		t.Errorf("calls=%d err=%v, want one call returning the network error", base.calls, err)
	}
}

func TestRetryTransport_ReadServerErrorExhausted_ReturnsReadableBody(t *testing.T) {
	base := &countingTransport{fn: func(int32) (*http.Response, error) {
		return statusResponse(http.StatusBadGateway, `{"Message":"down"}`), nil
	}}

	resp, err := newRetryTransport(base).RoundTrip(newJSONRequest(t, http.MethodGet, ""))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	body, readErr := io.ReadAll(resp.Body)
	if base.calls != 4 || readErr != nil || string(body) != `{"Message":"down"}` {
		t.Errorf("calls=%d body=%q readErr=%v, want 4 calls and the last body readable", base.calls, body, readErr)
	}
}

func TestRetryTransport_WriteUnauthorized_ResendsBody(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /connect/token", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"access_token": "fresh", "expires_in": 3600, "token_type": "Bearer"})
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	base := &countingTransport{fn: func(call int32) (*http.Response, error) {
		if call == 1 {
			return statusResponse(http.StatusUnauthorized, ""), nil
		}
		return statusResponse(http.StatusCreated, `{}`), nil
	}}
	rt := newRetryTransport(base)
	rt.auth = &tokenManager{token: "t", expiresAt: time.Now().Add(time.Hour), logger: testNopLogger(),
		identityURL: srv.URL, clientID: "id", clientSecret: "secret", httpClient: srv.Client()}

	_, err := rt.RoundTrip(newJSONRequest(t, http.MethodPost, `{"a":1}`))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(base.bodies) != 2 || base.bodies[1] != `{"a":1}` {
		t.Errorf("bodies=%q, want the body sent again after the 401", base.bodies)
	}
}

func TestRateLimitTransport_TooManyRequests_ResendsBody(t *testing.T) {
	base := &countingTransport{fn: func(call int32) (*http.Response, error) {
		if call == 1 {
			resp := statusResponse(http.StatusTooManyRequests, "")
			resp.Header.Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(time.Second).Unix(), 10))
			return resp, nil
		}
		return statusResponse(http.StatusCreated, `{}`), nil
	}}
	rt := &rateLimitTransport{base: base, logger: testNopLogger(), tracker: &rateLimitTracker{}, wait: true}

	_, err := rt.RoundTrip(newJSONRequest(t, http.MethodPost, `{"a":1}`))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(base.bodies) != 2 || base.bodies[1] != `{"a":1}` {
		t.Errorf("bodies=%q, want the body sent again after the 429", base.bodies)
	}
}

func TestClient_CreateODataRaw_PostsBodyAndReturnsEntity(t *testing.T) {
	var gotMethod, gotPath, gotBody string
	mux := http.NewServeMux()
	mux.HandleFunc("/odata/v2.0/12/EmployeeDependant", func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		gotMethod, gotPath, gotBody = r.Method, r.URL.Path, string(data)
		writeJSON(w, http.StatusCreated, map[string]any{"EmployeeDependantId": 7, "FirstName": "Ann", "Extra": []any{}})
	})
	_, client := testServerAndClient(t, mux)

	entity, _, err := client.CreateODataRaw(context.Background(), 12, "EmployeeDependant", map[string]any{"FirstName": "Ann"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPost || gotPath != "/odata/v2.0/12/EmployeeDependant" || gotBody != `{"FirstName":"Ann"}` {
		t.Errorf("request = %s %s %s", gotMethod, gotPath, gotBody)
	}
	if entity["EmployeeDependantId"] != float64(7) || entity["FirstName"] != "Ann" {
		t.Errorf("entity = %v, want the created entity", entity)
	}
}

func TestClient_UpdateODataRaw_PatchesKeyedEntity(t *testing.T) {
	var gotMethod, gotPath, gotBody string
	mux := http.NewServeMux()
	mux.HandleFunc("/odata/v2.0/12/EmployeeBankDetail(99)", func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		gotMethod, gotPath, gotBody = r.Method, r.URL.Path, string(data)
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	_, err := client.UpdateODataRaw(context.Background(), 12, "EmployeeBankDetail", 99, map[string]any{"Comments": "x"})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMethod != http.MethodPatch || gotPath != "/odata/v2.0/12/EmployeeBankDetail(99)" || gotBody != `{"Comments":"x"}` {
		t.Errorf("request = %s %s %s", gotMethod, gotPath, gotBody)
	}
}

func TestClient_GetODataRaw_ReturnsKeyedEntity(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/12/Employee(5)", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{"EmployeeId": 5, "CustomFields": []any{}})
	})
	_, client := testServerAndClient(t, mux)

	entity, _, err := client.GetODataRaw(context.Background(), 12, "Employee", 5)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if entity["EmployeeId"] != float64(5) {
		t.Errorf("entity = %v", entity)
	}
}

func TestClient_CreateODataRaw_ValidationError(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/odata/v2.0/12/Employee", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusBadRequest, map[string]any{"Message": "'LastName' must not be empty.",
			"Details": []map[string]any{{"Message": "'LastName' must not be empty."}}, "Success": false})
	})
	_, client := testServerAndClient(t, mux)

	_, _, err := client.CreateODataRaw(context.Background(), 12, "Employee", map[string]any{})

	var apiErr *APIError
	if !errors.As(err, &apiErr) || len(apiErr.Details) != 1 || apiErr.Details[0].Message != "'LastName' must not be empty." {
		t.Errorf("err = %v, want an APIError with the validation detail", err)
	}
}

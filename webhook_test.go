package payspace

import (
	"context"
	"net/http"
	"testing"
)

func TestWebhookService_ListErrors(t *testing.T) {
	mux := http.NewServeMux()
	// WebhookError uses apiURL: {base}/v2.0/{companyID}/WebhookError/{entityType}
	mux.HandleFunc("GET /v2.0/1/WebhookError/Employee", func(w http.ResponseWriter, r *http.Request) {
		// Verify query parameters.
		q := r.URL.Query()
		if got := q.Get("from"); got != "2024-01-01" {
			t.Errorf("expected from='2024-01-01', got %q", got)
		}
		if got := q.Get("to"); got != "2024-06-30" {
			t.Errorf("expected to='2024-06-30', got %q", got)
		}
		if got := q.Get("pageNumber"); got != "1" {
			t.Errorf("expected pageNumber='1', got %q", got)
		}
		if got := q.Get("pageSize"); got != "25" {
			t.Errorf("expected pageSize='25', got %q", got)
		}

		writeJSON(w, http.StatusOK, WebhookErrorResponse{
			Data: []WebhookError{
				{EntityType: "Employee", ErrorDate: "2024-03-15", Message: "Validation failed", Payload: `{"id":1}`},
				{EntityType: "Employee", ErrorDate: "2024-04-20", Message: "Timeout", Payload: `{"id":2}`},
			},
			PageNumber: 1,
			PageSize:   25,
			TotalCount: 2,
		})
	})
	_, client := testServerAndClient(t, mux)

	result, resp, err := client.Webhooks.ListErrors(context.Background(), 1, "Employee", "2024-01-01", "2024-06-30", 1, 25)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if result.TotalCount != 2 {
		t.Errorf("expected TotalCount=2, got %d", result.TotalCount)
	}
	if result.PageNumber != 1 {
		t.Errorf("expected PageNumber=1, got %d", result.PageNumber)
	}
	if result.PageSize != 25 {
		t.Errorf("expected PageSize=25, got %d", result.PageSize)
	}
	if len(result.Data) != 2 {
		t.Fatalf("expected 2 errors, got %d", len(result.Data))
	}
	if result.Data[0].Message != "Validation failed" {
		t.Errorf("expected Message='Validation failed', got %q", result.Data[0].Message)
	}
	if result.Data[1].ErrorDate != "2024-04-20" {
		t.Errorf("expected ErrorDate='2024-04-20', got %q", result.Data[1].ErrorDate)
	}
}

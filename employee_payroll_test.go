package payspace

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestEmployeePayrollService_ListPayRates(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeePayRate", func(w http.ResponseWriter, r *http.Request) {
		salary := 50000.0
		writeODataList(w, []EmployeePayRate{
			{PayRateId: 1, EmployeeNumber: "EMP001", BasicSalary: &salary, PayType: "Monthly"},
			{PayRateId: 2, EmployeeNumber: "EMP002", BasicSalary: &salary, PayType: "Hourly"},
		})
	})
	_, client := testServerAndClient(t, mux)

	rates, resp, err := client.EmployeePayroll.ListPayRates(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(rates) != 2 {
		t.Fatalf("expected 2 pay rates, got %d", len(rates))
	}
	if rates[0].PayType != "Monthly" {
		t.Errorf("expected PayType=Monthly, got %s", rates[0].PayType)
	}
}

func TestEmployeePayrollService_GetPayRate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeePayRate(42)", func(w http.ResponseWriter, r *http.Request) {
		salary := 75000.0
		writeJSON(w, http.StatusOK, EmployeePayRate{
			PayRateId:      42,
			EmployeeNumber: "EMP001",
			BasicSalary:    &salary,
		})
	})
	_, client := testServerAndClient(t, mux)

	rate, resp, err := client.EmployeePayroll.GetPayRate(context.Background(), 1, 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if rate.PayRateId != 42 {
		t.Errorf("expected PayRateId=42, got %d", rate.PayRateId)
	}
	if rate.BasicSalary == nil || *rate.BasicSalary != 75000.0 {
		t.Errorf("expected BasicSalary=75000, got %v", rate.BasicSalary)
	}
}

func TestEmployeePayrollService_CreatePayRate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeePayRate", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeePayRate
		json.NewDecoder(r.Body).Decode(&body)
		if body.EmployeeNumber != "EMP001" {
			t.Errorf("expected EmployeeNumber=EMP001, got %s", body.EmployeeNumber)
		}
		body.PayRateId = 99
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	salary := 60000.0
	input := &EmployeePayRate{EmployeeNumber: "EMP001", BasicSalary: &salary}
	created, resp, err := client.EmployeePayroll.CreatePayRate(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.PayRateId != 99 {
		t.Errorf("expected PayRateId=99, got %d", created.PayRateId)
	}
}

func TestEmployeePayrollService_UpdatePayRate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/EmployeePayRate(42)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	salary := 80000.0
	resp, err := client.EmployeePayroll.UpdatePayRate(context.Background(), 1, 42, &EmployeePayRate{BasicSalary: &salary})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestEmployeePayrollService_DeletePayRate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EmployeePayRate(42)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.EmployeePayroll.DeletePayRate(context.Background(), 1, 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestEmployeePayrollService_ListRecurring(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeComponent", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "1" {
			t.Errorf("expected period=1, got %s", r.URL.Query().Get("period"))
		}
		writeODataList(w, []map[string]any{
			{"ComponentId": 1, "ComponentCode": "ALLOW1", "Description": "Travel Allowance"},
		})
	})
	_, client := testServerAndClient(t, mux)

	items, resp, err := client.EmployeePayroll.ListRecurring(context.Background(), 1, "EmployeeComponent", "Monthly", "1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}
	if items[0]["Description"] != "Travel Allowance" {
		t.Errorf("expected Description=Travel Allowance, got %v", items[0]["Description"])
	}
}

func TestEmployeePayrollService_CreateRecurring(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeeComponent", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "2" {
			t.Errorf("expected period=2, got %s", r.URL.Query().Get("period"))
		}
		var body map[string]any
		json.NewDecoder(r.Body).Decode(&body)
		body["ComponentId"] = 50
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	input := map[string]any{"ComponentCode": "BONUS1", "Amount": 1000.0}
	created, resp, err := client.EmployeePayroll.CreateRecurring(context.Background(), 1, "EmployeeComponent", "Monthly", "2", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created["ComponentCode"] != "BONUS1" {
		t.Errorf("expected ComponentCode=BONUS1, got %v", created["ComponentCode"])
	}
}

func TestEmployeePayrollService_DeleteRecurring(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EmployeeComponent(10)", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.EmployeePayroll.DeleteRecurring(context.Background(), 1, "EmployeeComponent", 10, "Monthly", "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestEmployeePayrollService_ExtractComponents(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeeComponent/extract/start", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "3" {
			t.Errorf("expected period=3, got %s", r.URL.Query().Get("period"))
		}
		writeJSON(w, http.StatusCreated, ExtractionRequest{
			ExtractionId: "ext-123",
			Status:       "Started",
		})
	})
	_, client := testServerAndClient(t, mux)

	extraction, resp, err := client.EmployeePayroll.ExtractComponents(context.Background(), 1, "Monthly", "3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if extraction.ExtractionId != "ext-123" {
		t.Errorf("expected ExtractionId=ext-123, got %s", extraction.ExtractionId)
	}
	if extraction.Status != "Started" {
		t.Errorf("expected Status=Started, got %s", extraction.Status)
	}
}

func TestEmployeePayrollService_ExtractComponentsStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeComponent/extract/ext-123/status", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "3" {
			t.Errorf("expected period=3, got %s", r.URL.Query().Get("period"))
		}
		writeJSON(w, http.StatusOK, ExtractionStatus{
			ExtractionId: "ext-123",
			Status:       "Completed",
			DownloadUrl:  "https://example.com/download",
		})
	})
	_, client := testServerAndClient(t, mux)

	status, resp, err := client.EmployeePayroll.ExtractComponentsStatus(context.Background(), 1, "ext-123", "Monthly", "3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if status.Status != "Completed" {
		t.Errorf("expected Status=Completed, got %s", status.Status)
	}
	if status.DownloadUrl != "https://example.com/download" {
		t.Errorf("expected DownloadUrl, got %s", status.DownloadUrl)
	}
}

package payspace

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestCompanyPayrollService_ListFrequencies(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/CompanyFrequency", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CompanyFrequency{
			{FrequencyId: 1, Description: "Monthly", FrequencyCode: "M", Periods: 12},
			{FrequencyId: 2, Description: "Weekly", FrequencyCode: "W", Periods: 52},
		})
	})
	_, client := testServerAndClient(t, mux)

	freqs, resp, err := client.CompanyPayroll.ListFrequencies(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(freqs) != 2 {
		t.Fatalf("expected 2 frequencies, got %d", len(freqs))
	}
	if freqs[0].Description != "Monthly" {
		t.Errorf("expected Description='Monthly', got %q", freqs[0].Description)
	}
	if freqs[1].Periods != 52 {
		t.Errorf("expected Periods=52, got %d", freqs[1].Periods)
	}
}

func TestCompanyPayrollService_ListRuns(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/CompanyRun", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CompanyRun{
			{RunId: 100, FrequencyId: 1, Period: 1, Year: 2024, Status: "Open"},
			{RunId: 101, FrequencyId: 1, Period: 2, Year: 2024, Status: "Closed"},
		})
	})
	_, client := testServerAndClient(t, mux)

	runs, resp, err := client.CompanyPayroll.ListRuns(context.Background(), 1, 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(runs) != 2 {
		t.Fatalf("expected 2 runs, got %d", len(runs))
	}
	if runs[0].Status != "Open" {
		t.Errorf("expected Status='Open', got %q", runs[0].Status)
	}
}

func TestCompanyPayrollService_ListRunsByFrequency(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/CompanyRun", func(w http.ResponseWriter, r *http.Request) {
		freq := r.URL.Query().Get("frequency")
		if freq != "2" {
			t.Errorf("expected frequency query param '2', got %q", freq)
		}
		writeODataList(w, []CompanyRun{
			{RunId: 200, FrequencyId: 2, Period: 1, Year: 2024, Status: "Open"},
		})
	})
	_, client := testServerAndClient(t, mux)

	runs, resp, err := client.CompanyPayroll.ListRuns(context.Background(), 1, 2, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(runs) != 1 {
		t.Fatalf("expected 1 run, got %d", len(runs))
	}
	if runs[0].FrequencyId != 2 {
		t.Errorf("expected FrequencyId=2, got %d", runs[0].FrequencyId)
	}
}

func TestCompanyPayrollService_ListComponents(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/ComponentCompany", func(w http.ResponseWriter, r *http.Request) {
		freq := r.URL.Query().Get("frequency")
		period := r.URL.Query().Get("period")
		if freq != "1" {
			t.Errorf("expected frequency=1, got %q", freq)
		}
		if period != "3" {
			t.Errorf("expected period=3, got %q", period)
		}
		writeODataList(w, []ComponentCompany{
			{ComponentCompanyId: 50, ComponentCode: "BASIC", Description: "Basic Salary", ComponentType: "Earning"},
			{ComponentCompanyId: 51, ComponentCode: "TAX", Description: "Income Tax", ComponentType: "Deduction"},
		})
	})
	_, client := testServerAndClient(t, mux)

	components, resp, err := client.CompanyPayroll.ListComponents(context.Background(), 1, 1, 3, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(components) != 2 {
		t.Fatalf("expected 2 components, got %d", len(components))
	}
	if components[0].ComponentCode != "BASIC" {
		t.Errorf("expected ComponentCode='BASIC', got %q", components[0].ComponentCode)
	}
	if components[1].ComponentType != "Deduction" {
		t.Errorf("expected ComponentType='Deduction', got %q", components[1].ComponentType)
	}
}

func TestCompanyPayrollService_UpdateComponent(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/ComponentCompany(50)", func(w http.ResponseWriter, r *http.Request) {
		freq := r.URL.Query().Get("frequency")
		period := r.URL.Query().Get("period")
		if freq != "1" {
			t.Errorf("expected frequency=1, got %q", freq)
		}
		if period != "3" {
			t.Errorf("expected period=3, got %q", period)
		}

		body, _ := io.ReadAll(r.Body)
		var input ComponentCompany
		if err := json.Unmarshal(body, &input); err != nil {
			t.Errorf("failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if input.Description != "Updated Basic Salary" {
			t.Errorf("expected Description='Updated Basic Salary', got %q", input.Description)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	input := &ComponentCompany{Description: "Updated Basic Salary"}
	resp, err := client.CompanyPayroll.UpdateComponent(context.Background(), 1, 50, 1, 3, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", resp.StatusCode)
	}
}

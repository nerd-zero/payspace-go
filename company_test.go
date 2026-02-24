package payspace

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestCompanyService_ListDocuments(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CompanyDocument/2024/6", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CompanyDocument{
			{DocumentId: 10, Year: 2024, Month: 6},
			{DocumentId: 11, Year: 2024, Month: 6},
		})
	})
	_, client := testServerAndClient(t, mux)

	docs, resp, err := client.Companies.ListDocuments(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(docs) != 2 {
		t.Fatalf("expected 2 documents, got %d", len(docs))
	}
	if docs[0].DocumentId != 10 {
		t.Errorf("expected DocumentId=10, got %d", docs[0].DocumentId)
	}
	if docs[1].DocumentId != 11 {
		t.Errorf("expected DocumentId=11, got %d", docs[1].DocumentId)
	}
}

func TestCompanyService_GetDocument(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CompanyDocument(42)", func(w http.ResponseWriter, r *http.Request) {
		total := 1500.50
		writeJSON(w, http.StatusOK, CompanyDocument{
			DocumentId:    42,
			Year:          2024,
			Month:         3,
			DocumentTotal: &total,
		})
	})
	_, client := testServerAndClient(t, mux)

	doc, resp, err := client.Companies.GetDocument(context.Background(), 1, 42)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if doc.DocumentId != 42 {
		t.Errorf("expected DocumentId=42, got %d", doc.DocumentId)
	}
	if doc.DocumentTotal == nil || *doc.DocumentTotal != 1500.50 {
		t.Errorf("expected DocumentTotal=1500.50, got %v", doc.DocumentTotal)
	}
}

func TestCompanyService_ListDocumentLines(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CompanyDocumentLine/2024/6", func(w http.ResponseWriter, r *http.Request) {
		amt := 250.00
		writeODataList(w, []CompanyDocumentLine{
			{DocumentLineId: 100, DocumentId: 10, Description: "License Fee", Amount: &amt},
		})
	})
	_, client := testServerAndClient(t, mux)

	lines, resp, err := client.Companies.ListDocumentLines(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(lines) != 1 {
		t.Fatalf("expected 1 line, got %d", len(lines))
	}
	if lines[0].DocumentLineId != 100 {
		t.Errorf("expected DocumentLineId=100, got %d", lines[0].DocumentLineId)
	}
	if lines[0].Description != "License Fee" {
		t.Errorf("expected Description='License Fee', got %q", lines[0].Description)
	}
}

func TestCompanyService_ListCustomForms(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CompanyCustomForm", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CompanyCustomForm{
			{CustomFormId: 1, FormCategory: "HR", FormDescription: "Onboarding"},
			{CustomFormId: 2, FormCategory: "Payroll", FormDescription: "Tax Form"},
		})
	})
	_, client := testServerAndClient(t, mux)

	forms, resp, err := client.Companies.ListCustomForms(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(forms) != 2 {
		t.Fatalf("expected 2 forms, got %d", len(forms))
	}
	if forms[0].FormCategory != "HR" {
		t.Errorf("expected FormCategory='HR', got %q", forms[0].FormCategory)
	}
}

func TestCompanyService_GetCustomForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CompanyCustomForm(5)", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, CompanyCustomForm{
			CustomFormId:    5,
			FormCategory:    "HR",
			FormDescription: "Exit Interview",
		})
	})
	_, client := testServerAndClient(t, mux)

	form, resp, err := client.Companies.GetCustomForm(context.Background(), 1, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if form.CustomFormId != 5 {
		t.Errorf("expected CustomFormId=5, got %d", form.CustomFormId)
	}
	if form.FormDescription != "Exit Interview" {
		t.Errorf("expected FormDescription='Exit Interview', got %q", form.FormDescription)
	}
}

func TestCompanyService_CreateCustomForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/CompanyCustomForm", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var input CompanyCustomForm
		if err := json.Unmarshal(body, &input); err != nil {
			t.Errorf("failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if input.FormCategory != "Benefits" {
			t.Errorf("expected FormCategory='Benefits', got %q", input.FormCategory)
		}
		writeJSON(w, http.StatusCreated, CompanyCustomForm{
			CustomFormId:    99,
			FormCategory:    input.FormCategory,
			FormDescription: input.FormDescription,
		})
	})
	_, client := testServerAndClient(t, mux)

	input := &CompanyCustomForm{
		FormCategory:    "Benefits",
		FormDescription: "Health Insurance",
	}
	form, resp, err := client.Companies.CreateCustomForm(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}
	if form.CustomFormId != 99 {
		t.Errorf("expected CustomFormId=99, got %d", form.CustomFormId)
	}
	if form.FormDescription != "Health Insurance" {
		t.Errorf("expected FormDescription='Health Insurance', got %q", form.FormDescription)
	}
}

func TestCompanyService_UpdateCustomForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/CompanyCustomForm(5)", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var input CompanyCustomForm
		if err := json.Unmarshal(body, &input); err != nil {
			t.Errorf("failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if input.FormDescription != "Updated Description" {
			t.Errorf("expected FormDescription='Updated Description', got %q", input.FormDescription)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	input := &CompanyCustomForm{FormDescription: "Updated Description"}
	resp, err := client.Companies.UpdateCustomForm(context.Background(), 1, 5, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", resp.StatusCode)
	}
}

func TestCompanyService_DeleteCustomForm(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/CompanyCustomForm(5)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.Companies.DeleteCustomForm(context.Background(), 1, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", resp.StatusCode)
	}
}

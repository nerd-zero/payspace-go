package payspace

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestEmployeePayslipService_ListPayslips(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeePayslip/2024/6", func(w http.ResponseWriter, r *http.Request) {
		gross := 10000.0
		net := 7500.0
		writeODataList(w, []EmployeePayslip{
			{PayslipId: 1, EmployeeNumber: "EMP001", GrossPay: &gross, NetPay: &net, Year: 2024, Month: 6},
		})
	})
	_, client := testServerAndClient(t, mux)

	payslips, resp, err := client.EmployeePayslips.ListPayslips(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(payslips) != 1 {
		t.Fatalf("expected 1 payslip, got %d", len(payslips))
	}
	if payslips[0].EmployeeNumber != "EMP001" {
		t.Errorf("expected EmployeeNumber=EMP001, got %s", payslips[0].EmployeeNumber)
	}
	if payslips[0].GrossPay == nil || *payslips[0].GrossPay != 10000.0 {
		t.Errorf("expected GrossPay=10000, got %v", payslips[0].GrossPay)
	}
}

func TestEmployeePayslipService_ListPayslipLines(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeePayslipLine/2024/6", func(w http.ResponseWriter, r *http.Request) {
		amount := 5000.0
		writeODataList(w, []EmployeePayslipLine{
			{PayslipLineId: 1, EmployeeNumber: "EMP001", ComponentCode: "BASIC", Amount: &amount},
			{PayslipLineId: 2, EmployeeNumber: "EMP001", ComponentCode: "TAX", Amount: &amount},
		})
	})
	_, client := testServerAndClient(t, mux)

	lines, resp, err := client.EmployeePayslips.ListPayslipLines(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(lines) != 2 {
		t.Fatalf("expected 2 payslip lines, got %d", len(lines))
	}
	if lines[0].ComponentCode != "BASIC" {
		t.Errorf("expected ComponentCode=BASIC, got %s", lines[0].ComponentCode)
	}
}

func TestEmployeePayslipService_ListConsolidatedPayslips(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeePayslip/2024/6/consolidated", func(w http.ResponseWriter, r *http.Request) {
		gross := 20000.0
		writeODataList(w, []EmployeePayslip{
			{PayslipId: 1, EmployeeNumber: "EMP001", GrossPay: &gross},
		})
	})
	_, client := testServerAndClient(t, mux)

	payslips, resp, err := client.EmployeePayslips.ListConsolidatedPayslips(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(payslips) != 1 {
		t.Fatalf("expected 1 consolidated payslip, got %d", len(payslips))
	}
}

func TestEmployeePayslipService_ListPayslipPDFs(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeePayslip/2024/6/pdf", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []EmployeePayslipPdf{
			{PayslipPdfId: 1, EmployeeNumber: "EMP001", PdfContent: "base64data", Year: 2024, Month: 6},
		})
	})
	_, client := testServerAndClient(t, mux)

	pdfs, resp, err := client.EmployeePayslips.ListPayslipPDFs(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(pdfs) != 1 {
		t.Fatalf("expected 1 PDF, got %d", len(pdfs))
	}
	if pdfs[0].PdfContent != "base64data" {
		t.Errorf("expected PdfContent=base64data, got %s", pdfs[0].PdfContent)
	}
}

func TestEmployeePayslipService_ListEditPayslips(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EditPayslip", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "1" {
			t.Errorf("expected period=1, got %s", r.URL.Query().Get("period"))
		}
		amount := 1000.0
		writeODataList(w, []EditPayslip{
			{EditPayslipId: 1, EmployeeNumber: "EMP001", ComponentCode: "BONUS", Amount: &amount},
		})
	})
	_, client := testServerAndClient(t, mux)

	edits, resp, err := client.EmployeePayslips.ListEditPayslips(context.Background(), 1, "Monthly", "1", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(edits) != 1 {
		t.Fatalf("expected 1 edit payslip, got %d", len(edits))
	}
	if edits[0].ComponentCode != "BONUS" {
		t.Errorf("expected ComponentCode=BONUS, got %s", edits[0].ComponentCode)
	}
}

func TestEmployeePayslipService_CreateEditPayslip(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EditPayslip", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "1" {
			t.Errorf("expected period=1, got %s", r.URL.Query().Get("period"))
		}
		if r.URL.Query().Get("Action") != "Add" {
			t.Errorf("expected Action=Add, got %s", r.URL.Query().Get("Action"))
		}
		var body EditPayslip
		json.NewDecoder(r.Body).Decode(&body)
		if body.ComponentCode != "OVERTIME" {
			t.Errorf("expected ComponentCode=OVERTIME, got %s", body.ComponentCode)
		}
		body.EditPayslipId = 55
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	amount := 500.0
	input := &EditPayslip{EmployeeNumber: "EMP001", ComponentCode: "OVERTIME", Amount: &amount}
	created, resp, err := client.EmployeePayslips.CreateEditPayslip(context.Background(), 1, "Monthly", "1", "Add", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.EditPayslipId != 55 {
		t.Errorf("expected EditPayslipId=55, got %d", created.EditPayslipId)
	}
}

func TestEmployeePayslipService_BatchEditPayslips(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EditPayslip/Batch", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "1" {
			t.Errorf("expected period=1, got %s", r.URL.Query().Get("period"))
		}
		var items []EditPayslip
		json.NewDecoder(r.Body).Decode(&items)
		if len(items) != 2 {
			t.Errorf("expected 2 items in batch, got %d", len(items))
		}
		// Return the items with IDs assigned.
		for i := range items {
			items[i].EditPayslipId = i + 100
		}
		writeODataList(w, items)
	})
	_, client := testServerAndClient(t, mux)

	amount := 100.0
	items := []EditPayslip{
		{EmployeeNumber: "EMP001", ComponentCode: "BONUS1", Amount: &amount},
		{EmployeeNumber: "EMP002", ComponentCode: "BONUS2", Amount: &amount},
	}
	result, resp, err := client.EmployeePayslips.BatchEditPayslips(context.Background(), 1, "Monthly", "1", items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(result) != 2 {
		t.Fatalf("expected 2 results, got %d", len(result))
	}
}

func TestEmployeePayslipService_DeleteEditPayslip(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EditPayslip(55)", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "1" {
			t.Errorf("expected period=1, got %s", r.URL.Query().Get("period"))
		}
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.EmployeePayslips.DeleteEditPayslip(context.Background(), 1, 55, "Monthly", "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestEmployeePayslipService_ListTaxCertificates(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeTaxCertificate", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []EmployeeTaxCertificate{
			{TaxCertificateId: 1, EmployeeNumber: "EMP001", TaxYear: 2024, CertificateType: "IRP5"},
		})
	})
	_, client := testServerAndClient(t, mux)

	certs, resp, err := client.EmployeePayslips.ListTaxCertificates(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(certs) != 1 {
		t.Fatalf("expected 1 certificate, got %d", len(certs))
	}
	if certs[0].CertificateType != "IRP5" {
		t.Errorf("expected CertificateType=IRP5, got %s", certs[0].CertificateType)
	}
}

func TestEmployeePayslipService_ListLumpSums(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeLumpSum", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "3" {
			t.Errorf("expected period=3, got %s", r.URL.Query().Get("period"))
		}
		amount := 25000.0
		writeODataList(w, []EmployeeLumpSum{
			{LumpSumId: 1, EmployeeNumber: "EMP001", ComponentCode: "LUMP1", Amount: &amount},
		})
	})
	_, client := testServerAndClient(t, mux)

	lumpSums, resp, err := client.EmployeePayslips.ListLumpSums(context.Background(), 1, "Monthly", "3", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(lumpSums) != 1 {
		t.Fatalf("expected 1 lump sum, got %d", len(lumpSums))
	}
	if lumpSums[0].ComponentCode != "LUMP1" {
		t.Errorf("expected ComponentCode=LUMP1, got %s", lumpSums[0].ComponentCode)
	}
}

func TestEmployeePayslipService_CreateLumpSum(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeeLumpSum", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("frequency") != "Monthly" {
			t.Errorf("expected frequency=Monthly, got %s", r.URL.Query().Get("frequency"))
		}
		if r.URL.Query().Get("period") != "3" {
			t.Errorf("expected period=3, got %s", r.URL.Query().Get("period"))
		}
		var body EmployeeLumpSum
		json.NewDecoder(r.Body).Decode(&body)
		if body.ComponentCode != "LUMP2" {
			t.Errorf("expected ComponentCode=LUMP2, got %s", body.ComponentCode)
		}
		body.LumpSumId = 99
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	amount := 50000.0
	input := &EmployeeLumpSum{EmployeeNumber: "EMP001", ComponentCode: "LUMP2", Amount: &amount}
	created, resp, err := client.EmployeePayslips.CreateLumpSum(context.Background(), 1, "Monthly", "3", input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.LumpSumId != 99 {
		t.Errorf("expected LumpSumId=99, got %d", created.LumpSumId)
	}
}

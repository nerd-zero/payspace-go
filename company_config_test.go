package payspace

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestCompanyConfigService_ListOrganizationPositions(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/OrganizationPositionDetail", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []OrganizationPositionDetail{
			{OrganizationPositionDetailId: 1, PositionCode: "DEV01", PositionDescription: "Developer", Department: "Engineering"},
			{OrganizationPositionDetailId: 2, PositionCode: "MGR01", PositionDescription: "Manager", Department: "Operations"},
		})
	})
	_, client := testServerAndClient(t, mux)

	positions, resp, err := client.CompanyConfig.ListOrganizationPositions(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(positions) != 2 {
		t.Fatalf("expected 2 positions, got %d", len(positions))
	}
	if positions[0].PositionCode != "DEV01" {
		t.Errorf("expected PositionCode='DEV01', got %q", positions[0].PositionCode)
	}
	if positions[1].Department != "Operations" {
		t.Errorf("expected Department='Operations', got %q", positions[1].Department)
	}
}

func TestCompanyConfigService_GetOrganizationPosition(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/OrganizationPositionDetail(7)", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, OrganizationPositionDetail{
			OrganizationPositionDetailId: 7,
			PositionCode:                 "HR01",
			PositionDescription:          "HR Specialist",
			Department:                   "Human Resources",
			EffectiveDate:                "2024-01-01",
		})
	})
	_, client := testServerAndClient(t, mux)

	pos, resp, err := client.CompanyConfig.GetOrganizationPosition(context.Background(), 1, 7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if pos.OrganizationPositionDetailId != 7 {
		t.Errorf("expected OrganizationPositionDetailId=7, got %d", pos.OrganizationPositionDetailId)
	}
	if pos.PositionDescription != "HR Specialist" {
		t.Errorf("expected PositionDescription='HR Specialist', got %q", pos.PositionDescription)
	}
	if pos.EffectiveDate != "2024-01-01" {
		t.Errorf("expected EffectiveDate='2024-01-01', got %q", pos.EffectiveDate)
	}
}

func TestCompanyConfigService_CreateOrganizationPosition(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/OrganizationPositionDetail", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var input OrganizationPositionDetail
		if err := json.Unmarshal(body, &input); err != nil {
			t.Errorf("failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if input.PositionCode != "FIN01" {
			t.Errorf("expected PositionCode='FIN01', got %q", input.PositionCode)
		}
		writeJSON(w, http.StatusCreated, OrganizationPositionDetail{
			OrganizationPositionDetailId: 20,
			PositionCode:                 input.PositionCode,
			PositionDescription:          input.PositionDescription,
			Department:                   input.Department,
		})
	})
	_, client := testServerAndClient(t, mux)

	input := &OrganizationPositionDetail{
		PositionCode:        "FIN01",
		PositionDescription: "Financial Analyst",
		Department:          "Finance",
	}
	pos, resp, err := client.CompanyConfig.CreateOrganizationPosition(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}
	if pos.OrganizationPositionDetailId != 20 {
		t.Errorf("expected OrganizationPositionDetailId=20, got %d", pos.OrganizationPositionDetailId)
	}
	if pos.Department != "Finance" {
		t.Errorf("expected Department='Finance', got %q", pos.Department)
	}
}

func TestCompanyConfigService_UpdateOrganizationPosition(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/OrganizationPositionDetail(7)", func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var input OrganizationPositionDetail
		if err := json.Unmarshal(body, &input); err != nil {
			t.Errorf("failed to decode request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if input.PositionDescription != "Senior HR Specialist" {
			t.Errorf("expected PositionDescription='Senior HR Specialist', got %q", input.PositionDescription)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	input := &OrganizationPositionDetail{PositionDescription: "Senior HR Specialist"}
	resp, err := client.CompanyConfig.UpdateOrganizationPosition(context.Background(), 1, 7, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", resp.StatusCode)
	}
}

func TestCompanyConfigService_DeleteOrganizationPosition(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/OrganizationPositionDetail(7)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.CompanyConfig.DeleteOrganizationPosition(context.Background(), 1, 7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf("expected status 204, got %d", resp.StatusCode)
	}
}

func TestCompanyConfigService_ListRegions(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/CompanyRegion", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CompanyRegion{
			{RegionId: 1, Description: "Gauteng", RegionCode: "GP"},
			{RegionId: 2, Description: "Western Cape", RegionCode: "WC"},
		})
	})
	_, client := testServerAndClient(t, mux)

	regions, resp, err := client.CompanyConfig.ListRegions(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(regions) != 2 {
		t.Fatalf("expected 2 regions, got %d", len(regions))
	}
	if regions[0].RegionCode != "GP" {
		t.Errorf("expected RegionCode='GP', got %q", regions[0].RegionCode)
	}
	if regions[1].Description != "Western Cape" {
		t.Errorf("expected Description='Western Cape', got %q", regions[1].Description)
	}
}

func TestCompanyConfigService_ListGeneralLedger(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/GeneralLedger/2024/6", func(w http.ResponseWriter, r *http.Request) {
		amt := 50000.00
		writeODataList(w, []GeneralLedger{
			{GeneralLedgerId: 1, AccountCode: "4000", Description: "Salaries", Amount: &amt},
		})
	})
	_, client := testServerAndClient(t, mux)

	entries, resp, err := client.CompanyConfig.ListGeneralLedger(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(entries))
	}
	if entries[0].AccountCode != "4000" {
		t.Errorf("expected AccountCode='4000', got %q", entries[0].AccountCode)
	}
	if entries[0].Amount == nil || *entries[0].Amount != 50000.00 {
		t.Errorf("expected Amount=50000.00, got %v", entries[0].Amount)
	}
}

func TestCompanyConfigService_ListEFTOutbox(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EFTOutbox/2024/6", func(w http.ResponseWriter, r *http.Request) {
		amt := 25000.00
		writeODataList(w, []EFTOutbox{
			{EftOutboxId: 1, EmployeeNumber: "EMP001", BankName: "FNB", AccountNumber: "123456789", Amount: &amt, Status: "Pending"},
		})
	})
	_, client := testServerAndClient(t, mux)

	records, resp, err := client.CompanyConfig.ListEFTOutbox(context.Background(), 1, 2024, 6, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(records) != 1 {
		t.Fatalf("expected 1 record, got %d", len(records))
	}
	if records[0].EmployeeNumber != "EMP001" {
		t.Errorf("expected EmployeeNumber='EMP001', got %q", records[0].EmployeeNumber)
	}
	if records[0].BankName != "FNB" {
		t.Errorf("expected BankName='FNB', got %q", records[0].BankName)
	}
	if records[0].Status != "Pending" {
		t.Errorf("expected Status='Pending', got %q", records[0].Status)
	}
}

func TestCompanyConfigService_ListCostingProjectActivities(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CostingProjectActivity", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CostingProjectActivity{
			{ProjectActivityId: 1, ProjectCode: "PROJ01", ActivityCode: "ACT01", Description: "Development"},
			{ProjectActivityId: 2, ProjectCode: "PROJ01", ActivityCode: "ACT02", Description: "Testing"},
		})
	})
	_, client := testServerAndClient(t, mux)

	activities, resp, err := client.CompanyConfig.ListCostingProjectActivities(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(activities) != 2 {
		t.Fatalf("expected 2 activities, got %d", len(activities))
	}
	if activities[0].ProjectCode != "PROJ01" {
		t.Errorf("expected ProjectCode='PROJ01', got %q", activities[0].ProjectCode)
	}
	if activities[1].Description != "Testing" {
		t.Errorf("expected Description='Testing', got %q", activities[1].Description)
	}
}

func TestCompanyConfigService_ListComponentTableDefinitions(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/ComponentTableDefinition", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []ComponentTableDefinition{
			{TableDefinitionId: 1, TableName: "TaxTable", Description: "Annual Tax Table"},
			{TableDefinitionId: 2, TableName: "BonusTable", Description: "Bonus Calculation Table"},
		})
	})
	_, client := testServerAndClient(t, mux)

	defs, resp, err := client.CompanyConfig.ListComponentTableDefinitions(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(defs) != 2 {
		t.Fatalf("expected 2 definitions, got %d", len(defs))
	}
	if defs[0].TableName != "TaxTable" {
		t.Errorf("expected TableName='TaxTable', got %q", defs[0].TableName)
	}
	if defs[1].Description != "Bonus Calculation Table" {
		t.Errorf("expected Description='Bonus Calculation Table', got %q", defs[1].Description)
	}
}

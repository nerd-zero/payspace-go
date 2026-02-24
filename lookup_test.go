package payspace

import (
	"context"
	"net/http"
	"testing"
)

func TestLookupService_ListLookup(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/CompanyLegalBody", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []LookupValue{
			{Id: 1, Code: "CC", Description: "Close Corporation"},
			{Id: 2, Code: "PTY", Description: "Proprietary Limited"},
		})
	})
	_, client := testServerAndClient(t, mux)

	values, resp, err := client.Lookups.ListLookup(context.Background(), 1, "CompanyLegalBody", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(values) != 2 {
		t.Fatalf("expected 2 values, got %d", len(values))
	}
	if values[0].Code != "CC" {
		t.Errorf("expected Code='CC', got %q", values[0].Code)
	}
	if values[1].Description != "Proprietary Limited" {
		t.Errorf("expected Description='Proprietary Limited', got %q", values[1].Description)
	}
}

func TestLookupService_ListActivityCodes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/ActivityCode", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []LookupValue{
			{Id: 10, Code: "DEV", Description: "Development"},
			{Id: 11, Code: "QA", Description: "Quality Assurance"},
		})
	})
	_, client := testServerAndClient(t, mux)

	values, resp, err := client.Lookups.ListActivityCodes(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(values) != 2 {
		t.Fatalf("expected 2 values, got %d", len(values))
	}
	if values[0].Code != "DEV" {
		t.Errorf("expected Code='DEV', got %q", values[0].Code)
	}
}

func TestLookupService_ListLeaveReasons(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/LeaveReason", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []LookupValue{
			{Id: 20, Code: "SICK", Description: "Sick Leave"},
			{Id: 21, Code: "FAM", Description: "Family Responsibility"},
		})
	})
	_, client := testServerAndClient(t, mux)

	values, resp, err := client.Lookups.ListLeaveReasons(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(values) != 2 {
		t.Fatalf("expected 2 values, got %d", len(values))
	}
	if values[0].Description != "Sick Leave" {
		t.Errorf("expected Description='Sick Leave', got %q", values[0].Description)
	}
}

func TestLookupService_ListLeaveTypes(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Lookup/LeaveType", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []LookupValue{
			{Id: 30, Code: "ANN", Description: "Annual Leave"},
			{Id: 31, Code: "MAT", Description: "Maternity Leave"},
		})
	})
	_, client := testServerAndClient(t, mux)

	values, resp, err := client.Lookups.ListLeaveTypes(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(values) != 2 {
		t.Fatalf("expected 2 values, got %d", len(values))
	}
	if values[1].Code != "MAT" {
		t.Errorf("expected Code='MAT', got %q", values[1].Code)
	}
}

func TestLookupService_ListCustomFields(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CustomField/Employee", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CustomField{
			{CustomFieldId: 100, EntityType: "Employee", FieldName: "ShoeSize", FieldType: "Number", Caption: "Shoe Size"},
			{CustomFieldId: 101, EntityType: "Employee", FieldName: "FavColor", FieldType: "Text", Caption: "Favourite Colour"},
		})
	})
	_, client := testServerAndClient(t, mux)

	fields, resp, err := client.Lookups.ListCustomFields(context.Background(), 1, "Employee", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
	if fields[0].FieldName != "ShoeSize" {
		t.Errorf("expected FieldName='ShoeSize', got %q", fields[0].FieldName)
	}
	if fields[1].FieldType != "Text" {
		t.Errorf("expected FieldType='Text', got %q", fields[1].FieldType)
	}
}

func TestLookupService_ListCustomFieldOptions(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CustomFieldLookupOption/101", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CustomFieldOption{
			{OptionId: 1, FieldId: 101, Code: "RED", Description: "Red"},
			{OptionId: 2, FieldId: 101, Code: "BLU", Description: "Blue"},
			{OptionId: 3, FieldId: 101, Code: "GRN", Description: "Green"},
		})
	})
	_, client := testServerAndClient(t, mux)

	options, resp, err := client.Lookups.ListCustomFieldOptions(context.Background(), 1, 101, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(options) != 3 {
		t.Fatalf("expected 3 options, got %d", len(options))
	}
	if options[0].Code != "RED" {
		t.Errorf("expected Code='RED', got %q", options[0].Code)
	}
	if options[2].Description != "Green" {
		t.Errorf("expected Description='Green', got %q", options[2].Description)
	}
}

func TestLookupService_ListCustomFormFields(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/CustomFormField/5", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []CustomField{
			{CustomFieldId: 200, FieldName: "StartDate", FieldType: "Date", Caption: "Start Date"},
			{CustomFieldId: 201, FieldName: "Notes", FieldType: "Text", Caption: "Additional Notes"},
		})
	})
	_, client := testServerAndClient(t, mux)

	fields, resp, err := client.Lookups.ListCustomFormFields(context.Background(), 1, 5, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if len(fields) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(fields))
	}
	if fields[0].Caption != "Start Date" {
		t.Errorf("expected Caption='Start Date', got %q", fields[0].Caption)
	}
	if fields[1].FieldType != "Text" {
		t.Errorf("expected FieldType='Text', got %q", fields[1].FieldType)
	}
}

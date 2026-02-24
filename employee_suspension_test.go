package payspace

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestEmployeeSuspensionService_List(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeSuspension", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []EmployeeSuspension{
			{SuspensionId: 1, EmployeeNumber: "EMP001", SuspensionReason: "Disciplinary", SuspensionStatus: "Active"},
			{SuspensionId: 2, EmployeeNumber: "EMP002", SuspensionReason: "Investigation", SuspensionStatus: "Completed"},
		})
	})
	_, client := testServerAndClient(t, mux)

	suspensions, resp, err := client.EmployeeSuspensions.List(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(suspensions) != 2 {
		t.Fatalf("expected 2 suspensions, got %d", len(suspensions))
	}
	if suspensions[0].SuspensionReason != "Disciplinary" {
		t.Errorf("expected SuspensionReason=Disciplinary, got %s", suspensions[0].SuspensionReason)
	}
	if suspensions[1].SuspensionStatus != "Completed" {
		t.Errorf("expected SuspensionStatus=Completed, got %s", suspensions[1].SuspensionStatus)
	}
}

func TestEmployeeSuspensionService_Get(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeSuspension(10)", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, EmployeeSuspension{
			SuspensionId:     10,
			EmployeeNumber:   "EMP001",
			SuspensionDate:   "2024-06-01",
			SuspensionReason: "Disciplinary",
			SuspensionStatus: "Active",
			ReturnDate:       "2024-07-01",
		})
	})
	_, client := testServerAndClient(t, mux)

	suspension, resp, err := client.EmployeeSuspensions.Get(context.Background(), 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if suspension.SuspensionId != 10 {
		t.Errorf("expected SuspensionId=10, got %d", suspension.SuspensionId)
	}
	if suspension.SuspensionDate != "2024-06-01" {
		t.Errorf("expected SuspensionDate=2024-06-01, got %s", suspension.SuspensionDate)
	}
	if suspension.ReturnDate != "2024-07-01" {
		t.Errorf("expected ReturnDate=2024-07-01, got %s", suspension.ReturnDate)
	}
}

func TestEmployeeSuspensionService_ListAll(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeSuspension/all", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []EmployeeSuspension{
			{SuspensionId: 1, EmployeeNumber: "EMP001", SuspensionStatus: "Active"},
			{SuspensionId: 2, EmployeeNumber: "EMP002", SuspensionStatus: "Completed"},
			{SuspensionId: 3, EmployeeNumber: "EMP003", SuspensionStatus: "Cancelled"},
		})
	})
	_, client := testServerAndClient(t, mux)

	suspensions, resp, err := client.EmployeeSuspensions.ListAll(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(suspensions) != 3 {
		t.Fatalf("expected 3 suspensions, got %d", len(suspensions))
	}
}

func TestEmployeeSuspensionService_Create(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeeSuspension", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeeSuspension
		json.NewDecoder(r.Body).Decode(&body)
		if body.EmployeeNumber != "EMP001" {
			t.Errorf("expected EmployeeNumber=EMP001, got %s", body.EmployeeNumber)
		}
		if body.SuspensionReason != "Investigation" {
			t.Errorf("expected SuspensionReason=Investigation, got %s", body.SuspensionReason)
		}
		body.SuspensionId = 50
		body.SuspensionStatus = "Active"
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	input := &EmployeeSuspension{
		EmployeeNumber:   "EMP001",
		SuspensionDate:   "2024-07-01",
		SuspensionReason: "Investigation",
		PaymentMethod:    "Full Pay",
	}
	created, resp, err := client.EmployeeSuspensions.Create(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.SuspensionId != 50 {
		t.Errorf("expected SuspensionId=50, got %d", created.SuspensionId)
	}
	if created.SuspensionStatus != "Active" {
		t.Errorf("expected SuspensionStatus=Active, got %s", created.SuspensionStatus)
	}
}

func TestEmployeeSuspensionService_Update(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/EmployeeSuspension(50)", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeeSuspension
		json.NewDecoder(r.Body).Decode(&body)
		if body.ReturnDate != "2024-08-01" {
			t.Errorf("expected ReturnDate=2024-08-01, got %s", body.ReturnDate)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.EmployeeSuspensions.Update(context.Background(), 1, 50, &EmployeeSuspension{
		ReturnDate:       "2024-08-01",
		SuspensionStatus: "Completed",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestEmployeeSuspensionService_Delete(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EmployeeSuspension(50)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.EmployeeSuspensions.Delete(context.Background(), 1, 50)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

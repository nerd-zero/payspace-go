package payspace

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestEmployeeLeaveService_ListAdjustments(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeLeaveAdjustment", func(w http.ResponseWriter, r *http.Request) {
		val := 5.0
		writeODataList(w, []EmployeeLeaveAdjustment{
			{LeaveAdjustmentId: 1, EmployeeNumber: "EMP001", LeaveType: "Annual", AdjustmentValue: &val},
		})
	})
	_, client := testServerAndClient(t, mux)

	adjustments, resp, err := client.EmployeeLeave.ListLeaveAdjustments(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(adjustments) != 1 {
		t.Fatalf("expected 1 adjustment, got %d", len(adjustments))
	}
	if adjustments[0].LeaveType != "Annual" {
		t.Errorf("expected LeaveType=Annual, got %s", adjustments[0].LeaveType)
	}
}

func TestEmployeeLeaveService_ListAdjustmentsByPeriod(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeLeaveAdjustment/effective/2024-06-01", func(w http.ResponseWriter, r *http.Request) {
		val := 3.0
		writeODataList(w, []EmployeeLeaveAdjustment{
			{LeaveAdjustmentId: 2, EmployeeNumber: "EMP001", LeaveType: "Sick", AdjustmentValue: &val},
		})
	})
	_, client := testServerAndClient(t, mux)

	adjustments, resp, err := client.EmployeeLeave.ListLeaveAdjustmentsByDate(context.Background(), 1, "2024-06-01", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(adjustments) != 1 {
		t.Fatalf("expected 1 adjustment, got %d", len(adjustments))
	}
	if adjustments[0].LeaveType != "Sick" {
		t.Errorf("expected LeaveType=Sick, got %s", adjustments[0].LeaveType)
	}
}

func TestEmployeeLeaveService_GetAdjustment(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeLeaveAdjustment(10)", func(w http.ResponseWriter, r *http.Request) {
		val := 2.0
		writeJSON(w, http.StatusOK, EmployeeLeaveAdjustment{
			LeaveAdjustmentId: 10,
			EmployeeNumber:    "EMP001",
			LeaveType:         "Annual",
			AdjustmentValue:   &val,
		})
	})
	_, client := testServerAndClient(t, mux)

	adj, resp, err := client.EmployeeLeave.GetLeaveAdjustment(context.Background(), 1, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if adj.LeaveAdjustmentId != 10 {
		t.Errorf("expected LeaveAdjustmentId=10, got %d", adj.LeaveAdjustmentId)
	}
}

func TestEmployeeLeaveService_CreateAdjustment(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeeLeaveAdjustment", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeeLeaveAdjustment
		json.NewDecoder(r.Body).Decode(&body)
		if body.EmployeeNumber != "EMP001" {
			t.Errorf("expected EmployeeNumber=EMP001, got %s", body.EmployeeNumber)
		}
		if body.LeaveType != "Annual" {
			t.Errorf("expected LeaveType=Annual, got %s", body.LeaveType)
		}
		body.LeaveAdjustmentId = 20
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	val := 5.0
	input := &EmployeeLeaveAdjustment{EmployeeNumber: "EMP001", LeaveType: "Annual", AdjustmentValue: &val}
	created, resp, err := client.EmployeeLeave.CreateLeaveAdjustment(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.LeaveAdjustmentId != 20 {
		t.Errorf("expected LeaveAdjustmentId=20, got %d", created.LeaveAdjustmentId)
	}
}

func TestEmployeeLeaveService_DeleteAdjustment(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EmployeeLeaveAdjustment(20)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.EmployeeLeave.DeleteLeaveAdjustment(context.Background(), 1, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestEmployeeLeaveService_ListApplications(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeLeaveApplication", func(w http.ResponseWriter, r *http.Request) {
		days := 5.0
		writeODataList(w, []EmployeeLeaveApplication{
			{LeaveAdjustmentId: 1, EmployeeNumber: "EMP001", LeaveType: "Annual", NoOfDays: &days, LeaveStatus: "Approved"},
		})
	})
	_, client := testServerAndClient(t, mux)

	apps, resp, err := client.EmployeeLeave.ListLeaveApplications(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(apps) != 1 {
		t.Fatalf("expected 1 application, got %d", len(apps))
	}
	if apps[0].LeaveStatus != "Approved" {
		t.Errorf("expected LeaveStatus=Approved, got %s", apps[0].LeaveStatus)
	}
}

func TestEmployeeLeaveService_CreateApplication(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeeLeaveApplication", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeeLeaveApplication
		json.NewDecoder(r.Body).Decode(&body)
		if body.EmployeeNumber != "EMP001" {
			t.Errorf("expected EmployeeNumber=EMP001, got %s", body.EmployeeNumber)
		}
		if body.LeaveType != "Sick" {
			t.Errorf("expected LeaveType=Sick, got %s", body.LeaveType)
		}
		body.LeaveAdjustmentId = 30
		body.LeaveStatus = "Pending"
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	days := 2.0
	input := &EmployeeLeaveApplication{
		EmployeeNumber: "EMP001",
		LeaveType:      "Sick",
		LeaveStartDate: "2024-06-10",
		LeaveEndDate:   "2024-06-11",
		NoOfDays:       &days,
	}
	created, resp, err := client.EmployeeLeave.CreateLeaveApplication(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.LeaveAdjustmentId != 30 {
		t.Errorf("expected LeaveAdjustmentId=30, got %d", created.LeaveAdjustmentId)
	}
	if created.LeaveStatus != "Pending" {
		t.Errorf("expected LeaveStatus=Pending, got %s", created.LeaveStatus)
	}
}

func TestEmployeeLeaveService_DeleteApplication(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EmployeeLeaveApplication(30)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.EmployeeLeave.DeleteLeaveApplication(context.Background(), 1, 30)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestEmployeeLeaveService_ListSetups(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeLeaveSetup", func(w http.ResponseWriter, r *http.Request) {
		entitlement := 21.0
		balance := 15.0
		writeODataList(w, []EmployeeLeaveSetup{
			{LeaveSetupId: 1, EmployeeNumber: "EMP001", LeaveType: "Annual", Entitlement: &entitlement, Balance: &balance},
		})
	})
	_, client := testServerAndClient(t, mux)

	setups, resp, err := client.EmployeeLeave.ListLeaveSetups(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(setups) != 1 {
		t.Fatalf("expected 1 setup, got %d", len(setups))
	}
	if setups[0].LeaveType != "Annual" {
		t.Errorf("expected LeaveType=Annual, got %s", setups[0].LeaveType)
	}
	if setups[0].Entitlement == nil || *setups[0].Entitlement != 21.0 {
		t.Errorf("expected Entitlement=21, got %v", setups[0].Entitlement)
	}
}

func TestEmployeeLeaveService_ListEntitlements(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeLeaveSetupEntitlement", func(w http.ResponseWriter, r *http.Request) {
		ent := 21.0
		writeODataList(w, []EmployeeLeaveSetupEntitlement{
			{LeaveSetupEntitlementId: 1, EmployeeNumber: "EMP001", LeaveType: "Annual", Entitlement: &ent},
		})
	})
	_, client := testServerAndClient(t, mux)

	entitlements, resp, err := client.EmployeeLeave.ListLeaveSetupEntitlements(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(entitlements) != 1 {
		t.Fatalf("expected 1 entitlement, got %d", len(entitlements))
	}
	if entitlements[0].LeaveType != "Annual" {
		t.Errorf("expected LeaveType=Annual, got %s", entitlements[0].LeaveType)
	}
}

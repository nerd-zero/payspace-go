package payspace

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
)

func TestEmployeeService_List(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []Employee{
			{EmployeeId: 1, EmployeeNumber: "EMP001", FirstName: "John", Surname: "Doe"},
			{EmployeeId: 2, EmployeeNumber: "EMP002", FirstName: "Jane", Surname: "Smith"},
		})
	})
	_, client := testServerAndClient(t, mux)

	employees, resp, err := client.Employees.List(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(employees) != 2 {
		t.Fatalf("expected 2 employees, got %d", len(employees))
	}
	if employees[0].FirstName != "John" {
		t.Errorf("expected FirstName=John, got %s", employees[0].FirstName)
	}
	if employees[1].EmployeeNumber != "EMP002" {
		t.Errorf("expected EmployeeNumber=EMP002, got %s", employees[1].EmployeeNumber)
	}
}

func TestEmployeeService_Get(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee(123)", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, Employee{
			EmployeeId:     123,
			EmployeeNumber: "EMP123",
			FirstName:      "Alice",
			Surname:        "Wonder",
		})
	})
	_, client := testServerAndClient(t, mux)

	employee, resp, err := client.Employees.Get(context.Background(), 1, 123)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if employee.EmployeeId != 123 {
		t.Errorf("expected EmployeeId=123, got %d", employee.EmployeeId)
	}
	if employee.FirstName != "Alice" {
		t.Errorf("expected FirstName=Alice, got %s", employee.FirstName)
	}
}

func TestEmployeeService_Create(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/Employee", func(w http.ResponseWriter, r *http.Request) {
		var body Employee
		json.NewDecoder(r.Body).Decode(&body)
		if body.FirstName != "Bob" {
			t.Errorf("expected FirstName=Bob in request body, got %s", body.FirstName)
		}
		if body.Surname != "Builder" {
			t.Errorf("expected Surname=Builder in request body, got %s", body.Surname)
		}
		body.EmployeeId = 10
		body.EmployeeNumber = "EMP010"
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	input := &Employee{FirstName: "Bob", Surname: "Builder"}
	created, resp, err := client.Employees.Create(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.EmployeeId != 10 {
		t.Errorf("expected EmployeeId=10, got %d", created.EmployeeId)
	}
	if created.EmployeeNumber != "EMP010" {
		t.Errorf("expected EmployeeNumber=EMP010, got %s", created.EmployeeNumber)
	}
}

func TestEmployeeService_Update(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/Employee(123)", func(w http.ResponseWriter, r *http.Request) {
		var body Employee
		json.NewDecoder(r.Body).Decode(&body)
		if body.FirstName != "Updated" {
			t.Errorf("expected FirstName=Updated, got %s", body.FirstName)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.Employees.Update(context.Background(), 1, 123, &Employee{FirstName: "Updated"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestEmployeeService_ListAll(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee/all", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []Employee{
			{EmployeeId: 1, EmployeeNumber: "EMP001", FirstName: "Active"},
			{EmployeeId: 2, EmployeeNumber: "EMP002", FirstName: "Inactive"},
		})
	})
	_, client := testServerAndClient(t, mux)

	employees, resp, err := client.Employees.ListAll(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(employees) != 2 {
		t.Fatalf("expected 2 employees, got %d", len(employees))
	}
}

func TestEmployeeService_ListByEffectiveDate(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/Employee/effective/2024-01-15", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []Employee{
			{EmployeeId: 1, EmployeeNumber: "EMP001", FirstName: "Effective"},
		})
	})
	_, client := testServerAndClient(t, mux)

	employees, resp, err := client.Employees.ListByEffectiveDate(context.Background(), 1, "2024-01-15", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(employees) != 1 {
		t.Fatalf("expected 1 employee, got %d", len(employees))
	}
	if employees[0].FirstName != "Effective" {
		t.Errorf("expected FirstName=Effective, got %s", employees[0].FirstName)
	}
}

func TestEmployeeService_ListAddresses(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeAddress/EMP001", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []EmployeeAddress{
			{AddressId: 1, EmployeeNumber: "EMP001", AddressType: "Postal", AddressLine1: "123 Main St"},
		})
	})
	_, client := testServerAndClient(t, mux)

	addresses, resp, err := client.Employees.ListAddresses(context.Background(), 1, "EMP001", nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(addresses) != 1 {
		t.Fatalf("expected 1 address, got %d", len(addresses))
	}
	if addresses[0].AddressLine1 != "123 Main St" {
		t.Errorf("expected AddressLine1='123 Main St', got %s", addresses[0].AddressLine1)
	}
}

func TestEmployeeService_UpdateAddress(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/EmployeeAddress(5)", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeeAddress
		json.NewDecoder(r.Body).Decode(&body)
		if body.AddressLine1 != "456 New St" {
			t.Errorf("expected AddressLine1='456 New St', got %s", body.AddressLine1)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.Employees.UpdateAddress(context.Background(), 1, 5, &EmployeeAddress{AddressLine1: "456 New St"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestEmployeeService_ListEmploymentStatuses(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeeEmploymentStatus", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []EmployeeEmploymentStatus{
			{EmploymentStatusId: 1, EmployeeNumber: "EMP001", EmploymentStatus: "Active"},
		})
	})
	_, client := testServerAndClient(t, mux)

	statuses, resp, err := client.Employees.ListEmploymentStatuses(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(statuses) != 1 {
		t.Fatalf("expected 1 status, got %d", len(statuses))
	}
	if statuses[0].EmploymentStatus != "Active" {
		t.Errorf("expected EmploymentStatus=Active, got %s", statuses[0].EmploymentStatus)
	}
}

func TestEmployeeService_CreateEmploymentStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeeEmploymentStatus", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeeEmploymentStatus
		json.NewDecoder(r.Body).Decode(&body)
		if body.EmployeeNumber != "EMP001" {
			t.Errorf("expected EmployeeNumber=EMP001, got %s", body.EmployeeNumber)
		}
		body.EmploymentStatusId = 10
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	input := &EmployeeEmploymentStatus{EmployeeNumber: "EMP001", EmploymentStatus: "Active"}
	created, resp, err := client.Employees.CreateEmploymentStatus(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.EmploymentStatusId != 10 {
		t.Errorf("expected EmploymentStatusId=10, got %d", created.EmploymentStatusId)
	}
}

func TestEmployeeService_DeleteEmploymentStatus(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EmployeeEmploymentStatus(5)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.Employees.DeleteEmploymentStatus(context.Background(), 1, 5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

func TestEmployeeService_ListPositions(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /odata/v2.0/1/EmployeePosition", func(w http.ResponseWriter, r *http.Request) {
		writeODataList(w, []EmployeePosition{
			{EmployeePositionId: 1, EmployeeNumber: "EMP001", JobTitle: "Developer"},
		})
	})
	_, client := testServerAndClient(t, mux)

	positions, resp, err := client.Employees.ListPositions(context.Background(), 1, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	if len(positions) != 1 {
		t.Fatalf("expected 1 position, got %d", len(positions))
	}
	if positions[0].JobTitle != "Developer" {
		t.Errorf("expected JobTitle=Developer, got %s", positions[0].JobTitle)
	}
}

func TestEmployeeService_CreatePosition(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /odata/v2.0/1/EmployeePosition", func(w http.ResponseWriter, r *http.Request) {
		var body EmployeePosition
		json.NewDecoder(r.Body).Decode(&body)
		if body.JobTitle != "Manager" {
			t.Errorf("expected JobTitle=Manager, got %s", body.JobTitle)
		}
		body.EmployeePositionId = 20
		writeJSON(w, http.StatusCreated, body)
	})
	_, client := testServerAndClient(t, mux)

	input := &EmployeePosition{EmployeeNumber: "EMP001", JobTitle: "Manager"}
	created, resp, err := client.Employees.CreatePosition(context.Background(), 1, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected 201, got %d", resp.StatusCode)
	}
	if created.EmployeePositionId != 20 {
		t.Errorf("expected EmployeePositionId=20, got %d", created.EmployeePositionId)
	}
}

func TestEmployeeService_UpdatePosition(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /odata/v2.0/1/EmployeePosition(20)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.Employees.UpdatePosition(context.Background(), 1, 20, &EmployeePosition{JobTitle: "Senior Manager"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", resp.StatusCode)
	}
}

func TestEmployeeService_DeletePosition(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /odata/v2.0/1/EmployeePosition(20)", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	_, client := testServerAndClient(t, mux)

	resp, err := client.Employees.DeletePosition(context.Background(), 1, 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200, got %d", resp.StatusCode)
	}
}

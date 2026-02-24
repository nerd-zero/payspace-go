package payspace

import (
	"net/url"
	"strings"
	"testing"
)

func TestQueryParams_Filter(t *testing.T) {
	q := NewQuery().Filter("EmployeeNumber eq 'EMP001'")
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	if got := vals.Get("$filter"); got != "EmployeeNumber eq 'EMP001'" {
		t.Errorf("expected $filter=EmployeeNumber eq 'EMP001', got %q", got)
	}
}

func TestQueryParams_TopSkip(t *testing.T) {
	q := NewQuery().Top(50).Skip(10)
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	if got := vals.Get("$top"); got != "50" {
		t.Errorf("expected $top=50, got %q", got)
	}
	if got := vals.Get("$skip"); got != "10" {
		t.Errorf("expected $skip=10, got %q", got)
	}
}

func TestQueryParams_TopClamp(t *testing.T) {
	q := NewQuery().Top(500)
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	if got := vals.Get("$top"); got != "100" {
		t.Errorf("expected $top to be clamped to 100, got %q", got)
	}
}

func TestQueryParams_OrderBy(t *testing.T) {
	q := NewQuery().OrderBy("Surname asc")
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	if got := vals.Get("$orderby"); got != "Surname asc" {
		t.Errorf("expected $orderby=Surname asc, got %q", got)
	}
}

func TestQueryParams_Select(t *testing.T) {
	q := NewQuery().Select("EmployeeNumber", "FirstName", "Surname")
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	got := vals.Get("$select")
	if got != "EmployeeNumber,FirstName,Surname" {
		t.Errorf("expected $select=EmployeeNumber,FirstName,Surname, got %q", got)
	}
}

func TestQueryParams_Expand(t *testing.T) {
	q := NewQuery().Expand("OrganizationGroups", "Addresses")
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	got := vals.Get("$expand")
	if got != "OrganizationGroups,Addresses" {
		t.Errorf("expected $expand=OrganizationGroups,Addresses, got %q", got)
	}
}

func TestQueryParams_Count(t *testing.T) {
	q := NewQuery().Count()
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	if got := vals.Get("$count"); got != "true" {
		t.Errorf("expected $count=true, got %q", got)
	}
}

func TestQueryParams_Param(t *testing.T) {
	q := NewQuery().Param("frequency", "Monthly").Param("period", "1")
	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	if got := vals.Get("frequency"); got != "Monthly" {
		t.Errorf("expected frequency=Monthly, got %q", got)
	}
	if got := vals.Get("period"); got != "1" {
		t.Errorf("expected period=1, got %q", got)
	}
}

func TestQueryParams_Combined(t *testing.T) {
	q := NewQuery().
		Filter("Surname eq 'Smith'").
		Top(25).
		Skip(0).
		OrderBy("FirstName asc").
		Select("EmployeeNumber", "FirstName").
		Count().
		Param("frequency", "Monthly")

	encoded := q.Encode()

	vals, err := url.ParseQuery(encoded)
	if err != nil {
		t.Fatalf("failed to parse query: %v", err)
	}
	if got := vals.Get("$filter"); got != "Surname eq 'Smith'" {
		t.Errorf("expected $filter=Surname eq 'Smith', got %q", got)
	}
	if got := vals.Get("$top"); got != "25" {
		t.Errorf("expected $top=25, got %q", got)
	}
	if got := vals.Get("$skip"); got != "0" {
		t.Errorf("expected $skip=0, got %q", got)
	}
	if got := vals.Get("$orderby"); got != "FirstName asc" {
		t.Errorf("expected $orderby=FirstName asc, got %q", got)
	}
	if !strings.Contains(vals.Get("$select"), "EmployeeNumber") {
		t.Errorf("expected $select to contain EmployeeNumber, got %q", vals.Get("$select"))
	}
	if got := vals.Get("$count"); got != "true" {
		t.Errorf("expected $count=true, got %q", got)
	}
	if got := vals.Get("frequency"); got != "Monthly" {
		t.Errorf("expected frequency=Monthly, got %q", got)
	}
}

func TestQueryParams_NilEncode(t *testing.T) {
	var q *QueryParams
	if got := q.Encode(); got != "" {
		t.Errorf("expected empty string for nil QueryParams, got %q", got)
	}
}

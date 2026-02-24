// Example: employee CRUD operations -- create, get, update, and filtered listing.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/nerd-zero/payspace-go"
	"go.uber.org/zap/zapcore"
)

func main() {
	clientID := os.Getenv("PAYSPACE_CLIENT_ID")
	clientSecret := os.Getenv("PAYSPACE_CLIENT_SECRET")
	companyIDStr := os.Getenv("PAYSPACE_COMPANY_ID")

	if clientID == "" || clientSecret == "" || companyIDStr == "" {
		log.Fatal("Set PAYSPACE_CLIENT_ID, PAYSPACE_CLIENT_SECRET, and PAYSPACE_COMPANY_ID")
	}

	companyID, err := strconv.Atoi(companyIDStr)
	if err != nil {
		log.Fatalf("Invalid PAYSPACE_COMPANY_ID: %v", err)
	}

	client, err := payspace.NewClient(clientID, clientSecret,
		payspace.WithEnvironment(payspace.Staging),
		payspace.WithLogLevel(zapcore.WarnLevel),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// --- Create a new employee ---
	newEmp := &payspace.Employee{
		EmployeeNumber: "EMP099",
		FirstName:      "Jane",
		Surname:        "Doe",
		EmailAddress:   "jane.doe@example.com",
		Gender:         "Female",
	}

	created, _, err := client.Employees.Create(ctx, companyID, newEmp)
	if err != nil {
		if payspace.IsValidationError(err) {
			if apiErr, ok := err.(*payspace.APIError); ok {
				for _, d := range apiErr.Details {
					fmt.Printf("Validation error: %s\n", d.Message)
				}
			}
			return
		}
		log.Fatalf("Failed to create employee: %v", err)
	}
	fmt.Printf("Created employee: %s %s (ID: %d)\n", created.FirstName, created.Surname, created.EmployeeId)

	// --- Get employee by ID ---
	emp, _, err := client.Employees.Get(ctx, companyID, created.EmployeeId)
	if err != nil {
		log.Fatalf("Failed to get employee: %v", err)
	}
	fmt.Printf("Fetched employee: %s %s\n", emp.FirstName, emp.Surname)

	// --- Update an employee ---
	update := &payspace.Employee{
		KnownAs: "JD",
	}
	_, err = client.Employees.Update(ctx, companyID, created.EmployeeId, update)
	if err != nil {
		log.Fatalf("Failed to update employee: %v", err)
	}
	fmt.Println("Updated employee KnownAs to 'JD'")

	// --- Filter employees by surname ---
	query := payspace.NewQuery().
		Filter("Surname eq 'Doe'").
		Select("EmployeeNumber", "FirstName", "Surname", "EmailAddress").
		OrderBy("FirstName asc")

	employees, _, err := client.Employees.List(ctx, companyID, query)
	if err != nil {
		log.Fatalf("Failed to list employees: %v", err)
	}

	fmt.Printf("\nEmployees with surname 'Doe': %d\n", len(employees))
	for _, e := range employees {
		fmt.Printf("  %s %s (%s) - %s\n", e.FirstName, e.Surname, e.EmployeeNumber, e.EmailAddress)
	}

	// --- List employees effective as of a date ---
	dated, _, err := client.Employees.ListByEffectiveDate(ctx, companyID, "2025-01-01", nil)
	if err != nil {
		log.Fatalf("Failed to list employees by date: %v", err)
	}
	fmt.Printf("\nEmployees effective 2025-01-01: %d\n", len(dated))

	// --- List all employees including inactive ---
	all, _, err := client.Employees.ListAll(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list all employees: %v", err)
	}
	fmt.Printf("All employees (incl. inactive): %d\n", len(all))
}

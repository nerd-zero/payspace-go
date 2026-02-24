// Example: advanced OData query builder usage -- filtering, pagination, counting, and sorting.
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

	// --- Filter with complex expressions ---
	fmt.Println("=== Filter: employees with surname starting with 'S' ===")
	query := payspace.NewQuery().
		Filter("startswith(Surname, 'S')").
		OrderBy("Surname asc, FirstName asc").
		Top(5)

	employees, _, err := client.Employees.List(ctx, companyID, query)
	if err != nil {
		log.Fatalf("Failed to list employees: %v", err)
	}
	for _, emp := range employees {
		fmt.Printf("  %s %s (%s)\n", emp.FirstName, emp.Surname, emp.EmployeeNumber)
	}

	// --- Count total records ---
	fmt.Println("\n=== Count total employees ===")
	countQuery := payspace.NewQuery().Top(1).Count()
	_, resp, err := client.Employees.List(ctx, companyID, countQuery)
	if err != nil {
		log.Fatalf("Failed to count employees: %v", err)
	}
	if resp.ODataCount != nil {
		fmt.Printf("  Total employees: %d\n", *resp.ODataCount)
	}

	// --- Pagination loop ---
	fmt.Println("\n=== Paginate through all employees (page size 50) ===")
	pageSize := 50
	page := 0
	total := 0

	for {
		pageQuery := payspace.NewQuery().
			Top(pageSize).
			Skip(page * pageSize).
			OrderBy("EmployeeNumber asc")

		batch, _, err := client.Employees.List(ctx, companyID, pageQuery)
		if err != nil {
			log.Fatalf("Failed to list employees (page %d): %v", page, err)
		}

		total += len(batch)
		fmt.Printf("  Page %d: %d employees\n", page+1, len(batch))

		if len(batch) < pageSize {
			break // last page
		}
		page++
	}
	fmt.Printf("  Total fetched: %d\n", total)

	// --- Select specific fields ---
	fmt.Println("\n=== Select only EmployeeNumber and Surname ===")
	selectQuery := payspace.NewQuery().
		Select("EmployeeNumber", "Surname").
		Top(5)

	selected, _, err := client.Employees.List(ctx, companyID, selectQuery)
	if err != nil {
		log.Fatalf("Failed to list employees: %v", err)
	}
	for _, emp := range selected {
		// Only selected fields are populated; others will be zero values.
		fmt.Printf("  %s - %s\n", emp.EmployeeNumber, emp.Surname)
	}
}

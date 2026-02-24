// Example: basic client setup, listing employees, and accessing companies.
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

	// Create a client targeting the staging environment.
	client, err := payspace.NewClient(clientID, clientSecret,
		payspace.WithEnvironment(payspace.Staging),
		payspace.WithLogLevel(zapcore.WarnLevel),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// List the first 10 employees sorted by surname.
	query := payspace.NewQuery().Top(10).OrderBy("Surname asc")
	employees, resp, err := client.Employees.List(ctx, companyID, query)
	if err != nil {
		log.Fatalf("Failed to list employees: %v", err)
	}

	fmt.Printf("Fetched %d employees (rate limit: %d/%d)\n\n",
		len(employees), resp.RateLimit.Remaining, resp.RateLimit.Limit)

	for _, emp := range employees {
		fmt.Printf("  %s %s (ID: %d, Number: %s)\n",
			emp.FirstName, emp.Surname, emp.EmployeeId, emp.EmployeeNumber)
	}

	// After the first API call, you can access the companies list from the auth token.
	groups, err := client.Companies_()
	if err != nil {
		log.Fatalf("Failed to get companies: %v", err)
	}

	fmt.Printf("\nAccessible companies:\n")
	for _, group := range groups {
		fmt.Printf("  Group: %s\n", group.GroupDescription)
		for _, c := range group.Companies {
			fmt.Printf("    - %s (ID: %d, Code: %s)\n", c.CompanyName, c.CompanyID, c.CompanyCode)
		}
	}
}

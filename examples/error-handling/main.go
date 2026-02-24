// Example: error handling patterns -- all error types, helper functions, and validation details.
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

	// --- Handle not-found errors (404) ---
	fmt.Println("=== Not Found (404) ===")
	_, _, err = client.Employees.Get(ctx, companyID, 9999999)
	if err != nil {
		if payspace.IsNotFound(err) {
			fmt.Println("  Employee not found (expected)")
		} else {
			fmt.Printf("  Unexpected error: %v\n", err)
		}
	}

	// --- Handle validation errors (400) ---
	fmt.Println("\n=== Validation Error (400) ===")
	invalidEmp := &payspace.Employee{
		// Missing required fields to trigger validation.
		EmployeeNumber: "",
	}
	_, _, err = client.Employees.Create(ctx, companyID, invalidEmp)
	if err != nil {
		if payspace.IsValidationError(err) {
			fmt.Println("  Validation failed (expected):")
			if apiErr, ok := err.(*payspace.APIError); ok {
				fmt.Printf("  Status: %d\n", apiErr.StatusCode)
				fmt.Printf("  Message: %s\n", apiErr.Message)
				for _, detail := range apiErr.Details {
					fmt.Printf("    - %s\n", detail.Message)
				}
			}
		} else {
			fmt.Printf("  Error: %v\n", err)
		}
	}

	// --- Inspect rate limit info from responses ---
	fmt.Println("\n=== Rate Limit Info ===")
	_, resp, err := client.Employees.List(ctx, companyID, payspace.NewQuery().Top(1))
	if err != nil {
		fmt.Printf("  Error: %v\n", err)

		// Check if rate limited.
		if payspace.IsRateLimited(err) {
			if rlErr, ok := err.(*payspace.RateLimitError); ok {
				fmt.Printf("  Rate limited! Limit: %d, Remaining: %d, Reset: %s\n",
					rlErr.Limit, rlErr.Remaining, rlErr.Reset)
			}
		}
	} else {
		fmt.Printf("  Limit: %d\n", resp.RateLimit.Limit)
		fmt.Printf("  Remaining: %d\n", resp.RateLimit.Remaining)
		fmt.Printf("  Reset: %s\n", resp.RateLimit.Reset)
	}

	// --- Check for unauthorized (401) ---
	fmt.Println("\n=== Unauthorized Check ===")
	// This normally won't happen because the SDK auto-refreshes tokens,
	// but if credentials are invalid you'll see this.
	badClient, err := payspace.NewClient("bad-id", "bad-secret",
		payspace.WithEnvironment(payspace.Staging),
		payspace.WithLogLevel(zapcore.ErrorLevel),
		payspace.WithRetry(0, 0), // Disable retries for this demo.
	)
	if err != nil {
		log.Fatalf("Failed to create bad client: %v", err)
	}

	_, _, err = badClient.Employees.List(ctx, companyID, nil)
	if err != nil {
		if payspace.IsUnauthorized(err) {
			fmt.Println("  Unauthorized (expected with bad credentials)")
		} else {
			fmt.Printf("  Error: %v\n", err)
		}
	}

	// --- General error inspection ---
	fmt.Println("\n=== General APIError inspection ===")
	_, _, err = client.Employees.Get(ctx, companyID, 9999999)
	if err != nil {
		if apiErr, ok := err.(*payspace.APIError); ok {
			fmt.Printf("  Status Code: %d\n", apiErr.StatusCode)
			fmt.Printf("  Message: %s\n", apiErr.Message)
			fmt.Printf("  Error string: %s\n", apiErr.Error())
		} else {
			fmt.Printf("  Non-API error: %v\n", err)
		}
	}
}

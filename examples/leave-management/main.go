// Example: leave management -- lookup leave types, check balances, and create leave applications.
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

	// --- List available leave types ---
	fmt.Println("=== Leave Types ===")
	leaveTypes, _, err := client.Lookups.ListLeaveTypes(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list leave types: %v", err)
	}
	for _, lt := range leaveTypes {
		fmt.Printf("  [%s] %s\n", lt.Code, lt.Description)
	}

	// --- Check leave balances via leave setups ---
	fmt.Println("\n=== Employee Leave Balances (first 10) ===")
	query := payspace.NewQuery().Top(10)
	setups, _, err := client.EmployeeLeave.ListLeaveSetups(ctx, companyID, query)
	if err != nil {
		log.Fatalf("Failed to list leave setups: %v", err)
	}

	for _, setup := range setups {
		balance := 0.0
		if setup.Balance != nil {
			balance = *setup.Balance
		}
		entitlement := 0.0
		if setup.Entitlement != nil {
			entitlement = *setup.Entitlement
		}
		fmt.Printf("  %s - %s: Balance=%.1f / Entitlement=%.1f (Cycle: %s to %s)\n",
			setup.EmployeeNumber, setup.LeaveType, balance, entitlement,
			setup.CycleStartDate, setup.CycleEndDate)
	}

	// --- Create a leave application ---
	fmt.Println("\n=== Create Leave Application ===")
	days := 3.0
	application := &payspace.EmployeeLeaveApplication{
		EmployeeNumber: "EMP001",
		LeaveType:      "Annual",
		LeaveStartDate: "2025-06-01",
		LeaveEndDate:   "2025-06-03",
		NoOfDays:       &days,
		LeaveReason:    "Family vacation",
	}

	created, _, err := client.EmployeeLeave.CreateLeaveApplication(ctx, companyID, application)
	if err != nil {
		if payspace.IsValidationError(err) {
			if apiErr, ok := err.(*payspace.APIError); ok {
				fmt.Println("  Validation errors:")
				for _, d := range apiErr.Details {
					fmt.Printf("    - %s\n", d.Message)
				}
			}
			return
		}
		log.Fatalf("Failed to create leave application: %v", err)
	}
	fmt.Printf("  Created leave application ID: %d (Status: %s)\n",
		created.LeaveAdjustmentId, created.LeaveStatus)

	// --- List pending leave applications ---
	fmt.Println("\n=== Pending Leave Applications ===")
	pendingQuery := payspace.NewQuery().
		Filter("LeaveStatus eq 'Pending'").
		Top(10)

	applications, _, err := client.EmployeeLeave.ListLeaveApplications(ctx, companyID, pendingQuery)
	if err != nil {
		log.Fatalf("Failed to list leave applications: %v", err)
	}

	for _, app := range applications {
		days := 0.0
		if app.NoOfDays != nil {
			days = *app.NoOfDays
		}
		fmt.Printf("  %s: %s %s - %s (%.0f days) [%s]\n",
			app.EmployeeNumber, app.LeaveType, app.LeaveStartDate, app.LeaveEndDate, days, app.LeaveStatus)
	}
}

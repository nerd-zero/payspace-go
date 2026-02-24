// Example: edit payslips -- list, create, batch operations, and delete.
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
	frequency := "1" // Monthly
	period := "1"    // Period 1

	// --- List existing edit payslips ---
	fmt.Printf("=== Edit Payslips (frequency=%s, period=%s) ===\n", frequency, period)
	edits, _, err := client.EmployeePayslips.ListEditPayslips(ctx, companyID, frequency, period, nil)
	if err != nil {
		log.Fatalf("Failed to list edit payslips: %v", err)
	}

	for _, ep := range edits {
		amount := 0.0
		if ep.Amount != nil {
			amount = *ep.Amount
		}
		fmt.Printf("  ID: %d, Employee: %s, Component: %s, Amount: %.2f\n",
			ep.EditPayslipId, ep.EmployeeNumber, ep.ComponentCode, amount)
	}

	// --- Create a single edit payslip entry ---
	fmt.Println("\n=== Create Edit Payslip ===")
	amount := 500.0
	newEdit := &payspace.EditPayslip{
		EmployeeNumber: "EMP001",
		ComponentCode:  "BONUS",
		Amount:         &amount,
	}

	created, _, err := client.EmployeePayslips.CreateEditPayslip(ctx, companyID, frequency, period, "Add", newEdit)
	if err != nil {
		if payspace.IsValidationError(err) {
			if apiErr, ok := err.(*payspace.APIError); ok {
				for _, d := range apiErr.Details {
					fmt.Printf("  Validation: %s\n", d.Message)
				}
			}
			return
		}
		log.Fatalf("Failed to create edit payslip: %v", err)
	}
	fmt.Printf("  Created edit payslip ID: %d\n", created.EditPayslipId)

	// --- Batch edit payslips ---
	fmt.Println("\n=== Batch Edit Payslips ===")
	amount1, amount2 := 200.0, 300.0
	batch := []payspace.EditPayslip{
		{EmployeeNumber: "EMP001", ComponentCode: "ALLOW1", Amount: &amount1},
		{EmployeeNumber: "EMP002", ComponentCode: "ALLOW1", Amount: &amount2},
	}

	results, _, err := client.EmployeePayslips.BatchEditPayslips(ctx, companyID, frequency, period, batch)
	if err != nil {
		log.Fatalf("Failed to batch edit payslips: %v", err)
	}
	fmt.Printf("  Batch created %d entries\n", len(results))
	for _, r := range results {
		a := 0.0
		if r.Amount != nil {
			a = *r.Amount
		}
		fmt.Printf("    ID: %d, Employee: %s, Amount: %.2f\n",
			r.EditPayslipId, r.EmployeeNumber, a)
	}

	// --- Delete an edit payslip ---
	if created.EditPayslipId > 0 {
		fmt.Printf("\n=== Delete Edit Payslip (ID: %d) ===\n", created.EditPayslipId)
		_, err = client.EmployeePayslips.DeleteEditPayslip(ctx, companyID, created.EditPayslipId, frequency, period)
		if err != nil {
			log.Fatalf("Failed to delete edit payslip: %v", err)
		}
		fmt.Println("  Deleted successfully")
	}
}

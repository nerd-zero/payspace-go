// Example: working with pay rates, company payroll frequencies/runs, and recurring components.
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

	// --- List company payroll frequencies ---
	fmt.Println("=== Payroll Frequencies ===")
	frequencies, _, err := client.CompanyPayroll.ListFrequencies(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list frequencies: %v", err)
	}
	for _, freq := range frequencies {
		fmt.Printf("  [%s] %s (%d periods/year)\n",
			freq.FrequencyCode, freq.Description, freq.Periods)
	}

	// --- List payroll runs ---
	fmt.Println("\n=== Payroll Runs ===")
	runs, _, err := client.CompanyPayroll.ListRuns(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list runs: %v", err)
	}
	for _, run := range runs {
		fmt.Printf("  %s - Year: %d, Period: %d, Status: %s\n",
			run.Description, run.Year, run.Period, run.Status)
	}

	// --- List employee pay rates ---
	fmt.Println("\n=== Employee Pay Rates (first 10) ===")
	query := payspace.NewQuery().Top(10).OrderBy("EmployeeNumber asc")
	payRates, _, err := client.EmployeePayroll.ListPayRates(ctx, companyID, query)
	if err != nil {
		log.Fatalf("Failed to list pay rates: %v", err)
	}
	for _, pr := range payRates {
		salary := 0.0
		if pr.BasicSalary != nil {
			salary = *pr.BasicSalary
		}
		fmt.Printf("  %s - Basic: %.2f, Type: %s, Freq: %s, Effective: %s\n",
			pr.EmployeeNumber, salary, pr.PayType, pr.PayFrequency, pr.EffectiveDate)
	}

	// --- List recurring components (generic map pattern) ---
	// The ListRecurring method works with entity names like:
	//   "EmployeeBonusProvision", "EmployeeComponent", "EmployeeGarnishee",
	//   "EmployeeMedical", "EmployeePensionFund", "EmployeeRetirementAnnuity", "EmployeeLoan"
	fmt.Println("\n=== Recurring Bonus Provisions (frequency=1, period=1) ===")
	bonuses, _, err := client.EmployeePayroll.ListRecurring(ctx, companyID,
		"EmployeeBonusProvision", "1", "1", nil)
	if err != nil {
		log.Fatalf("Failed to list bonus provisions: %v", err)
	}
	for _, b := range bonuses {
		fmt.Printf("  Employee: %v, Component: %v, Amount: %v\n",
			b["EmployeeNumber"], b["ComponentCode"], b["Amount"])
	}

	// --- List company-level components for a frequency/period ---
	fmt.Println("\n=== Company Components (frequency=1, period=1) ===")
	components, _, err := client.CompanyPayroll.ListComponents(ctx, companyID, 1, 1, nil)
	if err != nil {
		log.Fatalf("Failed to list company components: %v", err)
	}
	for _, comp := range components {
		fmt.Printf("  [%s] %s (Type: %s)\n",
			comp.ComponentCode, comp.Description, comp.ComponentType)
	}
}

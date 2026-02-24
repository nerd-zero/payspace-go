// Example: retrieving payslips, payslip lines, consolidated payslips, and tax certificates.
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
	year, month := 2025, 1

	// --- List payslips for a specific period ---
	fmt.Printf("=== Payslips for %d/%02d ===\n", year, month)
	payslips, _, err := client.EmployeePayslips.ListPayslips(ctx, companyID, year, month, nil)
	if err != nil {
		log.Fatalf("Failed to list payslips: %v", err)
	}

	for _, ps := range payslips {
		netPay := 0.0
		if ps.NetPay != nil {
			netPay = *ps.NetPay
		}
		fmt.Printf("  %s %s (%s) - Net Pay: %.2f\n",
			ps.FirstName, ps.Surname, ps.EmployeeNumber, netPay)
	}

	// --- Filter payslips for a specific employee ---
	if len(payslips) > 0 {
		empNumber := payslips[0].EmployeeNumber
		fmt.Printf("\n=== Payslip lines for %s (%d/%02d) ===\n", empNumber, year, month)

		query := payspace.NewQuery().
			Filter(fmt.Sprintf("EmployeeNumber eq '%s'", empNumber))

		lines, _, err := client.EmployeePayslips.ListPayslipLines(ctx, companyID, year, month, query)
		if err != nil {
			log.Fatalf("Failed to list payslip lines: %v", err)
		}

		for _, line := range lines {
			amount := 0.0
			if line.Amount != nil {
				amount = *line.Amount
			}
			fmt.Printf("  [%s] %s: %.2f (%s)\n",
				line.ComponentCode, line.Description, amount, line.LineType)
		}
	}

	// --- Consolidated payslips ---
	fmt.Printf("\n=== Consolidated payslips for %d/%02d ===\n", year, month)
	consolidated, _, err := client.EmployeePayslips.ListConsolidatedPayslips(ctx, companyID, year, month, nil)
	if err != nil {
		log.Fatalf("Failed to list consolidated payslips: %v", err)
	}
	fmt.Printf("  Found %d consolidated payslips\n", len(consolidated))

	// --- Tax certificates ---
	fmt.Println("\n=== Tax certificates ===")
	certs, _, err := client.EmployeePayslips.ListTaxCertificates(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list tax certificates: %v", err)
	}

	for _, cert := range certs {
		fmt.Printf("  %s - Tax Year: %d, Type: %s, Issued: %s\n",
			cert.EmployeeNumber, cert.TaxYear, cert.CertificateType, cert.IssueDate)
	}
	fmt.Printf("  Total certificates: %d\n", len(certs))
}

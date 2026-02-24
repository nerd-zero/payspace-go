// Example: general ledger entries, extraction with polling, and EFT outbox.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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

	// --- List general ledger entries ---
	fmt.Printf("=== General Ledger (%d/%02d) ===\n", year, month)
	entries, _, err := client.CompanyConfig.ListGeneralLedger(ctx, companyID, year, month, nil)
	if err != nil {
		log.Fatalf("Failed to list GL entries: %v", err)
	}

	for _, gl := range entries {
		amount := 0.0
		if gl.Amount != nil {
			amount = *gl.Amount
		}
		fmt.Printf("  [%s] %s: %.2f\n", gl.AccountCode, gl.Description, amount)
	}
	fmt.Printf("  Total entries: %d\n", len(entries))

	// --- Start a GL extraction and poll for completion ---
	fmt.Printf("\n=== Start GL Extraction (%d/%02d) ===\n", year, month)
	extraction, _, err := client.CompanyConfig.ExtractGeneralLedger(ctx, companyID, year, month)
	if err != nil {
		log.Fatalf("Failed to start GL extraction: %v", err)
	}
	fmt.Printf("  Extraction ID: %d, Status: %s\n", extraction.ExtractionId, extraction.Status)

	// Poll for completion.
	fmt.Println("  Polling for completion...")
	for i := 0; i < 30; i++ {
		time.Sleep(2 * time.Second)

		status, _, err := client.CompanyConfig.ExtractGeneralLedgerStatus(ctx, companyID, extraction.ExtractionId)
		if err != nil {
			log.Fatalf("Failed to get extraction status: %v", err)
		}

		fmt.Printf("  [%d] Status: %s\n", i+1, status.Status)

		if status.Status == "Completed" || status.Status == "Failed" {
			break
		}
	}

	// --- List EFT outbox records ---
	fmt.Printf("\n=== EFT Outbox (%d/%02d) ===\n", year, month)
	eftRecords, _, err := client.CompanyConfig.ListEFTOutbox(ctx, companyID, year, month, nil)
	if err != nil {
		log.Fatalf("Failed to list EFT outbox: %v", err)
	}

	for _, eft := range eftRecords {
		amount := 0.0
		if eft.Amount != nil {
			amount = *eft.Amount
		}
		fmt.Printf("  %s - %s (%s): %.2f [%s]\n",
			eft.EmployeeNumber, eft.BankName, eft.AccountNumber, amount, eft.Status)
	}
	fmt.Printf("  Total EFT records: %d\n", len(eftRecords))
}

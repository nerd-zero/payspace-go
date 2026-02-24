// Example: company configuration -- organization positions, regions, lookups, and custom fields.
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

	// --- Organization positions ---
	fmt.Println("=== Organization Positions ===")
	positions, _, err := client.CompanyConfig.ListOrganizationPositions(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list positions: %v", err)
	}
	for _, pos := range positions {
		fmt.Printf("  [%s] %s - Dept: %s\n",
			pos.PositionCode, pos.PositionDescription, pos.Department)
	}

	// --- Regions ---
	fmt.Println("\n=== Company Regions ===")
	regions, _, err := client.CompanyConfig.ListRegions(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list regions: %v", err)
	}
	for _, r := range regions {
		fmt.Printf("  [%s] %s\n", r.RegionCode, r.Description)
	}

	// --- Generic lookup: any named dataset ---
	// Use ListLookup for any of the 30+ datasets:
	//   "ActivityCode", "LeaveReason", "LeaveType", "CompanyLegalBody",
	//   "AppealOutcome", "AwardFavour", "LeaveBucket", etc.
	fmt.Println("\n=== Lookup: Activity Codes ===")
	codes, _, err := client.Lookups.ListActivityCodes(ctx, companyID, nil)
	if err != nil {
		log.Fatalf("Failed to list activity codes: %v", err)
	}
	for _, code := range codes {
		fmt.Printf("  [%s] %s\n", code.Code, code.Description)
	}

	// Using the generic method directly:
	fmt.Println("\n=== Lookup: Leave Reasons (generic) ===")
	reasons, _, err := client.Lookups.ListLookup(ctx, companyID, "LeaveReason", nil)
	if err != nil {
		log.Fatalf("Failed to list leave reasons: %v", err)
	}
	for _, r := range reasons {
		fmt.Printf("  [%s] %s\n", r.Code, r.Description)
	}

	// --- Custom fields ---
	fmt.Println("\n=== Custom Fields for 'Employee' entity ===")
	fields, _, err := client.Lookups.ListCustomFields(ctx, companyID, "Employee", nil)
	if err != nil {
		log.Fatalf("Failed to list custom fields: %v", err)
	}
	for _, f := range fields {
		fmt.Printf("  %s (%s) - Caption: %s\n", f.FieldName, f.FieldType, f.Caption)

		// List dropdown options for this custom field.
		if f.FieldType == "Dropdown" {
			options, _, err := client.Lookups.ListCustomFieldOptions(ctx, companyID, f.CustomFieldId, nil)
			if err != nil {
				log.Printf("    Failed to list options: %v", err)
				continue
			}
			for _, opt := range options {
				fmt.Printf("    - [%s] %s\n", opt.Code, opt.Description)
			}
		}
	}
}

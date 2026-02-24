// Example: uploading a file and receiving a temporary attachment signature.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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

	// Upload from an in-memory reader (could be *os.File for real files).
	fileContent := strings.NewReader("Medical certificate content here...")
	filename := "medical-cert.pdf"

	fmt.Printf("Uploading %s...\n", filename)
	result, _, err := client.FileUpload.Upload(ctx, companyID, filename, fileContent)
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	fmt.Printf("Upload successful!\n")
	fmt.Printf("  TempAttachmentSignature: %s\n", result.TempAttachmentSignature)
	fmt.Println("\nUse this signature when creating entities that accept attachments,")
	fmt.Println("such as leave applications or custom forms.")

	// To upload an actual file from disk:
	//
	//   f, err := os.Open("/path/to/document.pdf")
	//   if err != nil {
	//       log.Fatal(err)
	//   }
	//   defer f.Close()
	//
	//   result, _, err := client.FileUpload.Upload(ctx, companyID, "document.pdf", f)
}

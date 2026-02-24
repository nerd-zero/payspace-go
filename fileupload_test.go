package payspace

import (
	"context"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"strings"
	"testing"
)

func TestFileUploadService_Upload(t *testing.T) {
	mux := http.NewServeMux()
	// FileUpload uses apiURL, not odataURL: {base}/v2.0/{companyID}/FileUpload
	mux.HandleFunc("POST /v2.0/1/FileUpload", func(w http.ResponseWriter, r *http.Request) {
		// Verify Content-Type is multipart/form-data.
		contentType := r.Header.Get("Content-Type")
		mediaType, params, err := mime.ParseMediaType(contentType)
		if err != nil {
			t.Errorf("failed to parse Content-Type: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if mediaType != "multipart/form-data" {
			t.Errorf("expected multipart/form-data, got %q", mediaType)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		boundary := params["boundary"]
		if boundary == "" {
			t.Error("missing multipart boundary")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Parse the multipart form to verify the file was received.
		reader := multipart.NewReader(r.Body, boundary)
		part, err := reader.NextPart()
		if err != nil {
			t.Errorf("failed to read multipart part: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if part.FormName() != "file" {
			t.Errorf("expected form field name 'file', got %q", part.FormName())
		}
		if part.FileName() != "document.pdf" {
			t.Errorf("expected filename 'document.pdf', got %q", part.FileName())
		}

		// Read file content.
		content, err := io.ReadAll(part)
		if err != nil {
			t.Errorf("failed to read file content: %v", err)
		}
		if string(content) != "fake-pdf-content" {
			t.Errorf("expected file content 'fake-pdf-content', got %q", string(content))
		}

		writeJSON(w, http.StatusOK, FileUploadResponse{
			TempAttachmentSignature: "abc-123-def-456",
		})
	})
	_, client := testServerAndClient(t, mux)

	fileData := strings.NewReader("fake-pdf-content")
	result, resp, err := client.FileUpload.Upload(context.Background(), 1, "document.pdf", fileData)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}
	if result.TempAttachmentSignature != "abc-123-def-456" {
		t.Errorf("expected TempAttachmentSignature='abc-123-def-456', got %q", result.TempAttachmentSignature)
	}
}

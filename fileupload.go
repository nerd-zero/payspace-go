package payspace

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"go.uber.org/zap"
)

// FileUploadResponse holds the result of a file upload operation.
type FileUploadResponse struct {
	TempAttachmentSignature string `json:"TempAttachmentSignature"`
}

// FileUploadService handles file upload API operations.
type FileUploadService struct {
	client *Client
}

// Upload uploads a file to PaySpace and returns a temporary attachment signature.
// The returned TempAttachmentSignature GUID can be used when creating leave applications
// or other entities that accept file attachments.
func (s *FileUploadService) Upload(ctx context.Context, companyID int, filename string, data io.Reader) (*FileUploadResponse, *Response, error) {
	s.client.logger.Info("uploading file",
		zap.Int("company_id", companyID),
		zap.String("filename", filename),
	)

	url := s.client.apiURL(companyID, "FileUpload")

	// Build multipart form data.
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, nil, fmt.Errorf("payspace: failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, data); err != nil {
		return nil, nil, fmt.Errorf("payspace: failed to copy file data: %w", err)
	}

	if err := writer.Close(); err != nil {
		return nil, nil, fmt.Errorf("payspace: failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, &buf)
	if err != nil {
		return nil, nil, fmt.Errorf("payspace: failed to create upload request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Accept", "application/json")

	resp, err := s.client.do(ctx, req)
	if err != nil {
		return nil, resp, err
	}

	var result FileUploadResponse
	if err := decodeJSON(resp, &result); err != nil {
		return nil, resp, err
	}

	s.client.logger.Debug("file uploaded successfully",
		zap.Int("company_id", companyID),
		zap.String("filename", filename),
		zap.String("signature", result.TempAttachmentSignature),
	)

	return &result, resp, nil
}

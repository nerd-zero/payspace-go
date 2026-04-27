package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeeAttachment represents an employee attachment record.
type EmployeeAttachment struct {
	AttachmentId          int    `json:"AttachmentId"`
	EmployeeNumber        string `json:"EmployeeNumber"`
	Classification        string `json:"Classification,omitempty"`
	AttachmentDescription string `json:"AttachmentDescription,omitempty"`
	AttachmentName        string `json:"AttachmentName,omitempty"`
	AttachmentUrl         string `json:"AttachmentUrl,omitempty"`
}

// EmployeeAttachmentCreateRequest is the payload for creating an employee attachment.
// Provide exactly one of:
//   - Attachment: base64-encoded file bytes (single-call upload, populates AttachmentUrl).
//   - TempAttachmentSignature: the GUID returned by FileUploadService.Upload
//     (two-call upload; leaves AttachmentUrl null on some tenants).
type EmployeeAttachmentCreateRequest struct {
	EmployeeNumber          string `json:"EmployeeNumber"`
	Classification          string `json:"Classification,omitempty"`
	AttachmentDescription   string `json:"AttachmentDescription"`
	AttachmentName          string `json:"AttachmentName"`
	Attachment              string `json:"Attachment,omitempty"`
	TempAttachmentSignature string `json:"TempAttachmentSignature,omitempty"`
}

// EmployeeAttachmentService handles employee attachment API operations.
type EmployeeAttachmentService struct {
	client *Client
}

// ListAttachments retrieves all employee attachments for a company.
func (s *EmployeeAttachmentService) ListAttachments(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeAttachment, *Response, error) {
	s.client.logger.Info("listing employee attachments", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list attachments query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeAttachment")
	var result ListResult[EmployeeAttachment]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateAttachment creates a new employee attachment record.
// The caller should first call FileUploadService.Upload to obtain the
// TempAttachmentSignature referenced in the request.
func (s *EmployeeAttachmentService) CreateAttachment(ctx context.Context, companyID int, req EmployeeAttachmentCreateRequest) (*EmployeeAttachment, *Response, error) {
	s.client.logger.Info("creating employee attachment",
		zap.Int("company_id", companyID),
		zap.String("employee_number", req.EmployeeNumber),
		zap.String("attachment_name", req.AttachmentName),
	)

	url := s.client.odataURL(companyID, "EmployeeAttachment")
	var result EmployeeAttachment
	resp, err := s.client.postAndDecode(ctx, url, req, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetAttachment retrieves a single employee attachment by ID.
func (s *EmployeeAttachmentService) GetAttachment(ctx context.Context, companyID, attachmentID int) (*EmployeeAttachment, *Response, error) {
	s.client.logger.Info("getting employee attachment", zap.Int("company_id", companyID), zap.Int("attachment_id", attachmentID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeAttachment"), attachmentID)
	var attachment EmployeeAttachment
	resp, err := s.client.getAndDecode(ctx, url, nil, &attachment)
	if err != nil {
		return nil, resp, err
	}
	return &attachment, resp, nil
}

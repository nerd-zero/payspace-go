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

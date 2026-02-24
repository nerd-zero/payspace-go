package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// WebhookError represents a single webhook error entry.
type WebhookError struct {
	EntityType string `json:"EntityType,omitempty"`
	ErrorDate  string `json:"ErrorDate,omitempty"`
	Message    string `json:"Message,omitempty"`
	Payload    string `json:"Payload,omitempty"`
}

// WebhookErrorResponse holds a paginated list of webhook errors.
type WebhookErrorResponse struct {
	Data       []WebhookError `json:"Data"`
	PageNumber int            `json:"PageNumber"`
	PageSize   int            `json:"PageSize"`
	TotalCount int            `json:"TotalCount"`
}

// WebhookService handles webhook-related API operations.
type WebhookService struct {
	client *Client
}

// ListErrors retrieves webhook errors for a given entity type and date range.
// This is NOT an OData endpoint; it uses the standard REST API with query parameters.
func (s *WebhookService) ListErrors(ctx context.Context, companyID int, entityType, from, to string, pageNumber, pageSize int) (*WebhookErrorResponse, *Response, error) {
	s.client.logger.Info("listing webhook errors",
		zap.Int("company_id", companyID),
		zap.String("entity_type", entityType),
	)
	s.client.logger.Debug("webhook errors query params",
		zap.String("from", from),
		zap.String("to", to),
		zap.Int("page_number", pageNumber),
		zap.Int("page_size", pageSize),
	)

	baseURL := s.client.apiURL(companyID, "WebhookError/"+entityType)
	url := addQueryParams(baseURL, map[string]string{
		"from":       from,
		"to":         to,
		"pageNumber": fmt.Sprintf("%d", pageNumber),
		"pageSize":   fmt.Sprintf("%d", pageSize),
	})

	var result WebhookErrorResponse
	resp, err := s.client.getAndDecode(ctx, url, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

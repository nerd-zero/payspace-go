package payspace

import (
	"context"
	"fmt"
	"io"

	"go.uber.org/zap"
)

// MetadataService handles OData metadata retrieval.
type MetadataService struct {
	client *Client
}

// Get retrieves the OData metadata document (XML) for a company.
// The response is raw XML, not JSON.
func (s *MetadataService) Get(ctx context.Context, companyID int) (string, *Response, error) {
	s.client.logger.Info("getting OData metadata", zap.Int("company_id", companyID))

	url := fmt.Sprintf("%s/odata/%s/%d/$metadata", s.client.apiBaseURL, s.client.apiVersion, companyID)
	resp, err := s.client.get(ctx, url, nil)
	if err != nil {
		return "", resp, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp, fmt.Errorf("payspace: failed to read metadata response: %w", err)
	}

	s.client.logger.Debug("retrieved OData metadata",
		zap.Int("company_id", companyID),
		zap.Int("body_length", len(data)),
	)

	return string(data), resp, nil
}

package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// CompanyFrequency represents a payroll frequency configuration for a company.
type CompanyFrequency struct {
	FrequencyId   int    `json:"FrequencyId"`
	Description   string `json:"Description,omitempty"`
	FrequencyCode string `json:"FrequencyCode,omitempty"`
	Periods       int    `json:"Periods,omitempty"`
}

// CompanyRun represents a payroll run within a company.
type CompanyRun struct {
	RunId       int    `json:"RunId"`
	FrequencyId int    `json:"FrequencyId,omitempty"`
	Period      int    `json:"Period,omitempty"`
	Year        int    `json:"Year,omitempty"`
	Status      string `json:"Status,omitempty"`
	Description string `json:"Description,omitempty"`
}

// ComponentCompany represents a payroll component configured at the company level.
type ComponentCompany struct {
	ComponentCompanyId int    `json:"ComponentCompanyId"`
	ComponentCode      string `json:"ComponentCode,omitempty"`
	Description        string `json:"Description,omitempty"`
	ComponentType      string `json:"ComponentType,omitempty"`
}

// CompanyPayrollService handles company payroll-related API operations.
type CompanyPayrollService struct {
	client *Client
}

// ListFrequencies retrieves all payroll frequencies for a company.
func (s *CompanyPayrollService) ListFrequencies(ctx context.Context, companyID int, params *QueryParams) ([]CompanyFrequency, *Response, error) {
	s.client.logger.Info("listing company frequencies", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list company frequencies query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "Lookup/CompanyFrequency")
	var result ListResult[CompanyFrequency]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListRuns retrieves all payroll runs for a company.
func (s *CompanyPayrollService) ListRuns(ctx context.Context, companyID int, params *QueryParams) ([]CompanyRun, *Response, error) {
	s.client.logger.Info("listing company runs", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list company runs query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "Lookup/CompanyRun")
	var result ListResult[CompanyRun]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListRunsByFrequency retrieves payroll runs filtered by frequency ID.
func (s *CompanyPayrollService) ListRunsByFrequency(ctx context.Context, companyID, frequencyID int, params *QueryParams) ([]CompanyRun, *Response, error) {
	s.client.logger.Info("listing company runs by frequency",
		zap.Int("company_id", companyID),
		zap.Int("frequency_id", frequencyID),
	)

	// Build a filter for the frequency ID and merge with any existing params.
	filterExpr := fmt.Sprintf("FrequencyId eq %d", frequencyID)
	if params == nil {
		params = NewQuery()
	}
	params = params.Filter(filterExpr)

	s.client.logger.Debug("list company runs by frequency query params", zap.String("params", params.Encode()))

	url := s.client.odataURL(companyID, "Lookup/CompanyRun")
	var result ListResult[CompanyRun]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListComponents retrieves company-level payroll components for a given frequency and period.
func (s *CompanyPayrollService) ListComponents(ctx context.Context, companyID, frequency, period int, params *QueryParams) ([]ComponentCompany, *Response, error) {
	s.client.logger.Info("listing company components",
		zap.Int("company_id", companyID),
		zap.Int("frequency", frequency),
		zap.Int("period", period),
	)
	if params != nil {
		s.client.logger.Debug("list company components query params", zap.String("params", params.Encode()))
	}

	baseURL := s.client.odataURL(companyID, "ComponentCompany")
	url := addQueryParams(baseURL, map[string]string{
		"frequency": fmt.Sprintf("%d", frequency),
		"period":    fmt.Sprintf("%d", period),
	})
	var result ListResult[ComponentCompany]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// UpdateComponent updates a company-level payroll component.
// Only ComponentCode and Description fields may be updated.
func (s *CompanyPayrollService) UpdateComponent(ctx context.Context, companyID, componentID, frequency, period int, component *ComponentCompany) (*Response, error) {
	s.client.logger.Info("updating company component",
		zap.Int("company_id", companyID),
		zap.Int("component_id", componentID),
		zap.Int("frequency", frequency),
		zap.Int("period", period),
	)

	baseURL := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "ComponentCompany"), componentID)
	url := addQueryParams(baseURL, map[string]string{
		"frequency": fmt.Sprintf("%d", frequency),
		"period":    fmt.Sprintf("%d", period),
	})
	return s.client.patch(ctx, url, component)
}

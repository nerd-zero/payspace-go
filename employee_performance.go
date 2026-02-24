package payspace

import (
	"context"

	"go.uber.org/zap"
)

// EmployeeReviewHeader represents a performance review header record.
type EmployeeReviewHeader struct {
	ReviewHeaderId int    `json:"ReviewHeaderId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	ReviewName     string `json:"ReviewName,omitempty"`
	ReviewDate     string `json:"ReviewDate,omitempty"`
	ReviewStatus   string `json:"ReviewStatus,omitempty"`
	ReviewerName   string `json:"ReviewerName,omitempty"`
	OverallScore   string `json:"OverallScore,omitempty"`
	OverallRating  string `json:"OverallRating,omitempty"`
}

// EmployeeReviewKpa represents a key performance area within a review.
type EmployeeReviewKpa struct {
	ReviewKpaId    int      `json:"ReviewKpaId"`
	ReviewHeaderId int      `json:"ReviewHeaderId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	KpaName        string   `json:"KpaName,omitempty"`
	Weight         *float64 `json:"Weight,omitempty"`
	Score          *float64 `json:"Score,omitempty"`
	Rating         string   `json:"Rating,omitempty"`
	Comment        string   `json:"Comment,omitempty"`
}

// EmployeeReviewTemplate represents a performance review template.
type EmployeeReviewTemplate struct {
	ReviewTemplateId int    `json:"ReviewTemplateId"`
	TemplateName     string `json:"TemplateName,omitempty"`
	Description      string `json:"Description,omitempty"`
	Status           string `json:"Status,omitempty"`
}

// EmployeePerformanceService handles employee performance review-related API operations.
// All endpoints use v1.0 and are read-only (GET only).
type EmployeePerformanceService struct {
	client *Client
}

// ListReviewHeaders retrieves all performance review headers using v1.0.
func (s *EmployeePerformanceService) ListReviewHeaders(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeReviewHeader, *Response, error) {
	s.client.logger.Info("listing review headers", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list review headers query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURLv1(companyID, "EmployeeReviewHeader")
	var result ListResult[EmployeeReviewHeader]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListReviewKPAs retrieves all key performance areas within reviews using v1.0.
func (s *EmployeePerformanceService) ListReviewKPAs(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeReviewKpa, *Response, error) {
	s.client.logger.Info("listing review KPAs", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list review KPAs query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURLv1(companyID, "EmployeeReviewKpa")
	var result ListResult[EmployeeReviewKpa]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListReviewTemplates retrieves all performance review templates using v1.0.
func (s *EmployeePerformanceService) ListReviewTemplates(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeReviewTemplate, *Response, error) {
	s.client.logger.Info("listing review templates", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list review templates query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURLv1(companyID, "EmployeeReviewTemplate")
	var result ListResult[EmployeeReviewTemplate]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

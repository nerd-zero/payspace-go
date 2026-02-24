package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeeSuspension represents an employee suspension record.
type EmployeeSuspension struct {
	SuspensionId     int    `json:"SuspensionId"`
	EmployeeNumber   string `json:"EmployeeNumber"`
	SuspensionDate   string `json:"SuspensionDate,omitempty"`
	SuspensionReason string `json:"SuspensionReason,omitempty"`
	SuspensionStatus string `json:"SuspensionStatus,omitempty"`
	ReturnDate       string `json:"ReturnDate,omitempty"`
	PaymentMethod    string `json:"PaymentMethod,omitempty"`
}

// EmployeeSuspensionService handles employee suspension-related API operations.
type EmployeeSuspensionService struct {
	client *Client
}

// List retrieves all employee suspensions.
func (s *EmployeeSuspensionService) List(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeSuspension, *Response, error) {
	s.client.logger.Info("listing employee suspensions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list suspensions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeSuspension")
	var result ListResult[EmployeeSuspension]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// Get retrieves a single suspension by ID.
func (s *EmployeeSuspensionService) Get(ctx context.Context, companyID, suspensionID int) (*EmployeeSuspension, *Response, error) {
	s.client.logger.Info("getting employee suspension", zap.Int("company_id", companyID), zap.Int("suspension_id", suspensionID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeSuspension"), suspensionID)
	var suspension EmployeeSuspension
	resp, err := s.client.getAndDecode(ctx, url, nil, &suspension)
	if err != nil {
		return nil, resp, err
	}
	return &suspension, resp, nil
}

// ListAll retrieves all employee suspensions including inactive ones.
func (s *EmployeeSuspensionService) ListAll(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeSuspension, *Response, error) {
	s.client.logger.Info("listing all employee suspensions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all suspensions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeSuspension/all")
	var result ListResult[EmployeeSuspension]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// Create creates a new employee suspension.
func (s *EmployeeSuspensionService) Create(ctx context.Context, companyID int, suspension *EmployeeSuspension) (*EmployeeSuspension, *Response, error) {
	s.client.logger.Info("creating employee suspension", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeSuspension")
	var created EmployeeSuspension
	resp, err := s.client.postAndDecode(ctx, url, suspension, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// Update updates an existing employee suspension.
func (s *EmployeeSuspensionService) Update(ctx context.Context, companyID, suspensionID int, suspension *EmployeeSuspension) (*Response, error) {
	s.client.logger.Info("updating employee suspension", zap.Int("company_id", companyID), zap.Int("suspension_id", suspensionID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeSuspension"), suspensionID)
	return s.client.patch(ctx, url, suspension)
}

// Delete deletes an employee suspension record.
func (s *EmployeeSuspensionService) Delete(ctx context.Context, companyID, suspensionID int) (*Response, error) {
	s.client.logger.Info("deleting employee suspension", zap.Int("company_id", companyID), zap.Int("suspension_id", suspensionID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeSuspension"), suspensionID)
	return s.client.del(ctx, url)
}

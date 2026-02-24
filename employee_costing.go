package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeeRecurringCostingSplit represents a costing split record.
type EmployeeRecurringCostingSplit struct {
	CostingSplitId int      `json:"CostingSplitId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	CostCentre     string   `json:"CostCentre,omitempty"`
	Percentage     *float64 `json:"Percentage,omitempty"`
	EffectiveDate  string   `json:"EffectiveDate,omitempty"`
	ComponentCode  string   `json:"ComponentCode,omitempty"`
	Description    string   `json:"Description,omitempty"`
}

// EmployeeRecurringCosting represents a recurring costing record.
type EmployeeRecurringCosting struct {
	RecurringCostingId int      `json:"RecurringCostingId"`
	EmployeeNumber     string   `json:"EmployeeNumber"`
	CostCentre         string   `json:"CostCentre,omitempty"`
	Percentage         *float64 `json:"Percentage,omitempty"`
	EffectiveDate      string   `json:"EffectiveDate,omitempty"`
	ComponentCode      string   `json:"ComponentCode,omitempty"`
	Description        string   `json:"Description,omitempty"`
}

// EmployeeCostingService handles employee costing-related API operations.
type EmployeeCostingService struct {
	client *Client
}

// --- Recurring Costing Split methods (read only) ---

// ListCostingSplits retrieves all costing split records.
func (s *EmployeeCostingService) ListCostingSplits(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeRecurringCostingSplit, *Response, error) {
	s.client.logger.Info("listing costing splits", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list costing splits query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeRecurringCostingSplit")
	var result ListResult[EmployeeRecurringCostingSplit]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListCostingSplitsByDate retrieves costing splits effective as of the given date.
func (s *EmployeeCostingService) ListCostingSplitsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeeRecurringCostingSplit, *Response, error) {
	s.client.logger.Info("listing costing splits by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list costing splits by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeRecurringCostingSplit/effective/%s", date))
	var result ListResult[EmployeeRecurringCostingSplit]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllCostingSplits retrieves all costing splits including inactive ones.
func (s *EmployeeCostingService) ListAllCostingSplits(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeRecurringCostingSplit, *Response, error) {
	s.client.logger.Info("listing all costing splits", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all costing splits query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeRecurringCostingSplit/all")
	var result ListResult[EmployeeRecurringCostingSplit]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// --- Recurring Costing methods (full CRUD) ---

// ListRecurringCostings retrieves all recurring costing records.
func (s *EmployeeCostingService) ListRecurringCostings(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeRecurringCosting, *Response, error) {
	s.client.logger.Info("listing recurring costings", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list recurring costings query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeRecurringCosting")
	var result ListResult[EmployeeRecurringCosting]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetRecurringCosting retrieves a single recurring costing by ID.
func (s *EmployeeCostingService) GetRecurringCosting(ctx context.Context, companyID, costingID int) (*EmployeeRecurringCosting, *Response, error) {
	s.client.logger.Info("getting recurring costing", zap.Int("company_id", companyID), zap.Int("costing_id", costingID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeRecurringCosting"), costingID)
	var costing EmployeeRecurringCosting
	resp, err := s.client.getAndDecode(ctx, url, nil, &costing)
	if err != nil {
		return nil, resp, err
	}
	return &costing, resp, nil
}

// ListRecurringCostingsByDate retrieves recurring costings effective as of the given date.
func (s *EmployeeCostingService) ListRecurringCostingsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeeRecurringCosting, *Response, error) {
	s.client.logger.Info("listing recurring costings by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list recurring costings by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeRecurringCosting/effective/%s", date))
	var result ListResult[EmployeeRecurringCosting]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllRecurringCostings retrieves all recurring costings including inactive ones.
func (s *EmployeeCostingService) ListAllRecurringCostings(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeRecurringCosting, *Response, error) {
	s.client.logger.Info("listing all recurring costings", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all recurring costings query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeRecurringCosting/all")
	var result ListResult[EmployeeRecurringCosting]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateRecurringCosting creates a new recurring costing record.
func (s *EmployeeCostingService) CreateRecurringCosting(ctx context.Context, companyID int, costing *EmployeeRecurringCosting) (*EmployeeRecurringCosting, *Response, error) {
	s.client.logger.Info("creating recurring costing", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeRecurringCosting")
	var created EmployeeRecurringCosting
	resp, err := s.client.postAndDecode(ctx, url, costing, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateRecurringCosting updates an existing recurring costing record.
func (s *EmployeeCostingService) UpdateRecurringCosting(ctx context.Context, companyID, costingID int, costing *EmployeeRecurringCosting) (*Response, error) {
	s.client.logger.Info("updating recurring costing", zap.Int("company_id", companyID), zap.Int("costing_id", costingID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeRecurringCosting"), costingID)
	return s.client.patch(ctx, url, costing)
}

// DeleteRecurringCosting deletes a recurring costing record.
func (s *EmployeeCostingService) DeleteRecurringCosting(ctx context.Context, companyID, costingID int) (*Response, error) {
	s.client.logger.Info("deleting recurring costing", zap.Int("company_id", companyID), zap.Int("costing_id", costingID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeRecurringCosting"), costingID)
	return s.client.del(ctx, url)
}

// DeleteRecurringCostingDetail deletes a specific detail from a recurring costing record.
func (s *EmployeeCostingService) DeleteRecurringCostingDetail(ctx context.Context, companyID, costingID, detailID int) (*Response, error) {
	s.client.logger.Info("deleting recurring costing detail",
		zap.Int("company_id", companyID),
		zap.Int("costing_id", costingID),
		zap.Int("detail_id", detailID),
	)

	url := fmt.Sprintf("%s(%d)/Detail(%d)", s.client.odataURL(companyID, "EmployeeRecurringCosting"), costingID, detailID)
	return s.client.del(ctx, url)
}

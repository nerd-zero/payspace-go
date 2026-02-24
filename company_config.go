package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// OrganizationPositionDetail represents a position within the company's organization structure.
type OrganizationPositionDetail struct {
	OrganizationPositionDetailId int    `json:"OrganizationPositionDetailId"`
	PositionCode                 string `json:"PositionCode,omitempty"`
	PositionDescription          string `json:"PositionDescription,omitempty"`
	Department                   string `json:"Department,omitempty"`
	EffectiveDate                string `json:"EffectiveDate,omitempty"`
}

// CompanyRegion represents a region configured for a company.
type CompanyRegion struct {
	RegionId    int    `json:"RegionId"`
	Description string `json:"Description,omitempty"`
	RegionCode  string `json:"RegionCode,omitempty"`
}

// GeneralLedger represents a general ledger entry for a company.
type GeneralLedger struct {
	GeneralLedgerId int      `json:"GeneralLedgerId"`
	AccountCode     string   `json:"AccountCode,omitempty"`
	Description     string   `json:"Description,omitempty"`
	Amount          *float64 `json:"Amount,omitempty"`
}

// EFTOutbox represents an EFT outbox record for electronic fund transfers.
type EFTOutbox struct {
	EftOutboxId    int      `json:"EftOutboxId"`
	EmployeeNumber string   `json:"EmployeeNumber,omitempty"`
	BankName       string   `json:"BankName,omitempty"`
	AccountNumber  string   `json:"AccountNumber,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	Status         string   `json:"Status,omitempty"`
}

// CostingProjectActivity represents a project activity used in costing.
type CostingProjectActivity struct {
	ProjectActivityId int    `json:"ProjectActivityId"`
	ProjectCode       string `json:"ProjectCode,omitempty"`
	ActivityCode      string `json:"ActivityCode,omitempty"`
	Description       string `json:"Description,omitempty"`
}

// ComponentTableDefinition represents a component table definition.
type ComponentTableDefinition struct {
	TableDefinitionId int    `json:"TableDefinitionId"`
	TableName         string `json:"TableName,omitempty"`
	Description       string `json:"Description,omitempty"`
}

// GeneralLedgerExtractResponse represents the response from starting a GL extraction.
type GeneralLedgerExtractResponse struct {
	ExtractionId int    `json:"ExtractionId"`
	Status       string `json:"Status,omitempty"`
}

// GeneralLedgerExtractStatus represents the status of a GL extraction.
type GeneralLedgerExtractStatus struct {
	ExtractionId int    `json:"ExtractionId"`
	Status       string `json:"Status,omitempty"`
}

// CompanyConfigService handles company configuration API operations.
type CompanyConfigService struct {
	client *Client
}

// --- Organization Positions ---

// ListOrganizationPositions retrieves all organization position details.
func (s *CompanyConfigService) ListOrganizationPositions(ctx context.Context, companyID int, params *QueryParams) ([]OrganizationPositionDetail, *Response, error) {
	s.client.logger.Info("listing organization positions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list organization positions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "OrganizationPositionDetail")
	var result ListResult[OrganizationPositionDetail]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetOrganizationPosition retrieves a single organization position by ID.
func (s *CompanyConfigService) GetOrganizationPosition(ctx context.Context, companyID, positionID int) (*OrganizationPositionDetail, *Response, error) {
	s.client.logger.Info("getting organization position",
		zap.Int("company_id", companyID),
		zap.Int("position_id", positionID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "OrganizationPositionDetail"), positionID)
	var position OrganizationPositionDetail
	resp, err := s.client.getAndDecode(ctx, url, nil, &position)
	if err != nil {
		return nil, resp, err
	}
	return &position, resp, nil
}

// ListOrganizationPositionsByDate retrieves organization positions effective as of the given date.
func (s *CompanyConfigService) ListOrganizationPositionsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]OrganizationPositionDetail, *Response, error) {
	s.client.logger.Info("listing organization positions by date",
		zap.Int("company_id", companyID),
		zap.String("date", date),
	)
	if params != nil {
		s.client.logger.Debug("list organization positions by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("OrganizationPositionDetail/effective/%s", date))
	var result ListResult[OrganizationPositionDetail]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllOrganizationPositions retrieves all organization positions including inactive ones.
func (s *CompanyConfigService) ListAllOrganizationPositions(ctx context.Context, companyID int, params *QueryParams) ([]OrganizationPositionDetail, *Response, error) {
	s.client.logger.Info("listing all organization positions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all organization positions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "OrganizationPositionDetail/all")
	var result ListResult[OrganizationPositionDetail]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateOrganizationPosition creates a new organization position.
func (s *CompanyConfigService) CreateOrganizationPosition(ctx context.Context, companyID int, position *OrganizationPositionDetail) (*OrganizationPositionDetail, *Response, error) {
	s.client.logger.Info("creating organization position", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "OrganizationPositionDetail")
	var created OrganizationPositionDetail
	resp, err := s.client.postAndDecode(ctx, url, position, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateOrganizationPosition updates an existing organization position.
func (s *CompanyConfigService) UpdateOrganizationPosition(ctx context.Context, companyID, positionID int, position *OrganizationPositionDetail) (*Response, error) {
	s.client.logger.Info("updating organization position",
		zap.Int("company_id", companyID),
		zap.Int("position_id", positionID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "OrganizationPositionDetail"), positionID)
	return s.client.patch(ctx, url, position)
}

// DeleteOrganizationPosition deletes an organization position.
func (s *CompanyConfigService) DeleteOrganizationPosition(ctx context.Context, companyID, positionID int) (*Response, error) {
	s.client.logger.Info("deleting organization position",
		zap.Int("company_id", companyID),
		zap.Int("position_id", positionID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "OrganizationPositionDetail"), positionID)
	return s.client.del(ctx, url)
}

// --- Regions ---

// ListRegions retrieves all regions configured for a company (read-only).
func (s *CompanyConfigService) ListRegions(ctx context.Context, companyID int, params *QueryParams) ([]CompanyRegion, *Response, error) {
	s.client.logger.Info("listing company regions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list company regions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "Lookup/CompanyRegion")
	var result ListResult[CompanyRegion]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// --- General Ledger ---

// ListGeneralLedger retrieves general ledger entries for a given year and month.
func (s *CompanyConfigService) ListGeneralLedger(ctx context.Context, companyID, year, month int, params *QueryParams) ([]GeneralLedger, *Response, error) {
	s.client.logger.Info("listing general ledger",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list general ledger query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("GeneralLedger/%d/%d", year, month))
	var result ListResult[GeneralLedger]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ExtractGeneralLedger starts a general ledger extraction for a given year and month.
func (s *CompanyConfigService) ExtractGeneralLedger(ctx context.Context, companyID, year, month int) (*GeneralLedgerExtractResponse, *Response, error) {
	s.client.logger.Info("extracting general ledger",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)

	url := s.client.odataURL(companyID, fmt.Sprintf("GeneralLedger/extract/start/%d/%d", year, month))
	var result GeneralLedgerExtractResponse
	resp, err := s.client.postAndDecode(ctx, url, nil, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// ExtractGeneralLedgerStatus retrieves the status of a general ledger extraction.
func (s *CompanyConfigService) ExtractGeneralLedgerStatus(ctx context.Context, companyID, extractionID int) (*GeneralLedgerExtractStatus, *Response, error) {
	s.client.logger.Info("getting general ledger extraction status",
		zap.Int("company_id", companyID),
		zap.Int("extraction_id", extractionID),
	)

	url := s.client.odataURL(companyID, fmt.Sprintf("GeneralLedger/extract/%d/status", extractionID))
	var status GeneralLedgerExtractStatus
	resp, err := s.client.getAndDecode(ctx, url, nil, &status)
	if err != nil {
		return nil, resp, err
	}
	return &status, resp, nil
}

// --- EFT Outbox ---

// ListEFTOutbox retrieves EFT outbox records for a given year and month.
func (s *CompanyConfigService) ListEFTOutbox(ctx context.Context, companyID, year, month int, params *QueryParams) ([]EFTOutbox, *Response, error) {
	s.client.logger.Info("listing EFT outbox",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list EFT outbox query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EFTOutbox/%d/%d", year, month))
	var result ListResult[EFTOutbox]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// --- Costing Project Activity ---

// ListCostingProjectActivities retrieves all costing project activities.
func (s *CompanyConfigService) ListCostingProjectActivities(ctx context.Context, companyID int, params *QueryParams) ([]CostingProjectActivity, *Response, error) {
	s.client.logger.Info("listing costing project activities", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list costing project activities query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "CostingProjectActivity")
	var result ListResult[CostingProjectActivity]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetCostingProjectActivity retrieves a single costing project activity by ID.
func (s *CompanyConfigService) GetCostingProjectActivity(ctx context.Context, companyID, activityID int) (*CostingProjectActivity, *Response, error) {
	s.client.logger.Info("getting costing project activity",
		zap.Int("company_id", companyID),
		zap.Int("activity_id", activityID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "CostingProjectActivity"), activityID)
	var activity CostingProjectActivity
	resp, err := s.client.getAndDecode(ctx, url, nil, &activity)
	if err != nil {
		return nil, resp, err
	}
	return &activity, resp, nil
}

// CreateCostingProjectActivity creates a new costing project activity.
func (s *CompanyConfigService) CreateCostingProjectActivity(ctx context.Context, companyID int, activity *CostingProjectActivity) (*CostingProjectActivity, *Response, error) {
	s.client.logger.Info("creating costing project activity", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "CostingProjectActivity")
	var created CostingProjectActivity
	resp, err := s.client.postAndDecode(ctx, url, activity, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateCostingProjectActivity updates an existing costing project activity.
func (s *CompanyConfigService) UpdateCostingProjectActivity(ctx context.Context, companyID, activityID int, activity *CostingProjectActivity) (*Response, error) {
	s.client.logger.Info("updating costing project activity",
		zap.Int("company_id", companyID),
		zap.Int("activity_id", activityID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "CostingProjectActivity"), activityID)
	return s.client.patch(ctx, url, activity)
}

// DeleteCostingProjectActivity deletes a costing project activity.
func (s *CompanyConfigService) DeleteCostingProjectActivity(ctx context.Context, companyID, activityID int) (*Response, error) {
	s.client.logger.Info("deleting costing project activity",
		zap.Int("company_id", companyID),
		zap.Int("activity_id", activityID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "CostingProjectActivity"), activityID)
	return s.client.del(ctx, url)
}

// --- Component Table Definition ---

// ListComponentTableDefinitions retrieves all component table definitions.
func (s *CompanyConfigService) ListComponentTableDefinitions(ctx context.Context, companyID int, params *QueryParams) ([]ComponentTableDefinition, *Response, error) {
	s.client.logger.Info("listing component table definitions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list component table definitions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "ComponentTableDefinition")
	var result ListResult[ComponentTableDefinition]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

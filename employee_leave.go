package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeeLeaveAdjustment represents a leave adjustment record.
type EmployeeLeaveAdjustment struct {
	LeaveAdjustmentId int      `json:"LeaveAdjustmentId"`
	EmployeeNumber    string   `json:"EmployeeNumber"`
	LeaveType         string   `json:"LeaveType,omitempty"`
	AdjustmentDate    string   `json:"AdjustmentDate,omitempty"`
	AdjustmentValue   *float64 `json:"AdjustmentValue,omitempty"`
	Reason            string   `json:"Reason,omitempty"`
}

// EmployeeLeaveSetup represents a leave setup configuration for an employee.
type EmployeeLeaveSetup struct {
	LeaveSetupId   int      `json:"LeaveSetupId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	LeaveType      string   `json:"LeaveType,omitempty"`
	Entitlement    *float64 `json:"Entitlement,omitempty"`
	Balance        *float64 `json:"Balance,omitempty"`
	AccrualMethod  string   `json:"AccrualMethod,omitempty"`
	CycleStartDate string   `json:"CycleStartDate,omitempty"`
	CycleEndDate   string   `json:"CycleEndDate,omitempty"`
}

// EmployeeLeaveSetupEntitlement represents a leave setup entitlement record.
type EmployeeLeaveSetupEntitlement struct {
	LeaveSetupEntitlementId int      `json:"LeaveSetupEntitlementId"`
	EmployeeNumber          string   `json:"EmployeeNumber"`
	LeaveType               string   `json:"LeaveType,omitempty"`
	Entitlement             *float64 `json:"Entitlement,omitempty"`
	EffectiveDate           string   `json:"EffectiveDate,omitempty"`
}

// EmployeeLeaveApplication represents a leave application record.
type EmployeeLeaveApplication struct {
	LeaveAdjustmentId    int      `json:"LeaveAdjustmentId"`
	EmployeeNumber       string   `json:"EmployeeNumber"`
	FullName             string   `json:"FullName,omitempty"`
	LeaveBucket          string   `json:"LeaveBucket,omitempty"`
	LeaveType            string   `json:"LeaveType,omitempty"`
	LeaveCompanyRun      string   `json:"LeaveCompanyRun,omitempty"`
	NoOfDays             *float64 `json:"NoOfDays,omitempty"`
	Comments             string   `json:"Comments,omitempty"`
	Reference            string   `json:"Reference,omitempty"`
	LeaveReason          string   `json:"LeaveReason,omitempty"`
	LeaveStartDate       string   `json:"LeaveStartDate,omitempty"`
	LeaveEndDate         string   `json:"LeaveEndDate,omitempty"`
	LeaveStatus          string   `json:"LeaveStatus,omitempty"`
	LeaveTransactionType string   `json:"LeaveTransactionType,omitempty"`
}

// EmployeeLeaveService handles leave-related API operations.
type EmployeeLeaveService struct {
	client *Client
}

// --- Leave Adjustment methods ---

// ListLeaveAdjustments retrieves all leave adjustments.
func (s *EmployeeLeaveService) ListLeaveAdjustments(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveAdjustment, *Response, error) {
	s.client.logger.Info("listing leave adjustments", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list leave adjustments query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveAdjustment")
	var result ListResult[EmployeeLeaveAdjustment]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetLeaveAdjustment retrieves a single leave adjustment by ID.
func (s *EmployeeLeaveService) GetLeaveAdjustment(ctx context.Context, companyID, adjustmentID int) (*EmployeeLeaveAdjustment, *Response, error) {
	s.client.logger.Info("getting leave adjustment", zap.Int("company_id", companyID), zap.Int("adjustment_id", adjustmentID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveAdjustment"), adjustmentID)
	var adjustment EmployeeLeaveAdjustment
	resp, err := s.client.getAndDecode(ctx, url, nil, &adjustment)
	if err != nil {
		return nil, resp, err
	}
	return &adjustment, resp, nil
}

// ListLeaveAdjustmentsByDate retrieves leave adjustments effective as of the given date.
func (s *EmployeeLeaveService) ListLeaveAdjustmentsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeeLeaveAdjustment, *Response, error) {
	s.client.logger.Info("listing leave adjustments by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list leave adjustments by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeLeaveAdjustment/effective/%s", date))
	var result ListResult[EmployeeLeaveAdjustment]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllLeaveAdjustments retrieves all leave adjustments including inactive ones.
func (s *EmployeeLeaveService) ListAllLeaveAdjustments(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveAdjustment, *Response, error) {
	s.client.logger.Info("listing all leave adjustments", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all leave adjustments query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveAdjustment/all")
	var result ListResult[EmployeeLeaveAdjustment]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateLeaveAdjustment creates a new leave adjustment.
func (s *EmployeeLeaveService) CreateLeaveAdjustment(ctx context.Context, companyID int, adjustment *EmployeeLeaveAdjustment) (*EmployeeLeaveAdjustment, *Response, error) {
	s.client.logger.Info("creating leave adjustment", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeLeaveAdjustment")
	var created EmployeeLeaveAdjustment
	resp, err := s.client.postAndDecode(ctx, url, adjustment, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateLeaveAdjustment updates an existing leave adjustment.
func (s *EmployeeLeaveService) UpdateLeaveAdjustment(ctx context.Context, companyID, adjustmentID int, adjustment *EmployeeLeaveAdjustment) (*Response, error) {
	s.client.logger.Info("updating leave adjustment", zap.Int("company_id", companyID), zap.Int("adjustment_id", adjustmentID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveAdjustment"), adjustmentID)
	return s.client.patch(ctx, url, adjustment)
}

// DeleteLeaveAdjustment deletes a leave adjustment record.
func (s *EmployeeLeaveService) DeleteLeaveAdjustment(ctx context.Context, companyID, adjustmentID int) (*Response, error) {
	s.client.logger.Info("deleting leave adjustment", zap.Int("company_id", companyID), zap.Int("adjustment_id", adjustmentID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveAdjustment"), adjustmentID)
	return s.client.del(ctx, url)
}

// --- Leave Setup methods ---

// ListLeaveSetups retrieves all leave setups.
func (s *EmployeeLeaveService) ListLeaveSetups(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveSetup, *Response, error) {
	s.client.logger.Info("listing leave setups", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list leave setups query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveSetup")
	var result ListResult[EmployeeLeaveSetup]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetLeaveSetup retrieves a single leave setup by ID.
func (s *EmployeeLeaveService) GetLeaveSetup(ctx context.Context, companyID, setupID int) (*EmployeeLeaveSetup, *Response, error) {
	s.client.logger.Info("getting leave setup", zap.Int("company_id", companyID), zap.Int("setup_id", setupID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveSetup"), setupID)
	var setup EmployeeLeaveSetup
	resp, err := s.client.getAndDecode(ctx, url, nil, &setup)
	if err != nil {
		return nil, resp, err
	}
	return &setup, resp, nil
}

// ListLeaveSetupsByDate retrieves leave setups effective as of the given date.
func (s *EmployeeLeaveService) ListLeaveSetupsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeeLeaveSetup, *Response, error) {
	s.client.logger.Info("listing leave setups by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list leave setups by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeLeaveSetup/effective/%s", date))
	var result ListResult[EmployeeLeaveSetup]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllLeaveSetups retrieves all leave setups including inactive ones.
func (s *EmployeeLeaveService) ListAllLeaveSetups(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveSetup, *Response, error) {
	s.client.logger.Info("listing all leave setups", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all leave setups query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveSetup/all")
	var result ListResult[EmployeeLeaveSetup]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateLeaveSetup creates a new leave setup.
func (s *EmployeeLeaveService) CreateLeaveSetup(ctx context.Context, companyID int, setup *EmployeeLeaveSetup) (*EmployeeLeaveSetup, *Response, error) {
	s.client.logger.Info("creating leave setup", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeLeaveSetup")
	var created EmployeeLeaveSetup
	resp, err := s.client.postAndDecode(ctx, url, setup, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateLeaveSetup updates an existing leave setup.
func (s *EmployeeLeaveService) UpdateLeaveSetup(ctx context.Context, companyID, setupID int, setup *EmployeeLeaveSetup) (*Response, error) {
	s.client.logger.Info("updating leave setup", zap.Int("company_id", companyID), zap.Int("setup_id", setupID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveSetup"), setupID)
	return s.client.patch(ctx, url, setup)
}

// DeleteLeaveSetup deletes a leave setup record.
func (s *EmployeeLeaveService) DeleteLeaveSetup(ctx context.Context, companyID, setupID int) (*Response, error) {
	s.client.logger.Info("deleting leave setup", zap.Int("company_id", companyID), zap.Int("setup_id", setupID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveSetup"), setupID)
	return s.client.del(ctx, url)
}

// --- Leave Setup Entitlement methods ---

// ListLeaveSetupEntitlements retrieves all leave setup entitlements.
func (s *EmployeeLeaveService) ListLeaveSetupEntitlements(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveSetupEntitlement, *Response, error) {
	s.client.logger.Info("listing leave setup entitlements", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list leave setup entitlements query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveSetupEntitlement")
	var result ListResult[EmployeeLeaveSetupEntitlement]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetLeaveSetupEntitlement retrieves a single leave setup entitlement by ID.
func (s *EmployeeLeaveService) GetLeaveSetupEntitlement(ctx context.Context, companyID, entitlementID int) (*EmployeeLeaveSetupEntitlement, *Response, error) {
	s.client.logger.Info("getting leave setup entitlement", zap.Int("company_id", companyID), zap.Int("entitlement_id", entitlementID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveSetupEntitlement"), entitlementID)
	var entitlement EmployeeLeaveSetupEntitlement
	resp, err := s.client.getAndDecode(ctx, url, nil, &entitlement)
	if err != nil {
		return nil, resp, err
	}
	return &entitlement, resp, nil
}

// ListLeaveSetupEntitlementsByDate retrieves leave setup entitlements effective as of the given date.
func (s *EmployeeLeaveService) ListLeaveSetupEntitlementsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeeLeaveSetupEntitlement, *Response, error) {
	s.client.logger.Info("listing leave setup entitlements by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list leave setup entitlements by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeLeaveSetupEntitlement/effective/%s", date))
	var result ListResult[EmployeeLeaveSetupEntitlement]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllLeaveSetupEntitlements retrieves all leave setup entitlements including inactive ones.
func (s *EmployeeLeaveService) ListAllLeaveSetupEntitlements(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveSetupEntitlement, *Response, error) {
	s.client.logger.Info("listing all leave setup entitlements", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all leave setup entitlements query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveSetupEntitlement/all")
	var result ListResult[EmployeeLeaveSetupEntitlement]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateLeaveSetupEntitlement creates a new leave setup entitlement.
func (s *EmployeeLeaveService) CreateLeaveSetupEntitlement(ctx context.Context, companyID int, entitlement *EmployeeLeaveSetupEntitlement) (*EmployeeLeaveSetupEntitlement, *Response, error) {
	s.client.logger.Info("creating leave setup entitlement", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeLeaveSetupEntitlement")
	var created EmployeeLeaveSetupEntitlement
	resp, err := s.client.postAndDecode(ctx, url, entitlement, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateLeaveSetupEntitlement updates an existing leave setup entitlement.
func (s *EmployeeLeaveService) UpdateLeaveSetupEntitlement(ctx context.Context, companyID, entitlementID int, entitlement *EmployeeLeaveSetupEntitlement) (*Response, error) {
	s.client.logger.Info("updating leave setup entitlement", zap.Int("company_id", companyID), zap.Int("entitlement_id", entitlementID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveSetupEntitlement"), entitlementID)
	return s.client.patch(ctx, url, entitlement)
}

// DeleteLeaveSetupEntitlement deletes a leave setup entitlement record.
func (s *EmployeeLeaveService) DeleteLeaveSetupEntitlement(ctx context.Context, companyID, entitlementID int) (*Response, error) {
	s.client.logger.Info("deleting leave setup entitlement", zap.Int("company_id", companyID), zap.Int("entitlement_id", entitlementID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveSetupEntitlement"), entitlementID)
	return s.client.del(ctx, url)
}

// --- Leave Application methods ---

// ListLeaveApplications retrieves all leave applications.
func (s *EmployeeLeaveService) ListLeaveApplications(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveApplication, *Response, error) {
	s.client.logger.Info("listing leave applications", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list leave applications query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveApplication")
	var result ListResult[EmployeeLeaveApplication]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetLeaveApplication retrieves a single leave application by ID.
func (s *EmployeeLeaveService) GetLeaveApplication(ctx context.Context, companyID, applicationID int) (*EmployeeLeaveApplication, *Response, error) {
	s.client.logger.Info("getting leave application", zap.Int("company_id", companyID), zap.Int("application_id", applicationID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveApplication"), applicationID)
	var application EmployeeLeaveApplication
	resp, err := s.client.getAndDecode(ctx, url, nil, &application)
	if err != nil {
		return nil, resp, err
	}
	return &application, resp, nil
}

// ListLeaveApplicationsByDate retrieves leave applications effective as of the given date.
func (s *EmployeeLeaveService) ListLeaveApplicationsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeeLeaveApplication, *Response, error) {
	s.client.logger.Info("listing leave applications by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list leave applications by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeLeaveApplication/effective/%s", date))
	var result ListResult[EmployeeLeaveApplication]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllLeaveApplications retrieves all leave applications including inactive ones.
func (s *EmployeeLeaveService) ListAllLeaveApplications(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeLeaveApplication, *Response, error) {
	s.client.logger.Info("listing all leave applications", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all leave applications query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLeaveApplication/all")
	var result ListResult[EmployeeLeaveApplication]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateLeaveApplication creates a new leave application.
func (s *EmployeeLeaveService) CreateLeaveApplication(ctx context.Context, companyID int, application *EmployeeLeaveApplication) (*EmployeeLeaveApplication, *Response, error) {
	s.client.logger.Info("creating leave application", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeLeaveApplication")
	var created EmployeeLeaveApplication
	resp, err := s.client.postAndDecode(ctx, url, application, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateLeaveApplication updates an existing leave application.
func (s *EmployeeLeaveService) UpdateLeaveApplication(ctx context.Context, companyID, applicationID int, application *EmployeeLeaveApplication) (*Response, error) {
	s.client.logger.Info("updating leave application", zap.Int("company_id", companyID), zap.Int("application_id", applicationID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveApplication"), applicationID)
	return s.client.patch(ctx, url, application)
}

// DeleteLeaveApplication deletes a leave application record.
func (s *EmployeeLeaveService) DeleteLeaveApplication(ctx context.Context, companyID, applicationID int) (*Response, error) {
	s.client.logger.Info("deleting leave application", zap.Int("company_id", companyID), zap.Int("application_id", applicationID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLeaveApplication"), applicationID)
	return s.client.del(ctx, url)
}

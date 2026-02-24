package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// Employee represents an employee record in PaySpace.
type Employee struct {
	EmployeeId     int    `json:"EmployeeId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	Title          string `json:"Title,omitempty"`
	FirstName      string `json:"FirstName,omitempty"`
	Surname        string `json:"Surname,omitempty"`
	KnownAs        string `json:"KnownAs,omitempty"`
	Initials       string `json:"Initials,omitempty"`
	DateOfBirth    string `json:"DateOfBirth,omitempty"`
	Gender         string `json:"Gender,omitempty"`
	EmailAddress   string `json:"EmailAddress,omitempty"`
	IdentityNumber string `json:"IdentityNumber,omitempty"`
	Language       string `json:"Language,omitempty"`
	MaritalStatus  string `json:"MaritalStatus,omitempty"`
	Nationality    string `json:"Nationality,omitempty"`
	WorkNumber     string `json:"WorkNumber,omitempty"`
	HomeNumber     string `json:"HomeNumber,omitempty"`
	CellNumber     string `json:"CellNumber,omitempty"`
	MiddleName     string `json:"MiddleName,omitempty"`
	MaidenName     string `json:"MaidenName,omitempty"`
	Custom1        string `json:"Custom1,omitempty"`
	Custom2        string `json:"Custom2,omitempty"`
	Custom3        string `json:"Custom3,omitempty"`
	Custom4        string `json:"Custom4,omitempty"`
	Custom5        string `json:"Custom5,omitempty"`
	Custom6        string `json:"Custom6,omitempty"`
	Custom7        string `json:"Custom7,omitempty"`
	Custom8        string `json:"Custom8,omitempty"`
	Custom9        string `json:"Custom9,omitempty"`
	Custom10       string `json:"Custom10,omitempty"`
}

// EmployeeAddress represents an employee's address record.
type EmployeeAddress struct {
	AddressId      int    `json:"AddressId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	AddressType    string `json:"AddressType,omitempty"`
	AddressLine1   string `json:"AddressLine1,omitempty"`
	AddressLine2   string `json:"AddressLine2,omitempty"`
	AddressLine3   string `json:"AddressLine3,omitempty"`
	AddressLine4   string `json:"AddressLine4,omitempty"`
	PostalCode     string `json:"PostalCode,omitempty"`
	Country        string `json:"Country,omitempty"`
}

// EmployeeEmploymentStatus represents an employee's employment status record.
type EmployeeEmploymentStatus struct {
	EmploymentStatusId int    `json:"EmploymentStatusId"`
	EmployeeNumber     string `json:"EmployeeNumber"`
	EmploymentDate     string `json:"EmploymentDate,omitempty"`
	TerminationDate    string `json:"TerminationDate,omitempty"`
	TerminationReason  string `json:"TerminationReason,omitempty"`
	EmploymentAction   string `json:"EmploymentAction,omitempty"`
	EncashLeave        *bool  `json:"EncashLeave,omitempty"`
	EmploymentStatus   string `json:"EmploymentStatus,omitempty"`
	PayrollNumber      string `json:"PayrollNumber,omitempty"`
	ReEmploymentDate   string `json:"ReEmploymentDate,omitempty"`
}

// EmployeePosition represents an employee's position record.
type EmployeePosition struct {
	EmployeePositionId           int    `json:"EmployeePositionId"`
	EmployeeNumber               string `json:"EmployeeNumber"`
	OrganizationPosition         string `json:"OrganizationPosition,omitempty"`
	OrganizationPositionWithCode string `json:"OrganizationPositionWithCode,omitempty"`
	Department                   string `json:"Department,omitempty"`
	JobTitle                     string `json:"JobTitle,omitempty"`
	EffectiveDate                string `json:"EffectiveDate,omitempty"`
	ManagerEmployeeNumber        string `json:"ManagerEmployeeNumber,omitempty"`
	CostCentre                   string `json:"CostCentre,omitempty"`
	Division                     string `json:"Division,omitempty"`
}

// EmployeeService handles employee-related API operations.
type EmployeeService struct {
	client *Client
}

// List retrieves all employees for a company.
func (s *EmployeeService) List(ctx context.Context, companyID int, params *QueryParams) ([]Employee, *Response, error) {
	s.client.logger.Info("listing employees", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list employees query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "Employee")
	var result ListResult[Employee]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListByEffectiveDate retrieves employees effective as of the given date.
func (s *EmployeeService) ListByEffectiveDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]Employee, *Response, error) {
	s.client.logger.Info("listing employees by effective date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list employees by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("Employee/effective/%s", date))
	var result ListResult[Employee]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAll retrieves all employees including inactive ones.
func (s *EmployeeService) ListAll(ctx context.Context, companyID int, params *QueryParams) ([]Employee, *Response, error) {
	s.client.logger.Info("listing all employees", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all employees query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "Employee/all")
	var result ListResult[Employee]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// Get retrieves a single employee by ID.
func (s *EmployeeService) Get(ctx context.Context, companyID, employeeID int) (*Employee, *Response, error) {
	s.client.logger.Info("getting employee", zap.Int("company_id", companyID), zap.Int("employee_id", employeeID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "Employee"), employeeID)
	var employee Employee
	resp, err := s.client.getAndDecode(ctx, url, nil, &employee)
	if err != nil {
		return nil, resp, err
	}
	return &employee, resp, nil
}

// Create creates a new employee.
func (s *EmployeeService) Create(ctx context.Context, companyID int, employee *Employee) (*Employee, *Response, error) {
	s.client.logger.Info("creating employee", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "Employee")
	var created Employee
	resp, err := s.client.postAndDecode(ctx, url, employee, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// Update updates an existing employee.
func (s *EmployeeService) Update(ctx context.Context, companyID, employeeID int, employee *Employee) (*Response, error) {
	s.client.logger.Info("updating employee", zap.Int("company_id", companyID), zap.Int("employee_id", employeeID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "Employee"), employeeID)
	return s.client.patch(ctx, url, employee)
}

// ListAddresses retrieves addresses for an employee.
func (s *EmployeeService) ListAddresses(ctx context.Context, companyID int, employeeNumber string, params *QueryParams) ([]EmployeeAddress, *Response, error) {
	s.client.logger.Info("listing employee addresses", zap.Int("company_id", companyID), zap.String("employee_number", employeeNumber))
	if params != nil {
		s.client.logger.Debug("list addresses query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeAddress/%s", employeeNumber))
	var result ListResult[EmployeeAddress]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// UpdateAddress updates an employee address.
func (s *EmployeeService) UpdateAddress(ctx context.Context, companyID, addressID int, address *EmployeeAddress) (*Response, error) {
	s.client.logger.Info("updating employee address", zap.Int("company_id", companyID), zap.Int("address_id", addressID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeAddress"), addressID)
	return s.client.patch(ctx, url, address)
}

// ListEmploymentStatuses retrieves all employment statuses.
func (s *EmployeeService) ListEmploymentStatuses(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeEmploymentStatus, *Response, error) {
	s.client.logger.Info("listing employment statuses", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list employment statuses query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeEmploymentStatus")
	var result ListResult[EmployeeEmploymentStatus]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetEmploymentStatus retrieves a single employment status by ID.
func (s *EmployeeService) GetEmploymentStatus(ctx context.Context, companyID, statusID int) (*EmployeeEmploymentStatus, *Response, error) {
	s.client.logger.Info("getting employment status", zap.Int("company_id", companyID), zap.Int("status_id", statusID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeEmploymentStatus"), statusID)
	var status EmployeeEmploymentStatus
	resp, err := s.client.getAndDecode(ctx, url, nil, &status)
	if err != nil {
		return nil, resp, err
	}
	return &status, resp, nil
}

// ListEmploymentStatusesByDate retrieves employment statuses effective as of the given date.
func (s *EmployeeService) ListEmploymentStatusesByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeeEmploymentStatus, *Response, error) {
	s.client.logger.Info("listing employment statuses by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list employment statuses by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeEmploymentStatus/effective/%s", date))
	var result ListResult[EmployeeEmploymentStatus]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllEmploymentStatuses retrieves all employment statuses including inactive ones.
func (s *EmployeeService) ListAllEmploymentStatuses(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeEmploymentStatus, *Response, error) {
	s.client.logger.Info("listing all employment statuses", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all employment statuses query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeEmploymentStatus/all")
	var result ListResult[EmployeeEmploymentStatus]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateEmploymentStatus creates a new employment status record.
func (s *EmployeeService) CreateEmploymentStatus(ctx context.Context, companyID int, status *EmployeeEmploymentStatus) (*EmployeeEmploymentStatus, *Response, error) {
	s.client.logger.Info("creating employment status", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeEmploymentStatus")
	var created EmployeeEmploymentStatus
	resp, err := s.client.postAndDecode(ctx, url, status, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateEmploymentStatus updates an existing employment status.
func (s *EmployeeService) UpdateEmploymentStatus(ctx context.Context, companyID, statusID int, status *EmployeeEmploymentStatus) (*Response, error) {
	s.client.logger.Info("updating employment status", zap.Int("company_id", companyID), zap.Int("status_id", statusID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeEmploymentStatus"), statusID)
	return s.client.patch(ctx, url, status)
}

// DeleteEmploymentStatus deletes an employment status record.
func (s *EmployeeService) DeleteEmploymentStatus(ctx context.Context, companyID, statusID int) (*Response, error) {
	s.client.logger.Info("deleting employment status", zap.Int("company_id", companyID), zap.Int("status_id", statusID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeEmploymentStatus"), statusID)
	return s.client.del(ctx, url)
}

// ListPositions retrieves all employee positions.
func (s *EmployeeService) ListPositions(ctx context.Context, companyID int, params *QueryParams) ([]EmployeePosition, *Response, error) {
	s.client.logger.Info("listing employee positions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list positions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeePosition")
	var result ListResult[EmployeePosition]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetPosition retrieves a single employee position by ID.
func (s *EmployeeService) GetPosition(ctx context.Context, companyID, positionID int) (*EmployeePosition, *Response, error) {
	s.client.logger.Info("getting employee position", zap.Int("company_id", companyID), zap.Int("position_id", positionID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePosition"), positionID)
	var position EmployeePosition
	resp, err := s.client.getAndDecode(ctx, url, nil, &position)
	if err != nil {
		return nil, resp, err
	}
	return &position, resp, nil
}

// ListPositionsByDate retrieves employee positions effective as of the given date.
func (s *EmployeeService) ListPositionsByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeePosition, *Response, error) {
	s.client.logger.Info("listing positions by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list positions by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePosition/effective/%s", date))
	var result ListResult[EmployeePosition]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllPositions retrieves all employee positions including inactive ones.
func (s *EmployeeService) ListAllPositions(ctx context.Context, companyID int, params *QueryParams) ([]EmployeePosition, *Response, error) {
	s.client.logger.Info("listing all employee positions", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all positions query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeePosition/all")
	var result ListResult[EmployeePosition]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreatePosition creates a new employee position.
func (s *EmployeeService) CreatePosition(ctx context.Context, companyID int, position *EmployeePosition) (*EmployeePosition, *Response, error) {
	s.client.logger.Info("creating employee position", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeePosition")
	var created EmployeePosition
	resp, err := s.client.postAndDecode(ctx, url, position, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdatePosition updates an existing employee position.
func (s *EmployeeService) UpdatePosition(ctx context.Context, companyID, positionID int, position *EmployeePosition) (*Response, error) {
	s.client.logger.Info("updating employee position", zap.Int("company_id", companyID), zap.Int("position_id", positionID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePosition"), positionID)
	return s.client.patch(ctx, url, position)
}

// DeletePosition deletes an employee position record.
func (s *EmployeeService) DeletePosition(ctx context.Context, companyID, positionID int) (*Response, error) {
	s.client.logger.Info("deleting employee position", zap.Int("company_id", companyID), zap.Int("position_id", positionID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePosition"), positionID)
	return s.client.del(ctx, url)
}

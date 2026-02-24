package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeePayRate represents an employee's pay rate record.
type EmployeePayRate struct {
	PayRateId      int      `json:"PayRateId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	BasicSalary    *float64 `json:"BasicSalary,omitempty"`
	AnnualSalary   *float64 `json:"AnnualSalary,omitempty"`
	HourlyRate     *float64 `json:"HourlyRate,omitempty"`
	DailyRate      *float64 `json:"DailyRate,omitempty"`
	EffectiveDate  string   `json:"EffectiveDate,omitempty"`
	PayFrequency   string   `json:"PayFrequency,omitempty"`
	PayType        string   `json:"PayType,omitempty"`
}

// EmployeeBonusProvision represents a bonus provision component.
type EmployeeBonusProvision struct {
	BonusProvisionId int      `json:"BonusProvisionId"`
	EmployeeNumber   string   `json:"EmployeeNumber"`
	ComponentCode    string   `json:"ComponentCode,omitempty"`
	Description      string   `json:"Description,omitempty"`
	Amount           *float64 `json:"Amount,omitempty"`
	Percentage       *float64 `json:"Percentage,omitempty"`
	EffectiveDate    string   `json:"EffectiveDate,omitempty"`
}

// EmployeeComponent represents a recurring payroll component.
type EmployeeComponent struct {
	ComponentId    int      `json:"ComponentId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	ComponentCode  string   `json:"ComponentCode,omitempty"`
	Description    string   `json:"Description,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	Percentage     *float64 `json:"Percentage,omitempty"`
	Value          *float64 `json:"Value,omitempty"`
	EffectiveDate  string   `json:"EffectiveDate,omitempty"`
}

// EmployeeGarnishee represents a garnishee order component.
type EmployeeGarnishee struct {
	GarnisheeId    int      `json:"GarnisheeId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	ComponentCode  string   `json:"ComponentCode,omitempty"`
	Description    string   `json:"Description,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	Balance        *float64 `json:"Balance,omitempty"`
	EffectiveDate  string   `json:"EffectiveDate,omitempty"`
}

// EmployeeMedical represents a medical aid component.
type EmployeeMedical struct {
	MedicalId        int      `json:"MedicalId"`
	EmployeeNumber   string   `json:"EmployeeNumber"`
	MedicalAidScheme string   `json:"MedicalAidScheme,omitempty"`
	MedicalAidPlan   string   `json:"MedicalAidPlan,omitempty"`
	MemberNumber     string   `json:"MemberNumber,omitempty"`
	Amount           *float64 `json:"Amount,omitempty"`
	EffectiveDate    string   `json:"EffectiveDate,omitempty"`
}

// EmployeePensionFund represents a pension fund component.
type EmployeePensionFund struct {
	PensionFundId      int      `json:"PensionFundId"`
	EmployeeNumber     string   `json:"EmployeeNumber"`
	ComponentCode      string   `json:"ComponentCode,omitempty"`
	Description        string   `json:"Description,omitempty"`
	EmployeeAmount     *float64 `json:"EmployeeAmount,omitempty"`
	CompanyAmount      *float64 `json:"CompanyAmount,omitempty"`
	EmployeePercentage *float64 `json:"EmployeePercentage,omitempty"`
	CompanyPercentage  *float64 `json:"CompanyPercentage,omitempty"`
	EffectiveDate      string   `json:"EffectiveDate,omitempty"`
}

// EmployeeRetirementAnnuity represents a retirement annuity component.
type EmployeeRetirementAnnuity struct {
	RetirementAnnuityId int      `json:"RetirementAnnuityId"`
	EmployeeNumber      string   `json:"EmployeeNumber"`
	ComponentCode       string   `json:"ComponentCode,omitempty"`
	Description         string   `json:"Description,omitempty"`
	Amount              *float64 `json:"Amount,omitempty"`
	Percentage          *float64 `json:"Percentage,omitempty"`
	EffectiveDate       string   `json:"EffectiveDate,omitempty"`
}

// EmployeeLoan represents a loan component.
type EmployeeLoan struct {
	LoanId         int      `json:"LoanId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	ComponentCode  string   `json:"ComponentCode,omitempty"`
	Description    string   `json:"Description,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	Balance        *float64 `json:"Balance,omitempty"`
	Instalment     *float64 `json:"Instalment,omitempty"`
	EffectiveDate  string   `json:"EffectiveDate,omitempty"`
}

// ExtractionRequest represents a response from starting an extraction process.
type ExtractionRequest struct {
	ExtractionId string `json:"ExtractionId"`
	Status       string `json:"Status"`
}

// ExtractionStatus represents the status of an extraction process.
type ExtractionStatus struct {
	ExtractionId string `json:"ExtractionId"`
	Status       string `json:"Status"`
	DownloadUrl  string `json:"DownloadUrl,omitempty"`
}

// EmployeePayrollService handles employee payroll-related API operations.
type EmployeePayrollService struct {
	client *Client
}

// --- Pay Rate methods ---

// ListPayRates retrieves all employee pay rates.
func (s *EmployeePayrollService) ListPayRates(ctx context.Context, companyID int, params *QueryParams) ([]EmployeePayRate, *Response, error) {
	s.client.logger.Info("listing employee pay rates", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list pay rates query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeePayRate")
	var result ListResult[EmployeePayRate]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetPayRate retrieves a single pay rate by ID.
func (s *EmployeePayrollService) GetPayRate(ctx context.Context, companyID, payRateID int) (*EmployeePayRate, *Response, error) {
	s.client.logger.Info("getting employee pay rate", zap.Int("company_id", companyID), zap.Int("pay_rate_id", payRateID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePayRate"), payRateID)
	var payRate EmployeePayRate
	resp, err := s.client.getAndDecode(ctx, url, nil, &payRate)
	if err != nil {
		return nil, resp, err
	}
	return &payRate, resp, nil
}

// ListPayRatesByDate retrieves pay rates effective as of the given date.
func (s *EmployeePayrollService) ListPayRatesByDate(ctx context.Context, companyID int, date string, params *QueryParams) ([]EmployeePayRate, *Response, error) {
	s.client.logger.Info("listing pay rates by date", zap.Int("company_id", companyID), zap.String("date", date))
	if params != nil {
		s.client.logger.Debug("list pay rates by date query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePayRate/effective/%s", date))
	var result ListResult[EmployeePayRate]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListAllPayRates retrieves all pay rates including inactive ones.
func (s *EmployeePayrollService) ListAllPayRates(ctx context.Context, companyID int, params *QueryParams) ([]EmployeePayRate, *Response, error) {
	s.client.logger.Info("listing all pay rates", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list all pay rates query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeePayRate/all")
	var result ListResult[EmployeePayRate]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreatePayRate creates a new pay rate record.
func (s *EmployeePayrollService) CreatePayRate(ctx context.Context, companyID int, payRate *EmployeePayRate) (*EmployeePayRate, *Response, error) {
	s.client.logger.Info("creating employee pay rate", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeePayRate")
	var created EmployeePayRate
	resp, err := s.client.postAndDecode(ctx, url, payRate, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdatePayRate updates an existing pay rate.
func (s *EmployeePayrollService) UpdatePayRate(ctx context.Context, companyID, payRateID int, payRate *EmployeePayRate) (*Response, error) {
	s.client.logger.Info("updating employee pay rate", zap.Int("company_id", companyID), zap.Int("pay_rate_id", payRateID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePayRate"), payRateID)
	return s.client.patch(ctx, url, payRate)
}

// DeletePayRate deletes a pay rate record.
func (s *EmployeePayrollService) DeletePayRate(ctx context.Context, companyID, payRateID int) (*Response, error) {
	s.client.logger.Info("deleting employee pay rate", zap.Int("company_id", companyID), zap.Int("pay_rate_id", payRateID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePayRate"), payRateID)
	return s.client.del(ctx, url)
}

// --- Generic recurring component methods ---

// ListRecurring retrieves recurring component records for a given entity.
func (s *EmployeePayrollService) ListRecurring(ctx context.Context, companyID int, entityName, frequency, period string, params *QueryParams) ([]map[string]any, *Response, error) {
	s.client.logger.Info("listing recurring components",
		zap.Int("company_id", companyID),
		zap.String("entity", entityName),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)
	if params != nil {
		s.client.logger.Debug("list recurring query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, entityName)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var result ListResult[map[string]any]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// CreateRecurring creates a new recurring component record.
func (s *EmployeePayrollService) CreateRecurring(ctx context.Context, companyID int, entityName, frequency, period string, body any) (map[string]any, *Response, error) {
	s.client.logger.Info("creating recurring component",
		zap.Int("company_id", companyID),
		zap.String("entity", entityName),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := s.client.odataURL(companyID, entityName)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var created map[string]any
	resp, err := s.client.postAndDecode(ctx, url, body, &created)
	if err != nil {
		return nil, resp, err
	}
	return created, resp, nil
}

// UpdateRecurring updates an existing recurring component record.
func (s *EmployeePayrollService) UpdateRecurring(ctx context.Context, companyID int, entityName string, id int, frequency, period string, body any) (*Response, error) {
	s.client.logger.Info("updating recurring component",
		zap.Int("company_id", companyID),
		zap.String("entity", entityName),
		zap.Int("id", id),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, entityName), id)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	return s.client.patch(ctx, url, body)
}

// DeleteRecurring deletes a recurring component record.
func (s *EmployeePayrollService) DeleteRecurring(ctx context.Context, companyID int, entityName string, id int, frequency, period string) (*Response, error) {
	s.client.logger.Info("deleting recurring component",
		zap.Int("company_id", companyID),
		zap.String("entity", entityName),
		zap.Int("id", id),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, entityName), id)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	return s.client.del(ctx, url)
}

// --- EmployeeComponent extraction methods ---

// ExtractComponents starts an extraction process for employee components.
func (s *EmployeePayrollService) ExtractComponents(ctx context.Context, companyID int, frequency, period string) (*ExtractionRequest, *Response, error) {
	s.client.logger.Info("extracting employee components",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := s.client.odataURL(companyID, "EmployeeComponent/extract/start")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var extraction ExtractionRequest
	resp, err := s.client.postAndDecode(ctx, url, nil, &extraction)
	if err != nil {
		return nil, resp, err
	}
	return &extraction, resp, nil
}

// ExtractComponentsStatus retrieves the status of a component extraction.
func (s *EmployeePayrollService) ExtractComponentsStatus(ctx context.Context, companyID int, extractionID, frequency, period string) (*ExtractionStatus, *Response, error) {
	s.client.logger.Info("getting component extraction status",
		zap.Int("company_id", companyID),
		zap.String("extraction_id", extractionID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeComponent/extract/%s/status", extractionID))
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var status ExtractionStatus
	resp, err := s.client.getAndDecode(ctx, url, nil, &status)
	if err != nil {
		return nil, resp, err
	}
	return &status, resp, nil
}

package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeePayslip represents an employee payslip record.
type EmployeePayslip struct {
	PayslipId       int      `json:"PayslipId"`
	EmployeeNumber  string   `json:"EmployeeNumber"`
	GrossPay        *float64 `json:"GrossPay,omitempty"`
	NetPay          *float64 `json:"NetPay,omitempty"`
	TotalDeductions *float64 `json:"TotalDeductions,omitempty"`
	TotalCompany    *float64 `json:"TotalCompany,omitempty"`
	Year            int      `json:"Year,omitempty"`
	Month           int      `json:"Month,omitempty"`
	Period          int      `json:"Period,omitempty"`
	Frequency       string   `json:"Frequency,omitempty"`
	PayDate         string   `json:"PayDate,omitempty"`
	FirstName       string   `json:"FirstName,omitempty"`
	Surname         string   `json:"Surname,omitempty"`
}

// EmployeePayslipLine represents a line item on a payslip.
type EmployeePayslipLine struct {
	PayslipLineId  int      `json:"PayslipLineId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	ComponentCode  string   `json:"ComponentCode,omitempty"`
	Description    string   `json:"Description,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	LineType       string   `json:"LineType,omitempty"`
	Year           int      `json:"Year,omitempty"`
	Month          int      `json:"Month,omitempty"`
}

// EmployeeCostedPayslipLine represents a costed payslip line item.
type EmployeeCostedPayslipLine struct {
	CostedPayslipLineId int      `json:"CostedPayslipLineId"`
	EmployeeNumber      string   `json:"EmployeeNumber"`
	ComponentCode       string   `json:"ComponentCode,omitempty"`
	Description         string   `json:"Description,omitempty"`
	Amount              *float64 `json:"Amount,omitempty"`
	CostCentre          string   `json:"CostCentre,omitempty"`
	Year                int      `json:"Year,omitempty"`
	Month               int      `json:"Month,omitempty"`
}

// EmployeePayslipPdf represents a payslip PDF record.
type EmployeePayslipPdf struct {
	PayslipPdfId   int    `json:"PayslipPdfId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	PdfContent     string `json:"PdfContent,omitempty"`
	Year           int    `json:"Year,omitempty"`
	Month          int    `json:"Month,omitempty"`
}

// EditPayslip represents an editable payslip record.
type EditPayslip struct {
	EditPayslipId  int      `json:"EditPayslipId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	ComponentCode  string   `json:"ComponentCode,omitempty"`
	Description    string   `json:"Description,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	Value          *float64 `json:"Value,omitempty"`
	Hours          *float64 `json:"Hours,omitempty"`
	Days           *float64 `json:"Days,omitempty"`
	LineType       string   `json:"LineType,omitempty"`
}

// EmployeePayslipTakeOn represents a payslip take-on record.
type EmployeePayslipTakeOn struct {
	PayslipTakeOnId int      `json:"PayslipTakeOnId"`
	EmployeeNumber  string   `json:"EmployeeNumber"`
	ComponentCode   string   `json:"ComponentCode,omitempty"`
	Description     string   `json:"Description,omitempty"`
	Amount          *float64 `json:"Amount,omitempty"`
	EffectiveDate   string   `json:"EffectiveDate,omitempty"`
}

// EmployeeLumpSum represents a lump sum record.
type EmployeeLumpSum struct {
	LumpSumId      int      `json:"LumpSumId"`
	EmployeeNumber string   `json:"EmployeeNumber"`
	ComponentCode  string   `json:"ComponentCode,omitempty"`
	Description    string   `json:"Description,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	EffectiveDate  string   `json:"EffectiveDate,omitempty"`
}

// EmployeeTaxCertificate represents a tax certificate record.
type EmployeeTaxCertificate struct {
	TaxCertificateId int    `json:"TaxCertificateId"`
	EmployeeNumber   string `json:"EmployeeNumber"`
	TaxYear          int    `json:"TaxYear,omitempty"`
	CertificateType  string `json:"CertificateType,omitempty"`
	IssueDate        string `json:"IssueDate,omitempty"`
	PdfContent       string `json:"PdfContent,omitempty"`
}

// EmployeePayslipService handles payslip-related API operations.
type EmployeePayslipService struct {
	client *Client
}

// --- Payslip read methods ---

// ListPayslips retrieves payslips for a given year and month.
func (s *EmployeePayslipService) ListPayslips(ctx context.Context, companyID, year, month int, params *QueryParams) ([]EmployeePayslip, *Response, error) {
	s.client.logger.Info("listing payslips",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list payslips query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePayslip/%d/%d", year, month))
	var result ListResult[EmployeePayslip]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListPayslipLines retrieves payslip line items for a given year and month.
func (s *EmployeePayslipService) ListPayslipLines(ctx context.Context, companyID, year, month int, params *QueryParams) ([]EmployeePayslipLine, *Response, error) {
	s.client.logger.Info("listing payslip lines",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list payslip lines query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePayslipLine/%d/%d", year, month))
	var result ListResult[EmployeePayslipLine]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListCostedPayslipLines retrieves costed payslip line items for a given year and month.
func (s *EmployeePayslipService) ListCostedPayslipLines(ctx context.Context, companyID, year, month int, params *QueryParams) ([]EmployeeCostedPayslipLine, *Response, error) {
	s.client.logger.Info("listing costed payslip lines",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list costed payslip lines query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeCostedPayslipLine/%d/%d", year, month))
	var result ListResult[EmployeeCostedPayslipLine]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListConsolidatedPayslips retrieves consolidated payslips for a given year and month.
func (s *EmployeePayslipService) ListConsolidatedPayslips(ctx context.Context, companyID, year, month int, params *QueryParams) ([]EmployeePayslip, *Response, error) {
	s.client.logger.Info("listing consolidated payslips",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list consolidated payslips query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePayslip/%d/%d/consolidated", year, month))
	var result ListResult[EmployeePayslip]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListPayslipPDFs retrieves payslip PDFs for a given year and month.
func (s *EmployeePayslipService) ListPayslipPDFs(ctx context.Context, companyID, year, month int, params *QueryParams) ([]EmployeePayslipPdf, *Response, error) {
	s.client.logger.Info("listing payslip PDFs",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list payslip PDFs query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePayslip/%d/%d/pdf", year, month))
	var result ListResult[EmployeePayslipPdf]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// UpdatePayslipComment updates a payslip comment using the v1.0 endpoint.
func (s *EmployeePayslipService) UpdatePayslipComment(ctx context.Context, companyID int, frequency, period string, body any) (*Response, error) {
	s.client.logger.Info("updating payslip comment",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := s.client.odataURLv1(companyID, "EmployeePayslip/comment")
	url = addQueryParams(url, map[string]string{
		"period":    period,
		"frequency": frequency,
	})
	return s.client.patch(ctx, url, body)
}

// --- Payslip line extraction methods ---

// ExtractPayslipLines starts an extraction of payslip lines for a given year and month.
func (s *EmployeePayslipService) ExtractPayslipLines(ctx context.Context, companyID, year, month int) (*ExtractionRequest, *Response, error) {
	s.client.logger.Info("extracting payslip lines",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePayslipLine/extract/start/%d/%d", year, month))
	var extraction ExtractionRequest
	resp, err := s.client.postAndDecode(ctx, url, nil, &extraction)
	if err != nil {
		return nil, resp, err
	}
	return &extraction, resp, nil
}

// ExtractPayslipLinesStatus retrieves the status of a payslip line extraction.
func (s *EmployeePayslipService) ExtractPayslipLinesStatus(ctx context.Context, companyID int, extractionID string) (*ExtractionStatus, *Response, error) {
	s.client.logger.Info("getting payslip line extraction status",
		zap.Int("company_id", companyID),
		zap.String("extraction_id", extractionID),
	)

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeePayslipLine/extract/%s/status", extractionID))
	var status ExtractionStatus
	resp, err := s.client.getAndDecode(ctx, url, nil, &status)
	if err != nil {
		return nil, resp, err
	}
	return &status, resp, nil
}

// ExtractCostedPayslipLines starts an extraction of costed payslip lines for a given year and month.
func (s *EmployeePayslipService) ExtractCostedPayslipLines(ctx context.Context, companyID, year, month int) (*ExtractionRequest, *Response, error) {
	s.client.logger.Info("extracting costed payslip lines",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeCostedPayslipLine/extract/start/%d/%d", year, month))
	var extraction ExtractionRequest
	resp, err := s.client.postAndDecode(ctx, url, nil, &extraction)
	if err != nil {
		return nil, resp, err
	}
	return &extraction, resp, nil
}

// ExtractCostedPayslipLinesStatus retrieves the status of a costed payslip line extraction.
func (s *EmployeePayslipService) ExtractCostedPayslipLinesStatus(ctx context.Context, companyID int, extractionID string) (*ExtractionStatus, *Response, error) {
	s.client.logger.Info("getting costed payslip line extraction status",
		zap.Int("company_id", companyID),
		zap.String("extraction_id", extractionID),
	)

	url := s.client.odataURL(companyID, fmt.Sprintf("EmployeeCostedPayslipLine/extract/%s/status", extractionID))
	var status ExtractionStatus
	resp, err := s.client.getAndDecode(ctx, url, nil, &status)
	if err != nil {
		return nil, resp, err
	}
	return &status, resp, nil
}

// --- EditPayslip methods ---

// ListEditPayslips retrieves edit payslip records.
func (s *EmployeePayslipService) ListEditPayslips(ctx context.Context, companyID int, frequency, period string, params *QueryParams) ([]EditPayslip, *Response, error) {
	s.client.logger.Info("listing edit payslips",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)
	if params != nil {
		s.client.logger.Debug("list edit payslips query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EditPayslip")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var result ListResult[EditPayslip]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetEditPayslip retrieves a single edit payslip by ID.
func (s *EmployeePayslipService) GetEditPayslip(ctx context.Context, companyID, editPayslipID int, frequency, period string) (*EditPayslip, *Response, error) {
	s.client.logger.Info("getting edit payslip",
		zap.Int("company_id", companyID),
		zap.Int("edit_payslip_id", editPayslipID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EditPayslip"), editPayslipID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var editPayslip EditPayslip
	resp, err := s.client.getAndDecode(ctx, url, nil, &editPayslip)
	if err != nil {
		return nil, resp, err
	}
	return &editPayslip, resp, nil
}

// CreateEditPayslip creates a new edit payslip record.
func (s *EmployeePayslipService) CreateEditPayslip(ctx context.Context, companyID int, frequency, period, action string, body *EditPayslip) (*EditPayslip, *Response, error) {
	s.client.logger.Info("creating edit payslip",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
		zap.String("action", action),
	)

	url := s.client.odataURL(companyID, "EditPayslip")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
		"Action":    action,
	})
	var created EditPayslip
	resp, err := s.client.postAndDecode(ctx, url, body, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateEditPayslip updates an existing edit payslip record.
func (s *EmployeePayslipService) UpdateEditPayslip(ctx context.Context, companyID, editPayslipID int, frequency, period, action string, body *EditPayslip) (*Response, error) {
	s.client.logger.Info("updating edit payslip",
		zap.Int("company_id", companyID),
		zap.Int("edit_payslip_id", editPayslipID),
		zap.String("frequency", frequency),
		zap.String("period", period),
		zap.String("action", action),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EditPayslip"), editPayslipID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
		"Action":    action,
	})
	return s.client.patch(ctx, url, body)
}

// BatchEditPayslips creates or updates multiple edit payslip records in a batch.
func (s *EmployeePayslipService) BatchEditPayslips(ctx context.Context, companyID int, frequency, period string, items []EditPayslip) ([]EditPayslip, *Response, error) {
	s.client.logger.Info("batch editing payslips",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
		zap.Int("item_count", len(items)),
	)

	url := s.client.odataURL(companyID, "EditPayslip/Batch")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var result ListResult[EditPayslip]
	resp, err := s.client.postAndDecode(ctx, url, items, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// DeleteEditPayslip deletes an edit payslip record.
func (s *EmployeePayslipService) DeleteEditPayslip(ctx context.Context, companyID, editPayslipID int, frequency, period string) (*Response, error) {
	s.client.logger.Info("deleting edit payslip",
		zap.Int("company_id", companyID),
		zap.Int("edit_payslip_id", editPayslipID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EditPayslip"), editPayslipID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	return s.client.del(ctx, url)
}

// --- EmployeePayslipTakeOn methods ---

// ListPayslipTakeOns retrieves payslip take-on records.
func (s *EmployeePayslipService) ListPayslipTakeOns(ctx context.Context, companyID int, frequency, period string, params *QueryParams) ([]EmployeePayslipTakeOn, *Response, error) {
	s.client.logger.Info("listing payslip take-ons",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)
	if params != nil {
		s.client.logger.Debug("list payslip take-ons query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeePayslipTakeOn")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var result ListResult[EmployeePayslipTakeOn]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetPayslipTakeOn retrieves a single payslip take-on by ID.
func (s *EmployeePayslipService) GetPayslipTakeOn(ctx context.Context, companyID, takeOnID int, frequency, period string) (*EmployeePayslipTakeOn, *Response, error) {
	s.client.logger.Info("getting payslip take-on",
		zap.Int("company_id", companyID),
		zap.Int("take_on_id", takeOnID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePayslipTakeOn"), takeOnID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var takeOn EmployeePayslipTakeOn
	resp, err := s.client.getAndDecode(ctx, url, nil, &takeOn)
	if err != nil {
		return nil, resp, err
	}
	return &takeOn, resp, nil
}

// CreatePayslipTakeOn creates a new payslip take-on record.
func (s *EmployeePayslipService) CreatePayslipTakeOn(ctx context.Context, companyID int, frequency, period string, takeOn *EmployeePayslipTakeOn) (*EmployeePayslipTakeOn, *Response, error) {
	s.client.logger.Info("creating payslip take-on",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := s.client.odataURL(companyID, "EmployeePayslipTakeOn")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var created EmployeePayslipTakeOn
	resp, err := s.client.postAndDecode(ctx, url, takeOn, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdatePayslipTakeOn updates an existing payslip take-on record.
func (s *EmployeePayslipService) UpdatePayslipTakeOn(ctx context.Context, companyID, takeOnID int, frequency, period string, takeOn *EmployeePayslipTakeOn) (*Response, error) {
	s.client.logger.Info("updating payslip take-on",
		zap.Int("company_id", companyID),
		zap.Int("take_on_id", takeOnID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePayslipTakeOn"), takeOnID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	return s.client.patch(ctx, url, takeOn)
}

// DeletePayslipTakeOn deletes a payslip take-on record.
func (s *EmployeePayslipService) DeletePayslipTakeOn(ctx context.Context, companyID, takeOnID int, frequency, period string) (*Response, error) {
	s.client.logger.Info("deleting payslip take-on",
		zap.Int("company_id", companyID),
		zap.Int("take_on_id", takeOnID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeePayslipTakeOn"), takeOnID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	return s.client.del(ctx, url)
}

// --- EmployeeLumpSum methods ---

// ListLumpSums retrieves lump sum records.
func (s *EmployeePayslipService) ListLumpSums(ctx context.Context, companyID int, frequency, period string, params *QueryParams) ([]EmployeeLumpSum, *Response, error) {
	s.client.logger.Info("listing lump sums",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)
	if params != nil {
		s.client.logger.Debug("list lump sums query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeLumpSum")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var result ListResult[EmployeeLumpSum]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetLumpSum retrieves a single lump sum by ID.
func (s *EmployeePayslipService) GetLumpSum(ctx context.Context, companyID, lumpSumID int, frequency, period string) (*EmployeeLumpSum, *Response, error) {
	s.client.logger.Info("getting lump sum",
		zap.Int("company_id", companyID),
		zap.Int("lump_sum_id", lumpSumID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLumpSum"), lumpSumID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var lumpSum EmployeeLumpSum
	resp, err := s.client.getAndDecode(ctx, url, nil, &lumpSum)
	if err != nil {
		return nil, resp, err
	}
	return &lumpSum, resp, nil
}

// CreateLumpSum creates a new lump sum record.
func (s *EmployeePayslipService) CreateLumpSum(ctx context.Context, companyID int, frequency, period string, lumpSum *EmployeeLumpSum) (*EmployeeLumpSum, *Response, error) {
	s.client.logger.Info("creating lump sum",
		zap.Int("company_id", companyID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := s.client.odataURL(companyID, "EmployeeLumpSum")
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	var created EmployeeLumpSum
	resp, err := s.client.postAndDecode(ctx, url, lumpSum, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateLumpSum updates an existing lump sum record.
func (s *EmployeePayslipService) UpdateLumpSum(ctx context.Context, companyID, lumpSumID int, frequency, period string, lumpSum *EmployeeLumpSum) (*Response, error) {
	s.client.logger.Info("updating lump sum",
		zap.Int("company_id", companyID),
		zap.Int("lump_sum_id", lumpSumID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLumpSum"), lumpSumID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	return s.client.patch(ctx, url, lumpSum)
}

// DeleteLumpSum deletes a lump sum record.
func (s *EmployeePayslipService) DeleteLumpSum(ctx context.Context, companyID, lumpSumID int, frequency, period string) (*Response, error) {
	s.client.logger.Info("deleting lump sum",
		zap.Int("company_id", companyID),
		zap.Int("lump_sum_id", lumpSumID),
		zap.String("frequency", frequency),
		zap.String("period", period),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeLumpSum"), lumpSumID)
	url = addQueryParams(url, map[string]string{
		"frequency": frequency,
		"period":    period,
	})
	return s.client.del(ctx, url)
}

// --- Tax Certificate methods ---

// ListTaxCertificates retrieves employee tax certificates.
func (s *EmployeePayslipService) ListTaxCertificates(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeTaxCertificate, *Response, error) {
	s.client.logger.Info("listing tax certificates", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list tax certificates query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeTaxCertificate")
	var result ListResult[EmployeeTaxCertificate]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

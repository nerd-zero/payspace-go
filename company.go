package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// CompanyInfo represents basic company information.
type CompanyInfo struct {
	CompanyId   int    `json:"CompanyId"`
	CompanyName string `json:"CompanyName"`
	CompanyCode string `json:"CompanyCode,omitempty"`
}

// CompanyDocument represents a company billing document.
type CompanyDocument struct {
	DocumentId     int      `json:"DocumentId"`
	EmployeeNumber string   `json:"EmployeeNumber,omitempty"`
	Year           int      `json:"Year,omitempty"`
	Month          int      `json:"Month,omitempty"`
	DocumentTotal  *float64 `json:"DocumentTotal,omitempty"`
	VatTotal       *float64 `json:"VatTotal,omitempty"`
}

// CompanyDocumentLine represents a line item in a company billing document.
type CompanyDocumentLine struct {
	DocumentLineId int      `json:"DocumentLineId"`
	DocumentId     int      `json:"DocumentId,omitempty"`
	Description    string   `json:"Description,omitempty"`
	Amount         *float64 `json:"Amount,omitempty"`
	VatAmount      *float64 `json:"VatAmount,omitempty"`
}

// CompanyCustomForm represents a custom form associated with a company.
type CompanyCustomForm struct {
	CustomFormId    int    `json:"CustomFormId"`
	EmployeeNumber  string `json:"EmployeeNumber,omitempty"`
	FormCategory    string `json:"FormCategory,omitempty"`
	FormDescription string `json:"FormDescription,omitempty"`
}

// CompanyService handles company-level API operations.
type CompanyService struct {
	client *Client
}

// --- Billing ---

// ListDocuments retrieves company billing documents for a given year and month.
func (s *CompanyService) ListDocuments(ctx context.Context, companyID, year, month int, params *QueryParams) ([]CompanyDocument, *Response, error) {
	s.client.logger.Info("listing company documents",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list company documents query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("CompanyDocument/%d/%d", year, month))
	var result ListResult[CompanyDocument]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetDocument retrieves a single company billing document by ID.
func (s *CompanyService) GetDocument(ctx context.Context, companyID, documentID int) (*CompanyDocument, *Response, error) {
	s.client.logger.Info("getting company document",
		zap.Int("company_id", companyID),
		zap.Int("document_id", documentID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "CompanyDocument"), documentID)
	var doc CompanyDocument
	resp, err := s.client.getAndDecode(ctx, url, nil, &doc)
	if err != nil {
		return nil, resp, err
	}
	return &doc, resp, nil
}

// ListDocumentLines retrieves company billing document lines for a given year and month.
func (s *CompanyService) ListDocumentLines(ctx context.Context, companyID, year, month int, params *QueryParams) ([]CompanyDocumentLine, *Response, error) {
	s.client.logger.Info("listing company document lines",
		zap.Int("company_id", companyID),
		zap.Int("year", year),
		zap.Int("month", month),
	)
	if params != nil {
		s.client.logger.Debug("list company document lines query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("CompanyDocumentLine/%d/%d", year, month))
	var result ListResult[CompanyDocumentLine]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// --- Custom Forms ---

// ListCustomForms retrieves all custom forms for a company.
func (s *CompanyService) ListCustomForms(ctx context.Context, companyID int, params *QueryParams) ([]CompanyCustomForm, *Response, error) {
	s.client.logger.Info("listing company custom forms", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list company custom forms query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "CompanyCustomForm")
	var result ListResult[CompanyCustomForm]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetCustomForm retrieves a single custom form by ID.
func (s *CompanyService) GetCustomForm(ctx context.Context, companyID, formID int) (*CompanyCustomForm, *Response, error) {
	s.client.logger.Info("getting company custom form",
		zap.Int("company_id", companyID),
		zap.Int("form_id", formID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "CompanyCustomForm"), formID)
	var form CompanyCustomForm
	resp, err := s.client.getAndDecode(ctx, url, nil, &form)
	if err != nil {
		return nil, resp, err
	}
	return &form, resp, nil
}

// CreateCustomForm creates a new custom form.
func (s *CompanyService) CreateCustomForm(ctx context.Context, companyID int, form *CompanyCustomForm) (*CompanyCustomForm, *Response, error) {
	s.client.logger.Info("creating company custom form", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "CompanyCustomForm")
	var created CompanyCustomForm
	resp, err := s.client.postAndDecode(ctx, url, form, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateCustomForm updates an existing custom form.
func (s *CompanyService) UpdateCustomForm(ctx context.Context, companyID, formID int, form *CompanyCustomForm) (*Response, error) {
	s.client.logger.Info("updating company custom form",
		zap.Int("company_id", companyID),
		zap.Int("form_id", formID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "CompanyCustomForm"), formID)
	return s.client.patch(ctx, url, form)
}

// DeleteCustomForm deletes a custom form.
func (s *CompanyService) DeleteCustomForm(ctx context.Context, companyID, formID int) (*Response, error) {
	s.client.logger.Info("deleting company custom form",
		zap.Int("company_id", companyID),
		zap.Int("form_id", formID),
	)

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "CompanyCustomForm"), formID)
	return s.client.del(ctx, url)
}

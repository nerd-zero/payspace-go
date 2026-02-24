package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// EmployeeOutOfOffice represents an out-of-office record.
type EmployeeOutOfOffice struct {
	OutOfOfficeId  int    `json:"OutOfOfficeId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	StartDate      string `json:"StartDate,omitempty"`
	EndDate        string `json:"EndDate,omitempty"`
	Reason         string `json:"Reason,omitempty"`
	Status         string `json:"Status,omitempty"`
}

// EmployeeNote represents an employee note record.
type EmployeeNote struct {
	NoteId         int    `json:"NoteId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	NoteDate       string `json:"NoteDate,omitempty"`
	NoteType       string `json:"NoteType,omitempty"`
	Subject        string `json:"Subject,omitempty"`
	NoteText       string `json:"NoteText,omitempty"`
	CreatedBy      string `json:"CreatedBy,omitempty"`
}

// EmployeeCustomForm represents a custom form record.
type EmployeeCustomForm struct {
	CustomFormId   int    `json:"CustomFormId"`
	EmployeeNumber string `json:"EmployeeNumber"`
	FormName       string `json:"FormName,omitempty"`
	FormData       string `json:"FormData,omitempty"`
	EffectiveDate  string `json:"EffectiveDate,omitempty"`
}

// EmployeeOtherService handles miscellaneous employee-related API operations.
type EmployeeOtherService struct {
	client *Client
}

// --- Out of Office methods ---

// ListOutOfOffice retrieves all out-of-office records.
func (s *EmployeeOtherService) ListOutOfOffice(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeOutOfOffice, *Response, error) {
	s.client.logger.Info("listing out of office records", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list out of office query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeOutOfOffice")
	var result ListResult[EmployeeOutOfOffice]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetOutOfOffice retrieves a single out-of-office record by ID.
func (s *EmployeeOtherService) GetOutOfOffice(ctx context.Context, companyID, outOfOfficeID int) (*EmployeeOutOfOffice, *Response, error) {
	s.client.logger.Info("getting out of office record", zap.Int("company_id", companyID), zap.Int("out_of_office_id", outOfOfficeID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeOutOfOffice"), outOfOfficeID)
	var record EmployeeOutOfOffice
	resp, err := s.client.getAndDecode(ctx, url, nil, &record)
	if err != nil {
		return nil, resp, err
	}
	return &record, resp, nil
}

// CreateOutOfOffice creates a new out-of-office record.
func (s *EmployeeOtherService) CreateOutOfOffice(ctx context.Context, companyID int, record *EmployeeOutOfOffice) (*EmployeeOutOfOffice, *Response, error) {
	s.client.logger.Info("creating out of office record", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeOutOfOffice")
	var created EmployeeOutOfOffice
	resp, err := s.client.postAndDecode(ctx, url, record, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateOutOfOffice updates an existing out-of-office record.
func (s *EmployeeOtherService) UpdateOutOfOffice(ctx context.Context, companyID, outOfOfficeID int, record *EmployeeOutOfOffice) (*Response, error) {
	s.client.logger.Info("updating out of office record", zap.Int("company_id", companyID), zap.Int("out_of_office_id", outOfOfficeID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeOutOfOffice"), outOfOfficeID)
	return s.client.patch(ctx, url, record)
}

// DeleteOutOfOffice deletes an out-of-office record.
func (s *EmployeeOtherService) DeleteOutOfOffice(ctx context.Context, companyID, outOfOfficeID int) (*Response, error) {
	s.client.logger.Info("deleting out of office record", zap.Int("company_id", companyID), zap.Int("out_of_office_id", outOfOfficeID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeOutOfOffice"), outOfOfficeID)
	return s.client.del(ctx, url)
}

// --- Note methods ---

// ListNotes retrieves all employee notes.
func (s *EmployeeOtherService) ListNotes(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeNote, *Response, error) {
	s.client.logger.Info("listing employee notes", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list notes query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeNote")
	var result ListResult[EmployeeNote]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetNote retrieves a single employee note by ID.
func (s *EmployeeOtherService) GetNote(ctx context.Context, companyID, noteID int) (*EmployeeNote, *Response, error) {
	s.client.logger.Info("getting employee note", zap.Int("company_id", companyID), zap.Int("note_id", noteID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeNote"), noteID)
	var note EmployeeNote
	resp, err := s.client.getAndDecode(ctx, url, nil, &note)
	if err != nil {
		return nil, resp, err
	}
	return &note, resp, nil
}

// CreateNote creates a new employee note.
func (s *EmployeeOtherService) CreateNote(ctx context.Context, companyID int, note *EmployeeNote) (*EmployeeNote, *Response, error) {
	s.client.logger.Info("creating employee note", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeNote")
	var created EmployeeNote
	resp, err := s.client.postAndDecode(ctx, url, note, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateNote updates an existing employee note.
func (s *EmployeeOtherService) UpdateNote(ctx context.Context, companyID, noteID int, note *EmployeeNote) (*Response, error) {
	s.client.logger.Info("updating employee note", zap.Int("company_id", companyID), zap.Int("note_id", noteID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeNote"), noteID)
	return s.client.patch(ctx, url, note)
}

// DeleteNote deletes an employee note.
func (s *EmployeeOtherService) DeleteNote(ctx context.Context, companyID, noteID int) (*Response, error) {
	s.client.logger.Info("deleting employee note", zap.Int("company_id", companyID), zap.Int("note_id", noteID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeNote"), noteID)
	return s.client.del(ctx, url)
}

// --- Custom Form methods ---

// ListCustomForms retrieves all employee custom forms.
func (s *EmployeeOtherService) ListCustomForms(ctx context.Context, companyID int, params *QueryParams) ([]EmployeeCustomForm, *Response, error) {
	s.client.logger.Info("listing employee custom forms", zap.Int("company_id", companyID))
	if params != nil {
		s.client.logger.Debug("list custom forms query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, "EmployeeCustomForm")
	var result ListResult[EmployeeCustomForm]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// GetCustomForm retrieves a single custom form by ID.
func (s *EmployeeOtherService) GetCustomForm(ctx context.Context, companyID, customFormID int) (*EmployeeCustomForm, *Response, error) {
	s.client.logger.Info("getting employee custom form", zap.Int("company_id", companyID), zap.Int("custom_form_id", customFormID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeCustomForm"), customFormID)
	var form EmployeeCustomForm
	resp, err := s.client.getAndDecode(ctx, url, nil, &form)
	if err != nil {
		return nil, resp, err
	}
	return &form, resp, nil
}

// CreateCustomForm creates a new custom form record.
func (s *EmployeeOtherService) CreateCustomForm(ctx context.Context, companyID int, form *EmployeeCustomForm) (*EmployeeCustomForm, *Response, error) {
	s.client.logger.Info("creating employee custom form", zap.Int("company_id", companyID))

	url := s.client.odataURL(companyID, "EmployeeCustomForm")
	var created EmployeeCustomForm
	resp, err := s.client.postAndDecode(ctx, url, form, &created)
	if err != nil {
		return nil, resp, err
	}
	return &created, resp, nil
}

// UpdateCustomForm updates an existing custom form record.
func (s *EmployeeOtherService) UpdateCustomForm(ctx context.Context, companyID, customFormID int, form *EmployeeCustomForm) (*Response, error) {
	s.client.logger.Info("updating employee custom form", zap.Int("company_id", companyID), zap.Int("custom_form_id", customFormID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeCustomForm"), customFormID)
	return s.client.patch(ctx, url, form)
}

// DeleteCustomForm deletes a custom form record.
func (s *EmployeeOtherService) DeleteCustomForm(ctx context.Context, companyID, customFormID int) (*Response, error) {
	s.client.logger.Info("deleting employee custom form", zap.Int("company_id", companyID), zap.Int("custom_form_id", customFormID))

	url := fmt.Sprintf("%s(%d)", s.client.odataURL(companyID, "EmployeeCustomForm"), customFormID)
	return s.client.del(ctx, url)
}

package payspace

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

// LookupValue represents a generic lookup entry shared by many lookup datasets.
type LookupValue struct {
	Id          int    `json:"Id"`
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
}

// CustomField represents a custom field definition for an entity type.
type CustomField struct {
	CustomFieldId int    `json:"CustomFieldId"`
	EntityType    string `json:"EntityType,omitempty"`
	FieldName     string `json:"FieldName,omitempty"`
	FieldType     string `json:"FieldType,omitempty"`
	Caption       string `json:"Caption,omitempty"`
}

// CustomFieldOption represents an option for a custom field dropdown.
type CustomFieldOption struct {
	OptionId    int    `json:"OptionId"`
	FieldId     int    `json:"FieldId,omitempty"`
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
}

// LookupService handles lookup-related API operations.
type LookupService struct {
	client *Client
}

// --- Generic Lookup ---

// ListLookup retrieves values from a named lookup dataset.
// This is the workhorse method: all 30+ lookup datasets use GET /Lookup/{DataSetName}.
// Examples: "ActivityCode", "LeaveReason", "LeaveType", "CompanyLegalBody",
// "AppealOutcome", "AppealReason", "AwardFavour", "LeaveBucket", etc.
func (s *LookupService) ListLookup(ctx context.Context, companyID int, lookupName string, params *QueryParams) ([]LookupValue, *Response, error) {
	s.client.logger.Info("listing lookup values",
		zap.Int("company_id", companyID),
		zap.String("lookup_name", lookupName),
	)
	if params != nil {
		s.client.logger.Debug("list lookup query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("Lookup/%s", lookupName))
	var result ListResult[LookupValue]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// --- Convenience Methods ---

// ListActivityCodes retrieves activity code lookup values.
func (s *LookupService) ListActivityCodes(ctx context.Context, companyID int, params *QueryParams) ([]LookupValue, *Response, error) {
	return s.ListLookup(ctx, companyID, "ActivityCode", params)
}

// ListLeaveReasons retrieves leave reason lookup values.
func (s *LookupService) ListLeaveReasons(ctx context.Context, companyID int, params *QueryParams) ([]LookupValue, *Response, error) {
	return s.ListLookup(ctx, companyID, "LeaveReason", params)
}

// ListLeaveTypes retrieves leave type lookup values.
func (s *LookupService) ListLeaveTypes(ctx context.Context, companyID int, params *QueryParams) ([]LookupValue, *Response, error) {
	return s.ListLookup(ctx, companyID, "LeaveType", params)
}

// ListCompanyLegalBodies retrieves company legal body lookup values.
func (s *LookupService) ListCompanyLegalBodies(ctx context.Context, companyID int, params *QueryParams) ([]LookupValue, *Response, error) {
	return s.ListLookup(ctx, companyID, "CompanyLegalBody", params)
}

// --- Custom Fields ---

// ListCustomFields retrieves custom field definitions for a given entity type.
func (s *LookupService) ListCustomFields(ctx context.Context, companyID int, entityType string, params *QueryParams) ([]CustomField, *Response, error) {
	s.client.logger.Info("listing custom fields",
		zap.Int("company_id", companyID),
		zap.String("entity_type", entityType),
	)
	if params != nil {
		s.client.logger.Debug("list custom fields query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("CustomField/%s", entityType))
	var result ListResult[CustomField]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListCustomFieldOptions retrieves dropdown options for a specific custom field.
func (s *LookupService) ListCustomFieldOptions(ctx context.Context, companyID, customFieldID int, params *QueryParams) ([]CustomFieldOption, *Response, error) {
	s.client.logger.Info("listing custom field options",
		zap.Int("company_id", companyID),
		zap.Int("custom_field_id", customFieldID),
	)
	if params != nil {
		s.client.logger.Debug("list custom field options query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("CustomFieldLookupOption/%d", customFieldID))
	var result ListResult[CustomFieldOption]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListCascadingOptions retrieves cascading custom field options filtered by a parent option.
func (s *LookupService) ListCascadingOptions(ctx context.Context, companyID, customFieldID, parentOptionID int, params *QueryParams) ([]CustomFieldOption, *Response, error) {
	s.client.logger.Info("listing cascading custom field options",
		zap.Int("company_id", companyID),
		zap.Int("custom_field_id", customFieldID),
		zap.Int("parent_option_id", parentOptionID),
	)
	if params != nil {
		s.client.logger.Debug("list cascading options query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("CascadingCustomFieldLookupOption/%d/%d", customFieldID, parentOptionID))
	var result ListResult[CustomFieldOption]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListCustomFormFields retrieves custom form field definitions for a form category.
func (s *LookupService) ListCustomFormFields(ctx context.Context, companyID, formCategoryID int, params *QueryParams) ([]CustomField, *Response, error) {
	s.client.logger.Info("listing custom form fields",
		zap.Int("company_id", companyID),
		zap.Int("form_category_id", formCategoryID),
	)
	if params != nil {
		s.client.logger.Debug("list custom form fields query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("CustomFormField/%d", formCategoryID))
	var result ListResult[CustomField]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

// ListCustomFormFieldOptions retrieves dropdown options for a custom form field.
func (s *LookupService) ListCustomFormFieldOptions(ctx context.Context, companyID, customFormFieldID int, params *QueryParams) ([]CustomFieldOption, *Response, error) {
	s.client.logger.Info("listing custom form field options",
		zap.Int("company_id", companyID),
		zap.Int("custom_form_field_id", customFormFieldID),
	)
	if params != nil {
		s.client.logger.Debug("list custom form field options query params", zap.String("params", params.Encode()))
	}

	url := s.client.odataURL(companyID, fmt.Sprintf("CustomFormFieldLookupOption/%d", customFormFieldID))
	var result ListResult[CustomFieldOption]
	resp, err := s.client.getAndDecode(ctx, url, params, &result)
	if err != nil {
		return nil, resp, err
	}
	return result.Value, resp, nil
}

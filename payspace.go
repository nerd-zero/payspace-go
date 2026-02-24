// Package payspace provides a Go client for the Deel Local Payroll (PaySpace) API v2.0.
//
// Usage:
//
//	client, err := payspace.NewClient("client_id", "client_secret",
//	    payspace.WithEnvironment(payspace.Production),
//	    payspace.WithLogLevel(zapcore.InfoLevel),
//	)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
//	employees, resp, err := client.Employees.List(ctx, companyID, nil)
package payspace

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

const (
	defaultUserAgent = "payspace.com"
	defaultTimeout   = 30 * time.Second
	defaultAPIVer    = "v2.0"
	defaultPageSize  = 100
	maxPageSize      = 100
)

// Environment selects between staging and production base URLs.
type Environment int

const (
	Staging Environment = iota
	Production
)

// Scope defines the API access scope requested during authentication.
type Scope string

const (
	ScopeFullAccess Scope = "api.full_access"
	ScopeReadOnly   Scope = "api.read_only"
	ScopeUpdate     Scope = "api.update"
	ScopeCreate     Scope = "api.create"
)

var envBaseURLs = map[Environment]struct{ identity, api string }{
	Staging:    {"https://staging-identity.yourhcm.com", "https://apistaging.payspace.com"},
	Production: {"https://identity.yourhcm.com", "https://api.payspace.com"},
}

// Client manages communication with the PaySpace API.
type Client struct {
	Employees           *EmployeeService
	EmployeePayroll     *EmployeePayrollService
	EmployeePayslips    *EmployeePayslipService
	EmployeeLeave       *EmployeeLeaveService
	EmployeeSuspensions *EmployeeSuspensionService
	EmployeeOther       *EmployeeOtherService
	EmployeeSkills      *EmployeeSkillsService
	EmployeeCosting     *EmployeeCostingService
	EmployeePerformance *EmployeePerformanceService
	Companies           *CompanyService
	CompanyPayroll      *CompanyPayrollService
	CompanyConfig       *CompanyConfigService
	Lookups             *LookupService
	Metadata            *MetadataService
	FileUpload          *FileUploadService
	Webhooks            *WebhookService

	httpClient  *http.Client
	auth        *tokenManager
	logger      *zap.Logger
	apiBaseURL  string
	identityURL string
	apiVersion  string
	userAgent   string
	rateLimit   *rateLimitTracker
}

// NewClient creates a new PaySpace API client.
// clientID and clientSecret are obtained from the PaySpace integrations settings.
func NewClient(clientID, clientSecret string, opts ...Option) (*Client, error) {
	if clientID == "" {
		return nil, fmt.Errorf("payspace: clientID is required")
	}
	if clientSecret == "" {
		return nil, fmt.Errorf("payspace: clientSecret is required")
	}

	cfg := defaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}

	logger := cfg.logger
	if logger == nil {
		var err error
		logger, err = newDefaultLogger(cfg.logLevel)
		if err != nil {
			return nil, fmt.Errorf("payspace: failed to create logger: %w", err)
		}
	}

	identityURL := cfg.identityURL
	apiBaseURL := cfg.apiBaseURL
	if identityURL == "" || apiBaseURL == "" {
		urls := envBaseURLs[cfg.environment]
		if identityURL == "" {
			identityURL = urls.identity
		}
		if apiBaseURL == "" {
			apiBaseURL = urls.api
		}
	}

	c := &Client{
		logger:      logger,
		apiBaseURL:  strings.TrimRight(apiBaseURL, "/"),
		identityURL: strings.TrimRight(identityURL, "/"),
		apiVersion:  cfg.apiVersion,
		userAgent:   cfg.userAgent,
		rateLimit:   newRateLimitTracker(),
	}

	c.auth = newTokenManager(tokenManagerConfig{
		clientID:     clientID,
		clientSecret: clientSecret,
		scope:        cfg.scope,
		identityURL:  c.identityURL,
		logger:       logger,
	})

	transport := buildTransport(cfg, c.auth, logger, c.rateLimit)
	c.httpClient = &http.Client{
		Timeout:   cfg.timeout,
		Transport: transport,
	}

	c.initServices()
	logger.Info("payspace client initialized",
		zap.String("api_base_url", c.apiBaseURL),
		zap.String("identity_url", c.identityURL),
		zap.String("api_version", c.apiVersion),
	)

	return c, nil
}

func (c *Client) initServices() {
	c.Employees = &EmployeeService{client: c}
	c.EmployeePayroll = &EmployeePayrollService{client: c}
	c.EmployeePayslips = &EmployeePayslipService{client: c}
	c.EmployeeLeave = &EmployeeLeaveService{client: c}
	c.EmployeeSuspensions = &EmployeeSuspensionService{client: c}
	c.EmployeeOther = &EmployeeOtherService{client: c}
	c.EmployeeSkills = &EmployeeSkillsService{client: c}
	c.EmployeeCosting = &EmployeeCostingService{client: c}
	c.EmployeePerformance = &EmployeePerformanceService{client: c}
	c.Companies = &CompanyService{client: c}
	c.CompanyPayroll = &CompanyPayrollService{client: c}
	c.CompanyConfig = &CompanyConfigService{client: c}
	c.Lookups = &LookupService{client: c}
	c.Metadata = &MetadataService{client: c}
	c.FileUpload = &FileUploadService{client: c}
	c.Webhooks = &WebhookService{client: c}
}

// odataURL builds an OData endpoint URL: {apiBase}/odata/{version}/{companyID}/{entity}
func (c *Client) odataURL(companyID int, entity string) string {
	return fmt.Sprintf("%s/odata/%s/%d/%s", c.apiBaseURL, c.apiVersion, companyID, entity)
}

// odataURLv1 builds an OData endpoint URL using API v1.0.
func (c *Client) odataURLv1(companyID int, entity string) string {
	return fmt.Sprintf("%s/odata/v1.0/%d/%s", c.apiBaseURL, companyID, entity)
}

// apiURL builds a non-OData API URL: {apiBase}/{version}/{companyID}/{path}
func (c *Client) apiURL(companyID int, path string) string {
	return fmt.Sprintf("%s/%s/%d/%s", c.apiBaseURL, c.apiVersion, companyID, path)
}

// Response wraps an http.Response and includes common PaySpace metadata.
type Response struct {
	*http.Response
	ODataContext string `json:"@odata.context,omitempty"`
	ODataCount   *int   `json:"@odata.count,omitempty"`
	RateLimit    RateLimitInfo
}

// newResponse wraps an http.Response.
func newResponse(r *http.Response) *Response {
	resp := &Response{Response: r}
	resp.RateLimit = parseRateLimitHeaders(r.Header)
	return resp
}

// ListResult holds a collection response with OData metadata.
type ListResult[T any] struct {
	Context string `json:"@odata.context,omitempty"`
	Count   *int   `json:"@odata.count,omitempty"`
	Value   []T    `json:"value"`
}

// do executes an HTTP request and returns the wrapped Response.
func (c *Client) do(ctx context.Context, req *http.Request) (*Response, error) {
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("payspace: request failed: %w", err)
	}

	response := newResponse(resp)
	c.rateLimit.update(response.RateLimit)

	if err := checkResponse(resp); err != nil {
		resp.Body.Close()
		return response, err
	}

	return response, nil
}

// get performs a GET request.
func (c *Client) get(ctx context.Context, rawURL string, query *QueryParams) (*Response, error) {
	if query != nil {
		rawURL = rawURL + "?" + query.Encode()
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("payspace: failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	return c.do(ctx, req)
}

// post performs a POST request with a JSON body.
func (c *Client) post(ctx context.Context, rawURL string, body any) (*Response, error) {
	return c.doWithBody(ctx, http.MethodPost, rawURL, body)
}

// patch performs a PATCH request with a JSON body.
func (c *Client) patch(ctx context.Context, rawURL string, body any) (*Response, error) {
	return c.doWithBody(ctx, http.MethodPatch, rawURL, body)
}

// del performs a DELETE request.
func (c *Client) del(ctx context.Context, rawURL string) (*Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("payspace: failed to create request: %w", err)
	}
	return c.do(ctx, req)
}

func (c *Client) doWithBody(ctx context.Context, method, rawURL string, body any) (*Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("payspace: failed to marshal request body: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, method, rawURL, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, fmt.Errorf("payspace: failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return c.do(ctx, req)
}

// decodeJSON reads and decodes the response body into v.
func decodeJSON(resp *Response, v any) error {
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("payspace: failed to read response body: %w", err)
	}
	if len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("payspace: failed to decode response: %w", err)
	}
	return nil
}

// getAndDecode is a convenience that performs a GET and decodes the JSON response.
func (c *Client) getAndDecode(ctx context.Context, rawURL string, query *QueryParams, v any) (*Response, error) {
	resp, err := c.get(ctx, rawURL, query)
	if err != nil {
		return resp, err
	}
	if err := decodeJSON(resp, v); err != nil {
		return resp, err
	}
	return resp, nil
}

// postAndDecode is a convenience that performs a POST and decodes the JSON response.
func (c *Client) postAndDecode(ctx context.Context, rawURL string, body, v any) (*Response, error) {
	resp, err := c.post(ctx, rawURL, body)
	if err != nil {
		return resp, err
	}
	if err := decodeJSON(resp, v); err != nil {
		return resp, err
	}
	return resp, nil
}

// patchAndDecode performs a PATCH and decodes the JSON response.
func (c *Client) patchAndDecode(ctx context.Context, rawURL string, body, v any) (*Response, error) {
	resp, err := c.patch(ctx, rawURL, body)
	if err != nil {
		return resp, err
	}
	if err := decodeJSON(resp, v); err != nil {
		return resp, err
	}
	return resp, nil
}

// addQueryParams appends arbitrary key=value parameters to a URL.
func addQueryParams(rawURL string, params map[string]string) string {
	if len(params) == 0 {
		return rawURL
	}
	v := url.Values{}
	for key, val := range params {
		v.Set(key, val)
	}
	sep := "?"
	if strings.Contains(rawURL, "?") {
		sep = "&"
	}
	return rawURL + sep + v.Encode()
}

// GroupCompany represents a group of companies returned during authentication.
type GroupCompany struct {
	GroupID          int       `json:"group_id"`
	GroupDescription string    `json:"group_description"`
	Companies        []Company `json:"companies"`
}

// Company is a company within a group.
type Company struct {
	CompanyID   int    `json:"company_id"`
	CompanyName string `json:"company_name"`
	CompanyCode string `json:"company_code"`
}

// Companies returns the list of companies the client has access to.
// This is populated after the first successful authentication.
func (c *Client) Companies_() ([]GroupCompany, error) {
	c.auth.mu.RLock()
	defer c.auth.mu.RUnlock()
	if c.auth.companies == nil {
		return nil, fmt.Errorf("payspace: not authenticated yet; make an API call first")
	}
	return c.auth.companies, nil
}

// TokenInfo returns current token metadata for debugging.
func (c *Client) TokenInfo() (expiresAt time.Time, hasToken bool) {
	c.auth.mu.RLock()
	defer c.auth.mu.RUnlock()
	return c.auth.expiresAt, c.auth.token != ""
}

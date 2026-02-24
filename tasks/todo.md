# PaySpace Go SDK - Implementation Plan

## Architecture Overview

A production-grade Go SDK for the Deel Local Payroll (PaySpace) API v2.0, following
patterns from popular Go SDKs (stripe-go, go-github). Service-oriented with functional
options, structured zap logging, auto-refreshing OAuth2 tokens, OData query builder,
and full API coverage.

---

## Directory Structure

```
payspace-go/
├── .tool-versions                 # asdf: golang 1.26
├── go.mod / go.sum
├── CLAUDE.md
├── docs/payspace-api/             # (existing API docs)
├── tasks/
│
├── payspace.go                    # Client, NewClient, top-level types
├── payspace_test.go               # Client tests
├── options.go                     # Functional options (WithLogger, WithBaseURL, etc.)
├── auth.go                        # Token management (auto-refresh)
├── auth_test.go
├── errors.go                      # Error types (APIError, RateLimitError, AuthError)
├── errors_test.go
├── odata.go                       # OData query builder ($filter, $top, $skip, etc.)
├── odata_test.go
├── transport.go                   # HTTP transport with middleware chain
├── transport_test.go
├── logging.go                     # Zap logging helpers & request/response logging
├── ratelimit.go                   # Rate limit tracking from response headers
│
├── employee.go                    # EmployeeService (biographical, address, status, position)
├── employee_test.go
├── employee_payroll.go            # EmployeePayrollService (pay rates, recurring components)
├── employee_payroll_test.go
├── employee_payslip.go            # EmployeePayslipService (payslips, PDFs, edit payslip)
├── employee_payslip_test.go
├── employee_leave.go              # EmployeeLeaveService (adjustments, setup, applications)
├── employee_leave_test.go
├── employee_suspension.go         # EmployeeSuspensionService
├── employee_suspension_test.go
├── employee_other.go              # EmployeeOtherService (notes, OOO, custom forms)
├── employee_other_test.go
├── employee_skills.go             # EmployeeSkillsService (training, qualifications, skills)
├── employee_skills_test.go
├── employee_costing.go            # EmployeeCostingService (recurring costing)
├── employee_costing_test.go
├── employee_performance.go        # EmployeePerformanceService (reviews, KPAs, templates)
├── employee_performance_test.go
│
├── company.go                     # CompanyService (billing, custom forms, config)
├── company_test.go
├── company_payroll.go             # CompanyPayrollService (frequencies, runs, components)
├── company_payroll_test.go
├── company_config.go              # CompanyConfigService (org positions, regions, GL)
├── company_config_test.go
│
├── lookup.go                      # LookupService (all lookup endpoints)
├── lookup_test.go
├── metadata.go                    # MetadataService ($metadata)
├── metadata_test.go
├── fileupload.go                  # FileUploadService
├── fileupload_test.go
├── webhook.go                     # WebhookService (webhook errors)
├── webhook_test.go
│
└── testutil_test.go               # Shared test helpers (mock server, fixtures)
```

All files in the root package `payspace` — flat layout, no sub-packages.
This follows the stripe-go pattern (single package, service-oriented).

---

## Core Design

### 1. Client (payspace.go)

```go
type Client struct {
    Employees          *EmployeeService
    EmployeePayroll    *EmployeePayrollService
    EmployeePayslips   *EmployeePayslipService
    EmployeeLeave      *EmployeeLeaveService
    EmployeeSuspension *EmployeeSuspensionService
    EmployeeOther      *EmployeeOtherService
    EmployeeSkills     *EmployeeSkillsService
    EmployeeCosting    *EmployeeCostingService
    EmployeePerformance *EmployeePerformanceService
    Companies          *CompanyService
    CompanyPayroll     *CompanyPayrollService
    CompanyConfig      *CompanyConfigService
    Lookups            *LookupService
    Metadata           *MetadataService
    FileUpload         *FileUploadService
    Webhooks           *WebhookService

    // internal
    httpClient  *http.Client
    auth        *tokenManager
    logger      *zap.Logger
    apiBaseURL  string
    identityURL string
    rateLimiter *rateLimitTracker
}

func NewClient(clientID, clientSecret string, opts ...Option) (*Client, error)
```

### 2. Options (options.go)

```go
type Option func(*clientConfig)

func WithEnvironment(env Environment)     Option  // Staging or Production
func WithAPIBaseURL(url string)           Option  // Custom API base URL
func WithIdentityURL(url string)          Option  // Custom identity URL
func WithScope(scope Scope)               Option  // API scope
func WithHTTPClient(client *http.Client)  Option  // Custom HTTP client
func WithLogger(logger *zap.Logger)       Option  // Custom zap logger
func WithLogLevel(level zapcore.Level)    Option  // Set log level
func WithTimeout(d time.Duration)         Option  // HTTP timeout
func WithRetry(maxRetries int, backoff time.Duration) Option  // Retry config
func WithRateLimitWait(wait bool)         Option  // Auto-wait on 429
func WithAPIVersion(version string)       Option  // "v2.0" or "v1.0"
func WithUserAgent(ua string)             Option  // Custom user agent (default: payspace.com)
```

### 3. Auth (auth.go) — Auto-refresh tokens

```go
type tokenManager struct {
    mu           sync.RWMutex
    token        string
    expiresAt    time.Time
    companies    []GroupCompany
    clientID     string
    clientSecret string
    scope        Scope
    identityURL  string
    httpClient   *http.Client
    logger       *zap.Logger
    refreshBefore time.Duration  // refresh 5 min before expiry
}

func (tm *tokenManager) Token(ctx context.Context) (string, error)
// Returns cached token if valid, otherwise refreshes.
// Proactive: refreshes when < refreshBefore from expiry.
// Reactive: called again on 401 with force=true.
```

### 4. Transport (transport.go) — Middleware chain

The HTTP client uses a `RoundTripper` chain:
1. **authTransport** — injects Bearer token, handles 401 retry
2. **loggingTransport** — logs request/response at appropriate levels
3. **retryTransport** — exponential backoff for 429/5xx
4. **rateLimitTransport** — tracks rate limit headers

### 5. OData Query Builder (odata.go)

```go
type QueryParams struct { ... }

func NewQuery() *QueryParams
func (q *QueryParams) Filter(expr string) *QueryParams
func (q *QueryParams) Top(n int) *QueryParams
func (q *QueryParams) Skip(n int) *QueryParams
func (q *QueryParams) OrderBy(field string) *QueryParams
func (q *QueryParams) Select(fields ...string) *QueryParams
func (q *QueryParams) Expand(fields ...string) *QueryParams
func (q *QueryParams) Count() *QueryParams
func (q *QueryParams) Encode() string
```

### 6. Error Types (errors.go)

```go
type APIError struct {
    StatusCode int
    Message    string
    Details    []ErrorDetail
    RequestID  string
}

type RateLimitError struct {
    *APIError
    Limit     int
    Remaining int
    Reset     time.Time
}

type AuthError struct { ... }
```

### 7. Service Pattern (example: employee.go)

```go
type EmployeeService struct {
    client *Client
}

// List methods return ([]T, *Response, error)
// Get methods return (*T, *Response, error)
// Create/Update return (*T, *Response, error)
// Delete returns (*Response, error)

type Response struct {
    *http.Response
    ODataContext string
    ODataCount   *int
    RateLimit    RateLimitInfo
}

type ListOptions struct {
    QueryParams *QueryParams
    CompanyID   int
}

func (s *EmployeeService) List(ctx context.Context, companyID int, opts *QueryParams) ([]Employee, *Response, error)
func (s *EmployeeService) Get(ctx context.Context, companyID, employeeID int) (*Employee, *Response, error)
func (s *EmployeeService) Create(ctx context.Context, companyID int, emp *Employee) (*Employee, *Response, error)
func (s *EmployeeService) Update(ctx context.Context, companyID, employeeID int, emp *Employee) (*Response, error)
// ... etc
```

### 8. Logging Strategy

All operations logged at appropriate levels:
- **Debug**: Full request/response bodies, headers, OData params
- **Info**: API call method + path, status code, duration
- **Warn**: Rate limit approaching (< 10%), retry attempts, token refresh
- **Error**: Failed requests, auth failures, unexpected errors

---

## Implementation Phases

### Phase 1: Foundation
- [x] Set up project (go.mod, .tool-versions, asdf golang 1.26.0)
- [x] Core types & configuration (payspace.go, options.go)
- [x] Logging infrastructure (logging.go)
- [x] Error types (errors.go)
- [x] OData query builder (odata.go)
- [x] HTTP transport with middleware chain (transport.go)
- [x] Auth with auto-refresh (auth.go)
- [x] Rate limit tracking (ratelimit.go)
- [x] Test utilities (testutil_test.go)
- [x] Tests for all foundation code

### Phase 2: Employee Services
- [x] EmployeeService — biographical, address, status, position, photo
- [x] EmployeePayrollService — pay rates, all recurring components
- [x] EmployeePayslipService — payslips, PDFs, edit, take-on, lump sum, tax certs
- [x] EmployeeLeaveService — adjustments, setup, entitlements, applications
- [x] EmployeeSuspensionService
- [x] EmployeeOtherService — notes, out of office, custom forms
- [x] EmployeeSkillsService — training, qualifications, skills
- [x] EmployeeCostingService — recurring costing splits/headers
- [x] EmployeePerformanceService — reviews, KPAs, templates
- [x] Tests for all employee services

### Phase 3: Company & Other Services
- [x] CompanyService — billing, custom forms
- [x] CompanyPayrollService — frequencies, runs, components
- [x] CompanyConfigService — org positions, regions, GL, EFT, table builder
- [x] LookupService — all standard + custom field lookups
- [x] MetadataService — $metadata endpoint
- [x] FileUploadService — multipart file upload
- [x] WebhookService — webhook error retrieval
- [x] Tests for all company & other services

### Phase 4: Verification
- [x] All tests pass (124/124)
- [x] go vet clean
- [x] Verify logging at all levels works correctly
- [x] Review: "Would a staff engineer approve this?" ✓

---

## Results

- **24 source files** (4,864 lines)
- **18 test files** (3,595 lines)
- **124 tests**, all passing
- `go build ./...` — clean
- `go vet ./...` — clean

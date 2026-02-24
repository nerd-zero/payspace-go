# payspace-go

Go client library for the [Deel Local Payroll (PaySpace)](https://www.payspace.com/) API v2.0.

## Features

- OAuth 2.0 client credentials authentication with automatic token refresh
- Fluent OData query builder for filtering, pagination, sorting, and field selection
- Automatic rate limiting with configurable wait behavior
- Retries with exponential backoff for transient errors
- Structured logging via [zap](https://github.com/uber-go/zap)
- Full coverage of employee, payroll, leave, and company configuration endpoints

## Installation

```bash
go get github.com/nerd-zero/payspace-go
```

Requires Go 1.25 or later.

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/nerd-zero/payspace-go"
    "go.uber.org/zap/zapcore"
)

func main() {
    client, err := payspace.NewClient(
        "your-client-id",
        "your-client-secret",
        payspace.WithEnvironment(payspace.Production),
        payspace.WithLogLevel(zapcore.WarnLevel),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx := context.Background()
    companyID := 12345

    employees, _, err := client.Employees.List(ctx, companyID, nil)
    if err != nil {
        log.Fatal(err)
    }

    for _, emp := range employees {
        fmt.Printf("%s %s (%s)\n", emp.FirstName, emp.Surname, emp.EmployeeNumber)
    }
}
```

## Authentication

The SDK uses OAuth 2.0 client credentials flow. Obtain your `clientID` and `clientSecret` from the PaySpace integrations settings.

Authentication is handled transparently -- the SDK acquires a token on the first API call and refreshes it automatically (5-minute buffer before expiry). You never need to manage tokens manually.

After the first API call, you can access the list of companies your credentials have access to:

```go
groups, err := client.Companies_()
if err != nil {
    log.Fatal(err)
}
for _, group := range groups {
    for _, company := range group.Companies {
        fmt.Printf("%s (ID: %d)\n", company.CompanyName, company.CompanyID)
    }
}
```

### Scopes

Control the access level with `WithScope`:

```go
payspace.WithScope(payspace.ScopeFullAccess) // default
payspace.WithScope(payspace.ScopeReadOnly)
payspace.WithScope(payspace.ScopeUpdate)
payspace.WithScope(payspace.ScopeCreate)
```

## Configuration

All options are passed to `NewClient` as functional options:

| Option | Description | Default |
|--------|-------------|---------|
| `WithEnvironment(env)` | `Staging` or `Production` | `Staging` |
| `WithScope(scope)` | API access scope | `ScopeFullAccess` |
| `WithTimeout(d)` | HTTP request timeout | `30s` |
| `WithRetry(max, backoff)` | Max retries and base backoff for transient errors | `3`, `500ms` |
| `WithRateLimitWait(bool)` | Automatically wait when rate limited (429) | `true` |
| `WithLogLevel(level)` | Minimum zap log level | `InfoLevel` |
| `WithLogger(logger)` | Custom `*zap.Logger` instance | auto-created |
| `WithHTTPClient(client)` | Custom `*http.Client` | default |
| `WithAPIVersion(ver)` | API version string | `"v2.0"` |
| `WithUserAgent(ua)` | User-Agent header | `"payspace-go/{version}"` |
| `WithAPIBaseURL(url)` | Override API base URL | environment-based |
| `WithIdentityURL(url)` | Override identity/auth URL | environment-based |

```go
client, err := payspace.NewClient(clientID, clientSecret,
    payspace.WithEnvironment(payspace.Production),
    payspace.WithTimeout(60 * time.Second),
    payspace.WithRetry(5, time.Second),
    payspace.WithLogLevel(zapcore.WarnLevel),
)
```

## Services

All API methods follow the pattern `client.ServiceName.Method(ctx, companyID, ...)`.

| Service | Description |
|---------|-------------|
| `Employees` | Employee biographical data, addresses, employment statuses, positions |
| `EmployeePayroll` | Pay rates, recurring components (bonus, medical, pension, loans, garnishees), component extraction |
| `EmployeePayslips` | Payslips, payslip lines, PDFs, edit payslips, batch operations, take-ons, lump sums, tax certificates |
| `EmployeeLeave` | Leave adjustments, setups, entitlements, applications |
| `EmployeeSuspensions` | Employee suspension records |
| `EmployeeOther` | Out-of-office, notes, custom forms |
| `EmployeeSkills` | Training, qualifications, skills |
| `EmployeeCosting` | Costing splits, recurring costings |
| `EmployeePerformance` | Performance review headers, KPAs, templates (read-only) |
| `Companies` | Billing documents, custom forms |
| `CompanyPayroll` | Payroll frequencies, runs, company-level components |
| `CompanyConfig` | Organization positions, regions, general ledger, EFT outbox, costing project activities, component table definitions |
| `Lookups` | 30+ named lookup datasets, custom fields and options |
| `Metadata` | OData `$metadata` XML document |
| `FileUpload` | Multipart file uploads |
| `Webhooks` | Webhook error listing |

## OData Query Builder

Build queries with the fluent `QueryParams` builder:

```go
// Filter and sort
query := payspace.NewQuery().
    Filter("Surname eq 'Smith'").
    OrderBy("FirstName asc").
    Top(10)

employees, _, err := client.Employees.List(ctx, companyID, query)

// Pagination with count
query := payspace.NewQuery().
    Top(50).
    Skip(100).
    Count()

// Select specific fields
query := payspace.NewQuery().
    Select("EmployeeNumber", "FirstName", "Surname")

// Expand related entities
query := payspace.NewQuery().
    Expand("OrganizationGroups")

// Custom parameters (e.g., frequency and period for payroll endpoints)
query := payspace.NewQuery().
    Param("frequency", "1").
    Param("period", "6")
```

The maximum page size is 100 records (enforced by `Top`).

## Effective Date Queries

Many entities support three listing variants:

```go
// Active records only (default)
employees, _, _ := client.Employees.List(ctx, companyID, nil)

// Records effective as of a specific date
employees, _, _ := client.Employees.ListByEffectiveDate(ctx, companyID, "2025-01-01", nil)

// All records including inactive
employees, _, _ := client.Employees.ListAll(ctx, companyID, nil)
```

This pattern applies to employees, employment statuses, positions, pay rates, leave setups, leave adjustments, leave applications, organization positions, and more.

## Error Handling

The SDK returns typed errors with helper functions for common checks:

```go
employees, _, err := client.Employees.List(ctx, companyID, nil)
if err != nil {
    if payspace.IsNotFound(err) {
        // 404 - resource not found
    }
    if payspace.IsUnauthorized(err) {
        // 401 - usually auto-retried with token refresh
    }
    if payspace.IsRateLimited(err) {
        // 429 - usually auto-handled if WithRateLimitWait is true
    }
    if payspace.IsValidationError(err) {
        // 400 with validation details
        if apiErr, ok := err.(*payspace.APIError); ok {
            for _, detail := range apiErr.Details {
                fmt.Println("Validation:", detail.Message)
            }
        }
    }
    log.Fatal(err)
}
```

### Error Types

- **`APIError`** -- returned for all non-2xx responses. Contains `StatusCode`, `Message`, and `Details` (for validation errors).
- **`RateLimitError`** -- returned for 429 responses. Embeds `APIError` and adds `Limit`, `Remaining`, and `Reset` fields.

## Rate Limiting

Rate limiting is handled automatically by default (`WithRateLimitWait(true)`). The SDK:

1. Tracks `X-RateLimit-Limit`, `X-RateLimit-Remaining`, and `X-RateLimit-Reset` headers
2. Pre-emptively waits when remaining requests reach 0
3. Reactively waits and retries on 429 responses

Access rate limit info from any response:

```go
employees, resp, err := client.Employees.List(ctx, companyID, nil)
if err == nil {
    fmt.Printf("Rate limit: %d/%d (resets %s)\n",
        resp.RateLimit.Remaining,
        resp.RateLimit.Limit,
        resp.RateLimit.Reset.Format(time.RFC3339),
    )
}
```

## Retries

The SDK automatically retries requests with exponential backoff for:

- **5xx server errors** -- transient server failures
- **Network errors** -- connection timeouts, DNS failures
- **401 Unauthorized** -- triggers a token refresh and retries once

Configure with `WithRetry(maxRetries, baseBackoff)`. Default: 3 retries with 500ms base backoff (doubling each attempt).

## Logging

The SDK uses [zap](https://github.com/uber-go/zap) structured logging.

```go
// Simple: set log level
client, _ := payspace.NewClient(id, secret,
    payspace.WithLogLevel(zapcore.WarnLevel),
)

// Advanced: custom logger
logger, _ := payspace.NewLogger(payspace.LoggerConfig{
    Level:       zapcore.DebugLevel,
    Development: true,
    OutputPaths: []string{"stdout", "/var/log/payspace.log"},
})

client, _ := payspace.NewClient(id, secret,
    payspace.WithLogger(logger),
)
```

## Environments

```go
payspace.WithEnvironment(payspace.Staging)    // staging-identity.yourhcm.com / apistaging.payspace.com
payspace.WithEnvironment(payspace.Production) // identity.yourhcm.com / api.payspace.com
```

Override individual URLs for custom environments:

```go
payspace.WithAPIBaseURL("https://custom-api.example.com")
payspace.WithIdentityURL("https://custom-identity.example.com")
```

## Examples

See the [`examples/`](examples/) directory for complete, runnable programs:

| Example | Description |
|---------|-------------|
| [basic](examples/basic/) | Client setup, list employees, access companies |
| [employees](examples/employees/) | Employee CRUD operations |
| [odata-queries](examples/odata-queries/) | Advanced query builder usage |
| [payslips](examples/payslips/) | Payslip retrieval and payslip lines |
| [leave-management](examples/leave-management/) | Leave types, balances, and applications |
| [payroll-components](examples/payroll-components/) | Pay rates and recurring components |
| [company-config](examples/company-config/) | Company configuration and lookups |
| [error-handling](examples/error-handling/) | Error types and handling patterns |
| [edit-payslip](examples/edit-payslip/) | Edit payslips and batch operations |
| [file-upload](examples/file-upload/) | File uploads |
| [general-ledger](examples/general-ledger/) | General ledger and EFT outbox |

## Contributing

```bash
make test    # run tests with race detector
make vet     # run go vet
make lint    # run vet + test
```

## Releasing

The SDK uses [CalVer](https://calver.org/) (`0.YYYYMM.MICRO`) for versioning. The version is stored in the `VERSION` file and embedded into the SDK at compile time.

To create a new release:

```bash
make release
```

This will:
1. Bump the version based on the current date (same month increments MICRO, new month resets to `.001`)
2. Commit the updated `VERSION` file
3. Create a git tag (`v0.YYYYMM.MICRO`)
4. Push the commit and tag to GitHub

GitHub Actions will then automatically create a GitHub Release with a changelog.

To preview the next version without making changes:

```bash
./scripts/version.sh --dry-run
```

## License

See [LICENSE](LICENSE) for details.

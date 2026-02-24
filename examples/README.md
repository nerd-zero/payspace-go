# Examples

Runnable examples demonstrating common payspace-go SDK use cases.

## Setup

All examples read credentials from environment variables:

```bash
export PAYSPACE_CLIENT_ID="your-client-id"
export PAYSPACE_CLIENT_SECRET="your-client-secret"
export PAYSPACE_COMPANY_ID="12345"
```

## Running

```bash
go run examples/<name>/main.go
```

For example:

```bash
go run examples/basic/main.go
```

## Examples

| Directory | Description |
|-----------|-------------|
| [basic](basic/) | Client setup, list employees, access companies from auth token |
| [employees](employees/) | Employee CRUD: create, get by ID, update, filter by surname |
| [odata-queries](odata-queries/) | Query builder: filter, pagination, count, sort, select fields |
| [payslips](payslips/) | List payslips by year/month, payslip lines, tax certificates |
| [leave-management](leave-management/) | Leave types lookup, leave balances, create leave application |
| [payroll-components](payroll-components/) | Pay rates, company frequencies/runs, recurring components |
| [company-config](company-config/) | Organization positions, regions, lookups, custom fields |
| [error-handling](error-handling/) | Error types, helper functions, validation details |
| [edit-payslip](edit-payslip/) | Edit payslips: list, create, batch operations |
| [file-upload](file-upload/) | Upload files, receive temp attachment signature |
| [general-ledger](general-ledger/) | General ledger entries, extraction, EFT outbox |

All examples default to the **Staging** environment for safety. Change to `payspace.Production` for production use.

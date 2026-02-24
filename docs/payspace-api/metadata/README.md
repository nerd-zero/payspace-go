# Payspace API Metadata Reference

*Auto-generated from OData $metadata endpoint for company 6846 (Takealot staging)*

**Total entity types:** 176

## Entity Categories

- [Employee Entities](01-employee-entities.md) (61 types)
- [Company Entities](02-company-entities.md) (59 types)
- [Organization Entities](03-organization-entities.md) (6 types)
- [Component Entities](04-component-entities.md) (4 types)
- [Lookup Entities](05-lookup-datasets.md) (30 types)
- [Other Entities](06-other-entities.md) (16 types)

## Column Legend

| Column | Description |
|--------|-------------|
| **Type** | OData EDM type (String, Int32, Int64, Boolean, Date, DateTimeOffset, Decimal, etc.) |
| **Nullable** | Whether the field can be null |
| **Required** | `Yes` = always required, `Conditional` = required depending on other fields, `No` = optional |
| **Access** | `Read-only` = cannot set, `Editable` = can update, `Create-only` = set on POST only |
| **Lookup** | Name of the lookup endpoint that provides valid values for this field |
| **Max Length** | Maximum string length |
| **(PK)** | Primary key field |

## PaySpace Custom Metadata Attributes

| Attribute | Description |
|-----------|-------------|
| `ps:Required` | true, false, or conditional |
| `ps:Editable` | Whether field can be modified after creation |
| `ps:ReadOnly` | Field cannot be set on POST or PATCH |
| `ps:Lookup` | Field uses lookup values; value = lookup endpoint name |
| `ps:Hidden` | Field disabled by company setting |
| `ps:MaxLength` / `ps:MinLength` | String length constraints |
| `ps:MinValue` / `ps:MaxValue` | Numeric range constraints |
| `ps:Unique` | Value must be unique across records |
| `ps:RegExFormat` | Regex validation pattern |
| `ps:CompanySetting` | Field depends on a specific company setting code |
| `ps:Caption` | Display label |

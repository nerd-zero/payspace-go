# Employee > Costing

## Recurring Costing Split

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCostingSplit?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeRecurringCostingSplit` records based on the specified `company_id.`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRecurringCostingSplit",
    "value": [
        {
            "RecurringCostingSplitDetailId": 0,                     // Type="Int64"
            "OrganizationGroup": "string",                          // Type="String"
            "OrganizationGroupDescription": "string",               // Type="String"
            "EmployeeNumber": "string",                             // Type="String"
            "Percentage": 0.0,                                      // Type="Decimal"
            "EffectiveDate": "2019-01-01T00:00:00+02:00"            // Type="DateTime"
        }
    ]
}
```

---

### Get a collection as of an effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCostingSplit/effective/:effectivedate?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeRecurringCostingSplit` records as of the specified effective date for a `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection as of an effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRecurringCostingSplit",
    "value": [
        {
            "RecurringCostingSplitDetailId": 0,                     // Type="Int64"
            "OrganizationGroup": "string",                          // Type="String"
            "OrganizationGroupDescription": "string",               // Type="String"
            "EmployeeNumber": "string",                             // Type="String"
            "Percentage": 0.0,                                      // Type="Decimal"
            "EffectiveDate": "2019-01-01T00:00:00+02:00"            // Type="DateTime"
        }
    ]
}
```

---

### Get a collection of all records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCostingSplit/All?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of all `EmployeeRecurringCostingSplit` records for a `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection as of an effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRecurringCostingSplit",
    "value": [
        {
            "RecurringCostingSplitDetailId": 0,                     // Type="Int64"
            "OrganizationGroup": "string",                          // Type="String"
            "OrganizationGroupDescription": "string",               // Type="String"
            "EmployeeNumber": "string",                             // Type="String"
            "Percentage": 0.0,                                      // Type="Decimal"
            "EffectiveDate": "2019-01-01T00:00:00+02:00"            // Type="DateTime"
        }
    ]
}
```

---
## Recurring Costing

Contains endpoints related to employee recurring costing split management.

List of available `EmployeeRecurringCosting.BasedOnSplitOption` can be found on the `BasedOnSplitOption` Lookup Endpoint.

> Note. This endpoint only works from v2.0

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCosting?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeRecurringCostingSplit` records based on the specified `company_id.`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$orderby` | `{{$recurring-costing-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRecurringCosting",
    "value": [
        {
            "RecurringCostingSplitHeaderId": 0,                 // Type="Int64"
            "EmployeeNumber": "string",                         // Type="String"
            "FullName": "string",                               // Type="String"
            "EffectiveDate": "2024-01-01",                      // Type="Date"
            "BasedOnSplitOption": "string",                     // Type="String"
            "RecurringCostingSplitDetails": [
                {
                    "RecurringCostingSplitDetailId": 0,         // Type="Int64"
                    "OrganizationGroup": "string",              // Type="String"
                    "OrganizationGroupDescription": "string",   // Type="String"
                    "ProjectCode": "string",                    // Type="String"
                    "ProjectActivityCode": "string",            // Type="String"
                    "Percentage": 0.0                           // Type="Decimal"
                }
            ]
        }
    ]
}
```

---

### Get a collection as of an effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCosting/effective/:effectivedate?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeRecurringCostingSplit` records as of the specified effective date for a `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$orderby` | `{{$recurring-costing-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get a collection as of an effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRecurringCosting",
    "value": [
        {
            "RecurringCostingSplitHeaderId": 0,                 // Type="Int64"
            "EmployeeNumber": "string",                         // Type="String"
            "FullName": "string",                               // Type="String"
            "EffectiveDate": "2024-01-01",                      // Type="Date"
            "BasedOnSplitOption": "string",                     // Type="String"
            "RecurringCostingSplitDetails": [
                {
                    "RecurringCostingSplitDetailId": 0,         // Type="Int64"
                    "OrganizationGroup": "string",              // Type="String"
                    "OrganizationGroupDescription": "string",   // Type="String"
                    "ProjectCode": "string",                    // Type="String"
                    "ProjectActivityCode": "string",            // Type="String"
                    "Percentage": 0.0                           // Type="Decimal"
                }
            ]
        }
    ]
}
```

---

### Get a collection of all records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCosting/All?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of all `EmployeeRecurringCostingSplit` records for a `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$orderby` | `{{$recurring-costing-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get a collection of all records** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRecurringCosting",
    "value": [
        {
            "RecurringCostingSplitHeaderId": 0,                 // Type="Int64"
            "EmployeeNumber": "string",                         // Type="String"
            "FullName": "string",                               // Type="String"
            "EffectiveDate": "2024-01-01",                      // Type="Date"
            "BasedOnSplitOption": "string",                     // Type="String"
            "RecurringCostingSplitDetails": [
                {
                    "RecurringCostingSplitDetailId": 0,         // Type="Int64"
                    "OrganizationGroup": "string",              // Type="String"
                    "OrganizationGroupDescription": "string",   // Type="String"
                    "ProjectCode": "string",                    // Type="String"
                    "ProjectActivityCode": "string",            // Type="String"
                    "Percentage": 0.0                           // Type="Decimal"
                }
            ]
        }
    ]
}
```

---

### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCosting`

Create a single employee position record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
    // see "EmployeeRecurringCosting" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRecurringCosting/$entity",
    "RecurringCostingSplitHeaderId": 0,                 // Type="Int64"
    "EmployeeNumber": "string",                         // Type="String"
    "FullName": "string",                               // Type="String"
    "EffectiveDate": "2024-01-01",                      // Type="Date"
    "BasedOnSplitOption": "string",                     // Type="String"
    "RecurringCostingSplitDetails": [
        {
            "RecurringCostingSplitDetailId": 0,         // Type="Int64"
            "OrganizationGroup": "string",              // Type="String"
            "OrganizationGroupDescription": "string",   // Type="String"
            "ProjectCode": "string",                    // Type="String"
            "ProjectActivityCode": "string",            // Type="String"
            "Percentage": 0.0                           // Type="Decimal"
        }
    ]
}
```

---

### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCosting({{RecurringCostingSplitHeaderId}})`

Update a single employee position record based on the specified `EmployeePositionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
    // see "EmployeeRecurringCosting" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCosting({{RecurringCostingSplitHeaderId}})`

Retrieves employee position records based on the specified `PositionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single record** (Status: 200 OK)

---

### Delete a single detail record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringCosting/detail({{RecurringCostingSplitDetailId}})`

Retrieves employee position records based on the specified `PositionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single detail record** (Status: 200 OK)

---

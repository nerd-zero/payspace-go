# Company > Costing Project Activity

## Get a Single Record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CostingProjectActivity({{ProjectId}})`

Retrieves a single record of `CostingProjectActivity` information based on the specified `ProjectId`

**Example Response: Get a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CostingProjectActivity/$entity",
    "ProjectId": 0,                                     // Type="Int64"
    "CompanyId": 0,                                     // Type="Int64"
    "ProjectActivityDescription": "string",             // Type="String"
    "ProjectActivityCode": "string",                    // Type="String"
    "Glcode": "string",                                 // Type="String"
    "OrganizationGroup": "string",                      // Type="String"
    "OrganizationRegion": "string",                     // Type="String"
    "KeyAccountManagerEmployeeNumber": "string",        // Type="String"
    "MonthlyStaffBudget": 0,                            // Type="Int32"
    "MonthlyBudgetedHours": 0.00,                       // Type="Decimal"
    "HourlyCost": 0.00,                                 // Type="Decimal"
    "InactiveDate": "DateTime",                         // Type="DateTime"
    "DoNotShow": "Boolean",                             // Type="Boolean"
    "ActivityHourlyRate": 0.00,                         // Type="Decimal"
    "TsdefaultHours": 0.00,                             // Type="Decimal"
    "MaxHours": 0.00,                                   // Type="Decimal"
    "OrganizationCategory": "string",                   // Type="String"
    "CustomFields": [
        {
            "Code": "string",                           // Type="String"
            "Label": "string",                          // Type="String"
            "Value": "string"                           // Type="String"
        }
    ]
}
```

---
## Get a Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CostingProjectActivity?$filter={{filter}}&$skip={{skip}}&$top={{top}}`

StartFragment

Retrieves a list of `CostingProjectActivity` records based on the specified `company_id.`

EndFragment

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{filter}}` | Optional - Filter by field eg. "ProjectId eq 1001". See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$skip` | `{{skip}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$top` | `{{top}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CostingProjectActivity/$entity",
    "value": [
        {
            "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CostingProjectActivity/$entity",
            "ProjectId": 0,                                     // Type="Int64"
            "CompanyId": 0,                                     // Type="Int64"
            "ProjectActivityDescription": "string",             // Type="String"
            "ProjectActivityCode": "string",                    // Type="String"
            "Glcode": "string",                                 // Type="String"
            "OrganizationGroup": "string",                      // Type="String"
            "OrganizationRegion": "string",                     // Type="String"
            "KeyAccountManagerEmployeeNumber": "string",        // Type="String"
            "MonthlyStaffBudget": 0,                            // Type="Int32"
            "MonthlyBudgetedHours": 0.00,                       // Type="Decimal"
            "HourlyCost": 0.00,                                 // Type="Decimal"
            "InactiveDate": "DateTime",                         // Type="DateTime"
            "DoNotShow": "Boolean",                             // Type="Boolean"
            "ActivityHourlyRate": 0.00,                         // Type="Decimal"
            "TsdefaultHours": 0.00,                             // Type="Decimal"
            "MaxHours": 0.00,                                   // Type="Decimal"
            "OrganizationCategory": "string",                   // Type="String"
            "CustomFields": [
                {
                    "Code": "string",                           // Type="String"
                    "Label": "string",                          // Type="String"
                    "Value": "string"                           // Type="String"
                }
            ]
        }
    ]
}
```

---
## Create a Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CostingProjectActivity`

Creates a single `CostingProjectActivity` record for the specified `companyId`

**Request Body:**

```json
{
  //see "CostingProjectActivity" in metadata endpoint for available fields
}
```

**Example Response: Create a Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CostingProjectActivity/$entity",
    "ProjectId": 0,                                     // Type="Int64"
    "CompanyId": 0,                                     // Type="Int64"
    "ProjectActivityDescription": "string",             // Type="String"
    "ProjectActivityCode": "string",                    // Type="String"
    "Glcode": "string",                                 // Type="String"
    "OrganizationGroup": "string",                      // Type="String"
    "OrganizationRegion": "string",                     // Type="String"
    "KeyAccountManagerEmployeeNumber": "string",        // Type="String"
    "MonthlyStaffBudget": 0,                            // Type="Int32"
    "MonthlyBudgetedHours": 0.00,                       // Type="Decimal"
    "HourlyCost": 0.00,                                 // Type="Decimal"
    "InactiveDate": "DateTime",                         // Type="DateTime"
    "DoNotShow": "Boolean",                             // Type="Boolean"
    "ActivityHourlyRate": 0.00,                         // Type="Decimal"
    "TsdefaultHours": 0.00,                             // Type="Decimal"
    "MaxHours": 0.00,                                   // Type="Decimal"
    "OrganizationCategory": "string",                   // Type="String"
    "CustomFields": [
        {
            "Code": "string",                           // Type="String"
            "Label": "string",                          // Type="String"
            "Value": "string"                           // Type="String"
        }
    ]
}
```

---
## Update a Single Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CostingProjectActivity({{ProjectId}})`

Updates a single `CostingProjectActivity` record for the specified `ProjectId`

**Request Body:**

```json
{
  //see "CostingProjectActivity" in metadata endpoint for available fields
}
```

**Example Response: Update a Single Record** (Status: 204 No Content)

---
## Delete a Single Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CostingProjectActivity({{ProjectId}})`

Deletes a single `CostingProjectActivity` record for the specified `ProjectId`

**Example Response: Delete a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CostingProjectActivity/$entity",
    "Success": true             // Type="boolean"
}
```

---

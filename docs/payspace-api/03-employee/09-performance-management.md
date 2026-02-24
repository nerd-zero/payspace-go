# Employee > Performance Management

## Employee Review Header

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeReviewHeader?$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeReviewHeader` records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeReviewHeader",
    "value": [
        {
            "ProcessId": 0,                                     // Type="Int64"
            "EmployeeReviewHeaderId": 0,                        // Type="Int64"
            "RaterId": 0,                                       // Type="Int64"
            "RaterEmployeeNumber": "String",                    // Type="String"
            "RaterFirstName": "String",                         // Type="String"
            "RaterLastName": "String",                          // Type="String"
            "RaterType": "Manager",                             // Type="Enum" see: RaterTypes
            "OverallScore": 0.00,                               // Type="Decimal"
            "Weighting": 0.00                                   // Type="Decimal"
        }
    ]
}
```

---
## Employee Review KPA

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeReviewKpa?$skip={{$skip}}&$top={{$top}}`

Retrieves a list of `EmployeeReviewKpa` records based on the specified `company_id.`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `Content` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$top` | `{{$top}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/1{{company-id}}/$metadata#EmployeeReviewKpa",
    "value": [
        {
            "EmployeeReviewKpaId": 0,                           // Type="Int64"
            "EmployeeNumber": "String",                         // Type="String"
            "EffectiveDate": "DateTime",                        // Type="DateTime"
            "SectionHeader": "String",                          // Type="String"
            "Weighting": 0.00,                                  // Type="Decimal"
            "KpaManagerStatusId": "Pending",                    // Type="Enum" See: ItemStatus
            "KpaEmployeeStatusId": "Pending",                   // Type="Enum" See: ItemStatus
            "EmployeeReviewKpaDetails": [
                {
                    "EmployeeReviewKpaDetailId": 0,             // Type="Int64"
                    "KpaTitle": "String",                       // Type="String"
                    "KpaDescription": "String",                 // Type="String"
                    "Weighting": 0,                             // Type="Int64"
                    "KpaMeansOfVerification": "String"          // Type="String"
                }
            ]
        }
    ]
}
```

---
## Employee Review Template

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeReviewTemplate?$top={{$top}}&$skip={{$skip}}&$orderby={{$employee-review-template-field}}`

Retrieves a list of active `EmployeeReviewTemplate` records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$orderby` | `{{$employee-review-template-field}}` | Optional (string) - Specifies the order in which items are returned |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeReviewTemplate",
    "value": [
        {
            "EmployeeReviewTemplateId": 0,                      // Type="Int64"
            "EmployeeNumber": "String",                         // Type="String"
            "EffectiveDate": "DateTime",                        // Type="DateTime"
            "TemplateTitle": "String",                          // Type="String"
            "ProcessTitle": "String",                           // Type="String"
            "Status": "In Progress",                            // Type="Enum" see: ReviewTemplate
            "OverallScore": 0                                   // Type="Decimal"
        }
    ]
}
```

---

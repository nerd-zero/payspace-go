# Company > Regions

## Region

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegion?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of `CompanyRegion`records based on the specified `company_id.`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "RegionId eq '1001'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegion",
    "value": [
        {
            "RegionId": 0,                                                      // Type="Int64"
            "RegionDescription": "string",                                      // Type="String"
            "RegionCode": "string",                                             // Type="String"
            "InactiveDate": "2000-00-00T00:00:00+02:00"                         // Type="DateTime"
        }
    ]
}
```

---

### Get a Single Record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegion({{RegionId}})`

Retrieves a single record of`CompanyRegion`information based on the specified `RegionId`.

**Example Response: Get a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegion",
    "HistoryEntityDetailId": 0,                                         // Type="Int64"
    "RegionDescription": "string",                                      // Type="String"
    "RegionCode": "string",                                             // Type="String"
    "InactiveDate": "2000-00-00T00:00:00+02:00"                         // Type="DateTime"
}
```

---

### Create a Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegion`

Creates a single `CompanyRegion` record for the specified `company_id`.

**Request Body:**

```json
{
  //see "CompanyRegion" in metadata endpoint for available fields
}
```

**Example Response: Create a Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegion",
    "HistoryEntityDetailId": 0,                                         // Type="Int64"
    "RegionDescription": "string",                                      // Type="String"
    "RegionCode": "string",                                             // Type="String"
    "InactiveDate": "2000-00-00T00:00:00+02:00"                         // Type="DateTime"
}
```

---

### Update Single Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegion({{RegionId}})`

Update a single `CompanyRegion`record for the specified `RegionId`.

**Request Body:**

```json
{
  //see "CompanyRegion" in metadata endpoint for available fields
}
```

**Example Response: Update Single Record** (Status: 204 No Content)

---

### Delete a Single Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegion({{RegionId}})`

Delete a single `CompanyRegion` record for the specified `RegionId`.

**Example Response: Delete a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegion",
    "Success": true             // Type="boolean"
}
```

---
## Region History

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegionHistory?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of CompanyRegionHistory records based on the specified `company_id.`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "HistoryEntityDetailIdeq '1001'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegionHistory",
    "value": [
        {
            "HistoryEntityDetailId": 0,                                         // Type="Int64"
            "RegionDescription": "string",                                      // Type="String"
            "RegionCode": "string",                                             // Type="String"
            "InactiveDate": "2000-00-00T00:00:00+02:00",                        // Type="DateTime"
            "EffectiveDate": "2000-00-00T00:00:00+02:00",                       // Type="DateTime"
            "CustomFields": [                                                   // CustomField
                // List of custom fields relevent to company region history
            ]
        }
    ]
}
```

---

### Get a collection of all company region history records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegionHistory/all?$orderby={{company-region-history-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{company-region-history-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all company region history records**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegionHistory",
    "value": [
        {
            "SuspensionId": 1,                                                  // Type="Int64"
            "HistoryEntityDetailId": 1,                                         // Type="Int64"
            "RegionDescription": "string",                                      // Type="String"
            "RegionCode": "string",                                             // Type="String"
            "InactiveDate": "2019-01-01T00:00:00+02:00",                        // Type="DateTime"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",                       // Type="DateTime"
            "CustomFields": [
                {
                    "Code": "string",                                           // Type="String"
                    "Label": "string",                                          // Type="String"
                    "Value": "string",                                          // Type="String"
                    "OptionCode": "string"                                      // Type="String"
                }
            ]
        }
    ]
}
```

---

### Get a Single Record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegionHistory({{HistoryEntityDetailId}})`

Retrieves a single record of`CompanyRegionHistory`information based on the specified `HistoryEntityDetailId`.

**Example Response: Get a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegionHistory",
    "HistoryEntityDetailId": 0,                                         // Type="Int64"
    "RegionDescription": "string",                                      // Type="String"
    "RegionCode": "string",                                             // Type="String"
    "InactiveDate": "2000-00-00T00:00:00+02:00",                        // Type="DateTime"
    "EffectiveDate": "2000-00-00T00:00:00+02:00",                       // Type="DateTime"
    "CustomFields": [                                                   // CustomField
        // List of custom fields relevent to company region history
    ]
}
```

---

### Create a Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegionHistory`

Creates a single `CompanyRegionHistory` record for the specified `company_id`.

**Request Body:**

```json
{
  //see "CompanyRegionHistory" in metadata endpoint for available fields
}
```

**Example Response: Create a Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegionHistory",
    "HistoryEntityDetailId": 0,                                         // Type="Int64"
    "RegionDescription": "string",                                      // Type="String"
    "RegionCode": "string",                                             // Type="String"
    "InactiveDate": "2000-00-00T00:00:00+02:00",                        // Type="DateTime"
    "EffectiveDate": "2000-00-00T00:00:00+02:00",                       // Type="DateTime"
    "CustomFields": [                                                   // CustomField
        // List of custom fields relevent to company region history
    ]
}
```

---

### Update Single Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegionHistory({{HistoryEntityDetailId}})`

Update a single `CompanyRegionHistory`record for the specified `HistoryEntityDetailId`.

**Request Body:**

```json
{
  //see "CompanyRegionHistory" in metadata endpoint for available fields
}
```

**Example Response: Update Single Record** (Status: 204 No Content)

---

### Delete a Single Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyRegionHistory({{HistoryEntityDetailId}})`

Delete a single `CompanyRegionHistory` record for the specified `HistoryEntityDetailId`.

**Example Response: Delete a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegionHistory",
    "Success": true             // Type="boolean"
}
```

---

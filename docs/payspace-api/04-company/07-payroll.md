# Company > Payroll

## Payroll Components

Retrieve list of `ComponentCompany` records.

> Note. This endpoint only works from v2.0 and require a open run.  
Only updating of ComponentCode and Description is supported. More functionality may be added in the future

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/ComponentCompany?frequency={{frequency}}&period={{period}}&$orderby={{$field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. Must be an open Run |
| `$orderby` | `{{$field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "GroupId eq 123". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ComponentCompany",
    "@odata.count": 1,
     "value": [
        {
            "CompanyComponentId": 0,     // Type="Int64"	
            "Description": "String",     // Type="String"
            "ComponentCode": "String"    // Type="String"
        }
    ]
}
```

---

### Update Component Company record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/ComponentCompany?frequency={{frequency}}&period={{period}}`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. Must be an open Run |

**Request Body:**

```json
{
  //see "ComponentCompany" in metadata endpoint for available fields
}
```

**Example Response: Get a collection of financial transactions** (Status: 204 No Content)

---

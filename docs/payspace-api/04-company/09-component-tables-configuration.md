# Company > Component Tables Configuration

## Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyTableBuilder?level={{$level}}&categoryCode={{$categoryCode}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&$orderby={{$company-table-builder-field}}&$top={{$top}}`

Retrieves a list of `CompanyTableBuilder` entities based on the specified `level` and `categoryCode`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `level` | `{{$level}}` | Required (string) - Options [Agency / CompanyGroup] |
| `categoryCode` | `{{$categoryCode}}` | Required(string) - Returns items with the applicable Table Builder Category Code |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$orderby` | `{{$company-table-builder-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |

**Example Response: Get a collection**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRegion",
    "value": [
        {
            "TablebuilderId": 0,				            // Type="Int64"
            "Name": "string",			        	        // Type="String"
            "CategoryId": 0,				                // Type="Int64"
            "CategoryCode": "string",			        	// Type="String"
            "TableBuilderCode": "string",			        // Type="String"
            "Level": "string",			        	        // Type="String"
            "CustomFields": [
                {
                    "Code": "string",			        	// Type="String"
                    "Label": "string",			        	// Type="String"
                    "Value": "string",			        	// Type="String"
                    "OptionCode": "string"			        // Type="String"
                }
            ]
        }
    ]
}
```

---

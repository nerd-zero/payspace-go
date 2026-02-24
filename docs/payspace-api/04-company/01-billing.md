# Company > Billing

## Get a collection of financial transactions

**`GET`** `{{api-base-url}}/odata/v2.0/FinancialTrans/:year/:month?$orderby={{$field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&$select={{$select}}`

Retrieves a list of financial transactions. The DocumentTotal field is the sum total of all the document lines, and is calculated by using the following formula: ((Quantity X Amount) + VatAmount) for each document line.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "GroupId eq 123". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$select` | `{{$select}}` | Optional (string) - Returns only the fields specified |

**Example Response: Get a collection of financial transactions** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/$metadata#Document",
    "@odata.count": 1,
    "value": [
        {
            "DocumentId": 0,						// Type="Int64"
            "GroupId": 0,							// Type="Int64"
            "GroupName": "string",					// Type="String"
            "CompanyId": 0,							// Type="Int64"
            "CompanyName": "string",				// Type="String"
            "DocumentType": "string",				// Type="String"
            "Date": "2021-03-08T00:00:00+02:00",	// Type="DateTime"
            "ProFormaCurrency": false,				// Type="Boolean"
            "PoNumber": "string",					// Type="String"
            "DocumentTotal": 0,						// Type="Decimal"
            "DocumentLines": [
                {
                    "DocumentLineId": 0,							// Type="Int64"
                    "DocumentId": 0,								// Type="Int64"
                    "Quantity": 0,									// Type="Decimal"
                    "LineItemDescription": "string",				// Type="String"
                    "Amount": 0,									// Type="Decimal"
                    "VatAmount": 0,									// Type="Decimal"
                    "LineItemClassificationCode": "string",			// Type="String"
                    "LineItemClassicicationDescription": "string",	// Type="String"
                    "RelatedCompanyId": 0,							// Type="Int64"
                    "ApplicablePeriodCode": "string"				// Type="String"
                }
            ]
        }
    ]
}
```

---

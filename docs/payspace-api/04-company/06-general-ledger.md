# Company > General Ledger

## Start General Ledger Data Processing

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGeneralLedger/:year/:month?groupingCode={{groupingCode}}&frequency={{frequency}}&period={{run}}&script={{script}}`

This starts processing the data for a general ledger extract with the given parameters. A `GeneralLedgerExtractionProgress` object is returned with an extraction id used to fetch the results.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `groupingCode` | `{{groupingCode}}` | A string to specify the grouping rules of the general ledger extract. Use the `value` field from the GLGrouping Enum(see Metadata) |
| `frequency` | `{{frequency}}` | Optional. Filter data by frequency. |
| `period` | `{{run}}` | Optional. Filter data by run |
| `script` | `{{script}}` | Optional. Execute a script that has been written in the custom script editor. <br/>Must be the same name as the file name in the custom script editor. |

**Example Response: Start General Ledger Data Processing** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",       // Type="Guid"
    "CompanyId": 0,                 // Type="Int64"    
    "GLGroupingCode": "String",     // Type="String"
    "Year": 0,                      // Type="Int32"
    "Month": 0,                     // Type="Int32"
    "Status": "String",             // Type="String"
    "Error": "String",               // Type="String"
    "DownloadUrl": "String"         // Type="String"
}
```

---
## Get General Ledger Extraction Status

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGeneralLedger/:extractionId/status`

Returns a `GeneralLedgerExtractionProgress` object with the current status of the general ledger processing.

**Example Response: Get General Ledger Extraction Status** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",         // Type="Guid"
    "CompanyId": 0,                 // Type="Int64"    
    "GLGroupingCode": "String",     // Type="String"
    "Year": 0,                      // Type="Int32"
    "Month": 0,                     // Type="Int32"
    "Status": "String",              // Type="String"
    "Error": "String",               // Type="String"
    "DownloadUrl": "String"         // Type="String"
}
```

---
## Get General Ledger Extraction Results

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGeneralLedger/:extractionId/result`

Returns a `GeneralLedgerExtraction` object with the results of the general ledger extraction process.

**Example Response: Get General Ledger Extraction Results** (Status: 200 OK)

```json
{
    "Key": "String",                                        // Type="String"
    "GeneralLedgerLines": [
        {
            "CompanyCode": "String",                        // Type="String"
            "CompanyName": "String",                        // Type="String"
            "Frequency": "String",                          // Type="String"
            "FrequencyId": 0,                               // Type="Int64"
            "GeneralLedgerAccountNumber": "String",         // Type="String"
            "PeriodEndDate": "2019-06-30T00:00:00+02:00",   // Type="DateTime"
            "DebitOrCredit": "DR",                          // Type="String"
            "Description": "String",                        // Type="String"
            "RunDescription": "String",                     // Type="String"
            "TaxCode": "String",                            // Type="String"
            "Amount": 1.00000,                              // Type="Decimal"               
            "CurrencyCode": "ZAR",                          // Type="String"
            "EmployeeNumber": "String"                      // Type="String"
        }
    ]
}
```

---

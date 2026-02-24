# Company > EFT Outbox

## Get EFT Files in publish outbox

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEFT/download?frequency={{frequency}}`

Returns a list of all `CompanyEFTFile` objects in waiting status in the EFT publish outbox.

This method should only be invoked once a file has been manually published through the PaySpace application front end i.e. a user has selected a file format and entered all of the necessary parameters on the salary payment file screen based on the file format selected such as payment date etc.

The file status will be set to downloaded once successfully retrieved.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Optional. Filter data by frequency. |

**Example Response: Get EFT Files in publish outbox**

```json
[
    {
        "EFTFileId": 0,                         // Type="Int64",
        "FileName": "String",                   // Type="String"
        "FileContents": "String",               // Type="String"
        "CompanyCode": "String",                // Type="String"
        "Frequency": "String",                  // Type="String"
        "PayDate": "2019-01-01T00:00:00+02:00"  // Type="DateTime"
    }
]
```

---

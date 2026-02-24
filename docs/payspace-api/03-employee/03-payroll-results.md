# Employee > Payroll Results

## Payslips

Retrieve employee payslip information associated with the `company_id`, `year` and `month` - separate payslips will be returned for each run in the requested period.

### Get a collection of payslips

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslip/:year/:month?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&componentCodes={{ComponentCode}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `componentCodes` | `{{ComponentCode}}` | Optional - A comma seperated list of component codes to return<br>eg. BAS,COMM<br>By default all payslip components will be returned |

**Example Response: Get a collection of payslips** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayslip",
    "value": [        
        {
            "PayslipId": 0,										                    // Type="Int64"
            "CompanyRun": "string",							                        // Type="String"			
            "RunDescription": "string",							                    // Type="String"
            "OrderNumber": 0,									                    // Type="Int64"	
            "PeriodStartDate": "2019-01-01T00:00:00+02:00",		                    // Type="DateTime"	
            "PeriodEndDate": "2019-01-31T00:00:00+02:00",		                    // Type="DateTime"
            "PayDate": "2019-01-31T00:00:00+02:00",		                            // Type="DateTime"	
            "PeriodCode": "string",								                    // Type="String"
            "Frequency": "string",								                    // Type="String"
            "Paid": true,				                                            // Type="Boolean"
            "PayslipComments": "string",						                    // Type="String"
            "EmployeeNumber": "string",							                    // Type="String"
            "TotalValue": 0,								                        // Type="Double"
            "NetPay": 0,								                            // Type="Double"
            "PayslipLines": [
                {
                    "AliasDescription": "string",				                    // Type="String"
                    "AlternateComponentName": "string",			                    // Type="String"
                    "ComponentCode": "string",					                    // Type="String"
                    "PayslipAction": "string",					                    // Type="String"
                    "TaxCode": "string",						                    // Type="String"
                    "ValidTaxCode": true,	                                        // Type="Boolean"
                    "Comments": "string",						                    // Type="String"
                    "PayslipLineValue": 0,						                    // Type="Decimal"
                    "Quantity": 0,								                    // Type="Decimal"
                    "YtdAmount": 0,								                    // Type="Decimal"
                    "DoNotShowOnPayslip": true,	                                    // Type="Boolean"
                    "InPackage": true,	                                            // Type="Boolean"
                    "PayslipMessage": "string"                                      // Type="String"
                }
            ],
            "LeaveBalances": [
                {
                    "Description": "string",					                    // Type="String"
                    "Balance": 0,								                    // Type="Double"
                    "LeaveType": "string",						                    // Type="String"
                    "Accrual": 0,								                    // Type="Double"
       
... (truncated)
```

---

### Get a collection of payslips lines

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslipLine/:year/:month?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Get collection of `EmployeePayslipLine` for the specified `year` and `month`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of payslips lines** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayslipLine",
    "value": [
        {
            "PayslipLineId": 0,                 // Type="Int64"	
            "EmployeeNumber": "string",         // Type="String"
            "RunDescription": "string",         // Type="String"
            "PeriodCode": "string",             // Type="String"
            "Comments": "string",               // Type="String"
            "ComponentName": "string",          // Type="String"
            "AliasDescription": "string",       // Type="String"
            "CompanyComponentId": 0,            // Type="Int64"	
            "AlternateComponentName": null,     // Type="String"
            "ComponentCode": "string",          // Type="String"
            "PayslipAction": "string",          // Type="String"
            "TaxCode": "string",                // Type="String"
            "ValidTaxCode": true,			    // Type="Boolean"
            "PayslipLineValue": 0,              // Type="Decimal"
            "YtdAmount": 0,                     // Type="Decimal"
            "Quantity": 0,                      // Type="Decimal"
            "DoNotShowOnPayslip": true,		    // Type="Boolean"
            "InPackage": true,		            // Type="Boolean"
            "Balance": 0,					    // Type="Double"
            "SpecialComponentCode": "string",   // Type="String"
            "StatutoryComponent": false		    // Type="Boolean"
        }
    ]
}
```

---

### Get a collection of costed payslips lines

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCostedPayslipLine/:year/:month?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Get collection of `EmployeeCostedPayslipLine` for the specified `year` and `month`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of costed payslips lines** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeCostedPayslipLine",
    "value": [
        {
            "CostedPayslipLineId": 0,                 // Type="Int64"	
            "EmployeeNumber": "string",               // Type="String"
            "RunDescription": "string",               // Type="String"
            "PeriodCode": "string",                   // Type="String"
            "Comments": "string",                     // Type="String"
            "ComponentName": "string",                // Type="String"
            "AliasDescription": "string",             // Type="String"
            "CompanyComponentId": 0,                  // Type="Int64"	
            "AlternateComponentName": "string",       // Type="String"
            "ComponentCode": "string",                // Type="String"
            "PayslipAction": "string",                // Type="String"
            "YtdAmount": 0,                           // Type="Decimal"
            "Quantity": 0,			                  // Type="Decimal"
            "CostedProject": "string",                // Type="String"
            "CostedOrgGroup": "string",               // Type="String"
            "OrganizationLevel": "string",            // Type="String"
            "CostedPayslipLineValue": 0		          // Type="Decimal"
        }
    ]
}
```

---

### Get a collection of consolidated payslips

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslip/:year/:month/consolidated?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&componentCodes={{ComponentCode}}`

This endpoint will consolidate an employee's payslips into one if there is more than one payslip in the specified `year` and `month`.

> **Note:**
>
> The consolidated endpoint's entity type is not available in the metadata as it is not a standard odata endpoint.
> Please use the snipit below and the request example as reference. 

~~~
{
	"PeriodCode": "string",						// Type="String"
	"Frequency": "string",						// Type="String"
	"EmployeeNumber": "string", 				  // Type="String"
	"PayslipLines": [
			{
				"AliasDescription": "string",	 // Type="String"
				"ComponentCode": "string",		// Type="String"
				"PayslipAction": "string",		// Type="String"
				"TaxCode": "string",			  // Type="String"
				"PayslipLineValue": 0,			// Type="Decimal"
				"Quantity": 0					 // Type="Decimal"
			}
		],
	"LeaveBalances": [
			{
				"Description": "string",		 // Type="String"
				"Balance": 0,					// Type="Double"
				"LeaveType": "string"			// Type="String"
			}
		]
}
~~~

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `componentCodes` | `{{ComponentCode}}` | Optional - A comma seperated list of component codes to return<br>eg. BAS,COMM<br>By default all payslip components will be returned |

**Example Response: Get a collection of consolidated payslips** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayslip",
    "value": [        
        {
            "PeriodCode": "string",								// Type="String"
            "Frequency": "string",								// Type="String"
            "EmployeeNumber": "string",							// Type="String"
            "PayslipLines": [
                {
                    "AliasDescription": "string",				// Type="String"
                    "AlternateComponentName": "string",			// Type="String"
                    "ComponentCode": "string",					// Type="String"
                    "PayslipAction": "string",					// Type="String"
                    "TaxCode": "string",						// Type="String"
                    "ValidTaxCode": true,	                    // Type="Boolean"
                    "Comments": "string",						// Type="String"
                    "YtdAmount": 0,						        // Type="Decimal"
                    "PayslipLineValue": 0,						// Type="Decimal"
                    "DoNotShowOnPayslip": true,	                // Type="Boolean"
                    "InPackage": true,	                        // Type="Boolean"
                    "PayslipMessage": "string",			        // Type="String"
                    "Quantity": 0								// Type="Decimal"
                }
            ],
            "LeaveBalances": [
                {
                    "Description": "string",					// Type="String"
                    "Balance": 0,								// Type="Double"
                    "LeaveType": "string"						// Type="String"
                }
            ]
        }
    ]
}
```

---

### Get a collection of payslip pdfs

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslip/:year/:month/pdf?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&altLanguage=true`

This endpoint will return a collection of payslips with a `DownloadUrl` for the payslip in PDF format.

> **Note:**
> 
> This endpoint has a custom rate limit of 100 requests per minute.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `altLanguage` | `true` | Optional - Retrieves the payslip components in the alternative language. |

**Example Response: Get a collection of payslip pdfs** (Status: 200 )

```json
{
    "@odata.context": "http://localhost:5000/odata/v2.0/158/$metadata#EmployeePayslipPdf",
    "value": [
        {
            "PayslipId": 0,                                 // Type="Int64"
            "RunDescription": "string",                     // Type="String"
            "OrderNumber": 0,                               // Type="Int64"
            "PeriodStartDate": "2019-01-01T00:00:00+02:00", // Type="DateTime"
            "PeriodEndDate": "2019-01-31T00:00:00+02:00",   // Type="DateTime"
            "PeriodCode": "string",                         // Type="String"
            "Frequency": "string",                          // Type="String"
            "EmployeeNumber": "string",                     // Type="String"
            "DownloadUrl": "string"                         // Type="String"
        }
    ]
}
```

---

### Update payslip comment

**`PATCH`** `{{api-base-url}}/odata/v1.0/{{company-id}}/EmployeePayslip/comment?period={{period}}&frequency={{frequency}}`

Update payslip comment for an employee with the given `EmployeeNumber` on an existing `EmployeePayslip` record for the specified `period` and `frequency`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{period}}` | Use the `value` field from the `CompanyRun` lookup. |
| `frequency` | `{{frequency}}` | Use the `value` field from the `CompanyFrequency` lookup. |

**Request Body:**

```json
{
    "EmployeeNumber": "string",                     // Type="String"
    "PayslipComment" : "string"                     // Type="String"
}
```

**Example Response: Update payslip comment** (Status: 204 No Content)

---

### Get a single payslip pdf

**`GET`** `{{api-base-url}}/odata/v1.0/:company-id/employeepayslip/:payslip-id/download?altLanguage=true&culture={{culture}}`

This endpoint generates a pdf version of an employee's payslip based on the specified `PayslipId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `altLanguage` | `true` | Optional - Retrieves the payslip components in the alternative language. |
| `culture` | `{{culture}}` | Optional - see Culture lookup for list of cultures |

**Example Response: Get a single payslip pdf**

```json
// Returns Employee Payslip in pdf
```

---

### Extract all payslip lines: Start

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslipLine/extract/start/:year/:month`

Starts the process to create a file containing all `EmployeePayslipLine` records

**Example Response: Extract all payslip lines: Start**

```json
{
  "ExtractionId": "f3dc35b0-9f4c-4a42-b579-1b3cb33e8be1",
  "ExtractionType": "EmployeePayslipLine",
  "CompanyId": 158,
  "Frequency": null,
  "Period": null,
  "StartTime": null,
  "EndTime": null,
  "Status": "Queued",
  "Message": null,
  "DownloadUrl": null
}
```

---

### Extract all payslip lines: Status

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslipLine/extract/:extractionId/status`

Returns the extraction process status with download url once completed

**Example Response: Extract all payslip lines: Status**

```json
{
    "ExtractionId": "Guid",               // Type="Guid"
    "ExtractionType": "EmployeePayslipLine", // Type="String"    
    "CompanyId": 0,                       // Type="Int64"    
    "Frequency": "String",                // Type="String"    
    "Period": "String",                   // Type="String"    
    "StartTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",     // Type="Int64"    
    "Status": "String",                   // Type="String"    
    "Message": "String",                  // Type="String"    
    "DownloadUrl": "download url"         // Type="String"    
}
```

---

### Extract all costed payslip lines: Start

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCostedPayslipLine/extract/start/:year/:month`

**Example Response: Extract all costed payslip lines: Start**

```json
{
  "ExtractionId": "1e8bdb47-3aad-4b71-bb23-4a0fd2c9a3a3",
  "ExtractionType": 4,
  "CompanyId": 45178,
  "Frequency": null,
  "Period": "20254",
  "StartTime": null,
  "EndTime": null,
  "Status": 0,
  "Message": null,
  "DownloadUrl": null
}
```

---

### Extract all costed payslip lines: Status

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCostedPayslipLine/extract/:extractionId/status`

**Example Response: Extract all costed payslip lines: Status**

```json
{
    "ExtractionId": "Guid",               // Type="Guid"
    "ExtractionType": "EmployeePayslipLine", // Type="String"    
    "CompanyId": 0,                       // Type="Int64"    
    "Frequency": "String",                // Type="String"    
    "Period": "String",                   // Type="String"    
    "StartTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",     // Type="Int64"    
    "Status": "String",                   // Type="String"    
    "Message": "String",                  // Type="String"    
    "DownloadUrl": "download url"         // Type="String"    
}
```

---
## Edit Payslip

Retrieve or post payslip components for a specific `company-id`, `company-frequency`, and `company-run-period`.

> Note: Duplicate `EditPayslip` records are not allowed. With POST requests the existing record with the same `ComponentCode` and `InputType` will be updated. If there is no existing record then it will be created.

### Get a collection of edit payslip records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EditPayslip?frequency={{frequency}}&period={{period}}&$orderby={{$edit-payslip-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |
| `$orderby` | `{{$edit-payslip-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of edit payslip records**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EditPayslip",
    "value": [
        {
            "EditPayslipId": 0,			    // Type="Int64"	
            "EmployeeNumber": "string",		// Type="String"
            "ComponentCode": "string",		// Type="String"
            "InputType": "string",		    // Type="String"
            "InputValue": 0,			    // Type="Decimal"	
            "CostCentreCode": "string",		// Type="String"
            "ProjectCode": "string",		// Type="String"
            "ActivityCode": "string",		// Type="String"
            "RecoveryAmount": 0,            // Type="Decimal"
            "RecoveryAmountOverride": 0,    // Type="Decimal"
            "Comments": "string"		    // Type="String"
        }
    ]
}
```

---

### Get a single edit payslip record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EditPayslip({{EditPayslipId}})?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Example Response: Get a single edit payslip record**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EditPayslip/$entity",    
    "EditPayslipId": 0,			    // Type="Int64"	
    "EmployeeNumber": "string",		// Type="String"
    "ComponentCode": "string",		// Type="String"
    "InputType": "string",		    // Type="String"
    "InputValue": 0,			    // Type="Decimal"	
    "CostCentreCode": "string",		// Type="String"
    "ProjectCode": "string",		// Type="String"
    "ActivityCode": "string",		// Type="String"
    "RecoveryAmount": 0,            // Type="Decimal"
    "RecoveryAmountOverride": 0,    // Type="Decimal"
    "Comments": "string"		    // Type="String"
}
```

---

### Create a single edit payslip record

**`POST`** `{{api-base-url}}/odata/v2.0/{{company-id}}/EditPayslip?frequency={{frequency}}&period={{period}}&Action={{Action}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |
| `Action` | `{{Action}}` | Optional - Add to the existing value or Overwrite it. eg `Add`, `Overwrite` or `Delete` |

**Request Body:**

```json
{
  //see "EditPayslip" in metadata endpoint for available fields
}
```

**Example Response: Create a single edit payslip record**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EditPayslip/$entity",
    "value": [
        {
            "EditPayslipId": 0,			    // Type="Int64"	
            "EmployeeNumber": "string",		// Type="String"
            "ComponentCode": "string",		// Type="String"
            "InputType": "string",		    // Type="String"
            "InputValue": 0,			    // Type="Decimal"	
            "CostCentreCode": "string",		// Type="String"
            "ProjectCode": "string",		// Type="String"
            "ActivityCode": "string",		// Type="String"
            "RecoveryAmount": 0,            // Type="Decimal"
            "RecoveryAmountOverride": 0,    // Type="Decimal"
            "Comments": "string"		    // Type="String"
        }
    ]
}
```

---

### Update a single edit payslip record

**`PATCH`** `{{api-base-url}}/odata/v2.0/{{company-id}}/EditPayslip({{EditPayslipId}})?frequency={{frequency}}&period={{period}}&Action={{action}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |
| `Action` | `{{action}}` | Optional - Add to the existing value or Overwrite it. eg `Add`, `Overwrite` or `Delete` |

**Request Body:**

```json
{
  //see "EditPayslip" in metadata endpoint for available fields
}
```

**Example Response: Update a single edit payslip record**

---

### Batch updates

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EditPayslip/Batch?frequency={{frequency}}&period={{period}}`

The `EditPayslip` `Batch` endpoint allows you to send an array of multiple POST/PATCH `EditPayslip` requests.

\*   The maximum number of records allowed is `50`.
\*   Each record will be created/updated individually.
\*   The response will contain a list of successfully saved records and records with errors.
\*   You must correct and resubmit the records with errors.
\*   See an example for the response structure.
    

The following fields are required:

| Field Name | Description | Allowed Values |
| --- | --- | --- |
| RequestId | Any text you can use to identify the record in the response back | Any string value |
| Action | Indicate if you want to add to the existing value or overwrite | `Add`, `Overwrite` or `Delete` |
| Method | Indicate if the record should be updated or created | `POST` or `PATCH` |
| Data | The `EditPayslip` record | See `EditPayslip` in metadata |

## Validation

The batch request will 1st validate all the records for the following

\*   Correct request body format
\*   Number of records
\*   Required fields for batch to be successful
    

After the batch validation succeeds each record will be validated individually before saving.

##### Example batch validation response

```
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EditPayslip/$entity",
    "Message": "RequestId 1, 2: Duplicate ComponentCode (CBI001) not allowed for employee SKY1",
    "Details": [
        {
            "Message": "RequestId 1, 2: Duplicate ComponentCode (CBI001) not allowed for employee SKY1"
        }
    ],
    "Success": false
}

```

##### Example batch successful response with successful records and failed records

```
{
    "Items": [
        {
            "Errors": [
                "Invalid employee number 'BCA0010a'."
            ],
            "RequestId": "abc",
            "IsSuccess": false
        },
        {
            "Data": {
                "EditPayslipId": 530256469,
                "EmployeeNumber": "BCA0010",
                "FullName": "Princess Diana",
                "ComponentCode": "Medical Aid Savings Employee",
                "InputType": "",
                "InputValue": 552.00000,
                "CostCentreCode": null,
                "ProjectCode": null,
                "ActivityCode": null,
                "Comments": ""
            },
            "RequestId": "def",
            "IsSuccess": true
        }
    ]
}

```

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Request Body:**

```json
[
  {
    "RequestId": "1",
    "Method": "POST",
    "Action": "Add",
    "Data": {
      "EditPayslipId": 0,
      "EmployeeNumber": "BCA0009",
      "ComponentCode": "ETI Wage Hours",
      "InputType": "",
      "InputValue": 173.36,
      "CostCentreCode": null,
      "ProjectCode": null,
      "ActivityCode": null,
      "Comments": ""
    }
  },
  {
    "RequestId": "2",
    "Method": "POST",
    "Action": "Overwrite",
    "Data": {
      "EditPayslipId": 0,
      "EmployeeNumber": "BCA0009",
      "ComponentCode": "ETI Wage",
      "InputType": "",
      "InputValue": 50000.0,
      "CostCentreCode": null,
      "ProjectCode": null,
      "ActivityCode": null,
      "Comments": ""
    }
  },
  {
    "RequestId": "2",
    "Method": "POST",
    "Action": "Delete",
    "Data": {
      "EditPayslipId": 0,
      "EmployeeNumber": "BCA0009",
      "ComponentCode": "ETI Wage",
      "InputType": "",
      "InputValue": 0.0,
      "CostCentreCode": null,
      "ProjectCode": null,
      "ActivityCode": null,
      "Comments": ""
    }
  }
]
```

**Example Response: Batch Updates**

```json
{
  "Items": [
    {
      "Data": {
        "EditPayslipId": 530256418,
        "EmployeeNumber": "BCA0009",
        "FullName": null,
        "ComponentCode": "ETI Wage Hours",
        "InputType": "",
        "InputValue": 173.36,
        "CostCentreCode": null,
        "ProjectCode": null,
        "ActivityCode": null,
        "RecoveryAmount": 0,
        "RecoveryAmountOverride": 0,
        "Comments": ""
      },
      "RequestId": "1",
      "IsSuccess": true
    },
    {
      "Errors": [
        "Invalid value for field Input Type."
      ],
      "RequestId": "2",
      "IsSuccess": false
    }
  ]
}
```

---

### Delete a single edit payslip record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EditPayslip({{EditPayslipId}})?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Request Body:**

```json
{
  //see "EditPayslip" in metadata endpoint for available fields
}
```

**Example Response: Delete a single edit payslip record**

---

### Extract all Records

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EditPayslip/extract/start?frequency={{frequency}}&period={{period}}`

Starts the process to create a file containing all `EditPayslip` records

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Example Response: Extract all Records**

```json
{
    "ExtractionId": "Guid",               // Type="Guid"
    "ExtractionType": "EditPayslip",      // Type="String"    
    "CompanyId": 0,                       // Type="Int64"    
    "Frequency": "String",                // Type="String"    
    "Period": "String",                   // Type="String"    
    "StartTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",     // Type="Int64"    
    "Status": "String",                   // Type="String"    
    "Message": "String",                  // Type="String"    
    "DownloadUrl": "download url"         // Type="String"    
}
```

---

### Extract all Records: Status

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EditPayslip/extract/:extractionId/status?frequency={{frequency}}&period={{period}}`

Returns the extraction process status with download url once completed

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` |  |
| `period` | `{{period}}` |  |

**Example Response: Extract all Records: Status**

```json
{
    "ExtractionId": "Guid",               // Type="Guid"
    "ExtractionType": "EditPayslip", // Type="String"    
    "CompanyId": 0,                       // Type="Int64"    
    "Frequency": "String",                // Type="String"    
    "Period": "String",                   // Type="String"    
    "StartTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",     // Type="Int64"    
    "Status": "String",                   // Type="String"    
    "Message": "String",                  // Type="String"    
    "DownloadUrl": "download url"         // Type="String"    
}
```

---
## Payslip Take On

Retrieve or post Payslip Take On components for a specific `company-id`, `company-frequency`, and `company-run-period`.

> Note: Duplicate `EmployeePayslipTakeOn` records are not allowed. With POST requests the existing record with the same `ComponentCode` and `InputType` will be updated. If there is no existing record then it will be created.

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslipTakeOn?frequency={{frequency}}&period={{period}}&$orderby={{$fieldname}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |
| `$orderby` | `{{$fieldname}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayslipTakeOn",
    "value": [
        {
            "EmployeePayslipTakeOnId": 0,   // Type="Int64"	
            "EmployeeNumber": "string",	    // Type="String"
            "FullName": "string",		    // Type="String"
            "ComponentCode": "string",	    // Type="String" use lookup TakeOnComponentCode
            "InputType": "string",		    // Type="String" use lookup TakeOnInputType
            "Quantity": 0,			        // Type="Decimal"	
            "Amount": 0		                // Type="Decimal"
        }
    ]
}
```

---

### Get a single record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslipTakeOn({{EditPayslipId}})?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Example Response: Get a single record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayslipTakeOn/$entity",    
    "EmployeePayslipTakeOnId": 0,   // Type="Int64"	
    "EmployeeNumber": "string",	    // Type="String"
    "FullName": "string",		    // Type="String"
    "ComponentCode": "string",	    // Type="String" use lookup TakeOnComponentCode
    "InputType": "string",		    // Type="String" use lookup TakeOnInputType
    "Quantity": 0,			        // Type="Decimal"	
    "Amount": 0		                // Type="Decimal"
}
```

---

### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/{{company-id}}/EmployeePayslipTakeOn?frequency={{frequency}}&period={{period}}&Action={{Action}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |
| `Action` | `{{Action}}` | Optional - Add to the existing value or Overwrite it. eg `Add`, `Overwrite` or `Delete` |

**Request Body:**

```json
{
  //see "EmployeePayslipTakeOn" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayslipTakeOn/$entity",
    "value": [
        {
            "EmployeePayslipTakeOnId": 0,   // Type="Int64"	
            "EmployeeNumber": "string",	    // Type="String"
            "FullName": "string",		    // Type="String"
            "ComponentCode": "string",	    // Type="String"
            "InputType": "string",		    // Type="String"
            "Quantity": 0,			        // Type="Decimal"	
            "Amount": 0		                // Type="Decimal"
        }
    ]
}
```

---

### Update existing record

**`PATCH`** `{{api-base-url}}/odata/v2.0/{{company-id}}/EmployeePayslipTakeOn({{EmployeePayslipTakeOnId}})?frequency={{frequency}}&period={{period}}&Action={{action}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |
| `Action` | `{{action}}` | Optional - Add to the existing value or Overwrite it. eg `Add`, `Overwrite` or `Delete` |

**Request Body:**

```json
{
  //see "EditPayslip" in metadata endpoint for available fields
}
```

**Example Response: Update existing record** (Status: 204 No Content)

---

### Delete existing record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayslipTakeOn({{EmployeePayslipTakeOnId}})?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Example Response: Delete existing record** (Status: 200 OK)

---
## Lump Sum Directive

Retrieve or post lump sum directives for a specific `company-id`, `company-frequency`, and `company-run-period`.

### Get a collection of lump sum directive records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLumpSum?frequency={{frequency}}&period={{period}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of lump sum directive records**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLumpSum",
    "value": [
        {
            "EmployeeLumpSumId": 0,			                                    // Type="Int64"	
            "EmployeeNumber": "string",		                                    // Type="String"
            "FullName": "string",       		                                // Type="String"
            "DirectiveIssuedDate": "2000-00-00T00:00:00+02:00",                 // Type="DateTime"
            "TaxCode": "string",		                                        // Type="String"
            "DirectiveAmount": 0,                                               // Type="Decimal"
            "DirectiveTax": 0,                                                  // Type="Decimal"
            "DirectiveNumber": "string",		                                // Type="String"
            "ReferenceNumber": "string"		                                    // Type="String"
        }
    ]
}
```

---

### Get a single lump sum directive record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLumpSum({{EmployeeLumpSumId}})?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Example Response: Get a single lump sum directive record**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLumpSum/$entity",
    "value": [
        {
            "EmployeeLumpSumId": 0,			                                    // Type="Int64"	
            "EmployeeNumber": "string",		                                    // Type="String"
            "FullName": "string",       		                                // Type="String"
            "DirectiveIssuedDate": "2000-00-00T00:00:00+02:00",                 // Type="DateTime"
            "TaxCode": "string",		                                        // Type="String"
            "DirectiveAmount": 0,                                               // Type="Decimal"
            "DirectiveTax": 0,                                                  // Type="Decimal"
            "DirectiveNumber": "string",		                                // Type="String"
            "ReferenceNumber": "string"		                                    // Type="String"
        }
    ]
}
```

---

### Create a single lump sum directive record

**`POST`** `{{api-base-url}}/odata/v2.0/{{company-id}}/EmployeeLumpSum?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Request Body:**

```json
{
  //see "EmployeeLumpSum" in metadata endpoint for available fields
}
```

**Example Response: Create a single lump sum directive record**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLumpSum/$entity",
    "value": [
        {
            "EmployeeLumpSumId": 0,			                                    // Type="Int64"	
            "EmployeeNumber": "string",		                                    // Type="String"
            "FullName": "string",       		                                // Type="String"
            "DirectiveIssuedDate": "2000-00-00T00:00:00+02:00",                 // Type="DateTime"
            "TaxCode": "string",		                                        // Type="String"
            "DirectiveAmount": 0,                                               // Type="Decimal"
            "DirectiveTax": 0,                                                  // Type="Decimal"
            "DirectiveNumber": "string",		                                // Type="String"
            "ReferenceNumber": "string"		                                    // Type="String"
        }
    ]
}
```

---

### Update a single lump sum directive record

**`PATCH`** `{{api-base-url}}/odata/v2.0/{{company-id}}/EmployeeLumpSum({{EmployeeLumpSumId}})?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Request Body:**

```json
{
  //see "EmployeeLumpSum" in metadata endpoint for available fields
}
```

**Example Response: Update a single lump sum directive record**

---

### Delete a single lump sum directive record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLumpSum({{EmployeeLumpSumId}})?frequency={{frequency}}&period={{period}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Use the `value` field from the CompanyFrequency lookup. |
| `period` | `{{period}}` | Use the `value` field from the CompanyRun lookup. |

**Request Body:**

```json
{
  //see "EmployeeLumpSum" in metadata endpoint for available fields
}
```

**Example Response: Delete a single lump sum directive record**

---
## Tax Certificates

### Get a collection of tax certificates

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTaxCertificate?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of tax certificates** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTaxCertificate",
    "value": [        
        {
            "TaxCertificateId": "string",						// Type="String"			
            "EmployeeNumber": "string",							// Type="String"	
            "YearStartDate": "2019-02-01T00:00:00+02:00",		// Type="DateTime"	
            "YearEndDate": "2020-01-31T00:00:00+02:00",			// Type="DateTime"
            "Irp5Number": "string",								// Type="String"
            "ManualNumber": "string",							// Type="String"
            "DirectiveNumber": "string",						// Type="String"
            "Irp5SourceCode": "string",	                        // Type="String"
            "NewIrp5Numbers": "string",					        // Type="String"
            "DownloadUrl": "string"							    // Type="String"
        }
    ]
}
```

---

# Company > Configuration

Contains several endpoints related to Company information.

## Organization Position Details

Retrieve organization position information for all positions associated with the `company_id`.

### Get a collection of Organization Positions

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationPositionDetail?$orderby={{$organization-position-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$organization-position-field}}` |  |
| `$top` | `{{$top}}` |  |
| `$skip` | `{{$skip}}` |  |
| `$count` | `{{$count}}` |  |
| `$filter` | `{{$filter}}` |  |

**Example Response: Get a collection of Organization Positions** (Status: 200 OK)

```json
{
   "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationPositionDetail",
    "@odata.count": 2,
    "value": [
        {
            "PositionDetailId": 0,                                // Type="Int64"
            "PositionTitle": "string",                            // Type="String"
            "OrganizationPositionId": 0,                          // Type="Int64" - ReadOnly field
            "EffectiveDate": "2020-08-21T00:00:00+02:00",         // Type="DateTime"
            "Grade": "string",                                    // Type="String"
            "PositionPurpose": "string",                          // Type="String"
            "SpecialFeatures": "string",                          // Type="String"
            "PositionType": "string",                             // Type="String"
            "SSTLevel": "string",                                 // Type="String"
            "PositionCode": "string",                             // Type="String"
            "DirectlyReportsPositionOverride": "string",          // Type="String"
            "CompanyTemplateDefinition": "string",                // Type="String"
            "InactiveDate": "string",                             // Type="DateTime"
            "OccupationalLevel": "string",                        // Type="String"
            "EeFunction": "string",                               // Type="String"
            "BeeLevel": "string",                                 // Type="String"
            "OfoCode": "string",                                  // Type="String"
            "CompanyPositionFamily": "string",                    // Type="String"
            "FunctionArea": "string",                             // Type="String"
            "SalarySurveyCode1": "string",                        // Type="String"
            "SalarySurveyCode2": "string"                         // Type="String"
        }
    ]
}
```

---

### Get a single Organization Position record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationPositionDetail({{OrganizationPositionDetailId}})`

**Example Response: Get a single Organization Position record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationPositionDetail/$entity",	
        "PositionDetailId": 0,                                // Type="Int64"
        "PositionTitle": "string",                            // Type="String"
        "EffectiveDate": "2020-08-21T00:00:00+02:00",         // Type="DateTime"
        "Grade": "string",                                    // Type="String"
        "PositionPurpose": "string",                          // Type="String"
        "SpecialFeatures": "string",                          // Type="String"
        "PositionType": "string",                             // Type="String"
        "SSTLevel": "string",                                 // Type="String"
        "PositionCode": "string",                             // Type="String"
        "DirectlyReportsPositionOverride": "string",          // Type="String"
        "CompanyTemplateDefinition": "string",                // Type="String"
        "InactiveDate": "string",                             // Type="DateTime"
        "OccupationalLevel": "string",                        // Type="String"
        "EeFunction": "string",                               // Type="String"
        "BeeLevel": "string",                                 // Type="String"
        "OfoCode": "string",                                  // Type="String"
        "CompanyPositionFamily": "string",                    // Type="String"
        "FunctionArea": "string",                             // Type="String"
        "SalarySurveyCode1": "string",                        // Type="String"
        "SalarySurveyCode2": "string"                         // Type="String"
}
```

---

### Create a single organization position record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationPositionDetail`

Create a single organization position record for the specified `OrganizationPositionDetailId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "OrganizationPositionDetail" in metadata endpoint for available fields
}
```

**Example Response: Create a single organization positoin record** (Status: 201 Created)

```json
{
    "@odata.context": "https://apiuat.payspace.com/odata/v2.0/158/$metadata#OrganizationPositionDetail/$entity",
    "PositionDetailId": 0,                                  // Type="Int64"
    "PositionTitle": "Test2",                               // Type="String"                               
    "EffectiveDate": "2021-01-01T00:00:00+02:00",           // Type="DateTime"
    "Grade": "01",                                          // Type="String"
    "PositionPurpose": null,                                // Type="String"  
    "SpecialFeatures": null,                                // Type="String"  
    "PositionType": null,                                   // Type="String"
    "SSTLevel": null,                                       // Type="String"
    "PositionCode": null,                                   // Type="String"  
    "DirectlyReportsPositionOverride": null,                // Type="String"
    "CompanyTemplateDefinition": null,                      // Type="String"
    "InactiveDate": null,                                   // Type="DateTime"
    "OccupationalLevel": null,                              // Type="String"
    "EeFunction": null,                                     // Type="Enum"
    "BeeLevel": null,                                       // Type="Enum"
    "OfoCode": null,                                        // Type="String"
    "CompanyPositionFamily": null,                          // Type="String"
    "FunctionArea": null,                                   // Type="String"
    "SalarySurveyCode1": "null",                            // Type="String"
    "SalarySurveyCode2": "null"                             // Type="String"
}
```

---

### Update a single organization position record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationPositionDetail({{OrganizationPositionDetailId}})`

Update a single organization position record based on the specified `OrganizationPositionDetailId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "OrganizationPositionDetail" in metadata endpoint for available fields
}
```

**Example Response: Update a single organization position record** (Status: 204 No Content)

```json
{
    {
    "@odata.context": "https://apiuat.payspace.com/odata/v2.0/158/$metadata#OrganizationPositionDetail/$entity",
    "PositionDetailId": 0,                                  // Type="Int64"
    "PositionTitle": "Test2",                               // Type="String"                               
    "EffectiveDate": "2021-01-01T00:00:00+02:00",           // Type="DateTime"
    "Grade": "01",                                          // Type="String"
    "PositionPurpose": null,                                // Type="String"  
    "SpecialFeatures": null,                                // Type="String"  
    "PositionType": null,                                   // Type="String"
    "SSTLevel": null,                                       // Type="String"
    "PositionCode": null,                                   // Type="String"  
    "DirectlyReportsPositionOverride": null,                // Type="String"
    "CompanyTemplateDefinition": null,                      // Type="String"
    "InactiveDate": null,                                   // Type="DateTime"
    "OccupationalLevel": null,                              // Type="String"
    "EeFunction": null,                                     // Type="Enum"
    "BeeLevel": null,                                       // Type="Enum"
    "OfoCode": null,                                        // Type="String"
    "CompanyPositionFamily": null,                          // Type="String"
    "FunctionArea": null,                                   // Type="String"
    "SalarySurveyCode1": "null",                            // Type="String"
    "SalarySurveyCode2": "null"                             // Type="String"
    }
}
```

---

### Delete a single organization position record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationPositionDetail({{OrganizationPositionDetailId}})`

Deletes organization position information based on the specified `OrganizationPositionDetailId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single organization position record** (Status: 200 OK)

---
## Organization Hierarchy Units

Retrieve organization unit information for all units associated with the `company_id`.

### Get a collection of organization units

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationUnit?$orderby={{$organization-unit-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of organization units based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$organization-unit-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of organization units** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationUnit",
    "@odata.count": 0,
    "value": [
		{
            "OrganizationUnitId": 0,						// Type="Int64"
            "ParentOrganizationUnitId": 0,					// Type="Int64"
            "UploadCode": "string",							// Type="String"
            "Description": "string",						// Type="String"
            "CostCentre": false,							// Type="Boolean"
            "OrganizationLevel": "string",					// Type="String"
            "GroupGlKey": "string",							// Type="String"
            "Budget": 0.00000,								// Type="Decimal"
            "Reference": "string",							// Type="String"
            "ManagerEmployeeNumber": "string",				// Type="String"
            "InactiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
            "Address": {
		        "AddressId": 0,								// Type="Int64"
		        "AddressLine1": "string",					// Type="String"
		        "AddressLine2": "string",					// Type="String"
		        "AddressLine3": "string",					// Type="String"
		        "AddressCode": "string",					// Type="String"
		        "Province": "string",						// Type="String"
		        "AddressCountry": "string",					// Type="String"
		        "UnitNumber": "string",						// Type="String"
		        "Complex": "string",						// Type="String"
		        "StreetNumber": "string"					// Type="String"
		    }
        }
    ]
}
```

---

### Get a single organization unit record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationUnit({{OrganizationUnitId}})`

Retrieves organization unit record based on the specified `OrganizationUnitId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single organization unit** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationUnit/$entity",
    "OrganizationUnitId": 0,						// Type="Int64"
    "ParentOrganizationUnitId": 0,					// Type="Int64"
    "UploadCode": "string",							// Type="String"
    "Description": "string",						// Type="String"
    "CostCentre": false,							// Type="Boolean"
    "OrganizationLevel": "string",					// Type="String"
    "GroupGlKey": "string",							// Type="String"
    "Budget": 0.00000,								// Type="Decimal"
    "Reference": "string",							// Type="String"
    "ManagerEmployeeNumber": "string",				// Type="String"
    "InactiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
     "Address": {
        "AddressId": 0,								// Type="Int64"
        "AddressLine1": "string",					// Type="String"
        "AddressLine2": "string",					// Type="String"
        "AddressLine3": "string",					// Type="String"
        "AddressCode": "string",					// Type="String"
        "Province": "string",						// Type="String"
        "AddressCountry": "string",					// Type="String"
        "UnitNumber": "string",						// Type="String"
        "Complex": "string",						// Type="String"
        "StreetNumber": "string"					// Type="String"
    }
}
```

---

### Create a single organization unit record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationUnit`

Create a single organization unit record for the specified `company-id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "OrganizationUnit" in metadata endpoint for available fields
}
```

**Example Response: Create a single organization unit record** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationUnit/$entity",
    "OrganizationUnitId": 0,						// Type="Int64"
    "ParentOrganizationUnitId": 0,					// Type="Int64"
    "UploadCode": "string",							// Type="String"
    "Description": "string",						// Type="String"
    "CostCentre": false,							// Type="Boolean"
    "OrganizationLevel": "string",					// Type="String"
    "GroupGlKey": "string",							// Type="String"
    "Budget": 0.00000,								// Type="Decimal"
    "Reference": "string",							// Type="String"
    "ManagerEmployeeNumber": "string",				// Type="String"
    "InactiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
     "Address": {
        "AddressId": 0,								// Type="Int64"
        "AddressLine1": "string",					// Type="String"
        "AddressLine2": "string",					// Type="String"
        "AddressLine3": "string",					// Type="String"
        "AddressCode": "string",					// Type="String"
        "Province": "string",						// Type="String"
        "AddressCountry": "string",					// Type="String"
        "UnitNumber": "string",						// Type="String"
        "Complex": "string",						// Type="String"
        "StreetNumber": "string"					// Type="String"
    }
}
```

---

### Update a single organization unit record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationUnit({{OrganizationUnitId}})`

Update a single organization unit record based on the specified `OrganizationUnitId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "OrganizationUnit" in metadata endpoint for available fields
}
```

**Example Response: Update a single organization unit record** (Status: 204 No Content)

---

### Delete a organization unit project record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationUnit({{OrganizationUnitId}})`

Delete organization unit record based on the specified `OrganizationUnitId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single organization unit record** (Status: 200 OK)

```json
{}
```

---
## Job Management

### Get a collection of Job Management records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/JobManagement?$orderby={{$job-management-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of job management records based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$job-management-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of job management records** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#JobManagement",
    "@odata.count": 0,
    "value": [
        {
            "JobManagementId": 0,									// Type="Int64"
            "Status": "Active",										// Type="Enum" see: JobManagementStatus in metadata
            "IsPlannedJob": false,									// Type="Boolean"
            "JobNumber": "string",									// Type="String"
            "ExternalParty": false,									// Type="Boolean"
            "CompanyFrequency": "string",							// Type="String"
            "OrganizationPosition": "string",						// Type="String"
            "OrganizationPositionWithCode": "string",               // Type="String"
            "OrganizationPositionId": 0,							// Type="Int64"
            "OrganizationGroup": "string",							// Type="String"
            "OrganizationGroupId": 0,								// Type="Int64" aka OrganizationUnitId in organization unit endpoint
            "OrganizationRegion": "string",							// Type="String"
            "JobTitle": "string",									// Type="String"
            "Grade": "string",										// Type="String"
            "BudgetGroup": "string",								// Type="String"
            "EmploymentSubCategory": "string",						// Type="String"
            "EmployeeFirstName": "string",							// Type="String"
            "EmployeeLastName": "string",							// Type="String"
            "EmployeeNumber": "string",                             // Type="String"
            "EmployementStartDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
            "EmployementEndDate": "2019-01-01T00:00:00+02:00",	    // Type="DateTime"
            "OverrideEmployeeName": false,							// Type="Boolean"
            "PositionType": "string",								// Type="String"
            "GeneralLedger": "string",								// Type="String"
            "ReportsToEffectiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
            "ReportsToJob": "string",								// Type="String"
            "ReportsToJobManagementId": 0,							// Type="Int64"
            "PreferredGender": "Male",							    // Type="Enum" see: Gender in metadata
            "PreferredRace": "string",								// Type="String"
            "PreferredLocalization": "string",						// Type="String"
            "OverrideRoster": "string",							    // Type="String"
            "OverrideOfoCode": "string",							// Type="String"
            "OverrideOccupationalLevel": "string",					// Type="String"
            "AbolishDate": "2019-01-01T00:00:00+02:00",	            // Type="DateTime"
            "UploadCode": "string",                                 // Type="String"
            "CustomFields": [
                {
                    "Code": "string",								// Type="String"
                    "Label": "string",								// Type="String"
                    "Value": "string"								// Type="String"
                }
            ]
        }
    ]
}
```

---

### Get a single job management record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/JobManagement({{JobManagementId}})`

Retrieves job management record based on the specified `OrganizationUnitId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single job management** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#JobManagement/$entity",
    "JobManagementId": 0,									// Type="Int64"
    "Status": "Active",										// Type="Enum" see: JobManagementStatus in metadata
    "IsPlannedJob": false,									// Type="Boolean"
    "JobNumber": "string",									// Type="String"
    "ExternalParty": false,									// Type="Boolean"
    "CompanyFrequency": "string",							// Type="String"
    "OrganizationPosition": "string",						// Type="String"
    "OrganizationPositionWithCode": "string",               // Type="String"
    "OrganizationPositionId": 0,							// Type="Int64"
    "OrganizationGroup": "string",							// Type="String"
    "OrganizationGroupId": 0,								// Type="Int64" aka OrganizationUnitId in organization unit endpoint
    "OrganizationRegion": "string",							// Type="String"
    "JobTitle": "string",									// Type="String"
    "Grade": "string",										// Type="String"
    "BudgetGroup": "string",								// Type="String"
    "EmploymentSubCategory": "string",						// Type="String"
    "EmployeeFirstName": "string",							// Type="String"
    "EmployeeLastName": "string",							// Type="String"
    "EmployeeNumber": "string",                             // Type="String"
    "EmployementStartDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
    "EmployementEndDate": "2019-01-01T00:00:00+02:00",	    // Type="DateTime"
    "OverrideEmployeeName": false,							// Type="Boolean"
    "PositionType": "string",								// Type="String"
    "GeneralLedger": "string",								// Type="String"
    "ReportsToEffectiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
    "ReportsToJob": "string",								// Type="String"
    "ReportsToJobManagementId": 0,							// Type="Int64"
    "PreferredGender": "Male",							    // Type="Enum" see: Gender in metadata
    "PreferredRace": "string",								// Type="String"
    "PreferredLocalization": "string",						// Type="String"
    "OverrideRoster": "string",							    // Type="String"
    "OverrideOfoCode": "string",							// Type="String"
    "OverrideOccupationalLevel": "string",					// Type="String"
    "AbolishDate": "2019-01-01T00:00:00+02:00",	            // Type="DateTime"
    "UploadCode": "string",                                 // Type="String"
    "CustomFields": [
        {
            "Code": "string",								// Type="String"
            "Label": "string",								// Type="String"
            "Value": "string"								// Type="String"
        }
    ]
}
```

---

### Get the history for a single job

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/JobManagement/{{JobManagementId}}/history`

Retrieves a list of job management history records based on the specified JobManagementId.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get the history for a single job** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#JobManagementHistory",
    "value": [
        {
            "EmployeePositionId": 0, // Type="Int64"
            "EmployeeNumber": "string", // Type="String"
            "FullName": "string", // Type="String"
            "StartDate": "2021-01-01T00:00:00+02:00", // Type="DateTime"
            "EndDate": "2023-12-31T00:00:00+02:00" // Type="DateTime"
        }
    ]
}
```

---
## Job Management Budget Cost

### Get a collection of Job Management Budget Cost records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/JobManagementBudgetCost/:job-management-id/:company-budget-period-id?$orderby={{$job-management-budget-cost-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

StartFragment

Retrieves a list of job management budget cost records based on the specified `CompanyBudgetPeriod` for a specific `JobManagementId`.

EndFragment..

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$job-management-budget-cost-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of Job Management Budget Cost records** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#JobManagementBudgetCost",
    "@odata.count": 0,
    "value": [
        {
            "CompanyComponentId": 0,                                // Type="Int64"
            "JobNumber": "string",                                  // Type="String"
            "ActionType": "string",                                 // Type="String"
            "EmployeeNumber": "string",                             // Type="String"
            "CompanyComponent": "string",                           // Type="String"
            "CostAmount": 0                                         // Type="Decimal"
        }
    ]
}
```

---
## Budget Archive Report

### Start snapshot period based Budget Archive Employee data processing.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveEmployee/:company-budget-period/snap-shot/extract`

This starts processing the data for a budget archive employee extract based on the specified snapshot `CompanyBudgetPeriod` for a specific `company-id`.

A `BudgetArchiveExtractionProgress` object is returned with an extraction id used to fetch the status. This object also contains a `DownloadUrl` which can be used to retrieve the result once the `Status` is completed.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Start snapshot period based Budget Archive Employee data processing.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---

### Get Budget Archive Employee Extraction Status.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveEmployee/extract/:extraction-id/status`

This retrieves a `BudgetArchiveExtractionProgress` object with a `DownloadUrl` which can be used to retrieve the result once the `Status` is `Completed`.

Example of extraction result
```
[
	{
		"JobManagementId": 0,                                       // Type="Int64"
		"JobNumber": "string",                                      // Type="String"
		"JobTitle": "string",                                       // Type="String"
		"OrganizationGroup": "string",                              // Type="String"
		"OrganizationRegion": "string",                             // Type="String"
		"Grade": "string",                                          // Type="String"
		"EmploymentSubCategory": "string",                          // Type="String"
		"EmployeeFirstName": "string",                              // Type="String"
		"EmployeeLastName": "string",                               // Type="String"
		"EmployeeNumber": "string",                                 // Type="String"
		"PositionType": "string",                                   // Type="String"
		"Comments": "string",                                       // Type="String"
		"OverrideEmployeeName": "string",                           // Type="String"
		"OverrideFigures": "string",                                // Type="String"
		"AbolishDate": "string",                                    // Type="String"
		"PlannedJobTitle": "string",                                // Type="String"
		"PlannedGrade": "string",                                   // Type="String"
		"FilledCount": 0,                                           // Type="Int32"
		"DoNotApplyIncrease": false,                                // Type="Boolean"
		"Figures": [                                                // Type="Collection"
			{
				"PeriodCode": "string",                     // Type="String"
				"Amount": 0                                 // Type="Decimal"
			},
			{
				"PeriodCode": "string",                     // Type="Decimal"
				"Amount": 0                                 // Type="Decimal"
			},
			{
				"PeriodCode": "string",                     // Type="String"
				"Amount": 0                                 // Type="Decimal"
			}
		]
	}
]

```

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get Budget Archive Employee Extraction Status.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---

### Start archive period based Budget Archive Employee data processing.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveEmployee/:company-budget-snap-shot-archive/archive/extract`

This starts processing the data for a budget archive employee extract based on the specified `CompanyBudgetSnapshotArchive`for a specific `company-id`.

A `BudgetArchiveExtractionProgress` object is returned with an extraction id used to fetch the status. This object also contains a `DownloadUrl` which can be used to retrieve the result once the `Status` is completed.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Start archive period based Budget Archive Employee data processing.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---

### Start date range based Budget Archive Employee data processing.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveEmployee/actuals/extract?startDate={{startDate}}&endDate={{endDate}}`

This starts processing the data for a budget archive employee extract based on the specified `Start date` and `End date` for a specific `company-id`.

A `BudgetArchiveExtractionProgress` object is returned with an extraction id used to fetch the status. This object also contains a `DownloadUrl` which can be used to retrieve the result once the `Status` is completed.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `startDate` | `{{startDate}}` | Required (Date) - Specifies start date from which figures are returned e.g. 2023-01-01 |
| `endDate` | `{{endDate}}` | Required (Date) - Specifies end date from which figures are returned e.g. 2023-01-01 |

**Example Response: Start date range based Budget Archive Employee data processing.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---

### Start snapshot period based Budget Archive Organization data processing.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveOrganization/:company-budget-period-id/snap-shot/extract`

Retrieves a list of budget archive organization records based on the specified snapshot `CompanyBudgetPeriod` for a specific `company-id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Start snapshot period based Budget Archive Organization data processing.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---

### Get Budget Archive Organization Extraction Status.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveOrganization/extract/:extraction-id/status`

This retrieves a `BudgetArchiveExtractionProgress` object with a `DownloadUrl` which can be used to retrieve the result once the `Status` is `Completed`.

Example of extraction result

```
[
	{
		"OrganizationGroupId": 0,                                   // Type="Int64"
		"OrganizationGroup": "string",                              // Type="String"
		"FilledCount": 0,                                           // Type="Int32"
		"JobCount": 0,                                              // Type="Int32"
		"Figures": [                                                // Type="Collection"
			{
				"PeriodCode": "string",                             // Type="String"
				"Amount": 0                                         // Type="Decimal"
			},
			{
				"PeriodCode": "string",                             // Type="Decimal"
				"Amount": 0                                         // Type="Decimal"
			},
			{
				"PeriodCode": "string",                             // Type="String"
				"Amount": 0                                         // Type="Decimal"
			}
		]
	}
]

```

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get Budget Archive Organization Extraction Status.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---

### Start archive period based Budget Archive Organization data processing.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveOrganization/:company-budget-snap-shot-archive-id/archive/extract`

This starts processing the data for a budget archive organization extract based on the specified `CompanyBudgetSnapshotArchive`for a specific `company-id`.

A `BudgetArchiveExtractionProgress` object is returned with an extraction id used to fetch the status. This object also contains a `DownloadUrl` which can be used to retrieve the result once the `Status` is completed.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Start archive period based Budget Archive Organization data processing.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---

### Start date range based Budget Archive Organization data processing.

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/BudgetArchiveOrganization/actuals/extract?startDate={{startDate}}&endDate={{endDate}}`

This starts processing the data for a budget archive organization extract based on the specified `Start date` and `End date` for a specific `company-id`.

A `BudgetArchiveExtractionProgress` object is returned with an extraction id used to fetch the status. This object also contains a `DownloadUrl` which can be used to retrieve the result once the `Status` is completed.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `startDate` | `{{startDate}}` | Required (Date) - Specifies start date from which figures are returned |
| `endDate` | `{{endDate}}` | Required (Date) - Specifies end date from which figures are returned |

**Example Response: Start date range based Budget Archive Organization data processing.** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",                             // Type="Guid"
    "CompanyId": 0,                                     // Type="Int64"
    "BudgetPeriodId": 0,                                // Type="Int64"
    "CompanyBudgetSnapshotsArchiveId": 0,               // Type="Int64"  
    "StartDate": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndDate": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "StartTime": "2023-01-01T00:00:00+02:00",           // Type="DateTime"
    "EndTime": "2023-01-01T00:00:00+02:00",             // Type="DateTime"
    "Status": "String",                                 // Type="String"
    "Message": "String",                                // Type="String"
    "DownloadUrl": "String"                             // Type="String"
}
```

---
## Currency Exchange Rates

### Get collection of Currency Exchange Rate records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupExchangeRate`

Retrieves a list of currency exchange rate records based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get collection of Currency Exchange Rate records** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupExchangeRate",
    "@odata.count": 0,
    "value": [
        {
            "ExchangeRateId": 0,                                    // Type="Int64"
            "Currency": "string",                                   // Type="String"
            "Country": "string",                                    // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "OneDollarInCurrency": 0                                // Type="Decimal"
        }
    ]
}
```

---

### Get collection on Currency Exchange Rate Records by Effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupExchangeRate/effective/:effectivedate?$orderby={{$EffectiveDate}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$EffectiveDate}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get collection on Currency Exchange Rate Records by Effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupExchangeRate",
    "@odata.count": 0,
    "value": [
        {
            "ExchangeRateId": 0,                                    // Type="Int64"
            "Currency": "string",                                   // Type="String"
            "Country": "string",                                    // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "OneDollarInCurrency": 0                                // Type="Decimal"
        }
    ]
}
```

---

### Get a collection of all Currency Exhange Rate records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupExchangeRate/all?$orderby={{$company-group-exchange-rate-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$company-group-exchange-rate-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all Currency Exhange Rate records**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupExchangeRate",
    "value": [
        {
            "SuspensionId": 1,                                                  // Type="Int64"
            "ExchangeRateId": 1,                                                // Type="Int64"
            "Currency": "string",                                               // Type="String"
            "Country": "string",                                                // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",                       // Type="DateTime"
            "OneDollarInCurrency": 1.00000                                      // Type="Decimal"
        }
    ]
}
```

---

### Get single Currency Exchange Rate Record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupExchangeRate({{ExchangeRateId}})`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get single Currency Exchange Rate Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupExchangeRate",
    "value": [
        {
            "ExchangeRateId": 0,                                    // Type="Int64"
            "Currency": "string",                                   // Type="String"
            "Country": "string",                                    // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "OneDollarInCurrency": 0                                // Type="Decimal"
        }
    ]
}
```

---

### Create a single Currency Exchange Rate record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupExchangeRate`

**Request Body:**

```json
{
  //see "CompanyGroupExchangeRate" in metadata endpoint for available fields
}
```

**Example Response: Create a single company group exchange rate** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupExchangeRate",
    "value": [
        {
            "Currency": "string",                                   // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "OneDollarInCurrency": 0                                // Type="Decimal"
        }
    ]
}
```

---

### Update a single Currency Exchange Rate record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupExchangeRate({{ExchangeRateId}})`

**Request Body:**

```json
{
  //see "OrganizationUnit" in metadata endpoint for available fields
}
```

**Example Response: Update a single Currency Exchange Rate record** (Status: 204 No Content)

---

### Delete a single Currecny Exchange Rate record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupExchangeRate({{ExchangeRateId}})`

**Example Response: Delete a single Currecny Exchange Rate record** (Status: 200 OK)

---
## Workday

The Workday API Endpoint allows you to access the Audit Report data in order to utilize in your business environment.

### Audit Results

**`GET`** `{{api-base-url}}/v2.0/workday/results?date={{date}}&frequency={{frequency}}`

Returns a list of downloadable links for audit reports received on a specific day

> Please note that in order to access the endpoint, you need to have the `api.workday` scope. Therefore, when you authenticate, make sure to request the `api.workday` scope.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `date` | `{{date}}` | Required - Filter results by date. Format: yyyy-MM-dd. Only one date can be chosen. |
| `frequency` | `{{frequency}}` | Optional  - filter results by frequency<br/> The results for all frequencies can be retrieved or a filter of one frequency can be applied. |

**Example Response: Results** (Status: 200 OK)

```json
{
  "ResultFiles": [
    "url to log file"
  ]
}
```

---
## Training Course

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyTrainingCourse?$orderby={{$companytrainingcourse-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$companytrainingcourse-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of Company Training Courses records** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}//$metadata#CompanyTrainingCourse",
    "@odata.count": 0,
    "value": [
        {
            "CourseId": 1,									// Type="Int64"
            "EffectiveDate": "2007-06-12T00:00:00+02:00",   // Type="DateTime"
            "StopDate": "2007-06-12T00:00:00+02:00",        // Type="DateTime"
            "CourseDescription": "String",                  // Type="String"
            "CourseOverview": "String",                     // Type="String"
            "CourseRequirements": "String",                 // Type="String"
            "CourseCostNotes": "String",                    // Type="String"
            "InHouseOutsideProvider": "String",             // Type="Enum" see: InHouseOutsideProvider in metadata
            "Institution": "String",                        // Type="String"
            "Comments": "String",                           // Type="String"
            "MaxCandidates": 9,                             // Type="Int64"
            "MinCandidates": 3,                             // Type="Int64"
            "CourseDuration": "5 days",                     // Type="String"
            "Credits": 7,                                   // Type="Int64"
            "NqfLevel": "ABET 1",                           // Type="Enum" see: NqfLevel in metadata
            "SaqaRegistered": true,                         // Type="Boolean"
            "SaqaIdNumber": "String",                       // Type="String"
            "CourseType": "Seminar",                        // Type="Enum" see: CourseType in metadata
            "SkillsPriority": "05",                         // Type="Enum" see: TrainingSkillsPriority in metadata
            "CourseDurationHours": 0,                       // Type="Int64"
            "DirectCourseCost": 0.00,                       // Type="Decimal"
            "ProgramCategory": "B",                         // Type="Enum" see: TrainingProgramCategory in metadata
            "IsEquityTraining": true,                       // Type="Boolean"
            "SchedulesCount": 1                             // Type="Int64"
        }
    ]
}
```

---

### Get a single record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyTrainingCourse({{CourseId}})`

Retrieves a single record of `CompanyTrainingCourse` information based on the specified `CourseId`.

**Example Response: Get a single record** (Status: 200 OK)

```json
{
    "@odata.context": "https://apiuat.payspace.com/odata/v2.0/6464/$metadata#CompanyTrainingCourse",
    "CourseId": 1,									// Type="Int64"
    "EffectiveDate": "2007-06-12T00:00:00+02:00",   // Type="DateTime"
    "StopDate": "2007-06-12T00:00:00+02:00",        // Type="DateTime"
    "CourseDescription": "String",                  // Type="String"
    "CourseOverview": "String",                     // Type="String"
    "CourseRequirements": "String",                 // Type="String"
    "CourseCostNotes": "String",                    // Type="String"
    "InHouseOutsideProvider": "String",             // Type="Enum" see: InHouseOutsideProvider in metadata
    "Institution": "String",                        // Type="String"
    "Comments": "String",                           // Type="String"
    "MaxCandidates": 9,                             // Type="Int64"
    "MinCandidates": 3,                             // Type="Int64"
    "CourseDuration": "5 days",                     // Type="String"
    "Credits": 7,                                   // Type="Int64"
    "NqfLevel": "ABET 1",                           // Type="Enum" see: NqfLevel in metadata
    "SaqaRegistered": true,                         // Type="Boolean"
    "SaqaIdNumber": "String",                       // Type="String"
    "CourseType": "Seminar",                        // Type="Enum" see: CourseType in metadata
    "SkillsPriority": "05",                         // Type="Enum" see: TrainingSkillsPriority in metadata
    "CourseDurationHours": 0,                       // Type="Int64"
    "DirectCourseCost": 0.00,                       // Type="Decimal"
    "ProgramCategory": "B",                         // Type="Enum" see: TrainingProgramCategory in metadata
    "IsEquityTraining": true,                       // Type="Boolean"
    "SchedulesCount": 1                             // Type="Int64"        
}
```

---

### Create a record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyTrainingCourse`

Creates a single `CompanyTrainingCourse` record for the specified `company_id`.

**Request Body:**

```json
{
  //see "CompanyTrainingCourse" in metadata endpoint for available fields
}
```

**Example Response: Create a record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}//$metadata#CompanyTrainingCourse",
    "CourseId": 1,									// Type="Int64"
    "EffectiveDate": "2007-06-12T00:00:00+02:00",   // Type="DateTime"
    "StopDate": "2007-06-12T00:00:00+02:00",        // Type="DateTime"
    "CourseDescription": "String",                  // Type="String"
    "CourseOverview": "String",                     // Type="String"
    "CourseRequirements": "String",                 // Type="String"
    "CourseCostNotes": "String",                    // Type="String"
    "InHouseOutsideProvider": "String",             // Type="Enum" see: InHouseOutsideProvider in metadata
    "Institution": "String",                        // Type="String"
    "Comments": "String",                           // Type="String"
    "MaxCandidates": 9,                             // Type="Int64"
    "MinCandidates": 3,                             // Type="Int64"
    "CourseDuration": "5 days",                     // Type="String"
    "Credits": 7,                                   // Type="Int64"
    "NqfLevel": "ABET 1",                           // Type="Enum" see: NqfLevel in metadata
    "SaqaRegistered": true,                         // Type="Boolean"
    "SaqaIdNumber": "String",                       // Type="String"
    "CourseType": "Seminar",                        // Type="Enum" see: CourseType in metadata
    "SkillsPriority": "05",                         // Type="Enum" see: TrainingSkillsPriority in metadata
    "CourseDurationHours": 0,                       // Type="Int64"
    "DirectCourseCost": 0.00,                       // Type="Decimal"
    "ProgramCategory": "B",                         // Type="Enum" see: TrainingProgramCategory in metadata
    "IsEquityTraining": true                        // Type="Boolean"
}
```

---

### Update Single Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyTrainingCourse({{CourseId}})`

Update a single `CompanyTrainingCourse`record for the specified `CourseId`.

**Request Body:**

```json
{
  //see "CompanyTrainingCourse" in metadata endpoint for available fields
}
```

**Example Response: Update Single Record** (Status: 204 No Content)

---

### Delete a Single Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyTrainingCourse({{CourseId}})`

Delete a single `CompanyTrainingCourse` record for the specified `CourseId`.

**Example Response: Delete a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "https://apiuat.payspace.com/odata/v2.0/6464/$metadata#CompanyTrainingCourse",
    "Success": true             // Type="boolean"
}
```

---
## Company Review Process

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyReviewProcess?$top={{$top}}&$skip={{$skip}}&$orderby={{$company-review-process-field}}&$filter={{$filter}}`

Retrieves a list of active `CompanyReviewProcess` records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$orderby` | `{{$company-review-process-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#CompanyReviewProcess",
    "value": [
        {
            "ProcessId": 0,                                     // Type="Int64"
            "CompanyId": 0,                                     // Type="Int64"
            "ProcessTypeDescription": "String",                 // Type="String"
            "ProcessStatus": "In Progress",                     // Type="Enum" see: ReviewProcessStatus
            "Title": "String",                                  // Type="String"
            "StartDate": "DateTime",                            // Type="DateTime"
            "EndDate": "DateTime",                              // Type="DateTime"
            "ProcessWeighting": 0.00                            // Type="Decimal"
        }
    ]
}
```

---
## Settings

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanySetting?$top={{$top}}&$skip={{$skip}}&$orderby={{fieldname}}&$filter={{$filter}}`

Retrieves Company Settings for the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$orderby` | `{{fieldname}}` | Optional (string) - Specifies the order in which items are returned |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Code eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanySetting",
    "value": [
        {
            "Code": 0,                 // Type="String"
            "Enabled": false,          // Type="Boolean"
            "Value": "String",         // Type="String"
            "Description": "string",   // Type="String"
            "Screen": "String"         // Type="String"
        }
    ]
}
```

---
## Organization Grade

### Get a collection of Organization  Grades

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationGrade?$orderby={{$organization-grade-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$organization-grade-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Description eq 'Grade 01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of Organization  Grades**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationGrade",
    "@odata.count": 1,
    "value": [
        {
            "GradeId": 0,                                           // Type="Int64"
            "Description": "string",                                // Type="String"
            "Code": "string",                                       // Type="String"
            "Minimum": 0,                                           // Type="Decimal"
            "Maximum": 0,                                           // Type="Decimal"
            "BonusPercentage": 0,                                   // Type="Decimal"
            "Average": 0,                                           // Type="Decimal"
            "IsEligiblePm": true,                                   // Type="Boolean"
            "AdditionalValues": [
                {
                    "Description": "string",                        // Type="String"
                    "Value": "string"                               // Type="String"
                }
            ]
        }
    ]
}
```

---

### Get a single Organization Grades record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationGrade({{OrganizationGradeId}})`

**Example Response: Get a single Organization Grades record**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationGrade/$entity",
    "GradeId": 0,                                           // Type="Int64"
    "Description": "string",                                // Type="String"
    "Code": "string",                                       // Type="String"
    "Minimum": 0,                                           // Type="Decimal"
    "Maximum": 0,                                           // Type="Decimal"
    "BonusPercentage": 0,                                   // Type="Decimal"
    "Average": 0,                                           // Type="Decimal"
    "IsEligiblePm": true,                                   // Type="Boolean"
    "AdditionalValues": [
        {
            "Description": "string",                        // Type="String"
            "Value": "string"                               // Type="String"
        }
    ]
}
```

---

### Delete a single Organization Grades record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationGrade({{OrganizationGradeId}})`

**Example Response: Get a single Organization Grades record**

---

### Update a single Organization Grades record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationGrade({{OrganizationGradeId}})`

**Request Body:**

```json
{
      //see "OrganizationGrade" in metadata endpoint for available fields
}
```

**Example Response: Update a single Organization Grades record** (Status: 204 No Content)

---

### Create a single Organization Grade record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/OrganizationGrade`

**Request Body:**

```json
{
  //see "OrganizationGrade" in metadata endpoint for available fields

    "AdditionalValues": [
                // See list of fields available in CompanyGradeFieldValue lookup endpoint
            ]    
}
```

**Example Response: Create a single Organization Grade record**

```json
 {
        "GradeId": 0,                                           // Type="Int64"
        "Description": "string",                                // Type="String"
        "Code": "string",                                       // Type="String"
        "Minimum": 0,                                           // Type="Decimal"
        "Maximum": 0,                                           // Type="Decimal"
        "BonusPercentage": 0,                                   // Type="Decimal"
        "Average": 0,                                           // Type="Decimal"
        "IsEligiblePm": true,                                   // Type="Boolean"
        "AdditionalValues": [
            {
                "Description": "string",                        // Type="String"
                "Value": "string"                                     // Type="String"
            }
        ]
}
```

---
## Other Dropdowns

### Action Types

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupActionType?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of action type records based on the specified `company_id.`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "ActionTypeCode eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupActionType",
    "value": [
        {
            "ActionTypeId": 0,                            // Type="Int64"
            "ActionTypeCode": "string",                   // Type="String"
            "ActionTypeDescription": "string"             // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupActionType`

Create a single action type record

**Request Body:**

```json
{
  // see "CompanyGroupActionType" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "ActionTypeCode": "string",                   // Type="String"
    "ActionTypeDescription": "string"             // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupActionType({{ActionTypeId}})`

Update a single action type record based on the specified `ActionTypeId`.

**Request Body:**

```json
{
  // see "CompanyGroupActionType" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGroupActionType({{ActionTypeId}})`

Deletes action type record based on the specified `ActionTypeId`

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupActionType/$entity",
  "Success": true
}
```

---

### Attachment Classification

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyAttachmentClassification?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of attachment classification records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Classification eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyAttachmentClassification",
    "value": [
        {
            "ClassificationId": 0,                  // Type="Int64"
            "Classification": "string",             // Type="String"
            "ExcludeSecurityRoles": "role1;role2"   // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyAttachmentClassification`

Create a single attachment classification record.

**Request Body:**

```json
{
  // see "CompanyAttachmentClassification" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyAttachmentClassification/$entity",

    "ClassificationId": 0,                  // Type="Int64"
    "Classification": "string",             // Type="String"
    "ExcludeSecurityRoles": "role1;role2"   // Type="String"
}

```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyAttachmentClassification({{CompanyAttachmentClassificationId}})`

Update a single attachment classification record based on the specified `CompanyAttachmentClassificationId`.

**Request Body:**

```json
{
    // see "CompanyAttachmentClassification" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyAttachmentClassification({{CompanyAttachmentClassificationId}})`

Deletes attachment classification record based on the specified `CompanyAttachmentClassificationId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyAttachmentClassification/$entity",
  "Success": true
}
```

---

### Bank Names

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyBankName?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of bank name records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "BankName eq 'FNB'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyBankName",
    "value": [
        {
            "BankNameId": 0,                      // Type="Int64"
            "BankName": "string",                 // Type="String"
            "DefaultBranchCode": "string"         // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyBankName`

Create a single bank name record

**Request Body:**

```json
{
  // see "CompanyBankName" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "BankName": "string",                   // Type="String"
    "DefaultBranchCode": "string"             // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyBankName({{BankNameId}})`

Update a single bank name record based on the specified `BankNameId`.

**Request Body:**

```json
{
  // see "CompanyBankName" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyBankName({{BankNameId}})`

Deletes bank name record based on the specified `BankNameId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyBankName/$entity",
  "Success": true
}
```

---

### Component Group

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyComponentGroup?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of component group records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "ComponentGroupCode eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyComponentGroup",
    "value": [
        {
            "ComponentGroupId": 0,                      // Type="Int64"
            "ComponentGroupCode": "string",                   // Type="String"
            "ComponentGroupDescription": "string"             // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyComponentGroup`

Create a single component group record

**Request Body:**

```json
{
  // see "CompanyComponentGroup" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "ComponentGroupCode": "string",                   // Type="String"
    "ComponentGroupDescription": "string"             // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyComponentGroup({{ComponentGroupId}})`

Update a single component group record based on the specified `ComponentGroupId`.

**Request Body:**

```json
{
  // see "CompanyComponentGroup" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyComponentGroup({{ComponentGroupId}})`

Deletes component group record based on the specified `ComponentGroupId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyComponentGroup/$entity",
  "Success": true
}
```

---

### Employee Bank Detail Indicator

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEbdIndicator?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of employee bank detail indicator records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EbdIndicator eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEbdIndicator",
    "value": [
        {
            "EbdIndicatorId": 0,                      // Type="Int64"
            "EbdIndicator": "string"                  // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEbdIndicator`

Create a single employee bank detail indicator record

**Request Body:**

```json
{
  // see "CompanyEbdIndicator" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "EbdIndicator": "string"                  // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEbdIndicator({{EbdIndicatorId}})`

Update a single employee bank detail indicator record based on the specified `EbdIndicatorId`.

**Request Body:**

```json
{
  // see "CompanyEbdIndicator" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEbdIndicator({{EbdIndicatorId}})`

Deletes employee bank detail indicator record based on the specified `EbdIndicatorId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEbdIndicator/$entity",
  "Success": true
}
```

---

### Employment Category

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentCategory?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of employement category records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "CategoryCode eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEmploymentCategory",
    "value": [
        {
            "CategoryId": 0,                // Type="Int64"
            "CategoryCode": "string",       // Type="String"
            "CategoryDescription": "string" // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentCategory`

Create a single employement category record.

**Request Body:**

```json
{
  // see "CompanyEmploymentCategory" in metadata endpoint for available fields
}

```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEmploymentCategory/$entity",

    "CategoryId": 0,                // Type="Int64"
    "CategoryCode": "string",       // Type="String"
    "CategoryDescription": "string" // Type="String"
}

```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentCategory({{EmploymentCategoryId}})`

Update a single employement category record based on the specified `CategoryId`.

**Request Body:**

```json
{
  // see "CompanyEmploymentCategory" in metadata endpoint for available fields
}

```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentCategory({{EmploymentCategoryId}})`

Deletes employement category record based on the specified `CategoryId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEmploymentCategory/$entity",
  "Success": true
}
```

---

### Employment SubCategory

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentSubCategory?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of employement subcategory records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "SubCategoryCode eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEmploymentSubCategory",
    "value": [
        {
            "SubCategoryId": 0,             // Type="Int64"
            "SubCategoryCode": "string",    // Type="String"
            "Description": "string"         // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentSubCategory`

Create a single employement subcategory record.

**Request Body:**

```json
{
  // see "CompanyEmploymentSubCategory" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEmploymentSubCategory/$entity",
    "SubCategoryId": 0,             // Type="Int64"
    "SubCategoryCode": "string",    // Type="String"
    "Description": "string"         // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentSubCategory({{EmploymentSubCategoryId}})`

Update a single employement subcategory record based on the specified `SubCategoryId`.

**Request Body:**

```json
{
  // see "CompanyEmploymentSubCategory" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyEmploymentSubCategory({{EmploymentSubCategoryId}})`

Deletes employement subcategory record based on the specified `SubCategoryId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEmploymentSubCategory/$entity",
  "Success": true
}
```

---

### Incident Legal Bodies

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLegalBody?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of incident legal body records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "LegalBodyCode eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLegalBody",
    "value": [
        {
            "CompanyLegalBodyId": 0,            // Type="Int64"
            "LegalBodyDescription": "string",   // Type="String"
            "LegalBodyCode": "string"           // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLegalBody`

Create a single incident legal body record.

**Request Body:**

```json
{
  // see "CompanyLegalBody" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLegalBody/$entity",
    "CompanyLegalBodyId": 0,            // Type="Int64"
    "LegalBodyDescription": "string",   // Type="String"
    "LegalBodyCode": "string"           // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLegalBody({{CompanyLegalBodyId}})`

Update a single incident legal body record based on the specified `CompanyLegalBodyId`.

**Request Body:**

```json
{
  // see "CompanyLegalBody" in metadata endpoint for available fields
}

```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLegalBody({{CompanyLegalBodyId}})`

Deletes incident legal body record based on the specified `CompanyLegalBodyId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLegalBody/$entity",
  "Success": true
}
```

---

### Incident Offences

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyOffence?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of incident offence records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "OffenceCode eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyOffence",
    "value": [
        {
            "CompanyOffenceId": 0,          // Type="Int64"
            "OffenceCode": "string",        // Type="String"
            "OffenceDescription": "string"  // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyOffence`

Create a single incident offence record.

**Request Body:**

```json
{
  // see "CompanyOffence" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyOffence/$entity",
    "CompanyOffenceId": 0,          // Type="Int64"
    "OffenceCode": "string",        // Type="String"
    "OffenceDescription": "string"  // Type="String"
}

```

---

#### Update Company Offence

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyOffence({{CompanyOffenceId}})`

Update a single incident offence record based on the specified `CompanyOffenceId`.

**Request Body:**

```json
{
  // see "CompanyOffence" in metadata endpoint for available fields
}
```

**Example Response: Update Company Offence** (Status: 204 No Content)

---

#### Delete Company Offence

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyOffence({{CompanyOffenceId}})`

Deletes incident offence record based on the specified `CompanyOffenceId`.

**Example Response: Delete Company Offence** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyOffence/$entity",
  "Success": true
}
```

---

### Incident Types

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncidentType?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of incident type records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "IncidentType eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyIncidentType",
    "value": [
        {
            "CompanyIncidentTypeId": 0,     // Type="Int64"
            "IncidentType": "string"        // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncidentType`

Create a single incident type record.

**Request Body:**

```json
{
  // see "CompanyIncidentType" in metadata endpoint for available fields
}

```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyIncidentType/$entity",
    "CompanyIncidentTypeId": 0,     // Type="Int64"
    "IncidentType": "string"        // Type="String"
}

```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncidentType({{CompanyIncidentTypeId}})`

Update a single incident type record based on the specified `CompanyIncidentTypeId`.

**Request Body:**

```json
{
  // see "CompanyIncidentType" in metadata endpoint for available fields
}

```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncidentType({{CompanyIncidentTypeId}})`

Deletes incident type record based on the specified `CompanyIncidentTypeId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyIncidentType/$entity",
  "Success": true
}
```

---

### Increase Reasons

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncreaseReason?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of increase reason records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "ReasonCode eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyIncreaseReason",
    "value": [
        {
            "ReasonId": 0,          // Type="Int64"
            "ReasonCode": "string", // Type="String"
            "Description": "string" // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncreaseReason`

Create a single increase reason record.

**Request Body:**

```json
{
  // see "CompanyIncreaseReason" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyIncreaseReason/$entity",
    "ReasonId": 0,          // Type="Int64"
    "ReasonCode": "string", // Type="String"
    "Description": "string" // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncreaseReason({{IncreaseReasonId}})`

Update a single increase reason record based on the specified `ReasonId`.

**Request Body:**

```json

{
  // see "CompanyIncreaseReason" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyIncreaseReason({{IncreaseReasonId}})`

Deletes increase reason record based on the specified `ReasonId`.

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyIncreaseReason/$entity",
  "Success": true
}
```

---

### Leave Reasons

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLeaveReason?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of leave reason records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "LeaveReason eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLeaveReason",
    "value": [
        {
            "LeaveReasonId": 0,             // Type="Int64"
            "CompanyLeaveSetup": "string",  // Type="String"
            "LeaveType": "string",          // Type="String": {{api-base-url}}/odata/v2.0/:company-id/Lookup/LeaveType
            "LeaveReason": "string"         // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLeaveReason`

Create a single leave reason record.

**Request Body:**

```json
{
  // see "CompanyLeaveReason" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLeaveReason/$entity",
    "LeaveReasonId": 0,             // Type="Int64"
    "CompanyLeaveSetup": "string",  // Type="String"
    "LeaveType": "string",          // Type="String": {{api-base-url}}/odata/v2.0/:company-id/Lookup/LeaveType
    "LeaveReason": "string"         // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLeaveReason({{LeaveReasonId}})`

Update a single leave reason record based on the specified `LeaveReasonId`.

**Request Body:**

```json
{
  // see "CompanyLeaveReason" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyLeaveReason({{LeaveReasonId}})`

Deletes leave reason record based on the specified `LeaveReasonId`

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLeaveReason/$entity",
  "Success": true
}
```

---

### Position Family

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFamily?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of position family records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Code eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyPositionFamily",
    "value": [
        {
            "FamilyId": 0,                      // Type="Int64"
            "Code": "string",                   // Type="String"
            "Description": "string"             // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFamily`

Create a single position family record.

**Request Body:**

```json
{
  // see "CompanyPositionFamily" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "Code": "string",                   // Type="String"
    "Description": "string"             // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFamily({{PositionFamilyId}})`

Update a single position family record based on the specified `LeaveReasonId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  // see "CompanyPositionFamily" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFamily({{PositionFamilyId}})`

Deletes position family record based on the specified `LeaveReasonId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyPositionFamily/$entity",
  "Success": true
}
```

---

### Position Functional Area

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFunctionArea?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of position functional area records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Code eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyPositionFunctionArea",
    "value": [
        {
            "FunctionAreaId": 0,        // Type="Int64"
            "Code": "string",           // Type="String"
            "Description": "string"     // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFunctionArea`

Create a single position functional area record.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  // see "CompanyPositionFunctionArea" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "Code": "string",                   // Type="String"
    "Description": "string"             // Type="String"
}
```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFunctionArea({{FunctionAreaId}})`

Update a single position functional area record based on the specified `FunctionAreaId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  // see "CompanyPositionFunctionArea" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyPositionFunctionArea({{FunctionAreaId}})`

Deletes position family record based on the specified `FunctionAreaId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyPositionFunctionArea/$entity",
  "Success": true
}
```

---

### Request Type

#### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyManagerRequestType?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}`

Retrieves a list of request type records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyManagerRequestType",
    "value": [
        {
            "RequestTypeId": 0,                 // Type="Int64"
            "RequestTypeName": "string",        // Type="String"
            "IsAttachmentReq": false,           // Type="Boolean"
            "RequestTypeDescription": "string", // Type="String"
            "ExcludeSecurityRoles": "ESS",      // Type="String"
            "RequestTypeAttachments": [
                0                               // Type="Int64"
            ]
        }
    ]
}
```

---

#### Get Company Attachments collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyManagerRequestType/company-attachments?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}`

Retrieves a list of company attachment records based on the specified `company_id.`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get Company Attachments collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyAttachment",
    "value": [
        {
            "CompanyAttachmentId": 0,       // Type="Int64"
            "SectionName": "string",        // Type="String"
            "AttachmentTitle": "string",    // Type="String"
            "AttachmentName": "string"      // Type="String"
        }
    ]
}
```

---

#### Create a single record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyManagerRequestType`

Create a single request types record.

**Request Body:**

```json
{
  // see "CompanyManagerRequestType" in metadata endpoint for available fields
}
```

**Example Response: Create a single record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyManagerRequestType/$entity",
    "RequestTypeId": 0,                     // Type="Int64"
    "RequestTypeName": "string",            // Type="String"
    "IsAttachmentReq": false,               // Type="Boolean"
    "RequestTypeDescription": "string",     // Type="String"
    "ExcludeSecurityRoles": "role1;role2",  // Type="String"
    "RequestTypeAttachments": [
        0                                   // Type="Int64"
    ]
}

```

---

#### Update a single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyManagerRequestType({{CompanyManagerRequestTypeId}})`

Update a single request type record based on the specified `RequestTypeId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
    // see "CompanyManagerRequestType" in metadata endpoint for available fields
}
```

**Example Response: Update a single record** (Status: 204 No Content)

---

#### Delete a single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyManagerRequestType({{CompanyManagerRequestTypeId}})`

Deletes request type record based on the specified `RequestTypeId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyManagerRequestType/$entity",
  "Success": true
}
```

---

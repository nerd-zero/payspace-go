# Employee > Leave

## Adjustments

> When `OverrideBalance` is set to `true` the employees leave balance will be equal to the value set on the `NoOfDays` field. The system will change the `NoOfDays` value provided to be the difference of what the existing balance is and what the `NoOfDays` value is, that is being advised.

### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveAdjustment?$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Get collection of leave adjustments based on the `companyId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveAdjustment",
    "value": [
        {
            "OverrideBalance": true,            // Type="Boolean"
            "LeaveCompanyRun": "String",        // Type="String"   
            "NoOfDays": "Decimal",              // Type="String"
            "LeaveAdjustmentId": 0,             // Type="Int64"
            "EmployeeNumber": "String",         // Type="String"
            "LeaveBucket": "String",            // Type="String"
            "LeaveType": "String",              // Type="String"
            "Comments": "String",               // Type="String"
            "Reference": "String",              // Type="String"
            "LastModifiedDate": "DateTime"      // Type="DateTime" 
        }
    ]
}
```

---

### Get Collection By Period

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveAdjustment/:year/:month?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Returns a collection of `EmployeeLeaveAdjustment` records applicable (influenced the balance) to the `year` and `month` specified.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get by period**

```json

{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveAdjustment",
    "value": [
        {           
            "OverrideBalance": false,       // Type="Boolean"
            "LeaveCompanyRun": "String",    // Type="String"
            "NoOfDays": 0.0,                // Type="Decimal"
            "LeaveAdjustmentId": 0,         // Type="Int64"
            "EmployeeNumber": "String",     // Type="String"
            "LeaveBucket": "String",        // Type="String"
            "LeaveType": "String",          // Type="String"
            "Comments": "String",           // Type="String"
            "Reference": "String",          // Type="String"
            "LastModifiedDate": "DateTime"  // Type="DateTime"
        }
    ]
}
```

---

### Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveAdjustment({{LeaveAdjustmentId}})`

Get single leave adjustment record based on the specified `LeaveAdjustmentId`.

**Example Response: Get Single** (Status: 200 OK)

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0//{{company-id}}/$metadata#EmployeeLeaveAdjustment/$entity",
    "OverrideBalance": true,            // Type="Boolean"
    "LeaveCompanyRun": "String",        // Type="String"   
    "NoOfDays": "Decimal",              // Type="String"
    "LeaveAdjustmentId": 0,             // Type="Int64"
    "EmployeeNumber": "String",         // Type="String"
    "LeaveBucket": "String",            // Type="String"
    "LeaveType": "String",              // Type="String"
    "Comments": "String",               // Type="String"
    "Reference": "String",              // Type="String"
    "LastModifiedDate": "DateTime"      // Type="DateTime"
}
```

---

### Patch exisiting

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveAdjustment({{LeaveAdjustmentId}})`

Update a single leave adjustment record based on the specified `LeaveAdjustmentId`.

**Example Response: Patch exisiting**

---

### Create new

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveAdjustment`

Create new leave adjustment record.

**Example Response: Create new**

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeLeaveAdjustment/$entity",
    "OverrideBalance": true,        	// Type="Boolean"
    "LeaveCompanyRun": "String", 	    // Type="String"   
    "NoOfDays": "Decimal",	       		// Type="String"
    "LeaveAdjustmentId": 0,	            // Type="Int64"
    "EmployeeNumber": "String",	       	// Type="String"
    "LeaveBucket": "String",	       	// Type="String"
    "LeaveType": "String",	       		// Type="String"
    "Comments": "String",	       		// Type="String"
    "Reference": "String"	       		// Type="String"
}
```

---

### Delete exisiting

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveAdjustment({{EmployeeLeaveAdjustmentId}})`

Delete single leave adjustment record based on the specified `LeaveAdjustmentId`.

---

### [OBSOLETE] Get a collection of leave transactions

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/LeaveAdjustment/:year/:month?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

> This API endpoint is Deprecated. It should not be used

Returns a collection of leave adjustments active during the run in the specified period.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get by period**

```json

{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveAdjustment",
    "value": [
        {
            "LeaveAdjustmentId": 0,                             // Type="Int64"
            "EmployeeNumber": "String",                         // Type="String"
            "LeaveBucket": "String",                            // Type="String"
            "LeaveType": "String",                              // Type="String"
            "LeaveStartDate": "2018-07-27T00:00:00+02:00",      // Type="DateTime"
            "LeaveEndDate": "2018-07-27T00:00:00+02:00",        // Type="DateTime"
            "NoOfDays": 1.0,                                    // Type="Decimal"
            "LeaveStatus": "Approved",                          // Type="String"
            "HasAttachment": false,                             // Type="Boolean"
            "Comments": "String",                               // Type="String"
            "LeaveReason": "String",                            // Type="String"
            "Reference": "String",                              // Type="String"
            "TransactionType": "Leave Application"              // Type="String"
        }
    ]
}
```

---
## Setup Entitlements

Leave Setup Entitlements can only be viewed and modified if the Setup being referenced has the `Employee Defined` checkbox enabled on the "Company Leave Scheme Parameters" screen.

### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetupEntitlement?$top={{top}}&$skip={{$skip}}&$count={{$count}}&$filter={{filter}}`

Get collection of leave setup entitlement records based on the `companyId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get Collection**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveSetupEntitlement",
    "value": [
        {
            "LeaveSetupEntitlementId": 0,   // Type="Int64"
            "EmployeeNumber": "String",     // Type="String"
            "LeaveSetupId": 0,              // Type="Decimal"
            "CompanyLeaveSetup": "String",  // Type="String"
            "Accrual Value": 0,             // Type="Decimal"
            "FullName": "String"            // Type="String"
        }
    ]
}

```

---

### Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetupEntitlement({{LeaveSetupEntitlementId}})`

Get single leave setup entitlement record based on the specified `LeaveSetupEntitlementId`.

**Example Response: Get Single**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveSetupEntitlement/$entity",
    "LeaveSetupEntitlementId": 0,   // Type="Int64"
    "EmployeeNumber": "String",     // Type="String"
    "LeaveSetupId": 0,              // Type="Int64"
    "CompanyLeaveSetup": "String",  // Type="String"
    "Accrual Value": 0,             // Type="Decimal"
    "FullName": "String"            // Type="String"
}
```

---

### Get a collection of all employee leave setup entitlements

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetupEntitlement/all?$orderby={{employee-leave-setup-entitlement-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-leave-setup-entitlement-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee leave setup entitlements**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveSetup",
    "value": [
        {
            "LeaveSetupEntitlementId": 1,                                   // Type="Int64"
            "EmployeeNumber": "string",                                     // Type="String"
            "LeaveSetupId": 0,                                              // Type="Int64"
            "CompanyLeaveSetup": "string",                                  // Type="String"
            "AccrualValue": 1.00000,                                        // Type="Decimal"
            "FullName": "string"                                            // Type="String"
        }
    ]
}
```

---

### Patch Existing

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetupEntitlement({{LeaveSetupEntitlementId}})`

Update a single employee leave setup entitlement record based on the specified `LeaveSetupEntitlementId`.

**Example Response: Patch Existing**

---

### Create New

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetupEntitlement`

To create a new employee leave setup entitlement record the following endpoints should be used:

- The `EmployeeLeaveSetup` endpoint to obtain the `LeaveSetupId`.
- The `CompanyLeaveSetup` endpoint to obtain the correct `CompanyLeaveSetup` format (i.e. "Scheme Name - Setup Description").

**Example Response: Create New**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveSetupEntitlement/$entity",
    "LeaveSetupEntitlementId": 0,   // Type="Int64"
    "EmployeeNumber": "String",     // Type="String"
    "LeaveSetupId": 0,              // Type="Int64"
    "CompanyLeaveSetup": "String",  // Type="String"
    "Accrual Value": 0,             // Type="Decimal"
    "FullName": "String"            // Type="String"
}

```

---
## Setup

> If no Leave Bucket Mapping is configured and an employee was previously attached to a leave scheme, please remember to migrate all leave balances from the employee’s previous leave scheme accordingly.

### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetup?$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Get collection of leave setups based on the `companyId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveSetup",
    "value": [
        {
            "LeaveSetupId": 0,                  // Type="Int64"
            "CompanyLeaveScheme": "String",     // Type="String"
            "EmployeeNumber": "String",         // Type="String"
            "FullName": "String",               // Type="String"
            "EffectiveDate": "DateTime",        // Type="DateTime"
            "PublicHolidayCategory": "String"   // Type="String"
        }
    ]
}
```

---

### Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetup({{LeaveSetupId}})`

Get single leave setup record based on the specified `LeaveSetupId`.

**Example Response: Get Single** (Status: 200 OK)

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0//{{company-id}}/$metadata#EmployeeLeaveSetup/$entity",
    "LeaveSetupId": 0,                  // Type="Int64"
    "CompanyLeaveScheme": "String",     // Type="String"
    "EmployeeNumber": "String",         // Type="String"
    "FullName": "String",               // Type="String"
    "EffectiveDate": "DateTime",        // Type="DateTime"
    "PublicHolidayCategory": "String"   // Type="String"
}
```

---

### Get a collection of all employee leave setups

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetup/all?$orderby={{employee-leave-setup-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-leave-setup-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee leave setups**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveSetup",
    "value": [
        {
            "LeaveSetupId": 1,                                      // Type="Int64"
            "CompanyLeaveScheme": "string",                         // Type="String"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "PublicHolidayCategory": "string"                       // Type="String"
        }
    ]
}
```

---

### Create new

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetup`

Create new leave setup record.

**Example Response: Create new**

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeLeaveSetup/$entity",
    "LeaveSetupId": 0,                  // Type="Int64"
    "CompanyLeaveScheme": "String",     // Type="String"
    "EmployeeNumber": "String",         // Type="String"
    "FullName": "String",               // Type="String"
    "EffectiveDate": "DateTime",        // Type="DateTime"
    "PublicHolidayCategory": "String"   // Type="String"
}
```

---

### Delete exisiting

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveSetup({{EmployeeLeaveSetupId}})`

Delete single leave setup record based on the specified `LeaveSetupId`.

---
## Application

### Skip Leave Balance Validation

- You can skip leave balance validation by setting `SkipValidation` to `true`.
- It will disable the leave balance validation and commit what is advised.
    

### Workflow

- Workflow can be enabled by adding the `allowWorkflow` query parameter and setting it to `true`
- This action will only initialize a workflow on `POST`
- Enabling this will also enforce requirements for attachments that may have been enabled for the leave type (follow attachment instructions below for this).
    

### Attachments

- Saving of attachments can be enabled by adding the `validateAttachments` query parameter and setting it to `true`
- Setting this to `true` will enforce attachment requirements for the leave type
- An attachment will only be saved to a leave application if this parameter is set to true.
- Attachment must be posted to the `FileUpload` and the resulting guid response of that post must be added to the leave application as a value for the `"TempAttachmentSignature"` property

### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveApplication?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Get collection of leave applications based on the `companyId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` |  |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get Collection**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLeaveApplication",
    "value": [
    {
        "LeaveReason": "String",            // Type="String"   
        "LeaveStartDate": "DateTime",       // Type="DateTime"   
        "LeaveEndDate": "DateTime",         // Type="DateTime"   
        "LeaveStatus": "string",            // Type="String"   
        "LeaveCompanyRun": "String",        // Type="String"   
        "LeaveTransactionType": "String",   // Type="String"   
        "SkipValidation": true,             // Type="Boolean"   
        "LeaveAdjustmentId": 0,             // Type="Int64"   
        "EmployeeNumber": "String",         // Type="String"   
        "LeaveBucket": "String",            // Type="String"   
        "LeaveType": "String",              // Type="String"   
        "NoOfDays": 0,                      // Type="Double" 
        "Comments": "String",               // Type="String"
        "CancellationId": 0,                // Type="Int64"   
        "Reference": "String",              // Type="String"
        "AttachmentUrl": "String",          // Type="String"
        "LastModifiedDate": "DateTime"      // Type="DateTime"  
    }]
}
```

---

### Get Collection By Period

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveApplication/:year/:month?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

GET a collection of `EmployeeLeaveApplication` records applicable (influenced the balance) to the `year` and `month` specified.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'".<br>See https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get by period**

```json

{
    "@odata.context": "https://localhost:44393/odata/v2.0//{{company-id}}/$metadata#EmployeeLeaveApplication/$entity",
    "value": [
        {
            "LeaveReason": "String",            // Type="String"   
            "LeaveStartDate": "DateTime",       // Type="DateTime"   
            "LeaveEndDate": "DateTime",         // Type="DateTime"   
            "LeaveStatus": "string",            // Type="String"   
            "LeaveCompanyRun": "String",        // Type="String"   
            "LeaveTransactionType": "String",   // Type="String"   
            "SkipValidation": true,             // Type="Boolean"   
            "LeaveAdjustmentId": 0,             // Type="Int64"   
            "EmployeeNumber": "String",         // Type="String"   
            "LeaveBucket": "String",            // Type="String"   
            "LeaveType": "String",              // Type="String"   
            "NoOfDays": 0,                      // Type="Double" 
            "Comments": "String",               // Type="String"
            "CancellationId": 0,                // Type="Int64"  
            "Reference": "String",              // Type="String"  
            "AttachmentUrl": "String",          // Type="String"
            "LastModifiedDate": "DateTime"      // Type="DateTime"
        }
    ]
}
```

---

### Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveApplication({{LeaveAdjustmentId}})`

Get single leave application record based on the specified `LeaveAdjustmentId`.

**Example Response: Get Single** (Status: 200 )

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0//{{company-id}}/$metadata#EmployeeLeaveApplication/$entity",
    "LeaveReason": "String",            // Type="String"   
    "LeaveStartDate": "DateTime",       // Type="DateTime"   
    "LeaveEndDate": "DateTime",         // Type="DateTime"   
    "LeaveStatus": "string",            // Type="String"   
    "LeaveCompanyRun": "String",        // Type="String"   
    "LeaveTransactionType": "String",   // Type="String"   
    "SkipValidation": true,             // Type="Boolean"   
    "LeaveAdjustmentId": 0,             // Type="Int64"   
    "EmployeeNumber": "String",         // Type="String"   
    "LeaveBucket": "String",            // Type="String"   
    "LeaveType": "String",              // Type="String"   
    "NoOfDays": 0,                      // Type="Double" 
    "Comments": "String",               // Type="String"
    "CancellationId": 0,                // Type="Int64"   
    "Reference": "String",              // Type="String"
    "AttachmentUrl": "String",          // Type="String"
    "LastModifiedDate": "DateTime"      // Type="DateTime"
}
```

---

### Patch exisiting

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveApplication({{LeaveAdjustmentId}})?validateAttachments={{validate-attachments}}`

Update single leave application record based on the specified `LeaveAdjustmentId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `validateAttachments` | `{{validate-attachments}}` | Enables the attachment validations for a leave application based on the leave bucket |

**Example Response: Patch exisiting** (Status: 204 No Content)

---

### Create new

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveApplication?allowWorkflow={{allow-workflow}}&validateAttachments={{validate-attachments}}`

Create new leave application record.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `allowWorkflow` | `{{allow-workflow}}` | Allows workflow processing on a leave application |
| `validateAttachments` | `{{validate-attachments}}` | Enables the attachment validations for a leave application based on the leave bucket |

**Example Response: Create new** (Status: 201 Created)

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0//{{company-id}}/$metadata#EmployeeLeaveApplication/$entity",
    "LeaveReason": "String",            // Type="String"   
    "LeaveStartDate": "DateTime",       // Type="DateTime"   
    "LeaveEndDate": "DateTime",         // Type="DateTime"   
    "LeaveStatus": "string",            // Type="String"   
    "LeaveCompanyRun": "String",        // Type="String"   
    "LeaveTransactionType": "String",   // Type="String"   
    "SkipValidation": true,             // Type="Boolean"   
    "LeaveAdjustmentId": 0,             // Type="Int64"   
    "EmployeeNumber": "String",         // Type="String"   
    "LeaveBucket": "String",            // Type="String"   
    "LeaveType": "String",              // Type="String"   
    "NoOfDays": 0,                      // Type="Double" 
    "Comments": "String",               // Type="String"   
    "Reference": "String",              // Type="String"
    "AttachmentUrl": "String"           // Type="String" 
}
```

---

### Delete exisiting

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLeaveApplication({{LeaveAdjustmentId}})`

Delete single leave application record based on the specified `LeaveAdjustmentId`.

---

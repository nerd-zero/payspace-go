# Employee > Suspension

Retrieve suspension information for all employees associated with the `company_id`.

## Get a collection of suspensions

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSuspension?$orderby={{$employee-suspension-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee suspension based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-suspension-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeSuspensionId eq 1234". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of suspensions** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSuspension",
    "@odata.count": 2,
    "value": [
        {
            "EmployeeNumber": "string",                             // Type="String"
            "SuspensionId": 0,                                      // Type="Int64"
            "SuspensionReason": "string",                           // Type="String"
            "EffectiveDate": "2020-09-01T00:00:00+02:00",           // Type="DateTime"     
            "ReturnDate": "2020-09-22T00:00:00+02:00",              // Type="DateTime"
            "MinimumPeriods": 0,                                    // Type="Int64"
            "PercentageOfPay": 0,                                   // Type="Int64"
            "OverridingAmount": null,                               // Type="Int64"
            "SuspensionPaymentMethod": "string",                    // Type="String"
            "DoNotCalculateRecurringComponents": true,              // Type="Boolean"
            "UI19DownloadUrl": "string",                            // Type="String"
            "CustomFields": [
                {			
                    "Code": "string",				                // Type="String"
                    "Label": "string",			                    // Type="String"
                    "Value": "string"			        	        // Type="String"
				}
            ]
        }
    ]
}
```

---
## Get a single suspension record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSuspension({{SuspensionId}})`

Retrieves suspension information based on the specified `SuspensionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single suspension record**

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeSuspension/$entity",
    "EmployeeNumber": "string",                             // Type="String"
    "SuspensionId": 0,                                      // Type="Int64"
    "SuspensionReason": "string",                           // Type="String"
    "EffectiveDate": "2020-09-01T00:00:00+02:00",           // Type="DateTime"     
    "ReturnDate": "2020-09-22T00:00:00+02:00",              // Type="DateTime"
    "MinimumPeriods": 0,                                    // Type="Int64"
    "PercentageOfPay": 0,                                   // Type="Int64"
    "OverridingAmount": null,                               // Type="Int64"
    "SuspensionPaymentMethod": "string",                    // Type="String"
    "DoNotCalculateRecurringComponents": true,              // Type="Boolean"
    "UI19DownloadUrl": "string",                            // Type="String"
    "CustomFields": [
        {			
            "Code": "string",				                // Type="String"
            "Label": "string",			                    // Type="String"
            "Value": "string"			        	        // Type="String"
        }
    ]
}
```

---
## Get a collection of all employee suspensions

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSuspension/all?$orderby={{$employee-suspension-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-suspension-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee suspensions**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSuspension",
    "value": [
        {
            "SuspensionId": 1,                                                  // Type="Int64"
            "EmployeeNumber": "string",                                         // Type="String"
            "FullName": "string",                                               // Type="String"
            "SuspensionReason": "string",                                       // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",                       // Type="DateTime"
            "ReturnDate": "2019-01-01T00:00:00+02:00",                          // Type="DateTime"
            "MinimumPeriods": 1.00000,                                          // Type="Decimal"
            "PercentageOfPay": 1.00000,                                         // Type="Decimal"
            "OverridingAmount": 1.00000,                                        // Type="Decimal"
            "SuspensionPaymentMethod": "string",                                // Type="String"
            "DoNotCalculateRecurringComponents": false,                         // Type="Boolean"
            "DenyOnceOffInput": false,                                          // Type="Boolean"
            "UI19DownloadUrl": "string",                                        // Type="String"
            "ExpectedReturnType": "string",                                     // Type="String"
            "FirstDateWorked": "2019-01-01T00:00:00+02:00",                     // Type="DateTime"
            "LastDateForWhichPaid": "2019-01-01T00:00:00+02:00",                // Type="DateTime"
            "UpdateDataSnapshot": false,                                        // Type="Boolean"
            "CustomFields": [		
                {			
                    "Code": "string",				                            // Type="String"
                    "Label": "string",			                                // Type="String"
                    "Value": "string"			        	                    // Type="String"
                }	
			]
        }
    ]
}
```

---
## Create a single suspension record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSuspension`

Create a single employee suspension record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeSuspension" in metadata endpoint for available fields
}
```

**Example Response: Create a single suspension record**

---
## Update a single suspension record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSuspension({{SuspensionId}})`

Update a single suspension record based on the specified `SuspensionId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeSuspension" in metadata endpoint for available fields
}
```

**Example Response: Update a single suspension record** (Status: 204 No Content)

---
## Delete a single suspension record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSuspension({{SuspensionId}})`

Retrieves suspension information based on the specified `SuspensionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single suspension record** (Status: 200 OK)

---

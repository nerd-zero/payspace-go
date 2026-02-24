# Employee > Payroll Processing

## Pay Rate Details

Retrieve employees' pay rate information associated with the `company_id`.

### Get a collection of pay rates

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayRate?$orderby={{$pay-rate-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee pay rate details based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$pay-rate-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of pay rates** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayRate",
    "@odata.count": 1,
    "value": [
        {
			"PayRateId": 0,                                 // Type="Int64"
			"EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String"
			"EffectiveDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
			"AutomaticPayInd": true,                        // Type="Boolean"
			"CompanyFrequency": "string",                   // Type="String"
			"PayFrequency": "per month",                    // Type="String"
			"HoursPerDay": 0,                               // Type="Double"
			"HoursPerMonth": 0,                             // Type="Double"
			"HoursPerWeek": 0,                              // Type="Double"
			"DaysPerPeriod": 0,                             // Type="Double"
			"WeeksPerMonth": 0,                             // Type="Double"
			"HourlyRate": 0,                                // Type="Decimal"
			"DailyRate": 0,                                 // Type="Decimal"
			"WeeklyRate": 0,                                // Type="Decimal"
			"MonthlyRate": 0,                               // Type="Decimal"
            "PercentageOfPreviousPackage": 0,               // Type="Double"
			"Package": 0,                                   // Type="Decimal"
			"HoursPerFortnight": 0,                         // Type="Double"
			"FortnightlyRate": 0,                           // Type="Decimal"
			"AnnualPackage": 0,                             // Type="Decimal"
			"ThirteenCheque": true,                         // Type="Boolean"
			"IsAnnual": true,                               // Type="Boolean"
			"Package2": 0,                                  // Type="Decimal"
			"Reason": "string",                             // Type="String"
			"Currency": "string",                           // Type="String"
			"OrganizationCategory": "string",               // Type="String"
			"WorkingDayMonday": true,                       // Type="Boolean"
			"WorkingDayTuesday": true,                      // Type="Boolean"
			"WorkingDayWednesday": true,                    // Type="Boolean"
			"WorkingDayThursday": true,                     // Type="Boolean"
			"WorkingDayFriday": true,                       // Type="Boolean"
			"WorkingDaySaturday": true,                     // Type="Boolean"
			"WorkingDaySunday": true,                       // Type="Boolean"
			"DateAdded": "2019-05-03T09:14:41.717Z",        // Type="DateTime"
			"Comments": "string",                           // Type="String"
			"DateCreated": "2019-05-03T09:14:41.717Z",       // Type="DateTime"
            "OneDollarInCurrency": 0,                       // Type="Double"
            "CustomFields": [		
                {			
                    "Code": "string",				        // Type="String"
                    "Label": "string",			        // Type="String"
                    "Value": "string
... (truncated)
```

---

### Get a single pay rate record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayRate({{PayrateId}})`

Retrieves employee pay rate information based on the specified `PayrateId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single pay rate record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayRate/$entity",
	"PayRateId": 0,                                 // Type="Int64"
	"EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String"
	"EffectiveDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
	"AutomaticPayInd": true,                        // Type="Boolean"
	"CompanyFrequency": "string",                   // Type="String"
	"PayFrequency": "per month",                    // Type="String"
	"HoursPerDay": 0,                               // Type="Double"
	"HoursPerMonth": 0,                             // Type="Double"
	"HoursPerWeek": 0,                              // Type="Double"
	"DaysPerPeriod": 0,                             // Type="Double"
	"WeeksPerMonth": 0,                             // Type="Double"
	"HourlyRate": 0,                                // Type="Decimal"
	"DailyRate": 0,                                 // Type="Decimal"
	"WeeklyRate": 0,                                // Type="Decimal"
	"MonthlyRate": 0,                               // Type="Decimal"
    "PercentageOfPreviousPackage": 0,               // Type="Double"
	"Package": 0,                                   // Type="Decimal"
	"HoursPerFortnight": 0,                         // Type="Double"
	"FortnightlyRate": 0,                           // Type="Decimal"
	"AnnualPackage": 0,                             // Type="Decimal"
	"ThirteenCheque": true,                         // Type="Boolean"
	"IsAnnual": true,                               // Type="Boolean"
	"Package2": 0,                                  // Type="Decimal"
	"Reason": "string",                             // Type="String"
	"Currency": "string",                           // Type="String"
	"OrganizationCategory": "string",               // Type="String"
	"WorkingDayMonday": true,                       // Type="Boolean"
	"WorkingDayTuesday": true,                      // Type="Boolean"
	"WorkingDayWednesday": true,                    // Type="Boolean"
	"WorkingDayThursday": true,                     // Type="Boolean"
	"WorkingDayFriday": true,                       // Type="Boolean"
	"WorkingDaySaturday": true,                     // Type="Boolean"
	"WorkingDaySunday": true,                       // Type="Boolean"
	"DateAdded": "2019-05-03T09:14:41.717Z",        // Type="DateTime"
	"Comments": "string",                           // Type="String"
	"DateCreated": "2019-05-03T09:14:41.717Z",       // Type="DateTime"
    "OneDollarInCurrency": 0,                       // Type="Double"
    "CustomFields": [		
        {			
            "Code": "string",				        // Type="String"
            "Label": "string",			        // Type="String"
            "Value": "string"			        	// Type="String"
        }	
    ]
}
```

---

### Get a collection of pay rates as of an effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayRate/effective/:effectivedate?$orderby={{$pay-rate-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}`

Retrieves a list of employee pay rate records as of the specified effective date and based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$pay-rate-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get a collection of pay rates by effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayRate",
    "@odata.count": 1,
    "value": [
        {
			"PayRateId": 0,                                 // Type="Int64"
			"EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String"
			"EffectiveDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
			"AutomaticPayInd": true,                        // Type="Boolean"
			"CompanyFrequency": "string",                   // Type="String"
			"PayFrequency": "per month",                    // Type="String"
			"HoursPerDay": 0,                               // Type="Double"
			"HoursPerMonth": 0,                             // Type="Double"
			"HoursPerWeek": 0,                              // Type="Double"
			"DaysPerPeriod": 0,                             // Type="Double"
			"WeeksPerMonth": 0,                             // Type="Double"
			"HourlyRate": 0,                                // Type="Decimal"
			"DailyRate": 0,                                 // Type="Decimal"
			"WeeklyRate": 0,                                // Type="Decimal"
			"MonthlyRate": 0,                               // Type="Decimal"
            "PercentageOfPreviousPackage": 0,               // Type="Double"
			"Package": 0,                                   // Type="Decimal"
			"HoursPerFortnight": 0,                         // Type="Double"
			"FortnightlyRate": 0,                           // Type="Decimal"
			"AnnualPackage": 0,                             // Type="Decimal"
			"ThirteenCheque": true,                         // Type="Boolean"
			"IsAnnual": true,                               // Type="Boolean"
			"Package2": 0,                                  // Type="Decimal"
			"Reason": "string",                             // Type="String"
			"Currency": "string",                           // Type="String"
			"OrganizationCategory": "string",               // Type="String"
			"WorkingDayMonday": true,                       // Type="Boolean"
			"WorkingDayTuesday": true,                      // Type="Boolean"
			"WorkingDayWednesday": true,                    // Type="Boolean"
			"WorkingDayThursday": true,                     // Type="Boolean"
			"WorkingDayFriday": true,                       // Type="Boolean"
			"WorkingDaySaturday": true,                     // Type="Boolean"
			"WorkingDaySunday": true,                       // Type="Boolean"
			"DateAdded": "2019-05-03T09:14:41.717Z",        // Type="DateTime"
			"Comments": "string",                           // Type="String"
			"DateCreated": "2019-05-03T09:14:41.717Z",       // Type="DateTime"
            "OneDollarInCurrency": 0,                       // Type="Double"
            "CustomFields": [		
                {			
                    "Code": "string",				        // Type="String"
                    "Label": "string",			        // Type="String"
                    "Value": "string
... (truncated)
```

---

### Get a collection of all employee payrates

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayRate/all?$orderby={{employee-payrate-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-payrate-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee payrates**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayRate",
    "value": [
        {
            "PayRateId": 1,                                         // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "EffectiveDate": "2019-01-01",                          // Type="DateTime"
            "AutomaticPayInd": false,                               // Type="Boolean"
            "CompanyFrequency": "string",                           // Type="String"
            "PayFrequency": "string",                               // Type="String"
            "HoursPerDay": 1.00,                                    // Type="Decimal"
            "HoursPerMonth": 1.00,                                  // Type="Decimal"
            "HoursPerWeek": 1.00,                                   // Type="Decimal"
            "DaysPerPeriod": 1.00,                                  // Type="Decimal"
            "WeeksPerMonth": 1.00,                                  // Type="Decimal"
            "HourlyRate": 1.00000,                                  // Type="Decimal"
            "DailyRate": 1.00000,                                   // Type="Decimal"
            "WeeklyRate": 1.00000,                                  // Type="Decimal"
            "MonthlyRate": 1.00000,                                 // Type="Decimal"
            "PercentageOfPreviousPackage": 1.00,                    // Type="Decimal"
            "Package": 1.00000,                                     // Type="Decimal"
            "HoursPerFortnight": 1.00000,                           // Type="Decimal"
            "FortnightlyRate": 1.00000,                             // Type="Decimal"
            "HoursPerSemiMonth": 1.00000,                           // Type="Decimal"
            "SemiMonthlyRate": 1.00000,                             // Type="Decimal"
            "AnnualPackage": 1.00000,                               // Type="Decimal"
            "ThirteenCheque": false,                                // Type="Boolean"
            "IsAnnual": false,                                      // Type="Boolean"
            "Package2": 1.00000,                                    // Type="Decimal"
            "Reason": "string",                                     // Type="String"
            "Currency": "string",                                   // Type="String"
            "OrganizationCategory": "string",                       // Type="String"
            "WorkingDayMonday": false,                              // Type="Boolean"
            "WorkingDayTuesday": false,                             // Type="Boolean"
            "WorkingDayWednesday": false,                           // Type="Boolean"
            "WorkingDayThursday": false,                            // Type="Boolean"
            
... (truncated)
```

---

### Delete a single pay rate record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayRate({{PayrateId}})`

Retrieves employee pay rate information based on the specified `PayrateId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single pay rate record** (Status: 200 OK)

```json
{}
```

---

### Create a single pay rate record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayRate`

Create a single employee pay rate record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeePayRate" in metadata endpoint for available fields
}
```

**Example Response: Create a single pay rate record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePayRate/$entity",
	"PayRateId": 0,                                 // Type="Int64"
	"EmployeeNumber": "string",                     // Type="String"
	"EffectiveDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
	"AutomaticPayInd": true,                        // Type="Boolean"
	"CompanyFrequency": "string",                   // Type="String"
	"PayFrequency": "per month",                    // Type="String"
	"HoursPerDay": 0,                               // Type="Double"
	"HoursPerMonth": 0,                             // Type="Double"
	"HoursPerWeek": 0,                              // Type="Double"
	"DaysPerPeriod": 0,                             // Type="Double"
	"WeeksPerMonth": 0,                             // Type="Double"
	"HourlyRate": 0,                                // Type="Decimal"
	"DailyRate": 0,                                 // Type="Decimal"
	"WeeklyRate": 0,                                // Type="Decimal"
	"MonthlyRate": 0,                               // Type="Decimal"
	"Package": 0,                                   // Type="Decimal"
	"HoursPerFortnight": 0,                         // Type="Double"
	"FortnightlyRate": 0,                           // Type="Decimal"
	"AnnualPackage": 0,                             // Type="Decimal"
	"ThirteenCheque": true,                         // Type="Boolean"
	"IsAnnual": true,                               // Type="Boolean"
	"Package2": 0,                                  // Type="Decimal"
	"Reason": "string",                             // Type="String"
	"Currency": "string",                           // Type="String"
	"OrganizationCategory": "string",               // Type="String"
	"WorkingDayMonday": true,                       // Type="Boolean"
	"WorkingDayTuesday": true,                      // Type="Boolean"
	"WorkingDayWednesday": true,                    // Type="Boolean"
	"WorkingDayThursday": true,                     // Type="Boolean"
	"WorkingDayFriday": true,                       // Type="Boolean"
	"WorkingDaySaturday": true,                     // Type="Boolean"
	"WorkingDaySunday": true,                       // Type="Boolean"
	"DateAdded": "2019-05-03T09:14:41.717Z",        // Type="DateTime"
	"Comments": "string",                           // Type="String"
	"DateCreated": "2019-05-03T09:14:41.717Z"       // Type="DateTime"
}
```

---

### Update a single pay rate record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePayRate({{PayrateId}})`

Update a single pay rate record based on the specified `PayrateId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeePayRate" in metadata endpoint for available fields
}
```

**Example Response: Update a single pay rate record** (Status: 204 No Content)

---
## Recurring Components

> NOTE: The `EmployeeNumber` `$filter` is required for Recurring components. eg. `EmployeeNumber eq '123'`

The `period` parameter can be found under the `Lookup Values` as the `Value` of CompanyRun.

The `frequency` parameter can be found under the `Lookup Values` as the `Value` of CompanyFrequency.

### Bonus Provision

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBonusProvision?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeBonusProvision` records based on the specified `run`, `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeBonusProvision",
    "@odata.count": 1,
    "value": [
       {
            "EmployeeBonusProvisionId": 0,                                      // Type="Int64"
            "IsBonusTaxSpread": true,                                           // Type="Boolean"
            "SpreadToDateOrTaxEnd": "Spread across the whole tax year",         // Type="String"   
            "BonusPayDateType": "Specify a day and month",                      // Type="String"   
            "BonusPayDate": "2014-12-15T00:00:00+02:00",                        // Type="DateTime"
            "OverrideBonusAccAmount": null,                                     // Type="Decimal"
            "AutomaticPay": true,                                               // Type="Boolean"
            "BonusPayOnTermination": true,                                      // Type="Boolean"
            "EmployeeNumber": "SKY1",                                           // Type="String"
            "FullName": "string",                                               // Type="String" 
            "ComponentCompany": "Automatic Bonus / Tax Spread",                 // Type="String"   
            "ComponentCompanyId": 0,                                            // Type="Int64" (read only)
            "ComponentIndicatorLine": null,                                     // Type="String"   
            "InPackage": false,                                                 // Type="Boolean"
            "BcoePercentage": null,                                             // Type="Decimal"
            "StartDate": null,                                                  // Type="DateTime"
            "EndDate": null,                                                    // Type="DateTime"
            "Comments": "",                                                     // Type="String"   
            "Values":  [
                {
                    "ComponentValueId": 0,                                      // Type="Int64"
                    "AliasDescription": "string",                               // Type="String"
                    "ComponentValue": 2000,                                     // Type="Decimal"
                    "ComponentEmployeeId": 0,                                   // Type="Int64"
                    "ValueType": "string"                                       // Type="String"
                }
            ]
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBonusProvision?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeBonusProvision` record for the specified `run`, `frequency` and `ComponentCompany`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeBonusProvision" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeBonusProvision",     
    "EmployeeBonusProvisionId": 0,                                      // Type="Int64"
    "IsBonusTaxSpread": true,                                           // Type="Boolean"
    "SpreadToDateOrTaxEnd": "Spread across the whole tax year",         // Type="String"   
    "BonusPayDateType": "Specify a day and month",                      // Type="String"   
    "BonusPayDate": "2014-12-15T00:00:00+02:00",                        // Type="DateTime"
    "OverrideBonusAccAmount": null,                                     // Type="Decimal"
    "AutomaticPay": true,                                               // Type="Boolean"
    "BonusPayOnTermination": true,                                      // Type="Boolean"
    "EmployeeNumber": "SKY1",                                           // Type="String"
    "FullName": "string",                                               // Type="String" 
    "ComponentCompany": "Automatic Bonus / Tax Spread",                 // Type="String"   
    "ComponentCompanyId": 544159,                                       // Type="Int64" (read only)
    "ComponentIndicatorLine": null,                                     // Type="String"   
    "InPackage": false,                                                 // Type="Boolean"
    "BcoePercentage": null,                                             // Type="Decimal"
    "StartDate": null,                                                  // Type="DateTime"
    "EndDate": null,                                                    // Type="DateTime"
    "Comments": "",                                                     // Type="String"   
    "Values":  [
        {
            "ComponentValueId": 0,                                      // Type="Int64"
            "AliasDescription": "string",                               // Type="String"
            "ComponentValue": 2000,                                     // Type="Decimal"
            "ComponentEmployeeId": 0,                                   // Type="Int64"
            "ValueType": "string"                                       // Type="String"
        }
    ]
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBonusProvision({{EmployeeBonusProvisionId}})?period={{run}}&frequency={{frequency}}`

Update a single `EmployeeBonusProvision` record based on the specified `run`, `frequency` and `EmployeeBonusProvisionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |

**Request Body:**

```json
{
  //see "EmployeeBonusProvision" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBonusProvision({{EmployeeBonusProvisionId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeBonusProvision` record for the specified `run`, `frequency` and `EmployeeBonusProvisionId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{}
```

---

### Company Car

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCompanyCarDetail?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeCompanyCarDetail` records based on the specified `run`, `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeCompanyCarDetail",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeCompanyCarId": 0,                              // Type="Int64"
            "CarDescription": "Audi",                               // Type="String" 
            "EffectiveDate": "2014-08-21T00:00:00+02:00",           // Type="DateTime"
            "PurchaseDate": "2014-08-21T00:00:00+02:00",            // Type="DateTime"
            "WasSubjectToMaintenancePlan": true,                    // Type="Boolean"
            "CarValue": 100000.00000,                               // Type="Decimal"
            "CarType": "Car",                                       // Type="String"
            "EmployerMaintenance": true,                            // Type="Boolean"
            "EmployerFuel": true,                                   // Type="Boolean"
            "CompensationAmount": null,                             // Type="Decimal"
            "CarPaymentType": null,                                 // Type="String" 
            "KenyaPurchasedCarOption": null,                        // Type="String" 
            "EmployeeBenefitPayments": null,                        // Type="Decimal"
            "CustomField1": null,                                   // Type="String" 
            "CustomField2": null,                                   // Type="String" 
            "CustomField3": null,                                   // Type="String" 
            "CompanyCarTax": "tax company car at 80% (default)",    // Type="String" 
            "Is80BusinessPurpose": null,                            // Type="Boolean"
            "OneHundredPercentTax": null,                           // Type="Boolean"
            "Inactive": null,                                       // Type="Boolean"
            "EmployeeNumber": "SKY1",                               // Type="String"
            "FullName": "string",                                   // Type="String"  
            "ComponentCompany": "Company Car Fringe Benefit",       // Type="String" 
            "ComponentCompanyId": 0,                                // Type="Int64" (READ ONLY)
            "ComponentIndicatorLine": null,                         // Type="String" 
            "InPackage": true,                                      // Type="Boolean"
            "BcoePercentage": 100.00000,                            // Type="Decimal"
            "StartDate": null,                                      // Type="DateTime"
            "EndDate": null,                                        // Type="DateTime"
            "Comments": "",                                         // Type="String" 
            "Values": [
                {
                    "ComponentValueId": 0,              // Type="Int64"
                    "AliasDescription": "string",       // Type="String"
                    "ComponentValue": 2000,             // Type="Deci
... (truncated)
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCompanyCarDetail?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeCompanyCarDetail` record for the specified `run`, `frequency` and `ComponentCompany`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeCompanyCarDetail" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeCompanyCarDetail", 
    "EmployeeCompanyCarId": 0,                              // Type="Int64"
    "CarDescription": "Audi",                               // Type="String" 
    "EffectiveDate": "2014-08-21T00:00:00+02:00",           // Type="DateTime"
    "PurchaseDate": "2014-08-21T00:00:00+02:00",            // Type="DateTime"
    "WasSubjectToMaintenancePlan": true,                    // Type="Boolean"
    "CarValue": 100000.00000,                               // Type="Decimal"
    "CarType": "Car",                                       // Type="String"
    "EmployerMaintenance": true,                            // Type="Boolean"
    "EmployerFuel": true,                                   // Type="Boolean"
    "CompensationAmount": null,                             // Type="Decimal"
    "CarPaymentType": null,                                 // Type="String" 
    "KenyaPurchasedCarOption": null,                        // Type="String" 
    "EmployeeBenefitPayments": null,                        // Type="Decimal"
    "CustomField1": null,                                   // Type="String" 
    "CustomField2": null,                                   // Type="String" 
    "CustomField3": null,                                   // Type="String" 
    "CompanyCarTax": "tax company car at 80% (default)",    // Type="String" 
    "Is80BusinessPurpose": null,                            // Type="Boolean"
    "OneHundredPercentTax": null,                           // Type="Boolean"
    "Inactive": null,                                       // Type="Boolean"
    "EmployeeNumber": "SKY1",                               // Type="String"
    "FullName": "string",                                   // Type="String" 
    "ComponentCompany": "Company Car Fringe Benefit",       // Type="String" 
    "ComponentCompanyId": 0,                                // Type="Int64" (READ ONLY)
    "ComponentIndicatorLine": null,                         // Type="String" 
    "InPackage": true,                                      // Type="Boolean"
    "BcoePercentage": 100.00000,                            // Type="Decimal"
    "StartDate": null,                                      // Type="DateTime"
    "EndDate": null,                                        // Type="DateTime"
    "Comments": "",                                         // Type="String" 
    "Values": [
        {
            "ComponentValueId": 0,              // Type="Int64"
            "AliasDescription": "string",       // Type="String"
            "ComponentValue": 2000,             // Type="Decimal"
            "ComponentEmployeeId": 0,           // Type="Int64"
            "ValueType": "string"               // Type="String"
        }
    ]
} 
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCompanyCarDetail({{EmployeeCompanyCarId}})?period={{run}}&frequency={{frequency}}`

Update a single `EmployeeCompanyCarDetail` record based on the specified `run`, `frequency` and `EmployeeCompanyCarId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |

**Request Body:**

```json
{
  //see "EmployeeCompanyCarDetail" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCompanyCarDetail({{EmployeeCompanyCarId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeCompanyCarDetail` record for the specified `run`, `frequency` and `EmployeeCompanyCarId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{}
```

---

### Disability

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDisability?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieve a list of active `EmployeeDisability` records for the specified `frequency`, `run` and `$filter`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeDisability",
     "value": [
        {
            "EmployeeDisabilityId": 0,                      // Type="Int64"
            "CompanyDisability": "string",                  // Type="String"
            "DisabilityComponentCompany": "string",         // Type="String"
            "EffectiveDate": "2013-10-01T00:00:00+02:00",   // Type="DateTime"
            "EmployerPercent": 0,                           // Type="Decimal"
            "EmployeePercent": 5.00000,                     // Type="Decimal"
            "EmployerExtraAdvised": 0,                      // Type="Decimal"
            "EmployeeExtraAdvised": 0,                      // Type="Decimal"
            "OverrideAmount": 0,                            // Type="Decimal"
            "EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String" 
            "ComponentCompany": "string",                   // Type="String"
            "ComponentCompanyId": 0,                        // Type="Int64"
            "ComponentIndicatorLine": null,                 // Type="String"
            "InPackage": true,                              // Type="Boolean"
            "BcoePercentage": 0,                            // Type="Decimal"
            "StartDate": null,                              // Type="DateTime"    
            "EndDate": null,                                // Type="DateTime"
            "Comments": "string",                           // Type="String"
            "Values": []
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDisability?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeDisability` record for the specified `frequency` and `run`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeDisability/$entity",
    "EmployeeDisabilityId": 0,                      // Type="Int64"
    "CompanyDisability": "string",                  // Type="String"
    "DisabilityComponentCompany": "string",         // Type="String"
    "EffectiveDate": "2013-10-01T00:00:00+02:00",   // Type="DateTime"
    "EmployerPercent": 0,                           // Type="Decimal"
    "EmployeePercent": 5.00000,                     // Type="Decimal"
    "EmployerExtraAdvised": 0,                      // Type="Decimal"
    "EmployeeExtraAdvised": 0,                      // Type="Decimal"
    "OverrideAmount": 0,                            // Type="Decimal"
    "EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String" 
    "ComponentCompany": "string",                   // Type="String"
    "ComponentCompanyId": 0,                        // Type="Int64"
    "ComponentIndicatorLine": null,                 // Type="String"
    "InPackage": true,                              // Type="Boolean"
    "BcoePercentage": 0,                            // Type="Decimal"
    "StartDate": null,                              // Type="DateTime"    
    "EndDate": null,                                // Type="DateTime"
    "Comments": "string",                           // Type="String"
    "Values": []
}
```

---

#### Updating Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDisability({{EmployeeDisabilityId}})?frequency={{frequency}}&period={{run}}`

Update an existing `EmployeeDisability` record for the specified `frequency` and `run`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Updating Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDisability({{EmployeeDisabilityId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeDisability` record for the specified `frequency` and `run`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

---

### Employee Components

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeComponent?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeComponent` records based on the specified `run`, `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeComponent",
    "@odata.count": 1,
    "value": [
        {
            "ComponentEmployeeId": 0,                   // Type="Int64"
            "EmployeeNumber": "string",                 // Type="String"
            "FullName": "string",                    // Type="String"  
            "ComponentCompany": "string",               // Type="String"
            "ComponentCompanyId": 0,                    // Type="Int64"
            "ComponentIndicatorLine": "string",         // Type="String"
            "InPackage": false,                         // Type="Boolean"
            "BcoePercentage": 0,                        // Type="Decimal"
            "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
            "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
            "Comments": "string",                       // Type="String"
            "Values": [
                {
                    "ComponentValueId": 0,              // Type="Int64"
                    "AliasDescription": "string",       // Type="String"
                    "ComponentValue": 2000,             // Type="Decimal"
                    "ComponentEmployeeId": 0,           // Type="Int64"
                    "ValueType": "string"               // Type="String"
                }
            ]
        }        
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeComponent?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeComponent` record for the specified `run`, `frequency` and `ComponentCompany`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeComponent" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeComponent", 
    "ComponentEmployeeId": 0,                   // Type="Int64"
    "EmployeeNumber": "string",                 // Type="String" 
    "FullName": "string",                       // Type="String"   
    "ComponentCompany": "string",               // Type="String"
    "ComponentCompanyId": 0,                    // Type="Int64"
    "ComponentIndicatorLine": "string",         // Type="String"
    "InPackage": false,                         // Type="Boolean"
    "BcoePercentage": 0,                        // Type="Decimal"
    "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
    "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
    "Comments": "string",                       // Type="String"
    "Values": [
        {
            "ComponentValueId": 0,              // Type="Int64"
            "AliasDescription": "string",       // Type="String"
            "ComponentValue": 0,             // Type="Decimal"
            "ComponentEmployeeId": 0,           // Type="Int64"
            "ValueType": "string"               // Type="String"
        }
    ]
} 
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeComponent({{ComponentEmployeeId}})?period={{run}}&frequency={{frequency}}`

Update a single `EmployeeComponent` record based on the specified `run`, `frequency` and `ComponentEmployeeId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |

**Request Body:**

```json
{
  //see "EmployeeComponent" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeComponent({{ComponentEmployeeId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeComponent` record for the specified `run`, `frequency` and `ComponentEmployeeId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{}
```

---

#### Extract all Employee Components: Start

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeComponent/extract/start?frequency={{frequency}}&period={{run}}`

Starts the process to create a file containing all `EmployeeComponent` records

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Extract all Employee Components: Start** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",             // Type="Guid"
    "ComponentName": "String",          // Type="String"    
    "CompanyId": 0,                     // Type="Int64"    
    "Frequency": "String",              // Type="String"    
    "Period": "String",                 // Type="String"    
    "StartTime": "2024-01-01T00:00:00", // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "Status": "Queued",                 // Type="String"    
    "Message": "String",                // Type="String"    
    "DownloadUrl": "download url"       // Type="String"    
}
```

---

#### Extract all Employee Components: Status

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeComponent/extract/:extractionId/status?frequency={{frequency}}&period={{run}}`

Returns the extraction process status with download url once completed

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Extract all Employee Components: Status**

```json
{
    "ExtractionId": "Guid",             // Type="Guid"
    "ComponentName": "String",          // Type="String"    
    "CompanyId": 0,                     // Type="Int64"    
    "Frequency": "String",              // Type="String"    
    "Period": "String",                 // Type="String"    
    "StartTime": "2024-01-01T00:00:00", // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "Status": "String",                 // Type="String"    
    "Message": "String",                // Type="String"    
    "DownloadUrl": "download url"       // Type="String"    
}
```

---

### Garnishees

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGarnishee?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeGarnishee` records based on the specified `run`, `frequency`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeGarnishee",
    "@odata.count": 1,
    "value": [
        {            
            "EmployeeGarnisheeId": 0,                   // Type="Int64"
            "CreditorName": "string",                   // Type="String"
            "BankName": "string",                       // Type="String"
            "BranchCode": "string",                     // Type="String"
            "AccountType": "string",                    // Type="String"
            "AccountNumber": "string",                  // Type="String"
            "Reference": "string",                      // Type="String"
            "Description": "string",                    // Type="String"
            "CapitalBalance": "string",                 // Type="String"
            "DeductionAmount": 0,                       // Type="Decimal"
            "IncreaseDecrease": 0,                      // Type="Decimal"
            "BalanceBroughtForward": 0,                 // Type="Decimal"
            "SubtractPercentage": 0,                    // Type="Decimal"
            "DebitAccountNumber": "string",             // Type="String"
            "CreditAccountNumber": "string",            // Type="String"
            "EmployeeNumber": "string",                 // Type="String"
            "FullName": "string",                       // Type="String" 
            "ComponentCompany": "string",               // Type="String"
            "ComponentCompanyId": 0,                    // Type="Int64"
            "ComponentIndicatorLine": null,             // Type="String"
            "InPackage": false,                         // Type="Boolean"
            "BcoePercentage": 0,                        // Type="Decimal"
            "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
            "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
            "Comments": "string",                       // Type="String"
            "Values": []
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGarnishee?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeGarnishee` record for the specified `run`, `frequency` and `ComponentCompany`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeGarnishee/$entity",
    "EmployeeGarnisheeId": 0,                   // Type="Int64"
    "CreditorName": "string",                   // Type="String"
    "BankName": "string",                       // Type="String"
    "BranchCode": "string",                     // Type="String"
    "AccountType": "string",                    // Type="String"
    "AccountNumber": "string",                  // Type="String"
    "Reference": "string",                      // Type="String"
    "Description": "string",                    // Type="String"
    "CapitalBalance": "string",                 // Type="String"
    "DeductionAmount": 0,                       // Type="Decimal"
    "IncreaseDecrease": 0,                      // Type="Decimal"
    "BalanceBroughtForward": 0,                 // Type="Decimal"
    "SubtractPercentage": 0,                    // Type="Decimal"
    "DebitAccountNumber": "string",             // Type="String"
    "CreditAccountNumber": "string",            // Type="String"
    "EmployeeNumber": "string",                 // Type="String"
    "FullName": "string",                       // Type="String" 
    "ComponentCompany": "string",               // Type="String"
    "ComponentCompanyId": 0,                    // Type="Int64"
    "ComponentIndicatorLine": null,             // Type="String"
    "InPackage": false,                         // Type="Boolean"
    "BcoePercentage": 0,                        // Type="Decimal"
    "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
    "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
    "Comments": "string",                       // Type="String"
    "Values": []
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGarnishee({{EmployeeGarnisheeId}})?frequency={{frequency}}&period={{run}}`

Update a single `EmployeeGarnishee` record based on the specified `run`, `frequency` and `EmployeeGarnisheeId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGarnishee({{EmployeeGarnisheeId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeGarnishee` record for the specified `run`, `frequency` and `EmployeeGarnisheeId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeGarnishee/$entity",
  "Success": true
}
```

---

### Group Life

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGroupLife?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeGroupLife` records based on the specified `run` and `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeGroupLife",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeGroupLifeId": 0,                                           // Type="Int64"
            "CompanyGroupLife": "GROUP LIFE A - Instance 1",                    // Type="String"   
            "GroupLifeComponentCompany": "Group Life Employer Contribution",    // Type="String"   
            "EffectiveDate": "2004-08-21T00:00:00+02:00",                       // Type="DateTime"
            "EmployeeExtraRecurring": null,                                     // Type="Decimal"
            "EmployerExtraRecurring": null,                                     // Type="Decimal"
            "EmployerPercent": null,                                            // Type="Decimal"
            "EmployeePercent": null,                                            // Type="Decimal"
            "OverrideAmount": null,                                             // Type="Decimal"
            "EmployeeNumber": "SKY1",                                           // Type="String"
            "FullName": "string",                                               // Type="String" 
            "ComponentCompany": "Company Car Fringe Benefit",                   // Type="String" 
            "ComponentCompanyId": 0,                                            // Type="Int64" (READ ONLY)
            "ComponentIndicatorLine": null,                                     // Type="String" 
            "InPackage": true,                                                  // Type="Boolean"
            "BcoePercentage": 100.00000,                                        // Type="Decimal"
            "StartDate": null,                                                  // Type="DateTime"
            "EndDate": null,                                                    // Type="DateTime"
            "Comments": "",                                                     // Type="String" 
            "Values": [
                {
                    "ComponentValueId": 0,              // Type="Int64"
                    "AliasDescription": "string",       // Type="String"
                    "ComponentValue": 2000,             // Type="Decimal"
                    "ComponentEmployeeId": 0,           // Type="Int64"
                    "ValueType": "string"               // Type="String"
                }
            ]
        }        
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGroupLife?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeGroupLife` record for the specified `run` and `frequency`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup |

**Request Body:**

```json
{
  //see "EmployeeGroupLife" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeGroupLife", 
    "EmployeeGroupLifeId": 0,                                           // Type="Int64"
    "CompanyGroupLife": "GROUP LIFE A - Instance 1",                    // Type="String"   
    "GroupLifeComponentCompany": "Group Life Employer Contribution",    // Type="String"   
    "EffectiveDate": "2004-08-21T00:00:00+02:00",                       // Type="DateTime"
    "EmployeeExtraRecurring": null,                                     // Type="Decimal"
    "EmployerExtraRecurring": null,                                     // Type="Decimal"
    "EmployerPercent": null,                                            // Type="Decimal"
    "EmployeePercent": null,                                            // Type="Decimal"
    "OverrideAmount": null,                                             // Type="Decimal"
    "EmployeeNumber": "SKY1",                                           // Type="String"
    "FullName": "string",                                               // Type="String" 
    "ComponentCompany": "Company Car Fringe Benefit",                   // Type="String" 
    "ComponentCompanyId": 0,                                            // Type="Int64" (READ ONLY)
    "ComponentIndicatorLine": null,                                     // Type="String" 
    "InPackage": true,                                                  // Type="Boolean"
    "BcoePercentage": 100.00000,                                        // Type="Decimal"
    "StartDate": null,                                                  // Type="DateTime"
    "EndDate": null,                                                    // Type="DateTime"
    "Comments": "",                                                     // Type="String" 
    "Values": [
        {
            "ComponentValueId": 0,              // Type="Int64"
            "AliasDescription": "string",       // Type="String"
            "ComponentValue": 2000,             // Type="Decimal"
            "ComponentEmployeeId": 0,           // Type="Int64"
            "ValueType": "string"               // Type="String"
        }
    ]
}  
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGroupLife({{EmployeeGroupLifeId}})?period={{run}}&frequency={{frequency}}`

Update a single `EmployeeGroupLife` record based on the specified `run`, `frequency` and `EmployeeGroupLifeId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |

**Request Body:**

```json
{
  //see "EmployeeGroupLife" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeGroupLife({{EmployeeGroupLifeId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeGroupLife` record for the specified `run`, `frequency` and `EmployeeGroupLifeId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{}
```

---

### House Payment

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeHousePayment?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeHousePayment` records based on the specified `run`, `frequency`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get a Collection**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeHousePayment",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeFinancialHousePaymentId": 0,               // Type="Int64"
            "Description": "string",                            // Type="String"
            "ReferenceNumber": "string",                        // Type="String"
            "Value": 0,                                         // Type="Decimal"
            "EffectiveDate": "2022-06-02T00:00:00+02:00",       // Type="DateTime"
            "EmployeeNumber": "string",                         // Type="String"
            "FullName": "string",                               // Type="String"
            "ComponentCompany": "string",                       // Type="String"
            "ComponentCompanyId": 0,                            // Type="Int64"
            "ComponentIndicatorLine": null,                     // Type="String"
            "InPackage": false,                                 // Type="Boolean"
            "BcoePercentage": 0,                                // Type="Decimal"
            "StartDate": null,                                  // Type="DateTime"
            "EndDate": null,                                    // Type="DateTime"
            "Comments": "string",                               // Type="String"
            "Values": []
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeHousePayment?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeHousePayment` record for the specified `run`, `frequency` and `ComponentCompany`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Create Record**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeHousePayment/$entity",
    "Description": "string",                            // Type="String"
    "ReferenceNumber": "string",                        // Type="String"
    "Value": 0,                                         // Type="Decimal"
    "EffectiveDate": "2022-06-02T00:00:00+02:00",       // Type="DateTime"
    "EmployeeNumber": "string",                         // Type="String"
    "FullName": "string",                               // Type="String"
    "ComponentCompany": "string",                       // Type="String"
    "ComponentCompanyId": 0,                            // Type="Int64"
    "ComponentIndicatorLine": null,                     // Type="String"
    "InPackage": false,                                 // Type="Boolean"
    "BcoePercentage": 0,                                // Type="Decimal"
    "StartDate": null,                                  // Type="DateTime"
    "EndDate": null,                                    // Type="DateTime"
    "Comments": "string"                                // Type="String"
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeHousePayment({{EmployeeFinancialHousePaymentId}})?frequency={{frequency}}&period={{run}}`

Update a single `EmployeeHousePayment` record based on the specified `run`, `frequency` and `EmployeeHousePaymentId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Update Existing Record**

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeHousePayment({{EmployeeFinancialHousePaymentId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeHousePayment` record for the specified `run`, `frequency` and `EmployeeHousePaymentId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record**

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeHousePayment/$entity",
  "Success": true
}
```

---

### Income Protection

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncomeProtection?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeIncomeProtection` records based on the specified `run`, `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeIncomeProtection",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeIncomeProtectionId": 0,                            // Type="Int64"
            "CompanyIncomeProtectionLink": "Test A - Instance 1",       // Type="String" 
            "EffectiveDate": "2013-10-01T00:00:00+02:00",               // Type="DateTime"
            "EmployerExtraRecurring": null,                             // Type="Decimal"
            "EmployeeExtraRecurring": null,                             // Type="Decimal"
            "EmployerPercent": null,                                    // Type="Decimal"
            "EmployeePercent": null,                                    // Type="Decimal"
            "OverrideAmount": null,                                     // Type="Decimal"
            "EmployeeNumber": "string",                                 // Type="String" 
            "FullName": "string",                                       // Type="String"   
            "ComponentCompany": "string",                               // Type="String"
            "ComponentCompanyId": 0,                                    // Type="Int64"
            "ComponentIndicatorLine": "string",                         // Type="String"
            "InPackage": false,                                         // Type="Boolean"
            "BcoePercentage": 0,                                        // Type="Decimal"
            "StartDate": "2019-05-03T09:14:41.717Z",                    // Type="DateTime"
            "EndDate": "2019-05-03T09:14:41.717Z",                      // Type="DateTime"
            "Comments": "string",                                       // Type="String"
            "Values": [
                {
                    "ComponentValueId": 0,              // Type="Int64"
                    "AliasDescription": "string",       // Type="String"
                    "ComponentValue": 2000,             // Type="Decimal"
                    "ComponentEmployeeId": 0,           // Type="Int64"
                    "ValueType": "string"               // Type="String"
                }
            ]
        }        
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncomeProtection?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeIncomeProtection` record for the specified `run`, `frequency` and `ComponentCompany`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeIncomeProtection" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeIncomeProtection", 
    "EmployeeIncomeProtectionId": 0,                            // Type="Int64"
    "CompanyIncomeProtectionLink": "Test A - Instance 1",       // Type="String" 
    "EffectiveDate": "2013-10-01T00:00:00+02:00",               // Type="DateTime"
    "EmployerExtraRecurring": null,                             // Type="Decimal"
    "EmployeeExtraRecurring": null,                             // Type="Decimal"
    "EmployerPercent": null,                                    // Type="Decimal"
    "EmployeePercent": null,                                    // Type="Decimal"
    "OverrideAmount": null,                                     // Type="Decimal"
    "EmployeeNumber": "string",                                 // Type="String"
    "FullName": "string",                                    // Type="String" 
    "ComponentCompany": "string",                               // Type="String"
    "ComponentCompanyId": 0,                                    // Type="Int64"
    "ComponentIndicatorLine": "string",                         // Type="String"
    "InPackage": false,                                         // Type="Boolean"
    "BcoePercentage": 0,                                        // Type="Decimal"
    "StartDate": "2019-05-03T09:14:41.717Z",                    // Type="DateTime"
    "EndDate": "2019-05-03T09:14:41.717Z",                      // Type="DateTime"
    "Comments": "string",                                       // Type="String"
    "Values": [
        {
            "ComponentValueId": 0,              // Type="Int64"
            "AliasDescription": "string",       // Type="String"
            "ComponentValue": 2000,             // Type="Decimal"
            "ComponentEmployeeId": 0,           // Type="Int64"
            "ValueType": "string"               // Type="String"
        }
    ]
}  
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncomeProtection({{EmployeeIncomeProtectionId}})?period={{run}}&frequency={{frequency}}`

Update a single `EmployeeIncomeProtection` record based on the specified `run`, `frequency` and `EmployeeIncomeProtectionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |

**Request Body:**

```json
{
  //see "EmployeeIncomeProtection" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncomeProtection({{EmployeeIncomeProtectionId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeIncomeProtection` record for the specified `run`, `frequency` and `EmployeeIncomeProtectionId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{}
```

---

### Loans

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLoan?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'BCA0003'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLoan",
    "value": [
        {
            "EmployeeLoanId": 0,                        // Type="Int64"
            "Customfield1": "string",                   // Type="String"
            "LoanAmount": 0.00000,                      // Type="Decimal"
            "BalanceBroughtForward": 0.00000,           // Type="Decimal"
            "PreviousBalance": null,                    // Type="Decimal"
            "InterestRate": 0.00000,                    // Type="Decimal"
            "InstallmentType": "Amount",                // Type="String"
            "LoanCompanyComponent": null,               // Type="String"
            "Installment": 0.00000,                     // Type="Decimal"
            "IncreaseDecrease": 0.00000,                // Type="Decimal"
            "NonFringeBenefitLoan": false,              // Type="Boolean"
            "DebitAccNo": "string",                     // Type="String"
            "CreditAccNo": "string",                    // Type="String"
            "PreviousInterestAmount": 0.00000,          // Type="Decimal"
            "OverrideOutstandingBalance": null,         // Type="Decimal"
            "EmployeeNumber": "string",                 // Type="String"
            "FullName": "string",                       // Type="String"
            "ComponentCompany": "string",               // Type="String"
            "ComponentCompanyId": 0,                    // Type="Int64"
            "ComponentIndicatorLine": null,             // Type="String"
            "InPackage": false,                         // Type="Boolean"
            "BcoePercentage": null,                     // Type="Decimal"
            "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
            "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
            "Comments": "string",                       // Type="String"
            "Values": [
                {
                    "ComponentValueId": 0,              // Type="Int64"
                    "AliasDescription": "string",       // Type="String"
                    "ComponentValue": 2000,             // Type="Decimal"
                    "ComponentEmployeeId": 0,           // Type="Int64"
                    "ValueType": "string"               // Type="String"
                }
            ]
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLoan?frequency={{frequency}}&period={{run}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Create Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeLoan",
    "EmployeeLoanId": 0,                        // Type="Int64"
    "Customfield1": "string",                   // Type="String"
    "LoanAmount": 0.00000,                      // Type="Decimal"
    "BalanceBroughtForward": 0.00000,           // Type="Decimal"
    "PreviousBalance": null,                    // Type="Decimal"
    "InterestRate": 0.00000,                    // Type="Decimal"
    "InstallmentType": "Amount",                // Type="String"
    "LoanCompanyComponent": null,               // Type="String"
    "Installment": 0.00000,                     // Type="Decimal"
    "IncreaseDecrease": 0.00000,                // Type="Decimal"
    "NonFringeBenefitLoan": false,              // Type="Boolean"
    "DebitAccNo": "string",                     // Type="String"
    "CreditAccNo": "string",                    // Type="String"
    "PreviousInterestAmount": 0.00000,          // Type="Decimal"
    "OverrideOutstandingBalance": null,         // Type="Decimal" - optional: used to specify a new Loan Balance, it returns the calculated IncreaseDecrease amount for the new Loan Balance
    "EmployeeNumber": "string",                 // Type="String"
    "FullName": "string",                       // Type="String"
    "ComponentCompany": "string",               // Type="String"
    "ComponentCompanyId": 0,                    // Type="Int64"
    "ComponentIndicatorLine": null,             // Type="String"
    "InPackage": false,                         // Type="Boolean"
    "BcoePercentage": null,                     // Type="Decimal"
    "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
    "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
    "Comments": "string",                       // Type="String"
    "Values": [
        {
            "ComponentValueId": 0,              // Type="Int64"
            "AliasDescription": "string",       // Type="String"
            "ComponentValue": 2000,             // Type="Decimal"
            "ComponentEmployeeId": 0,           // Type="Int64"
            "ValueType": "string"               // Type="String"
        }
    ]
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLoan({{EmployeeLoanId}})?frequency={{frequency}}&period={{run}}`

The `OverrideOutstandingBalance` field can be used to specify a new Loan Balance, it returns the calculated `IncreaseDecrease` amount for the new Loan Balance.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeLoan({{EmployeeLoanId}})?frequency={{frequency}}&period={{run}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 )

---

### Medical Aid

Retrieve employees' medical aid records

A medical aid component can be linked to multiple components (eg. Employee and Employer contributions), so you will need the `companyComponentId` as well as the `EmployeeMedicalId` in order to patch or delete the correct component.

The `componentCompanyId` value can be found on the `ComponentId` property in ComponentCompanyDetail lookup




#### Get a Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeMedical?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a collection of active `EmployeeMedical` records based on the specified `run`, `frequency`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get a Collection** (Status: 200 )

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeMedical",
    "value": [
        {
            "EmployeeMedicalId": 219012,                                            // Type="Int64"
            "CompanyMedicalAid": "Discovery - Classic Comprehensive",               // Type="String"
            "MedicalCategory": "Normal",                                            // Type="String"
            "EffectiveDate": "2018-09-01T00:00:00+02:00",                           // Type="DateTime"
            "TableOverride": false,                                                 // Type="Boolean"
            "PercentageAmount": "Percentage",                                       // Type="String"
            "TotalContribution": 3000.00000,                                        // Type="Decimal"
            "ReferenceNumber": "DCC001",                                            // Type="String"
            "IsRetired": false,                                                     // Type="Boolean"
            "Maximum": null,                                                        // Type="Decimal"
            "EmployeeAdvisedAmount": null,                                          // Type="Decimal"
            "EmployerAdvisedAmount": null,                                          // Type="Decimal"
            "EmployeePercentage": 50.00000,                                         // Type="Decimal"
            "EmployerPercentage": 50.00000,                                         // Type="Decimal"
            "EmployeeExtraAdvised": null,                                           // Type="Decimal"
            "EmployerExtraAdvised": null,                                           // Type="Decimal"
            "Advised": false,                                                       // Type="Boolean"
            "EmployeeNumber": "SKY1",                                               // Type="String"
            "FullName": "string",                                                   // Type="String" 
            "ComponentCompany": "Medical Aid Employee Deduction",                   // Type="String"
            "ComponentCompanyId": 6729,                                             // Type="Int64"
            "ComponentIndicatorLine": null,                                         // Type="String"
            "InPackage": false,                                                     // Type="Boolean"
            "BcoePercentage": null,                                                 // Type="Decimal"
            "StartDate": null,                                                      // Type="DateTime"
            "EndDate": null,                                                        // Type="DateTime"
            "Comments": "",                                                         // Type="String"
            "Values": [ ]
        }
    ]
}
```

---

#### Get Single Record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeMedical/:ComponentCompanyId/({{EmployeeMedicalId}})?frequency={{frequency}}&period={{run}}`

Retrieves an `EmployeeMedical` record based on the `componentCompanyId` and `employeeMedicalId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Get Single Record** (Status: 200 )

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeMedical/$entity",
    "EmployeeMedicalId": 171665,                                    // Type="Int64"
    "CompanyMedicalAid": "TEST - Advised SUBS",                     // Type="String"
    "MedicalCategory": "Normal",                                    // Type="String"
    "EffectiveDate": "2017-04-01T00:00:00+02:00",                   // Type="DateTime"
    "TableOverride": true,                                          // Type="Boolean"
    "PercentageAmount": "Amount",                                   // Type="String"
    "TotalContribution": 3844.33000,                                // Type="Decimal"
    "ReferenceNumber": "123",                                       // Type="String"
    "IsRetired": false,                                             // Type="Boolean"
    "Maximum": null,                                                // Type="Decimal"
    "EmployeeAdvisedAmount": 2500.00000,                            // Type="Decimal"
    "EmployerAdvisedAmount": 1344.33000,                            // Type="Decimal"
    "EmployeePercentage": null,                                     // Type="Decimal"
    "EmployerPercentage": null,                                     // Type="Decimal"
    "EmployeeExtraAdvised": null,                                   // Type="Decimal"
    "EmployerExtraAdvised": null,                                   // Type="Decimal"
    "Advised": false,                                               // Type="Boolean"
    "EmployeeNumber": "01112",                                      // Type="String"
    "FullName": "string",                                           // Type="String" 
    "ComponentCompany": "Medical Aid Employee Deduction",           // Type="String"
    "ComponentCompanyId": 6729,                                     // Type="Int64"
    "ComponentIndicatorLine": null,                                 // Type="String"
    "InPackage": false,                                             // Type="Boolean"
    "BcoePercentage": null,                                         // Type="Decimal"
    "StartDate": null,                                              // Type="DateTime"
    "EndDate": null,                                                // Type="DateTime"
    "Comments": "",                                                 // Type="String"
    "Values": [ ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeMedical?frequency={{frequency}}&period={{run}}`

Update an `EmployeeMedical` record based on the `employeeMedicalId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeMedical/$entity",
    "EmployeeMedicalId": 171665,                                    // Type="Int64"
    "CompanyMedicalAid": "TEST - Advised SUBS",                     // Type="String"
    "MedicalCategory": "Normal",                                    // Type="String"
    "EffectiveDate": "2017-04-01T00:00:00+02:00",                   // Type="DateTime"
    "TableOverride": true,                                          // Type="Boolean"
    "PercentageAmount": "Amount",                                   // Type="String"
    "TotalContribution": 3844.33000,                                // Type="Decimal"
    "ReferenceNumber": "123",                                       // Type="String"
    "IsRetired": false,                                             // Type="Boolean"
    "Maximum": null,                                                // Type="Decimal"
    "EmployeeAdvisedAmount": 2500.00000,                            // Type="Decimal"
    "EmployerAdvisedAmount": 1344.33000,                            // Type="Decimal"
    "EmployeePercentage": null,                                     // Type="Decimal"
    "EmployerPercentage": null,                                     // Type="Decimal"
    "EmployeeExtraAdvised": null,                                   // Type="Decimal"
    "EmployerExtraAdvised": null,                                   // Type="Decimal"
    "Advised": false,                                               // Type="Boolean"
    "EmployeeNumber": "01112",                                      // Type="String"
    "FullName": "string",                                           // Type="String" 
    "ComponentCompany": "Medical Aid Employee Deduction",           // Type="String"
    "ComponentCompanyId": 6729,                                     // Type="Int64"
    "ComponentIndicatorLine": null,                                 // Type="String"
    "InPackage": false,                                             // Type="Boolean"
    "BcoePercentage": null,                                         // Type="Decimal"
    "StartDate": null,                                              // Type="DateTime"
    "EndDate": null,                                                // Type="DateTime"
    "Comments": "",                                                 // Type="String"
    "Values": [ ]
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeMedical/:ComponentCompanyId/({{EmployeeMedicalId}})?frequency={{frequency}}&period={{run}}`

Update an `EmployeeMedical` record based on the `componentCompanyId` and `employeeMedicalId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeMedical/:ComponentCompanyId/({{EmployeeMedicalId}})?frequency={{frequency}}&period={{run}}`

Delete an `EmployeeMedical` record based on the `componentCompanyId` and `employeeMedicalId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 )

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeMedical/$entity",
  "Success": true
}
```

---

#### Extract all Records: Start

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeMedical/extract/start?frequency={{frequency}}&period={{run}}`

Starts the process to create a file containing all `EmployeeMedical` records

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Extract all Records: Start** (Status: 200 OK)

```json
{
    "ExtractionId": "Guid",             // Type="Guid"
    "ComponentName": "String",          // Type="String"    
    "CompanyId": 0,                     // Type="Int64"    
    "Frequency": "String",              // Type="String"    
    "Period": "String",                 // Type="String"    
    "StartTime": "2024-01-01T00:00:00", // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "Status": "Queued",                 // Type="String"    
    "Message": "String",                // Type="String"    
    "DownloadUrl": "download url"       // Type="String"    
}
```

---

#### Extract all Records: Status

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeMedical/extract/:extractionId/status?frequency={{frequency}}&period={{run}}`

Returns the extraction process status with download url once completed

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Extract all Records: Status**

```json
{
    "ExtractionId": "Guid",             // Type="Guid"
    "ComponentName": "String",          // Type="String"    
    "CompanyId": 0,                     // Type="Int64"    
    "Frequency": "String",              // Type="String"    
    "Period": "String",                 // Type="String"    
    "StartTime": "2024-01-01T00:00:00", // Type="Int64"    
    "EndTime": "2024-01-01T00:00:00",   // Type="Int64"    
    "Status": "String",                 // Type="String"    
    "Message": "String",                // Type="String"    
    "DownloadUrl": "download url"       // Type="String"    
}
```

---

### Pension

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePensionFund?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeePensionFund` records based on the specified `run`, `frequency`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePensionFund",
    "@odata.count": 1,
    "value": [
        {
            "EmployeePensionFundId": 0,                     // Type="Int64"
            "CompanyPensionFund": "string",                 // Type="String"
            "ComponentCompany": "string",                   // Type="String"
            "ComponentCompanyId": 0,                        // Type="Int64"
            "ComponentIndicatorLine": "string",             // Type="String"
            "CompanyPensionFundLink": "string",             // Type="String"
            "EffectiveDate": "2003-12-01T00:00:00+02:00",   // Type="DateTime"
            "EmployeeFixedAmount": 0,                       // Type="Decimal"
            "EmployerFixedAmount": 0,                       // Type="Decimal"
            "EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String"
            "OverrideAmount": null,                         // Type="Decimal"
            "ReferenceNumber": null,                        // Type="String"
            "InPackage": false,                             // Type="Boolean"
            "BcoePercentage": null,                         // Type="Decimal"
            "StartDate": "2003-12-01T00:00:00+02:00",       // Type="DateTime"
            "EndDate": "2003-12-01T00:00:00+02:00",         // Type="DateTime"
            "Comments": "string",                           // Type="String"
            "Values": []
        }
    ]    
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePensionFund?frequency={{frequency}}&period={{run}}`

Create a single `EmployeePensionFund` record for the specified `run`, `frequency` and `ComponentCompany`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePensionFund/$entity",
    "EmployeePensionFundId": 0,                     // Type="Int64"
    "CompanyPensionFund": "string",                 // Type="String"
    "ComponentCompany": "string",                   // Type="String"
    "ComponentCompanyId": 0,                        // Type="Int64"
    "ComponentIndicatorLine": "string",             // Type="String"
    "CompanyPensionFundLink": "string",             // Type="String"
    "EffectiveDate": "2003-12-01T00:00:00+02:00",   // Type="DateTime"
    "EmployeeFixedAmount": 0,                       // Type="Decimal"
    "EmployerFixedAmount": 0,                       // Type="Decimal"
    "EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String"
    "OverrideAmount": null,                         // Type="Decimal"
    "ReferenceNumber": null,                        // Type="String"
    "InPackage": false,                             // Type="Boolean"
    "BcoePercentage": null,                         // Type="Decimal"
    "StartDate": "2003-12-01T00:00:00+02:00",       // Type="DateTime"
    "EndDate": "2003-12-01T00:00:00+02:00",         // Type="DateTime"
    "Comments": "string",                           // Type="String"
    "Values": []   
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePensionFund({{EmployeePensionFundId}})?frequency={{frequency}}&period={{run}}`

Update a single `EmployeePensionFund` record based on the specified `run`, `frequency` and `EmployeePensionFundId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePensionFund({{EmployeePensionFundId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeePensionFund` record for the specified `run`, `frequency` and `EmployeePensionFundId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeePensionFund/$entity",
  "Success": true
}
```

---

### Retirement Annuity

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRetirementAnnuity?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeRetirementAnnuity` records based on the specified `run`, `frequency`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` |  |
| `period` | `{{run}}` |  |
| `$filter` | `{{$filter}}` |  |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRetirementAnnuity",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeRetirementAnnuityId": 0,               // Type="Int64"
            "EmployeeRetirementAnnuityName": "string",      // Type="String"
            "Reference": "string",                          // Type="String"
            "EffectiveDate": "2004-08-21T00:00:00+02:00",   // Type="DateTime"
            "Amount": 0,                                    // Type="Decimal"
            "EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String" 
            "ComponentCompany": "string",                   // Type="String"
            "ComponentCompanyId": 0,                        // Type="Int64"
            "ComponentIndicatorLine": null,                 // Type="String"
            "InPackage": false,                             // Type="Boolean"
            "BcoePercentage": null,                         // Type="Decimal"
            "StartDate": null,                              // Type="DateTime"
            "EndDate": null,                                // Type="DateTime"
            "Comments": "string",                           // Type="String"
            "Values": []
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRetirementAnnuity?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeRetirementAnnuity` record for the specified `run`, `frequency` and `ComponentCompany`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` |  |
| `period` | `{{run}}` |  |

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeRetirementAnnuity/$entity",
    "EmployeeRetirementAnnuityId": 0,               // Type="Int64"
    "EmployeeRetirementAnnuityName": "string",      // Type="String"
    "Reference": "string",                          // Type="String"
    "EffectiveDate": "2004-08-21T00:00:00+02:00",   // Type="DateTime"
    "Amount": 0,                                    // Type="Decimal"
    "EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String"
    "ComponentCompany": "string",                   // Type="String"
    "ComponentCompanyId": 0,                        // Type="Int64"
    "ComponentIndicatorLine": null,                 // Type="String"
    "InPackage": false,                             // Type="Boolean"
    "BcoePercentage": null,                         // Type="Decimal"
    "StartDate": null,                              // Type="DateTime"
    "EndDate": null,                                // Type="DateTime"
    "Comments": "string",                           // Type="String"
    "Values": []
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRetirementAnnuity({{EmployeeRetirementAnnuityId}})?frequency={{frequency}}&period={{run}}`

Update a single `EmployeeRetirementAnnuity` record based on the specified `run`, `frequency` and `EmployeeRetirementAnnuityId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` |  |
| `period` | `{{run}}` |  |

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRetirementAnnuity({{EmployeeRetirementAnnuityId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeRetiementAnnuity` record for the specified `run`, `frequency` and `EmployeeRetirementAnnuityId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` |  |
| `period` | `{{run}}` |  |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeRetirementAnnuity/$entity",
  "Success": true
}
```

---

### Saving

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSaving?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of `EmployeeSaving` records based on the specified `run`, `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSaving",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeSavingId": 0,                      // Type="Int64"
            "SavingAmount": 0.00000,                    // Type="Decimal"
            "SavingBalance": 0.00000,                   // Type="Decimal"
            "Installment": 0.00000,                     // Type="Decimal"
            "IncreaseDecrease": null,                   // Type="Decimal"
            "EmployeeNumber": "string",                 // Type="String"
            "FullName": "string",                       // Type="String"   
            "ComponentCompany": "string",               // Type="String"
            "ComponentCompanyId": 0,                    // Type="Int64"
            "ComponentIndicatorLine": "string",         // Type="String"
            "InPackage": false,                         // Type="Boolean"
            "BcoePercentage": 0,                        // Type="Decimal"
            "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
            "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
            "Comments": "string",                       // Type="String"
            "Values": [
                {
                    "ComponentValueId": 0,              // Type="Int64"
                    "AliasDescription": "string",       // Type="String"
                    "ComponentValue": 2000,             // Type="Decimal"
                    "ComponentEmployeeId": 0,           // Type="Int64"
                    "ValueType": "string"               // Type="String"
                }
            ]
        }        
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSaving?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeSaving` record for the specified `run`, `frequency` and `ComponentCompany`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeSaving" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSaving", 
    "EmployeeSavingId": 0,                      // Type="Int64"
    "SavingAmount": 0.00000,                    // Type="Decimal"
    "SavingBalance": 0.00000,                   // Type="Decimal"
    "Installment": 0.00000,                     // Type="Decimal"
    "IncreaseDecrease": null,                   // Type="Decimal"
    "EmployeeNumber": "string",                 // Type="String" 
    "FullName": "string",                       // Type="String"   
    "ComponentCompany": "string",               // Type="String"
    "ComponentCompanyId": 0,                    // Type="Int64"
    "ComponentIndicatorLine": "string",         // Type="String"
    "InPackage": false,                         // Type="Boolean"
    "BcoePercentage": 0,                        // Type="Decimal"
    "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
    "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
    "Comments": "string",                       // Type="String"
    "Values": [
        {
            "ComponentValueId": 0,              // Type="Int64"
            "AliasDescription": "string",       // Type="String"
            "ComponentValue": 2000,             // Type="Decimal"
            "ComponentEmployeeId": 0,           // Type="Int64"
            "ValueType": "string"               // Type="String"
        }
    ]
} 
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSaving({{EmployeeSavingId}})?period={{run}}&frequency={{frequency}}`

Update a single `EmployeeSaving` record based on the specified `run`, `frequency` and `EmployeeSavingId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |

**Request Body:**

```json
{
  //see "EmployeeSaving" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSaving({{EmployeeSavingId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeSaving` record for the specified `run`, `frequency` and `EmployeeSavingId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{}
```

---

### Travel Business Usage

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTravelBusinessUsage?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of `EmployeeTravelBusinessUsage` records based on the specified `run` and `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTravelBusinessUsage",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeTravelBusinessUsageId": 0,         // Type="Int64"
            "TravelBusinessUsageTaxPercentage": 0,      // Type="String" Enum TravelBusinessUsageTaxPercentage
            "EmployeeNumber": "string",                 // Type="String"  
            "FullName": "string",                       // Type="String"  
            "ComponentCompany": "string",               // Type="String"
            "ComponentCompanyId": 0,                    // Type="Int64"
            "ComponentIndicatorLine": "string",         // Type="String"
            "InPackage": false,                         // Type="Boolean"
            "BcoePercentage": 0,                        // Type="Decimal"
            "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
            "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
            "Comments": "string",                       // Type="String"
            "Values": [
                {
                    "ComponentValueId": 0,              // Type="Int64"
                    "AliasDescription": "string",       // Type="String"
                    "ComponentValue": 2000,             // Type="Decimal"
                    "ComponentEmployeeId": 0,           // Type="Int64"
                    "ValueType": "string"               // Type="String"
                }
            ]
        }        
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTravelBusinessUsage?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeTravelBusinessUsage` record for the specified `run` and `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeTravelBusinessUsage" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTravelBusinessUsage", 
    "EmployeeTravelBusinessUsageId": 0,         // Type="Int64"
    "TravelBusinessUsageTaxPercentage": 0,      // Type="String" Enum TravelBusinessUsageTaxPercentage
    "EmployeeNumber": "string",                 // Type="String"
    "FullName": "Jono xxxx",                    // Type="String"    
    "ComponentCompany": "string",               // Type="String"
    "ComponentCompanyId": 0,                    // Type="Int64"
    "ComponentIndicatorLine": "string",         // Type="String"
    "InPackage": false,                         // Type="Boolean"
    "BcoePercentage": 0,                        // Type="Decimal"
    "StartDate": "2019-05-03T09:14:41.717Z",    // Type="DateTime"
    "EndDate": "2019-05-03T09:14:41.717Z",      // Type="DateTime"
    "Comments": "string",                       // Type="String"
    "Values": [
        {
            "ComponentValueId": 0,              // Type="Int64"
            "AliasDescription": "string",       // Type="String"
            "ComponentValue": 2000,             // Type="Decimal"
            "ComponentEmployeeId": 0,           // Type="Int64"
            "ValueType": "string"               // Type="String"
        }
    ]
}  
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTravelBusinessUsage({{EmployeeTravelBusinessUsageId}})?period={{run}}&frequency={{frequency}}`

Update a single `EmployeeTravelBusinessUsage` record based on the specified `run`, `frequency` and `EmployeeTravelBusinessUsageId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |

**Request Body:**

```json
{
  //see "EmployeeTravelBusinessUsage" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTravelBusinessUsage({{EmployeeTravelBusinessUsageId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeTravelBusinessUsage` record for the specified `run`, `frequency` and `EmployeeTravelBusinessUsageId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{}
```

---

### Table Builder

1. For a list of all the `TableBuilder` components use `TableBuilderComponentCompany` lookup.
2. For a list of custom fields use `TableBuilderCategory` lookup and filter the value by the `TableBuilderCategory` value in  `TableBuilderComponentCompany` lookup.
3. For `TableBuilderCode` options use `TableBuilderCode` lookup and use `TableBuilderComponentCompany` value in the lookup url.

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTableBuilder?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'BCA0003'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTableBuilder",
    "value": [
         {
            "EmployeeTableBuilderId": 0,                    // Type="Int64"
            "EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String"
            "TableBuilderCode": "string",                   // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",   // Type="DateTime"
            "ComponentCompany": "string",                   // Type="String"
            "ComponentCompanyId": 0,                        // Type="Int64"
            "ComponentIndicatorLine": "string",             // Type="String"
            "InPackage": false,                             // Type="Boolean"
            "BcoePercentage": null,                         // Type="Decimal"
            "StartDate": "2019-01-01T00:00:00+02:00",       // Type="DateTime"
            "EndDate": "2019-01-01T00:00:00+02:00",         // Type="DateTime"
            "Comments": "string",                           // Type="String"
            "CustomFields": [
               {			
                    "Code": "string",				        // Type="String"
                    "Label": "string",			            // Type="String"
                    "Value": "string"			        	// Type="String"
                }	
            ]
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTableBuilder?frequency={{frequency}}&period={{run}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeTableBuilder" in metadata endpoint for available fields
}
```

**Example Response: Create Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTableBuilder",
    "EmployeeTableBuilderId": 0,                    // Type="Int64"
    "EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String"
    "TableBuilderCode": "string",                   // Type="String"
    "EffectiveDate": "2019-01-01T00:00:00+02:00",   // Type="DateTime"
    "ComponentCompany": "string",                   // Type="String"
    "ComponentCompanyId": 0,                        // Type="Int64"
    "ComponentIndicatorLine": "string",             // Type="String"
    "InPackage": false,                             // Type="Boolean"
    "BcoePercentage": null,                         // Type="Decimal"
    "StartDate": "2019-01-01T00:00:00+02:00",       // Type="DateTime"
    "EndDate": "2019-01-01T00:00:00+02:00",         // Type="DateTime"
    "Comments": "string",                           // Type="String"
    "CustomFields": [
        {			
            "Code": "string",				        // Type="String"
            "Label": "string",			            // Type="String"
            "Value": "string"			        	// Type="String"
        }	
    ]
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTableBuilder({{EmployeeTableBuilderId}})?frequency={{frequency}}&period={{run}}`

The `OverrideOutstandingBalance` field can be used to specify a new Loan Balance, it returns the calculated `IncreaseDecrease` amount for the new Loan Balance.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Request Body:**

```json
{
  //see "EmployeeTableBuilder" in metadata endpoint for available fields
}
```

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTableBuilder({{EmployeeTableBuilderId}})?frequency={{frequency}}&period={{run}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record**

---

### Unions

#### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeUnion?frequency={{frequency}}&period={{run}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeUnion` records based on the specified `run`, `frequency`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeUnion",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeUnionId": 0,                           // Type="Int64"
            "Union": "string",                              // Type="String"
            "EffectiveDate": "2020-01-01T00:00:00+02:00",   // Type="DateTime"
            "EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String" 
            "ComponentCompany": "string",                   // Type="String"
            "ComponentCompanyId": 0,                        // Type="Int64"
            "ComponentIndicatorLine": null,                 // Type="String"
            "InPackage": false,                             // Type="Boolean"
            "BcoePercentage": null,                         // Type="Decimal"
            "StartDate": null,                              // Type="DateTime"
            "EndDate": null,                                // Type="DateTime"
            "Comments": "string",                           // Type="String"
            "Values": [
                {
                    "ComponentValueId": 0,                  // Type="Int64"
                    "Description": "string",                // Type="String"
                    "ComponentValue": 0,                    // Type="Decimal"
                    "ComponentEmployeeId": 0,               // Type="Int64"
                    "ValueType": "string"                   // Type="String"
                }
            ]
        }
    ]
}
```

---

#### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeUnion?frequency={{frequency}}&period={{run}}`

Create a single `EmployeeUnion` record for the specified `run`, `frequency` and `ComponentCompany`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeUnion/$entity",
    "EmployeeUnionId": 0,                           // Type="Int64"
    "Union": "string",                              // Type="String"
    "EffectiveDate": "2020-01-01T00:00:00+02:00",   // Type="DateTime"
    "EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String" 
    "ComponentCompany": "string",                   // Type="String"
    "ComponentCompanyId": 0,                        // Type="Int64"
    "ComponentIndicatorLine": null,                 // Type="String"
    "InPackage": false,                             // Type="Boolean"
    "BcoePercentage": null,                         // Type="Decimal"
    "StartDate": null,                              // Type="DateTime"
    "EndDate": null,                                // Type="DateTime"
    "Comments": "string",                           // Type="String"
    "Values": [
        {
            "ComponentValueId": 0,                  // Type="Int64"
            "Description": "string",                // Type="String"
            "ComponentValue": 0,                    // Type="Decimal"
            "ComponentEmployeeId": 0,               // Type="Int64"
            "ValueType": "string"                   // Type="String"
        }
    ]
}
```

---

#### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeUnion({{EmployeeUnionId}})?frequency={{frequency}}&period={{run}}`

Update a single `EmployeeUnion` record based on the specified `run`, `frequency` and `EmployeeUnionId`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Update Existing Record** (Status: 204 No Content)

---

#### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeUnion({{EmployeeUnionId}})?frequency={{frequency}}&period={{run}}`

Delete a single `EmployeeUnion` record for the specified `run`, `frequency` and `EmployeeUnionId`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | `value` from `CompanyFrequency` lookup required |
| `period` | `{{run}}` | `value` from `CompanyRun` lookup required |

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeUnion/$entity",
  "Success": true
}
```

---
## Take On Year To Date Figures

Retrieve employees' take on year to date information associated with the `company_id` and `frequency`.

### Get a collection of take on records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTakeOn?frequency={{frequency}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of active employee's take on year to date records, for the latest take on run in the specified `company_id` and `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Required `value` from `CompanyFrequency` lookup |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of take on records** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTakeOn",
    "@odata.count": 1,
    "value": [
        {
            "Identifier": "string",                     // Type="String"
            "EmployeeNumber": "string",                 // Type="String"
            "EmployeeTakeOnComponentCompany": "string", // Type="String" see EmployeeTakeOnComponentCompany lookup
            "PayslipAction": "string",                  // Type="String"
            "TaxCode": "string",                        // Type="String"
            "ComponentValue": 0.00000                   // Type="Decimal"
        }
    ]
}
```

---

### Get a single take on record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTakeOn({{Identifier}})?frequency={{frequency}}`

Retrieves employee's take on year to date record based on the specified `Identifier` and `frequency`.


**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Required `value` from `CompanyFrequency` lookup |

**Example Response: Get a single take on record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTakeOn/$entity",
    "Identifier": "string",                     // Type="String"
    "EmployeeNumber": "string",                 // Type="String"
    "EmployeeTakeOnComponentCompany": "string", // Type="String" see EmployeeTakeOnComponentCompany lookup
    "PayslipAction": "string",                  // Type="String"
    "TaxCode": "string",                        // Type="String"
    "ComponentValue": 0.00000                   // Type="Decimal"
}
```

---

### Create a single take on record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTakeOn?frequency={{frequency}}`

Create a single employee's take on year to date record for the specified `EmployeeNumber` , `EmployeeTakeOnComponentCompany` and `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Required `value` from `CompanyFrequency` lookup |

**Request Body:**

```json
{
  //see "EmployeeTakeOn" in metadata endpoint for available fields
}
```

**Example Response: Create a single take on record** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTakeOn/$entity",
    "Identifier": "string",                     // Type="String"
    "EmployeeNumber": "string",                 // Type="String"
    "EmployeeTakeOnComponentCompany": "string", // Type="String" see EmployeeTakeOnComponentCompany lookup
    "PayslipAction": "string",                  // Type="String"
    "TaxCode": "string",                        // Type="String"
    "ComponentValue": 0.00000                   // Type="Decimal"
}
```

---

### Update a single take on record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTakeOn({{Identifier}})?frequency={{frequency}}`

Update a single employee's take on year to date record for the specified `Identifier` and `frequency`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Required `value` from `CompanyFrequency` lookup |

**Request Body:**

```json
{
  //see "EmployeeTakeOn" in metadata endpoint for available fields
}
```

**Example Response: Update a single take on record** (Status: 204 No Content)

---
## Recurring Templates

Apply Recurring Templates to employees.

### Add a single Recurring Template to an employee

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeRecurringTemplate?frequency={{frequency}}&period={{run}}`

Apply Recurring Templates to employees.

> Note: When `ReplaceExisting` is set to true, all of the employees current components will be removed before the template is applied.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` |  |
| `period` | `{{run}}` |  |

**Request Body:**

```json
{
  //see "EmployeeRecurringTemplate" in metadata endpoint for available fields
}
```

**Example Response: Add a single Recurring Template to an employee**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeBonusProvision",     
    "EmployeeNumber": "SKY1",                                           // Type="String"   
    "TemplateName": Bonus Prov",                                        // Type="String"
    "ReplaceExisting": false                                            // Type="Boolean"
}
```

---
## Claims

Retrieve employees' claim information associated with the `company_id`.

For claim workflow examples have a look at the collection [here](https://documenter.getpostman.com/view/7757277/2s83ziPPVh).

### Get a collection of claims

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim?$orderby={{$claim-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee claim details based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$claim-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of claim** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeClaim",
    "@odata.count": 1,
    "value": [
        {
            "ClaimDetailId": 0,                                     // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"                                             
            "ClaimBatchId": 0,                                      // Type="Int64"
            "Status": "string",                                     // Type="String"
            "DateCaptured": "2022-12-17T08:26:14.633+02:00",        // Type="DateTime"
            "ClaimComponent": "string",                             // Type="String"
            "OrganizationGroup": "string",                          // Type="String"
            "ProjectCode": "string",                                // Type="String"
            "ActivityCode": "string",                               // Type="String"
            "CompanyRun": "20134",                                  // Type="String"
            "Date": "2022-12-17T08:26:14.633+02:00",                // Type="DateTime"
            "Description": "string",                                // Type="String"
            "Rate": 0,                                              // Type="Decimal"
            "Quantity": 0,                                          // Type="Decimal"
            "TotalAmount": 0,                                       // Type="Decimal"
            "OpeningOdo": 0,                                        // Type="Int32"
            "ClosingOdo": 0,                                        // Type="Int32"
            "PrivateTravel": 0,                                     // Type="Int32"
            "RegistrationNumber": "string",                         // Type="String"
            "DepartureDate": "2022-12-17T08:26:14.633+02:00",       // Type="DateTime"
            "ReturnDate": "2022-12-18T08:26:14.633+02:00",          // Type="DateTime"
            "Destination": "string",                                // Type="String"
            "AttachmentUrls": ["string", "string"],			        // Type="Collection(String)"
            "CustomFields": [
                {
                    "Code": "string",                               // Type="String"
                    "Label": "string",                              // Type="String"
                    "Value": "string"                               // Type="String"
                }
            ]
        }
    ]
}
```

---

### Get a single employee claim record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim({{EmployeeClaimId}})`

Retrieves employee claim information based on the specified `EmployeeClaimId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single employee claim record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeClaim/$entity",
	"EmployeeClaimId": 0,                                   // Type="Int64"
    "EmployeeNumber": "string",                             // Type="String"
    "FullName": "string",                                   // Type="String"                                             
    "ClaimBatchId": 0,                                      // Type="Int64"
    "Status": "string",                                     // Type="String"
    "DateCaptured": "2022-12-17T08:26:14.633+02:00",        // Type="DateTime"
    "ClaimComponent": "string",                             // Type="String"
    "OrganizationGroup": "string",                          // Type="String"
    "ProjectCode": "string",                                // Type="String"
    "ActivityCode": "string",                               // Type="String"
    "CompanyRun": "20134",                                  // Type="String"
    "Date": "2022-12-17T08:26:14.633+02:00",                // Type="DateTime"
    "Description": "string",                                // Type="String"
    "Rate": 0,                                              // Type="Decimal"
    "Quantity": 0,                                          // Type="Decimal"
    "TotalAmount": 0,                                       // Type="Decimal"
    "OpeningOdo": 0,                                        // Type="Int32"
    "ClosingOdo": 0,                                        // Type="Int32"
    "PrivateTravel": 0,                                     // Type="Int32"
    "RegistrationNumber": "string",                         // Type="String"
    "DepartureDate": "2022-12-17T08:26:14.633+02:00",       // Type="DateTime"
    "ReturnDate": "2022-12-18T08:26:14.633+02:00",          // Type="DateTime"
    "Destination": "string",                                // Type="String"
    "AttachmentUrls": [                                     // Type="Collection(String)"
        "string"                                            // Type="String"
    ],
    "CustomFields": [
        {
            "Code": "string",                               // Type="String"
            "Label": "string",                              // Type="String"
            "Value": "string"                               // Type="String"
        }
    ]
}
```

---

### Get a collection of all employee claims

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim/all?$orderby={{employee-claim-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-claim-field}}` |  |
| `$top` | `{{$top}}` |  |
| `$skip` | `{{$skip}}` |  |
| `$count` | `{{$count}}` |  |
| `$filter` | `{{$filter}}` |  |

**Example Response: Get a collection of all employee claims**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeClaim",
    "value": [
        {
            "EmployeeClaimId": 1,                                   // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "ClaimBatchId": 1,                                      // Type="Int64"
            "Status": "string",                                     // Type="String"
            "DateCaptured": "2019-01-01T00:00:00+02:00",            // Type="DateTime"
            "ClaimComponent": "string",                             // Type="String"
            "OrganizationGroup": "string",                          // Type="String"
            "ProjectCode": "string",                                // Type="String"
            "ActivityCode": "string",                               // Type="String"
            "CompanyRun": "string",                                 // Type="String"
            "Date": "2019-01-01T00:00:00+02:00",                    // Type="DateTime"
            "Description": "string",                                // Type="String"
            "Rate": 1.00000,                                        // Type="Decimal"
            "Quantity": 1.00000,                                    // Type="Decimal"
            "TotalAmount": 1.00000,                                 // Type="Decimal"
            "OpeningOdo": 0,                                        // Type="Int64"
            "ClosingOdo": 0,                                        // Type="Int64"
            "PrivateTravel": 0,                                     // Type="Int64"
            "RegistrationNumber": "string",                         // Type="String"
            "DepartureDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "ReturnDate": "2019-01-01T00:00:00+02:00",              // Type="DateTime"
            "Destination": "string",                                // Type="String"
            "AttachmentUrls": [],
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

### Create a single employee claim record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim`

Create a single employee claim record for the specified `EmployeeNumber`.

> NOTE: If you specify a `ClaimBatchId`, this `EmployeeClaim` will be linked to specified claim batch. This will only work with batches with a "To Be Finalized" status. `CompanyRun` is optional depending on company settings.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeClaim" in metadata endpoint for available fields
}
```

**Example Response: Create a single employee claim record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeClaim/$entity",
	"EmployeeClaimId": 0,                                   // Type="Int64"
    "EmployeeNumber": "string",                             // Type="String"
    "FullName": "string",                                   // Type="String"                                             
    "ClaimBatchId": 0,                                      // Type="Int64"
    "Status": "string",                                     // Type="String"
    "DateCaptured": "2022-12-17T08:26:14.633+02:00",        // Type="DateTime"
    "ClaimComponent": "string",                             // Type="String"
    "OrganizationGroup": "string",                          // Type="String"
    "ProjectCode": "string",                                // Type="String"
    "ActivityCode": "string",                               // Type="String"
    "CompanyRun": "20134",                                  // Type="String"
    "Date": "2022-12-17T08:26:14.633+02:00",                // Type="DateTime"
    "Description": "string",                                // Type="String"
    "Rate": 0,                                              // Type="Decimal"
    "Quantity": 0,                                          // Type="Decimal"
    "TotalAmount": 0,                                       // Type="Decimal"
    "OpeningOdo": 0,                                        // Type="Int32"
    "ClosingOdo": 0,                                        // Type="Int32"
    "PrivateTravel": 0,                                     // Type="Int32"
    "RegistrationNumber": "string",                         // Type="String"
    "DepartureDate": "2022-12-17T08:26:14.633+02:00",       // Type="DateTime"
    "ReturnDate": "2022-12-18T08:26:14.633+02:00",          // Type="DateTime"
    "Destination": "string",                                // Type="String"
    "AttachmentUrls": ["string", "string"],			        // Type="Collection(String)"
    "CustomFields": [
        {
            "Code": "string",                               // Type="String"
            "Label": "string",                              // Type="String"
            "Value": "string"                               // Type="String"
        }
    ]
}
```

---

### Upload claim attachment

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim/:employee-claim-id/attachment/upload`

Uploads an attachment for the specified `EmployeeClaimId`. These attachments are shared across all employee claims within the same batch that share the same `ClaimComponent`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body (form-data):**

| Key | Type | Value | Description |
|-----|------|-------|-------------|
| `File` | file | `` | Key name "File" As "File" type |

**Example Response: Upload claim attachment** (Status: 201 Created)

```json
{}
```

---

### Update a single employee claim record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim({{EmployeeClaimId}})`

Update a single employee claim record based on the specified `EmployeeClaimId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeClaim" in metadata endpoint for available fields
}
```

**Example Response: Update a single employee claim record** (Status: 204 No Content)

---

### Delete a single employee claim record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim({{EmployeeClaimId}})`

Delete a single employee claim record based on the specified `EmployeeClaimId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single employee claim record** (Status: 200 OK)

---

### Submit employee claim batch for workflow

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeClaim/:claim-batch-id/submit`

Submits all employee claim records to workflow for the specified `ClaimBatchId`.

> NOTE: The request for approval will be sent to the `EmployeeInbox` of each employee listed as `Approvers` provided in the response.

The response will include the `EmployeeWorkflow` that will provide the `UserWorkflowStepId` to advance the workflow. This value can also be retrieved from the `EmployeeInbox` endpoint.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Submit employee claim batch for workflow** (Status: 201 Created)

```json
{}
```

---
## Claim Employee Workflow

### Get a single employee workflow record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeWorkflow/claim/:user-workflow-step-id/:claim-batch-id`

Retrieves employee workflow information based on the specified `UserWorkflowStepId` and `ClaimBatchId`.

> NOTE: View `EmployeeInbox` endpoint to get the `UserWorkflowStepId` used for this request.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single employee workflow record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-Id}}/$metadata#EmployeeWorkflow/$entity",
    "UserWorkflowStepId": 0,                             // Type="Int64" - Represents current user workflow step id.
    "WorkflowComplete": false,                           // Type="Boolean" - Represents whether or not the workflow is complete.
    "PreviousApprovers": [                               // Type="Collection(String)"
        "string"                                         // Type="String" - Employee Number
    ],
    "Approvers": [                                       // Type="Collection(String)"
        "string"                                         // Type="String" - Employee Number
    ],
    "NextApprovers": [                                   // Type="Collection(String)"
        "string",                                        // Type="String" - Employee Number
        "string"                                         // Type="String" - Employee Number
    ],
    "ActionDate": "2023-03-16T07:29:51.813+02:00"         // Type="DateTime"
}
```

---

### Submit Workflow Step

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeWorkflow/claim/:user-workflow-step-id/submit`

Processes `WorkflowAction` for specified `UserWorkflowStepId`.

> NOTE: `ApproverEmployeeNumber` must be selected from the employee numbers provided on the `Approvers` property when requesting information on this workflow step.
> 
> `SelectedNextApprover` must be selected from the employee numbers provided on the `NextApprovers` property when requesting information on this workflow step.
> 
> When calling the this endpoint after the step is processed, then you will get a No Workflow Found response as the step you are attempting to manipulate has already been processed.

Depending on company settings, you may be required to specify `CompanyRun` and `SelectedNextApprover`.

`Comments` are always required upon selection of the `Reject` `WorkflowAction` upon submit.

Upon a successful processing of the final step of a workflow, the response will include property `WorkflowComplete` as true.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  "ClaimBatchId": 0,
  "ApproverEmployeeNumber": "string",
  "WorkflowAction": "string",
  "CompanyRun": "string",
  "SelectedNextApprover": "string",
  "Comments": ""
}
```

**Example Response: Submit Workflow Step** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-Id}}/$metadata#EmployeeWorkflow/$entity",
    "UserWorkflowStepId": 0,                             // Type="Int64" - Represents current user workflow step id.
    "WorkflowComplete": false,                           // Type="Boolean" - Represents whether or not the workflow is complete.
    "PreviousApprovers": [                               // Type="Collection(String)"
        "string"                                         // Type="String" - Employee Number
    ],
    "Approvers": [                                       // Type="Collection(String)"
        "string"                                         // Type="String" - Employee Number
    ],
    "NextApprovers": [                                   // Type="Collection(String)"
        "string",                                        // Type="String" - Employee Number
        "string"                                         // Type="String" - Employee Number
    ],
    "ActionDate": "2023-03-16T07:29:51.813+02:00"         // Type="DateTime"
}
```

---

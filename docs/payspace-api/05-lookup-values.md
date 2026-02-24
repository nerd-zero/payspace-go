# Lookup Values

Contains several endpoints related to lookup information.

## Custom Fields

### Custom Fields

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomField/:entityTypeName?$top=100&$skip=0&$count=true`

Lookup custom fields for a `CompanyId` and `Entity type name` eg. `Employee`, `EmployeeAsset`, `EmployeeDependant`. Use the entity type name as it is in the `$metadata`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Custom fields for a type** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFieldDetail",
    "@odata.count": 1,
    "value": [       
         {
            "Code": "string",			// Type="String"
            "Label": "string",			// Type="String"
            "Type": "string",			// Type="String"
            "IsRequired": false,		// Type="Boolean"
            "Options": [
                {
                    "Code": "string",       // Type="String"
                    "Description": "string" // Type="String"
                }
            ]
        }
    ]
}
```

---

### Custom Fields Lookup Options

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomField/:entityTypeName/lookup?$top=100&$skip=0&$count=true&fieldCode={{custom-field-code}}`

Custom field lookup options for `lookup` field types.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |
| `fieldCode` | `{{custom-field-code}}` | The custom field code from the custom field lookup endpoint |

**Example Response: Get Custom fields for a type** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFieldDetail",
    "@odata.count": 1,
    "value": [       
         {
            "Code": "string",			// Type="String"
            "Label": "string",			// Type="String"
            "Type": "string",			// Type="String"
            "IsRequired": false,		// Type="Boolean"
            "Options": [
                {
                    "Code": "string",       // Type="String"
                    "Description": "string" // Type="String"
                }
            ]
        }
    ]
}
```

---

### Custom Fields Cascading Lookup Options

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomField/:entityTypeName/lookup/cascade?$top=100&$skip=0&$count=true&fieldCode={{custom-field-code}}&parentValue={{parent-field-value}}&filterParentByCode=false`

Custom field lookup options for `lookup` field types where the options depends on another field's selected option.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |
| `fieldCode` | `{{custom-field-code}}` | The custom field code from the custom field lookup endpoint |
| `parentValue` | `{{parent-field-value}}` | The parent field selected description or option code |
| `filterParentByCode` | `false` | Optional, By default it checks the selected parent description and not the code.</br>`true` Match with the option Code.</br>`false` Match with option description |

**Example Response: Get Custom fields for a type** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFieldDetail",
    "@odata.count": 1,
    "value": [       
         {
            "Code": "string",			// Type="String"
            "Label": "string",			// Type="String"
            "Type": "string",			// Type="String"
            "IsRequired": false,		// Type="Boolean"
            "Options": [
                {
                    "Code": "string",       // Type="String"
                    "Description": "string" // Type="String"
                }
            ]
        }
    ]
}
```

---
## Custom Forms

### Custom Form Category Code

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomFormCategoryCode?$top=100&$skip=0&$count=true`

Custom form Category codes

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Custom fields for a type** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFieldDetail",
    "@odata.count": 1,
    "value": [       
         {
            "Code": "string",			// Type="String"
            "Label": "string",			// Type="String"
            "Type": "string",			// Type="String"
            "IsRequired": false,		// Type="Boolean"
            "Options": [
                {
                    "Code": "string",       // Type="String"
                    "Description": "string" // Type="String"
                }
            ]
        }
    ]
}
```

---

### Custom Form Custom Fields

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomField/CustomForm?$top=100&$skip=0&$count=true&CustomFormCategoryCode={{custom-form-category-code}}`

Lookup custom fields for Custom forms

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |
| `CustomFormCategoryCode` | `{{custom-form-category-code}}` | Value from the CustomFormCategoryCode lookup |

**Example Response: Get Custom fields for a type** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFieldDetail",
    "@odata.count": 1,
    "value": [       
         {
            "Code": "string",			// Type="String"
            "Label": "string",			// Type="String"
            "Type": "string",			// Type="String"
            "IsRequired": false,		// Type="Boolean"
            "Options": [
                {
                    "Code": "string",       // Type="String"
                    "Description": "string" // Type="String"
                }
            ]
        }
    ]
}
```

---

### Custom Form Lookup Options

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomField/CustomForm/:categoryCode/lookup?$top=100&$skip=0&$count=true&fieldCode={{custom-field-code}}`

Custom form field lookup options for `lookup` field types

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |
| `fieldCode` | `{{custom-field-code}}` | The custom field code from the custom field lookup endpoint |

**Example Response: Get Custom fields for a type** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFieldDetail",
    "@odata.count": 1,
    "value": [       
         {
            "Code": "string",			// Type="String"
            "Label": "string",			// Type="String"
            "Type": "string",			// Type="String"
            "IsRequired": false,		// Type="Boolean"
            "Options": [
                {
                    "Code": "string",       // Type="String"
                    "Description": "string" // Type="String"
                }
            ]
        }
    ]
}
```

---

### Custom Form Cascading Lookup Field Options

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomField/CustomForm/:categoryCode/lookup/cascade?$top=100&$skip=0&$count=true&fieldCode={{custom-field-code}}&parentValue={{parent-field-value}}&filterParentByCode=false`

Custom form field lookup options for `lookup` field types where the options depends on another field's selected option.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |
| `fieldCode` | `{{custom-field-code}}` | The code from the custom field lookup endpoint |
| `parentValue` | `{{parent-field-value}}` | The parent field selected description or option code |
| `filterParentByCode` | `false` | Optional, By default it checks the selected parent description and not the code.</br>`true` Match with the option Code.</br>`false` Match with option description |

**Example Response: Get Custom fields for a type** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFieldDetail",
    "@odata.count": 1,
    "value": [       
         {
            "Code": "string",			// Type="String"
            "Label": "string",			// Type="String"
            "Type": "string",			// Type="String"
            "IsRequired": false,		// Type="Boolean"
            "Options": [
                {
                    "Code": "string",       // Type="String"
                    "Description": "string" // Type="String"
                }
            ]
        }
    ]
}
```

---
## ActivityCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/ActivityCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get ActivityCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ActivityCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## LeaveReason

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/LeaveReason?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get LeaveReason Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#LeaveReason",
    "@odata.count": 1,
    "value": [
        {
            "LeaveType": "string",              // Type="String"
            "LeaveBucket": "string",            // Type="String"
            "CompanyLeaveScheme": "string",     // Type="String"
            "Value": "string",			        // Type="String"
            "Description": "string"		        // Type="String"
        }
    ]
}
```

---
## LeaveType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/LeaveType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get LeaveType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#LeaveType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CompanyLegalBody

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyLegalBody?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyLegalBody Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLegalBody",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## AppealOutcome

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/AppealOutcome?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get AppealOutcome Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#AppealOutcome",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## AppealReason

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/AppealReason?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get AppealReason Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#AppealReason",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## AwardFavour

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/AwardFavour?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get AwardFavour Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#AwardFavour",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## LeaveBucket

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/LeaveType/{{LeaveType}}?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get LeaveBucket Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#LeaveType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## LeaveCompanyRun

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/LeaveCompanyRun?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` |  |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get LeaveBucket Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#LeaveCompanyRun",
    "@odata.count": 1,
    "value": [
        {
            "PeriodStartDate": "2019-03-01T00:00:00+02:00",	// Type="DateTime"
            "PeriodEndDate": "2019-03-31T00:00:00+02:00",	// Type="DateTime"
            "PayDate": "2019-07-31T00:00:00+02:00",			// Type="DateTime"
            "RunStatus": "Closed",							// Type="String" [Open, Closed, Future]
            "OrderNumber": 1,	        					// Type="Int64"
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## AccountType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/AccountType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get AccountType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#AccountType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## AddressCountry

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/AddressCountry?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get AddressCountry Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#AddressCountry",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## AddressCountry Province

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/AddressCountry/{{CountryValue}}?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get AddressCountry Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#AddressCountry",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## BasedOnSplitOption

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/BasedOnSplitOption?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get BasedOnSplitOption Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#BasedOnSplitOption",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",		    // Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## AdministratorEmployeeNumber

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/AdministratorEmployeeNumber?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get AdministratorEmployeeNumber Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#LookupDataSet",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## BudgetGroup

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/BudgetGroup?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get BudgetGroup Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#BudgetGroup",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CarPaymentType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CarPaymentType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CarPaymentType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CarPaymentType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Citizenship

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Citizenship?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Citizenship Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Citizenship",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string",	// Type="String"
            "ISO2DigitCode": "string",	// Type="String"
            "ISO3DigitCode": "string"	// Type="String"
        }
    ]
}
```

---
## Classification

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Classification?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Classification Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Classification",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CompanyComponent

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyComponent?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyComponent Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyComponent",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## CompanyDisability

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyDisability?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyDisability Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyDisability",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CompanyEdbIndicator

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyEdbIndicator?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` | Required |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyEdbIndicator Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyEdbIndicator",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## CompanyFrequency

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyFrequency?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyFrequency Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyFrequency",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string",	// Type="String"
            "Frequency": "string"       // Type="String" (Monthly, Weekly or Fortnightly)
        }
    ]
}
```

---
## CompanyGlAccount

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OverrideOfoCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyGlAccount Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGlAccount",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CompanyGroupLife

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyGroupLife?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` | Required |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyGroupLife Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGroupLife",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## CompanyIncidentType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyIncidentType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` | Required |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyIncidentType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyIncidentType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## CompanyMedicalAid

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyMedicalAid?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` | Required |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyMedicalAid Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyMedicalAid",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## CompanyOffence

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyOffence?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyOffence Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyOffence",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CompanyPensionFund

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyPensionFund?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` | Required |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyPensionFund Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyPensionFund",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## CompanyBudgetPeriod

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyBudgetPeriod?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` | Required |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyPensionFund Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyBudgetPeriod",
    "@odata.count": 1,
    "value": [
        {
            "PeriodStartDate": "2019-03-01T00:00:00+02:00",                     // Type="DateTime"
            "PeriodEndDate": "2019-03-01T00:00:00+02:00",                       // Type="DateTime"
            "Value": "string",                                                  // Type="String"
            "Description": "string"                                             // Type="String"
        }
    ]
}
```

---
## CompanyPositionFamily

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyPositionFamily?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` | Required |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyPositionFamily Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyPositionFamily",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## CompanyRun

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyRun?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyRun Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyRun",
    "@odata.count": 1,
    "value": [
        {
        	"PeriodStartDate": "2019-03-01T00:00:00+02:00",	// Type="DateTime"
            "PeriodEndDate": "2019-03-31T00:00:00+02:00",	// Type="DateTime"
            "PayDate": "2019-07-31T00:00:00+02:00",			// Type="DateTime"
            "RunStatus": "Closed",							// Type="String" [Open, Closed, Future]
            "OrderNumber": 1,	        					// Type="Int64"
            "RunType": "string",							// Type="String"
            "Value": "string",								// Type="String"
            "Description": "string",						// Type="String"
            "CutOffDate": "2019-07-31T00:00:00+02:00"       // Type="DateTime"
        }
    ]
}
```

---
## CompanyTemplateDefinition

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyRun?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CompanyTemplateDefinition Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyTemplateDefinition",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## ComponentCompany

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/ComponentCompany?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get ComponentCompany Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ComponentCompany",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",								// Type="String"
            "Description": "string"							// Type="String"
        }
    ]
}
```

---
## ComponentCompanyDetail

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/ComponentCompanyDetail?frequency={{frequency}}&period={{run}}&$count={{$count}}&$filter={{$filter}}&allowOnceOffComponents={{allowOnceOffComponents}}`

Returns the list of components that can be used with the recurring components endpoints.

> To include onceoff components in the result use the `allowOnceOffComponent` query string value


**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Company frequency in company to query against |
| `period` | `{{run}}` | company run to query against |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Actoion eq 'Deduction'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `allowOnceOffComponents` | `{{allowOnceOffComponents}}` | Optional (true/false) - Indicates to return onceoff components in the results. |

**Example Response: ComponentCompanyDetail** (Status: 200 )

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ComponentCompanyDetail",
    "@odata.count": 1,
    "value": [
        {
            "ComponentId": 0,                               // Type="Int64"
            "Action": "string",                             // Type="String"
            "ComponentType": "string",                      // Type="String"
            "ComponentCode": "string",                      // Type="String"
            "FormulaTableType": "string",                   // Type=Enums.FormulaTables
            "TaxCode": "string",                            // Type="String"
            "ComponentLines": [
                {
                    "CalculationElement": "string",         // Type="String"
                    "Description": "string",                // Type="String"
                    "ValueType": "string",                  // Type="String"
                    "Indicators": [
                        {
                            "Description": "Item 1",        // Type="String"
                            "Variable": 5.00000             // Type="Decimal"
                        },
                        {
                            "Description": "Item 2",        // Type="String"
                            "Variable": 15.00000            // Type="Decimal"
                        }
                    ],
                    "ComponentValue": 0.0                   // Type="Decimal"
                }
            ],
            "ComponentSubCodes": [
                {
                    "SubCode": "string",                    // Type="String"
                    "SubCodeDescription": "string",         // Type="String"
                    "Component": "string",                  // Type="String"
                    "EffectiveDate": "2019-05-03T00:00:00"  // Type="DateTime"                    
                }
            ],
            "InPackage": false,                             // Type="Boolean"
            "EnforcePackageRule": false,                    // Type="Boolean"
            "FieldCode": "string",                          // Type="String"
            "EnableBcoe": false,                            // Type="Boolean"
            "MedicalCategory": "string",                    // Type="String"
            "ComponentCategory": "string",                  // Type="String"
            "Value": "string",                              // Type="String"
            "Description": "string",                        // Type="String"
            "IsStatutoryComponent": false,                  // Type="Boolean"
            "IsOnceOffComponent": false,                    // Type="Boolean"
            "CannotEditInPaySlip": false,                   // Type="Boolean"
            "DontDisplayInCompanyRecurring": false,         // Type="Boolean"
            "DontDisplayInEmployeeRecurring": false        // Type="Boolean"
        }
    ]
}
```

---
## ComponentCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/ComponentCode?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get ComponentCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ComponentCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## ComponentCode InputType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/ComponentCode/:component-code?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get InputType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ComponentCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CostCentreCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CostCentreCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CostCentreCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CostCentreCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Country

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Country?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Country Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Country",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string",	// Type="String"
            "ISO2DigitCode": "string",	// Type="String"
            "ISO3DigitCode": "string"	// Type="String"
        }
    ]
}
```

---
## Currency

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Currency?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Currency Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Currency",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Culture

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Culture?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Culture Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Lookup_Culture",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CustomTradeUnion

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomTradeUnion?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CustomTradeUnion Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomTradeUnion",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## DisabledType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/DisabledType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get DisabledType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#DisabledType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## DirectlyReportsEmployeeNumber

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/DirectlyReportsEmployeeNumber?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get DirectlyReportsEmployeeNumber Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#DirectlyReportsEmployeeNumber",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## DirectlyReportsPositionOverride

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/DirectlyReportsPositionOverride?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get DirectlyReportsPositionOverride Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#DirectlyReportsPositionOverride",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## EmployeeNumber

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/EmployeeNumber?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get EmployeeNumber Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeNumber",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## EmployeeTakeOnComponentCompany

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/EmployeeTakeOnComponentCompany?frequency={{frequency}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Required `value` from `CompanyFrequency` lookup |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get EmployeeTakeOnComponentCompany  Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTakeOnComponentCompany",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## EmploymentAction

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/EmploymentAction?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get EmploymentAction Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmploymentAction",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## EmploymentCategory

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/EmploymentCategory?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get EmploymentCategory Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmploymentCategory",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## FinancialYear

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/FinancialYear?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get FinancialYear Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#FinancialYear",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## EmploymentSubCategory

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/EmploymentSubCategory?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get EmploymentSubCategory Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmploymentSubCategory",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## FunctionArea

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/FunctionArea?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get FunctionArea Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#FunctionArea",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## GeneralLedger

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/GeneralLedger?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get GeneralLedger Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#GeneralLedger",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Grade

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Grade?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Grade Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Grade",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## IdentityType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/IdentityType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get IdentityType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#IdentityType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## InputType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/nputType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get InputType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#InputType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Job

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Job?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Job Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Job",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Language

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Language?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Language Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Language",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## MaritalStatus

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/MaritalStatus?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get MaritalStatus Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#MaritalStatus",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## ManagerEmployeeNumber

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/ManagerEmployeeNumber?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get ManagerEmployeeNumber Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ManagerEmployeeNumber",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Nationality

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Nationality?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Nationality Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Nationality",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string",	// Type="String"
            "ISO2DigitCode": "string",	// Type="String"
            "ISO3DigitCode": "string"	// Type="String"
        }
    ]
}
```

---
## NatureOfPerson

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/NatureOfPerson?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get NatureOfPerson Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#NatureOfPerson",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OccupationalLevel

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OccupationalLevel?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OccupationalLevel Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OccupationalLevel",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OfoCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OfoCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OfoCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OfoCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OrganizationGroup

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OrganizationGroup?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OrganizationGroup Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationGroup",
    "@odata.count": 1,
    "value": [
        {
            "OrganizationGroupDescription": "string", // Type="String"
            "Value": "string",			              // Type="String" (UploadCode)
            "Description": "string"		              // Type="String"
        }
    ]
}
```

---
## OrganizationCategory

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OrganizationCategory?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OrganizationCategory Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationCategory",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OrganizationLevel

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OrganizationLevel?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OrganizationLevel Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationLevel",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OrganizationPosition

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OrganizationPosition?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OrganizationPosition Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationPosition",
    "@odata.count": 1,
    "value": [
        {
            "UploadCode": "string",     // Type="String"
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OrganizationPositionWithCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OrganizationPositionWithCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OrganizationPosition Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationPosition",
    "@odata.count": 1,
    "value": [
        {
            "UploadCode": "string",     // Type="String"
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OrganizationRegion

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OrganizationRegion?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OrganizationRegion Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OrganizationRegion",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OverrideOfoCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OverrideOfoCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OverrideOfoCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OverrideOfoCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OverrideRoster

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OverrideOfoCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OverrideRoster Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OverrideRoster",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OverrideOccupationalLevel

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OverrideOccupationalLevel?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OverrideOccupationalLevel Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OverrideOccupationalLevel",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OtherOutcome

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OtherOutcome?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OtherOutcome Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OtherOutcome",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OffenceOutcome

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OffenceOutcome?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OffenceOutcome Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OffenceOutcome",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## OffenceCategory

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/OffenceCategory?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get OffenceCategory Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#OffenceCategory",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## PassportCountry

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/PassportCountry?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get PassportCountry Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#PassportCountry",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Postal Code - Company Lookup

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/PostalCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get PostalCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Lookup_PostalCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Postal Code - Country Lookup

**`GET`** `{{api-base-url}}/odata/v2.0/country-lookup/:country-id/PostalCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get PostalCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Lookup_PostalCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## PayPoint

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/PayPoint?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get PayPoint Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#PayPoint",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## ProjectCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/ProjectCode?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get ProjectCode Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ProjectCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## PreferredRace

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/PreferredRace?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get PreferredRace Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#PreferredRace",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## PreferredLocalization

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/PreferredLocalization?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get PreferredLocalization Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#PreferredRace",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Race

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Race?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Race Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Race",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Reason

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Reason?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Reason Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Reason",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## ReportsToJob

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Reason?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get ReportsToJob Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#ReportsToJob",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Roster

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Roster?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Roster Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Roster",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## SecurityRole

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/SecurityRole?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get SecurityRole Lookup values**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Lookup_SecurityRole",
    "@odata.count": 0,
    "value": [
        {
            "Value": "string",          // Type="String"
            "Description": "string",    // Type="String"
            "Id": 0                     // Type="Int64"
        }
   ]
}
```

---
## SdlExemption

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/SdlExemption?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get SdlExemption Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#SdlExemption",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## SSTLevel

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/SSTLevel?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get SSTLevel Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#SSTLevel",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## SuspensionReason

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/SuspensionReason?$top=100&$skip=0&$count=true`

> The `Description` or `Value` can be used when POST/PATCH `EmployeeSuspension` record

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get SuspensionReason Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#SuspensionReason",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## SettlementReinstate

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/SettlementReinstate?$top=100&$skip=0&$count=true`

> The `Description` or `Value` can be used when POST/PATCH `EmployeeSuspension` record

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get SettlementReinstate Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#SettlementReinstate",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## TableBuilderComponentCompany

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TableBuilderComponentCompany?frequency={{frequency}}&period={{run}}&$count={{$count}}&$filter={{$filter}}`

Retrieves employee pay rate information based on the specified `PayrateId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Company frequency in company to query against |
| `period` | `{{run}}` | company run to query against |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Actoion eq 'Deduction'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: TableBuilderComponentCompany**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TableBuilderComponentCompany",
    "@odata.count": 1,
    "value": [
        {
            "ComponentId": 0,                               // Type="Int64"
            "Action": "string",                             // Type="String"
            "ComponentType": "string",                      // Type="String"
            "ComponentCode": "string",                      // Type="String"
            "FormulaTableType": "string",                   // Type=Enums.FormulaTables
            "TableBuilderCategory": "string",               // Type="String"
            "TaxCode": "string",                            // Type="String"
            "ComponentLines": [
                {
                    "CalculationElement": "string",         // Type="String"
                    "Description": "string",                // Type="String"
                    "ValueType": "string",                  // Type="String"
                    "Indicators": [
                        {
                            "Description": "Item 1",        // Type="String"
                            "Variable": 5.00000             // Type="Decimal"
                        },
                        {
                            "Description": "Item 2",        // Type="String"
                            "Variable": 15.00000            // Type="Decimal"
                        }
                    ],
                    "ComponentValue": 0.0                   // Type="Decimal"
                }
            ],
            "ComponentSubCodes": [
                {
                    "SubCode": "string",                    // Type="String"
                    "SubCodeDescription": "string",         // Type="String"
                    "Component": "string",                  // Type="String"
                    "EffectiveDate": "2019-05-03T00:00:00"  // Type="DateTime"                    
                }
            ],
            "InPackage": false,                             // Type="Boolean"
            "EnforcePackageRule": false,                    // Type="Boolean"
            "FieldCode": "string",                          // Type="String"
            "EnableBcoe": false,                            // Type="Boolean"
            "MedicalCategory": "string",                    // Type="String"
            "ComponentCategory": "string",                  // Type="String"
            "Value": "string",                              // Type="String"
            "Description": "string",                        // Type="String"
            "IsStatutoryComponent": false,                  // Type="Boolean"
            "IsOnceOffComponent": false                     // Type="Boolean"
        }
    ]
}
```

---
## TableBuilderCategory

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TableBuilderCategory?frequency={{frequency}}&period={{run}}&$count={{$count}}&$filter={{$filter}}`

Retrieves employee pay rate information based on the specified `PayrateId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Company frequency in company to query against |
| `period` | `{{run}}` | company run to query against |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Actoion eq 'Deduction'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: TableBuilderCategory**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TableBuilderCategory",
    "@odata.count": 1,
    "value": [
        {
            "Fields": [               
                {
                    "Code": "string",           // Type="String"
                    "Label": "string",          // Type="String"
                    "Type": "string",           // Type="String"
                    "IsRequired": false,        // Type="Boolean"
                    "Options": [
                         {
                            "Code": "string",           // Type="String"
                            "Description": "string"     // Type="String"
                        }
                    ]
                }
            ],
            "IsEmployeeLevelOnly": false,       // Type="Boolean"
            "Value": "string",                  // Type="String"
            "Description": "string"             // Type="String"
        }
    ]
}
```

---
## TableBuilderCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TableBuilderComponentCompany/:TableBuilderComponentCompany?frequency={{frequency}}&period={{run}}&$count={{$count}}&$filter={{$filter}}`

Retrieves employee pay rate information based on the specified `PayrateId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Company frequency in company to query against |
| `period` | `{{run}}` | company run to query against |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "Actoion eq 'Deduction'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: TableBuilderComponentCompany**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TableBuilderCode",
    "@odata.count": 1,
    "value": [
        {            
            "Value": "string",          // Type=string
            "Description": "string"     // Type=string
        }
    ]
}
```

---
## TaxCode

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TaxCode?frequency={{frequency}}&period={{run}}&$count={{$count}}`

Retrieves tax codes based on the specified `company-id`, `company-frequency`, and `company-run-period`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{frequency}}` | Company frequency in company to query against |
| `period` | `{{run}}` | company run to query against |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get TaxCode Lookup values**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TaxCode",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## TaxStatus

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TaxStatus?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get TaxStatus Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TaxStatus",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## TerminationCompanyRun

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TerminationCompanyRun?frequency={{FequencyName}}&$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `frequency` | `{{FequencyName}}` | Required |
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get TerminationCompanyRun Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TerminationCompanyRun",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## TerminationReason

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TerminationReason?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get TerminationReason Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TerminationReason",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## Title

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Title?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Title Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Title",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## TradeUnion

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TradeUnion?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get TradeUnion Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TradeUnion",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## UifExemption

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/UifExemption?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get UifExemption Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#UifExemption",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## WorkflowRole

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/WorkflowRole?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get WorkflowRole Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#WorkflowRole",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CustomFormCategory

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CustomFormCategoryCode`

**Example Response: CustomFormCategory** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CustomFormCategoryCode",
    "value": [
        {
            "CustomFormLevel": "string",        // Type="String"
            "CompanyLevel": true,               // Type="Boolean"
            "EmployeeLevel": false,             // Type="Boolean"
            "Value": "string",                  // Type="String"
            "Description": "string"             // Type="String"
        }
    ]
}
```

---
## RecurringTemplate

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/RecurringTemplate?$top=100&$skip=0&$count=true&frequency={{FequencyName}}&period={{run}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |
| `frequency` | `{{FequencyName}}` |  |
| `period` | `{{run}}` |  |

**Example Response: Get RecurringTemplate Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#RecurringTemplate",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## PositionType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/PositionType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get PositionType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#PositionType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"		// Type="String"
        }
    ]
}
```

---
## CourseType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CourseType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get CourseType Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CourseType",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"     // Type="String"
        }
    ]
}
```

---
## NqfLevel

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/NqfLevel?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get NqfLevel Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#NqfLevel",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"     // Type="String"
        }
    ]
}
```

---
## TrainingSkillsPriority

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TrainingSkillsPriority?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get TrainingSkillsPriority Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TrainingSkillsPriority",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"     // Type="String"
        }
    ]
}
```

---
## TrainingProgramCategory

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/TrainingProgramCategory?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get TrainingProgramCategory Lookup values** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#TrainingProgramCategory",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"     // Type="String"
        }
    ]
}
```

---
## CompanyLeaveSetup

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyLeaveSetup?$top=100&$skip=0&$count=true`

Note that `Value` is a concatenated result of the Leave Scheme Name and the Leave Setup Description (e.g. "Standard Leave - Annual Leave")

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: CompanyLeaveSetup**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyLeaveSetup",
    "@odata.count": 1,
    "value": [
        {
            "LeaveType": "string",              // Type="String"
            "LeaveBucket": "string",            // Type="String"
            "CompanyLeaveScheme": "string",     // Type="String"
            "Value": "string",                  // Type="String" (Scheme Name - Setup Description)
            "Description": "string"             // Type="String"
        }
    ]
}
```

---
## Gender

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/Gender?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Gender Lookup values**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Gender",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"	    // Type="String"
        }
    ]
}
```

---
## CompanyGradeFieldValue

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyGradeFieldValue`

**Example Response: CompanyGradeFieldValue**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyGradeFieldValue",
    "@odata.count": 1,
    "value": [
        {
            "Value": "string",			// Type="String"
            "Description": "string"     // Type="String"
        }
    ]
}
```

---
## CompanyGroupActionType

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyGroupActionType?$top=100&$skip=0&$count=true`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `100` |  |
| `$skip` | `0` |  |
| `$count` | `true` |  |

**Example Response: Get Group Action Type Lookup values**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Lookup_CompanyGroupActionType",
    "value": [
        {
            "Value": "string",      // Type="String"
            "Description": "string" // Type="String"
        }
    ]
}
```

---
## CompanyLeaveSetup

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyLeaveSetup`

**Example Response: CompanyLeaveSetup** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Lookup_CompanyLeaveSetup",
    "@odata.count": 1,
    "value": [
        {
            "LeaveType": "string",              // Type="String"
            "LeaveBucket": "string",            // Type="String"
            "CompanyLeaveScheme": "string",     // Type="String"
            "Value": "string",                  // Type="String"
            "Description": "string"             // Type="String"
        }
    ]
}
```

---

# Employee > Basic Information


Contains several endpoints related to employee information. 




## Biographical

Retrieve employees' biographical information associated with the `company_id`.

### Get a collection of employees

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Employee?$orderby={{$employee-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&$select={{$select}}`

Retrieves a list of active employee biographical details as of the today's date and based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$select` | `{{$select}}` | Optional (string) - Returns only the fields specified |

**Example Response: Get a collection of employees** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Employee",
    "@odata.count": 1,
    "value": [
        {
				"EmployeeId": 0,				            // Type="Int64"
                "InitiateWorkFlow": false,                  // Type="Boolean" - optional values used with POST and Patch to initiate the employee workflow
			    "EmployeeNumber": "string",		       		// Type="String"
				"Title": "string",  			        	// Type="String"
				"FirstName": "string",			        	// Type="String"
				"LastName": "string",			        	// Type="String"
				"PreferredName": "string",		    		// Type="String"
				"MaidenName": "string",			        	// Type="String"
				"MiddleName": "string",			        	// Type="String"
				"Initials": "string",		        		// Type="String"
				"Email": "string",			            	// Type="String"
				"Birthday": "2019-05-03T08:57:54.266Z",		// Type="DateTime"
				"HomeNumber": "string",			        	// Type="String"
				"WorkNumber": "string",			        	// Type="String"
				"CellNumber": "string",			        	// Type="String"
				"WorkExtension": "string",		    		// Type="String"
				"Language": "string",				        // Type="String"
				"Gender": "Male",		  		            // Type="String"
				"MaritalStatus": "string",		    		// Type="String"
				"Race": "string",				            // Type="String"
				"Nationality": "string",			    	// Type="String"
				"Citizenship": "string",		    		// Type="String"
				"DisabledType": "string",				    // Type="String"
				"ForeignNational": true,				    // Type="Boolean"
				"DateCreated": "2019-05-03T08:57:54.266Z",	// Type="DateTime"
				"EmergencyContactName": "string",			// Type="String"
				"EmergencyContactNumber": "string",			// Type="String"
				"EmergencyContactAddress": "string",		// Type="String"
				"IsRetired": true,			            	// Type="Boolean"
				"CustomFieldValue": "string",				// Type="String"
				"CustomFieldValue2": "string",				// Type="String"
				"UifExemption": "string",					// Type="String"
				"SdlExemption": "string",					// Type="String"
				"EtiExempt": false,							// Type="Boolean"
                "ImageDownloadUrl": "string",                // Type="String"
				"CustomFields": [		
					{			
					  "Code": "string",				        // Type="String"
					  "Label": "string",			        // Type="String"
					  "Value": "string"			        	// Type="String"
					}	
				],        
                "Address": [
                    {
                        "AddressId": 0,					// Type="Int64"
                        "AddressType": "string",		// Type="String"
                        "EmployeeNumber": "string",		// Type="String"
                        "AddressLine1": "string",		// Type="String"
                        "AddressLine2": "string",		// Type="String"
                        "AddressLine3": "string",		// Type="String"
                        "AddressCode": "string",		// Type="String"
                        "AddressCountry": "string",		// Type="Strin
... (truncated)
```

---

### Get a collection of employees  as of an effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Employee/effective/:effectivedate?$orderby={{$employee-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of active employee biographical details as of the specified effective date and based on the specified `company_id`.


**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of employees  as of an effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Employee",
    "@odata.count": 1,
    "value": [
        {
			"EmployeeId": 0,				            // Type="Int64"
            "InitiateWorkFlow": false,                  // Type="Boolean" - optional values used with POST and Patch to initiate the employee workflow
		    "EmployeeNumber": "string",		       		// Type="String"
			"Title": "string",  			        	// Type="String"
			"FirstName": "string",			        	// Type="String"
			"LastName": "string",			        	// Type="String"
			"PreferredName": "string",		    		// Type="String"
			"MaidenName": "string",			        	// Type="String"
			"MiddleName": "string",			        	// Type="String"
			"Initials": "string",		        		// Type="String"
			"Email": "string",			            	// Type="String"
			"Birthday": "2019-05-03T08:57:54.266Z",		// Type="DateTime"
			"HomeNumber": "string",			        	// Type="String"
			"WorkNumber": "string",			        	// Type="String"
			"CellNumber": "string",			        	// Type="String"
			"WorkExtension": "string",		    		// Type="String"
			"Language": "string",				        // Type="String"
			"Gender": "Male",				            // Type="String"
			"MaritalStatus": "string",		    		// Type="String"
			"Race": "string",				            // Type="String"
			"Nationality": "string",			    	// Type="String"
			"Citizenship": "string",		    		// Type="String"
			"TaxRefNumber": "string",		    		// Type="String"
			"Disabled": true,			            	// Type="Boolean"
			"DisabledType": "string",				    // Type="String"
			"EthnicGroup": "string",				    // Type="String"
			"ForeignNational": true,				    // Type="Boolean"
			"GroupDate": "2019-05-03T08:57:54.266Z",	// Type="DateTime"
			"DateCreated": "2019-05-03T08:57:54.266Z",	// Type="DateTime"
			"EmergencyContactName": "string",			// Type="String"
			"EmergencyContactNumber": "string",			// Type="String"
			"EmergencyContactAddress": "string",		// Type="String"
			"IsMockEmployee": true,			        	// Type="Boolean"
			"IsRetired": true,			            	// Type="Boolean"
			"CustomFieldValue": "string",				// Type="String"
			"CustomFieldValue2": "string",				// Type="String"
			"UifExemption": "string",					// Type="String"
			"SdlExemption": "string",					// Type="String"
			"EtiExempt": false,							// Type="Boolean"
            "ImageDownloadUrl": "string",                // Type="String"
			"CustomFields": [		
				{			
				  "Code": "string",				        // Type="String"
				  "Label": "string",			        // Type="String"
				  "Value": "string"			        	// Type="String"
				}	
			],        
            "Address": [
                {
                    "AddressId": 0,					// Type="Int64"
                    "AddressType": "string",		// Type="String"
                    "EmployeeNumber": "string",		// Type="String"
                    "AddressLine1": "string",		// Type="String"
                    "AddressLine2": "string",		// Type="String"
  
... (truncated)
```

---

### Get a single employee record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Employee({{EmployeeId}})`

Retrieves employee information based on the specified `EmployeeId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single employee record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Employee/$entity",
	"EmployeeId": 0,				            // Type="Int64"
    "InitiateWorkFlow": false,                  // Type="Boolean" - optional values used with POST and Patch to initiate the employee workflow
    "EmployeeNumber": "string",		       		// Type="String"
	"Title": "string",  			        	// Type="String"
	"FirstName": "string",			        	// Type="String"
	"LastName": "string",			        	// Type="String"
	"PreferredName": "string",		    		// Type="String"
	"MaidenName": "string",			        	// Type="String"
	"MiddleName": "string",			        	// Type="String"
	"Initials": "string",		        		// Type="String"
	"Email": "string",			            	// Type="String"
	"Birthday": "2019-05-03T08:57:54.266Z",		// Type="DateTime"
	"HomeNumber": "string",			        	// Type="String"
	"WorkNumber": "string",			        	// Type="String"
	"CellNumber": "string",			        	// Type="String"
	"WorkExtension": "string",		    		// Type="String"
	"Language": "string",				        // Type="String"
	"Gender": "Male",				            // Type="Enum" see "Gender" in metadata endpoint
	"MaritalStatus": "string",		    		// Type="String"
	"Race": "string",				            // Type="String"
	"Nationality": "string",			    	// Type="String"
	"Citizenship": "string",		    		// Type="String"
	"DisabledType": "string",				    // Type="String"
	"ForeignNational": true,				    // Type="Boolean"
	"DateCreated": "2019-05-03T08:57:54.266Z",	// Type="DateTime"
	"EmergencyContactName": "string",			// Type="String"
	"EmergencyContactNumber": "string",			// Type="String"
	"EmergencyContactAddress": "string",		// Type="String"
	"IsRetired": true,			            	// Type="Boolean"
	"CustomFieldValue": "string",				// Type="String"
	"CustomFieldValue2": "string",				// Type="String"
	"UifExemption": "string",					// Type="String"
	"SdlExemption": "string",					// Type="String"
	"EtiExempt": false,							// Type="Boolean"    
    "ImageDownloadUrl": "string",                // Type="String"
	"CustomFields": [		
		{			
		  "Code": "string",				        // Type="String"
		  "Label": "string",			        // Type="String"
		  "Value": "string"			        	// Type="String"
		}	
	],        
    "Address": [
        {
            "AddressId": 0,					// Type="Int64"
            "AddressType": "string",		// Type="String"
            "EmployeeNumber": "string",		// Type="String"
            "AddressLine1": "string",		// Type="String"
            "AddressLine2": "string",		// Type="String"
            "AddressLine3": "string",		// Type="String"
            "AddressCode": "string",		// Type="String"
            "AddressCountry": "string",		// Type="String"
            "Province": "string",			// Type="String"
            "UnitNumber": "string",			// Type="String"
            "Complex": "string",			// Type="String"
            "StreetNumber": "string",		// Type="Strin
... (truncated)
```

---

### Get a collection of all employees

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Employee/all?$orderby={{$employee-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&$select={{$select}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$select` | `{{$select}}` | Optional (string) - Returns only the fields specified |

**Example Response: Get a collection of all employees**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#Employee",
    "@odata.count": 1,
    "value": [
        {
				"EmployeeId": 0,				            // Type="Int64"
			    "EmployeeNumber": "string",		       		// Type="String"
                "CompanyId": 0,                             // Type="Int64"
                "CompanyFrequency": "string",               // Type="String"
                "DefaultPayslip": "string",                 // Type="String"
                "InitiateWorkFlow": false,                  // Type="Boolean" - optional values used with POST and Patch to initiate the employee workflow
				"Title": "string",  			        	// Type="String"
				"FirstName": "string",			        	// Type="String"
				"LastName": "string",			        	// Type="String"
				"PreferredName": "string",		    		// Type="String"
				"MaidenName": "string",			        	// Type="String"
				"MiddleName": "string",			        	// Type="String"
				"Initials": "string",		        		// Type="String"
				"Email": "string",			            	// Type="String"
				"Birthday": "2019-05-03T08:57:54.266Z",		// Type="DateTime"
				"HomeNumber": "string",			        	// Type="String"
				"WorkNumber": "string",			        	// Type="String"
				"CellNumber": "string",			        	// Type="String"
				"WorkExtension": "string",		    		// Type="String"
				"Language": "string",				        // Type="String"
				"Gender": "Male",		  		            // Type="String"
				"MaritalStatus": "string",		    		// Type="String"
				"Race": "string",				            // Type="String"
				"Nationality": "string",			    	// Type="String"
				"Citizenship": "string",		    		// Type="String"
				"DisabledType": "string",				    // Type="String"
				"ForeignNational": true,				    // Type="Boolean"
				"DateCreated": "2019-05-03T08:57:54.266Z",	// Type="DateTime"
				"EmergencyContactName": "string",			// Type="String"
				"EmergencyContactNumber": "string",			// Type="String"
				"EmergencyContactAddress": "string",		// Type="String"
				"IsRetired": true,			            	// Type="Boolean"
				"CustomFieldValue": "string",				// Type="String"
				"CustomFieldValue2": "string",				// Type="String"
				"UifExemption": "string",					// Type="String"
				"SdlExemption": "string",					// Type="String"
				"EtiExempt": false,							// Type="Boolean"
                "ImageDownloadUrl": "string",               // Type="String"
				"CustomFields": [		
					{			
					  "Code": "string",				        // Type="String"
					  "Label": "string",			        // Type="String"
					  "Value": "string",			        // Type="String"
                      "OptionCode": "string"                // Type="String"
					}	
				],        
                "Address": [
                    {
                        "EmployeeId": 0,				    // Type="Int64"
                        "EmployeeNumber": "string",		    // Type="String"
                        "AddressId":
... (truncated)
```

---

### Create a single employee record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/Employee`

Create a single employee record for the specified `EmployeeNumber`.

---

### Employee Workflow

Too initiate the employee workflow set the field `InitiateWorkFlow` to true. The `InitiateWorkFlow` field can be used with `POST` and `PATCH` endpoints.

---

### Employee Address

The following rules apply when posting a new employee:

- Use the `Employee` endpoint to create addresses.
- When posting a new `Employee` at least one address is required.
- The 1st address in the list must always be a `Physical` address.
- When posting only one address it must be a `Physical` address. A second `Street` address will automatically be added and `SameAsPhysical` will be set to true.
- Duplicate address types are not allowed.
- It is recommended to use the `EmployeeAddress` endpoint to `PATCH` the address, but it is also possible to patch the address using the `Employee`.
- If the `Employee` endpoint is used to `PATCH` the address then you must ensure that the `AddressId` is also present. This will ensure that the correct address is patched.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "Employee" in metadata endpoint for available fields
}
```

**Example Response: Create a single employee record** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeWithAddress/$entity",
	"EmployeeId": 0,				            // Type="Int64"
    "InitiateWorkFlow": false,                  // Type="Boolean" - optional values used with POST and Patch to initiate the employee workflow
    "EmployeeNumber": "string",		       		// Type="String"
	"Title": "string",  			        	// Type="String"
	"FirstName": "string",			        	// Type="String"
	"LastName": "string",			        	// Type="String"
	"PreferredName": "string",		    		// Type="String"
	"MaidenName": "string",			        	// Type="String"
	"MiddleName": "string",			        	// Type="String"
	"Initials": "string",		        		// Type="String"
	"Email": "string",			            	// Type="String"
	"Birthday": "2019-05-03T08:57:54.266Z",		// Type="DateTime"
	"HomeNumber": "string",			        	// Type="String"
	"WorkNumber": "string",			        	// Type="String"
	"CellNumber": "string",			        	// Type="String"
	"WorkExtension": "string",		    		// Type="String"
	"Language": "string",				        // Type="String"
	"Gender": "Male",				            // Type="String"
	"MaritalStatus": "string",		    		// Type="String"
	"Race": "string",				            // Type="String"
	"Nationality": "string",			    	// Type="String"
	"Citizenship": "string",		    		// Type="String"
	"DisabledType": "string",				    // Type="String"
	"ForeignNational": true,				    // Type="Boolean"
	"DateCreated": "2019-05-03T08:57:54.266Z",	// Type="DateTime"
	"EmergencyContactName": "string",			// Type="String"
	"EmergencyContactNumber": "string",			// Type="String"
	"EmergencyContactAddress": "string",		// Type="String"
	"IsRetired": true,			            	// Type="Boolean"
	"CustomFieldValue": "string",				// Type="String"
	"CustomFieldValue2": "string",				// Type="String"
	"UifExemption": "string",					// Type="String"
	"SdlExemption": "string",					// Type="String"
	"EtiExempt": false,							// Type="Boolean"    
    "ImageDownloadUrl": "string",                // Type="String"
	"CustomFields": [		
		{			
		  "Code": "string",				        // Type="String"
		  "Label": "string",			        // Type="String"
		  "Value": "string"			        	// Type="String"
		}	
	],
	"Address": [
	        {
	        	 "AddressId": 0,					// Type="Int64"
		         "AddressType": "string",			// Type="String"
		         "EmployeeNumber": "string",		// Type="String"
		         "AddressLine1": "string",			// Type="String"
		         "AddressLine2": "string",			// Type="String"
		         "AddressLine3": "string",			// Type="String"
		         "AddressCode": "string",			// Type="String"
		         "AddressCountry": "string",		// Type="String"
		         "Province": "string",				// Type="String"
		         "UnitNumber": "string",			// Type="String"
		         "Complex": "string",				// Type="String"
		         "StreetNumber": "string",			// Type="String"
		         "SameAsPhysical": tr
... (truncated)
```

---

### Update a single employee record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/Employee({{EmployeeId}})`

Update a single employee record based on the specified `EmployeeId`.

---

### Employee Workflow

Too initiate the employee workflow set the field `InitiateWorkFlow` to true. The `InitiateWorkFlow` field can be used with `POST` and `PATCH` endpoints.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "Employee" in metadata endpoint for available fields
}
```

**Example Response: Update a single employee record** (Status: 204 No Content)

---

### Download employee photo

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/Employee/{{EmployeeId}}/image/download`

Retrieves a single employee photo based on the specified `EmployeeId`.

**Example Response: Download employee photo** (Status: 200 OK)

```json
// Returns Downloaded Employee Photo
```

---

### Upload employee photo

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/Employee/{{EmployeeId}}/image/upload`

Update a single employee photo based on the specified `EmployeeId`.

**Request Body (form-data):**

| Key | Type | Value | Description |
|-----|------|-------|-------------|
| `image` | file | `` | Key name "Image" As "File" type |

**Example Response: Upload employee photo** (Status: 200 OK)

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#Employee/$entity",
  "Success": true
}
```

---
## Employee Address

Retrieve employee address information associated with an `EmployeeNumber`.

<br/>

---

### Employee Address

The following rules apply when posting a new employee:
 * Use the `Employee` endpoint to create addresses. 
 * When posting a new `Employee` at least one address is required.
 * The 1st address in the list must always be a `Physical` address. 
 * When posting only one address it must be a `Physical` address. A second `Postal` address will automatically be added and `SameAsPhysical` will be set to true.
 * Duplicate address types are not allowed.
 * It is recommended to use the `EmployeeAddress` endpoint to `PATCH` the address, but it is also possible to patch the address using the `Employee`. 
 * If the `Employee` endpoint is used to `PATCH` the address then you must ensure that the `AddressId` is also present. This will ensure that the correct address is patched. 

### Get an employee address

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAddress/{{EmployeeNumber}}?$select={{$select}}`

Retrieves a list of employee addresses for the given `EmployeeNumber`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$select` | `{{$select}}` |  |

**Example Response: Get an employee address** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAddress",
    "value": [
		        {
		            "AddressId": 0						// Type="Int64"
		            "AddressType": "string",			// Type="String"
		            "EmployeeNumber": "string",			// Type="String"
		            "AddressLine1": "string",			// Type="String"
		            "AddressLine2": "string",			// Type="String"
		            "AddressLine3": "string",			// Type="String"
		            "AddressCode": "string",			// Type="String"
		            "AddressCountry": "string",			// Type="String"
		            "Province": "string",				// Type="String"
		            "UnitNumber": "string",				// Type="String"
		            "Complex": "string",				// Type="String"
		            "StreetNumber": "string",			// Type="String"
		            "SameAsPhysical": true,				// Type="Boolean"
		            "IsCareofAddress": true,		   	// Type="Boolean"
		            "CareOfIntermediary": "string",		// Type="String"
		            "SpecialServices": "string"			// Type="String"
		        }
	    ]
}
```

---

### Update a single employee address record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAddress({{AddressId}})`

Update a single employee record based on the specified `AddressId`.

**Request Body:**

```json
{
  //see "EmployeeAddress" in metadata endpoint for available fields
}
```

**Example Response: Update a single employee address record** (Status: 204 No Content)

---
## Tax Profile

Retrieve active employees' status information associated with the `company_id`.

Reinstate is called with a `PATCH` where only the `EmploymentAction` is required and can have one of two values for a reinstate:

- ‘reinstate this employee resuming this tax record’ – here if pass in only the `EmploymentAction` then will clear the Terminate fields and update the `EmploymentAction`

- ‘reinstate this employee starting a new tax record’ – here if pass in only the `EmploymentAction` then will overwrite the `EmploymentAction` but leave the terminate fields in place. 
A new `EmploymentStatus` will be created which is a copy of the previous one without the terminate fields and `EmploymentAction` will also be NULL.

> **For ZA companies: Due to legislation changes in South Africa, the following fields have been made obsolete and are no longer relevant. The fields will not return values from 01/05/2021 and any changes to them will not be saved**
> - `DeemedMonthlyRemuneration`
> - `Deemed75Indicator`
> - `DeemedRecoveryMonthly`

### Get a collection of employment statuses

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus?$orderby={{$employment-status-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of current active employee employment statuses based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employment-status-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of employment statuses** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeEmploymentStatus",
    "@odata.count": 1,
    "value": [
				{
					"EmploymentStatusId": 0,								// Type="Int64"
					"EmployeeNumber": "string",								// Type="String"
					"GroupJoinDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"EmploymentDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"TerminationDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"TerminationReason": "string",							// Type="String"
					"TaxStatus": "string",									// Type="String"
					"TaxReferenceNumber": "string",							// Type="String"
					"NatureOfPerson": "string",								// Type="String"
					"TaxOffice": 0,											// Type="Int32"
					"TaxDirectiveNumber": "string",							// Type="String"
					"IT3AReason": 0,										// Type="Int32"
					"EmploymentAction": "string", 							// Type="String"
					"TerminationCompanyRun": "string",						// Type="String"
					"IdentityType": "string",								// Type="String"
					"IdNumber": "string",									// Type="String"
					"PassportNumber": "string",								// Type="String"
					"PercentageAmount": "string", 							// Type="String"
					"Amount": 0,											// Type="Decimal"
					"Percentage": 0,										// Type="Decimal"
					"DeemedMonthlyRemuneration": 0,							// Type="Decimal"
					"Deemed75Indicator": true,								// Type="Boolean"
					"DeemedRecoveryMonthly": true,							// Type="Boolean"
					"EncashLeave": true,									// Type="Boolean"
					"Irp30": true,											// Type="Boolean"
					"FinalizeIssueTaxCert": true,							// Type="Boolean"
					"PassportCountry": true,								// Type="String"
					"PassportIssued": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"PassportExpiry": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"PermitIssued": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
					"PermitExpiry": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
					"AdditionalDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"EmploymentCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
					"TerminationCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
					"TempWorker": true,										// Type="Boolean"
					"AdditionalDate1": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"NotReEmployable": true,								// Type="Boolean"
					"ReferenceNumber": "string",							// Type="String"
					"OldEmployeeId": 0,										// Type="Int64" - Creates a transfer record for ESS users to be able to access previous payslips
                    "CustomFields": [		
                        {			
                        "Code": "string",				        // Type="String"
                        "Label": "string",			        // Type="String"
                        "Value": "string"			        	// Type="String"
                        }	
				    ]
				}
    ]
}
```

---

### Get a single employment status record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus({{StatusId}})`

Retrieves employee employment status information based on the specified `StatusId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single employment status record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeEmploymentStatus/$entity",
	"EmploymentStatusId": 0,								// Type="Int64"
	"EmployeeNumber": "string",								// Type="String"
	"GroupJoinDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"EmploymentDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"TerminationDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"TerminationReason": "string",							// Type="String"
	"TaxStatus": "string",									// Type="String"
	"TaxReferenceNumber": "string",							// Type="String"
	"NatureOfPerson": "string",								// Type="String"
	"TaxOffice": 0,											// Type="Int32"
	"TaxDirectiveNumber": "string",							// Type="String"
	"IT3AReason": 0,										// Type="Int32"
	"EmploymentAction": "string", 							// Type="String"
	"TerminationCompanyRun": "string",						// Type="String"
	"IdentityType": "string",								// Type="String"
	"IdNumber": "string",									// Type="String"
	"PassportNumber": "string",								// Type="String"
	"PercentageAmount": "string", 							// Type="String"
	"Amount": 0,											// Type="Decimal"
	"Percentage": 0,										// Type="Decimal"
	"DeemedMonthlyRemuneration": 0,							// Type="Decimal"
	"Deemed75Indicator": true,								// Type="Boolean"
	"DeemedRecoveryMonthly": true,							// Type="Boolean"
	"EncashLeave": true,									// Type="Boolean"
	"Irp30": true,											// Type="Boolean"
	"FinalizeIssueTaxCert": true,							// Type="Boolean"
	"PassportCountry": true,								// Type="String"
	"PassportIssued": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"PassportExpiry": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"PermitIssued": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
	"PermitExpiry": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
	"AdditionalDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"EmploymentCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
	"TerminationCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
	"TempWorker": true,										// Type="Boolean"
	"AdditionalDate1": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"NotReEmployable": true,								// Type="Boolean"
	"ReferenceNumber": "string",							// Type="String"
	"OldEmployeeId": 0,										// Type="Int64" - Creates a transfer record for ESS users to be able to access previous payslips
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

### Get a collection of employment statuses as of an effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus/effective/:effectivedate?$orderby={{$employment-status-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}`

Retrieves a list of employee employment Statuses as of the specififed effective date and based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employment-status-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |

**Example Response: Get a collection of of employment statuses as of an effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeEmploymentStatus",
    "@odata.count": 1,
    "value": [
				{
					"EmploymentStatusId": 0,								// Type="Int64"
					"EmployeeNumber": "string",								// Type="String"
					"GroupJoinDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"EmploymentDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"TerminationDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"TerminationReason": "string",							// Type="String"
					"TaxStatus": "string",									// Type="String"
					"TaxReferenceNumber": "string",							// Type="String"
					"NatureOfPerson": "string",								// Type="String"
					"TaxOffice": 0											// Type="Int32"
					"TaxDirectiveNumber": "string",							// Type="String"
					"IT3AReason": 0,										// Type="Int32"
					"EmploymentAction": "string", 							// Type="String"
					"TerminationCompanyRun": "string",						// Type="String"
					"IdentityType": "string",								// Type="String"
					"IdNumber": "string",									// Type="String"
					"PassportNumber": "string",								// Type="String"
					"PercentageAmount": "string", 							// Type="String"
					"Amount": 0,											// Type="Decimal"
					"Percentage": 0,										// Type="Decimal"
					"DeemedMonthlyRemuneration": 0,							// Type="Decimal"
					"Deemed75Indicator": true,								// Type="Boolean"
					"DeemedRecoveryMonthly": true,							// Type="Boolean"
					"EncashLeave": true,									// Type="Boolean"
					"Irp30": true,											// Type="Boolean"
					"FinalizeIssueTaxCert": true,							// Type="Boolean"
					"PassportCountry": true,								// Type="String"
					"PassportIssued": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"PassportExpiry": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"PermitIssued": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
					"PermitExpiry": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
					"AdditionalDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"EmploymentCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
					"TerminationCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
					"TempWorker": true,										// Type="Boolean"
					"AdditionalDate1": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"NotReEmployable": true,								// Type="Boolean"
					"ReferenceNumber": "string",							// Type="String"
					"OldEmployeeId": 0,										// Type="Int64" - Creates a transfer record for ESS users to be able to access previous payslips
                    "CustomFields": [		
                        {			
                        "Code": "string",				        // Type="String"
                        "Label": "string",			        // Type="String"
                        "Value": "string"			        	// Type="String"
                        }	
				    ]
				}
    ]
}
```

---

### Get a collection of all employment statuses

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus/all?$orderby={{$employment-status-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of all employee employment statuses based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employment-status-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of employment statuses** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeEmploymentStatus",
    "@odata.count": 1,
    "value": [
				{
					"EmploymentStatusId": 0,								// Type="Int64"
					"EmployeeNumber": "string",								// Type="String"
					"GroupJoinDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"EmploymentDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"TerminationDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"TerminationReason": "string",							// Type="String"
					"TaxStatus": "string",									// Type="String"
					"TaxReferenceNumber": "string",							// Type="String"
					"NatureOfPerson": "string",								// Type="String"
					"TaxOffice": 0											// Type="Int32"
					"TaxDirectiveNumber": "string",							// Type="String"
					"IT3AReason": 0,										// Type="Int32"
					"EmploymentAction": "string", 							// Type="String"
					"TerminationCompanyRun": "string",						// Type="String"
					"IdentityType": "string",								// Type="String"
					"IdNumber": "string",									// Type="String"
					"PassportNumber": "string",								// Type="String"
					"PercentageAmount": "string", 							// Type="String"
					"Amount": 0,											// Type="Decimal"
					"Percentage": 0,										// Type="Decimal"
					"DeemedMonthlyRemuneration": 0,							// Type="Decimal"
					"Deemed75Indicator": true,								// Type="Boolean"
					"DeemedRecoveryMonthly": true,							// Type="Boolean"
					"EncashLeave": true,									// Type="Boolean"
					"Irp30": true,											// Type="Boolean"
					"FinalizeIssueTaxCert": true,							// Type="Boolean"
					"PassportCountry": true,								// Type="String"
					"PassportIssued": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"PassportExpiry": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"PermitIssued": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
					"PermitExpiry": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
					"AdditionalDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"EmploymentCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
					"TerminationCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
					"TempWorker": true,										// Type="Boolean"
					"AdditionalDate1": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
					"NotReEmployable": true,								// Type="Boolean"
					"ReferenceNumber": "string",							// Type="String"
					"OldEmployeeId": 0,										// Type="Int64" - Creates a transfer record for ESS users to be able to access previous payslips
                    "CustomFields": [		
                        {			
                        "Code": "string",				        // Type="String"
                        "Label": "string",			        // Type="String"
                        "Value": "string"			        	// Type="String"
                        }	
				    ]
				}
    ]
}
```

---

### Create a single employment status record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus`

Create a single employee employment status record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeEmploymentStatus" in metadata endpoint for available fields
}
```

**Example Response: Create a single employment status record** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeEmploymentStatus/$entity"
	"EmploymentStatusId": 0,								// Type="Int64"
	"EmployeeNumber": "string",								// Type="String"
	"GroupJoinDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"EmploymentDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"TerminationDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"TerminationReason": "string",							// Type="String"
	"TaxStatus": "string",									// Type="String"
	"TaxReferenceNumber": "string",							// Type="String"
	"NatureOfPerson": "string",								// Type="String"
	"TaxOffice": 0											// Type="Int32"
	"TaxDirectiveNumber": "string",							// Type="String"
	"IT3AReason": 0,										// Type="Int32"
	"EmploymentAction": "string", 							// Type="String"
	"TerminationCompanyRun": "string",						// Type="String"
	"IdentityType": "string",								// Type="String"
	"IdNumber": "string",									// Type="String"
	"PassportNumber": "string",								// Type="String"
	"PercentageAmount": "string", 							// Type="String"
	"Amount": 0,											// Type="Decimal"
	"Percentage": 0,										// Type="Decimal"
	"DeemedMonthlyRemuneration": 0,							// Type="Decimal"
	"Deemed75Indicator": true,								// Type="Boolean"
	"DeemedRecoveryMonthly": true,							// Type="Boolean"
	"EncashLeave": true,									// Type="Boolean"
	"Irp30": true,											// Type="Boolean"
	"FinalizeIssueTaxCert": true,							// Type="Boolean"
	"PassportCountry": true,								// Type="String"
	"PassportIssued": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"PassportExpiry": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"PermitIssued": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
	"PermitExpiry": "2019-01-01T01:01:00.000Z",				// Type="DateTime"
	"AdditionalDate": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"EmploymentCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
	"TerminationCaptureDate": "2019-01-01T01:01:00.000Z",	// Type="DateTime"
	"TempWorker": true,										// Type="Boolean"
	"AdditionalDate1": "2019-01-01T01:01:00.000Z",			// Type="DateTime"
	"NotReEmployable": true,								// Type="Boolean"
	"ReferenceNumber": "string",							// Type="String"
	"OldEmployeeId": 0										// Type="Int64" - Creates a transfer record for ESS users to be able to access previous payslips
}
```

---

### Update a single employment status record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus({{EmploymentStatusId}})`

Update a single employee employment status record based on the specified `EmploymentStatusId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeEmploymentStatus" in metadata endpoint for available fields
}
```

**Example Response: Update a single employment status record** (Status: 204 No Content)

---

### EmploymentStatus: EmployeeTermination

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus({{EmploymentStatusId}})`

This is employee termination.

**Request Body:**

```json
{
    "TerminationDate": "2019-01-01",					// Required
    "TerminationReason": "string",						// Required
    "EmploymentAction": "string",						// Required
    "TerminationCompanyRun": "string",					// Optional
    "EncashLeave": true 								// Required
}
```

---

### EmploymentStatus: Reinstate Same Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus({{EmploymentStatusId}})`

This is to reinstate an employee with the same tax record. Where only the Employment action is required for reinstatement but any other changes will also be applied.

**Request Body:**

```json
{
    "EmploymentAction": "reinstate this employee resuming this tax record" //Required
}
```

---

### EmploymentStatus: Reinstate With New Tax Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus({{EmploymentStatusId}})`

This is to reinstate an employee with a new tax record. Where only the Employment action is required for reinstatement but any other changes will also be applied to the new tax record.

**Request Body:**

```json
{
    "EmploymentAction": "reinstate this employee starting a new tax record",	//Required
	"EmploymentDate": "2019-01-01"												//Required
}
```

---

### Delete a single employment status record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeEmploymentStatus({{EmploymentStatusId}})`

Retrieves employee employment status information based on the specified `EmploymentStatusId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single employment status record** (Status: 200 OK)

```json
{}
```

---
## Positions

### API v1.1

###### Create/Update

- To create or update an `EmployeePosition` record linked to a Organization Position, use the `OrganizationPositionWithCode` Field.
    
- The options for the field `OrganizationPositionWithCode` can be retrieved from the `OrganizationPositionWithCode` lookup.
    

> From API v1.1 the field `OrganizationPosition` is readonly. 
  

###### OrganizationGroups

- The `OrganizationGroups` field will no longer be returned by default and is no longer a string collection.
    
- To return the `OrganizationGroups` field data, use the $expand query string functionality. eg. $expand=OrganizationGroups
    
- The `OrganizationGroups` field is a readonly colection of `OrganizationUnit` records
    

### API v1.0

###### Create/Update

- To create or update a `EmployeePosition` record linked to a Organization Position, use the `OrganizationPosition` Field.
    
- The options for the field `OrganizationPosition` can be retrieved from the `OrganizationPosition` lookup.

### Get a collection of positions

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}&$expand=OrganizationGroups`

Retrieves a list of employee position records as of the today's date and based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq '1234'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$expand` | `OrganizationGroups` | Optional - to return the OrganizationGroups field data, use the value 'OrganizationGroups' |

**Example Response: Get a collection of positions** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePosition",
    "@odata.count": 10,
    "value": [
        {
        	"EmployeePositionId": 0,                            // Type="Int64"
			"EmployeeNumber": "string",                         // Type="String"
            "FullName": "string",                               // Type="String"
			"EffectiveDate": "2019-05-03T09:28:15.342Z",        // Type="DateTime"
			"OrganizationPosition": "string",                   // Type="String" - ReadOnly field
            "OrganizationPositionId": 0,                        // Type="Int64" - ReadOnly field
            "OrganizationPositionWithCode": "string",           // Type="String"
			"PositionType": "Permanent",						// Type="String"
			"Grade": "string",                                  // Type="String"
			"OccupationalLevel": "string",                       // Type="String"
			"DirectlyReportsPositionOverride": "string",        // Type="String"
			"DirectlyReportsPosition": "string",    			// Type="String"
			"OrganizationGroup": "string",                     	// Type="String"
            "OrganizationGroupDescription": "string",           // Type="String"
			"OrganizationGroups": [                             
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
                }
            ],                    // Type="String Array"
			"OrganizationRegion": "string",                     // Type="String"
			"PayPoint": "string",                               // Type="String"
			"DirectlyReportsEmployee": "string",                // Type="String"
			"DirectlyReportsEmployeeNumber": "string",          // Type="String"
			"EmploymentCategory": "string",                     // Type="String"
			"EmploymentSubCategory": "string",                  // Type="String"
			"Administrator": "string",                          // Type="String"
			"AdministratorEmployeeNumber": "string",            // Type="String"
			"WorkflowRole": "string",                           // Type="String"
			"GeneralLedger": "string",                          // Type="String"
			"TradeUnion": "string",                             // Type="String"
			"IsPromotion": true,                                // Type="Boolean"
			"Roster": "stri
... (truncated)
```

---

### Get a single position record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition({{EmployeePositionId}})`

Retrieves employee position records based on the specified `EmployeePositionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

---

### Get a collection of positions as of an effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition/effective/:effectivedate?$orderby={{$employee-position-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$expand=OrganizationGroups`

Retrieves a list of employee position records as of the specified effective date and based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `Authorization ` | `Bearer {{payspace-access-token}}` | The Authorization bearer token |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-position-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$expand` | `OrganizationGroups` | Optional - to return the OrganizationGroups field data, use the value 'OrganizationGroups' |

**Example Response: Get a collection of positions as of an effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePosition",
    "@odata.count": 10,
    "value": [
        {
        	"EmployeePositionId": 0,                            // Type="Int64"
			"EmployeeNumber": "string",                         // Type="String"
            "FullName": "string",                               // Type="String"
			"EffectiveDate": "2019-05-03T09:28:15.342Z",        // Type="DateTime"
			"OrganizationPosition": "string",                   // Type="String" - ReadOnly field
            "OrganizationPositionId": 0,                        // Type="Int64" - ReadOnly field
            "OrganizationPositionWithCode": "string",           // Type="String"
			"PositionType": "Permanent",						// Type="String"
			"Grade": "string",                                  // Type="String"
			"OccupationalLevel": "string",                      // Type="String"
			"DirectlyReportsPositionOverride": "string",        // Type="String"
			"DirectlyReportsPosition": "string",    			// Type="String"
			"OrganizationGroup": "string",                     	// Type="String"
            "OrganizationGroupDescription": "string",           // Type="String"
			"OrganizationGroups": [                             
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
                    "InactiveDate": "2019-01-01T00:00:00+02:00"	    // Type="DateTime"
                }
            ],   
			"OrganizationRegion": "string",                     // Type="String"
			"PayPoint": "string",                               // Type="String"
			"DirectlyReportsEmployee": "string",                // Type="String"
			"DirectlyReportsEmployeeNumber": "string",          // Type="String"
			"EmploymentCategory": "string",                     // Type="String"
			"EmploymentSubCategory": "string",                  // Type="String"
			"Administrator": "string",                          // Type="String"
			"AdministratorEmployeeNumber": "string",            // Type="String"
			"WorkflowRole": "string",                           // Type="String"
			"GeneralLedger": "string",                          // Type="String"
			"TradeUnion": "string",                             // Type="String"
			"IsPromotion": true,                                // Type="Boolean"
			"Roster": "string",                                 
... (truncated)
```

---

### Get a collection of all employee positions

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition/all?$orderby={{employee-positions-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-positions-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee positions**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePosition",
    "value": [
        {
            "EmployeePositionId": 1,                                // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "OrganizationPosition": "string",                       // Type="String"
            "OrganizationPositionId": 1,                            // Type="Int64"
            "OrganizationPositionWithCode": "string",               // Type="String"
            "PositionType": "string",                               // Type="String"
            "Grade": "string",                                      // Type="String"
            "OccupationalLevel": "string",                          // Type="String"
            "DirectlyReportsPositionOverride": "string",            // Type="String"
            "DirectlyReportsPosition": "string",                    // Type="String"
            "OrganizationGroup": "string",                          // Type="String"
            "OrganizationGroupDescription": "string",               // Type="String"
            "OrganizationGroups": [],                               // Type="String"
            "OrganizationRegion": "string",                         // Type="String"
            "PayPoint": "string",                                   // Type="String"
            "DirectlyReportsEmployee": "string",                    // Type="String"
            "DirectlyReportsEmployeeNumber": "string",              // Type="String"
            "EmploymentCategory": "string",                         // Type="String"
            "EmploymentSubCategory": "string",                      // Type="String"
            "Administrator": "string",                              // Type="String"
            "AdministratorEmployeeNumber": "string",                // Type="String"
            "WorkflowRole": "string",                               // Type="String"
            "GeneralLedger": "string",                              // Type="String"
            "TradeUnion": "string",                                 // Type="String"
            "IsPromotion": false,                                   // Type="Boolean"
            "KeepJob": false,                                       // Type="Boolean"
            "Roster": "string",                                     // Type="String"
            "Job": "string",                                        // Type="String"
            "Comments": "string",                                   // Type="String"
            "AltPositionName": "string",                            // Type="String"
            "PositionEffectiveDate": "2019-01-01T00:00:00+02:00",   // Type="DateTime"
            "CustomTradeUnion": 
... (truncated)
```

---

### Create a single position record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition`

Create a single employee position record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeePosition" in metadata endpoint for available fields
}
```

**Example Response: Create a single position record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeePosition/$entity",
    "EmployeePositionId": 0,                            // Type="Int64"
    "EmployeeNumber": "string",                         // Type="String"
    "EffectiveDate": "2019-05-03T09:28:15.342Z",        // Type="DateTime"
    "OrganizationPosition": "string",                   // Type="String" - ReadOnly field
    "OrganizationPositionWithCode": "string",           // Type="String"
    "PositionType": "Permanent",						// Type="String"
    "Grade": "string",                                  // Type="String"
    "OccupationalLevel": "string",                      // Type="String"
    "DirectlyReportsPositionOverride": "string",        // Type="String"
    "DirectlyReportsPosition": "string",    			// Type="String"
    "OrganizationGroup": "string",                     	// Type="String"
    "OrganizationGroups": [],                           
    "OrganizationRegion": "string",                     // Type="String"
    "PayPoint": "string",                               // Type="String"
    "DirectlyReportsEmployee": "string",                // Type="String"
    "DirectlyReportsEmployeeNumber": "string",          // Type="String"
    "EmploymentCategory": "string",                     // Type="String"
    "EmploymentSubCategory": "string",                  // Type="String"
    "Administrator": "string",                          // Type="String"
    "AdministratorEmployeeNumber": "string",            // Type="String"
    "WorkflowRole": "string",                           // Type="String"
    "GeneralLedger": "string",                          // Type="String"
    "TradeUnion": "string",                             // Type="String"
    "IsPromotion": true,                                // Type="Boolean"
    "Roster": "string",                                 // Type="String"
    "Job": "string",                                    // Type="String"
    "Comments": "string",                               // Type="String"
    "AltPositionName": "string",                        // Type="String"
    "DateAdded": "2019-05-03T09:28:15.342Z",            // Type="DateTime"
    "PositionEffectiveDate": "2019-05-03T09:28:15.342Z", // Type="DateTime"
    "CustomTradeUnion": "string",                       // Type="String"
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

### Update a single position record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition({{EmployeePositionId}})`

Update a single employee position record based on the specified `EmployeePositionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeePosition" in metadata endpoint for available fields
}
```

**Example Response: Update a single position record** (Status: 204 No Content)

---

### Delete a single position record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition({{PositionId}})`

Retrieves employee position records based on the specified `PositionId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single position record** (Status: 200 OK)

```json
{}
```

---

### Extract all positions: Start

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition/extract/start`

Starts the process to create a file containing all `EmployeePosition` records

**Example Response: Extract all positions: Start**

```json
{
  "ExtractionId": "f3dc35b0-9f4c-4a42-b579-1b3cb33e8be1",
  "ExtractionType": "EmployeePosition",
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

### Extract all positions: Status

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeePosition/extract/:extractionId/status`

Returns the extraction process status with download url once completed

**Example Response: Extract all positions: Status**

```json
{
    "ExtractionId": "Guid",               // Type="Guid"
    "ExtractionType": "EmployeePosition", // Type="String"    
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
## Banking Details

Retrieve employee bank details information associated with the `company_id`.

### Get a collection of bank detail records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBankDetail?$orderby={{$bank-detail-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee bank details based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$bank-detail-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of bank detail records** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeBankDetail",
    "@odata.count": 0,
    "value": [
	    {
			"BankDetailId": 0,					// Type="Int64"
			"EmployeeNumber": "string",			// Type="String"
			"PaymentMethod": "string",			// Type="String"
			"SplitType": "string", 			    // Type="String"
			"BankAccountOwner": "string", 		// Type="String"
			"BankAccountOwnerName": "",			// Type="String"
			"AccountType": "string",	        // Type="String"
			"BankName": "string",				// Type="String"
			"BankBranchNo": "string",		    // Type="String"
			"BankAccountNo": "string",	        // Type="String"
			"Reference": "string",				// Type="String"
			"IsMainAccount": true, 			    // Type="Boolean"
			"Amount": 100,				   	    // Type="Double"
			"Comments": "string",				// Type="String"
			"SwiftCode": "string", 				// Type="Boolean"
			"RoutingCode": "string", 			// Type="Boolean"
			"ComponentId": 0					// Type="Int64"
		}
	  ]
}
```

---

### Get a single bank detail record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBankDetail({{BankDetailId}})`

Retrieves employee bank detail information based on the specified `BankDetailId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single bank detail record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeBankDetail/$entity",	
	"BankDetailId": 0,					// Type="Int64"
	"EmployeeNumber": "string",			// Type="String"
	"PaymentMethod": "string",			// Type="String"
	"SplitType": "string", 			    // Type="String"
	"BankAccountOwner": "string", 		// Type="String"
	"BankAccountOwnerName": "",			// Type="String"
	"AccountType": "string",	        // Type="String"
	"BankName": "string",				// Type="String"
	"BankBranchNo": "string",		    // Type="String"
	"BankAccountNo": "string",	        // Type="String"
	"Reference": "string",				// Type="String"
	"IsMainAccount": true, 			    // Type="Boolean"
	"Amount": 100,				   	    // Type="Double"
	"Comments": "string",				// Type="String"
	"SwiftCode": "string", 				// Type="Boolean"
	"RoutingCode": "string", 			// Type="Boolean"
	"ComponentId": 0					// Type="Int64"
}
```

---

### Create a single bank detail record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBankDetail`

Create a single employee bank details record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeBankDetail" in metadata endpoint for available fields
}
```

**Example Response: Create a single bank detail** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeBankDetail/$entity",
	"BankDetailId": 0,					// Type="Int64"
	"EmployeeNumber": "string",			// Type="String"
	"PaymentMethod": "string",			// Type="String"
	"SplitType": "string", 			    // Type="String"
	"BankAccountOwner": "string", 		// Type="String"
	"BankAccountOwnerName": "",			// Type="String"
	"AccountType": "string",	        // Type="String"
	"BankName": "string",				// Type="String"
	"BankBranchNo": "string",		    // Type="String"
	"BankAccountNo": "string",	        // Type="String"
	"Reference": "string",				// Type="String"
	"IsMainAccount": true, 			    // Type="Boolean"
	"Amount": 100,				   	    // Type="Double"
	"Comments": "string",				// Type="String"
	"SwiftCode": "string", 				// Type="Boolean"
	"RoutingCode": "string", 			// Type="Boolean"
	"ComponentId": 0					// Type="Int64"
}
```

---

### Update a single bank detail record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBankDetail({{BankDetailId}})`

Update a single bank detail record based on the specified `BankdetailId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeBankDetail" in metadata endpoint for available fields
}
```

**Example Response: Update a single bank detail record** (Status: 204 No Content)

---

### Delete a single bank detail record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeBankDetail({{BankDetailId}})`

Retrieves employee bank detail information based on the specified `BankDetailId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single bank detail record** (Status: 200 OK)

```json
{}
```

---
## Dependants

Retrieve dependant information for all employees associated with the `company_id`.

### Get a collection of dependants

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDependant?$orderby={{$employee-dependant-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee dependants based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-dependant-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of Dependants** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeDependant",
    "@odata.count": 102,
    "value": [
        {
			"EmployeeDependantId": 0 ,                       // Type="Int64"
			"EmployeeNumber": "string",                      // Type="String"
			"AdultRelationship": "LegalSpouse",              // Type="String"
			"DependantType": "Adult",                        // Type="String"
			"ChildBirthDate": "2019-05-03T08:06:38.470Z",    // Type="DateTime"
			"MedicalAidDependant": true,                     // Type="Boolean"
			"FirstName": "string",                           // Type="String"
			"LastName": "string",                            // Type="String"
			"ContactNumber": "string",                       // Type="String"
			"MaritalStatus": "Re-Married",                   // Type="String"
			"MaritalDate": "2019-05-03T08:06:38.470Z",       // Type="DateTime"
			"WorkingSpouse": true,                           // Type="Boolean"
			"IsBreadWinner": true,                           // Type="Boolean"
			"RestrictEmployerContribution": true,            // Type="Boolean"
			"EffectiveDate": "2019-05-03T08:06:38.470Z",     // Type="DateTime"
			"InactiveDate": "2019-05-03T08:06:38.470Z",      // Type="DateTime"
			"CustomFields": [
			  {
			    "Code": "string",                            // Type="String"
			    "Label": "string",                           // Type="String"
			    "Value": "string",                           // Type="String"
			  }
			]
		}
    ]
}
```

---

### Get a single dependant record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDependant({{EmployeeDependantId}})`

Retrieves dependant information based on the specified `DependentId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single Dependant record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeDependant/$entity",
	"EmployeeDependantId": 0 ,                       // Type="Int64"
	"EmployeeNumber": "string",                      // Type="String"
	"AdultRelationship": "LegalSpouse",              // Type="String"
	"DependantType": "Adult",                        // Type="String"
	"ChildBirthDate": "2019-05-03T08:06:38.470Z",    // Type="DateTime"
	"MedicalAidDependant": true,                     // Type="Boolean"
	"FirstName": "string",                           // Type="String"
	"LastName": "string",                            // Type="String"
	"ContactNumber": "string",                       // Type="String"
	"MaritalStatus": "Re-Married",                   // Type="String"
	"MaritalDate": "2019-05-03T08:06:38.470Z",       // Type="DateTime"
	"WorkingSpouse": true,                           // Type="Boolean"
	"IsBreadWinner": true,                           // Type="Boolean"
	"RestrictEmployerContribution": true,            // Type="Boolean"
	"EffectiveDate": "2019-05-03T08:06:38.470Z",     // Type="DateTime"
	"InactiveDate": "2019-05-03T08:06:38.470Z",      // Type="DateTime"
	"CustomFields": [
	  {
	    "Code": "string",                            // Type="String"
	    "Label": "string",                           // Type="String"
	    "Value": "string",                           // Type="String"
	  }
	]
}
```

---

### Get a collection of all employee dependants

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDependant/all?$orderby={{employee-dependants-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-dependants-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee dependants**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeDependant",
    "value": [
        {
            "EmployeeDependantId": 1,                               // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "AdultRelationship": "string",                          // Type="String"
            "DependantType": "string",                              // Type="String"
            "ChildBirthDate": "2019-01-01T00:00:00+02:00",          // Type="DateTIme"
            "MedicalAidDependant": false,                           // Type="Boolean"
            "FirstName": "string",                                  // Type="String"
            "LastName": "string",                                   // Type="String"
            "ContactNumber": "string",                              // Type="String"
            "MaritalStatus": "string",                              // Type="String"
            "MaritalDate": "2019-01-01T00:00:00+02:00",             // Type="DateTIme"
            "WorkingSpouse": false,                                 // Type="Boolean"
            "IsBreadWinner": true,                                  // Type="Boolean"
            "RestrictEmployerContribution": false,                  // Type="Boolean"
            "ChildMedicalAidCalculations": false,                   // Type="Boolean"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTIme"
            "InactiveDate": "2019-01-01T00:00:00+02:00",            // Type="DateTIme"
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

### Create a single dependant record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDependant`

Create a single employee dependant record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeDependant" in metadata endpoint for available fields
}
```

**Example Response: Create a single dependant record** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeDependant/$entity",
	"EmployeeDependantId": 0 ,                       // Type="Int64"
	"EmployeeNumber": "string",                      // Type="String"
	"AdultRelationship": "LegalSpouse",              // Type="String"
	"DependantType": "Adult",                        // Type="String"
	"ChildBirthDate": "2019-05-03T08:06:38.470Z",    // Type="DateTime"
	"MedicalAidDependant": true,                     // Type="Boolean"
	"FirstName": "string",                           // Type="String"
	"LastName": "string",                            // Type="String"
	"ContactNumber": "string",                       // Type="String"
	"MaritalStatus": "Re-Married",                   // Type="String"
	"MaritalDate": "2019-05-03T08:06:38.470Z",       // Type="DateTime"
	"WorkingSpouse": true,                           // Type="Boolean"
	"IsBreadWinner": true,                           // Type="Boolean"
	"RestrictEmployerContribution": true,            // Type="Boolean"
	"EffectiveDate": "2019-05-03T08:06:38.470Z",     // Type="DateTime"
	"InactiveDate": "2019-05-03T08:06:38.470Z",      // Type="DateTime"
	"CustomFields": [
	  {
	    "Code": "string",                            // Type="String"
	    "Label": "string",                           // Type="String"
	    "Value": "string",                           // Type="String"
	  }
	]
}
```

---

### Update a single dependant record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDependant({{EmployeeDependantId}})`

Update a single dependant record based on the specified `EmployeeDependantId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeDependant" in metadata endpoint for available fields
}
```

**Example Response: Update a single dependant record** (Status: 204 No Content)

---

### Delete a single dependant record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeDependant({{EmployeeDependantId}})`

Retrieves dependant information based on the specified `DependentId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single Dependant record** (Status: 200 OK)

```json
{}
```

---
## Attachments

Retrieve employee attachment information associated with the `company_id`.

### Get a collection of employee attachment records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAttachment?$orderby={{$attachment-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of active employee attachment records based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$attachment-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of employee attachment records** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAttachment",
	"@odata.count": 0,
	"value": [
		{
			"AttachmentId": 0,										// Type="Int64"
			"EmployeeNumber": "string",			        			// Type="String"
			"Classification": "string",   							// Type="String"
			"AttachmentDescription": "string",			        	// Type="String"
			"AttachmentName": "string",			        			// Type="String"
			"AttachmentUrl": "string",			        			// Type="String"
			"Attachment": null,										// Type="base64 byte[]"
            "CustomFields": [		
                {			
                    "Code": "string",				        // Type="String"
                    "Label": "string",			        // Type="String"
                    "Value": "string"			        	// Type="String"
                }	
            ]
		}
	]
}
```

---

### Get a single employee attachment record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAttachment({{AttachmentId}})`

Retrieves employee attachment information based on the specified `AttachmentId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single employee attachment record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAttachment/$entity",	
	"AttachmentId": 0,										// Type="Int64"
	"EmployeeNumber": "string",			        			// Type="String"
	"Classification": "string",   							// Type="String"
	"AttachmentDescription": "string",			        	// Type="String"
	"AttachmentName": "string",			        			// Type="String"
	"AttachmentUrl": "string",			        			// Type="String"
	"Attachment": null,										// Type="base64 byte[]"
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

### Get a collection of all employee attachments

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAttachment/all?$orderby={{employee-attachment-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-attachment-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee attachment**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAttachment",
    "value": [
        {
            "AttachmentId": 0,                                      // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "Classification": "string",                             // Type="String"
            "AttachmentDescription": "string",                      // Type="String"
            "AttachmentName": "string",                             // Type="String"
            "AttachmentUrl": "string",                              // Type="String"
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

### Create a single employee attachment record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAttachment`

Create a single attachment record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeAttachment" in metadata endpoint for available fields
}
```

**Example Response: Create a single employee attachment record** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAttachment/$entity",
	"AttachmentId": 0,										// Type="Int64"
	"EmployeeNumber": "string",			        			// Type="String"
	"Classification": "string",   							// Type="String"
	"AttachmentDescription": "string",			        	// Type="String"
	"AttachmentName": "string",			        			// Type="String"
	"AttachmentUrl": "string",			        			// Type="String"
	"Attachment": null,										// Type="base64 byte[]"
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

### Update a single employee attachment record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAttachment({{AttachmentId}})`

Update a single attachment record based on the specified `AttachmentId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeAttachment" in metadata endpoint for available fields
}
```

**Example Response: Update a single employee attachment record** (Status: 204 No Content)

---

### Delete a single employee attachment record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAttachment({{AttachmentId}})`

Delete a single employee attachment record based on the specified `AttachmentId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single employee attachment record** (Status: 200 OK)

---
## Project Details

Retrieve project information for all employees associated with the `company_id`.

### Get a collection of projects

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeProject?$orderby={{$employee-project-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee projects based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-project-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of Projects** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeProject",
    "@odata.count": 102,
    "value": [
        {
			"EmployeeProjectId": 0,							// Type="Int64"
			"EmployeeNumber": "string",						// Type="String"
			"ProjectCode": "string",						// Type="String"
			"ProjectActivityCode": "string",				// Type="String"
			"EffectiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
			"EndDate": "2019-01-01T00:00:00+02:00",			// Type="DateTime"
			"ProjectBased": false,							// Type="Boolean"
            "ReminderOption": "string",                     // Type="String"
			"PaidByInvoice": false,							// Type="Boolean"
			"PayFrequency": "string",						// Type="String"
			"ProjectPayRate": 0,							// Type="Decimal"
			"HoursPerDay": 0,								// Type="Decimal"
			"DaysPerMonth": 0,								// Type="Decimal"
			"Comments": "string",							// Type="String"
            "CustomFields": [		
                {			
                    "Code": "string",				        // Type="String"
                    "Label": "string",			        // Type="String"
                    "Value": "string"			        	// Type="String"
                }	
            ]
		}
    ]
}
```

---

### Get a collection of projects as of an effective date

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeProject/effective/:effectivedate?$orderby={{$employee-project-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee projects based on the specified `company_id`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-project-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of projects as of an effective date** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeProject",
    "@odata.count": 102,
    "value": [
        {
			"EmployeeProjectId": 0,							// Type="Int64"
			"EmployeeNumber": "string",						// Type="String"
			"ProjectCode": "string",						// Type="String"
			"ProjectActivityCode": "string",				// Type="String"
			"EffectiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
			"EndDate": "2019-01-01T00:00:00+02:00",			// Type="DateTime"
			"ProjectBased": false,							// Type="Boolean"
            "ReminderOption": "string",                     // Type="String"
			"PaidByInvoice": false,							// Type="Boolean"
			"PayFrequency": "string",						// Type="String"
			"ProjectPayRate": 0,							// Type="Decimal"
			"HoursPerDay": 0,								// Type="Decimal"
			"DaysPerMonth": 0,								// Type="Decimal"
			"Comments": "string",							// Type="String"
            "CustomFields": [		
                {			
                    "Code": "string",				        // Type="String"
                    "Label": "string",			        // Type="String"
                    "Value": "string"			        	// Type="String"
                }	
            ]
		}
    ]
}
```

---

### Get a single project record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeProject({{EmployeeProjectId}})`

Retrieves project information based on the specified `EmployeeProjectId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get a single project record** (Status: 200 OK)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeProject/$entity",
	"EmployeeProjectId": 0,							// Type="Int64"
	"EmployeeNumber": "string",						// Type="String"
	"ProjectCode": "string",						// Type="String"
	"ProjectActivityCode": "string",				// Type="String"
	"EffectiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
	"EndDate": "2019-01-01T00:00:00+02:00",			// Type="DateTime"
	"ProjectBased": false,							// Type="Boolean"
    "ReminderOption": "string",                     // Type="String"
	"PaidByInvoice": false,							// Type="Boolean"
	"PayFrequency": "string",						// Type="String"
	"ProjectPayRate": 0,							// Type="Decimal"
	"HoursPerDay": 0,								// Type="Decimal"
	"DaysPerMonth": 0,								// Type="Decimal"
	"Comments": "string",							// Type="String"
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

### Get a collection of all projects

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeProject/all?$orderby={{employee-project-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-project-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all projects**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeProject",
    "value": [
        {
            "EmployeeProjectId": 1,                                 // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "ProjectCode": "string",                                // Type="String"
            "ProjectDescription": "string",                         // Type="String"
            "ProjectActivityCode": "string",                        // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "EndDate": "2019-01-01T12:00:00.00+02:00",              // Type="DateTime"
            "ProjectBased": false,                                  // Type="Boolean"
            "ReminderOption": "string",                             // Type="String"
            "PaidByInvoice": false,                                 // Type="Boolean"
            "PayFrequency": "string",                               // Type="String"
            "ProjectPayRate": 1.00000,                              // Type="Decimal"
            "HoursPerDay": 1.00000,                                 // Type="Decimal"
            "DaysPerMonth": 1.00000,                                // Type="Decimal"
            "Comments": "string",                                   // Type="String"
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

### Create a single project record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeProject`

Create a single employee project record for the specified `EmployeeNumber`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeProject" in metadata endpoint for available fields
}
```

**Example Response: Create a single project record** (Status: 201 Created)

```json
{
	"@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeProject/$entity",
	"EmployeeProjectId": 0,							// Type="Int64"
	"EmployeeNumber": "string",						// Type="String"
	"ProjectCode": "string",						// Type="String"
	"ProjectActivityCode": "string",				// Type="String"
	"EffectiveDate": "2019-01-01T00:00:00+02:00",	// Type="DateTime"
	"EndDate": "2019-01-01T00:00:00+02:00",			// Type="DateTime"
	"ProjectBased": false,							// Type="Boolean"
    "ReminderOption": "string",                     // Type="String"
	"PaidByInvoice": false,							// Type="Boolean"
	"PayFrequency": "string",						// Type="String"
	"ProjectPayRate": 0,							// Type="Decimal"
	"HoursPerDay": 0,								// Type="Decimal"
	"DaysPerMonth": 0,								// Type="Decimal"
	"Comments": "string",							// Type="String"
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

### Update a single project record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeProject({{EmployeeProjectId}})`

Update a single project record based on the specified `EmployeeProjectId`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
  //see "EmployeeProject" in metadata endpoint for available fields
}
```

**Example Response: Update a single project record** (Status: 204 No Content)

---

### Delete a single project record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeProject({{EmployeeProjectId}})`

Retrieves project information based on the specified `EmployeeProjectId`.

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete a single project record** (Status: 200 OK)

```json
{}
```

---
## Asset

Retrieve asset information for all employees associated with the `company_id`.

### Get collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAsset?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq '001'". |
| `$top` | `{{$top}}` | Optional - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAsset",
     "value": [
        {
            "AssetId": 0,                                   // Type="Int64"
            "EmployeeNumber": "string",                     // Type="String"
            "FullName": "string",                           // Type="String"
            "Description": "string",                        // Type="String"
            "SerialNo": "string",                           // Type="String"
            "IssuedDate": "2018-10-01T00:00:00+02:00",      // Type="DateTime"
            "AssetStatus": 1,                               // Type="Enum"
            "ReturnDate": "2019-10-01T00:00:00+02:00",      // Type="DateTime"
            "Quantity": 1,                                  // Type="Int32"
            "RandValue": 10.00,                             // Type="Decimal"
            "StockRoomNo": "string",                        // Type="String"    
            "Model": "string",                              // Type="String"
            "Notes": "string",                              // Type="String"
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

### Get single record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAsset({{AssetId}})`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `` | `` |  |

**Example Response: Get single record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAsset/$entity",
    "AssetId": 0,                                   // Type="Int64"
    "EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String"
    "Description": "string",                        // Type="String"
    "SerialNo": "string",                           // Type="String"
    "IssuedDate": "2018-10-01T00:00:00+02:00",      // Type="DateTime"
    "AssetStatus": 1,                               // Type="Enum"
    "ReturnDate": "2019-10-01T00:00:00+02:00",      // Type="DateTime"
    "Quantity": 1,                                  // Type="Int32"
    "RandValue": 10.00,                             // Type="Decimal"
    "StockRoomNo": "string",                        // Type="String"    
    "Model": "string",                              // Type="String"
    "Notes": "string",                              // Type="String"
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

### Get a collection of all employee assets

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAsset/all?$orderby={{employee-asset-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-asset-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of all employee assets**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAsset",
    "value": [
        {
            "AssetId": 0,                                           // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "Description": "string",                                // Type="String"
            "SerialNo": "string",                                   // Type="String"
            "IssuedDate": "2019-01-01T00:00:00+02:00",              // Type="DateTime"
            "AssetStatus": "string",                                // Type="String"
            "ReturnDate": "2019-01-01T00:00:00+02:00",              // Type="DateTime"
            "Quantity": 0,                                          // Type="Int32"
            "RandValue": null,                                      // Type="Int32"
            "StockRoomNo": "string",                                // Type="String"
            "Model": "string",                                      // Type="String"
            "Notes": "string",                                      // Type="String"
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

### Create record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAsset`

**Example Response: Create record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAsset/$entity",
    "EmployeeNumber": "string",                     // Type="String"
    "FullName": "string",                           // Type="String"
    "Description": "string",                        // Type="String"
    "SerialNo": "string",                           // Type="String"
    "IssuedDate": "2018-10-01T00:00:00+02:00",      // Type="DateTime"
    "AssetStatus": 1,                               // Type="Enum"
    "ReturnDate": "2019-10-01T00:00:00+02:00",      // Type="DateTime"
    "Quantity": 1,                                  // Type="Int32"
    "RandValue": 10.00,                             // Type="Decimal"
    "StockRoomNo": "string",                        // Type="String"    
    "Model": "string",                              // Type="String"
    "Notes": "string",                              // Type="String"
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

### Update single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAsset({{AssetId}})`

**Example Response: Update single record** (Status: 204 No Content)

---

### Delete single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeAsset({{AssetId}})`

**Example Response: Delete single record** (Status: 200 OK)

```json
{
  "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeAsset/$entity",
  "Success": true
}
```

---
## Employee Custom Forms

Working with custom forms defined for use on employee level

List of available custom fields can be found on the customfields endpoint using the customFormCategory (see `Custom form custom fields` lookup)


### Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCustomForm?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `{{$top}}` |  |
| `$skip` | `{{$skip}}` |  |
| `$count` | `{{$count}}` |  |
| `$filter` | `{{$filter}}` |  |

**Example Response: Get Collection**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeCustomForm/",
    "@odata.count": 1,
    "value":[{
        "CustomFormId": 0,                              // Type="int64"
        "EmployeeNumber": "Emp1",                       // Type="String"
        "CustomFormCategoryCode": "Test",               // Type="String"
        "EffectiveDate": "2022-07-01",                  // Type="DateTime"
        "CustomFields": [                               // CustomField
            // List of custom fields relevent to custom form category
        ]
    }]
}
```

---

### Get Collection by category

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCustomForm/:category?$top=10`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `10` |  |

**Example Response: Get Collection by category**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeCustomForm/",
    "@odata.count": 1,
    "value":[{
        "CustomFormId": 0,                              // Type="int64"
        "EmployeeNumber": "Emp1",                       // Type="String"
        "CustomFormCategoryCode": "Test",               // Type="String"
        "EffectiveDate": "2022-07-01",                  // Type="DateTime"
        "CustomFields": [                               // CustomField
            // List of custom fields relevent to custom form category
        ]
    }]
}
```

---

### Get a collection of all employee custom forms

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCustomForm/all?$orderby={{employee-custom-form-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-custom-form-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a collection of all employee custom forms**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeCustomForm",
    "value": [
        {
            "CustomFormId": 1,                                      // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "CustomFormCategoryCode": "string",                     // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "ParentCustomFormId": 1,                                // Type="Int64"
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

### Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCustomForm({{CustomFormId}})`

**Example Response: Get Single** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeCustomForm/",
    "CustomFormId": 0,                              // Type="int64"
    "EmployeeNumber": "Emp1",                       // Type="String"
    "CustomFormCategoryCode": "Test",               // Type="String"
    "EffectiveDate": "2022-07-01",                  // Type="DateTime"
    "CustomFields": [                               // CustomField
        // List of custom fields relevent to custom form category
    ]
}
```

---

### Post

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCustomForm`

**Request Body:**

```json
{
    "EmployeeNumber": "Emp1",                       // Type="String"
    "CustomFormCategoryCode": "Test",                   // Type="String"
    "EffectiveDate": "2022-07-01",                  // Type="DateTime"
    "CustomFields": [                               // CustomField
        // See list of fields available in customfields endpoint for the required customFormCategory, (see customfield with category)
    ]
}
```

**Example Response: Post** (Status: 201 Created)

---

### PATCH

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCustomForm({{CustomFormId}})`

**Request Body:**

```json
{
  "CustomFields": [
    {
      "Code": "01",
      "Value": "123456789"
    }
  ]
}
```

**Example Response: PATCH** (Status: 200 OK)

---

### Delete

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeCustomForm({{CustomFormId}})`

**Example Response: Delete** (Status: 200 OK)

---
## Employee Inbox

Retrieves a list of active `EmployeeInbox` records based on the specified `company_id`.

### Get Collection

**`GET`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeInbox?$filter={{$filter}}`

Retrieves a list of active `EmployeeInbox` records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |

**Example Response: Get Collection**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeInbox",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeInboxId": 0,                   // Type="Int64"
            "EmployeeNumber": "string",             // Type="String"
            "EmployeeFullName": "string",           // Type="String"
            "ApplicantEmployeeNumber": "string",    // Type="String"
            "ApplicantFullName": "string",          // Type="String"
            "OnBehalfOfEmployeeNumber": "string",   // Type="String"
            "OnBehalfOfEmployeeFullName": "string", // Type="String"
            "ClaimBatchId": 0,                      // Type="Int64"
            "Subject": "string",                    // Type="String"
            "InboxEntryType": "string",             // Type="String"
            "UserWorkflowStepId": 0,                // Type="Int64"
            "EffectiveDate": 0                      // Type="DateTime"
        }
    ]
}
```

---
## Incident Management

Retrieve incident information for all employees associated with the `company_id`.

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncident?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeIncident` records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeIncident",
    "value": [
        {
            "EmployeeIncidentId": 0,                                        // Type="Int64"
            "EmployeeNumber": "string",                                     // Type="String"
            "FullName": "string",                                           // Type="String"
            "LinkedIncidentId": 0,                                          // Type="Int64"
            "IncidentType": 1,                                              // Type="Enum" see: IncidentType in metadata
            "CompanyOffence": "string",                                     // Type="String"
            "OffenceOutcome": "string",                                     // Type="String"
            "OffenceCategory": "string",                                    // Type="String"
            "AppealReason": "string",                                       // Type="String"
            "CompanyLegalBody": "string",                                   // Type="String"
            "CompanyIncidentType": "string",                                // Type="String"
            "AppealOutcome": "string",                                      // Type="String" 
            "AwardFavour": "string",                                        // Type="String" 
            "OtherOutcome": "string",                                       // Type="String" 
            "SettlementReinstate": "string",                                // Type="String" 
            "IncidentDate": "2000-00-00T00:00:00+02:00",                    // Type="DateTime"
            "IncidentNotes": "string",                                      // Type="String" 
            "EmployeeRepresentative": "string",                             // Type="String" 
            "CompanyRepresentative": "string",                              // Type="String" 
            "OutcomeDetails": "string",                                     // Type="String" 
            "DateCharged": "2000-00-00T00:00:00+02:00",                     // Type="DateTime"
            "HearingDate": "2000-00-00T00:00:00+02:00",                     // Type="DateTime"
            "Chairperson": "string",                                        // Type="String" 
            "CompanyWitnesses": "string",                                   // Type="String" 
            "EmployeeWitnesses": "string",                                  // Type="String" 
            "ValidityDate": "2000-00-00T00:00:00+02:00",                    // Type="DateTime"
            "CaseRefNumber": "string",                                      // Type="String" 
            "SRDescription": "string",                                      // Type="String"
            "SRMonetaryValue": 0,                                           // Type="Decimal"
            "SRReinstatementDate": "2000-00-00T00:00:00+02:00",             // Type="DateTime"
     
... (truncated)
```

---

### Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncident({{IncidentId}})`

Retrieves `EmployeeIncident` information based on the specified `IncidentId`.

**Example Response: Get Single** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeIncident",
    "EmployeeIncidentId": 0,                                        // Type="Int64"
    "EmployeeNumber": "string",                                     // Type="String"
    "FullName": "string",                                           // Type="String"
    "LinkedIncidentId": 0,                                          // Type="Int64"
    "IncidentType": 1,                                              // Type="Enum" see: IncidentType in metadata
    "CompanyOffence": "string",                                     // Type="String"
    "OffenceOutcome": "string",                                     // Type="String"
    "OffenceCategory": "string",                                    // Type="String"
    "AppealReason": "string",                                       // Type="String"
    "CompanyLegalBody": "string",                                   // Type="String"
    "CompanyIncidentType": "string",                                // Type="String"
    "AppealOutcome": "string",                                      // Type="String" 
    "AwardFavour": "string",                                        // Type="String" 
    "OtherOutcome": "string",                                       // Type="String" 
    "SettlementReinstate": "string",                                // Type="String" 
    "IncidentDate": "2000-00-00T00:00:00+02:00",                    // Type="DateTime"
    "IncidentNotes": "string",                                      // Type="String" 
    "EmployeeRepresentative": "string",                             // Type="String" 
    "CompanyRepresentative": "string",                              // Type="String" 
    "OutcomeDetails": "string",                                     // Type="String" 
    "DateCharged": "2000-00-00T00:00:00+02:00",                     // Type="DateTime"
    "HearingDate": "2000-00-00T00:00:00+02:00",                     // Type="DateTime"
    "Chairperson": "string",                                        // Type="String" 
    "CompanyWitnesses": "string",                                   // Type="String" 
    "EmployeeWitnesses": "string",                                  // Type="String" 
    "ValidityDate": "2000-00-00T00:00:00+02:00",                    // Type="DateTime"
    "CaseRefNumber": "string",                                      // Type="String" 
    "SRDescription": "string",                                      // Type="String"
    "SRMonetaryValue": 0,                                           // Type="Decimal"
    "SRReinstatementDate": "2000-00-00T00:00:00+02:00",             // Type="DateTime"
    "SRConditions": "string",                                       // Type="String"
    "CommissionerName": "string",                                   // Type="String"
    "LegalRepresentative": "string",                                // Type="String"
    "Settlemen
... (truncated)
```

---

### Create record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncident`

Create a single `EmployeeIncident` record for the specified `EmployeeNumber`.

**Request Body:**

```json
{
  //see "EmployeeIncident" in metadata endpoint for available fields
}
```

**Example Response: Create record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeIncident",
    "EmployeeIncidentId": 0,                                        // Type="Int64"
    "EmployeeNumber": "string",                                     // Type="String"
    "FullName": "string",                                           // Type="String"
    "LinkedIncidentId": 0,                                          // Type="Int64"
    "IncidentType": 1,                                              // Type="Enum" see: IncidentType in metadata
    "CompanyOffence": "string",                                     // Type="String"
    "OffenceOutcome": "string",                                     // Type="String"
    "OffenceCategory": "string",                                    // Type="String"
    "AppealReason": "string",                                       // Type="String"
    "CompanyLegalBody": "string",                                   // Type="String"
    "CompanyIncidentType": "string",                                // Type="String"
    "AppealOutcome": "string",                                      // Type="String" 
    "AwardFavour": "string",                                        // Type="String" 
    "OtherOutcome": "string",                                       // Type="String" 
    "SettlementReinstate": "string",                                // Type="String" 
    "IncidentDate": "2000-00-00T00:00:00+02:00",                    // Type="DateTime"
    "IncidentNotes": "string",                                      // Type="String" 
    "EmployeeRepresentative": "string",                             // Type="String" 
    "CompanyRepresentative": "string",                              // Type="String" 
    "OutcomeDetails": "string",                                     // Type="String" 
    "DateCharged": "2000-00-00T00:00:00+02:00",                     // Type="DateTime"
    "HearingDate": "2000-00-00T00:00:00+02:00",                     // Type="DateTime"
    "Chairperson": "string",                                        // Type="String" 
    "CompanyWitnesses": "string",                                   // Type="String" 
    "EmployeeWitnesses": "string",                                  // Type="String" 
    "ValidityDate": "2000-00-00T00:00:00+02:00",                    // Type="DateTime"
    "CaseRefNumber": "string",                                      // Type="String" 
    "SRDescription": "string",                                      // Type="String"
    "SRMonetaryValue": 0,                                           // Type="Decimal"
    "SRReinstatementDate": "2000-00-00T00:00:00+02:00",             // Type="DateTime"
    "SRConditions": "string",                                       // Type="String"
    "CommissionerName": "string",                                   // Type="String"
    "LegalRepresentative": "string",                                // Type="String"
    "Settlemen
... (truncated)
```

---

### Update single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncident({{IncidentId}})`

Update a single `EmployeeIncident` record for the specified `IncidentId`.

**Request Body:**

```json
{
  //see "EmployeeIncident" in metadata endpoint for available fields
}
```

**Example Response: Update single record** (Status: 204 No Content)

---

### Delete single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeIncident({{IncidentId}})`

Delete a single `EmployeeIncident` record for the specified `IncidentId`.

**Example Response: Delete single record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeIncident",
    "Success": true							                        // Type="boolean"
}
```

---

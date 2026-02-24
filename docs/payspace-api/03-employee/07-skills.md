# Employee > Skills

## Employee Training

Retrieve training information for all employees associated with the `company_id`.

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeTraining?$orderby={{$employee-dependant-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of active `EmployeeTraining` records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{$employee-dependant-field}}` | "EmployeeNumber eq 'Emp01'" - EmployeeNumber filter required |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeTraining",
    "value": [
        {
            "TrainingId": 0, 								// Type="Int64"
            "EmployeeNumber": "string",						// Type="String"
            "FullName": "string",							// Type="String"
            "EffectiveDate": "DateTimeOffset",				// Type="DateTimeOffset"
            "EndDate": "DateTimeOffset",					// Type="DateTimeOffset"
            "Course": "string",								// Type="String"
            "Status": "string",								// Type="String"
            "TrainingNotCompleteStatus": "string",			// Type="String"
            "CourseResult": 0.00000,						// Type="Decimal"
            "DirectCourseCost": 0.00000,					// Type="Decimal"
            "InvoiceNumber": "string",						// Type="String"
            "Facilitator": "string",						// Type="String"
            "Assessor": "string",							// Type="String"
            "CertificateNo": "string",						// Type="String"
            "CertificateExpiryDate": "DateTimeOffset",		// Type="DateTimeOffset"
            "EmployeeRate": 0.00000,						// Type="Decimal"
            "TravelCost": 0.00000,							// Type="Decimal"
            "Accommodation": 0.00000,						// Type="Decimal"
            "FacilitatorCost": 0.00000,						// Type="Decimal"
            "FoodBeverage": 0.00000,						// Type="Decimal"
            "Additional5": 0.00000,							// Type="Decimal"
            "Additional6": 0.00000,							// Type="Decimal"
            "Notes": "string",								// Type="String"
            "AttachmentName": "string",						// Type="String"
            "AttachmentUrl": "string",						// Type="String"
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

### Get Single

**`GET`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeTraining({{TrainingId}})`

Retrieves `EmployeeTraining` information based on the specified `TrainingId`.

**Example Response: Get Single** (Status: 200 OK)

```json
{

    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeTraining",
    "TrainingId": 0, 								// Type="Int64"
    "EmployeeNumber": "string",						// Type="String"
    "FullName": "string",							// Type="String"
    "EffectiveDate": "DateTimeOffset",				// Type="DateTimeOffset"
    "EndDate": "DateTimeOffset",					// Type="DateTimeOffset"
    "Course": "string",								// Type="String"
    "Status": "string",								// Type="String"
    "TrainingNotCompleteStatus": "string",			// Type="String"
    "CourseResult": 0.00000,						// Type="Decimal"
    "DirectCourseCost": 0.00000,					// Type="Decimal"
    "InvoiceNumber": "string",						// Type="String"
    "Facilitator": "string",						// Type="String"
    "Assessor": "string",							// Type="String"
    "CertificateNo": "string",						// Type="String"
    "CertificateExpiryDate": "DateTimeOffset",		// Type="DateTimeOffset"
    "EmployeeRate": 0.00000,						// Type="Decimal"
    "TravelCost": 0.00000,							// Type="Decimal"
    "Accommodation": 0.00000,						// Type="Decimal"
    "FacilitatorCost": 0.00000,						// Type="Decimal"
    "FoodBeverage": 0.00000,						// Type="Decimal"
    "Additional5": 0.00000,							// Type="Decimal"
    "Additional6": 0.00000,							// Type="Decimal"
    "Notes": "string",								// Type="String"
    "AttachmentName": "string",						// Type="String"
    "AttachmentUrl": "string",						// Type="String"
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

### Get a collection of all employee training records

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeTraining/all?$orderby={{employee-training-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-training-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee training records**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeTraining",
    "value": [
        {
            "TrainingId": 1,
            "EmployeeNumber": "string",                                         // Type="String"
            "FullName": "string",                                               // Type="String"
            "EffectiveDate": "2019-01-01T00:00:00+02:00",                       // Type="DateTime"
            "EndDate": "2019-01-01T00:00:00+02:00",                             // Type="DateTime"
            "Course": "string",                                                 // Type="String"
            "Status": "string",                                                 // Type="String"
            "TrainingNotCompleteStatus": 0,                                     // Type="Int32"
            "CourseResult": 1.00000,                                            // Type="Decimal"
            "DirectCourseCost": 1.00000,                                        // Type="Decimal"
            "InvoiceNumber": "string",                                          // Type="String"
            "Facilitator": "string",                                            // Type="String"
            "Assessor": "string",                                               // Type="String"
            "CertificateNo": "string",                                          // Type="String"
            "CertificateExpiryDate": "2019-01-01T00:00:00+02:00",               // Type="DateTime"
            "EmployeeRate": 1.00000,                                            // Type="Decimal"
            "TravelCost": 1.00000,                                              // Type="Decimal"
            "Accommodation": 1.00000,                                           // Type="Decimal"
            "FacilitatorCost": 1.00000,                                         // Type="Decimal"
            "FoodBeverage": 1.00000,                                            // Type="Decimal"
            "Additional5": 1.00000,                                             // Type="Decimal"
            "Additional6": 1.00000,                                             // Type="Decimal"
            "Notes": "string",                                                  // Type="String"
            "AttachmentName": "string",                                         // Type="String"
            "AttachmentUrl": null,
            "CustomFields": [		
                {			
                    "Code": "string",				                            // Type="String"
                    "Label": "string",			                                // Type="String"
                    "Value": "string",			        	                    // Type="String"
                    "OptionCode": "string"                                      // Type="String"
                }	
			]
        }
    ]
}
```

---

### Create record

**`POST`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeTraining`

Create a single `EmployeeTraining` record for the specified `EmployeeNumber`.

**Request Body:**

```json
{
  //see "EmployeeTraining" in metadata endpoint for available fields
}
```

**Example Response: Create record** (Status: 201 Created)

```json
{

    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeTraining",
    "TrainingId": 0, 								// Type="Int64"
    "EmployeeNumber": "string",						// Type="String"
    "FullName": "string",							// Type="String"
    "EffectiveDate": "DateTimeOffset",				// Type="DateTimeOffset"
    "EndDate": "DateTimeOffset",					// Type="DateTimeOffset"
    "Course": "string",								// Type="String"
    "Status": "string",								// Type="String"
    "TrainingNotCompleteStatus": "string",			// Type="String"
    "CourseResult": 0.00000,						// Type="Decimal"
    "DirectCourseCost": 0.00000,					// Type="Decimal"
    "InvoiceNumber": "string",						// Type="String"
    "Facilitator": "string",						// Type="String"
    "Assessor": "string",							// Type="String"
    "CertificateNo": "string",						// Type="String"
    "CertificateExpiryDate": "DateTimeOffset",		// Type="DateTimeOffset"
    "EmployeeRate": 0.00000,						// Type="Decimal"
    "TravelCost": 0.00000,							// Type="Decimal"
    "Accommodation": 0.00000,						// Type="Decimal"
    "FacilitatorCost": 0.00000,						// Type="Decimal"
    "FoodBeverage": 0.00000,						// Type="Decimal"
    "Additional5": 0.00000,							// Type="Decimal"
    "Additional6": 0.00000,							// Type="Decimal"
    "Notes": "string",								// Type="String"
    "AttachmentName": "string",						// Type="String"
    "AttachmentUrl": "string",						// Type="String"
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

### Update single record

**`PATCH`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeTraining({{TrainingId}})`

Update a single `EmployeeTraining` record for the specified `TrainingId`.

**Request Body:**

```json
{
  //see "EmployeeTraining" in metadata endpoint for available fields
}
```

**Example Response: Update single record** (Status: 204 No Content)

---

### Delete single record

**`DELETE`** `{{api-base-url}}/odata/v1.0/:company-id/EmployeeTraining({{TrainingId}})`

Delete a single `EmployeeTraining` record for the specified `TrainingId`.

**Example Response: Delete Employee Training Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeTraining",
    "Success": true							// Type="boolean"
}
```

---
## Employee Qualifications

Retrieve qualification information for all employees associated with the `company_id`.

### Get a Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeQualification?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeQualification` records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeQualification",
    "value": [
        {
            "EmployeeQualificationId": 0,                                   // Type="Int64"
            "EmployeeNumber": "string",                                     // Type="String"
            "FullName": "string",                                           // Type="String"
            "DateCompleted": "2000-00-00T00:00:00+02:00",                   // Type="DateTime"
            "Qualification": "string",                                      // Type="String"
            "EducationLevel": "string",                                     // Type="String"
            "Subjects": "string",                                           // Type="String"
            "Institution": "string",                                        // Type="String"
            "InstituteType": "string",                                      // Type="String"
            "HighestQualification": false                                   // Type="Boolean"
        }
    ]
}
```

---

### Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeQualification({{QualificationId}})`

Retrieves `EmployeeQualification` information based on the specified `QualificationId`.

**Example Response: Get Single** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeQualification",
    "EmployeeQualificationId": 0,                                   // Type="Int64"
    "EmployeeNumber": "string",                                     // Type="String"
    "FullName": "string",                                           // Type="String"
    "DateCompleted": "2000-00-00T00:00:00+02:00",                   // Type="DateTime"
    "Qualification": "string",                                      // Type="String"
    "EducationLevel": "string",                                     // Type="String"
    "Subjects": "string",                                           // Type="String"
    "Institution": "string",                                        // Type="String"
    "InstituteType": "string",                                      // Type="String"
    "HighestQualification": false                                   // Type="Boolean"
}
```

---

### Get a collection of all employee qualifications

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeQualification/all?$orderby={{employee-qualification-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-qualification-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee qualifications**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeQualification",
    "value": [
        {
            "EmployeeQualificationId": 1,                           // Type="Int64"
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "DateCompleted": "2019-01-01T00:00:00+02:00",           // Type="DateTime"
            "Qualification": "string",                              // Type="String"
            "EducationLevel": "string",                             // Type="String"
            "Subjects": "string",                                   // Type="String"
            "Institution": "string",                                // Type="String"
            "InstituteType": "string",                              // Type="String"
            "HighestQualification": false,                          // Type="Boolean"
            "AttachmentUrls": [],
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

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeQualification`

Create a single `EmployeeQualification` record for the specified `EmployeeNumber`.

**Request Body:**

```json
{
  //see "EmployeeQualification" in metadata endpoint for available fields
}
```

**Example Response: Create record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeQualification",
    "EmployeeQualificationId": 0,                                   // Type="Int64"
    "EmployeeNumber": "string",                                     // Type="String"
    "FullName": "string",                                           // Type="String"
    "DateCompleted": "2000-00-00T00:00:00+02:00",                   // Type="DateTime"
    "Qualification": "string",                                      // Type="String"
    "EducationLevel": "string",                                     // Type="String"
    "Subjects": "string",                                           // Type="String"
    "Institution": "string",                                        // Type="String"
    "InstituteType": "string",                                      // Type="String"
    "HighestQualification": false                                   // Type="Boolean"
}
```

---

### Update single record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeQualification({{QualificationId}})`

Update a single `EmployeeQualification` record for the specified `QualificationId`.

**Request Body:**

```json
{
  //see "EmployeeQualification" in metadata endpoint for available fields
}
```

**Example Response: Update single record** (Status: 204 No Content)

---

### Delete single record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeQualification({{QualificationId}})`

Delete a single `EmployeeQualification` record for the specified `QualificationId`.

**Example Response: Delete single record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v1.0/{{company-id}}/$metadata#EmployeeQualification",
    "Success": true							                        // Type="boolean"
}
```

---
## Employee Skills

Retrieve skill information for all employees associated with the `company_id`.

### Get a collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSkill?$filter={{$filter}}&$top={{$top}}&$skip={{$skip}}`

Retrieves a list of active `EmployeeSkill` records based on the specified `company_id.`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |

**Example Response: Get a collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSkill",
    "value": [
        {
            "EmployeeSkillId": 0, 						    // Type="Int64"
            "EmployeeNumber": "string",						// Type="String"
            "SkillCategory": "string",						// Type="String"
            "Skill": "string",                              // Type="String"
            "Competency": "string",                         // Type="String"
            "Experience": "string"                          // Type="String"
        }
    ]
}
```

---

### Get a Single Record

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSkill({{EmployeeSkillId}})`

Retrieves a single record of `EmployeeSkill` information based on the specified `EmployeeSkillId`.

**Example Response: Get a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSkill",
    "EmployeeSkillId": 0, 						    // Type="Int64"
    "EmployeeNumber": "string",						// Type="String"
    "SkillCategory": "string",						// Type="String"
    "Skill": "string",                              // Type="String"
    "Competency": "string",                         // Type="String"
    "Experience": "string"                          // Type="String"
}
```

---

### Create a Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSkill`

Creates a single `EmployeeSkill` record for the specified `EmployeeNumber`.

**Request Body:**

```json
{
  //see "EmployeeSkill" in metadata endpoint for available fields
}
```

**Example Response: Create a Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSkill",
    "EmployeeSkillId": 0, 						    // Type="Int64"
    "EmployeeNumber": "string",						// Type="String"
    "SkillCategory": "string",						// Type="String"
    "Skill": "string",                              // Type="String"
    "Competency": "string",                         // Type="String"
    "Experience": "string"                          // Type="String"
}
```

---

### Update Single Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSkill({{EmployeeSkillId}})`

Update a single `EmployeeSkill` record for the specified `EmployeeSkillId`.

**Request Body:**

```json
{
  //see "EmployeeSkill" in metadata endpoint for available fields
}
```

**Example Response: Update Single Record** (Status: 204 No Content)

---

### Delete a Single Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeSkill({{EmployeeSkillId}})`

Delete a single `EmployeeSkill` record for the specified `EmployeeSkillId`.

**Example Response: Delete a Single Record** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeSkill",
    "Success": true             // Type="boolean"
}
```

---

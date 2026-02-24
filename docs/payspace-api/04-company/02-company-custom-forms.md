# Company > Company Custom Forms

Working with custom forms defined for use on company level

List of available custom fields can be found on the customfields endpoint using the customFormCategory (see `Custom form custom fields` lookup)


## Get Collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyCustomForm?$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

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
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyCustomForm/",
    "@odata.count": 1,
    "value":[{
        "CustomFormId": 0,                              // Type="Int64"
        "CustomFormCategoryCode": "Test",               // Type="String"
        "EffectiveDate": "2022-07-01",                  // Type="DateTime"
        "ProjectCode": "string",                        // Type="String"
        "CustomFields": [                               // CustomField
            // List of custom fields relevent to custom form category
        ]
    }]
}
```

---
## Get Collection by category

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyCustomForm/:category?$top=10`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$top` | `10` |  |

**Example Response: Get Collection by category**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyCustomForm/",
    "@odata.count": 1,
    "value":[{
        "CustomFormId": 0,                              // Type="Int64"
        "CustomFormCategoryCode": "Test",               // Type="String"
        "EffectiveDate": "2022-07-01",                  // Type="DateTime"
        "ProjectCode": "string",                        // Type="String"
        "CustomFields": [                               // CustomField
            // List of custom fields relevent to custom form category
        ]
    }]
}
```

---
## Get Single

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyCustomForm({{CustomFormId}})`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Get Single** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#CompanyCustomForm/",
    "CustomFormId": 0,                              // Type="Int64"
    "CustomFormCategoryCode": "Test",               // Type="String"
    "EffectiveDate": "2022-07-01",                  // Type="DateTime"
    "ProjectCode": "string",                        // Type="String"
    "CustomFields": [                               // CustomField
        // List of custom fields relevent to custom form category
    ]
}
```

---
## Post

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyCustomForm`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Request Body:**

```json
{
    "CustomFormCategoryCode": "Test",               // Type="String"
    "EffectiveDate": "2022-07-01",                  // Type="DateTime"
    "CustomFields": [                               // CustomField
        // See list of fields available in customfields endpoint for the required customFormCategory, (see customfield with category)
    ]
}
```

**Example Response: Post** (Status: 201 Created)

---
## PATCH

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyCustomForm({{CustomFormId}})`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

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
## Delete

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/CompanyCustomForm({{CustomFormId}})`

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Example Response: Delete** (Status: 200 OK)

---

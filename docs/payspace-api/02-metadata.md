# Metadata

The `Metadata` endpoint returns a list of available entities that can be requested, along with any applicable enumeration values.
The exact content of the response may vary depending on company configuration, user permissions, and the company’s country of operation.

The following table outlines XML attributes that are not part of the OData standard. These attributes provide additional metadata about each field.

| Attribute                    | Description                                                                                                                                                                                                                                                                         |
| ---------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ps:Caption`                 | Specifies the display label of the field.                                                                                                                                                                                                                                           |
| `ps:Required`                | Indicates whether the field is required. The value may vary based on company settings.<br>- Possible values: `true`, `false`, `conditional`.<br>- `conditional`: The field may be required depending on the values of other fields.<br>- If not present, the field is not required. |
| `ps:Editable`                | Indicates whether the field can be modified after creation. If set to `false`, the field can only be set during creation.                                                                                                                                                           |
| `ps:Lookup`                  | Indicates that the field is a lookup. The value specifies the lookup endpoint to use.                                                                                                                                                                                               |
| `ps:Hidden`                  | Indicates whether the field is hidden. If set to `true`, the field is typically disabled by a company setting.                                                                                                                                                                      |
| `ps:MaxLength`               | Specifies the maximum allowed length for string values.                                                                                                                                                                                                                             |
| `ps:MinLength`               | Specifies the minimum required length for string values.                                                                                                                                                                                                                            |
| `ps:Unique`                  | Indicates that the field value must be unique.                                                                                                                                                                                                                                      |
| `ps:CompanySetting`          | Specifies the company setting code on which the field depends.                                                                                                                                                                                                                      |
| `ps:ReadOnly`                | Indicates that the field is read-only and cannot be set during creation (`POST`) or modified (`PATCH`).                                                                                                                                                                             |
| `ps:HasAdditionalValidation` | Indicates that additional validation applies, which is too complex to represent in metadata. For example, it may define rules that depend on values in other fields, potentially in different entities.                                                                             |
| `ps:MinValue`                | Specifies the minimum allowed numeric value.                                                                                                                                                                                                                                        |
| `ps:MaxValue`                | Specifies the maximum allowed numeric value.                                                                                                                                                                                                                                        |
| `ps:RegExFormat`             | Specifies a regular expression pattern that the field value must match.                                                                                                                                                                                                             |



## Get metadata

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/$metadata`

**Example Response: Get metadata** (Status: 200 OK)

```json
<?xml version="1.0" encoding="utf-8"?>
<edmx:Edmx Version="4.0" xmlns:edmx="http://docs.oasis-open.org/odata/ns/edmx">
    <edmx:DataServices>
        <Schema Namespace="PaySpace.Venuta.Data.Models.Dto" xmlns:ps="PaySpace" xmlns="http://docs.oasis-open.org/odata/ns/edm">
            <EntityType Name="CompanyGroupExchangeRate" ps:Caption="CompanyGroupExchangeRate">
                <Key>
                    <PropertyRef Name="ExchangeRateId" />
                </Key>
                <Property Name="ExchangeRateId" Type="Edm.Int64" Nullable="false" ps:Caption="ExchangeRateId" />
                <Property Name="Currency" Type="Edm.String" Nullable="false" ps:Caption="Currency" ps:Editable="false" ps:Lookup="ExchangeRateCurrency" ps:Required="true" />
                <Property Name="Country" Type="Edm.String" ps:Caption="Country" ps:Editable="false" ps:Lookup="Country" />
                <Property Name="EffectiveDate" Type="Edm.DateTimeOffset" Nullable="false" ps:Caption="EffectiveDate" ps:Required="true" />
                <Property Name="OneDollarInCurrency" Type="Edm.Decimal" Nullable="false" Scale="variable" ps:Caption="OneDollarInCurrency" ps:Required="true" />
            </EntityType>
            <EntityType Name="Employee" ps:Caption="Basic Profile">
                <Key>
                    <PropertyRef Name="EmployeeId" />
                </Key>
                <Property Name="Birthday" Type="Edm.Date" ps:Caption="Birth Date" ps:HasAdditionalValidation="true" ps:Required="true" />
                <Property Name="ImageDownloadUrl" Type="Edm.String" ps:Caption="ImageDownloadUrl" ps:Hidden="true" />
                <Property Name="EmployeeId" Type="Edm.Int64" Nullable="false" ps:Caption="EmployeeId" />
                <Property Name="OrganizationRegion" Type="Edm.String" ps:Caption="Organization Region" />
                <Property Name="InitiateWorkFlow" Type="Edm.Boolean" Nullable="false" ps:Caption="InitiateWorkFlow" ps:Hidden="true" />
                <Property Name="EmployeeNumber" Type="Edm.String" ps:Caption="Employee Number" ps:MaxLength="20" ps:Unique="true" ps:Required="conditional" />
                <Property Name="CompanyFrequency" Type="Edm.String" ps:Caption="Company frequency" ps:Editable="false" ps:CompanySetting="FREQEMP" ps:Hidden="true" ps:Lookup="CompanyFrequency" ps:Required="conditional" />
                <Property Name="Title" Type="Edm.String" ps:Caption="Title" ps:Lookup="Title" ps:Required="true" />
                <Property Name="FirstName" Type="Edm.String" Nullable="false" ps:Caption="First Name" ps:MaxLength="100" ps:HasAdditionalValidation="true" ps:Required="true" />
                <Property Name="LastName" Type="Edm.String" Nullable="false" ps:Caption="Last Name" ps:MaxLength="100" ps:HasAdditionalValidation="true" ps:Required="true" />
                <Property Name="PreferredName" Type="Edm.String" ps:Caption="Preferred Name" ps:MaxLength="100" />
                <Property Name="MaidenName" Type="Edm.String
... (truncated)
```

---

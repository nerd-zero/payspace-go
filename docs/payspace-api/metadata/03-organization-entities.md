# Organization Entities (Payspace Metadata)

*6 entity types — auto-generated from OData $metadata*

## Contents

- [OrganizationGrade](#organizationgrade) (9 properties)
- [OrganizationPayPoint](#organizationpaypoint) (2 properties)
- [OrganizationPositionDetail](#organizationpositiondetail) (22 properties)
- [OrganizationUnit](#organizationunit) (12 properties)
- [OrganizationUnitAddress](#organizationunitaddress) (15 properties)
- [PositionOrganizationUnitAddress](#positionorganizationunitaddress) (25 properties)

---

## OrganizationGrade
*Caption: Grades*

**Primary Key:** `GradeId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `GradeId` (PK) | Int64 | No |  |  |  |
| `Description` | String | No | Yes | 200 | Unique |
| `Code` | String | No | Yes | 64 | Unique |
| `Minimum` | Decimal | Yes |  |  |  |
| `Maximum` | Decimal | Yes |  |  |  |
| `BonusPercentage` | Decimal | Yes |  |  |  |
| `Average` | Decimal | Yes |  |  |  |
| `IsEligiblePm` | Boolean | Yes |  |  |  |
| `AdditionalValues` | Collection(PaySpace.Venuta.Data.Models.Dto.CompanyGradeFieldValue) | Yes |  |  |  |
---

## OrganizationPayPoint

**Primary Key:** `PayPointId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `PayPointId` (PK) | Int64 | No |  |  |  |
| `Description` | String | No | Yes | 100 |  |
---

## OrganizationPositionDetail
*Caption: Position Management*

**Primary Key:** `PositionDetailId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | Yes | Yes |  | |  |  |
| `InactiveDate` | Date | Yes |  |  | |  |  |
| `PositionDetailId` (PK) | Int64 | No |  |  | |  |  |
| `PositionTitle` | String | No | Yes |  | | 80 |  |
| `OrganizationPositionId` | Int64 | Yes |  | Create-only | |  | Hidden |
| `Grade` | String | Yes | Conditional |  | `Grade` |  |  |
| `PositionPurpose` | String | Yes |  |  | | 1000 |  |
| `SpecialFeatures` | String | Yes |  |  | | 1000 |  |
| `PositionType` | String | Yes |  |  | `PositionType` |  |  |
| `SSTLevel` | String | Yes |  |  | `SSTLevel` |  |  |
| `PositionCode` | String | Yes |  |  | | 50 |  |
| `OrganizationPositionWithCode` | String | Yes |  | Read-only | |  | Hidden |
| `DirectlyReportsPositionOverride` | String | Yes |  | Read-only | |  | Hidden |
| `CompanyTemplateDefinition` | String | Yes |  |  | `CompanyTemplateDefinition` |  |  |
| `OccupationalLevel` | String | Yes |  |  | `OccupationalLevel` |  |  |
| `EeFunction` | PaySpace.Venuta.Data.Models.Enums.EEFunction | Yes |  | Read-only | |  | Hidden |
| `BeeLevel` | PaySpace.Venuta.Data.Models.Enums.BEELevel | Yes |  |  | |  |  |
| `OfoCode` | String | Yes |  |  | `OfoCode` |  |  |
| `CompanyPositionFamily` | String | Yes |  |  | `CompanyPositionFamily` |  |  |
| `FunctionArea` | String | Yes |  |  | `FunctionArea` |  |  |
| `SalarySurveyCode1` | String | Yes |  |  | | 50 |  |
| `SalarySurveyCode2` | String | Yes |  |  | | 50 |  |
---

## OrganizationUnit
*Caption: Organisation Hierarchy Units*

**Primary Key:** `OrganizationUnitId`

**Navigation Properties:** `Address`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `InactiveDate` | Date | Yes |  | |  |  |
| `OrganizationUnitId` (PK) | Int64 | No |  | |  |  |
| `ParentOrganizationUnitId` | Int64 | Yes |  | |  |  |
| `UploadCode` | String | No | Yes | |  | Unique |
| `Description` | String | No | Yes | |  | Unique |
| `CostCentre` | Boolean | Yes |  | |  |  |
| `OrganizationLevel` | String | No | Yes | `OrganizationLevel` |  |  |
| `GroupGlKey` | String | Yes |  | |  |  |
| `Budget` | Decimal | Yes |  | |  |  |
| `Reference` | String | Yes |  | | 100 |  |
| `ManagerEmployeeNumber` | String | Yes |  | `ManagerEmployeeNumber` |  |  |
| `ActiveEmployees` | Int32 | No |  | |  |  |
---

## OrganizationUnitAddress

**Primary Key:** `AddressId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `AddressId` (PK) | Int64 | No |  | |  |  |
| `AddressLine1` | String | Yes | Conditional | |  |  |
| `DistributionCodeAbroad` | String | Yes | Conditional | | 50 | Hidden |
| `CompanyImplantation` | String | Yes | Conditional | `CompanyImplantation` |  | Hidden |
| `DistributionService` | String | Yes |  | | 50 | Hidden |
| `AddressLine2` | String | Yes |  | |  |  |
| `AddressLine3` | String | Yes | Conditional | |  |  |
| `AddressCode` | String | Yes |  | |  |  |
| `Province` | String | Yes | Yes | `AddressCountry/{AddressCountry}` |  |  |
| `AddressCountry` | String | Yes | Yes | `AddressCountry` |  |  |
| `UnitNumber` | String | Yes |  | | 8 |  |
| `Complex` | String | Yes |  | |  |  |
| `StreetNumber` | String | Yes |  | |  |  |
| `Municipality` | String | Yes |  | `Municipality` |  | Hidden |
| `AddressStreetType` | String | Yes |  | `AddressStreetType` |  | Hidden |
---

## PositionOrganizationUnitAddress

**Primary Key:** `AddressId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `AddressType` | PaySpace.Venuta.Data.Models.Enums.AddressType | No | Yes | |  | Unique |
| `SameAsPhysical` | Boolean | Yes |  | |  |  |
| `IsCareofAddress` | Boolean | Yes |  | |  |  |
| `CareOfIntermediary` | String | Yes |  | |  |  |
| `SpecialServices` | String | Yes |  | |  |  |
| `Block` | String | Yes |  | | 50 | Hidden |
| `Entrance` | String | Yes |  | | 50 | Hidden |
| `Staircase` | String | Yes |  | | 50 | Hidden |
| `Floor` | String | Yes |  | | 50 | Hidden |
| `Door` | String | Yes |  | | 50 | Hidden |
| `AddressId` (PK) | Int64 | No |  | |  |  |
| `AddressLine1` | String | Yes | Conditional | |  |  |
| `DistributionCodeAbroad` | String | Yes | Conditional | | 50 | Hidden |
| `CompanyImplantation` | String | Yes | Conditional | `CompanyImplantation` |  | Hidden |
| `DistributionService` | String | Yes |  | | 50 | Hidden |
| `AddressLine2` | String | Yes |  | |  |  |
| `AddressLine3` | String | Yes | Conditional | |  |  |
| `AddressCode` | String | Yes |  | |  |  |
| `Province` | String | Yes | Yes | `AddressCountry/{AddressCountry}` |  |  |
| `AddressCountry` | String | Yes | Yes | `AddressCountry` |  |  |
| `UnitNumber` | String | Yes |  | | 8 |  |
| `Complex` | String | Yes |  | |  |  |
| `StreetNumber` | String | Yes |  | |  |  |
| `Municipality` | String | Yes |  | `Municipality` |  | Hidden |
| `AddressStreetType` | String | Yes |  | `AddressStreetType` |  | Hidden |
---


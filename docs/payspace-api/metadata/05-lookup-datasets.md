# Lookup Entities (Payspace Metadata)

*30 entity types — auto-generated from OData $metadata*

## Contents

- [BureauComponentLookupDataSet](#bureaucomponentlookupdataset) (22 properties)
- [BureauTaxabilityOptionLookupDataSet](#bureautaxabilityoptionlookupdataset) (1 properties)
- [ComponentCodeLookupDataSet](#componentcodelookupdataset) (2 properties)
- [ComponentLookupDataSet](#componentlookupdataset) (29 properties)
- [ComponentVariableDataSet](#componentvariabledataset) (1 properties)
- [CountryLookupDataSet](#countrylookupdataset) (2 properties)
- [CustomFormCategoryLookupDataSet](#customformcategorylookupdataset) (6 properties)
- [DescriptionOnlyLookupDataSet](#descriptiononlylookupdataset) (0 properties)
- [EnumLookupDataSet](#enumlookupdataset) (1 properties)
- [ExchangeRateCurrencyLookupDataSet](#exchangeratecurrencylookupdataset) (1 properties)
- [ForfeitLeaveLookupDataSet](#forfeitleavelookupdataset) (4 properties)
- [IncidentTypeLookupDataSet](#incidenttypelookupdataset) (1 properties)
- [LeaveLookupDataSet](#leavelookupdataset) (4 properties)
- [LookupDataSet](#lookupdataset) (2 properties)
- [OccupationalLevelFieldValueLookDataSet](#occupationallevelfieldvaluelookdataset) (1 properties)
- [OnceOffComponentLookupDataSet](#onceoffcomponentlookupdataset) (1 properties)
- [OrganizationCategoryLookupDataSet](#organizationcategorylookupdataset) (2 properties)
- [OrganizationGroupLookupDataSet](#organizationgrouplookupdataset) (2 properties)
- [OrganizationGroupUploadCodeLookupDataSet](#organizationgroupuploadcodelookupdataset) (0 properties)
- [OrganizationLevelLookupDataSet](#organizationlevellookupdataset) (1 properties)
- [OrganizationRegionLookupDataSet](#organizationregionlookupdataset) (1 properties)
- [ProvinceLookupDataSet](#provincelookupdataset) (1 properties)
- [PurchasedCarOptionLookupDataSet](#purchasedcaroptionlookupdataset) (2 properties)
- [TableBuilderCodeDataSet](#tablebuildercodedataset) (1 properties)
- [TableBuilderLookupDataSet](#tablebuilderlookupdataset) (3 properties)
- [TaxCountryLookupDataset](#taxcountrylookupdataset) (1 properties)
- [TaxYearLookupDataSet](#taxyearlookupdataset) (3 properties)
- [TemplatesDataSet](#templatesdataset) (2 properties)
- [UploadCodeLookupDataSet](#uploadcodelookupdataset) (1 properties)
- [ValueOnlyLookupDataSet](#valueonlylookupdataset) (0 properties)

---

## BureauComponentLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Formula` | String | Yes |  |
| `OrderNumber` | Decimal | Yes |  |
| `ComponentType` | String | Yes |  |
| `SpecialComponentCode` | String | Yes |  |
| `IsOnceOff` | Boolean | Yes |  |
| `TaxCode` | String | Yes |  |
| `RecoverFigures` | Boolean | No |  |
| `AutoRecoveryType` | String | Yes |  |
| `AllowCompanyLevelAutoRecoveryChange` | Boolean | No |  |
| `PayslipAction` | PaySpace.Venuta.Data.Models.Enums.PayslipAction | Yes |  |
| `InPackagePayslipActionId` | Int32 | Yes |  |
| `InPackage` | Boolean | No |  |
| `EnforcePackageRule` | Boolean | No |  |
| `DoNotShowCurrency` | Boolean | No |  |
| `FormulaLines` | Collection(PaySpace.Venuta.Data.Models.Dto.ComponentFormulaLine) | Yes |  |
| `HasIndicator` | Boolean | No |  |
| `DisplayReadOnlyComponent` | Boolean | No |  |
| `PreventAddOnCompanyComponents` | Boolean | Yes |  |
| `FormulaTable` | PaySpace.Venuta.Data.Models.Enums.FormulaTables | Yes |  |
| `AllowFormulaOverride` | Boolean | No |  |
| `StartDate` | DateTimeOffset | Yes |  |
| `StopDate` | DateTimeOffset | Yes |  |
---

## BureauTaxabilityOptionLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CountryId` | Int32 | Yes |  |
---

## ComponentCodeLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `ComponentCode` | String | Yes |  |
| `AliasDescription` | String | Yes |  |
---

## ComponentLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `ComponentId` | Int64 | No |  |
| `Action` | PaySpace.Venuta.Data.Models.Enums.PayslipAction | No |  |
| `ComponentType` | String | Yes |  |
| `ComponentCode` | String | Yes |  |
| `TaxCode` | String | Yes |  |
| `TableBuilderCategory` | String | Yes |  |
| `IsTableBuilderCategoryRestricted` | Boolean | No |  |
| `FormulaTableType` | PaySpace.Venuta.Data.Models.Enums.FormulaTables | Yes |  |
| `ComponentLines` | Collection(PaySpace.Venuta.Data.Models.Dto.ComponentLine) | Yes |  |
| `ComponentSubCodes` | Collection(PaySpace.Venuta.Data.Models.Dto.ComponentSubCode) | Yes |  |
| `InPackage` | Boolean | No |  |
| `EnforcePackageRule` | Boolean | No |  |
| `FieldCode` | String | Yes |  |
| `EnableBcoe` | Boolean | No |  |
| `MedicalCategory` | String | Yes |  |
| `ComponentCategory` | String | Yes |  |
| `IsMonthlyByWeeks` | Boolean | Yes |  |
| `SpecialComponentCode` | String | Yes |  |
| `IsOnceOffComponent` | Boolean | No |  |
| `IsStatutoryComponent` | Boolean | No |  |
| `CannotEditInPaySlip` | Boolean | No |  |
| `DontDisplayInCompanyRecurring` | Boolean | No |  |
| `DontDisplayInEmployeeRecurring` | Boolean | No |  |
| `IsTableBuilderBalanceComponent` | Boolean | No |  |
| `TableBuilderCategoryId` | Int64 | No |  |
| `CurrencyCode` | String | Yes |  |
| `IncludeInBcoe` | Boolean | No |  |
| `BureauComponent` | String | Yes |  |
| `TableBuilderCategoryCode` | String | Yes |  |
---

## ComponentVariableDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `TaxCountry` | String | Yes |  |
---

## CountryLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `ISO2DigitCode` | String | Yes |  |
| `ISO3DigitCode` | String | Yes |  |
---

## CustomFormCategoryLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CustomFormLevel` | String | Yes |  |
| `CompanyLevel` | Boolean | Yes |  |
| `EmployeeLevel` | Boolean | Yes |  |
| `CustomFormScreenType` | PaySpace.Venuta.Data.Models.Enums.CustomFormScreenTypes | Yes |  |
| `CustomFormType` | PaySpace.Venuta.Data.Models.Enums.CustomFormType | Yes |  |
| `ShowPayrollCycle` | Boolean | Yes |  |
---

## DescriptionOnlyLookupDataSet

*No properties defined.*
---

## EnumLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EnumValue` | Int32 | No |  |
---

## ExchangeRateCurrencyLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CurrencyId` | Int64 | No |  |
---

## ForfeitLeaveLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `OrderNumber` | Int32 | No |  |
| `CompanyLeaveSetupId` | Int64 | No |  |
| `CompanyLeaveSchemeId` | Int64 | No |  |
| `LeaveType` | PaySpace.Venuta.Data.Models.Enums.LeaveType | No |  |
---

## IncidentTypeLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `IncidentEntityType` | String | Yes |  |
---

## LeaveLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `LeaveType` | PaySpace.Venuta.Data.Models.Enums.LeaveType | No |  |
| `LeaveBucket` | String | Yes |  |
| `CompanyLeaveScheme` | String | Yes |  |
| `IsEmployeeDefined` | Boolean | No |  |
---

## LookupDataSet

**Primary Key:** `Value`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Value` (PK) | String | No |  |
| `Description` | String | Yes |  |
---

## OccupationalLevelFieldValueLookDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `OrderNumber` | Int32 | Yes |  |
---

## OnceOffComponentLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `InputType` | String | Yes |  |
---

## OrganizationCategoryLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `PayRateHoursPerDay` | Double | Yes |  |
| `PayRateDaysPerPeriod` | Double | Yes |  |
---

## OrganizationGroupLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `OrganizationGroupDescription` | String | Yes |  |
| `OrganizationGroupId` | Int64 | No |  |
---

## OrganizationGroupUploadCodeLookupDataSet

*No properties defined.*
---

## OrganizationLevelLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Level` | Int32 | No |  |
---

## OrganizationRegionLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `RegionId` | Int64 | No |  |
---

## ProvinceLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CountryId` | Int32 | Yes |  |
---

## PurchasedCarOptionLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `OrderNumber` | Int32 | No |  |
| `Group` | String | Yes |  |
---

## TableBuilderCodeDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `DefaultValues` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomFieldDefaultValue) | Yes |  |
---

## TableBuilderLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Fields` | Collection(PaySpace.Venuta.Data.Models.Dto.LookupCustomFieldDetail) | Yes |  |
| `IsEmployeeLevelOnly` | Boolean | No |  |
| `CategoryCode` | String | Yes |  |
---

## TaxCountryLookupDataset

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CountryId` | Int32 | No |  |
---

## TaxYearLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `StartDate` | DateTimeOffset | No |  |
| `EndDate` | DateTimeOffset | No |  |
| `TaxYearId` | Int32 | No |  |
---

## TemplatesDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `TaxCountry` | String | Yes |  |
| `ParentCode` | String | Yes |  |
---

## UploadCodeLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `UploadCode` | String | Yes |  |
---

## ValueOnlyLookupDataSet

*No properties defined.*
---


# Company Entities (Payspace Metadata)

*59 entity types — auto-generated from OData $metadata*

## Contents

- [Company](#company) (35 properties)
- [CompanyAddress](#companyaddress) (25 properties)
- [CompanyAttachmentClassification](#companyattachmentclassification) (3 properties)
- [CompanyBankName](#companybankname) (3 properties)
- [CompanyBudgetPeriodLookupDataSet](#companybudgetperiodlookupdataset) (2 properties)
- [CompanyCategoryField](#companycategoryfield) (3 properties)
- [CompanyCategoryFieldValue](#companycategoryfieldvalue) (3 properties)
- [CompanyCloudRoom](#companycloudroom) (3 properties)
- [CompanyComponentGroup](#companycomponentgroup) (3 properties)
- [CompanyComponentLookupDataSet](#companycomponentlookupdataset) (1 properties)
- [CompanyCourseDelegate](#companycoursedelegate) (6 properties)
- [CompanyCourseSchedule](#companycourseschedule) (13 properties)
- [CompanyCustomForm](#companycustomform) (10 properties)
- [CompanyDisabilityLookupDataSet](#companydisabilitylookupdataset) (2 properties)
- [CompanyEbdIndicator](#companyebdindicator) (2 properties)
- [CompanyEmploymentCategory](#companyemploymentcategory) (3 properties)
- [CompanyEmploymentSubCategory](#companyemploymentsubcategory) (3 properties)
- [CompanyFrequency](#companyfrequency) (34 properties)
- [CompanyFrequencyDataSet](#companyfrequencydataset) (4 properties)
- [CompanyGL](#companygl) (3 properties)
- [CompanyGLDetail](#companygldetail) (8 properties)
- [CompanyGLExtra](#companyglextra) (7 properties)
- [CompanyGLInterAccount](#companyglinteraccount) (5 properties)
- [CompanyGradeField](#companygradefield) (3 properties)
- [CompanyGradeFieldValueLookDataSet](#companygradefieldvaluelookdataset) (1 properties)
- [CompanyGroupActionType](#companygroupactiontype) (3 properties)
- [CompanyGroupActionTypeReason](#companygroupactiontypereason) (5 properties)
- [CompanyGroupDataset](#companygroupdataset) (1 properties)
- [CompanyGroupExchangeRate](#companygroupexchangerate) (5 properties)
- [CompanyGroupLifeLookupDataSet](#companygrouplifelookupdataset) (2 properties)
- [CompanyIncidentType](#companyincidenttype) (2 properties)
- [CompanyIncreaseReason](#companyincreasereason) (3 properties)
- [CompanyLeaveAveragingIncome](#companyleaveaveragingincome) (9 properties)
- [CompanyLeaveBucketMapping](#companyleavebucketmapping) (3 properties)
- [CompanyLeaveBucketMappingDetail](#companyleavebucketmappingdetail) (6 properties)
- [CompanyLeavePurchase](#companyleavepurchase) (2 properties)
- [CompanyLeaveReason](#companyleavereason) (4 properties)
- [CompanyLegalBody](#companylegalbody) (3 properties)
- [CompanyManagerRequestType](#companymanagerrequesttype) (6 properties)
- [CompanyMedicalAidLookupDataSet](#companymedicalaidlookupdataset) (2 properties)
- [CompanyOffence](#companyoffence) (3 properties)
- [CompanyPensionFundLinkLookupDataSet](#companypensionfundlinklookupdataset) (5 properties)
- [CompanyPositionFamily](#companypositionfamily) (3 properties)
- [CompanyPositionFunctionArea](#companypositionfunctionarea) (3 properties)
- [CompanyPublicHoliday](#companypublicholiday) (7 properties)
- [CompanyQualification](#companyqualification) (4 properties)
- [CompanyRegion](#companyregion) (4 properties)
- [CompanyRegionHistory](#companyregionhistory) (6 properties)
- [CompanyReviewProcess](#companyreviewprocess) (8 properties)
- [CompanyRun](#companyrun) (29 properties)
- [CompanyRunLookupDataSet](#companyrunlookupdataset) (9 properties)
- [CompanySetting](#companysetting) (5 properties)
- [CompanySettingValue](#companysettingvalue) (2 properties)
- [CompanySkill](#companyskill) (5 properties)
- [CompanySkillCategory](#companyskillcategory) (4 properties)
- [CompanyTableBuilder](#companytablebuilder) (8 properties)
- [CompanyTrainingCourse](#companytrainingcourse) (24 properties)
- [EditPayslip](#editpayslip) (13 properties)
- [EditPayslipAdvanced](#editpayslipadvanced) (1 properties)

---

## Company

**Primary Key:** `CompanyId`

**Navigation Properties:** `CompanyGroupExchangeRates`, `Settings`, `Address`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `CompanyId` (PK) | Int64 | No |  | |  |  |
| `CompanyName` | String | Yes |  | | 128 |  |
| `TradingName` | String | Yes | Yes | |  |  |
| `RegistrationNo` | String | Yes | Yes | | 50 |  |
| `CompanyContactNo` | String | Yes | Yes | |  |  |
| `UIFReferenceNo` | String | Yes |  | |  |  |
| `CompanyGroup` | String | Yes |  | |  |  |
| `Country` | String | Yes |  | `TaxCountry` |  |  |
| `Currency` | String | Yes |  | `Currency` |  |  |
| `VatNumber` | String | Yes |  | |  |  |
| `CompanyContactPerson` | String | Yes |  | |  |  |
| `CompanyContactTitle` | String | Yes |  | `Title` |  | Hidden |
| `CompanyContactEmail` | String | Yes |  | |  | Hidden |
| `AgencyUrlOverride` | String | Yes |  | |  |  |
| `AgencyNameOverride` | String | Yes |  | |  |  |
| `CompanyContactFaxNo` | String | Yes |  | |  |  |
| `CompanyMobileNumber` | String | Yes |  | |  |  |
| `CompanyAlternateContactNo` | String | Yes |  | |  |  |
| `CompanyCode` | String | Yes |  | |  |  |
| `GenerateEmpNo` | Boolean | No |  | |  |  |
| `CompanyEmpNoGeneratedLength` | Int32 | Yes |  | |  |  |
| `CompanyEmpNoStart` | Int64 | Yes |  | |  |  |
| `CompanyEmployeeNumberPrefix` | String | Yes |  | |  |  |
| `TaxCalculationMethod` | String | Yes |  | `TaxCalculationMethodCode` |  |  |
| `PaymentModule` | String | Yes |  | `PaymentModule` |  |  |
| `PasswordExpiryPeriodOverride` | Int32 | Yes |  | |  |  |
| `KeepEmployeeNumOnTransfer` | Boolean | Yes |  | |  |  |
| `CompanyEmailPayslips` | Boolean | Yes |  | |  |  |
| `AdvancedPositionManagement` | Boolean | Yes |  | |  |  |
| `CompanyTaxReferenceNumber` | String | Yes | Yes | | 20 |  |
| `IncludeAnalytics` | Boolean | Yes |  | |  |  |
| `FinancialYearEndMonth` | Int32 | Yes |  | |  |  |
| `InitiateWorkFlow` | Boolean | No |  | |  |  |
| `CompanyContactGender` | String | Yes |  | |  | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  | |  |  |
---

## CompanyAddress

**Primary Key:** `AddressId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `AddressId` (PK) | Int64 | No |  | |  |  |
| `AddressType` | PaySpace.Venuta.Data.Models.Enums.AddressType | No | Yes | |  | Unique |
| `AddressLine1` | String | Yes | Conditional | |  |  |
| `Block` | String | Yes |  | | 50 | Hidden |
| `Entrance` | String | Yes |  | | 50 | Hidden |
| `Staircase` | String | Yes |  | | 50 | Hidden |
| `Floor` | String | Yes |  | | 50 | Hidden |
| `Door` | String | Yes |  | | 50 | Hidden |
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
| `SameAsPhysical` | Boolean | Yes |  | |  |  |
| `IsCareofAddress` | Boolean | Yes |  | |  |  |
| `CareOfIntermediary` | String | Yes |  | |  |  |
| `SpecialServices` | String | Yes |  | |  |  |
| `Municipality` | String | Yes |  | `Municipality` |  | Hidden |
| `AddressStreetType` | String | Yes |  | `AddressStreetType` |  | Hidden |
---

## CompanyAttachmentClassification

**Primary Key:** `ClassificationId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `ClassificationId` (PK) | Int64 | No |  |  |  |
| `Classification` | String | No | Yes | 50 |  |
| `ExcludeSecurityRoles` | String | Yes |  | 800 |  |
---

## CompanyBankName

**Primary Key:** `CompanyBankNameId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `CompanyBankNameId` (PK) | Int64 | No |  |  |  |
| `BankName` | String | No | Yes | 100 |  |
| `DefaultBranchCode` | String | Yes |  | 10 |  |
---

## CompanyBudgetPeriodLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `PeriodStartDate` | DateTimeOffset | Yes |  |
| `PeriodEndDate` | DateTimeOffset | Yes |  |
---

## CompanyCategoryField
*Caption: Pay Rate Categories*

**Primary Key:** `CompanyCategoryFieldId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CompanyCategoryFieldId` (PK) | Int64 | No |  |
| `Description` | String | Yes |  |
| `OrderNumber` | Int32 | No |  |
---

## CompanyCategoryFieldValue

**Primary Key:** `Description`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Description` (PK) | String | No |  |
| `Value` | String | Yes |  |
| `CompanyCategoryFieldId` | Int64 | No |  |
---

## CompanyCloudRoom

**Primary Key:** `CompanyCloudRoomId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CompanyCloudRoomId` (PK) | Int64 | No |  |
| `ZohoTicketId` | Int64 | Yes |  |
| `ZohoTicketNumber` | Int64 | Yes |  |
---

## CompanyComponentGroup

**Primary Key:** `ComponentGroupId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `ComponentGroupId` (PK) | Int32 | No |  |  |  |
| `ComponentGroupCode` | String | No | Yes | 10 |  |
| `ComponentGroupDescription` | String | No | Yes | 50 |  |
---

## CompanyComponentLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `ComponentId` | Int64 | No |  |
---

## CompanyCourseDelegate

**Primary Key:** `DelegateId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `DelegateId` (PK) | Int64 | No |  | |  |  |
| `ScheduleId` | Int64 | No | Yes | |  |  |
| `EmployeeNumber` | String | Yes |  | `EmployeeNumber` |  |  |
| `DelegateName` | String | No | Yes | | 100 |  |
| `EmailAddress` | String | No | Yes | | 100 |  |
| `DelegateType` | String | No | Yes | `InHouseOutsideProvider` | 100 |  |
---

## CompanyCourseSchedule

**Primary Key:** `ScheduleId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `ScheduleId` (PK) | Int64 | No |  | |  |  |
| `CourseDescription` | String | No | Yes | `CompanyCourseDescription` |  |  |
| `StartDate` | DateTimeOffset | No | Yes | |  |  |
| `EndDate` | DateTimeOffset | No | Yes | |  |  |
| `ApplicationCutOffDate` | DateTimeOffset | No | Yes | |  |  |
| `Venue` | String | No | Yes | | 200 |  |
| `VenueAddress` | String | Yes |  | | 500 |  |
| `MaxCandidates` | Int32 | No | Yes | |  |  |
| `NoCandidates` | Int32 | No |  | |  |  |
| `CoordinatorType` | String | No | Yes | `InHouseOutsideProvider` |  |  |
| `CoordinatorName` | String | No | Yes | | 100 |  |
| `CoordinatorEmail` | String | No | Yes | | 100 |  |
| `CoordinatorNumber` | String | Yes |  | | 100 |  |
---

## CompanyCustomForm
*Caption: Custom Forms*

**Primary Key:** `CustomFormId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EffectiveDate` | Date | Yes |  |  | |  |
| `InactiveDate` | Date | Yes |  |  | |  |
| `CustomFormId` (PK) | Int64 | No |  |  | |  |
| `CustomFormCategoryCode` | String | No | Yes | Create-only | `CustomFormCategoryCode` | Hidden |
| `RelatedEntityId` | Int64 | Yes |  |  | | Hidden |
| `ParentCustomFormId` | Int64 | Yes |  |  | | Hidden |
| `ProjectCode` | String | Yes |  |  | `ProjectCode` |  |
| `Code` | String | Yes |  |  | |  |
| `Description` | String | Yes |  |  | |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
---

## CompanyDisabilityLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EmployerPercentage` | Decimal | Yes |  |
| `EmployeePercentage` | Decimal | Yes |  |
---

## CompanyEbdIndicator

**Primary Key:** `EbdIndicatorId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `EbdIndicatorId` (PK) | Int64 | No |  |  |  |
| `EbdIndicator` | String | No | Yes | 50 |  |
---

## CompanyEmploymentCategory

**Primary Key:** `CategoryId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `CategoryId` (PK) | Int64 | No |  |  |  |
| `CategoryCode` | String | No | Yes | 10 |  |
| `CategoryDescription` | String | No | Yes | 100 |  |
---

## CompanyEmploymentSubCategory

**Primary Key:** `SubCategoryId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `SubCategoryId` (PK) | Int64 | No |  |  |  |
| `SubCategoryCode` | String | No | Yes | 10 |  |
| `Description` | String | No | Yes | 100 |  |
---

## CompanyFrequency

**Primary Key:** `CompanyFrequencyId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `WeeklyTaxYearStartDate` | Date | Yes | Conditional |  | |  |
| `InactiveFromDate` | Date | Yes |  |  | |  |
| `PreventAutoClosureUntil` | Date | Yes |  |  | |  |
| `PostRunByRunBeforeDate` | Date | Yes |  |  | |  |
| `PeriodStartDate` | Date | Yes |  |  | |  |
| `FirstProcessingPeriodEndDate` | Date | Yes | Conditional |  | | Hidden |
| `CompanyFrequencyId` (PK) | Int64 | No |  |  | |  |
| `RunFrequency` | String | No | Yes | Create-only | `RunFrequency` |  |
| `PayDay` | String | Yes | Yes |  | `RunFrequency/{RunFrequency}` |  |
| `PayBeforeWeekendOrHoliday` | Boolean | Yes |  | Create-only | |  |
| `LastWeekPartOfCurrentMonth` | Boolean | Yes |  | Create-only | |  |
| `AllowPayOnWeekends` | Boolean | Yes |  |  | |  |
| `HoursPerDay` | Decimal | Yes |  |  | |  |
| `DaysPerPeriod` | Decimal | Yes |  |  | |  |
| `FrequencyName` | String | No | Yes |  | |  |
| `TransferNetPay` | Int32 | Yes |  |  | |  |
| `TaxCertificatePrefix` | String | Yes |  |  | |  |
| `MajorBank` | Boolean | Yes |  |  | |  |
| `PostTimeSheet` | Boolean | Yes |  |  | |  |
| `AutomaticPayInd` | Boolean | Yes |  |  | |  |
| `WorkingDayMonday` | Boolean | Yes |  |  | |  |
| `WorkingDayTuesday` | Boolean | Yes |  |  | |  |
| `WorkingDayWednesday` | Boolean | Yes |  |  | |  |
| `WorkingDayThursday` | Boolean | Yes |  |  | |  |
| `WorkingDayFriday` | Boolean | Yes |  |  | |  |
| `WorkingDaySaturday` | Boolean | Yes |  |  | |  |
| `WorkingDaySunday` | Boolean | Yes |  |  | |  |
| `ExcludeFromBudget` | Boolean | Yes |  |  | |  |
| `HasAlternativeComponentNames` | Boolean | Yes |  |  | |  |
| `DefaultBasicPayQtyToDays` | Boolean | Yes |  |  | |  |
| `FirstCycleEndDay` | Int32 | Yes | Conditional |  | | Hidden |
| `FirstCyclePayDay` | Int32 | Yes | Conditional |  | | Hidden |
| `SecondCycleEndDay` | Int32 | Yes | Conditional |  | | Hidden |
| `SecondCyclePayDay` | Int32 | Yes | Conditional |  | | Hidden |
---

## CompanyFrequencyDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `FrequencyId` | Int64 | Yes |  |
| `Frequency` | PaySpace.Venuta.Data.Models.Enums.PayslipFrequency | Yes |  |
| `HoursPerDay` | Decimal | Yes |  |
| `DaysPerPeriod` | Decimal | Yes |  |
---

## CompanyGL

**Primary Key:** `CompanyGlId`

| Property | Type | Nullable | Required | Notes |
|----------|------|----------|----------|-------|
| `CompanyGlId` (PK) | Int64 | No |  |  |
| `HeaderName` | String | No | Yes |  |
| `EffectiveDate` | DateTimeOffset | No | Yes |  |
---

## CompanyGLDetail
*Caption: General Ledger Parameters*

**Primary Key:** `CompanyGLDetailId`

| Property | Type | Nullable | Access | Lookup | Notes |
|----------|------|----------|--------|--------|-------|
| `CompanyGLDetailId` (PK) | Int64 | No |  | |  |
| `ComponentCompany` | String | Yes | Create-only | `CompanyGLComponent` |  |
| `TaxCode` | String | Yes | Create-only | |  |
| `ComponentCode` | String | Yes |  | |  |
| `GLAccountNumber` | String | Yes |  | |  |
| `GLContraAccountNumber` | String | Yes |  | |  |
| `PayslipAction` | String | Yes | Create-only | |  |
| `GeneralLedgerHeader` | String | Yes | Create-only | `CompanyGlAccount` |  |
---

## CompanyGLExtra

**Primary Key:** `CompanyExtraGLId`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `CompanyExtraGLId` (PK) | Int64 | No | |  |
| `EffectiveDate` | DateTimeOffset | Yes | |  |
| `ComponentDescription` | String | Yes | |  |
| `GLAccountNo` | String | Yes | |  |
| `GLContraAccountNo` | String | Yes | |  |
| `GLAmount` | Decimal | No | |  |
| `GeneralLedgerHeader` | String | Yes | `CompanyGlAccount` |  |
---

## CompanyGLInterAccount

**Primary Key:** `CompanyGlInterAccId`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `CompanyGlInterAccId` (PK) | Int64 | No | |  |
| `InterGLAccount` | String | Yes | |  |
| `GLAccountNo` | String | Yes | |  |
| `GLContraAccountNo` | String | Yes | |  |
| `GeneralLedgerHeader` | String | Yes | `CompanyGlAccount` |  |
---

## CompanyGradeField

**Primary Key:** `GradeFieldId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `GradeFieldId` (PK) | Int64 | No |  |
| `Description` | String | Yes |  |
| `OrderNumber` | Int32 | Yes |  |
---

## CompanyGradeFieldValueLookDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `OrderNumber` | Int32 | Yes |  |
---

## CompanyGroupActionType

**Primary Key:** `ActionTypeId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `ActionTypeId` (PK) | Int64 | No |  |  |  |
| `ActionTypeCode` | String | No | Yes | 10 |  |
| `ActionTypeDescription` | String | No | Yes | 50 |  |
---

## CompanyGroupActionTypeReason

**Primary Key:** `ActionTypeReasonId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `InactiveDate` | Date | Yes |  | |  |  |
| `ActionTypeReasonId` (PK) | Int64 | No |  | |  |  |
| `CompanyGroupActionType` | String | No | Yes | `CompanyGroupActionType` |  |  |
| `ActionReasonCode` | String | No | Yes | | 10 |  |
| `ActionReasonDescription` | String | No | Yes | | 50 |  |
---

## CompanyGroupDataset

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CompanyGroupId` | Int64 | No |  |
---

## CompanyGroupExchangeRate

**Primary Key:** `ExchangeRateId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `ExchangeRateId` (PK) | Int64 | No |  |  | |  |
| `Currency` | String | No | Yes | Create-only | `ExchangeRateCurrency` |  |
| `Country` | String | Yes |  | Create-only | `Country` |  |
| `EffectiveDate` | DateTimeOffset | No | Yes |  | |  |
| `OneDollarInCurrency` | Decimal | No | Yes |  | |  |
---

## CompanyGroupLifeLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EmployerExtraRecurring` | Decimal | Yes |  |
| `EmployeeExtraRecurring` | Decimal | Yes |  |
---

## CompanyIncidentType

**Primary Key:** `CompanyIncidentTypeId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `CompanyIncidentTypeId` (PK) | Int64 | No |  |  |  |
| `IncidentType` | String | No | Yes | 100 |  |
---

## CompanyIncreaseReason

**Primary Key:** `ReasonId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `ReasonId` (PK) | Int64 | No |  |  |  |
| `ReasonCode` | String | No | Yes | 50 |  |
| `Description` | String | No | Yes | 100 |  |
---

## CompanyLeaveAveragingIncome

**Primary Key:** `CompanyLeaveAveragingIncomeId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `CompanyLeaveAveragingIncomeId` (PK) | Int64 | No |  |  | |  |
| `ComponentName` | String | Yes |  | Read-only | |  |
| `ComponentType` | String | Yes |  | Read-only | |  |
| `AveragingIncomePayslipAction` | String | Yes |  | Read-only | `AveragingIncomePayslipAction` |  |
| `MonthsToAverageOver` | Int32 | No | Yes |  | |  |
| `GuaranteeBonus` | Boolean | No |  |  | |  |
| `TaxCode` | String | Yes |  | Read-only | |  |
| `IncludeInAveragingIncome` | Boolean | No |  | Read-only | |  |
| `OrderNumber` | Int32 | Yes |  | Read-only | |  |
---

## CompanyLeaveBucketMapping

**Primary Key:** `LeaveBucketMappingId`

**Navigation Properties:** `LeaveBucketMappingDetails`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `LeaveBucketMappingId` (PK) | Int64 | No | |  |
| `FromLeaveScheme` | String | Yes | `LeaveBucketMappingScheme` |  |
| `ToLeaveScheme` | String | Yes | `LeaveBucketMappingScheme` |  |
---

## CompanyLeaveBucketMappingDetail

**Primary Key:** `LeaveBucketMappingDetailId`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `LeaveBucketMappingDetailId` (PK) | Int64 | No | |  |
| `LeaveMappingId` | Int64 | No | |  |
| `FromLeaveScheme` | String | Yes | |  |
| `ToLeaveScheme` | String | Yes | |  |
| `FromLeaveBucket` | String | Yes | `FromLeaveScheme/{FromLeaveScheme}` |  |
| `ToLeaveBucket` | String | Yes | `ToLeaveScheme/{ToLeaveScheme}` |  |
---

## CompanyLeavePurchase

**Primary Key:** `CompanyLeaveSetup`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `CompanyLeaveSetup` (PK) | String | No | Yes | Create-only | `CompanyLeaveSetup` |  |
| `IsSelected` | Boolean | No | Yes |  | |  |
---

## CompanyLeaveReason

**Primary Key:** `LeaveReasonId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `LeaveReasonId` (PK) | Int64 | No |  | |  |  |
| `CompanyLeaveSetup` | String | Yes |  | `CompanyLeaveSetup` |  |  |
| `LeaveType` | PaySpace.Venuta.Data.Models.Enums.LeaveType | No | Yes | |  |  |
| `LeaveReason` | String | No | Yes | | 100 |  |
---

## CompanyLegalBody

**Primary Key:** `CompanyLegalBodyId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `CompanyLegalBodyId` (PK) | Int64 | No |  |  |  |
| `LegalBodyDescription` | String | No | Yes | 100 |  |
| `LegalBodyCode` | String | No | Yes | 50 |  |
---

## CompanyManagerRequestType

**Primary Key:** `RequestTypeId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `RequestTypeId` (PK) | Int64 | No |  |  |  |
| `RequestTypeName` | String | No | Yes | 100 |  |
| `IsAttachmentReq` | Boolean | No | Yes |  |  |
| `RequestTypeDescription` | String | Yes |  | 500 |  |
| `ExcludeSecurityRoles` | String | Yes |  | 200 |  |
| `RequestTypeAttachments` | Collection(Edm.Int64) | No |  |  |  |
---

## CompanyMedicalAidLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Advised` | Boolean | Yes |  |
| `Percentage` | Decimal | Yes |  |
---

## CompanyOffence

**Primary Key:** `CompanyOffenceId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `CompanyOffenceId` (PK) | Int64 | No |  |  |  |
| `OffenceCode` | String | No | Yes | 50 |  |
| `OffenceDescription` | String | No | Yes | 100 |  |
---

## CompanyPensionFundLinkLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EmployerPercent` | Decimal | Yes |  |
| `EmployerFixedAmount` | Decimal | Yes |  |
| `EmployeePercent` | Decimal | Yes |  |
| `EmployeeFixedAmount` | Decimal | Yes |  |
| `EffectiveDate` | DateTimeOffset | Yes |  |
---

## CompanyPositionFamily

**Primary Key:** `FamilyId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `FamilyId` (PK) | Int64 | No |  |  |  |
| `Code` | String | No | Yes | 10 |  |
| `Description` | String | No | Yes | 50 |  |
---

## CompanyPositionFunctionArea

**Primary Key:** `FunctionAreaId`

| Property | Type | Nullable | Required | Max Length | Notes |
|----------|------|----------|----------|------------|-------|
| `FunctionAreaId` (PK) | Int64 | No |  |  |  |
| `Code` | String | No | Yes | 10 |  |
| `Description` | String | No | Yes | 50 |  |
---

## CompanyPublicHoliday
*Caption: Public Holidays*

**Primary Key:** `CompanyHolidayId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `HolidayDate` | Date | No | Yes | |  |  |
| `CompanyHolidayId` (PK) | Int32 | No |  | |  |  |
| `HolidayDescription` | String | No | Yes | | 100 |  |
| `PublicHolidayLevel` | PaySpace.Venuta.Data.Models.Enums.PublicHolidayLevel | No | Yes | |  |  |
| `PublicHolidayProvinces` | Collection(Edm.String) | Yes | Conditional | `CompanyProvince` |  |  |
| `PublicHolidayCategories` | Collection(Edm.String) | Yes | Conditional | `PublicHolidayCategory` |  |  |
| `PublicHolidayMunicipalities` | Collection(Edm.String) | Yes | Conditional | `Municipality` |  | Hidden |
---

## CompanyQualification
*Caption: Qualifications*

**Primary Key:** `CompanyQualificationId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `CompanyQualificationId` (PK) | Int64 | No |  | |  |  |
| `Description` | String | No | Yes | | 150 |  |
| `NqfLevel` | String | Yes |  | `NqfLevel` |  |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  | |  |  |
---

## CompanyRegion
*Caption: Regions*

**Primary Key:** `RegionId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `InactiveDate` | Date | Yes |  |
| `RegionId` (PK) | Int64 | No |  |
| `RegionDescription` | String | Yes |  |
| `RegionCode` | String | Yes |  |
---

## CompanyRegionHistory
*Caption: Regions*

**Primary Key:** `HistoryEntityDetailId`

| Property | Type | Nullable | Required | Notes |
|----------|------|----------|----------|-------|
| `EffectiveDate` | Date | Yes | Yes |  |
| `InactiveDate` | Date | Yes |  |  |
| `HistoryEntityDetailId` (PK) | Int64 | No |  |  |
| `RegionDescription` | String | No | Yes |  |
| `RegionCode` | String | No | Yes |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  |
---

## CompanyReviewProcess

**Primary Key:** `ProcessId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `StartDate` | Date | No |  |
| `EndDate` | Date | No |  |
| `ProcessId` (PK) | Int64 | No |  |
| `CompanyId` | Int64 | No |  |
| `ProcessTypeDescription` | String | Yes |  |
| `ProcessStatus` | PaySpace.Venuta.Data.Models.Enums.ReviewProcessStatus | No |  |
| `Title` | String | Yes |  |
| `ProcessWeighting` | Decimal | Yes |  |
---

## CompanyRun

**Primary Key:** `RunId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `PayDate` | Date | Yes |  |  | |  |
| `PayslipDate` | Date | Yes |  |  | | Hidden |
| `PeriodStartDate` | Date | No |  |  | |  |
| `PeriodEndDate` | Date | No |  |  | |  |
| `CutOffDate` | Date | Yes |  |  | |  |
| `TaxYearStartDate` | Date | Yes |  | Read-only | |  |
| `TaxYearEndDate` | Date | Yes |  | Read-only | |  |
| `RunType` | PaySpace.Venuta.Data.Models.Enums.RunType | No |  | Create-only | |  |
| `RunId` (PK) | Int64 | No |  |  | |  |
| `RunDescription` | String | No | Yes |  | |  |
| `OrderNumber` | Int32 | No | Yes |  | |  |
| `Status` | PaySpace.Venuta.Data.Models.Enums.RunStatus | No | Yes |  | |  |
| `TakeOnRun` | Boolean | No |  |  | |  |
| `PeriodCode` | String | Yes |  |  | |  |
| `HasBeenInvoiced` | Boolean | Yes |  |  | |  |
| `DenyTimesheetAccess` | Boolean | Yes |  |  | |  |
| `CutOff` | Boolean | Yes |  |  | |  |
| `EnablePayslipAccess` | Boolean | Yes |  |  | |  |
| `EnableEmpSelfServiceChanges` | Boolean | Yes |  |  | |  |
| `IsInterimTaxedPrevious` | Boolean | Yes |  |  | |  |
| `DisableEssAccess` | Boolean | Yes |  |  | |  |
| `RetroRun` | Boolean | Yes |  |  | |  |
| `BackdatedTakeOn` | Boolean | Yes |  |  | |  |
| `SendReportsPayDay` | Boolean | Yes |  |  | |  |
| `RunTypeExt` | String | Yes |  |  | `RunTypeExt` | Hidden |
| `TotalPeriodsInTaxYear` | Int32 | Yes |  | Read-only | |  |
| `CurrentPeriodInTaxYear` | Int32 | Yes |  | Read-only | |  |
| `TotalPeriodsInMonth` | Int32 | Yes |  | Read-only | |  |
| `CurrentPeriodInMonth` | Int32 | Yes |  | Read-only | |  |
---

## CompanyRunLookupDataSet

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `PeriodStartDate` | DateTimeOffset | No |  |
| `PeriodEndDate` | DateTimeOffset | No |  |
| `PayDate` | DateTimeOffset | Yes |  |
| `RunStatus` | PaySpace.Venuta.Data.Models.Enums.RunStatus | No |  |
| `OrderNumber` | Int32 | No |  |
| `IsMainRun` | Boolean | No |  |
| `RunType` | PaySpace.Venuta.Data.Models.Enums.RunType | No |  |
| `CutOffDate` | DateTimeOffset | Yes |  |
| `RunId` | Int64 | No |  |
---

## CompanySetting

**Primary Key:** `Code`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Code` (PK) | String | No |  |
| `Enabled` | Boolean | No |  |
| `Value` | String | Yes |  |
| `Description` | String | Yes |  |
| `Screen` | String | Yes |  |
---

## CompanySettingValue

**Primary Key:** `SettingTypeCode`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `SettingTypeCode` (PK) | String | No |  |
| `SettingValue` | String | Yes |  |
---

## CompanySkill
*Caption: Skills*

**Primary Key:** `CompanySkillId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `CompanySkillId` (PK) | Int64 | No |  | |  |  |
| `Description` | String | No | Yes | | 100 |  |
| `SkillCategory` | String | No | Yes | `SkillCategory` | 50 |  |
| `Overview` | String | Yes |  | | 500 |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  | |  |  |
---

## CompanySkillCategory
*Caption: Skills Categories*

**Primary Key:** `CompanySkillCategoryId`

| Property | Type | Nullable | Required | Access | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `CompanySkillCategoryId` (PK) | Int64 | No |  |  |  |  |
| `CompanySkillCategoryCode` | String | No | Yes | Create-only | 50 |  |
| `Description` | String | No | Yes |  | 100 |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  |  |  |
---

## CompanyTableBuilder

**Primary Key:** `EffectiveDate`, `TablebuilderId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EffectiveDate` (PK) | Date | No |  |
| `TablebuilderId` (PK) | Int64 | No |  |
| `Name` | String | Yes |  |
| `CategoryId` | Int64 | No |  |
| `CategoryCode` | String | Yes |  |
| `TableBuilderCode` | String | Yes |  |
| `Level` | PaySpace.Venuta.Data.Models.Enums.CompanyTableBuilderLevel | No |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |
---

## CompanyTrainingCourse
*Caption: Training Courses*

**Primary Key:** `CourseId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `EffectiveDate` | Date | Yes | Yes | |  |  |
| `StopDate` | Date | Yes |  | |  |  |
| `CourseId` (PK) | Int64 | No |  | |  |  |
| `CourseDescription` | String | No | Yes | | 1000 |  |
| `CourseOverview` | String | Yes |  | | 1000 |  |
| `CourseRequirements` | String | Yes |  | | 1000 |  |
| `CourseCostNotes` | String | Yes |  | | 50 |  |
| `InHouseOutsideProvider` | String | Yes |  | `InHouseOutsideProvider` |  |  |
| `Institution` | String | Yes |  | | 200 |  |
| `Comments` | String | Yes |  | | 200 |  |
| `MaxCandidates` | Int32 | Yes |  | |  |  |
| `MinCandidates` | Int32 | Yes |  | |  |  |
| `CourseDuration` | String | Yes |  | | 50 |  |
| `Credits` | Int32 | Yes |  | |  |  |
| `NqfLevel` | String | Yes |  | `NqfLevel` |  |  |
| `SaqaRegistered` | Boolean | Yes |  | |  |  |
| `SaqaIdNumber` | String | Yes |  | | 50 |  |
| `CourseType` | String | Yes |  | `CourseType` |  |  |
| `SkillsPriority` | String | Yes |  | `TrainingSkillsPriority` |  |  |
| `CourseDurationHours` | Int32 | Yes |  | |  |  |
| `DirectCourseCost` | Decimal | Yes |  | |  |  |
| `ProgramCategory` | String | Yes |  | `TrainingProgramCategory` |  |  |
| `IsEquityTraining` | Boolean | Yes |  | |  |  |
| `SchedulesCount` | Int32 | No |  | |  |  |
---

## EditPayslip
*Caption: Payslips*

**Primary Key:** `EditPayslipId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EditPayslipId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `ComponentCode` | String | No | Yes | Create-only | `ComponentCode` |  |
| `InputType` | String | Yes |  | Create-only | `ComponentCode/{ComponentCode}` |  |
| `InputValue` | Decimal | No | Yes |  | |  |
| `CostCentreCode` | String | Yes |  |  | `CostCentreCode` |  |
| `ProjectCode` | String | Yes |  |  | `ProjectCode` |  |
| `ActivityCode` | String | Yes |  |  | `ActivityCode` |  |
| `RecoveryAmount` | Decimal | Yes |  | Read-only | |  |
| `RecoveryAmountOverride` | Decimal | Yes |  |  | |  |
| `Comments` | String | Yes |  |  | |  |
| `RegionCode` | String | Yes |  | Read-only | `RegionCode` | Hidden |
---

## EditPayslipAdvanced
*Caption: Payslips – Multi Run*

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `CompanyRun` | String | Yes | `CompanyRun` |  |
---


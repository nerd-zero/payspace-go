# Employee Entities (Payspace Metadata)

*61 entity types — auto-generated from OData $metadata*

## Contents

- [Employee](#employee) (40 properties)
- [EmployeeAddress](#employeeaddress) (26 properties)
- [EmployeeAsset](#employeeasset) (14 properties)
- [EmployeeAttachment](#employeeattachment) (8 properties)
- [EmployeeBankDetail](#employeebankdetail) (30 properties)
- [EmployeeBonusProvision](#employeebonusprovision) (18 properties)
- [EmployeeClaim](#employeeclaim) (25 properties)
- [EmployeeCompanyCarDetail](#employeecompanycardetail) (32 properties)
- [EmployeeComponent](#employeecomponent) (11 properties)
- [EmployeeComponentValue](#employeecomponentvalue) (5 properties)
- [EmployeeCostedPayslipLine](#employeecostedpayslipline) (18 properties)
- [EmployeeCustomForm](#employeecustomform) (9 properties)
- [EmployeeDependant](#employeedependant) (19 properties)
- [EmployeeDisability](#employeedisability) (19 properties)
- [EmployeeEmploymentStatus](#employeeemploymentstatus) (54 properties)
- [EmployeeEtiTakeOn](#employeeetitakeon) (10 properties)
- [EmployeeGarnishee](#employeegarnishee) (25 properties)
- [EmployeeGroupLife](#employeegrouplife) (19 properties)
- [EmployeeHousePayment](#employeehousepayment) (15 properties)
- [EmployeeInbox](#employeeinbox) (12 properties)
- [EmployeeIncident](#employeeincident) (46 properties)
- [EmployeeIncomeProtection](#employeeincomeprotection) (18 properties)
- [EmployeeJournal](#employeejournal) (9 properties)
- [EmployeeLeaveAdjustment](#employeeleaveadjustment) (12 properties)
- [EmployeeLeaveApplication](#employeeleaveapplication) (25 properties)
- [EmployeeLeaveSetup](#employeeleavesetup) (7 properties)
- [EmployeeLeaveSetupEntitlement](#employeeleavesetupentitlement) (6 properties)
- [EmployeeLoan](#employeeloan) (25 properties)
- [EmployeeLumpSum](#employeelumpsum) (11 properties)
- [EmployeeMedical](#employeemedical) (27 properties)
- [EmployeeNote](#employeenote) (23 properties)
- [EmployeeNumberLookupDataSet](#employeenumberlookupdataset) (0 properties)
- [EmployeeOutOfOffice](#employeeoutofoffice) (6 properties)
- [EmployeePayRate](#employeepayrate) (44 properties)
- [EmployeePayslip](#employeepayslip) (17 properties)
- [EmployeePayslipLine](#employeepayslipline) (21 properties)
- [EmployeePayslipPdf](#employeepayslippdf) (9 properties)
- [EmployeePayslipTakeOn](#employeepaysliptakeon) (8 properties)
- [EmployeePensionFund](#employeepensionfund) (18 properties)
- [EmployeePosition](#employeeposition) (36 properties)
- [EmployeeProject](#employeeproject) (17 properties)
- [EmployeeQualification](#employeequalification) (13 properties)
- [EmployeeRecurringCosting](#employeerecurringcosting) (5 properties)
- [EmployeeRecurringCostingDetail](#employeerecurringcostingdetail) (6 properties)
- [EmployeeRecurringCostingSplit](#employeerecurringcostingsplit) (9 properties)
- [EmployeeRecurringTemplate](#employeerecurringtemplate) (4 properties)
- [EmployeeRetirementAnnuity](#employeeretirementannuity) (15 properties)
- [EmployeeReviewHeader](#employeereviewheader) (8 properties)
- [EmployeeReviewKpa](#employeereviewkpa) (7 properties)
- [EmployeeReviewKpaDetail](#employeereviewkpadetail) (5 properties)
- [EmployeeReviewTemplate](#employeereviewtemplate) (7 properties)
- [EmployeeSaving](#employeesaving) (15 properties)
- [EmployeeSkill](#employeeskill) (7 properties)
- [EmployeeSuspension](#employeesuspension) (21 properties)
- [EmployeeTableBuilder](#employeetablebuilder) (17 properties)
- [EmployeeTakeOn](#employeetakeon) (7 properties)
- [EmployeeTaxCertificate](#employeetaxcertificate) (11 properties)
- [EmployeeTraining](#employeetraining) (27 properties)
- [EmployeeTravelBusinessUsage](#employeetravelbusinessusage) (12 properties)
- [EmployeeUnion](#employeeunion) (13 properties)
- [EmployeeWorkflow](#employeeworkflow) (6 properties)

---

## Employee
*Caption: Basic Profile*

**Primary Key:** `EmployeeId`

**Navigation Properties:** `Address`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `Birthday` | Date | Yes | Yes |  | |  |  |
| `ImageDownloadUrl` | String | Yes |  |  | |  | Hidden |
| `EmployeeId` (PK) | Int64 | No |  |  | |  |  |
| `OrganizationRegion` | String | Yes |  |  | |  |  |
| `InitiateWorkFlow` | Boolean | No |  |  | |  | Hidden |
| `EmployeeNumber` | String | Yes | Conditional |  | | 20 | Unique |
| `CompanyFrequency` | String | Yes | Conditional | Create-only | `CompanyFrequency` |  | Hidden; Setting: `FREQEMP` |
| `Title` | String | Yes | Yes |  | `Title` |  |  |
| `FirstName` | String | No | Yes |  | | 100 |  |
| `LastName` | String | No | Yes |  | | 100 |  |
| `PreferredName` | String | Yes |  |  | | 100 |  |
| `MaidenName` | String | Yes |  |  | | 50 |  |
| `MiddleName` | String | Yes |  |  | | 100 |  |
| `Initials` | String | Yes | Yes |  | | 5 |  |
| `Email` | String | Yes |  |  | | 150 | Unique |
| `HomeNumber` | String | Yes |  |  | | 100 |  |
| `WorkNumber` | String | Yes |  |  | | 100 |  |
| `CellNumber` | String | Yes |  |  | | 100 |  |
| `WorkExtension` | String | Yes |  |  | | 6 |  |
| `Language` | String | Yes |  |  | `Language` |  |  |
| `Gender` | String | Yes | Yes |  | `Gender` |  |  |
| `MaritalStatus` | String | Yes |  |  | `MaritalStatus` |  | Setting: `MaritalStatus` |
| `Race` | String | Yes | Yes |  | `Race` |  |  |
| `Nationality` | String | Yes | Yes |  | `Nationality` |  |  |
| `Citizenship` | String | Yes |  |  | `Citizenship` |  |  |
| `DisabledType` | String | Yes |  |  | `DisabledType` |  |  |
| `ForeignNational` | Boolean | Yes |  |  | |  |  |
| `DateCreated` | DateTimeOffset | Yes |  | Create-only | |  | Hidden |
| `EmergencyContactName` | String | Yes |  |  | | 200 |  |
| `EmergencyContactNumber` | String | Yes |  |  | | 100 |  |
| `EmergencyContactAddress` | String | Yes |  |  | | 400 |  |
| `IsRetired` | Boolean | Yes |  |  | |  |  |
| `UifExemption` | String | Yes | Conditional |  | `UifExemption` |  |  |
| `SdlExemption` | String | Yes | Conditional |  | `SdlExemption` |  | Hidden |
| `EtiExempt` | Boolean | Yes |  |  | |  |  |
| `CustomFieldValue` | String | Yes |  |  | | 100 | Hidden; Setting: `CUSTFIELD` |
| `CustomFieldValue2` | String | Yes |  |  | | 100 | Hidden; Setting: `CUSTFIELD2` |
| `DefaultPayslip` | String | Yes |  |  | `DefaultPayslip` |  |  |
| `CompanyId` | Int64 | No |  |  | |  | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
---

## EmployeeAddress
*Caption: Employee Address*

**Primary Key:** `AddressId`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `EmployeeNumber` | String | Yes |  | `EmployeeNumber` |  |  |
| `Province` | String | Yes | Yes | `AddressCountry/{AddressCountry}` |  |  |
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

## EmployeeAsset
*Caption: Asset Register*

**Primary Key:** `AssetId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `IssuedDate` | Date | No | Yes |  | |  |
| `ReturnDate` | Date | Yes |  |  | |  |
| `AssetId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `Description` | String | No | Yes |  | |  |
| `SerialNo` | String | Yes |  |  | |  |
| `AssetStatus` | PaySpace.Venuta.Data.Models.Enums.AssetStatus | No | Yes |  | |  |
| `Quantity` | Int32 | No | Yes |  | |  |
| `RandValue` | Decimal | Yes |  |  | |  |
| `StockRoomNo` | String | Yes |  |  | |  |
| `Model` | String | Yes |  |  | |  |
| `Notes` | String | Yes |  |  | |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
---

## EmployeeAttachment

**Primary Key:** `AttachmentId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `AttachmentId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `Classification` | String | Yes |  |  | `Classification` |  |  |
| `AttachmentDescription` | String | No | Yes |  | | 256 |  |
| `Attachment` | Binary | Yes |  |  | |  |  |
| `AttachmentName` | String | No | Yes |  | | 100 |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
| `AttachmentUrl` | String | Yes |  | Read-only | |  |  |
---

## EmployeeBankDetail
*Caption: Banking Details*

**Primary Key:** `BankDetailId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `BankDetailId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `PaymentMethod` | String | No | Yes |  | `PaymentMethod` |  |  |
| `SplitType` | PaySpace.Venuta.Data.Models.Enums.BankDetailSplitType | Yes |  |  | |  |  |
| `BankAccountOwner` | PaySpace.Venuta.Data.Models.Enums.BankAccountOwnerType | No | Conditional |  | |  |  |
| `BankAccountOwnerName` | String | Yes | Conditional |  | | 100 |  |
| `AccountType` | PaySpace.Venuta.Data.Models.Enums.AccountType | No | Conditional |  | |  |  |
| `BankName` | String | Yes | Conditional |  | `BankName` |  |  |
| `BankBranchNo` | String | Yes | Conditional |  | | 6 |  |
| `AgencyCheckDigit` | String | Yes | Conditional |  | |  | Hidden |
| `BankAccountNo` | String | Yes | Conditional |  | | 16 |  |
| `AccountCheckDigit` | String | Yes | Conditional |  | |  | Hidden |
| `Reference` | String | Yes |  |  | | 40 |  |
| `Amount` | Decimal | Yes | Conditional |  | |  |  |
| `Comments` | String | Yes |  |  | | 200 |  |
| `SwiftCode` | String | Yes |  |  | | 50 |  |
| `RoutingCode` | String | Yes |  |  | | 100 |  |
| `CompanyComponent` | String | Yes |  |  | `CompanyComponent` |  |  |
| `Currency` | String | Yes |  |  | `Currency` |  |  |
| `CompanyEdbIndicator` | String | Yes |  |  | `CompanyEdbIndicator` |  |  |
| `SkipValidation` | Boolean | No |  |  | |  | Hidden |
| `BankingKey` | PaySpace.Venuta.Data.Models.Enums.BankingKeyType | Yes | Conditional |  | |  | Hidden |
| `CpfNumber` | String | Yes | Conditional |  | |  | Hidden |
| `CnpjNumber` | String | Yes | Conditional |  | |  | Hidden |
| `CellphoneNumber` | String | Yes | Conditional |  | |  | Hidden |
| `EmailAddress` | String | Yes | Conditional |  | |  | Hidden |
| `RandomKey` | String | Yes | Conditional |  | |  | Hidden |
| `DomesticClearingCode` | String | Yes |  | Read-only | | 50 | Hidden |
| `DomesticAuthorityBankCode` | String | Yes |  | Read-only | | 50 | Hidden |
---

## EmployeeBonusProvision

**Primary Key:** `EmployeeBonusProvisionId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `BonusPayDate` | Date | Yes |  |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeBonusProvisionId` (PK) | Int64 | No |  |  | |  |  |
| `IsBonusTaxSpread` | Boolean | No |  |  | |  |  |
| `SpreadToDateOrTaxEnd` | PaySpace.Venuta.Data.Models.Enums.SpreadToDateOrTaxEnd | Yes |  |  | |  |  |
| `BonusPayDateType` | PaySpace.Venuta.Data.Models.Enums.BonusPayDateType | No | Yes |  | |  |  |
| `OverrideBonusAccAmount` | Decimal | Yes |  |  | |  |  |
| `AutomaticPay` | Boolean | Yes |  | Read-only | |  | Hidden |
| `BonusPayOnTermination` | Boolean | Yes |  | Read-only | |  | Hidden |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeClaim

**Primary Key:** `EmployeeClaimId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EmployeeClaimId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `ClaimBatchId` | Int64 | No |  |  | |  |
| `Status` | PaySpace.Venuta.Data.Models.Enums.BatchStatus | No |  | Read-only | |  |
| `DateCaptured` | DateTimeOffset | Yes |  | Read-only | |  |
| `ClaimComponent` | String | Yes |  | Create-only | `ClaimComponent` |  |
| `OrganizationGroup` | String | Yes | Conditional | Create-only | `OrganizationGroup` | Hidden; Setting: `HIDECC` |
| `ProjectCode` | String | Yes | Conditional | Create-only | `ProjectCode` | Hidden; Setting: `HIDEPROJ` |
| `ActivityCode` | String | Yes | Conditional | Create-only | `ActivityCode` | Hidden; Setting: `HIDEACT` |
| `CompanyRun` | String | Yes |  | Read-only | |  |
| `Date` | DateTimeOffset | Yes |  |  | |  |
| `Description` | String | No | Yes |  | |  |
| `Rate` | Decimal | No |  |  | |  |
| `Quantity` | Decimal | No |  |  | |  |
| `TotalAmount` | Decimal | No | Yes |  | |  |
| `OpeningOdo` | Int32 | Yes |  |  | |  |
| `ClosingOdo` | Int32 | Yes |  |  | |  |
| `PrivateTravel` | Int32 | Yes |  |  | |  |
| `RegistrationNumber` | String | Yes |  |  | |  |
| `DepartureDate` | DateTimeOffset | Yes |  |  | |  |
| `ReturnDate` | DateTimeOffset | Yes |  |  | |  |
| `Destination` | String | Yes |  |  | |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
| `AttachmentUrls` | Collection(Edm.String) | Yes |  | Read-only | |  |
---

## EmployeeCompanyCarDetail

**Primary Key:** `EmployeeCompanyCarId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No |  |  | |  |  |
| `PurchaseDate` | Date | Yes | Yes |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeCompanyCarId` (PK) | Int64 | No |  |  | |  |  |
| `CarDescription` | String | No | Yes |  | |  |  |
| `WasSubjectToMaintenancePlan` | Boolean | Yes |  |  | |  |  |
| `CarValue` | Decimal | Yes |  |  | |  |  |
| `CarType` | PaySpace.Venuta.Data.Models.Enums.CarType | Yes | Yes |  | |  |  |
| `EmployerMaintenance` | Boolean | Yes |  |  | |  |  |
| `EmployerFuel` | Boolean | Yes |  |  | |  |  |
| `CompensationAmount` | Decimal | Yes |  |  | |  |  |
| `CarPaymentType` | String | Yes |  |  | `CarPaymentType` |  | Hidden |
| `KenyaPurchasedCarOption` | String | Yes |  |  | `KenyaPurchasedCarOption` |  | Hidden |
| `EmployeeBenefitPayments` | Decimal | Yes |  |  | |  | Hidden |
| `CustomField1` | String | Yes |  |  | |  | Hidden |
| `CustomField2` | String | Yes |  |  | |  | Hidden |
| `BodyType` | String | Yes |  |  | `BodyType` |  | Hidden |
| `CustomField3` | String | Yes |  |  | |  | Hidden |
| `CustomField4` | Decimal | Yes |  |  | |  |  |
| `CompanyCarTax` | PaySpace.Venuta.Data.Models.Enums.TaxableOption | Yes | Yes |  | |  |  |
| `Is80BusinessPurpose` | Boolean | Yes |  |  | |  | Hidden |
| `OneHundredPercentTax` | Boolean | Yes |  |  | |  | Hidden |
| `Inactive` | Boolean | Yes |  |  | |  | Hidden |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeComponent
*Caption: Recurring Payroll Components*

**Primary Key:** `ComponentEmployeeId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `ComponentEmployeeId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  | Read-only | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeComponentValue

**Primary Key:** `ComponentValueId`

| Property | Type | Nullable | Access | Notes |
|----------|------|----------|--------|-------|
| `ComponentValueId` (PK) | Int64 | No |  |  |
| `Description` | String | Yes | Read-only |  |
| `ComponentValue` | Decimal | Yes |  |  |
| `ComponentEmployeeId` | Int64 | No | Create-only |  |
| `ValueType` | PaySpace.Venuta.Data.Models.Enums.ValueType | Yes |  |  |
---

## EmployeeCostedPayslipLine

**Primary Key:** `CostedPayslipLineId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `CostedPayslipLineId` (PK) | Int64 | No |  |
| `PayslipLineId` | Int64 | No |  |
| `EmployeeNumber` | String | Yes |  |
| `RunDescription` | String | Yes |  |
| `PeriodCode` | String | Yes |  |
| `Comments` | String | Yes |  |
| `AliasDescription` | String | Yes |  |
| `CompanyComponentId` | Int64 | No |  |
| `AlternateComponentName` | String | Yes |  |
| `ComponentCode` | String | Yes |  |
| `PayslipAction` | PaySpace.Venuta.Data.Models.Enums.PayslipAction | Yes |  |
| `YtdAmount` | Decimal | Yes |  |
| `Quantity` | Decimal | No |  |
| `QuantityType` | String | Yes |  |
| `CostedProject` | String | Yes |  |
| `CostedOrgGroup` | String | Yes |  |
| `OrganizationLevel` | String | Yes |  |
| `CostedPayslipLineValue` | Decimal | No |  |
---

## EmployeeCustomForm
*Caption: Custom Forms*

**Primary Key:** `CustomFormId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EffectiveDate` | Date | Yes |  |  | | Unique |
| `CustomFormId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | Yes | Conditional |  | |  |
| `CustomFormCategoryCode` | String | No | Yes | Create-only | `CustomFormCategoryCode` | Hidden |
| `ParentCustomFormId` | Int64 | Yes |  |  | | Hidden |
| `CompanyRun` | String | Yes |  |  | `CompanyRun` | Hidden |
| `CustomFormType` | PaySpace.Venuta.Data.Models.Enums.CustomFormType | No |  |  | |  |
| `RelatedEntityId` | Int64 | Yes |  |  | | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
---

## EmployeeDependant
*Caption: Dependants*

**Primary Key:** `EmployeeDependantId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EffectiveDate` | Date | Yes | Yes |  | |  |
| `ChildBirthDate` | Date | Yes |  |  | |  |
| `MaritalDate` | Date | Yes |  |  | |  |
| `InactiveDate` | Date | Yes |  |  | |  |
| `EmployeeDependantId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `AdultRelationship` | PaySpace.Venuta.Data.Models.Enums.AdultRelationship | Yes |  |  | |  |
| `DependantType` | String | Yes | Yes |  | `EmployeeDependantType` |  |
| `MedicalAidDependant` | Boolean | No |  |  | |  |
| `FirstName` | String | Yes | Yes |  | |  |
| `LastName` | String | Yes | Yes |  | |  |
| `ContactNumber` | String | Yes |  |  | | Regex: `^[0-9]` |
| `MaritalStatus` | String | Yes |  |  | `DependantMaritalStatus` |  |
| `WorkingSpouse` | Boolean | Yes |  |  | |  |
| `IsBreadWinner` | Boolean | Yes |  |  | |  |
| `RestrictEmployerContribution` | Boolean | Yes |  |  | |  |
| `ChildMedicalAidCalculations` | Boolean | Yes |  |  | |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
---

## EmployeeDisability

**Primary Key:** `EmployeeDisabilityId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No |  |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeDisabilityId` (PK) | Int64 | No |  |  | |  |  |
| `CompanyDisability` | String | No | Yes |  | `DisabilityComponentCompany/{DisabilityComponentCompany}` |  |  |
| `DisabilityComponentCompany` | String | Yes |  |  | `DisabilityComponentCompany` |  | Hidden |
| `EmployerPercent` | Decimal | Yes |  |  | |  |  |
| `EmployeePercent` | Decimal | Yes |  |  | |  |  |
| `EmployerExtraAdvised` | Decimal | Yes |  |  | |  |  |
| `EmployeeExtraAdvised` | Decimal | Yes |  |  | |  |  |
| `OverrideAmount` | Decimal | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeEmploymentStatus
*Caption: Tax Profile*

**Primary Key:** `EmploymentStatusId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `GroupJoinDate` | Date | Yes | Yes |  | |  |  |
| `EmploymentDate` | Date | Yes | Yes |  | |  |  |
| `SeniorityDate` | Date | Yes |  |  | |  | Hidden |
| `AdditionalDate` | Date | Yes |  |  | |  | Hidden; Setting: `ADDTAXDATE,CUSTDATE1` |
| `AdditionalDate1` | Date | Yes |  |  | |  | Hidden; Setting: `ADDDATE1,CUSTDATE2` |
| `IdNumberExpiryDate` | Date | Yes | Conditional |  | |  | Hidden |
| `PermitIssued` | Date | Yes |  |  | |  |  |
| `PermitExpiry` | Date | Yes |  |  | |  |  |
| `PassportIssued` | Date | Yes |  |  | |  |  |
| `PassportExpiry` | Date | Yes |  |  | |  |  |
| `TerminationDate` | Date | Yes |  |  | |  |  |
| `OtherIssued` | Date | Yes |  |  | |  | Hidden |
| `OtherExpiry` | Date | Yes |  |  | |  | Hidden |
| `MyKadIssued` | Date | Yes |  |  | |  | Hidden |
| `MyKadExpiry` | Date | Yes |  |  | |  | Hidden |
| `EmploymentCaptureDate` | Date | Yes |  |  | |  | Hidden |
| `TerminationCaptureDate` | Date | Yes |  |  | |  | Hidden |
| `EmploymentStatusId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `TerminationReason` | String | Yes |  |  | `TaxStatus/{TaxStatus}` |  |  |
| `TaxStatus` | String | Yes | Yes |  | `TaxStatus` |  |  |
| `TaxReferenceNumber` | String | Yes | Conditional |  | | 10 | Regex: `^([0-9]{10,10})$` |
| `NatureOfPerson` | String | Yes | Yes |  | `NatureOfPerson` |  |  |
| `TaxDirectiveNumber` | String | Yes |  |  | | 64 |  |
| `EmploymentAction` | PaySpace.Venuta.Data.Models.Enums.EmploymentAction | Yes |  |  | |  |  |
| `TerminationCompanyRun` | String | Yes |  |  | `TerminationCompanyRun` |  |  |
| `IdentityType` | String | Yes |  |  | `NatureOfPerson/{NatureOfPerson}` |  |  |
| `IdNumber` | String | Yes | Conditional |  | | 30 | Regex: `^(18[0-9]{2}|19[0-9]{2}|20[0-9]{2})/\d{6}/(06|07|08|09|10|11|20|21|22|23|24|25|26|30|31)$` |
| `PassportNumber` | String | Yes |  |  | | 20 |  |
| `PercentageAmount` | PaySpace.Venuta.Data.Models.Enums.PercentageAmount | Yes |  |  | |  |  |
| `Amount` | Decimal | Yes |  |  | |  |  |
| `Percentage` | Decimal | Yes |  |  | |  |  |
| `DeemedMonthlyRemuneration` | Decimal | Yes |  |  | |  | Hidden |
| `Deemed75Indicator` | Boolean | Yes |  |  | |  | Hidden |
| `DeemedRecoveryMonthly` | Boolean | Yes |  |  | |  | Hidden |
| `EncashLeave` | Boolean | Yes |  |  | |  |  |
| `Irp30` | Boolean | Yes |  |  | |  |  |
| `FinalizeIssueTaxCert` | Boolean | Yes |  |  | |  | Hidden |
| `PassportCountry` | String | Yes |  |  | `PassportCountry` |  |  |
| `TempWorker` | Boolean | Yes |  |  | |  | Hidden; Setting: `CEPSS` |
| `ShadowPayroll` | Boolean | Yes |  |  | |  | Hidden |
| `NotReEmployable` | Boolean | Yes |  |  | |  |  |
| `ReferenceNumber` | String | Yes |  |  | | 50 |  |
| `OldEmployeeId` | Int64 | Yes |  | Create-only | |  | Hidden |
| `TaxOffice` | Int32 | Yes |  |  | |  | Hidden |
| `IT3AReason` | Int32 | Yes |  |  | |  | Hidden |
| `DentalBenefit` | String | Yes |  |  | `DentalBenefit` |  | Hidden |
| `Directorship` | PaySpace.Venuta.Data.Models.Enums.Directorship | Yes |  |  | |  | Hidden |
| `EmploymentIdentifier` | String | Yes |  |  | |  | Hidden |
| `TerminationPaySeverance` | Boolean | Yes |  |  | |  | Hidden |
| `SeveranceDay` | String | Yes | Conditional |  | `SeveranceDay` |  | Hidden |
| `RegisteredWithBirthDate` | Boolean | Yes |  |  | |  | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  | Read-only | |  |  |
---

## EmployeeEtiTakeOn
*Caption: Employee ETI Take On*

**Primary Key:** `EmployeeEtiTakeOnId`

| Property | Type | Nullable | Access | Lookup | Notes |
|----------|------|----------|--------|--------|-------|
| `PeriodCode` | Date | No |  | |  |
| `EmployeeEtiTakeOnId` (PK) | Int64 | No |  | |  |
| `EmployeeNumber` | String | Yes | Create-only | `TerminationEmployeeNumber` |  |
| `FullName` | String | Yes | Read-only | |  |
| `MonthlyRemuneration` | Decimal | Yes |  | |  |
| `MonthlyMinWage` | Decimal | Yes |  | |  |
| `WagePaid` | Decimal | Yes |  | |  |
| `EtiHours` | Decimal | Yes |  | |  |
| `EtiValue` | Decimal | Yes |  | |  |
| `EtiPeriodCount` | Int32 | Yes |  | |  |
---

## EmployeeGarnishee

**Primary Key:** `EmployeeGarnisheeId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeGarnisheeId` (PK) | Int64 | No |  |  | |  |  |
| `CreditorName` | String | No | Yes |  | |  |  |
| `BankName` | String | Yes |  |  | |  |  |
| `BranchCode` | String | Yes |  |  | |  |  |
| `AccountType` | String | Yes |  |  | `AccountType` |  |  |
| `AccountNumber` | String | Yes |  |  | |  |  |
| `Reference` | String | Yes |  |  | |  |  |
| `Description` | String | Yes |  |  | |  |  |
| `CapitalBalance` | String | No | Yes | Create-only | |  |  |
| `DeductionAmount` | Decimal | No | Yes |  | |  |  |
| `IncreaseDecrease` | Decimal | Yes |  |  | |  |  |
| `BalanceBroughtForward` | Decimal | Yes |  | Read-only | |  |  |
| `SubtractPercentage` | Decimal | Yes |  |  | |  |  |
| `DebitAccountNumber` | String | Yes |  |  | |  |  |
| `CreditAccountNumber` | String | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeGroupLife

**Primary Key:** `EmployeeGroupLifeId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No |  |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeGroupLifeId` (PK) | Int64 | No |  |  | |  |  |
| `CompanyGroupLife` | String | No | Yes |  | `GroupLifeComponentCompany/{GroupLifeComponentCompany}` |  |  |
| `GroupLifeComponentCompany` | String | Yes |  |  | `GroupLifeComponentCompany` |  | Hidden |
| `EmployeeExtraRecurring` | Decimal | Yes |  |  | |  |  |
| `EmployerExtraRecurring` | Decimal | Yes |  |  | |  |  |
| `EmployerPercent` | Decimal | Yes |  |  | |  |  |
| `EmployeePercent` | Decimal | Yes |  |  | |  |  |
| `OverrideAmount` | Decimal | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeHousePayment

**Primary Key:** `EmployeeFinancialHousePaymentId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | Yes | Yes |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeFinancialHousePaymentId` (PK) | Int64 | No |  |  | |  |  |
| `Description` | String | No | Yes |  | |  |  |
| `ReferenceNumber` | String | No | Yes |  | |  |  |
| `Value` | Decimal | No | Yes |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeInbox

**Primary Key:** `EmployeeInboxId`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `EmployeeInboxId` (PK) | Int64 | No | |  |
| `EmployeeNumber` | String | Yes | |  |
| `EmployeeFullName` | String | Yes | |  |
| `ApplicantEmployeeNumber` | String | Yes | |  |
| `ApplicantFullName` | String | Yes | |  |
| `OnBehalfOfEmployeeNumber` | String | Yes | |  |
| `OnBehalfOfEmployeeFullName` | String | Yes | |  |
| `ClaimBatchId` | Int64 | No | |  |
| `Subject` | String | Yes | |  |
| `InboxEntryType` | String | Yes | `InboxEntryType` |  |
| `UserWorkflowStepId` | Int64 | No | |  |
| `EffectiveDate` | DateTimeOffset | No | |  |
---

## EmployeeIncident
*Caption: Incident Management*

**Primary Key:** `EmployeeIncidentId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EmployeeIncidentId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `LinkedIncidentId` | Int64 | Yes |  | Read-only | |  |
| `IncidentType` | PaySpace.Venuta.Data.Models.Enums.EmployeeIncidentType | No | Yes | Create-only | |  |
| `CompanyOffence` | String | Yes |  |  | `CompanyOffence` |  |
| `OffenceOutcome` | String | Yes |  |  | `OffenceOutcome` |  |
| `OffenceCategory` | String | Yes |  |  | `OffenceCategory` |  |
| `AppealReason` | String | Yes |  |  | `AppealReason` |  |
| `CompanyLegalBody` | String | Yes |  |  | `CompanyLegalBody` |  |
| `CompanyIncidentType` | String | Yes |  |  | `CompanyIncidentType` |  |
| `AppealOutcome` | String | Yes |  |  | `AppealOutcome` |  |
| `AwardFavour` | String | Yes |  |  | `AwardFavour` |  |
| `OtherOutcome` | String | Yes |  |  | `OtherOutcome` |  |
| `SettlementReinstate` | String | Yes |  |  | `SettlementReinstate` |  |
| `IncidentDate` | DateTimeOffset | No | Yes |  | |  |
| `IncidentNotes` | String | Yes |  |  | |  |
| `EmployeeRepresentative` | String | Yes |  |  | |  |
| `CompanyRepresentative` | String | Yes |  |  | |  |
| `OutcomeDetails` | String | Yes |  |  | |  |
| `DateCharged` | DateTimeOffset | Yes |  |  | |  |
| `HearingDate` | DateTimeOffset | Yes |  |  | |  |
| `Chairperson` | String | Yes |  |  | |  |
| `CompanyWitnesses` | String | Yes |  |  | |  |
| `EmployeeWitnesses` | String | Yes |  |  | |  |
| `ValidityDate` | DateTimeOffset | Yes |  |  | |  |
| `CaseRefNumber` | String | Yes |  |  | |  |
| `SRDescription` | String | Yes |  |  | |  |
| `SRMonetaryValue` | Decimal | Yes |  |  | |  |
| `SRReinstatementDate` | DateTimeOffset | Yes |  |  | |  |
| `SRConditions` | String | Yes |  |  | |  |
| `CommissionerName` | String | Yes |  |  | |  |
| `LegalRepresentative` | String | Yes |  |  | |  |
| `SettlementDetails` | String | Yes |  |  | |  |
| `Accused` | String | Yes |  |  | |  |
| `Facilitator` | String | Yes |  |  | |  |
| `NatureOfGrievance` | String | Yes |  |  | |  |
| `GrievanceRecommendation` | String | Yes |  |  | |  |
| `DateOfOutcome` | DateTimeOffset | Yes |  |  | |  |
| `ReferralDate` | DateTimeOffset | Yes |  |  | |  |
| `OutcomeNotes` | String | Yes |  |  | |  |
| `IsGuilty` | Boolean | Yes |  |  | |  |
| `ReportedBy` | String | Yes |  |  | |  |
| `OutcomeDate` | DateTimeOffset | Yes |  |  | |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
| `AttachmentUrls` | Collection(Edm.String) | Yes |  | Read-only | |  |
---

## EmployeeIncomeProtection

**Primary Key:** `EmployeeIncomeProtectionId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No |  |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeIncomeProtectionId` (PK) | Int64 | No |  |  | |  |  |
| `CompanyIncomeProtectionLink` | String | Yes |  |  | `CompanyIncomeProtectionLink` |  |  |
| `EmployerExtraRecurring` | Decimal | Yes |  |  | |  |  |
| `EmployeeExtraRecurring` | Decimal | Yes |  |  | |  |  |
| `EmployerPercent` | Decimal | Yes |  |  | |  |  |
| `EmployeePercent` | Decimal | Yes |  |  | |  |  |
| `OverrideAmount` | Decimal | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeJournal
*Caption: Journal*

**Primary Key:** `JournalId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EffectiveDate` | Date | No | Yes |  | |  |
| `JournalId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `Subject` | String | No | Yes |  | |  |
| `Description` | String | No | Yes |  | |  |
| `Attachment` | Binary | Yes |  |  | |  |
| `AttachmentName` | String | Yes |  |  | | Hidden |
| `AttachmentDescription` | String | Yes |  |  | | Hidden |
| `AttachmentUrl` | String | Yes |  | Read-only | |  |
---

## EmployeeLeaveAdjustment
*Caption: Adjustment*

**Primary Key:** `LeaveAdjustmentId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `ConcessionYearEndDate` | Date | Yes |  |  | | Hidden |
| `LeaveAdjustmentId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `LeaveBucket` | String | No | Yes |  | `LeaveType/{LeaveType}` |  |
| `LeaveType` | String | No | Yes |  | `LeaveType` |  |
| `LeaveCompanyRun` | String | Yes |  |  | `LeaveCompanyRun` |  |
| `NoOfDays` | Double | No |  |  | |  |
| `Comments` | String | Yes |  |  | |  |
| `Reference` | String | Yes |  |  | |  |
| `OverrideBalance` | Boolean | No |  |  | |  |
| `LastModifiedDate` | DateTimeOffset | Yes |  | Read-only | |  |
---

## EmployeeLeaveApplication
*Caption: Application*

**Primary Key:** `LeaveAdjustmentId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `LeaveStartDate` | Date | Yes | Yes |  | |  |  |
| `LeaveEndDate` | Date | Yes | Yes |  | |  |  |
| `LeaveAdjustmentId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `LeaveBucket` | String | No | Yes |  | `LeaveType/{LeaveType}` |  |  |
| `LeaveType` | String | No | Yes |  | `LeaveType` |  |  |
| `LeaveCompanyRun` | String | Yes | Conditional |  | `LeaveCompanyRun` |  |  |
| `NoOfDays` | Double | No |  |  | |  |  |
| `TimeOfDay` | PaySpace.Venuta.Data.Models.Enums.TimeOfDay | Yes | Conditional |  | |  | Hidden; Setting: `LVETIME` |
| `Comments` | String | Yes |  |  | | 2000 |  |
| `Reference` | String | Yes |  |  | | 50 |  |
| `LeaveReason` | String | Yes | Conditional |  | `LeaveReason` |  |  |
| `LeaveStatus` | PaySpace.Venuta.Data.Models.Enums.LeaveStatus | No | Yes |  | |  |  |
| `LeaveTransactionType` | PaySpace.Venuta.Data.Models.Enums.LeaveEntryType | No | Yes |  | |  |  |
| `SkipValidation` | Boolean | No |  | Create-only | |  |  |
| `ThirteenCheque` | Boolean | Yes |  |  | |  | Hidden |
| `CancellationId` | Int64 | Yes |  | Read-only | |  | Hidden |
| `TempAttachmentSignature` | String | Yes |  |  | |  | Hidden |
| `ConcessionYearEnd` | String | Yes |  | Create-only | |  | Hidden |
| `LastModifiedDate` | DateTimeOffset | Yes |  | Read-only | |  |  |
| `SellVacationDays` | Boolean | No |  |  | |  | Hidden |
| `TotalDaysSell` | Double | Yes |  |  | |  | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
| `AttachmentUrl` | String | Yes |  | Read-only | |  | Hidden |
---

## EmployeeLeaveSetup
*Caption: Leave Setup*

**Primary Key:** `LeaveSetupId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EffectiveDate` | Date | Yes | Yes | Create-only | |  |
| `LeaveSetupId` (PK) | Int64 | No |  |  | |  |
| `CompanyLeaveScheme` | String | No | Yes | Create-only | `CompanyLeaveScheme` |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `PublicHolidayCategory` | String | Yes |  | Create-only | `PublicHolidayCategory` |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
---

## EmployeeLeaveSetupEntitlement
*Caption: Leave Setup Entitlement*

**Primary Key:** `LeaveSetupEntitlementId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `LeaveSetupEntitlementId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `LeaveSetupId` | Int64 | No |  | Create-only | | Hidden |
| `CompanyLeaveSetup` | String | No | Yes | Create-only | `CompanyLeaveSetup` |  |
| `AccrualValue` | Decimal | No | Yes |  | |  |
| `FullName` | String | Yes |  | Read-only | |  |
---

## EmployeeLoan

**Primary Key:** `EmployeeLoanId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeLoanId` (PK) | Int64 | No |  |  | |  |  |
| `Customfield1` | String | Yes |  |  | |  | Hidden |
| `LoanAmount` | Decimal | No | Yes | Create-only | |  |  |
| `BalanceBroughtForward` | Decimal | Yes |  | Read-only | |  |  |
| `PreviousBalance` | Decimal | Yes |  | Read-only | |  | Hidden |
| `InterestRate` | Decimal | Yes |  |  | |  |  |
| `InstallmentType` | PaySpace.Venuta.Data.Models.Enums.LoanInstallmentType | No | Yes |  | |  |  |
| `LoanCompanyComponent` | String | Yes | Conditional |  | `LoanCompanyComponent` |  |  |
| `Installment` | Decimal | Yes | Conditional |  | |  |  |
| `IncreaseDecrease` | Decimal | Yes |  |  | |  |  |
| `NonFringeBenefitLoan` | Boolean | Yes |  |  | |  |  |
| `DebitAccNo` | String | Yes |  |  | |  |  |
| `CreditAccNo` | String | Yes |  |  | |  |  |
| `PreviousInterestAmount` | Decimal | Yes |  | Read-only | |  | Hidden |
| `OverrideOutstandingBalance` | Decimal | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeLumpSum
*Caption: Lump Sum Directive*

**Primary Key:** `EmployeeLumpSumId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `DirectiveIssuedDate` | Date | Yes | Yes |  | |  |
| `EmployeeLumpSumId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `FinancialYear` | String | Yes |  |  | `FinancialYear` | Hidden |
| `TaxCode` | String | No | Yes | Create-only | `TaxCode` |  |
| `DirectiveAmount` | Decimal | No | Yes |  | |  |
| `DirectiveTax` | Decimal | No | Yes |  | |  |
| `DirectiveNumber` | String | Yes | Yes |  | |  |
| `ReferenceNumber` | String | Yes |  |  | |  |
| `TaxFreeDirectiveAmount` | Decimal | Yes |  |  | | Hidden |
---

## EmployeeMedical

**Primary Key:** `EmployeeMedicalId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No | Yes |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeMedicalId` (PK) | Int64 | No |  |  | |  |  |
| `CompanyMedicalAid` | String | No | Yes |  | `MedicalCategory/{MedicalCategory}` |  |  |
| `MedicalCategory` | String | Yes |  |  | `MedicalCategory` |  | Hidden |
| `TableOverride` | Boolean | No |  |  | |  |  |
| `PercentageAmount` | PaySpace.Venuta.Data.Models.Enums.PercentageAmount | No | Yes |  | |  |  |
| `TotalContribution` | Decimal | Yes |  |  | |  |  |
| `ReferenceNumber` | String | Yes |  |  | |  |  |
| `IsRetired` | Boolean | Yes |  |  | |  |  |
| `Maximum` | Decimal | Yes |  |  | |  |  |
| `EmployeeAdvisedAmount` | Decimal | Yes |  |  | |  |  |
| `EmployerAdvisedAmount` | Decimal | Yes |  |  | |  |  |
| `EmployeePercentage` | Decimal | Yes |  |  | |  |  |
| `EmployerPercentage` | Decimal | Yes |  |  | |  |  |
| `EmployeeExtraAdvised` | Decimal | Yes |  |  | |  |  |
| `EmployerExtraAdvised` | Decimal | Yes |  |  | |  |  |
| `Advised` | Boolean | Yes |  | Read-only | |  | Hidden |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeNote
*Caption: Notes and Reminders*

**Primary Key:** `EmployeeNoteId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EmployeeNoteId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `CaptureDateTime` | DateTimeOffset | No |  |  | | Hidden |
| `NoteTitle` | String | No | Yes |  | |  |
| `NoteDescription` | String | No | Yes |  | |  |
| `IsReminderAttached` | Boolean | Yes |  |  | |  |
| `ReminderDateTime` | DateTimeOffset | Yes |  |  | |  |
| `RemindEmployeeByEmail` | Boolean | Yes |  |  | |  |
| `RemindEmployeeBySms` | Boolean | Yes |  |  | | Hidden |
| `RemindMeBySMS` | Boolean | Yes |  |  | | Hidden |
| `IsPublic` | Boolean | Yes |  |  | | Hidden |
| `SendToDirectlyReportsTo` | Boolean | Yes |  |  | |  |
| `ReminderRecipientEmails` | String | Yes |  |  | |  |
| `IsRecurringReminder` | Boolean | Yes |  |  | |  |
| `ReminderEndDateTime` | DateTimeOffset | Yes |  |  | |  |
| `ReminderRepeatEveryNumberOfDays` | Int32 | Yes |  |  | |  |
| `RecurringFrequency` | PaySpace.Venuta.Data.Models.Enums.RecurrenceFrequency | Yes |  |  | |  |
| `ReminderFrequencyDay` | Int32 | Yes |  |  | |  |
| `RecurringWeekDays` | String | Yes |  |  | |  |
| `MonthsOfYear` | String | Yes |  |  | `MonthsOfYear` |  |
| `HaveActioned` | Boolean | No |  |  | | Hidden |
| `IsRemindersSent` | Boolean | No |  |  | | Hidden |
---

## EmployeeNumberLookupDataSet

*No properties defined.*
---

## EmployeeOutOfOffice

**Primary Key:** `OutOfOfficeId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `FromDate` | Date | Yes | Yes |  | |  |
| `ToDate` | Date | Yes | Yes |  | |  |
| `OutOfOfficeId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `AltEmployeeNumber` | String | No | Yes |  | `AltEmployeeNumber` |  |
---

## EmployeePayRate
*Caption: Pay Rate Details*

**Primary Key:** `PayRateId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No | Yes |  | |  | Unique |
| `PayRateId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `AutomaticPayInd` | Boolean | Yes |  |  | |  |  |
| `CompanyFrequency` | String | Yes | Conditional | Create-only | `CompanyFrequency` |  | Setting: `FREQEMP` |
| `PayFrequency` | PaySpace.Venuta.Data.Models.Enums.PayFrequency | No | Yes |  | |  |  |
| `HoursPerDay` | Double | Yes | Conditional |  | |  |  |
| `HoursPerMonth` | Double | Yes |  | Read-only | |  | Hidden |
| `HoursPerWeek` | Double | Yes |  | Read-only | |  | Hidden |
| `DaysPerPeriod` | Double | Yes | Conditional |  | |  |  |
| `WeeksPerMonth` | Double | Yes |  | Read-only | |  | Hidden |
| `HourlyRate` | Decimal | Yes |  | Read-only | |  | Hidden |
| `DailyRate` | Decimal | Yes |  | Read-only | |  | Hidden |
| `WeeklyRate` | Decimal | Yes |  | Read-only | |  | Hidden |
| `MonthlyRate` | Decimal | Yes |  | Read-only | |  | Hidden |
| `PercentageOfPreviousPackage` | Double | Yes |  |  | |  |  |
| `Package` | Decimal | No | Yes |  | |  |  |
| `HoursPerFortnight` | Double | Yes |  | Read-only | |  | Hidden |
| `FortnightlyRate` | Decimal | Yes |  | Read-only | |  | Hidden |
| `HoursPerSemiMonth` | Double | Yes |  | Read-only | |  | Hidden |
| `SemiMonthlyRate` | Decimal | Yes |  | Read-only | |  | Hidden |
| `AnnualPackage` | Decimal | Yes |  | Read-only | |  | Hidden |
| `ThirteenCheque` | Boolean | Yes |  |  | |  |  |
| `IsAnnual` | Boolean | Yes |  |  | |  |  |
| `Package2` | Decimal | Yes |  |  | |  | Hidden; Setting: `2NDPAYRATE` |
| `Reason` | String | Yes |  |  | `Reason` |  |  |
| `Currency` | String | Yes | Conditional |  | `Currency` |  | Setting: `PAYRATECUR` |
| `OrganizationCategory` | String | Yes |  |  | `OrganizationCategory` |  |  |
| `SalaryCoefficient` | Decimal | Yes |  |  | |  | Hidden |
| `PartialSalary` | Decimal | Yes |  | Read-only | |  | Hidden |
| `WorkingDayMonday` | Boolean | Yes |  |  | |  | Setting: `PROU` |
| `WorkingDayTuesday` | Boolean | Yes |  |  | |  | Setting: `PROU` |
| `WorkingDayWednesday` | Boolean | Yes |  |  | |  | Setting: `PROU` |
| `WorkingDayThursday` | Boolean | Yes |  |  | |  | Setting: `PROU` |
| `WorkingDayFriday` | Boolean | Yes |  |  | |  | Setting: `PROU` |
| `WorkingDaySaturday` | Boolean | Yes |  |  | |  | Setting: `PROU` |
| `WorkingDaySunday` | Boolean | Yes |  |  | |  | Setting: `PROU` |
| `DateAdded` | DateTimeOffset | Yes |  | Read-only | |  | Hidden |
| `Comments` | String | Yes |  |  | | 500 |  |
| `DateCreated` | DateTimeOffset | Yes |  | Read-only | |  | Hidden |
| `OneDollarInCurrency` | Decimal | Yes |  | Read-only | |  | Hidden |
| `BureauTaxabilityOption` | String | Yes |  |  | `BureauTaxabilityOption` |  | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
---

## EmployeePayslip

**Primary Key:** `PayslipId`

| Property | Type | Nullable | Access | Notes |
|----------|------|----------|--------|-------|
| `PeriodStartDate` | Date | No |  |  |
| `PeriodEndDate` | Date | No |  |  |
| `PayDate` | Date | No |  |  |
| `PayslipId` (PK) | Int64 | No |  |  |
| `CompanyRun` | String | Yes |  |  |
| `RunDescription` | String | Yes |  |  |
| `OrderNumber` | Int32 | No |  |  |
| `PeriodCode` | String | Yes |  |  |
| `Frequency` | String | Yes |  |  |
| `Paid` | Boolean | Yes |  |  |
| `PayslipComments` | String | Yes |  |  |
| `EmployeeNumber` | String | Yes | Create-only |  |
| `HourlyRate` | Decimal | Yes |  |  |
| `TotalValue` | Decimal | No |  |  |
| `NetPay` | Decimal | No |  |  |
| `PayslipLines` | Collection(PaySpace.Venuta.Data.Models.Dto.PayslipLine) | Yes |  |  |
| `LeaveBalances` | Collection(PaySpace.Venuta.Data.Models.Dto.LeaveBalance) | Yes |  |  |
---

## EmployeePayslipLine

**Primary Key:** `PayslipLineId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `PayslipLineId` (PK) | Int64 | No |  |
| `EmployeeNumber` | String | Yes |  |
| `RunDescription` | String | Yes |  |
| `PeriodCode` | String | Yes |  |
| `Comments` | String | Yes |  |
| `AliasDescription` | String | Yes |  |
| `CompanyComponentId` | Int64 | No |  |
| `AlternateComponentName` | String | Yes |  |
| `ComponentCode` | String | Yes |  |
| `PayslipAction` | PaySpace.Venuta.Data.Models.Enums.PayslipAction | Yes |  |
| `TaxCode` | String | Yes |  |
| `ValidTaxCode` | Boolean | Yes |  |
| `PayslipLineValue` | Decimal | No |  |
| `YtdAmount` | Decimal | Yes |  |
| `DoNotShowOnPayslip` | Boolean | Yes |  |
| `DoNotConvertCurrency` | Boolean | Yes |  |
| `InPackage` | Boolean | Yes |  |
| `Quantity` | Decimal | Yes |  |
| `Balance` | Decimal | Yes |  |
| `SpecialComponentCode` | String | Yes |  |
| `StatutoryComponent` | Boolean | No |  |
---

## EmployeePayslipPdf

**Primary Key:** `PayslipId`

| Property | Type | Nullable | Access | Notes |
|----------|------|----------|--------|-------|
| `PeriodStartDate` | Date | No |  |  |
| `PeriodEndDate` | Date | No |  |  |
| `DownloadUrl` | String | Yes |  |  |
| `PayslipId` (PK) | Int64 | No |  |  |
| `RunDescription` | String | Yes |  |  |
| `OrderNumber` | Int32 | No |  |  |
| `PeriodCode` | String | Yes |  |  |
| `Frequency` | String | Yes |  |  |
| `EmployeeNumber` | String | Yes | Create-only |  |
---

## EmployeePayslipTakeOn
*Caption: Take On Run By Run Figures*

**Primary Key:** `EmployeePayslipTakeOnId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EmployeePayslipTakeOnId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `ComponentCode` | String | No | Yes | Create-only | `TakeOnComponentCode` |  |
| `InputType` | String | Yes |  | Create-only | `TakeOnComponentCode/{ComponentCode}` |  |
| `Quantity` | Decimal | No |  |  | |  |
| `Amount` | Decimal | No |  |  | |  |
| `ComponentCompanyId` | Int64 | Yes |  | Read-only | |  |
---

## EmployeePensionFund

**Primary Key:** `EmployeePensionFundId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No |  |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeePensionFundId` (PK) | Int64 | No |  |  | |  |  |
| `CompanyPensionFund` | String | No | Yes |  | `PensionFundComponentCompany/{ComponentCompany}` |  |  |
| `ComponentCompany` | String | Yes |  |  | `PensionFundComponentCompany` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `CompanyPensionFundLink` | String | No | Yes |  | `CompanyPensionFund/{CompanyPensionFund}` |  |  |
| `EmployeeFixedAmount` | Decimal | Yes |  |  | |  |  |
| `EmployerFixedAmount` | Decimal | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `OverrideAmount` | Decimal | Yes |  |  | |  |  |
| `ReferenceNumber` | String | Yes |  |  | | 50 |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeePosition
*Caption: Position*

**Primary Key:** `EmployeePositionId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No | Yes |  | |  |  |
| `PositionEffectiveDate` | Date | Yes |  |  | |  |  |
| `OrganizationGroups` | Collection(PaySpace.Venuta.Data.Models.Dto.PositionOrganizationUnit) | Yes |  | Read-only | |  |  |
| `EmployeePositionId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `OrganizationPosition` | String | Yes |  | Read-only | |  | Hidden |
| `OrganizationPositionId` | Int64 | Yes |  |  | |  | Hidden |
| `OrganizationPositionWithCode` | String | Yes | Conditional |  | `OrganizationPositionWithCode` |  |  |
| `PositionType` | String | Yes |  |  | `PositionType` |  | Setting: `EMPSTATUS` |
| `Grade` | String | Yes |  |  | `Grade` |  | Setting: `POSGRADE` |
| `OccupationalLevel` | String | Yes |  | Read-only | |  | Hidden |
| `DirectlyReportsPositionOverride` | String | Yes |  |  | `DirectlyReportsPositionOverride` |  |  |
| `DirectlyReportsPosition` | String | Yes |  | Read-only | |  | Hidden |
| `OrganizationGroup` | String | Yes | Conditional |  | `OrganizationGroup` |  |  |
| `OrganizationGroupDescription` | String | Yes |  | Read-only | |  | Hidden |
| `OrganizationRegion` | String | Yes |  |  | `OrganizationRegion` |  | Setting: `REG` |
| `PayPoint` | String | Yes |  |  | `PayPoint` |  | Setting: `PAYPNT` |
| `DirectlyReportsEmployee` | String | Yes |  | Read-only | |  | Hidden |
| `DirectlyReportsEmployeeNumber` | String | Yes |  |  | `DirectlyReportsEmployeeNumberV2` |  | Setting: `DIRREPTOPERS` |
| `EmploymentCategory` | String | Yes |  |  | `EmploymentCategory` |  | Setting: `CAT` |
| `EmploymentSubCategory` | String | Yes |  |  | `EmploymentSubCategory` |  | Setting: `SUBCAT` |
| `Administrator` | String | Yes |  | Read-only | |  | Hidden |
| `AdministratorEmployeeNumber` | String | Yes |  |  | `AdministratorEmployeeNumber` |  |  |
| `WorkflowRole` | String | Yes |  |  | `WorkflowRole` |  |  |
| `GeneralLedger` | String | Yes |  |  | `GeneralLedger` |  |  |
| `TradeUnion` | String | Yes |  |  | `TradeUnion` |  | Hidden; Setting: `SHOWTRADEUNI` |
| `IsPromotion` | Boolean | Yes |  |  | |  |  |
| `KeepJob` | Boolean | No |  |  | |  | Hidden; Setting: `JOBMod` |
| `Roster` | String | Yes |  |  | `Roster` |  |  |
| `Job` | String | Yes | Conditional |  | `Job` |  | Hidden; Setting: `JOBMod` |
| `Comments` | String | Yes |  |  | | 400 |  |
| `AltPositionName` | String | Yes |  |  | | 200 | Setting: `EmpAltJobTitle` |
| `CustomTradeUnion` | String | Yes |  |  | `CustomTradeUnion` |  | Hidden; Setting: `SHOWTRADEUNI` |
| `OccupationCode` | String | Yes |  |  | `OccupationCode` |  | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
---

## EmployeeProject
*Caption: Project Details*

**Primary Key:** `EmployeeProjectId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | Yes | Yes |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeProjectId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ProjectCode` | String | No | Yes |  | `ProjectCode` |  |  |
| `ProjectDescription` | String | Yes |  | Read-only | |  | Hidden |
| `ProjectActivityCode` | String | Yes |  |  | `ProjectCode/{ProjectCode}` |  |  |
| `ProjectBased` | Boolean | Yes |  |  | |  |  |
| `ReminderOption` | PaySpace.Venuta.Data.Models.Enums.ReminderOption | Yes |  |  | |  |  |
| `PaidByInvoice` | Boolean | Yes |  |  | |  |  |
| `PayFrequency` | PaySpace.Venuta.Data.Models.Enums.PayFrequency | Yes |  |  | |  | Hidden; Setting: `PROJPAYRATE` |
| `ProjectPayRate` | Decimal | Yes |  |  | |  | Hidden; Setting: `PROJPAYRATE` |
| `HoursPerDay` | Decimal | Yes |  |  | |  | Hidden; Setting: `PROJPAYRATE` |
| `DaysPerMonth` | Decimal | Yes |  |  | |  | Hidden; Setting: `PROJPAYRATE` |
| `Comments` | String | Yes |  |  | | 500 |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
---

## EmployeeQualification
*Caption: Qualifications*

**Primary Key:** `EmployeeQualificationId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `DateCompleted` | Date | Yes | Yes |  | |  |  |
| `EmployeeQualificationId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `Qualification` | String | No | Yes |  | `Qualification` |  |  |
| `EducationLevel` | String | No | Yes |  | `EducationLevel` |  |  |
| `Subjects` | String | Yes |  |  | | 250 |  |
| `Institution` | String | Yes |  |  | | 100 |  |
| `InstituteType` | String | Yes |  |  | `InstituteType` |  |  |
| `HighestQualification` | Boolean | Yes |  |  | |  |  |
| `Attachment` | Binary | Yes |  |  | |  |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
| `AttachmentUrls` | Collection(Edm.String) | Yes |  | Read-only | |  |  |
---

## EmployeeRecurringCosting

**Primary Key:** `RecurringCostingSplitHeaderId`

**Navigation Properties:** `RecurringCostingSplitDetails`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EffectiveDate` | Date | No | Yes |  | |  |
| `RecurringCostingSplitHeaderId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `BasedOnSplitOption` | String | Yes |  | Create-only | `BasedOnSplitOption` |  |
---

## EmployeeRecurringCostingDetail

**Primary Key:** `RecurringCostingSplitDetailId`

| Property | Type | Nullable | Required | Lookup | Notes |
|----------|------|----------|----------|--------|-------|
| `RecurringCostingSplitDetailId` (PK) | Int64 | No |  | |  |
| `OrganizationGroup` | String | Yes |  | `OrganizationGroup` |  |
| `OrganizationGroupDescription` | String | Yes |  | |  |
| `ProjectCode` | String | Yes |  | `ProjectCode` |  |
| `ProjectActivityCode` | String | Yes |  | `ProjectCode/{ProjectCode}` |  |
| `Percentage` | Decimal | No | Yes | |  |
---

## EmployeeRecurringCostingSplit

**Primary Key:** `RecurringCostingSplitDetailId`

| Property | Type | Nullable | Access | Lookup | Notes |
|----------|------|----------|--------|--------|-------|
| `EffectiveDate` | Date | No |  | |  |
| `RecurringCostingSplitDetailId` (PK) | Int64 | No |  | |  |
| `OrganizationGroup` | String | Yes |  | `OrganizationGroup` |  |
| `OrganizationGroupDescription` | String | Yes |  | |  |
| `EmployeeNumber` | String | Yes | Create-only | |  |
| `Percentage` | Decimal | No |  | |  |
| `ProjectCode` | String | Yes |  | `ProjectCode` |  |
| `ProjectActivityCode` | String | Yes |  | `ProjectCode/{ProjectCode}` |  |
| `BasedOnSplitOption` | String | Yes | Create-only | |  |
---

## EmployeeRecurringTemplate
*Caption: Recurring Component Template*

**Primary Key:** `TemplateId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `TemplateId` (PK) | Int64 | No |  |  | | Hidden |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `RecurringTemplate` | String | No | Yes |  | `RecurringTemplate` |  |
| `ReplaceExisting` | Boolean | No |  |  | |  |
---

## EmployeeRetirementAnnuity

**Primary Key:** `EmployeeRetirementAnnuityId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No |  |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeRetirementAnnuityId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeRetirementAnnuityName` | String | No | Yes |  | |  |  |
| `Reference` | String | No | Yes |  | |  |  |
| `Amount` | Decimal | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeReviewHeader

**Primary Key:** `EmployeeReviewHeaderId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EmployeeReviewHeaderId` (PK) | Int64 | No |  |
| `RaterId` | Int64 | No |  |
| `RaterEmployeeNumber` | String | Yes |  |
| `RaterFirstName` | String | Yes |  |
| `RaterLastName` | String | Yes |  |
| `RaterType` | PaySpace.Venuta.Data.Models.Enums.RaterTypes | No |  |
| `OverallScore` | Decimal | Yes |  |
| `Weighting` | Decimal | Yes |  |
---

## EmployeeReviewKpa

**Primary Key:** `EmployeeReviewKpaId`

**Navigation Properties:** `EmployeeReviewKpaDetails`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EffectiveDate` | Date | No |  |
| `EmployeeReviewKpaId` (PK) | Int64 | No |  |
| `EmployeeNumber` | String | Yes |  |
| `SectionHeader` | String | Yes |  |
| `Weighting` | Decimal | Yes |  |
| `KpaManagerStatusId` | PaySpace.Venuta.Data.Models.Enums.ItemStatus | Yes |  |
| `KpaEmployeeStatusId` | PaySpace.Venuta.Data.Models.Enums.ItemStatus | Yes |  |
---

## EmployeeReviewKpaDetail

**Primary Key:** `EmployeeReviewKpaDetailId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EmployeeReviewKpaDetailId` (PK) | Int64 | No |  |
| `KpaTitle` | String | Yes |  |
| `KpaDescription` | String | Yes |  |
| `Weighting` | Decimal | No |  |
| `KpaMeansOfVerification` | String | Yes |  |
---

## EmployeeReviewTemplate

**Primary Key:** `EmployeeReviewTemplateId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EffectiveDate` | Date | No |  |
| `EmployeeReviewTemplateId` (PK) | Int64 | No |  |
| `EmployeeNumber` | String | Yes |  |
| `TemplateTitle` | String | Yes |  |
| `ProcessTitle` | String | Yes |  |
| `Status` | PaySpace.Venuta.Data.Models.Enums.ReviewTemplate | No |  |
| `OverallScore` | Decimal | Yes |  |
---

## EmployeeSaving

**Primary Key:** `EmployeeSavingId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeSavingId` (PK) | Int64 | No |  |  | |  |  |
| `SavingAmount` | Decimal | No |  | Create-only | |  |  |
| `SavingBalance` | Decimal | Yes |  | Create-only | |  |  |
| `Installment` | Decimal | No |  |  | |  |  |
| `IncreaseDecrease` | Decimal | Yes |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeSkill
*Caption: Skills*

**Primary Key:** `EmployeeSkillId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EmployeeSkillId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `SkillCategory` | String | No | Yes | Create-only | `SkillCategory` |  |
| `Skill` | String | No | Yes |  | `SkillCategory/{SkillCategory}` |  |
| `Competency` | String | No | Yes |  | `Competency` |  |
| `Experience` | String | No | Yes |  | `Experience` |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
---

## EmployeeSuspension
*Caption: Suspension*

**Primary Key:** `SuspensionId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `LastDateForWhichPaid` | Date | Yes |  |  | | Hidden |
| `FirstDateWorked` | Date | Yes |  |  | | Hidden |
| `UI19DownloadUrl` | String | Yes |  |  | | Hidden |
| `SuspensionId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `SuspensionReason` | String | No | Yes |  | `SuspensionReason` |  |
| `EffectiveDate` | DateTimeOffset | Yes | Conditional |  | |  |
| `ReturnDate` | DateTimeOffset | Yes | Conditional |  | |  |
| `MinimumPeriods` | Decimal | Yes |  |  | |  |
| `PercentageOfPay` | Decimal | Yes |  |  | |  |
| `OverridingAmount` | Decimal | Yes |  |  | |  |
| `SuspensionPaymentMethod` | PaySpace.Venuta.Data.Models.Enums.SuspensionPaymentMethod | Yes | Conditional |  | |  |
| `DoNotCalculateRecurringComponents` | Boolean | Yes |  |  | |  |
| `DenyOnceOffInput` | Boolean | Yes |  |  | |  |
| `ExpectedReturnType` | String | Yes |  |  | `ExpectedReturnType` | Hidden |
| `LastPayslipRun` | String | Yes |  |  | `LastPayslipRun` | Hidden |
| `FirstActivePayslipRun` | String | Yes |  |  | `FirstActivePayslipRun` | Hidden |
| `UpdateDataSnapshot` | Boolean | No |  |  | | Hidden |
| `TransactionType` | PaySpace.Venuta.Data.Models.Enums.TransactionType | Yes |  |  | | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
---

## EmployeeTableBuilder

**Primary Key:** `EmployeeTableBuilderId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No | Yes |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeTableBuilderId` (PK) | Int64 | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `TableBuilderCode` | String | Yes | Conditional |  | `TableBuilderComponentCompany/{ComponentCompany}` | 16 |  |
| `InitialBalance` | Decimal | No | Conditional |  | |  |  |
| `IncreaseDecrease` | Decimal | Yes |  |  | |  |  |
| `Balance` | Decimal | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `TableBuilderComponentCompany` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  | Read-only | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | |  |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
---

## EmployeeTakeOn
*Caption: Take On Year To Date Figures*

**Primary Key:** `Identifier`

| Property | Type | Nullable | Access | Lookup | Notes |
|----------|------|----------|--------|--------|-------|
| `Identifier` (PK) | String | No |  | |  |
| `EmployeeNumber` | String | Yes | Read-only | `EmployeeNumber` |  |
| `FullName` | String | Yes | Read-only | |  |
| `EmployeeTakeOnComponentCompany` | String | Yes | Read-only | `EmployeeTakeOnComponentCompany` |  |
| `PayslipAction` | PaySpace.Venuta.Data.Models.Enums.PayslipAction | No | Read-only | | Hidden |
| `TaxCode` | String | Yes | Read-only | |  |
| `ComponentValue` | Decimal | No |  | |  |
---

## EmployeeTaxCertificate

**Primary Key:** `TaxCertificateId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `YearStartDate` | Date | No |  |
| `YearEndDate` | Date | No |  |
| `TaxCertificateId` (PK) | String | No |  |
| `EmployeeNumber` | String | Yes |  |
| `CompanyName` | String | Yes |  |
| `Irp5Number` | String | Yes |  |
| `ManualNumber` | String | Yes |  |
| `DirectiveNumber` | String | Yes |  |
| `Irp5SourceCode` | String | Yes |  |
| `NewIrp5Numbers` | String | Yes |  |
| `DownloadUrl` | String | Yes |  |
---

## EmployeeTraining
*Caption: Training Records*

**Primary Key:** `TrainingId`

| Property | Type | Nullable | Required | Access | Lookup | Notes |
|----------|------|----------|----------|--------|--------|-------|
| `EndDate` | Date | Yes |  |  | |  |
| `CertificateExpiryDate` | Date | Yes |  |  | |  |
| `TrainingId` (PK) | Int64 | No |  |  | |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |
| `FullName` | String | Yes |  | Read-only | |  |
| `EffectiveDate` | DateTimeOffset | No | Yes |  | |  |
| `Course` | String | No | Yes |  | `Course` |  |
| `Status` | PaySpace.Venuta.Data.Models.Enums.EmployeeTrainingStatus | No | Yes | Create-only | |  |
| `TrainingNotCompleteStatus` | PaySpace.Venuta.Data.Models.Enums.TrainingNotCompleteStatus | Yes |  |  | |  |
| `CourseResult` | Decimal | Yes |  |  | |  |
| `DirectCourseCost` | Decimal | Yes |  |  | |  |
| `InvoiceNumber` | String | Yes |  |  | |  |
| `Facilitator` | String | Yes |  |  | |  |
| `Assessor` | String | Yes |  |  | |  |
| `CertificateNo` | String | Yes |  |  | |  |
| `EmployeeRate` | Decimal | Yes |  |  | |  |
| `TravelCost` | Decimal | Yes |  |  | |  |
| `Accommodation` | Decimal | Yes |  |  | |  |
| `FacilitatorCost` | Decimal | Yes |  |  | |  |
| `FoodBeverage` | Decimal | Yes |  |  | |  |
| `Additional5` | Decimal | Yes |  |  | |  |
| `Additional6` | Decimal | Yes |  |  | |  |
| `Notes` | String | Yes |  |  | |  |
| `Attachment` | Binary | Yes |  |  | |  |
| `AttachmentName` | String | Yes |  |  | | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |
| `AttachmentUrl` | String | Yes |  | Read-only | |  |
---

## EmployeeTravelBusinessUsage

**Primary Key:** `EmployeeTravelBusinessUsageId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeTravelBusinessUsageId` (PK) | Int64 | No |  |  | |  |  |
| `TravelBusinessUsageTaxPercentage` | PaySpace.Venuta.Data.Models.Enums.TravelBusinessUsageTaxPercentage | No |  |  | |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeUnion

**Primary Key:** `EmployeeUnionId`

**Navigation Properties:** `Values`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EffectiveDate` | Date | No | Yes |  | |  |  |
| `StartDate` | Date | Yes |  |  | |  |  |
| `EndDate` | Date | Yes |  |  | |  |  |
| `EmployeeUnionId` (PK) | Int64 | No |  |  | |  |  |
| `Union` | String | No | Yes |  | `Union` |  |  |
| `EmployeeNumber` | String | No | Yes | Create-only | `EmployeeNumber` |  |  |
| `FullName` | String | Yes |  | Read-only | |  |  |
| `ComponentCompany` | String | Yes |  |  | `ComponentCompanyDetail` |  | Hidden |
| `ComponentCompanyId` | Int64 | No |  |  | |  | Hidden |
| `ComponentIndicatorLine` | String | Yes |  |  | `ComponentIndicatorLine` |  |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `BcoePercentage` | Decimal | Yes |  |  | |  |  |
| `Comments` | String | Yes |  |  | | 250 |  |
---

## EmployeeWorkflow

**Primary Key:** `UserWorkflowStepId`

| Property | Type | Nullable | Access | Notes |
|----------|------|----------|--------|-------|
| `UserWorkflowStepId` (PK) | Int64 | No |  |  |
| `WorkflowComplete` | Boolean | No | Read-only |  |
| `PreviousApprovers` | Collection(Edm.String) | Yes | Read-only |  |
| `Approvers` | Collection(Edm.String) | Yes | Read-only |  |
| `NextApprovers` | Collection(Edm.String) | Yes | Read-only |  |
| `ActionDate` | DateTimeOffset | No |  |  |
---


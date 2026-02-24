# Other Entities (Payspace Metadata)

*16 entity types — auto-generated from OData $metadata*

## Contents

- [BudgetArchiveEmployee](#budgetarchiveemployee) (20 properties)
- [BudgetArchiveOrganization](#budgetarchiveorganization) (5 properties)
- [CostingProjectActivity](#costingprojectactivity) (18 properties)
- [CustomFieldDetail](#customfielddetail) (17 properties)
- [CustomFieldEditorOption](#customfieldeditoroption) (3 properties)
- [HousePaymentDetailComponentHistory](#housepaymentdetailcomponenthistory) (7 properties)
- [JobManagement](#jobmanagement) (38 properties)
- [JobManagementBudgetCost](#jobmanagementbudgetcost) (6 properties)
- [JobManagementHistory](#jobmanagementhistory) (5 properties)
- [LeaveAdjustment](#leaveadjustment) (13 properties)
- [PayRateCategory](#payratecategory) (10 properties)
- [ReportRequest](#reportrequest) (1 properties)
- [RequestTypeAttachment](#requesttypeattachment) (4 properties)
- [RosterDataset](#rosterdataset) (1 properties)
- [ShiftTypeDataset](#shifttypedataset) (9 properties)
- [TaxCalculationError](#taxcalculationerror) (5 properties)

---

## BudgetArchiveEmployee

**Primary Key:** `JobManagementId`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `AbolishDate` | Date | Yes | |  |
| `JobManagementId` (PK) | Int64 | No | |  |
| `JobNumber` | String | Yes | |  |
| `JobTitle` | String | Yes | |  |
| `OrganizationGroup` | String | Yes | `OrganizationGroup` |  |
| `OrganizationRegion` | String | Yes | `OrganizationRegion` |  |
| `Grade` | String | Yes | `Grade` |  |
| `EmploymentSubCategory` | String | Yes | `EmploymentSubCategory` |  |
| `EmployeeFirstName` | String | Yes | |  |
| `EmployeeLastName` | String | Yes | |  |
| `EmployeeNumber` | String | Yes | |  |
| `PositionType` | String | Yes | `PositionType` |  |
| `Comments` | String | Yes | |  |
| `OverrideEmployeeName` | Boolean | Yes | |  |
| `OverrideFigures` | Boolean | Yes | |  |
| `PlannedJobTitle` | String | Yes | |  |
| `PlannedGrade` | String | Yes | |  |
| `FilledCount` | Int32 | No | |  |
| `DoNotApplyIncrease` | Boolean | No | |  |
| `Figures` | Collection(PaySpace.Venuta.Data.Models.Dto.BudgetArchiveEmployeeFigure) | Yes | |  |
---

## BudgetArchiveOrganization

**Primary Key:** `OrganizationGroupId`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `OrganizationGroupId` (PK) | Int64 | No | |  |
| `OrganizationGroup` | String | Yes | `OrganizationGroup` |  |
| `FilledCount` | Int32 | No | |  |
| `JobCount` | Int32 | No | |  |
| `Figures` | Collection(PaySpace.Venuta.Data.Models.Dto.BudgetArchiveOrganizationFigure) | Yes | |  |
---

## CostingProjectActivity
*Caption: Project Costing*

**Primary Key:** `ProjectId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `InactiveDate` | Date | Yes |  |  | |  |  |
| `ProjectId` (PK) | Int64 | No |  |  | |  |  |
| `CompanyId` | Int64 | No |  | Create-only | |  | Hidden |
| `ProjectActivityDescription` | String | No | Yes |  | | 100 |  |
| `ProjectActivityCode` | String | No | Yes |  | | 50 |  |
| `Glcode` | String | Yes |  |  | | 50 |  |
| `OrganizationGroup` | String | Yes |  |  | `OrganizationGroup` |  |  |
| `OrganizationRegion` | String | Yes |  |  | `OrganizationRegion` |  |  |
| `KeyAccountManagerEmployeeNumber` | String | Yes |  |  | `KeyAccountManagerEmployeeNumber` |  |  |
| `MonthlyStaffBudget` | Int32 | Yes |  |  | |  |  |
| `MonthlyBudgetedHours` | Decimal | Yes |  |  | |  |  |
| `HourlyCost` | Decimal | Yes |  |  | |  |  |
| `DoNotShow` | Boolean | Yes |  | Create-only | |  | Hidden |
| `ActivityHourlyRate` | Decimal | Yes |  |  | |  | Hidden |
| `TsdefaultHours` | Decimal | Yes |  | Create-only | |  | Hidden |
| `MaxHours` | Decimal | Yes |  |  | |  |  |
| `OrganizationCategory` | String | Yes |  |  | `OrganizationCategory` | 50 |  |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
---

## CustomFieldDetail

**Primary Key:** `Code`

**Navigation Properties:** `Options`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Code` (PK) | String | No |  |
| `ParentFieldCode` | String | Yes |  |
| `Label` | String | Yes |  |
| `Type` | String | Yes |  |
| `IsRequired` | Boolean | No |  |
| `IsReadOnly` | Boolean | Yes |  |
| `ShowOnGrid` | Boolean | Yes |  |
| `CustomFieldId` | Int64 | No |  |
| `Level` | String | Yes |  |
| `Group` | String | Yes |  |
| `ValidationExpression` | String | Yes |  |
| `FieldGroupOrder` | Int32 | Yes |  |
| `OrderNumber` | Int32 | No |  |
| `CustomFormCategory` | String | Yes |  |
| `Tooltip` | String | Yes |  |
| `AdvancedSettings` | PaySpace.Venuta.Data.Models.Dto.AdvancedSettings | Yes |  |
| `CustomFormCode` | String | Yes |  |
---

## CustomFieldEditorOption

**Primary Key:** `Code`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Code` (PK) | String | No |  |
| `Description` | String | Yes |  |
| `DefaultValue` | Boolean | Yes |  |
---

## HousePaymentDetailComponentHistory

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Rated` | Boolean | Yes |  |
| `FloorArea` | Decimal | Yes |  |
| `PropertyValue` | Decimal | Yes |  |
| `AccommodationOption` | String | Yes |  |
| `HotelAccommodation` | String | Yes |  |
| `IsPropertyFurnished` | Boolean | Yes |  |
| `EmployeeDeduction` | Decimal | Yes |  |
---

## JobManagement
*Caption: Job Management*

**Primary Key:** `JobManagementId`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `EmployementStartDate` | Date | Yes |  | Read-only | |  |  |
| `EmployementEndDate` | Date | Yes |  | Read-only | |  |  |
| `ReportsToEffectiveDate` | Date | Yes |  |  | |  |  |
| `AbolishDate` | Date | Yes |  |  | |  |  |
| `JobManagementId` (PK) | Int64 | No |  |  | |  |  |
| `Status` | PaySpace.Venuta.Data.Models.Enums.JobManagementStatus | No |  | Read-only | |  |  |
| `IsPlannedJob` | Boolean | Yes |  |  | |  |  |
| `JobNumber` | String | Yes |  |  | |  |  |
| `ExternalParty` | Boolean | Yes |  | Read-only | |  |  |
| `CompanyFrequency` | String | Yes |  | Read-only | `CompanyFrequency` |  |  |
| `OrganizationPosition` | String | Yes |  | Read-only | |  | Hidden |
| `OrganizationPositionWithCode` | String | No | Yes |  | `OrganizationPositionWithCode` |  |  |
| `OrganizationPositionId` | Int64 | Yes |  |  | |  | Hidden |
| `OrganizationGroup` | String | Yes |  |  | `OrganizationGroup` |  | Hidden |
| `OrganizationGroupUploadCode` | String | No | Yes |  | `OrganizationGroupUploadCode` |  |  |
| `OrganizationGroupId` | Int64 | Yes |  |  | |  | Hidden |
| `OrganizationRegion` | String | Yes |  |  | `OrganizationRegion` |  |  |
| `JobTitle` | String | No | Yes |  | | 150 |  |
| `Grade` | String | Yes |  | Read-only | `Grade` |  | Hidden |
| `GradeCode` | String | No | Yes |  | `Grade` |  |  |
| `BudgetGroup` | String | No | Yes |  | `BudgetGroup` |  |  |
| `EmploymentSubCategory` | String | Yes |  |  | `EmploymentSubCategory` |  |  |
| `EmployeeFirstName` | String | Yes |  | Read-only | |  |  |
| `EmployeeLastName` | String | Yes |  | Read-only | |  |  |
| `EmployeeNumber` | String | Yes |  | Read-only | `EmployeeNumber` |  |  |
| `OverrideEmployeeName` | Boolean | Yes |  |  | |  | Hidden |
| `PositionType` | String | Yes |  |  | `PositionType` |  |  |
| `GeneralLedger` | String | Yes |  |  | `CompanyGlAccount` |  |  |
| `ReportsToJob` | String | Yes |  |  | `ReportsToJob` |  |  |
| `ReportsToJobManagementId` | Int64 | Yes |  |  | |  | Hidden |
| `PreferredGender` | String | Yes |  |  | `Gender` |  |  |
| `PreferredRace` | String | Yes |  |  | `PreferredRace` |  |  |
| `PreferredLocalization` | String | Yes |  |  | `PreferredLocalization` |  |  |
| `OverrideRoster` | String | Yes |  |  | `OverrideRoster` |  | Setting: `JOBROSTERLNK` |
| `OverrideOfoCode` | String | Yes |  |  | `OverrideOfoCode` |  |  |
| `OverrideOccupationalLevel` | String | Yes |  |  | `OverrideOccupationalLevel` |  |  |
| `UploadCode` | String | Yes |  |  | |  | Hidden |
| `CustomFields` | Collection(PaySpace.Venuta.Data.Models.Dto.CustomField) | Yes |  |  | |  |  |
---

## JobManagementBudgetCost

**Primary Key:** `CompanyComponentId`

| Property | Type | Nullable | Lookup | Notes |
|----------|------|----------|--------|-------|
| `CompanyComponentId` (PK) | Int64 | No | |  |
| `JobNumber` | String | Yes | |  |
| `ActionType` | String | Yes | |  |
| `EmployeeNumber` | String | Yes | `EmployeeNumber` |  |
| `CompanyComponent` | String | Yes | `CompanyComponent` |  |
| `CostAmount` | Decimal | No | |  |
---

## JobManagementHistory

**Primary Key:** `EmployeePositionId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `StartDate` | Date | No |  |
| `EndDate` | Date | Yes |  |
| `EmployeePositionId` (PK) | Int64 | No |  |
| `EmployeeNumber` | String | Yes |  |
| `FullName` | String | Yes |  |
---

## LeaveAdjustment

**Primary Key:** `LeaveAdjustmentId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `LeaveAdjustmentId` (PK) | Int64 | No |  |
| `EmployeeNumber` | String | Yes |  |
| `LeaveBucket` | String | Yes |  |
| `LeaveType` | PaySpace.Venuta.Data.Models.Enums.LeaveType | No |  |
| `LeaveStartDate` | DateTimeOffset | Yes |  |
| `LeaveEndDate` | DateTimeOffset | Yes |  |
| `NoOfDays` | Double | No |  |
| `LeaveStatus` | PaySpace.Venuta.Data.Models.Enums.LeaveStatus | No |  |
| `HasAttachment` | Boolean | No |  |
| `Comments` | String | Yes |  |
| `LeaveReason` | String | Yes |  |
| `Reference` | String | Yes |  |
| `TransactionType` | PaySpace.Venuta.Data.Models.Enums.LeaveEntryType | No |  |
---

## PayRateCategory
*Caption: Pay Rate Categories*

**Primary Key:** `CategoryId`

**Navigation Properties:** `AdditionalValues`

| Property | Type | Nullable | Required | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|------------|-------|
| `EffectiveDate` | Date | No | Yes | |  |  |
| `CategoryId` (PK) | Int64 | No |  | |  |  |
| `Code` | String | No | Yes | | 50 |  |
| `Description` | String | No | Yes | | 170 |  |
| `PayRateDefault` | Decimal | No | Yes | |  |  |
| `PayRateHoursPerDay` | Double | Yes |  | |  |  |
| `PayRateDaysPerPeriod` | Double | Yes |  | |  |  |
| `Grade` | String | Yes |  | `Grade` |  |  |
| `IsPayRateRecordLinked` | Boolean | Yes |  | |  |  |
| `Reason` | String | Yes |  | `Reason` |  |  |
---

## ReportRequest

**Primary Key:** `Id`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `Id` (PK) | String | No |  |
---

## RequestTypeAttachment

**Primary Key:** `CompanyAttachmentId`

| Property | Type | Nullable | Required | Notes |
|----------|------|----------|----------|-------|
| `CompanyAttachmentId` (PK) | Int64 | No |  |  |
| `SectionName` | String | No | Yes |  |
| `AttachmentTitle` | String | No | Yes |  |
| `AttachmentName` | String | No | Yes |  |
---

## RosterDataset

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `RosterType` | PaySpace.Venuta.Data.Models.Enums.RosterType | Yes |  |
---

## ShiftTypeDataset

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `MondayHours` | Decimal | Yes |  |
| `TuesdayHours` | Decimal | Yes |  |
| `WednesdayHours` | Decimal | Yes |  |
| `ThursdayHours` | Decimal | Yes |  |
| `FridayHours` | Decimal | Yes |  |
| `SaturdayHours` | Decimal | Yes |  |
| `SundayHours` | Decimal | Yes |  |
| `FullDayHours` | Decimal | Yes |  |
| `HoursPerWeek` | Decimal | Yes |  |
---

## TaxCalculationError

**Primary Key:** `TaxCalculationErrorId`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `LocalizedMessage` | String | Yes |  |
| `TaxCalculationErrorId` (PK) | Int64 | No |  |
| `CompanyRun` | String | Yes |  |
| `Message` | String | Yes |  |
| `IsError` | Boolean | No |  |
---


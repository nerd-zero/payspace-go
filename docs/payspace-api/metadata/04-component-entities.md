# Component Entities (Payspace Metadata)

*4 entity types — auto-generated from OData $metadata*

## Contents

- [ComponentCompany](#componentcompany) (49 properties)
- [ComponentCompanyIndicatorLine](#componentcompanyindicatorline) (4 properties)
- [ComponentCompanyValue](#componentcompanyvalue) (7 properties)
- [ComponentHistory](#componenthistory) (3 properties)

---

## ComponentCompany

**Primary Key:** `CompanyComponentId`

**Navigation Properties:** `IndicatorLines`, `Values`, `CompanyLeavePurchases`

| Property | Type | Nullable | Required | Access | Lookup | Max Length | Notes |
|----------|------|----------|----------|--------|--------|------------|-------|
| `BonusEffectiveDate` | Date | Yes |  |  | |  |  |
| `BonusPayDate` | Date | Yes |  |  | |  |  |
| `CompanyComponentId` (PK) | Int64 | No |  |  | |  |  |
| `ComponentBureau` | String | No | Yes | Read-only | `ComponentBureau` |  |  |
| `ComponentName` | String | No | Yes | Read-only | |  |  |
| `Formula` | String | Yes |  | Read-only | |  |  |
| `Description` | String | Yes |  |  | | 128 |  |
| `InPackage` | Boolean | Yes |  |  | |  |  |
| `ComponentCode` | String | Yes |  |  | | 50 |  |
| `MinValue` | Decimal | Yes |  |  | |  |  |
| `MaxValue` | Decimal | Yes |  |  | |  |  |
| `ActiveFromMonth` | String | Yes |  |  | `MonthsOfTheYear` |  |  |
| `ActiveToMonth` | String | Yes |  |  | `MonthsOfTheYear` |  |  |
| `AutoRecoveryType` | String | Yes |  |  | `AutoRecoveryType` |  |  |
| `MaxRatioRecovery` | Decimal | Yes |  |  | |  |  |
| `MultiplyByComponentCompany` | String | Yes |  |  | `Description/{Description}` |  |  |
| `RecoverFigures` | Boolean | Yes |  |  | |  |  |
| `Inactive` | Boolean | Yes |  |  | |  |  |
| `AddToEmployee` | Boolean | Yes |  |  | |  |  |
| `ProRata` | Boolean | Yes |  |  | |  |  |
| `DoNotConvertCurrency` | Boolean | Yes |  |  | |  |  |
| `Currency` | String | Yes |  |  | `Currency` |  |  |
| `DoNotShowOnPayslip` | Boolean | Yes |  |  | |  |  |
| `DoNotShowCurrency` | Boolean | Yes |  |  | |  |  |
| `IsOnceOff` | Boolean | Yes |  |  | |  |  |
| `AllowCompanyLevelAutoRecoveryChange` | Boolean | Yes |  |  | |  |  |
| `LanguageAliasDescription` | String | Yes |  |  | | 150 |  |
| `OverridingTaxCode` | String | Yes |  |  | `OverridingTaxCode` |  |  |
| `ComponentGroup` | String | Yes |  |  | `ComponentGroup` |  |  |
| `OnceOffValuesActSameAsPackage` | Boolean | Yes |  |  | |  |  |
| `ExclComponentRetrospective` | Boolean | Yes |  |  | |  |  |
| `IsCostToCompany` | Boolean | Yes |  |  | |  |  |
| `ShowOnMockPayslipOnly` | Boolean | Yes |  |  | |  |  |
| `HidePayslipComments` | Boolean | Yes |  |  | |  |  |
| `RetroDoNotCalc` | Boolean | Yes |  |  | |  |  |
| `CalcRegardlessOfBasicPosted` | Boolean | Yes |  |  | |  |  |
| `EnforcePackageRule` | Boolean | Yes |  |  | |  |  |
| `OverrideDecimalPlaces` | String | Yes |  |  | `DecimalPlace` |  |  |
| `PayslipMessage` | String | Yes |  |  | | 500 |  |
| `RunsToCompleteTheComponentIn` | Collection(Edm.String) | Yes |  |  | `PeriodsToCalculate` |  |  |
| `IsMonthlyByWeeks` | Boolean | Yes |  |  | |  |  |
| `IndicatorDescription` | String | Yes |  |  | |  |  |
| `WeeklyStartDay` | String | Yes |  |  | `WeeklyStartDay` |  |  |
| `CarVATPercentage` | Decimal | Yes |  |  | |  |  |
| `IsBonusTaxSpread` | Boolean | Yes |  |  | |  |  |
| `SpreadToDateOrTaxEnd` | PaySpace.Venuta.Data.Models.Enums.SpreadToDateOrTaxEnd | Yes |  |  | |  |  |
| `BonusPayDateType` | PaySpace.Venuta.Data.Models.Enums.BonusPayDateType | Yes |  |  | `BonusPayDateType` |  |  |
| `AutomaticallyPayBonus` | Boolean | Yes |  |  | |  |  |
| `PayBonusOnTermination` | Boolean | Yes |  |  | |  |  |
---

## ComponentCompanyIndicatorLine

**Primary Key:** `IndicatorLineId`

| Property | Type | Nullable | Required | Access | Notes |
|----------|------|----------|----------|--------|-------|
| `IndicatorLineId` (PK) | Int64 | No |  |  |  |
| `IndicatorId` | Int64 | No |  | Read-only |  |
| `IndicatorLineDescription` | String | No | Yes |  |  |
| `IndicatorLineValue` | Decimal | No | Yes |  |  |
---

## ComponentCompanyValue

**Primary Key:** `ComponentValueId`

| Property | Type | Nullable | Access | Notes |
|----------|------|----------|--------|-------|
| `ComponentValueId` (PK) | Int64 | No |  |  |
| `Description` | String | Yes | Read-only |  |
| `ComponentValue` | Decimal | Yes |  |  |
| `ComponentCompanyId` | Int64 | No | Create-only |  |
| `ValueTypeId` | Int32 | No |  |  |
| `Active` | Boolean | No |  |  |
| `CompanyRunId` | Int64 | Yes |  |  |
---

## ComponentHistory

**Primary Key:** `Id`

| Property | Type | Nullable | Notes |
|----------|------|----------|-------|
| `EffectiveDate` | Date | No |  |
| `Id` (PK) | Int64 | No |  |
| `Description` | String | Yes |  |
---


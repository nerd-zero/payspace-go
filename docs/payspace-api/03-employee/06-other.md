# Employee > Other

## Out of office

### Get collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeOutOfOffice?$$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee out of office records based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get collection** (Status: 200 )

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeOutOfOffice",
    "value": [
        {
            "OutOfOfficeId": 0,                         // Type="Int64"
            "EmployeeNumber": "string",                 // Type="String"
            "FullName": "string",                       // Type="String"
            "FromDate": "2022-01-01T00:00:00+02:00",    // Type="DateTime"
            "ToDate": "2022-01-01T00:00:00+02:00",      // Type="DateTime"
            "AltEmployeeNumber": "string"               // Type="String"
        }
    ]
}
```

---

### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeOutOfOffice`

Create a single employee out of office record for the specified `EmployeeNumber`.

**Example Response: Post** (Status: 200 OK)

```json
 {
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeOutOfOffice",
    "OutOfOfficeId": 0,                         // Type="Int64"
    "EmployeeNumber": "string",                 // Type="String"
    "FullName": "string",                       // Type="String"
    "FromDate": "2022-01-01T00:00:00+02:00",    // Type="DateTime"
    "ToDate": "2022-01-01T00:00:00+02:00",      // Type="DateTime"
    "AltEmployeeNumber": "string"               // Type="String"
}
```

---

### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeOutOfOffice({{OutOfOfficeId}})`

Update a single employee out of office record based on the specified `OutOfOfficeId`

**Example Response: Patch** (Status: 204 No Content)

---

### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeOutOfOffice({{OutOfOfficeId}})`

Delete single employee out of office record based on the specified `OutOfOfficeId`

**Example Response: Delete** (Status: 200 OK)

---
## Notes

### Get collection

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeNote?$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

Retrieves a list of employee notes based on the specified `company_id`.

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-uris |

**Example Response: Get a Collection** (Status: 200 OK)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeNote",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeNoteId": 0,                                    // Type="Int64"        
            "EmployeeNumber": "string",                             // Type="String"
            "FullName": "string",                                   // Type="String"
            "CaptureDateTime": "2010-11-09T07:37:53.047+02:00",     // Type="DateTime"
            "NoteTitle": "string",                                  // Type="String" 
            "NoteDescription": "string",                            // Type="String" 
            "IsReminderAttached": false,                            // Type="Boolean"
            "ReminderDateTime": "2010-11-09T07:39:29+02:00",        // Type="DateTime"
            "RemindEmployeeByEmail": false,                         // Type="Boolean"
            "RemindEmployeeBySms": false,                           // Type="Boolean"
            "RemindMeBySMS": false,                                 // Type="Boolean"
            "IsPublic": false,                                      // Type="Boolean"
            "SendToDirectlyReportsTo": null,                        // Type="Boolean"
            "ReminderRecipientEmails": "string",                    // Type="String" 
            "IsRecurringReminder": false,                           // Type="Boolean"
            "ReminderEndDateTime": null,                            // Type="DateTime"
            "ReminderRepeatEveryNumberOfDays": 0,                   // Type="Int32"
            "RecurringFrequency": null,                             // Type="String"
            "ReminderFrequencyDay": null,                           // Type="Int32"
            "RecurringWeekDays": null,                              // Type="String"
            "MonthsOfYear": null,                                   // Type="String"
            "HaveActioned": false,                                  // Type="Boolean"
            "IsRemindersSent": false                                // Type="Boolean"
        }
    ]
}
```

---

### Get a collection of all employee notes

**`GET`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeNote/all?$orderby={{employee-note-field}}&$top={{$top}}&$skip={{$skip}}&$count={{$count}}&$filter={{$filter}}`

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `$orderby` | `{{employee-note-field}}` | Optional (string) - Specifies the order in which items are returned |
| `$top` | `{{$top}}` | Optional (integer($int32)) - Limits the number of items returned from a collection |
| `$skip` | `{{$skip}}` | Optional  (integer($int32)) - Excludes the specified number of items of the queried collection from the result |
| `$count` | `{{$count}}` | Optional (true/false) - Return total count in response |
| `$filter` | `{{$filter}}` | Optional - Filter by field eg. "EmployeeNumber eq 'Emp01'". see https://docs.microsoft.com/en-us/dynamics-nav/using-filter-expressions-in-odata-urisa |

**Example Response: Get a collection of all employee notes**

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeNote",
    "value": [
        {
            "EmployeeNoteId": 1,                                        // Type="Int64"
            "EmployeeNumber": "string",                                 // Type="String"
            "FullName": "string",                                       // Type="String"
            "CaptureDateTime": "2019-01-01T00:00:00+02:00",             // Type="DateTime"
            "NoteTitle": "string",                                      // Type="String"
            "NoteDescription": "string",                                // Type="String"
            "IsReminderAttached": true,                                 // Type="Boolean"
            "ReminderDateTime": "2019-01-01T00:00:00+02:00",            // Type="DateTime"
            "RemindEmployeeByEmail": true,                              // Type="Boolean"
            "RemindEmployeeBySms": true,                                // Type="Boolean"
            "RemindMeBySMS": true,                                      // Type="Boolean"
            "IsPublic": true,                                           // Type="Boolean"
            "SendToDirectlyReportsTo": true,                            // Type="Boolean"
            "ReminderRecipientEmails": "string",                        // Type="String"
            "IsRecurringReminder": true,                                // Type="Boolean"
            "ReminderEndDateTime": "2019-01-01T00:00:00+02:00",         // Type="DateTime"
            "ReminderRepeatEveryNumberOfDays": 0,                       // Type="Int32"
            "RecurringFrequency": "string",                             // Type="String"
            "ReminderFrequencyDay": 0,                                  // Type="Int32"
            "RecurringWeekDays": "string",                              // Type="String"
            "MonthsOfYear": "string",                                   // Type="String"
            "HaveActioned": false,                                      // Type="Boolean"
            "IsRemindersSent": true                                     // Type="Boolean"
        }
    ]
}
```

---

### Create Record

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeNote`

Create a single `EmployeeNote` record for the specified `EmployeeNumber`.

**Example Response: Create Record** (Status: 201 Created)

```json
{
    "@odata.context": "{{api-base-url}}/odata/v2.0/{{company-id}}/$metadata#EmployeeNote",
    "@odata.count": 1,
    "value": [
        {
            "EmployeeNumber": "string",                                 // Type="String"
            "employeeNoteRecurring": null,                              // Type="Boolean"
            "haveActioned": false,                                      // Type="Boolean"
            "isPublic": false,                                          // Type="Boolean"
            "isRecurringReminder": false,                               // Type="Boolean"
            "isReminderAttached": false,                                // Type="Boolean"
            "isRemindersSent": false,                                   // Type="Boolean"
            "monthsOfYear": null,                                       // Type="String"
            "noteDescription": "string",                                // Type="String"
            "noteTitle": "string",                                      // Type="String"
            "RecurringFrequency": null,                                 // Type="String"
            "recurringWeekDays": null,                                  // Type="String"
            "remindEmployeeByEmail": false,                             // Type="Boolean"
            "remindEmployeeBySms": false,                               // Type="Boolean"
            "remindMeBySMS": false,                                     // Type="Boolean"
            "reminderDateTime": "2022-02-22T00:00:00+02:00",            // Type="DateTime"
            "reminderEndDateTime": null,                                // Type="DateTime"
            "reminderFrequencyDay": null,                               // Type="Int32"
            "reminderFrequencyMonth": null,                             // Type="Int32"
            "reminderRecipientEmails": "string",                        // Type="String"
            "reminderRepeatEveryNumberOfDays": null,                    // Type="Int32"
            "sendToDirectlyReportsTo": null                             // Type="Boolean"
        }
    ]
}
```

---

### Update Existing Record

**`PATCH`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeNote({{EmployeeNoteId}})`

Update a single `EmployeeNote` record for the specified `EmployeeNoteId`.

**Example Response: Update Existing Record** (Status: 204 No Content)

```json
{
  //see "EmployeeNote" in metadata endpoint for available fields
}
```

---

### Delete Existing Record

**`DELETE`** `{{api-base-url}}/odata/v2.0/:company-id/EmployeeNote({{EmployeeNoteId}})`

Delete a single `EmployeeNote` record for the specified `EmployeeNoteId`.

**Example Response: Delete Existing Record** (Status: 200 OK)

```json
{
  "@odata.context": "https://localhost:44393/odata/v2.0/158/$metadata#EmployeeNote/$entity",
  "Success": true
}
```

---

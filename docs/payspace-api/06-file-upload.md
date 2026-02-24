# File upload

This allows users to post attachments to a temporary location in order for the attachments to be attached to a record of another entity and save it permanenlty

- The temp attachment will only be valid for 1 hour after posting (It must be used on a supported endpoint within that time)
- The temp attachment can only be attached to a single record of any endpoint that supports attachments
- Only `EmployeeLeaveApplication` is supported at the moment
    

For the temp file to be attached to the intended record, the guid result of the file upload must be set as the value for the property `TempAttachmentSignature` on a supported endpoint, upon saving that record, the temp file will be moved to a permanent location

EndFragment

## File upload

**`POST`** `{{api-base-url}}/odata/v2.0/:company-id/FileUpload`

**Request Body (form-data):**

| Key | Type | Value | Description |
|-----|------|-------|-------------|
| `file` | file | `` | Form post file |

**Example Response: File upload** (Status: 200 )

```json
{
    "@odata.context": "https://localhost:44393/odata/v2.0//{{company-id}}/$metadata#FileUpload/$entity",
    "TempAttachmentSignature": "String" // Type="String" 
}
```

---

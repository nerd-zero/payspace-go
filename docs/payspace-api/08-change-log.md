# Change Log

# V2.0

###### Lookup endpoint

Lookups are now on separate enpoints that is prefixed with `Lookup`

eg. `{{api-base-url}}/odata/v2.0/:company-id/CompanyFrequency` changed to `{api-base-url}}/odata/v2.0/:company-id/Lookup/CompanyFrequency`

###### Date format

All dates with the Type `Edm.Date`(see Metadata) will be returned in the format `yyyy-MM-dd` and the same format is expected when adding or updating records

> An invalid date format will result in "Invalid date format for 'fieldName'. Expected 'yyyy-MM-dd'" validation error. 
  

---

# v1.1

###### EmployeePosition Create/Update

- To create or update an `EmployeePosition` record linked to a Organization Position, use the `OrganizationPositionWithCode` Field.
    
- The options for the field `OrganizationPositionWithCode` can be retrieved from the `OrganizationPositionWithCode` lookup.
    

> From API v1.1 the field `OrganizationPosition` is readonly. 
  

###### EmployeePosition OrganizationGroups

- The `OrganizationGroups` field will no longer be returned by default and is no longer a string collection.
    
- To return the `OrganizationGroups` field data, use the $expand query string functionality. eg. $expand=OrganizationGroups
    
- The `OrganizationGroups` field is a readonly colection of `OrganizationUnit` records


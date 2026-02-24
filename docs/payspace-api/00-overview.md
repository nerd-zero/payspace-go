# Deel Local Payroll API v2.0

## Documentation Index

### Endpoint Reference (from Postman collection)
- [Authentication](01-authentication.md)
- [Metadata](02-metadata.md)
- [Employee Endpoints](03-employee.md) (261 endpoints across 9 sub-sections)
- [Company Endpoints](04-company.md) (133 endpoints across 9 sub-sections)
- [Lookup Values](05-lookup-values.md) (121 endpoints)
- [File Upload](06-file-upload.md)
- [Webhooks](07-webhooks.md)
- [Change Log](08-change-log.md)

### Entity Metadata (from live OData $metadata - 176 entities)
- [Metadata Index](metadata/README.md)
- [Employee Entities](metadata/01-employee-entities.md) (61 types - field definitions, types, required/editable/lookup)
- [Company Entities](metadata/02-company-entities.md) (59 types)
- [Organization Entities](metadata/03-organization-entities.md) (6 types)
- [Component Entities](metadata/04-component-entities.md) (4 types)
- [Lookup Datasets](metadata/05-lookup-datasets.md) (30 types)
- [Other Entities](metadata/06-other-entities.md) (16 types)

---

Welcome to the Deel Local Payroll API.

Our powerful API allows you to access your employee data in order to utilize in your business environment. You can also push data into Deel Local Payroll from third party systems.

> This documentation pertains to v2.0 of the Deel Local Payroll API and all URL's references v2.0.
> 
> Documentation for version 1.1 can be found [HERE](https://developer-v1-1.payspace.com/).  
> Documentation for version 1.0 can be found [HERE](https://developer-v1.payspace.com/).  
> For breaking changes in API versions read [HERE](#6dc9fb4b-f0ee-4f21-a6e6-c0b018f23f6c).  
> Older version may not be updated 
  

The respective endpoint's support [ODATA](https://www.odata.org/getting-started/basic-tutorial/) queries which allows you to filter your queries so that only the necessary data is returned.

An access token is required to access all endpoints and can be obtained using the Authentication endpoint.

Example responses are provided down the right hand side for each endpoint.

If you would like to have Deel Local Payroll notify you of a change on an employee's profile, please refer to our [Webhooks documentation](#28ddfcdd-fe2e-4d29-b966-e063c5e04555).

---

# API Urls

### Testing

When testing or developing use the following base urls:

- Authentication: `https://staging-identity.yourhcm.com`
    
- API: `https://apistaging.payspace.com`
    

### Production

For production use the following base urls:

- Authentication: `https://identity.yourhcm.com`
    
- API: `https://api.payspace.com`
    

---

# Paging data

All `GET` endpoints by default returns the top **100** records. To get more than a **100** records use a combination of the `$skip` and `$top` query parameters. The maximum value for `$top` is **100**.

> Note that the total record count can be returned in the response body (`@odata.count`) by adding the $count parameter as described in each end point. 
  

Examples:

| Description | `$top` | `$skip` |
| --- | --- | --- |
| To get data on page 1 with a page size of 10 | 10 | 0 |
| To get data on page 2 with a page size of 10 | 10 | 10 |
| To get data on page 3 with a page size of 100 | 100 | 200 |

---

# Rate Limits

API access rate limits are applied at a per-key basis in unit time. Access to the production API using a key is limited to **2000 requests per minute**. In addition, every API response is accompanied by the following set of headers to identify the status of your consumption.

| Header | Description |
| --- | --- |
| `X-RateLimit-Limit` | The maximum number of requests that the consumer is permitted to make per minute. |
| `X-RateLimit-Remaining` | The number of requests remaining in the current rate limit window. |
| `X-RateLimit-Reset` | The time at which the current rate limit window resets. |

| Environment | Rate Limit |
| --- | --- |
| Production | 2000 |
| Staging | 200 |

Once you hit the rate limit, you will receive a response error with a status code of `429 Too Many Requests`.

---

# Response Errors

###### Not found Error

---

When a record could not by found. eg. incorrect key

Status Code: `404 Not found`

Response body: _Empty_

###### Unauthorized

---

Invalid or no token

Status Code: `401 Unauthorized`

Response body: _Empty_

###### Validation Errors

---

When creating or updating a record and not all required values are present or when any other validation fails

Status Code: `400 Bad Request`

Response body example:

```
{
    "@odata.context": "....",
    "Message": "'Quantity' must not be empty.\r\n'Description' must not be empty.",
    "Details": [
        {
            "Message": "'Quantity' must not be empty."
        },
        {
            "Message": "'Description' must not be empty."
        }
    ],
    "Success": false
}

 ```

###### Request body JSON Error

---

When posting an incorrect JSON payload

Status Code: `400 Bad Request`

Response body example:

```
{
    "@odata.context": "....",
    "Message": "Unexpected end when reading JSON. Path '', line 2, position 0.",
    "Details": [
        {
            "Message": "Unexpected end when reading JSON. Path '', line 2, position 0."
        }
    ],
    "Success": false
}

 ```

---

# User Agent

Include the value `payspace.com` as an agent string in the HTTP header `User-Agent` for all requests.

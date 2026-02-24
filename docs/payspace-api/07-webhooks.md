# Webhooks

## What are webhooks?

Webhooks are one way that applications can send automated messages or information to other applications when something happens. In the case of PaySpace, this "something" is when a change takes place on an employee's profile. Instead of polling our API endpoints periodically to check if something has changed on employee's profiles, you want to be notified when something changes - this is where webhooks come in. Webhooks have a message, or payload, and are sent to a unique URL, essentially the destination applications "phone number" or "address". 

Webhooks are almost always faster than polling and require much less work on your end. They're much like SMS notifications. Say your bank sends you an SMS when you make a new purchase. You already told the bank your phone number, so they knew where to send the message. They type out "You just spent R10 at NewStore" and send it to your phone number +27-82-000-0000. Something happened at your bank, and you got a message about it. All is well.

## Webhook configuration

You can configure a webhook on select employee screens in PaySpace. Visit our integrations setup screen (```Config > Basic Settings > General Company > Integrations```) and click on the "New Webhook Configuration" button, in the object drop down will be a list of available screens to setup a webhook on. The destination URL field is where you need to enter the secure URL you want PaySpace to send the payload to. The payload will be a JSON object of the full record that was added or changed. For example, if you setup a webhook on an Employee's Basic Profile screen, when an employee's address is changed by any user, a JSON object with all the employee fields on this screen will be sent to the URL you specified. You can then use this information to populate any other 3rd party system.

<a href="https://Webhook.site/" target="_blank">Webhook.site</a> is a free online toolset that generates a unique URL that allows you to inspect and test the JSON payload that is sent from PaySpace. Simply enter this URL into PaySpace and make a change on the relevant screen and you will be able to view the payload.
## Webhook security

The URL you specify must be a secure https URL which will ensure that all the information sent, is encrypted. The secret key that you can generate on the integrations setup screen is automatically included in the encrypted payload that is sent from PaySpace and can be used to ensure that you accept and process only payloads that have this secret key. This ensures that you don't accept and process any information from somebody with malicious intent.
## Get failed webhook requests

**`GET`** `{{api-base-url}}/v2.0/:companyid/WebhookError/:EntityType?from={{fromDate}}&to={{toDate}}&pageNumber={{pageNumber}}&pageSize={{pageSize}}`

Retrieve collection of failed webhook requests for a specified time period. An email address can be advised when configuring a webhook within the application. An email will be sent each day with a summary list of all errors for the previous day to notify of any failures.

> Note: This is not an OData Endpoint, Filtering on all the fields is not possible. 
  

**Response Example**

``` json
[
    {
        "EmployeeNumber": "string",                // Type="String"
        "CompanyId": 0,                            // Type="Int64"
        "EntityType": "Employee",                  // Type="String"
        "Timestamp": "2023-05-15T09:45:16+02:00",  // Type="DateTime"
        "Error": "string",                         // Type="string"
        "RequestTime": 0.0,                        // Type="String"
        "WebhookUrl": "string",                    // Type="string"
        "WebhookName": "string",                   // Type="String"
        "Data": {
            ...                                    // Webhook data. Can change depending on EntityType
        }
    }
]

```

**Custom Headers:**

| Key | Value | Description |
|-----|-------|-------------|
| `content-type` | `application/json` |  |

**Query Parameters:**

| Name | Value | Description |
|------|-------|-------------|
| `from` | `{{fromDate}}` | date. example: 2023-05-15 |
| `to` | `{{toDate}}` | date. example: 2023-05-15 |
| `pageNumber` | `{{pageNumber}}` |  |
| `pageSize` | `{{pageSize}}` | the maximum is 100 records |

---

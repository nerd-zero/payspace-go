# Authentication

The `client_id` and `client_secret` are unique and can be created by navigating to Config > Basic Settings > General Company > Integrations and then select the API Credentials tab. The API secret will be displayed after saving and must be copied and saved immediately as you will not be able to retrieve this secret once you navigate away from this page. If you require API credentials for the test environment, the same process must be followed on the test environment by logging into test.payspace.com.

The response will contain a list of companies that the access token has access to. The `company_id` or `company_code` must be supplied in all endpoints.

The access token will expire after 1 hour.

## Get Token

**`POST`** `{{identity-base-url}}/connect/token`

The `client_id`, `client_secret` and `scope` can be send via POST using a `form-data` or `x-www-form-urlencoded` body. Note that the `content-type` header depends on the body type.

**Request Body (x-www-form-urlencoded):**

| Key | Value | Description |
|-----|-------|-------------|
| `client_id` | `{{client-id}}` | Obtained from PaySpace |
| `client_secret` | `{{client-secret}}` | Obtained from PaySpace |
| `scope` | `{{api-scope}}` | Defaults to `api.full_access`<br>Options: `api.read_only`, `api.full_access`, `api.update`, `api.create` |

**Example Response: Get Token Example** (Status: 200 OK)

```json
{
  "access_token": "string",
  "expires_in": 3600,
  "token_type": "Bearer",
  "group_companies": [
    {
      "group_id": 0,
      "group_description": "string",
      "companies": [
        {
          "company_id": 0,
          "company_name": "string",
          "company_code": "string"
        },
        {
          "company_id": 0,
          "company_name": "string",
          "company_code": "string"
        }
      ]
    }
  ]
}
```

---

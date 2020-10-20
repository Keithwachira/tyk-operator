# API Definitions

## Migrating Existing APIs

If there are existing APIs that you want to link to a CRD, it's very easy to do.  Simply add the `api_id` from your API Definition to the YAML of your `ApiDefinition` type and the Operator will take care of the rest.

Example

1.  From the existing API Definition, grab the following field:
```
"api_id": "5e0fac4845bb46c77543be28300fd9d7"
```

2. And then simply add this value to your YAML, in the `spec.api_id` field:

```crd
apiVersion: tyk.tyk.io/v1alpha1
kind: ApiDefinition
metadata:
  name: my-existing-api
spec:
  api_id: 5e0fac4845bb46c77543be28300fd9d7
  name: existing API
  protocol: http
  active: true
  proxy:
    target_url: http://httpbin.org
    listen_path: /httpbin
    strip_listen_path: true
```

3. And then apply your changes:

```
$ kubectl apply -f config/samples/httpbin_protected.yaml
apidefinition.tyk.tyk.io/my-existing-api created
```

Note, the source of truth for the API definition is now the CRD, meaning it will override any differences in your existing API definition.

Any keys, policies and analytics linked to this API will continue to operate the same. 

## Support 
An API Definition describes the configuration of an API. It instructs Tyk Gateway how to configure the API.

✅ - Fully supported

⚠️ - Requires documentation / testing

❌️ - Not currently supported

## API Level

### Authentication

| Type | Supported |
| ----------- | --------- |
| [Static Bearer Token](#) | ✅ |
| [Keyless (open)](#) | ✅ |
| [JWT](#) | ⚠️ |
| [OpenID Connect](#) | ⚠️ |
| [OAuth2](#) | ⚠️ |

### Cross Origin Resource Sharing (CORS)

TBA

### API Webhooks

- [Webhooks](./api_definitions/webhooks.md) - [❌️](https://github.com/TykTechnologies/tyk-operator/issues/62) 

### Custom Plugins

| Type | Supported |
| ----------- | --------- |
| [Go](./api_definitions/custom_plugin.md) | ⚠️ |
| [gRPC](./api_definitions/custom_plugin.md) | ⚠️ |
| [Javascript](./api_definitions/custom_plugin.md) | ✅ |
| [Lua](./api_definitions/custom_plugin.md) | ⚠️ |
| [Python](./api_definitions/custom_plugin.md) | ⚠️ |

## Endpoint level

Here is the list of supported middleware that you can apply at the endpoint level.

| Endpoint Middleware  | Supported |
| ----------- | --------- |
| [Analytics - Endpoint Tracking](#) | ⚠️ |
| [Availability - Circuit Breaker](./../config/samples/httpbin_timeout.yaml) | [❌️](# "Incompatible types string vs float64") |
| [Availability - Enforced Timeouts](./../config/samples/httpbin_timeout.yaml) | ✅ |
| [Headers - Global Request Add](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [Headers - Global Request Remove](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [Headers - Global Response Add](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [Headers - Global Response Remove](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [Performance - Cache](./../config/samples/httpbin_cache.yaml) | ✅ |
| [Plugin - Virtual Endpoint](./api_definitions/custom_plugin.md) | ✅ |
| [Security - Allow list](#) | ⚠️ |
| [Security - Block list](#) | ⚠️ |
| [Security - Ignore list](#) | ⚠️ |
| [Transform - Internal](#) | ⚠️ |
| [Transform - Method](#) | ⚠️ |
| [Transform - Mock](#) | ⚠️ |
| [Transform - Request Body](../config/samples/httpbin_transform.yaml) | ✅ |
| [Transform - Response Body](../config/samples/httpbin_transform.yaml) | ✅ |
| [Transform - Request Body JQ](../config/samples/httpbin_transform.yaml) | [⚠️](# "Requires JQ on Gateway Host & Testing") |
| [Transform - Response Body JQ](../config/samples/httpbin_transform.yaml) | [⚠️](# "Requires JQ on Gateway Host & Testing") |
| [Transform - URL Rewrite](#) | ⚠️ |
| [Validate - JSON Schema](../config/samples/httpbin_validate.yaml) | [❌️](https://github.com/TykTechnologies/tyk-operator/issues/59) |
| [Validate - Limit Request Size](../config/samples/httpbin_validate.yaml) | [❌️](https://github.com/TykTechnologies/tyk-operator/issues/59) |
# API Definitions

An [API Definitions](https://tyk.io/docs/getting-started/key-concepts/what-is-an-api-definition/) is a reverse proxy definition in the world of Tyk.

Here is the list of currently supported plugins/middleware that you can apply in an [API definition](https://tyk.io/docs/getting-started/key-concepts/what-is-an-api-definition/).

| Middleware  | Supported |
| ----------- | --------- |
| [Cache](./../config/samples/httpbin_cache.yaml) | ✅ |
| [Custom Plugins](./api_definitions/custom_plugin.md) | ✅ |
| [Headers - Global Request Add](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [Headers - Global Request Remove](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [Headers - Global Response Add](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [Headers - Global Response Remove](../config/samples/httpbin_global-headers.yaml) | ✅ |
| [JSON Schema Validation](../config/samples/httpbin_validate.yaml) | [❌️](# "Unable to convert JSON Schema map[string]interface{} to CRD") |
| [Transform - Request Body](../config/samples/httpbin_transform.yaml) | ✅ |
| [Transform - Response Body](../config/samples/httpbin_transform.yaml) | ✅ |
| [Transform - Request Body JQ](../config/samples/httpbin_transform.yaml) | [⚠️](# "Requires JQ on Gateway Host & Testing") |
| [Transform - Response Body JQ](../config/samples/httpbin_transform.yaml) | [⚠️](# "Requires JQ on Gateway Host & Testing") |
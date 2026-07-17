# \RoutingRulesAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoutingRules**](RoutingRulesAPI.md#CreateRoutingRules) | **Post** /v1/routing-rules | Create a single routing rule
[**DeleteRoutingRules**](RoutingRulesAPI.md#DeleteRoutingRules) | **Delete** /v1/routing-rules/{name} | Delete a single routing rule
[**GetRoutingRules**](RoutingRulesAPI.md#GetRoutingRules) | **Get** /v1/routing-rules/{name} | Get a single routing rule
[**ListRoutingRules**](RoutingRulesAPI.md#ListRoutingRules) | **Get** /v1/routing-rules | List routing rules
[**UpdateRoutingRules**](RoutingRulesAPI.md#UpdateRoutingRules) | **Put** /v1/routing-rules/{name} | Update a single routing rule



## CreateRoutingRules

> CreateRoutingRules(ctx).RoutingRuleXO(routingRuleXO).Execute()

Create a single routing rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func main() {
	routingRuleXO := *sonatyperepo.NewRoutingRuleXO() // RoutingRuleXO | A routing rule configuration

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RoutingRulesAPI.CreateRoutingRules(context.Background()).RoutingRuleXO(routingRuleXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesAPI.CreateRoutingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoutingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **routingRuleXO** | [**RoutingRuleXO**](RoutingRuleXO.md) | A routing rule configuration | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoutingRules

> DeleteRoutingRules(ctx, name).Execute()

Delete a single routing rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func main() {
	name := "name_example" // string | The name of the routing rule to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RoutingRulesAPI.DeleteRoutingRules(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesAPI.DeleteRoutingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the routing rule to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRoutingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoutingRules

> GetRoutingRules(ctx, name).Execute()

Get a single routing rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func main() {
	name := "name_example" // string | The name of the routing rule to get

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RoutingRulesAPI.GetRoutingRules(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesAPI.GetRoutingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the routing rule to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoutingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoutingRules

> ListRoutingRules(ctx).Execute()

List routing rules

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RoutingRulesAPI.ListRoutingRules(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesAPI.ListRoutingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRoutingRulesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRoutingRules

> UpdateRoutingRules(ctx, name).RoutingRuleXO(routingRuleXO).Execute()

Update a single routing rule

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v3"
)

func main() {
	name := "name_example" // string | The name of the routing rule to update
	routingRuleXO := *sonatyperepo.NewRoutingRuleXO() // RoutingRuleXO | A routing rule configuration

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RoutingRulesAPI.UpdateRoutingRules(context.Background(), name).RoutingRuleXO(routingRuleXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RoutingRulesAPI.UpdateRoutingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the routing rule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRoutingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routingRuleXO** | [**RoutingRuleXO**](RoutingRuleXO.md) | A routing rule configuration | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


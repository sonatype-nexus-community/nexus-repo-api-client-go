# \SecurityManagementUserTokensAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List1**](SecurityManagementUserTokensAPI.md#List1) | **Get** /v1/security/user-tokens/tokens | List user tokens or look up a token owner by namecode
[**ResetAllUserTokens**](SecurityManagementUserTokensAPI.md#ResetAllUserTokens) | **Delete** /v1/security/user-tokens | 
[**ServiceStatus**](SecurityManagementUserTokensAPI.md#ServiceStatus) | **Get** /v1/security/user-tokens | Show if the user token capability is enabled or not
[**SetServiceStatus**](SecurityManagementUserTokensAPI.md#SetServiceStatus) | **Put** /v1/security/user-tokens | 



## List1

> List1(ctx).Realm(realm).UserId(userId).Namecode(namecode).IncludeExpired(includeExpired).Skip(skip).Limit(limit).Execute()

List user tokens or look up a token owner by namecode



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
	realm := "realm_example" // string | Filter by realm (optional). Also acts as a realm constraint in namecode lookup mode. (optional)
	userId := "userId_example" // string | Filter by user ID (optional). Ignored when 'namecode' is provided. (optional)
	namecode := "namecode_example" // string | Look up the token owner by nameCode. When present, the endpoint resolves a single token and returns 404 if not found. Expired tokens are always included in lookup results. (optional)
	includeExpired := true // bool | Include expired tokens in list results (default: false). Ignored in namecode lookup mode — expired tokens are always included. (optional)
	skip := int32(56) // int32 | Number of items to skip for pagination (default: 0). Ignored in lookup mode. (optional)
	limit := int32(56) // int32 | Maximum number of items to return (default: 25, max: 100). Ignored in lookup mode. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUserTokensAPI.List1(context.Background()).Realm(realm).UserId(userId).Namecode(namecode).IncludeExpired(includeExpired).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUserTokensAPI.List1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiList1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **realm** | **string** | Filter by realm (optional). Also acts as a realm constraint in namecode lookup mode. | 
 **userId** | **string** | Filter by user ID (optional). Ignored when &#39;namecode&#39; is provided. | 
 **namecode** | **string** | Look up the token owner by nameCode. When present, the endpoint resolves a single token and returns 404 if not found. Expired tokens are always included in lookup results. | 
 **includeExpired** | **bool** | Include expired tokens in list results (default: false). Ignored in namecode lookup mode — expired tokens are always included. | 
 **skip** | **int32** | Number of items to skip for pagination (default: 0). Ignored in lookup mode. | 
 **limit** | **int32** | Maximum number of items to return (default: 25, max: 100). Ignored in lookup mode. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetAllUserTokens

> ResetAllUserTokens(ctx).Execute()



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
	r, err := apiClient.SecurityManagementUserTokensAPI.ResetAllUserTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUserTokensAPI.ResetAllUserTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetAllUserTokensRequest struct via the builder pattern


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


## ServiceStatus

> UserTokensApiModel ServiceStatus(ctx).Execute()

Show if the user token capability is enabled or not

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
	resp, r, err := apiClient.SecurityManagementUserTokensAPI.ServiceStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUserTokensAPI.ServiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceStatus`: UserTokensApiModel
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementUserTokensAPI.ServiceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceStatusRequest struct via the builder pattern


### Return type

[**UserTokensApiModel**](UserTokensApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetServiceStatus

> UserTokensApiModel SetServiceStatus(ctx).Body(body).Execute()



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
	body := *sonatyperepo.NewUserTokensApiModel() // UserTokensApiModel |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementUserTokensAPI.SetServiceStatus(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUserTokensAPI.SetServiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetServiceStatus`: UserTokensApiModel
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementUserTokensAPI.SetServiceStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetServiceStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserTokensApiModel**](UserTokensApiModel.md) |  | 

### Return type

[**UserTokensApiModel**](UserTokensApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


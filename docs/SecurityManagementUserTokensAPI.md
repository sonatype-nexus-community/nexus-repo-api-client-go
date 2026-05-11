# \SecurityManagementUserTokensAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List1**](SecurityManagementUserTokensAPI.md#List1) | **Get** /v1/security/user-tokens/tokens | List all user tokens
[**ResetAllUserTokens**](SecurityManagementUserTokensAPI.md#ResetAllUserTokens) | **Delete** /v1/security/user-tokens | 
[**ServiceStatus**](SecurityManagementUserTokensAPI.md#ServiceStatus) | **Get** /v1/security/user-tokens | Show if the user token capability is enabled or not
[**SetServiceStatus**](SecurityManagementUserTokensAPI.md#SetServiceStatus) | **Put** /v1/security/user-tokens | 



## List1

> UserTokenListXO List1(ctx).Realm(realm).UserId(userId).IncludeExpired(includeExpired).Skip(skip).Limit(limit).Execute()

List all user tokens



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
	realm := "realm_example" // string | The realm of the user (optional for cloud, required for self-hosted) (optional)
	userId := "userId_example" // string | Filter by user ID (optional) (optional)
	includeExpired := true // bool | Include expired tokens (default: false) (optional)
	skip := int32(56) // int32 | Number of items to skip for pagination (default: 0) (optional)
	limit := int32(56) // int32 | Maximum number of items to return (default: 25, max: 100) (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementUserTokensAPI.List1(context.Background()).Realm(realm).UserId(userId).IncludeExpired(includeExpired).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUserTokensAPI.List1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List1`: UserTokenListXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementUserTokensAPI.List1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiList1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **realm** | **string** | The realm of the user (optional for cloud, required for self-hosted) | 
 **userId** | **string** | Filter by user ID (optional) | 
 **includeExpired** | **bool** | Include expired tokens (default: false) | 
 **skip** | **int32** | Number of items to skip for pagination (default: 0) | 
 **limit** | **int32** | Maximum number of items to return (default: 25, max: 100) | 

### Return type

[**UserTokenListXO**](UserTokenListXO.md)

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


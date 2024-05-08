# \SecurityManagementUserTokensAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResetAllUserTokens**](SecurityManagementUserTokensAPI.md#ResetAllUserTokens) | **Delete** /v1/security/user-tokens | Invalidate all existing user tokens.
[**ServiceStatus**](SecurityManagementUserTokensAPI.md#ServiceStatus) | **Get** /v1/security/user-tokens | Show if the user token capability is enabled or not
[**SetServiceStatus**](SecurityManagementUserTokensAPI.md#SetServiceStatus) | **Put** /v1/security/user-tokens | Enable/Disable the user token capability



## ResetAllUserTokens

> ResetAllUserTokens(ctx).Execute()

Invalidate all existing user tokens.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
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
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
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

Enable/Disable the user token capability

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
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


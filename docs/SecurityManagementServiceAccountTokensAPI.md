# \SecurityManagementServiceAccountTokensAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityServiceAccountTokens**](SecurityManagementServiceAccountTokensAPI.md#CreateSecurityServiceAccountTokens) | **Post** /v1/security/service-account-tokens | Create a new service account token
[**DeleteSecurityServiceAccountTokens**](SecurityManagementServiceAccountTokensAPI.md#DeleteSecurityServiceAccountTokens) | **Delete** /v1/security/service-account-tokens/{id} | Hard-delete a service account and its associated token
[**ListSecurityServiceAccountTokens**](SecurityManagementServiceAccountTokensAPI.md#ListSecurityServiceAccountTokens) | **Get** /v1/security/service-account-tokens | List all service account tokens (never exposes secret material)



## CreateSecurityServiceAccountTokens

> ServiceAccountCreatedXO CreateSecurityServiceAccountTokens(ctx).ServiceAccountCreateXO(serviceAccountCreateXO).Execute()

Create a new service account token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397"
)

func main() {
	serviceAccountCreateXO := *sonatyperepo.NewServiceAccountCreateXO("Name_example", "RoleId_example") // ServiceAccountCreateXO | Service account to create

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementServiceAccountTokensAPI.CreateSecurityServiceAccountTokens(context.Background()).ServiceAccountCreateXO(serviceAccountCreateXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementServiceAccountTokensAPI.CreateSecurityServiceAccountTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityServiceAccountTokens`: ServiceAccountCreatedXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementServiceAccountTokensAPI.CreateSecurityServiceAccountTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityServiceAccountTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountCreateXO** | [**ServiceAccountCreateXO**](ServiceAccountCreateXO.md) | Service account to create | 

### Return type

[**ServiceAccountCreatedXO**](ServiceAccountCreatedXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityServiceAccountTokens

> DeleteSecurityServiceAccountTokens(ctx, id).Execute()

Hard-delete a service account and its associated token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397"
)

func main() {
	id := "id_example" // string | Service account id

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementServiceAccountTokensAPI.DeleteSecurityServiceAccountTokens(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementServiceAccountTokensAPI.DeleteSecurityServiceAccountTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Service account id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityServiceAccountTokensRequest struct via the builder pattern


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


## ListSecurityServiceAccountTokens

> ServiceAccountXO ListSecurityServiceAccountTokens(ctx).Execute()

List all service account tokens (never exposes secret material)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementServiceAccountTokensAPI.ListSecurityServiceAccountTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementServiceAccountTokensAPI.ListSecurityServiceAccountTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityServiceAccountTokens`: ServiceAccountXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementServiceAccountTokensAPI.ListSecurityServiceAccountTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityServiceAccountTokensRequest struct via the builder pattern


### Return type

[**ServiceAccountXO**](ServiceAccountXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


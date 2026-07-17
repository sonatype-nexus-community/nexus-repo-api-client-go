# \SecurityManagementSAMLUsersAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecuritySamlUsers**](SecurityManagementSAMLUsersAPI.md#CreateSecuritySamlUsers) | **Post** /v1/security/saml/users | Create a new SAML user. This allows administrators to pre-create SAML users with roles before their first login.
[**DeleteSecuritySamlUsers**](SecurityManagementSAMLUsersAPI.md#DeleteSecuritySamlUsers) | **Delete** /v1/security/saml/users/{userId} | Delete a SAML user.
[**GetSecuritySamlUsers**](SecurityManagementSAMLUsersAPI.md#GetSecuritySamlUsers) | **Get** /v1/security/saml/users/{userId} | Retrieve a SAML user by userId.
[**ListSecuritySamlUsers**](SecurityManagementSAMLUsersAPI.md#ListSecuritySamlUsers) | **Get** /v1/security/saml/users | Retrieve a list of SAML users. The response is limited to 1,000 users.
[**UpdateSecuritySamlUsers**](SecurityManagementSAMLUsersAPI.md#UpdateSecuritySamlUsers) | **Put** /v1/security/saml/users/{userId} | Update a SAML user&#39;s roles and attributes.



## CreateSecuritySamlUsers

> CreateSecuritySamlUsers(ctx).SamlUserXO(samlUserXO).Execute()

Create a new SAML user. This allows administrators to pre-create SAML users with roles before their first login.

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
	samlUserXO := *sonatyperepo.NewSamlUserXO("Status_example", "UserId_example") // SamlUserXO | The SAML user to create

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSAMLUsersAPI.CreateSecuritySamlUsers(context.Background()).SamlUserXO(samlUserXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLUsersAPI.CreateSecuritySamlUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecuritySamlUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **samlUserXO** | [**SamlUserXO**](SamlUserXO.md) | The SAML user to create | 

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


## DeleteSecuritySamlUsers

> DeleteSecuritySamlUsers(ctx, userId).Execute()

Delete a SAML user.

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
	userId := "userId_example" // string | The userid the request should apply to.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSAMLUsersAPI.DeleteSecuritySamlUsers(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLUsersAPI.DeleteSecuritySamlUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The userid the request should apply to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecuritySamlUsersRequest struct via the builder pattern


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


## GetSecuritySamlUsers

> GetSecuritySamlUsers(ctx, userId).Execute()

Retrieve a SAML user by userId.

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
	userId := "userId_example" // string | The userid the request should apply to.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSAMLUsersAPI.GetSecuritySamlUsers(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLUsersAPI.GetSecuritySamlUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The userid the request should apply to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecuritySamlUsersRequest struct via the builder pattern


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


## ListSecuritySamlUsers

> ListSecuritySamlUsers(ctx).UserId(userId).Execute()

Retrieve a list of SAML users. The response is limited to 1,000 users.

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
	userId := "userId_example" // string | An optional term to search userids for. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSAMLUsersAPI.ListSecuritySamlUsers(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLUsersAPI.ListSecuritySamlUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecuritySamlUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | An optional term to search userids for. | 

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


## UpdateSecuritySamlUsers

> UpdateSecuritySamlUsers(ctx, userId).SamlUserXO(samlUserXO).Execute()

Update a SAML user's roles and attributes.

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
	userId := "userId_example" // string | The userid the request should apply to.
	samlUserXO := *sonatyperepo.NewSamlUserXO("Status_example", "UserId_example") // SamlUserXO | The SAML user to update

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSAMLUsersAPI.UpdateSecuritySamlUsers(context.Background(), userId).SamlUserXO(samlUserXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLUsersAPI.UpdateSecuritySamlUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The userid the request should apply to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecuritySamlUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **samlUserXO** | [**SamlUserXO**](SamlUserXO.md) | The SAML user to update | 

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


# \SecurityManagementUsersAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityUsers**](SecurityManagementUsersAPI.md#CreateSecurityUsers) | **Post** /v1/security/users | Create a new user in the default source.
[**CreateSecurityUsersUserToken**](SecurityManagementUsersAPI.md#CreateSecurityUsersUserToken) | **Post** /v1/security/users/{userId}/{realm}/user-token | Create a user token for the given user.
[**DeleteSecurityUsers**](SecurityManagementUsersAPI.md#DeleteSecurityUsers) | **Delete** /v1/security/users/{userId} | Delete a user.
[**DeleteSecurityUsersUserToken**](SecurityManagementUsersAPI.md#DeleteSecurityUsersUserToken) | **Delete** /v1/security/users/{userId}/{realm}/user-token | Delete the user token for the given user.
[**DeleteSecurityUsersUserTokenReset**](SecurityManagementUsersAPI.md#DeleteSecurityUsersUserTokenReset) | **Delete** /v1/security/users/{userId}/{realm}/user-token-reset | Reset the user token for the given user.
[**GetSecurityUsersUserToken**](SecurityManagementUsersAPI.md#GetSecurityUsersUserToken) | **Get** /v1/security/users/{userId}/{realm}/user-token | Get user token metadata for the given user.
[**ListSecurityUsers**](SecurityManagementUsersAPI.md#ListSecurityUsers) | **Get** /v1/security/users | Retrieve a list of users. For SAML user sources a limit of 1000 users will be applied.
[**UpdateSecurityUsers**](SecurityManagementUsersAPI.md#UpdateSecurityUsers) | **Put** /v1/security/users/{userId} | Update an existing user.
[**UpdateSecurityUsersChangePassword**](SecurityManagementUsersAPI.md#UpdateSecurityUsersChangePassword) | **Put** /v1/security/users/{userId}/change-password | Change a user&#39;s password.



## CreateSecurityUsers

> CreateSecurityUsers(ctx).ApiCreateUser(apiCreateUser).Execute()

Create a new user in the default source.

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
	apiCreateUser := *sonatyperepo.NewApiCreateUser("EmailAddress_example", "FirstName_example", "LastName_example", "Password_example", "Status_example", "UserId_example") // ApiCreateUser | A representation of the user to create.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUsersAPI.CreateSecurityUsers(context.Background()).ApiCreateUser(apiCreateUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.CreateSecurityUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCreateUser** | [**ApiCreateUser**](ApiCreateUser.md) | A representation of the user to create. | 

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


## CreateSecurityUsersUserToken

> UserTokenXO CreateSecurityUsersUserToken(ctx, userId, realm).Execute()

Create a user token for the given user.



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
	userId := "userId_example" // string | The userId of the user to create the token for
	realm := "realm_example" // string | The realm of the user to create the token for

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementUsersAPI.CreateSecurityUsersUserToken(context.Background(), userId, realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.CreateSecurityUsersUserToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityUsersUserToken`: UserTokenXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementUsersAPI.CreateSecurityUsersUserToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The userId of the user to create the token for | 
**realm** | **string** | The realm of the user to create the token for | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityUsersUserTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserTokenXO**](UserTokenXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityUsers

> DeleteSecurityUsers(ctx, userId).Realm(realm).Execute()

Delete a user.

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
	realm := "realm_example" // string | The realm the request should apply to. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUsersAPI.DeleteSecurityUsers(context.Background(), userId).Realm(realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.DeleteSecurityUsers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSecurityUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realm** | **string** | The realm the request should apply to. | 

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


## DeleteSecurityUsersUserToken

> DeleteSecurityUsersUserToken(ctx, userId, realm).Execute()

Delete the user token for the given user.



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
	userId := "userId_example" // string | The userId of the user to delete the token for
	realm := "realm_example" // string | The realm of the user to delete the token for

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUsersAPI.DeleteSecurityUsersUserToken(context.Background(), userId, realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.DeleteSecurityUsersUserToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The userId of the user to delete the token for | 
**realm** | **string** | The realm of the user to delete the token for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityUsersUserTokenRequest struct via the builder pattern


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


## DeleteSecurityUsersUserTokenReset

> DeleteSecurityUsersUserTokenReset(ctx, userId, realm).Execute()

Reset the user token for the given user.



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
	userId := "userId_example" // string | The userId of the user to reset the token for
	realm := "realm_example" // string | The realm of the user to reset the token for

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUsersAPI.DeleteSecurityUsersUserTokenReset(context.Background(), userId, realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.DeleteSecurityUsersUserTokenReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The userId of the user to reset the token for | 
**realm** | **string** | The realm of the user to reset the token for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityUsersUserTokenResetRequest struct via the builder pattern


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


## GetSecurityUsersUserToken

> UserTokenXO GetSecurityUsersUserToken(ctx, userId, realm).Execute()

Get user token metadata for the given user.



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
	userId := "userId_example" // string | The userId of the user to get the token for
	realm := "realm_example" // string | The realm of the user to get the token for

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementUsersAPI.GetSecurityUsersUserToken(context.Background(), userId, realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.GetSecurityUsersUserToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityUsersUserToken`: UserTokenXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementUsersAPI.GetSecurityUsersUserToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The userId of the user to get the token for | 
**realm** | **string** | The realm of the user to get the token for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityUsersUserTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserTokenXO**](UserTokenXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityUsers

> ListSecurityUsers(ctx).UserId(userId).Source(source).Execute()

Retrieve a list of users. For SAML user sources a limit of 1000 users will be applied.

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
	source := "source_example" // string | An optional user source to restrict the search to. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUsersAPI.ListSecurityUsers(context.Background()).UserId(userId).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.ListSecurityUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | An optional term to search userids for. | 
 **source** | **string** | An optional user source to restrict the search to. | 

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


## UpdateSecurityUsers

> UpdateSecurityUsers(ctx, userId).ApiUser(apiUser).Execute()

Update an existing user.

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
	apiUser := *sonatyperepo.NewApiUser("EmailAddress_example", "FirstName_example", "LastName_example", "Source_example", "Status_example", "UserId_example") // ApiUser | A representation of the user to update.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUsersAPI.UpdateSecurityUsers(context.Background(), userId).ApiUser(apiUser).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.UpdateSecurityUsers``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateSecurityUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiUser** | [**ApiUser**](ApiUser.md) | A representation of the user to update. | 

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


## UpdateSecurityUsersChangePassword

> UpdateSecurityUsersChangePassword(ctx, userId).Body(body).Execute()

Change a user's password.

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
	body := "body_example" // string | The new password to use.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementUsersAPI.UpdateSecurityUsersChangePassword(context.Background(), userId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementUsersAPI.UpdateSecurityUsersChangePassword``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateSecurityUsersChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The new password to use. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


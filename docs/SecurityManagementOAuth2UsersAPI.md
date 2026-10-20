# \SecurityManagementOAuth2UsersAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityOauth2Users**](SecurityManagementOAuth2UsersAPI.md#CreateSecurityOauth2Users) | **Post** /v1/security/oauth2/users | Create a cached OAuth2 user record
[**GetSecurityOauth2Users**](SecurityManagementOAuth2UsersAPI.md#GetSecurityOauth2Users) | **Get** /v1/security/oauth2/users/{userId} | Retrieve a cached OAuth2 user by user ID
[**ListSecurityOauth2Users**](SecurityManagementOAuth2UsersAPI.md#ListSecurityOauth2Users) | **Get** /v1/security/oauth2/users | Retrieve a list of cached OAuth2 users
[**UpdateSecurityOauth2Users**](SecurityManagementOAuth2UsersAPI.md#UpdateSecurityOauth2Users) | **Put** /v1/security/oauth2/users/{userId} | Update the status of a cached OAuth2 user



## CreateSecurityOauth2Users

> CreateSecurityOauth2Users(ctx).OAuth2UserXO(oAuth2UserXO).Execute()

Create a cached OAuth2 user record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {
	oAuth2UserXO := *sonatyperepo.NewOAuth2UserXO("Status_example", "UserId_example") // OAuth2UserXO | The OAuth2 user to create

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementOAuth2UsersAPI.CreateSecurityOauth2Users(context.Background()).OAuth2UserXO(oAuth2UserXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementOAuth2UsersAPI.CreateSecurityOauth2Users``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityOauth2UsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2UserXO** | [**OAuth2UserXO**](OAuth2UserXO.md) | The OAuth2 user to create | 

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


## GetSecurityOauth2Users

> OAuth2UserXO GetSecurityOauth2Users(ctx, userId).Execute()

Retrieve a cached OAuth2 user by user ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {
	userId := "userId_example" // string | The user ID the request should apply to.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementOAuth2UsersAPI.GetSecurityOauth2Users(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementOAuth2UsersAPI.GetSecurityOauth2Users``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityOauth2Users`: OAuth2UserXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementOAuth2UsersAPI.GetSecurityOauth2Users`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID the request should apply to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityOauth2UsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OAuth2UserXO**](OAuth2UserXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityOauth2Users

> []OAuth2UserXO ListSecurityOauth2Users(ctx).UserId(userId).Execute()

Retrieve a list of cached OAuth2 users



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {
	userId := "userId_example" // string | An optional term to search user IDs for. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementOAuth2UsersAPI.ListSecurityOauth2Users(context.Background()).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementOAuth2UsersAPI.ListSecurityOauth2Users``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityOauth2Users`: []OAuth2UserXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementOAuth2UsersAPI.ListSecurityOauth2Users`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityOauth2UsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | An optional term to search user IDs for. | 

### Return type

[**[]OAuth2UserXO**](OAuth2UserXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityOauth2Users

> UpdateSecurityOauth2Users(ctx, userId).OAuth2UserXO(oAuth2UserXO).Execute()

Update the status of a cached OAuth2 user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {
	userId := "userId_example" // string | The user ID the request should apply to.
	oAuth2UserXO := *sonatyperepo.NewOAuth2UserXO("Status_example", "UserId_example") // OAuth2UserXO | The OAuth2 user to update

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementOAuth2UsersAPI.UpdateSecurityOauth2Users(context.Background(), userId).OAuth2UserXO(oAuth2UserXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementOAuth2UsersAPI.UpdateSecurityOauth2Users``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The user ID the request should apply to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityOauth2UsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oAuth2UserXO** | [**OAuth2UserXO**](OAuth2UserXO.md) | The OAuth2 user to update | 

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


# \SecurityManagementAnonymousAccessAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSecurityAnonymous**](SecurityManagementAnonymousAccessAPI.md#ListSecurityAnonymous) | **Get** /v1/security/anonymous | Get Anonymous Access settings
[**UpdateSecurityAnonymous**](SecurityManagementAnonymousAccessAPI.md#UpdateSecurityAnonymous) | **Put** /v1/security/anonymous | Update Anonymous Access settings



## ListSecurityAnonymous

> AnonymousAccessSettingsXO ListSecurityAnonymous(ctx).Execute()

Get Anonymous Access settings

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

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementAnonymousAccessAPI.ListSecurityAnonymous(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAnonymousAccessAPI.ListSecurityAnonymous``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityAnonymous`: AnonymousAccessSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAnonymousAccessAPI.ListSecurityAnonymous`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAnonymousRequest struct via the builder pattern


### Return type

[**AnonymousAccessSettingsXO**](AnonymousAccessSettingsXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityAnonymous

> AnonymousAccessSettingsXO UpdateSecurityAnonymous(ctx).AnonymousAccessSettingsXO(anonymousAccessSettingsXO).Execute()

Update Anonymous Access settings

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
	anonymousAccessSettingsXO := *sonatyperepo.NewAnonymousAccessSettingsXO("RealmName_example", "UserId_example") // AnonymousAccessSettingsXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementAnonymousAccessAPI.UpdateSecurityAnonymous(context.Background()).AnonymousAccessSettingsXO(anonymousAccessSettingsXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAnonymousAccessAPI.UpdateSecurityAnonymous``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityAnonymous`: AnonymousAccessSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAnonymousAccessAPI.UpdateSecurityAnonymous`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityAnonymousRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **anonymousAccessSettingsXO** | [**AnonymousAccessSettingsXO**](AnonymousAccessSettingsXO.md) |  | 

### Return type

[**AnonymousAccessSettingsXO**](AnonymousAccessSettingsXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


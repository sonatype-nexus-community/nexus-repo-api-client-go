# \SecurityManagementSAMLUsersV2API

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSecuritySamlUsersV2**](SecurityManagementSAMLUsersV2API.md#ListSecuritySamlUsersV2) | **Get** /v2/security/saml/users | Retrieve a paginated list of SAML users.



## ListSecuritySamlUsersV2

> ListSecuritySamlUsersV2(ctx).UserId(userId).ContinuationToken(continuationToken).Execute()

Retrieve a paginated list of SAML users.



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
	userId := "userId_example" // string | An optional term to search userids for. (optional)
	continuationToken := "continuationToken_example" // string | A token returned by a prior request. Use this to get the next page of results. Omit the parameter to retrieve the first page; supplying it with an empty value is rejected. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSAMLUsersV2API.ListSecuritySamlUsersV2(context.Background()).UserId(userId).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLUsersV2API.ListSecuritySamlUsersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecuritySamlUsersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userId** | **string** | An optional term to search userids for. | 
 **continuationToken** | **string** | A token returned by a prior request. Use this to get the next page of results. Omit the parameter to retrieve the first page; supplying it with an empty value is rejected. | 

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


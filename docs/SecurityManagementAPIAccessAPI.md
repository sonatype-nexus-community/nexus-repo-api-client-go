# \SecurityManagementAPIAccessAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInternalUiSecurityAccessCheck**](SecurityManagementAPIAccessAPI.md#CreateInternalUiSecurityAccessCheck) | **Post** /internal/ui/security/access-check | Check if a user or role has access to an API endpoint



## CreateInternalUiSecurityAccessCheck

> ApiAccessResultXo CreateInternalUiSecurityAccessCheck(ctx).ApiAccessCheckXo(apiAccessCheckXo).Execute()

Check if a user or role has access to an API endpoint

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
	apiAccessCheckXo := *sonatyperepo.NewApiAccessCheckXo("/service/rest/v1/repositories", "GET") // ApiAccessCheckXo | Access check request

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementAPIAccessAPI.CreateInternalUiSecurityAccessCheck(context.Background()).ApiAccessCheckXo(apiAccessCheckXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAPIAccessAPI.CreateInternalUiSecurityAccessCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternalUiSecurityAccessCheck`: ApiAccessResultXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAPIAccessAPI.CreateInternalUiSecurityAccessCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternalUiSecurityAccessCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiAccessCheckXo** | [**ApiAccessCheckXo**](ApiAccessCheckXo.md) | Access check request | 

### Return type

[**ApiAccessResultXo**](ApiAccessResultXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


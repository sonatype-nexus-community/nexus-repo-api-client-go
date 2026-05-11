# \SecurityManagementAPIAccessAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAccess**](SecurityManagementAPIAccessAPI.md#CheckAccess) | **Post** /internal/ui/security/access-check | Check if a user or role has access to an API endpoint



## CheckAccess

> ApiAccessResultXo CheckAccess(ctx).Body(body).Execute()

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
	body := *sonatyperepo.NewApiAccessCheckXo("/service/rest/v1/repositories", "GET") // ApiAccessCheckXo | Access check request

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementAPIAccessAPI.CheckAccess(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAPIAccessAPI.CheckAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckAccess`: ApiAccessResultXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAPIAccessAPI.CheckAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiAccessCheckXo**](ApiAccessCheckXo.md) | Access check request | 

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


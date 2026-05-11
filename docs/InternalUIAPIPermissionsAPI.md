# \InternalUIAPIPermissionsAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](InternalUIAPIPermissionsAPI.md#List) | **Get** /internal/ui/api/permissions | List REST endpoints with permission metadata for the API documentation UI



## List

> ApiPermissionsResponse List(ctx).Method(method).Path(path).Permission(permission).Tag(tag).Execute()

List REST endpoints with permission metadata for the API documentation UI

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
	method := "method_example" // string |  (optional)
	path := "path_example" // string |  (optional)
	permission := "permission_example" // string |  (optional)
	tag := "tag_example" // string |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InternalUIAPIPermissionsAPI.List(context.Background()).Method(method).Path(path).Permission(permission).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InternalUIAPIPermissionsAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ApiPermissionsResponse
	fmt.Fprintf(os.Stdout, "Response from `InternalUIAPIPermissionsAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **method** | **string** |  | 
 **path** | **string** |  | 
 **permission** | **string** |  | 
 **tag** | **string** |  | 

### Return type

[**ApiPermissionsResponse**](ApiPermissionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


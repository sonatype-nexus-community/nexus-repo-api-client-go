# \SecurityManagementAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSources**](SecurityManagementAPI.md#GetUserSources) | **Get** /v1/security/user-sources | Retrieve a list of the available user sources.



## GetUserSources

> []ApiUserSource GetUserSources(ctx).Execute()

Retrieve a list of the available user sources.

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
	resp, r, err := apiClient.SecurityManagementAPI.GetUserSources(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAPI.GetUserSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserSources`: []ApiUserSource
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAPI.GetUserSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSourcesRequest struct via the builder pattern


### Return type

[**[]ApiUserSource**](ApiUserSource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


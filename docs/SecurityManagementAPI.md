# \SecurityManagementAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSecurityUserSources**](SecurityManagementAPI.md#ListSecurityUserSources) | **Get** /v1/security/user-sources | Retrieve a list of the available user sources.



## ListSecurityUserSources

> []ApiUserSource ListSecurityUserSources(ctx).Execute()

Retrieve a list of the available user sources.

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
	resp, r, err := apiClient.SecurityManagementAPI.ListSecurityUserSources(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAPI.ListSecurityUserSources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityUserSources`: []ApiUserSource
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAPI.ListSecurityUserSources`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityUserSourcesRequest struct via the builder pattern


### Return type

[**[]ApiUserSource**](ApiUserSource.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


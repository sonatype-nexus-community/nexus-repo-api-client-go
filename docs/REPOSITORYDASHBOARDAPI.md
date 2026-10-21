# \REPOSITORYDASHBOARDAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListRepositoryDashboard**](REPOSITORYDASHBOARDAPI.md#ListRepositoryDashboard) | **Get** /v1/repository-dashboard | 
[**ListRepositoryDashboardFormats**](REPOSITORYDASHBOARDAPI.md#ListRepositoryDashboardFormats) | **Get** /v1/repository-dashboard/formats | 



## ListRepositoryDashboard

> RepositoryDashboardResponse ListRepositoryDashboard(ctx).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Search(search).Format(format).Monitoring(monitoring).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397"
)

func main() {
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 50)
	sortBy := "sortBy_example" // string |  (optional) (default to "name")
	sortOrder := "sortOrder_example" // string |  (optional) (default to "asc")
	search := "search_example" // string |  (optional)
	format := "format_example" // string |  (optional)
	monitoring := "monitoring_example" // string |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.REPOSITORYDASHBOARDAPI.ListRepositoryDashboard(context.Background()).Page(page).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Search(search).Format(format).Monitoring(monitoring).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `REPOSITORYDASHBOARDAPI.ListRepositoryDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRepositoryDashboard`: RepositoryDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `REPOSITORYDASHBOARDAPI.ListRepositoryDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRepositoryDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 50]
 **sortBy** | **string** |  | [default to &quot;name&quot;]
 **sortOrder** | **string** |  | [default to &quot;asc&quot;]
 **search** | **string** |  | 
 **format** | **string** |  | 
 **monitoring** | **string** |  | 

### Return type

[**RepositoryDashboardResponse**](RepositoryDashboardResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRepositoryDashboardFormats

> []string ListRepositoryDashboardFormats(ctx).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.REPOSITORYDASHBOARDAPI.ListRepositoryDashboardFormats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `REPOSITORYDASHBOARDAPI.ListRepositoryDashboardFormats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRepositoryDashboardFormats`: []string
	fmt.Fprintf(os.Stdout, "Response from `REPOSITORYDASHBOARDAPI.ListRepositoryDashboardFormats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRepositoryDashboardFormatsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \DataStoreAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDataStore**](DataStoreAPI.md#ListDataStore) | **Get** /beta/data-store | Get the data store
[**UpdateDataStore**](DataStoreAPI.md#UpdateDataStore) | **Put** /beta/data-store | Update the data store



## ListDataStore

> ListDataStore(ctx).Execute()

Get the data store

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

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.DataStoreAPI.ListDataStore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStoreAPI.ListDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDataStoreRequest struct via the builder pattern


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


## UpdateDataStore

> UpdateDataStore(ctx).DataStoreApiUpdateXO(dataStoreApiUpdateXO).Execute()

Update the data store

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
	dataStoreApiUpdateXO := *sonatyperepo.NewDataStoreApiUpdateXO("JdbcUrl_example", "Name_example", "Source_example", "Type_example") // DataStoreApiUpdateXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.DataStoreAPI.UpdateDataStore(context.Background()).DataStoreApiUpdateXO(dataStoreApiUpdateXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataStoreAPI.UpdateDataStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dataStoreApiUpdateXO** | [**DataStoreApiUpdateXO**](DataStoreApiUpdateXO.md) |  | 

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


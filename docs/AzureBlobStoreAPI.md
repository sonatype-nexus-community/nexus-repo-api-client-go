# \AzureBlobStoreAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAzureblobstoreTestConnection**](AzureBlobStoreAPI.md#CreateAzureblobstoreTestConnection) | **Post** /v1/azureblobstore/test-connection | Verify connection using supplied Azure Blob Store settings



## CreateAzureblobstoreTestConnection

> CreateAzureblobstoreTestConnection(ctx).AzureConnectionXO(azureConnectionXO).Execute()

Verify connection using supplied Azure Blob Store settings

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
	azureConnectionXO := *sonatyperepo.NewAzureConnectionXO() // AzureConnectionXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.AzureBlobStoreAPI.CreateAzureblobstoreTestConnection(context.Background()).AzureConnectionXO(azureConnectionXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureBlobStoreAPI.CreateAzureblobstoreTestConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureblobstoreTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureConnectionXO** | [**AzureConnectionXO**](AzureConnectionXO.md) |  | 

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


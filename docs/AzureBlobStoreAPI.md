# \AzureBlobStoreAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VerifyConnection2**](AzureBlobStoreAPI.md#VerifyConnection2) | **Post** /v1/azureblobstore/test-connection | Verify connection using supplied Azure Blob Store settings



## VerifyConnection2

> VerifyConnection2(ctx).Body(body).Execute()

Verify connection using supplied Azure Blob Store settings

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
	body := *sonatyperepo.NewAzureConnectionXO() // AzureConnectionXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.AzureBlobStoreAPI.VerifyConnection2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzureBlobStoreAPI.VerifyConnection2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyConnection2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AzureConnectionXO**](AzureConnectionXO.md) |  | 

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


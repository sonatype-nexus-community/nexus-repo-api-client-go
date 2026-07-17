# \SupportAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSupportSupportzip**](SupportAPI.md#CreateSupportSupportzip) | **Post** /v1/support/supportzip | Creates and downloads a support zip
[**CreateSupportSupportzippath**](SupportAPI.md#CreateSupportSupportzippath) | **Post** /v1/support/supportzippath | Creates a support zip and returns the path



## CreateSupportSupportzip

> CreateSupportSupportzip(ctx).SupportZipGeneratorRequest(supportZipGeneratorRequest).Execute()

Creates and downloads a support zip

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
	supportZipGeneratorRequest := *sonatyperepo.NewSupportZipGeneratorRequest() // SupportZipGeneratorRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SupportAPI.CreateSupportSupportzip(context.Background()).SupportZipGeneratorRequest(supportZipGeneratorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.CreateSupportSupportzip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportSupportzipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportZipGeneratorRequest** | [**SupportZipGeneratorRequest**](SupportZipGeneratorRequest.md) |  | 

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


## CreateSupportSupportzippath

> CreateSupportSupportzippath(ctx).SupportZipGeneratorRequest(supportZipGeneratorRequest).Execute()

Creates a support zip and returns the path

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
	supportZipGeneratorRequest := *sonatyperepo.NewSupportZipGeneratorRequest() // SupportZipGeneratorRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SupportAPI.CreateSupportSupportzippath(context.Background()).SupportZipGeneratorRequest(supportZipGeneratorRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.CreateSupportSupportzippath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSupportSupportzippathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **supportZipGeneratorRequest** | [**SupportZipGeneratorRequest**](SupportZipGeneratorRequest.md) |  | 

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


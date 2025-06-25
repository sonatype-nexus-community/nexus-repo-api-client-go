# \FormatsAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get1**](FormatsAPI.md#Get1) | **Get** /v1/formats/{format}/upload-specs | Get upload field requirements for the desired format
[**Get2**](FormatsAPI.md#Get2) | **Get** /v1/formats/upload-specs | Get upload field requirements for each supported format



## Get1

> UploadDefinitionXO Get1(ctx, format).Execute()

Get upload field requirements for the desired format

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
	format := "format_example" // string | The desired repository format

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.FormatsAPI.Get1(context.Background(), format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormatsAPI.Get1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get1`: UploadDefinitionXO
	fmt.Fprintf(os.Stdout, "Response from `FormatsAPI.Get1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** | The desired repository format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UploadDefinitionXO**](UploadDefinitionXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get2

> []UploadDefinitionXO Get2(ctx).Execute()

Get upload field requirements for each supported format

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
	resp, r, err := apiClient.FormatsAPI.Get2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FormatsAPI.Get2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get2`: []UploadDefinitionXO
	fmt.Fprintf(os.Stdout, "Response from `FormatsAPI.Get2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGet2Request struct via the builder pattern


### Return type

[**[]UploadDefinitionXO**](UploadDefinitionXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


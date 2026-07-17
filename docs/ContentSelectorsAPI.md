# \ContentSelectorsAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityContentSelectors**](ContentSelectorsAPI.md#CreateSecurityContentSelectors) | **Post** /v1/security/content-selectors | Create a new content selector
[**DeleteSecurityContentSelectors**](ContentSelectorsAPI.md#DeleteSecurityContentSelectors) | **Delete** /v1/security/content-selectors/{name} | Delete a content selector
[**GetSecurityContentSelectors**](ContentSelectorsAPI.md#GetSecurityContentSelectors) | **Get** /v1/security/content-selectors/{name} | Get a content selector by name
[**ListSecurityContentSelectors**](ContentSelectorsAPI.md#ListSecurityContentSelectors) | **Get** /v1/security/content-selectors | List content selectors
[**UpdateSecurityContentSelectors**](ContentSelectorsAPI.md#UpdateSecurityContentSelectors) | **Put** /v1/security/content-selectors/{name} | Update a content selector



## CreateSecurityContentSelectors

> CreateSecurityContentSelectors(ctx).ContentSelectorApiCreateRequest(contentSelectorApiCreateRequest).Execute()

Create a new content selector

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
	contentSelectorApiCreateRequest := *sonatyperepo.NewContentSelectorApiCreateRequest("format == "maven2" and path =^ "/org/sonatype/nexus"", "Name_example") // ContentSelectorApiCreateRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ContentSelectorsAPI.CreateSecurityContentSelectors(context.Background()).ContentSelectorApiCreateRequest(contentSelectorApiCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSelectorsAPI.CreateSecurityContentSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityContentSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentSelectorApiCreateRequest** | [**ContentSelectorApiCreateRequest**](ContentSelectorApiCreateRequest.md) |  | 

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


## DeleteSecurityContentSelectors

> DeleteSecurityContentSelectors(ctx, name).Execute()

Delete a content selector

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
	name := "name_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ContentSelectorsAPI.DeleteSecurityContentSelectors(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSelectorsAPI.DeleteSecurityContentSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityContentSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetSecurityContentSelectors

> ContentSelectorApiResponse GetSecurityContentSelectors(ctx, name).Execute()

Get a content selector by name

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
	name := "name_example" // string | The content selector name

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentSelectorsAPI.GetSecurityContentSelectors(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSelectorsAPI.GetSecurityContentSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityContentSelectors`: ContentSelectorApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSelectorsAPI.GetSecurityContentSelectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The content selector name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityContentSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentSelectorApiResponse**](ContentSelectorApiResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityContentSelectors

> []ContentSelectorApiResponse ListSecurityContentSelectors(ctx).Execute()

List content selectors

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
	resp, r, err := apiClient.ContentSelectorsAPI.ListSecurityContentSelectors(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSelectorsAPI.ListSecurityContentSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityContentSelectors`: []ContentSelectorApiResponse
	fmt.Fprintf(os.Stdout, "Response from `ContentSelectorsAPI.ListSecurityContentSelectors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityContentSelectorsRequest struct via the builder pattern


### Return type

[**[]ContentSelectorApiResponse**](ContentSelectorApiResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityContentSelectors

> UpdateSecurityContentSelectors(ctx, name).ContentSelectorApiUpdateRequest(contentSelectorApiUpdateRequest).Execute()

Update a content selector

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
	name := "name_example" // string | The content selector name
	contentSelectorApiUpdateRequest := *sonatyperepo.NewContentSelectorApiUpdateRequest("format == "maven2" and path =^ "/org/sonatype/nexus"") // ContentSelectorApiUpdateRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ContentSelectorsAPI.UpdateSecurityContentSelectors(context.Background(), name).ContentSelectorApiUpdateRequest(contentSelectorApiUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentSelectorsAPI.UpdateSecurityContentSelectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The content selector name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityContentSelectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentSelectorApiUpdateRequest** | [**ContentSelectorApiUpdateRequest**](ContentSelectorApiUpdateRequest.md) |  | 

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


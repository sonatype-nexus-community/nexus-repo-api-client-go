# \ManageSonatypeHTTPSystemSettingsAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteHttp**](ManageSonatypeHTTPSystemSettingsAPI.md#DeleteHttp) | **Delete** /v1/http | Reset HTTP System Settings
[**ListHttp**](ManageSonatypeHTTPSystemSettingsAPI.md#ListHttp) | **Get** /v1/http | Get HTTP system settings
[**UpdateHttp**](ManageSonatypeHTTPSystemSettingsAPI.md#UpdateHttp) | **Put** /v1/http | Update HTTP system settings



## DeleteHttp

> DeleteHttp(ctx).Execute()

Reset HTTP System Settings

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
	r, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.DeleteHttp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeHTTPSystemSettingsAPI.DeleteHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHttpRequest struct via the builder pattern


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


## ListHttp

> HttpSettingsXo ListHttp(ctx).Execute()

Get HTTP system settings

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
	resp, r, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.ListHttp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeHTTPSystemSettingsAPI.ListHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHttp`: HttpSettingsXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeHTTPSystemSettingsAPI.ListHttp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHttpRequest struct via the builder pattern


### Return type

[**HttpSettingsXo**](HttpSettingsXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHttp

> UpdateHttp(ctx).HttpSettingsXo(httpSettingsXo).Execute()

Update HTTP system settings

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
	httpSettingsXo := *sonatyperepo.NewHttpSettingsXo("TODO", "TODO", int32(123), int32(123), "UserAgent_example") // HttpSettingsXo | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.UpdateHttp(context.Background()).HttpSettingsXo(httpSettingsXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeHTTPSystemSettingsAPI.UpdateHttp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHttpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **httpSettingsXo** | [**HttpSettingsXo**](HttpSettingsXo.md) |  | 

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


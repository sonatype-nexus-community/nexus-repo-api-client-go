# \ManageSonatypeHTTPSystemSettingsAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHttpSettings**](ManageSonatypeHTTPSystemSettingsAPI.md#GetHttpSettings) | **Get** /v1/http | Get HTTP system settings
[**ResetHttpSettings**](ManageSonatypeHTTPSystemSettingsAPI.md#ResetHttpSettings) | **Delete** /v1/http | Reset HTTP System Settings
[**UpdateHttpSettings**](ManageSonatypeHTTPSystemSettingsAPI.md#UpdateHttpSettings) | **Put** /v1/http | Update HTTP system settings



## GetHttpSettings

> HttpSettingsXo GetHttpSettings(ctx).Execute()

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
	resp, r, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.GetHttpSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeHTTPSystemSettingsAPI.GetHttpSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHttpSettings`: HttpSettingsXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeHTTPSystemSettingsAPI.GetHttpSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetHttpSettingsRequest struct via the builder pattern


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


## ResetHttpSettings

> ResetHttpSettings(ctx).Execute()

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
	r, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.ResetHttpSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeHTTPSystemSettingsAPI.ResetHttpSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiResetHttpSettingsRequest struct via the builder pattern


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


## UpdateHttpSettings

> UpdateHttpSettings(ctx).Body(body).Execute()

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
	body := *sonatyperepo.NewHttpSettingsXo(*sonatyperepo.NewProxySettingsXo(*sonatyperepo.NewAuthSettingsXo(false, "NtlmDomain_example", "NtlmHost_example", "Password_example", "Username_example"), false, "Host_example", "Port_example"), *sonatyperepo.NewProxySettingsXo(*sonatyperepo.NewAuthSettingsXo(false, "NtlmDomain_example", "NtlmHost_example", "Password_example", "Username_example"), false, "Host_example", "Port_example"), int32(123), int32(123), "UserAgent_example") // HttpSettingsXo |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeHTTPSystemSettingsAPI.UpdateHttpSettings(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeHTTPSystemSettingsAPI.UpdateHttpSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHttpSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HttpSettingsXo**](HttpSettingsXo.md) |  | 

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


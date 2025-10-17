# \InstanceConfigurationAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyConfiguration**](InstanceConfigurationAPI.md#ApplyConfiguration) | **Put** /v1/configuration | Apply instance configuration
[**ExportConfiguration**](InstanceConfigurationAPI.md#ExportConfiguration) | **Get** /v1/configuration | Export instance configuration



## ApplyConfiguration

> ConfigurationApplyResponseXO ApplyConfiguration(ctx).Body(body).Execute()

Apply instance configuration

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
	body := *sonatyperepo.NewInstanceConfigurationXO() // InstanceConfigurationXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.ApplyConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.ApplyConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyConfiguration`: ConfigurationApplyResponseXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.ApplyConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**InstanceConfigurationXO**](InstanceConfigurationXO.md) |  | 

### Return type

[**ConfigurationApplyResponseXO**](ConfigurationApplyResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportConfiguration

> InstanceConfigurationXO ExportConfiguration(ctx).Execute()

Export instance configuration

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
	resp, r, err := apiClient.InstanceConfigurationAPI.ExportConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.ExportConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportConfiguration`: InstanceConfigurationXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.ExportConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExportConfigurationRequest struct via the builder pattern


### Return type

[**InstanceConfigurationXO**](InstanceConfigurationXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


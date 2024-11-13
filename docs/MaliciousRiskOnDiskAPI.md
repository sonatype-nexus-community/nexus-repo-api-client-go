# \MaliciousRiskOnDiskAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnabledRegistries**](MaliciousRiskOnDiskAPI.md#GetEnabledRegistries) | **Get** /v1/malicious-risk/enabledRegistries | Get RHC Enabled registries for malicious risk scanning.
[**GetMaliciousRiskOnDiskCount**](MaliciousRiskOnDiskAPI.md#GetMaliciousRiskOnDiskCount) | **Get** /v1/malicious-risk/risk-on-disk | Get Malicious Risk On Disk Count



## GetEnabledRegistries

> GetEnabledRegistries(ctx).Execute()

Get RHC Enabled registries for malicious risk scanning.

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

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.MaliciousRiskOnDiskAPI.GetEnabledRegistries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaliciousRiskOnDiskAPI.GetEnabledRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnabledRegistriesRequest struct via the builder pattern


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


## GetMaliciousRiskOnDiskCount

> GetMaliciousRiskOnDiskCount(ctx).Execute()

Get Malicious Risk On Disk Count

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

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.MaliciousRiskOnDiskAPI.GetMaliciousRiskOnDiskCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaliciousRiskOnDiskAPI.GetMaliciousRiskOnDiskCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaliciousRiskOnDiskCountRequest struct via the builder pattern


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


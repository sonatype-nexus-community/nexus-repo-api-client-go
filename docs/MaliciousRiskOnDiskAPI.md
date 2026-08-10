# \MaliciousRiskOnDiskAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMaliciousRiskEnabledRegistries**](MaliciousRiskOnDiskAPI.md#ListMaliciousRiskEnabledRegistries) | **Get** /v1/malicious-risk/enabledRegistries | Get RHC Enabled registries for malicious risk scanning.
[**ListMaliciousRiskRiskOnDisk**](MaliciousRiskOnDiskAPI.md#ListMaliciousRiskRiskOnDisk) | **Get** /v1/malicious-risk/risk-on-disk | Get Malicious Risk On Disk Count



## ListMaliciousRiskEnabledRegistries

> map[string]interface{} ListMaliciousRiskEnabledRegistries(ctx).Execute()

Get RHC Enabled registries for malicious risk scanning.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.MaliciousRiskOnDiskAPI.ListMaliciousRiskEnabledRegistries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaliciousRiskOnDiskAPI.ListMaliciousRiskEnabledRegistries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMaliciousRiskEnabledRegistries`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `MaliciousRiskOnDiskAPI.ListMaliciousRiskEnabledRegistries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMaliciousRiskEnabledRegistriesRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMaliciousRiskRiskOnDisk

> MaliciousRiskOnDiskCountResponse ListMaliciousRiskRiskOnDisk(ctx).Execute()

Get Malicious Risk On Disk Count

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.MaliciousRiskOnDiskAPI.ListMaliciousRiskRiskOnDisk(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MaliciousRiskOnDiskAPI.ListMaliciousRiskRiskOnDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMaliciousRiskRiskOnDisk`: MaliciousRiskOnDiskCountResponse
	fmt.Fprintf(os.Stdout, "Response from `MaliciousRiskOnDiskAPI.ListMaliciousRiskRiskOnDisk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMaliciousRiskRiskOnDiskRequest struct via the builder pattern


### Return type

[**MaliciousRiskOnDiskCountResponse**](MaliciousRiskOnDiskCountResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


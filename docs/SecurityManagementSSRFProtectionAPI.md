# \SecurityManagementSSRFProtectionAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSecuritySsrfProtection**](SecurityManagementSSRFProtectionAPI.md#ListSecuritySsrfProtection) | **Get** /v1/security/ssrf-protection | Get SSRF protection settings
[**UpdateSecuritySsrfProtection**](SecurityManagementSSRFProtectionAPI.md#UpdateSecuritySsrfProtection) | **Put** /v1/security/ssrf-protection | Update SSRF protection settings



## ListSecuritySsrfProtection

> ListSecuritySsrfProtection(ctx).Execute()

Get SSRF protection settings

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
	r, err := apiClient.SecurityManagementSSRFProtectionAPI.ListSecuritySsrfProtection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSSRFProtectionAPI.ListSecuritySsrfProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecuritySsrfProtectionRequest struct via the builder pattern


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


## UpdateSecuritySsrfProtection

> UpdateSecuritySsrfProtection(ctx).SsrfProtectionConfigurationXO(ssrfProtectionConfigurationXO).Execute()

Update SSRF protection settings

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
	ssrfProtectionConfigurationXO := *sonatyperepo.NewSsrfProtectionConfigurationXO(true) // SsrfProtectionConfigurationXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSSRFProtectionAPI.UpdateSecuritySsrfProtection(context.Background()).SsrfProtectionConfigurationXO(ssrfProtectionConfigurationXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSSRFProtectionAPI.UpdateSecuritySsrfProtection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecuritySsrfProtectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ssrfProtectionConfigurationXO** | [**SsrfProtectionConfigurationXO**](SsrfProtectionConfigurationXO.md) |  | 

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


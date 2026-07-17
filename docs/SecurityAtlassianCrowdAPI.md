# \SecurityAtlassianCrowdAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityAtlassianCrowdClearCache**](SecurityAtlassianCrowdAPI.md#CreateSecurityAtlassianCrowdClearCache) | **Post** /v1/security/atlassian-crowd/clear-cache | Clear Atlassian Crowd cache
[**CreateSecurityAtlassianCrowdVerifyConnection**](SecurityAtlassianCrowdAPI.md#CreateSecurityAtlassianCrowdVerifyConnection) | **Post** /v1/security/atlassian-crowd/verify-connection | Verify connection using supplied Atlassian Crowd settings
[**ListSecurityAtlassianCrowd**](SecurityAtlassianCrowdAPI.md#ListSecurityAtlassianCrowd) | **Get** /v1/security/atlassian-crowd | Retrieve Atlassian Crowd settings configured in Nexus Repository Manager
[**UpdateSecurityAtlassianCrowd**](SecurityAtlassianCrowdAPI.md#UpdateSecurityAtlassianCrowd) | **Put** /v1/security/atlassian-crowd | Update Atlassian Crowd settings configured in Nexus Repository Manager



## CreateSecurityAtlassianCrowdClearCache

> CreateSecurityAtlassianCrowdClearCache(ctx).Execute()

Clear Atlassian Crowd cache

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
	r, err := apiClient.SecurityAtlassianCrowdAPI.CreateSecurityAtlassianCrowdClearCache(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAtlassianCrowdAPI.CreateSecurityAtlassianCrowdClearCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityAtlassianCrowdClearCacheRequest struct via the builder pattern


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


## CreateSecurityAtlassianCrowdVerifyConnection

> CreateSecurityAtlassianCrowdVerifyConnection(ctx).CrowdApiXO(crowdApiXO).Execute()

Verify connection using supplied Atlassian Crowd settings

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
	crowdApiXO := *sonatyperepo.NewCrowdApiXO("ApplicationName_example", "ApplicationPassword_example", false, false, "Url_example") // CrowdApiXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityAtlassianCrowdAPI.CreateSecurityAtlassianCrowdVerifyConnection(context.Background()).CrowdApiXO(crowdApiXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAtlassianCrowdAPI.CreateSecurityAtlassianCrowdVerifyConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityAtlassianCrowdVerifyConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crowdApiXO** | [**CrowdApiXO**](CrowdApiXO.md) |  | 

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


## ListSecurityAtlassianCrowd

> ListSecurityAtlassianCrowd(ctx).Execute()

Retrieve Atlassian Crowd settings configured in Nexus Repository Manager

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
	r, err := apiClient.SecurityAtlassianCrowdAPI.ListSecurityAtlassianCrowd(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAtlassianCrowdAPI.ListSecurityAtlassianCrowd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAtlassianCrowdRequest struct via the builder pattern


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


## UpdateSecurityAtlassianCrowd

> UpdateSecurityAtlassianCrowd(ctx).CrowdApiXO(crowdApiXO).Execute()

Update Atlassian Crowd settings configured in Nexus Repository Manager

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
	crowdApiXO := *sonatyperepo.NewCrowdApiXO("ApplicationName_example", "ApplicationPassword_example", false, false, "Url_example") // CrowdApiXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityAtlassianCrowdAPI.UpdateSecurityAtlassianCrowd(context.Background()).CrowdApiXO(crowdApiXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAtlassianCrowdAPI.UpdateSecurityAtlassianCrowd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityAtlassianCrowdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **crowdApiXO** | [**CrowdApiXO**](CrowdApiXO.md) |  | 

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


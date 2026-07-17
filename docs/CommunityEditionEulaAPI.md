# \CommunityEditionEulaAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemEula**](CommunityEditionEulaAPI.md#CreateSystemEula) | **Post** /v1/system/eula | Set the Community Eula status.
[**ListSystemEula**](CommunityEditionEulaAPI.md#ListSystemEula) | **Get** /v1/system/eula | Get the current Community Eula status.



## CreateSystemEula

> CreateSystemEula(ctx).EulaStatus(eulaStatus).Execute()

Set the Community Eula status.

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
	eulaStatus := *sonatyperepo.NewEulaStatus() // EulaStatus |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.CommunityEditionEulaAPI.CreateSystemEula(context.Background()).EulaStatus(eulaStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityEditionEulaAPI.CreateSystemEula``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemEulaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eulaStatus** | [**EulaStatus**](EulaStatus.md) |  | 

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


## ListSystemEula

> ListSystemEula(ctx).Execute()

Get the current Community Eula status.

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
	r, err := apiClient.CommunityEditionEulaAPI.ListSystemEula(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommunityEditionEulaAPI.ListSystemEula``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemEulaRequest struct via the builder pattern


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


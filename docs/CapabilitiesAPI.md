# \CapabilitiesAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create4**](CapabilitiesAPI.md#Create4) | **Post** /v1/capabilities | Create a capability
[**Delete5**](CapabilitiesAPI.md#Delete5) | **Delete** /v1/capabilities/{capabilityId} | Delete a capability
[**GetTypes**](CapabilitiesAPI.md#GetTypes) | **Get** /v1/capabilities/types | List all capability types available and exposed in the system
[**List2**](CapabilitiesAPI.md#List2) | **Get** /v1/capabilities | List the active capabilities
[**Update3**](CapabilitiesAPI.md#Update3) | **Put** /v1/capabilities/{capabilityId} | Update a capability



## Create4

> CapabilityDTO Create4(ctx).Body(body).Execute()

Create a capability

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
	body := *sonatyperepo.NewCapabilityDTO() // CapabilityDTO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.CapabilitiesAPI.Create4(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.Create4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create4`: CapabilityDTO
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.Create4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreate4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CapabilityDTO**](CapabilityDTO.md) |  | 

### Return type

[**CapabilityDTO**](CapabilityDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete5

> Delete5(ctx, capabilityId).Execute()

Delete a capability

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
	capabilityId := "capabilityId_example" // string | capabilityId

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.CapabilitiesAPI.Delete5(context.Background(), capabilityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.Delete5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**capabilityId** | **string** | capabilityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete5Request struct via the builder pattern


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


## GetTypes

> []CapabilityType GetTypes(ctx).Execute()

List all capability types available and exposed in the system

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
	resp, r, err := apiClient.CapabilitiesAPI.GetTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.GetTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTypes`: []CapabilityType
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.GetTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTypesRequest struct via the builder pattern


### Return type

[**[]CapabilityType**](CapabilityType.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List2

> []CapabilityDTO List2(ctx).Execute()

List the active capabilities

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
	resp, r, err := apiClient.CapabilitiesAPI.List2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.List2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List2`: []CapabilityDTO
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.List2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiList2Request struct via the builder pattern


### Return type

[**[]CapabilityDTO**](CapabilityDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update3

> Update3(ctx, capabilityId).Body(body).Execute()

Update a capability

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
	capabilityId := "capabilityId_example" // string | capabilityId
	body := *sonatyperepo.NewCapabilityDTO() // CapabilityDTO | capability (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.CapabilitiesAPI.Update3(context.Background(), capabilityId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.Update3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**capabilityId** | **string** | capabilityId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdate3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CapabilityDTO**](CapabilityDTO.md) | capability | 

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


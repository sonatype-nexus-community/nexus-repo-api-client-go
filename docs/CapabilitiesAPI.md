# \CapabilitiesAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCapabilities**](CapabilitiesAPI.md#CreateCapabilities) | **Post** /v1/capabilities | Create a capability
[**DeleteCapabilities**](CapabilitiesAPI.md#DeleteCapabilities) | **Delete** /v1/capabilities/{capabilityId} | Delete a capability
[**ListCapabilities**](CapabilitiesAPI.md#ListCapabilities) | **Get** /v1/capabilities | List the active capabilities
[**ListCapabilitiesTypes**](CapabilitiesAPI.md#ListCapabilitiesTypes) | **Get** /v1/capabilities/types | List all capability types available and exposed in the system
[**UpdateCapabilities**](CapabilitiesAPI.md#UpdateCapabilities) | **Put** /v1/capabilities/{capabilityId} | Update a capability



## CreateCapabilities

> CapabilityDTO CreateCapabilities(ctx).CapabilityDTO(capabilityDTO).Execute()

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
	capabilityDTO := *sonatyperepo.NewCapabilityDTO() // CapabilityDTO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.CapabilitiesAPI.CreateCapabilities(context.Background()).CapabilityDTO(capabilityDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.CreateCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCapabilities`: CapabilityDTO
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.CreateCapabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **capabilityDTO** | [**CapabilityDTO**](CapabilityDTO.md) |  | 

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


## DeleteCapabilities

> DeleteCapabilities(ctx, capabilityId).Execute()

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
	r, err := apiClient.CapabilitiesAPI.DeleteCapabilities(context.Background(), capabilityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.DeleteCapabilities``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCapabilities

> []CapabilityDTO ListCapabilities(ctx).Execute()

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
	resp, r, err := apiClient.CapabilitiesAPI.ListCapabilities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.ListCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCapabilities`: []CapabilityDTO
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.ListCapabilities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCapabilitiesRequest struct via the builder pattern


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


## ListCapabilitiesTypes

> []CapabilityTypeDTO ListCapabilitiesTypes(ctx).Execute()

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
	resp, r, err := apiClient.CapabilitiesAPI.ListCapabilitiesTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.ListCapabilitiesTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCapabilitiesTypes`: []CapabilityTypeDTO
	fmt.Fprintf(os.Stdout, "Response from `CapabilitiesAPI.ListCapabilitiesTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCapabilitiesTypesRequest struct via the builder pattern


### Return type

[**[]CapabilityTypeDTO**](CapabilityTypeDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCapabilities

> UpdateCapabilities(ctx, capabilityId).CapabilityDTO(capabilityDTO).Execute()

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
	capabilityDTO := *sonatyperepo.NewCapabilityDTO() // CapabilityDTO | capability (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.CapabilitiesAPI.UpdateCapabilities(context.Background(), capabilityId).CapabilityDTO(capabilityDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CapabilitiesAPI.UpdateCapabilities``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **capabilityDTO** | [**CapabilityDTO**](CapabilityDTO.md) | capability | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


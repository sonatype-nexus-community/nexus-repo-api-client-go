# \MandatoryTelemetryAcknowledgementAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSystemMandatoryTelemetryAcknowledgement**](MandatoryTelemetryAcknowledgementAPI.md#CreateSystemMandatoryTelemetryAcknowledgement) | **Post** /v1/system/mandatory-telemetry-acknowledgement | Record mandatory telemetry acknowledgement
[**ListSystemMandatoryTelemetryAcknowledgement**](MandatoryTelemetryAcknowledgementAPI.md#ListSystemMandatoryTelemetryAcknowledgement) | **Get** /v1/system/mandatory-telemetry-acknowledgement | Get the mandatory telemetry acknowledgement record



## CreateSystemMandatoryTelemetryAcknowledgement

> CreateSystemMandatoryTelemetryAcknowledgement(ctx).MandatoryTelemetryAcknowledgementStatus(mandatoryTelemetryAcknowledgementStatus).Execute()

Record mandatory telemetry acknowledgement



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397"
)

func main() {
	mandatoryTelemetryAcknowledgementStatus := *sonatyperepo.NewMandatoryTelemetryAcknowledgementStatus() // MandatoryTelemetryAcknowledgementStatus |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.MandatoryTelemetryAcknowledgementAPI.CreateSystemMandatoryTelemetryAcknowledgement(context.Background()).MandatoryTelemetryAcknowledgementStatus(mandatoryTelemetryAcknowledgementStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatoryTelemetryAcknowledgementAPI.CreateSystemMandatoryTelemetryAcknowledgement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSystemMandatoryTelemetryAcknowledgementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mandatoryTelemetryAcknowledgementStatus** | [**MandatoryTelemetryAcknowledgementStatus**](MandatoryTelemetryAcknowledgementStatus.md) |  | 

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


## ListSystemMandatoryTelemetryAcknowledgement

> MandatoryTelemetryAcknowledgementStatus ListSystemMandatoryTelemetryAcknowledgement(ctx).Execute()

Get the mandatory telemetry acknowledgement record

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v397"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.MandatoryTelemetryAcknowledgementAPI.ListSystemMandatoryTelemetryAcknowledgement(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MandatoryTelemetryAcknowledgementAPI.ListSystemMandatoryTelemetryAcknowledgement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSystemMandatoryTelemetryAcknowledgement`: MandatoryTelemetryAcknowledgementStatus
	fmt.Fprintf(os.Stdout, "Response from `MandatoryTelemetryAcknowledgementAPI.ListSystemMandatoryTelemetryAcknowledgement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemMandatoryTelemetryAcknowledgementRequest struct via the builder pattern


### Return type

[**MandatoryTelemetryAcknowledgementStatus**](MandatoryTelemetryAcknowledgementStatus.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


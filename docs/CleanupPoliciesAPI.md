# \CleanupPoliciesAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCleanupRun**](CleanupPoliciesAPI.md#CreateCleanupRun) | **Post** /v1/cleanup/run | Run cleanup on a repository (dry run or async execution)
[**CreateInternalCleanupPolicies**](CleanupPoliciesAPI.md#CreateInternalCleanupPolicies) | **Post** /internal/cleanup-policies | Create a cleanup policy
[**DeleteInternalCleanupPolicies**](CleanupPoliciesAPI.md#DeleteInternalCleanupPolicies) | **Delete** /internal/cleanup-policies/{name} | 
[**GetInternalCleanupPolicies**](CleanupPoliciesAPI.md#GetInternalCleanupPolicies) | **Get** /internal/cleanup-policies/{name} | 
[**ListInternalCleanupPolicies**](CleanupPoliciesAPI.md#ListInternalCleanupPolicies) | **Get** /internal/cleanup-policies | List cleanup policies
[**UpdateInternalCleanupPolicies**](CleanupPoliciesAPI.md#UpdateInternalCleanupPolicies) | **Put** /internal/cleanup-policies/{name} | 



## CreateCleanupRun

> CleanupExecutionStatusXO CreateCleanupRun(ctx).CleanupExecutionRequestXO(cleanupExecutionRequestXO).Execute()

Run cleanup on a repository (dry run or async execution)



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
	cleanupExecutionRequestXO := *sonatyperepo.NewCleanupExecutionRequestXO("Repository_example") // CleanupExecutionRequestXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.CleanupPoliciesAPI.CreateCleanupRun(context.Background()).CleanupExecutionRequestXO(cleanupExecutionRequestXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.CreateCleanupRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCleanupRun`: CleanupExecutionStatusXO
	fmt.Fprintf(os.Stdout, "Response from `CleanupPoliciesAPI.CreateCleanupRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCleanupRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleanupExecutionRequestXO** | [**CleanupExecutionRequestXO**](CleanupExecutionRequestXO.md) |  | 

### Return type

[**CleanupExecutionStatusXO**](CleanupExecutionStatusXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInternalCleanupPolicies

> CleanupPolicyXO CreateInternalCleanupPolicies(ctx).CleanupPolicyXO(cleanupPolicyXO).Execute()

Create a cleanup policy



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
	cleanupPolicyXO := *sonatyperepo.NewCleanupPolicyXO("Format_example", "Name_example") // CleanupPolicyXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.CleanupPoliciesAPI.CreateInternalCleanupPolicies(context.Background()).CleanupPolicyXO(cleanupPolicyXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.CreateInternalCleanupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInternalCleanupPolicies`: CleanupPolicyXO
	fmt.Fprintf(os.Stdout, "Response from `CleanupPoliciesAPI.CreateInternalCleanupPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInternalCleanupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cleanupPolicyXO** | [**CleanupPolicyXO**](CleanupPolicyXO.md) |  | 

### Return type

[**CleanupPolicyXO**](CleanupPolicyXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInternalCleanupPolicies

> DeleteInternalCleanupPolicies(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.CleanupPoliciesAPI.DeleteInternalCleanupPolicies(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.DeleteInternalCleanupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInternalCleanupPoliciesRequest struct via the builder pattern


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


## GetInternalCleanupPolicies

> CleanupPolicyXO GetInternalCleanupPolicies(ctx, name).Execute()



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
	name := "name_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.CleanupPoliciesAPI.GetInternalCleanupPolicies(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.GetInternalCleanupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInternalCleanupPolicies`: CleanupPolicyXO
	fmt.Fprintf(os.Stdout, "Response from `CleanupPoliciesAPI.GetInternalCleanupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInternalCleanupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CleanupPolicyXO**](CleanupPolicyXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInternalCleanupPolicies

> []CleanupPolicyXO ListInternalCleanupPolicies(ctx).Format(format).Execute()

List cleanup policies



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
	format := "format_example" // string |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.CleanupPoliciesAPI.ListInternalCleanupPolicies(context.Background()).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.ListInternalCleanupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInternalCleanupPolicies`: []CleanupPolicyXO
	fmt.Fprintf(os.Stdout, "Response from `CleanupPoliciesAPI.ListInternalCleanupPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInternalCleanupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 

### Return type

[**[]CleanupPolicyXO**](CleanupPolicyXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInternalCleanupPolicies

> CleanupPolicyXO UpdateInternalCleanupPolicies(ctx, name).CleanupPolicyXO(cleanupPolicyXO).Execute()



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
	name := "name_example" // string | 
	cleanupPolicyXO := *sonatyperepo.NewCleanupPolicyXO("Format_example", "Name_example") // CleanupPolicyXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.CleanupPoliciesAPI.UpdateInternalCleanupPolicies(context.Background(), name).CleanupPolicyXO(cleanupPolicyXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CleanupPoliciesAPI.UpdateInternalCleanupPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInternalCleanupPolicies`: CleanupPolicyXO
	fmt.Fprintf(os.Stdout, "Response from `CleanupPoliciesAPI.UpdateInternalCleanupPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInternalCleanupPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cleanupPolicyXO** | [**CleanupPolicyXO**](CleanupPolicyXO.md) |  | 

### Return type

[**CleanupPolicyXO**](CleanupPolicyXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


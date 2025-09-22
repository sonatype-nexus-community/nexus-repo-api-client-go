# \ReconcilePlanAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReconciliationPlans**](ReconcilePlanAPI.md#CreateReconciliationPlans) | **Post** /v1/plan | Create reconciliation plans with selected parameters
[**DeleteAllPlans**](ReconcilePlanAPI.md#DeleteAllPlans) | **Delete** /v1/plan | Delete all non executed reconciliation plans
[**DeletePlan**](ReconcilePlanAPI.md#DeletePlan) | **Delete** /v1/plan/{planId} | Delete a reconciliation plan based on its Id
[**ExecuteReconcileTaskById**](ReconcilePlanAPI.md#ExecuteReconcileTaskById) | **Put** /v1/plan/{planId} | Execute a reconciliation plan based on its Id
[**ExecuteReconcileTasks**](ReconcilePlanAPI.md#ExecuteReconcileTasks) | **Put** /v1/plan | Execute all non executed reconciliation plans
[**GetAvailablePlans**](ReconcilePlanAPI.md#GetAvailablePlans) | **Get** /v1/plan | Get list of currently available plans
[**GetPlanDetails**](ReconcilePlanAPI.md#GetPlanDetails) | **Get** /v1/plan/details | Get reconciliation plan details
[**GetSinglePlanWithDetails**](ReconcilePlanAPI.md#GetSinglePlanWithDetails) | **Get** /v1/plan/{planId} | Get single reconciliation plan with details



## CreateReconciliationPlans

> CreateReconciliationPlans(ctx).Repositories(repositories).BlobStores(blobStores).StartDate(startDate).EndDate(endDate).SinceDays(sinceDays).SinceHours(sinceHours).SinceMinutes(sinceMinutes).Execute()

Create reconciliation plans with selected parameters

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
	repositories := []string{"Inner_example"} // []string | repository(ies) which should be processed with high priority (optional)
	blobStores := []string{"Inner_example"} // []string |  (optional)
	startDate := "startDate_example" // string |  (optional)
	endDate := "endDate_example" // string |  (optional)
	sinceDays := int32(56) // int32 |  (optional)
	sinceHours := int32(56) // int32 |  (optional)
	sinceMinutes := int32(56) // int32 |  (optional) (default to 30)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ReconcilePlanAPI.CreateReconciliationPlans(context.Background()).Repositories(repositories).BlobStores(blobStores).StartDate(startDate).EndDate(endDate).SinceDays(sinceDays).SinceHours(sinceHours).SinceMinutes(sinceMinutes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.CreateReconciliationPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReconciliationPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repositories** | **[]string** | repository(ies) which should be processed with high priority | 
 **blobStores** | **[]string** |  | 
 **startDate** | **string** |  | 
 **endDate** | **string** |  | 
 **sinceDays** | **int32** |  | 
 **sinceHours** | **int32** |  | 
 **sinceMinutes** | **int32** |  | [default to 30]

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


## DeleteAllPlans

> DeleteAllPlans(ctx).Execute()

Delete all non executed reconciliation plans

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
	r, err := apiClient.ReconcilePlanAPI.DeleteAllPlans(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.DeleteAllPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllPlansRequest struct via the builder pattern


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


## DeletePlan

> DeletePlan(ctx, planId).Execute()

Delete a reconciliation plan based on its Id

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
	planId := int32(56) // int32 | Id of the plan to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ReconcilePlanAPI.DeletePlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.DeletePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int32** | Id of the plan to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


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


## ExecuteReconcileTaskById

> ExecuteReconcileTaskById(ctx, planId).Execute()

Execute a reconciliation plan based on its Id

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
	planId := int32(56) // int32 | Id of the plan to execute

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ReconcilePlanAPI.ExecuteReconcileTaskById(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.ExecuteReconcileTaskById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int32** | Id of the plan to execute | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteReconcileTaskByIdRequest struct via the builder pattern


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


## ExecuteReconcileTasks

> ExecuteReconcileTasks(ctx).Execute()

Execute all non executed reconciliation plans

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
	r, err := apiClient.ReconcilePlanAPI.ExecuteReconcileTasks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.ExecuteReconcileTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteReconcileTasksRequest struct via the builder pattern


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


## GetAvailablePlans

> GetAvailablePlans(ctx).ContinuationToken(continuationToken).Execute()

Get list of currently available plans

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
	continuationToken := "continuationToken_example" // string | A token returned by a prior request. If present, the next page of results are returned (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ReconcilePlanAPI.GetAvailablePlans(context.Background()).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.GetAvailablePlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **continuationToken** | **string** | A token returned by a prior request. If present, the next page of results are returned | 

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


## GetPlanDetails

> GetPlanDetails(ctx).PlanId(planId).State(state).Repository(repository).ContinuationToken(continuationToken).Execute()

Get reconciliation plan details

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
	planId := int32(56) // int32 | Id of the plan (optional)
	state := "state_example" // string |  (optional)
	repository := "repository_example" // string |  (optional)
	continuationToken := "continuationToken_example" // string | A token returned by a prior request. If present, the next page of results are returned (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ReconcilePlanAPI.GetPlanDetails(context.Background()).PlanId(planId).State(state).Repository(repository).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.GetPlanDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planId** | **int32** | Id of the plan | 
 **state** | **string** |  | 
 **repository** | **string** |  | 
 **continuationToken** | **string** | A token returned by a prior request. If present, the next page of results are returned | 

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


## GetSinglePlanWithDetails

> GetSinglePlanWithDetails(ctx, planId).Repository(repository).ContinuationToken(continuationToken).Execute()

Get single reconciliation plan with details

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
	planId := int32(56) // int32 | Id of the plan
	repository := "repository_example" // string |  (optional)
	continuationToken := "continuationToken_example" // string | A token returned by a prior request. If present, the next page of results are returned (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ReconcilePlanAPI.GetSinglePlanWithDetails(context.Background(), planId).Repository(repository).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReconcilePlanAPI.GetSinglePlanWithDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int32** | Id of the plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSinglePlanWithDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repository** | **string** |  | 
 **continuationToken** | **string** | A token returned by a prior request. If present, the next page of results are returned | 

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


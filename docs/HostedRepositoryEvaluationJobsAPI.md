# \HostedRepositoryEvaluationJobsAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHostedRepositoryEvaluationJobsSchedulerStart**](HostedRepositoryEvaluationJobsAPI.md#CreateHostedRepositoryEvaluationJobsSchedulerStart) | **Post** /v1/hosted-repository-evaluation/jobs/scheduler/start | Start/enable the evaluation scheduler
[**CreateHostedRepositoryEvaluationJobsSchedulerStop**](HostedRepositoryEvaluationJobsAPI.md#CreateHostedRepositoryEvaluationJobsSchedulerStop) | **Post** /v1/hosted-repository-evaluation/jobs/scheduler/stop | Stop/disable the evaluation scheduler
[**CreateHostedRepositoryEvaluationJobsSchedulerUpdateSchedule**](HostedRepositoryEvaluationJobsAPI.md#CreateHostedRepositoryEvaluationJobsSchedulerUpdateSchedule) | **Post** /v1/hosted-repository-evaluation/jobs/scheduler/update-schedule | Update the scheduler cron expression
[**ListHostedRepositoryEvaluationJobsCount**](HostedRepositoryEvaluationJobsAPI.md#ListHostedRepositoryEvaluationJobsCount) | **Get** /v1/hosted-repository-evaluation/jobs/count | Get total count of evaluation jobs
[**ListHostedRepositoryEvaluationJobsQueued**](HostedRepositoryEvaluationJobsAPI.md#ListHostedRepositoryEvaluationJobsQueued) | **Get** /v1/hosted-repository-evaluation/jobs/queued | Get all queued evaluation jobs
[**ListHostedRepositoryEvaluationJobsSchedulerStatus**](HostedRepositoryEvaluationJobsAPI.md#ListHostedRepositoryEvaluationJobsSchedulerStatus) | **Get** /v1/hosted-repository-evaluation/jobs/scheduler/status | Get scheduler status and schedule



## CreateHostedRepositoryEvaluationJobsSchedulerStart

> CreateHostedRepositoryEvaluationJobsSchedulerStart(ctx).Execute()

Start/enable the evaluation scheduler

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
	r, err := apiClient.HostedRepositoryEvaluationJobsAPI.CreateHostedRepositoryEvaluationJobsSchedulerStart(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationJobsAPI.CreateHostedRepositoryEvaluationJobsSchedulerStart``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostedRepositoryEvaluationJobsSchedulerStartRequest struct via the builder pattern


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


## CreateHostedRepositoryEvaluationJobsSchedulerStop

> CreateHostedRepositoryEvaluationJobsSchedulerStop(ctx).Execute()

Stop/disable the evaluation scheduler

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
	r, err := apiClient.HostedRepositoryEvaluationJobsAPI.CreateHostedRepositoryEvaluationJobsSchedulerStop(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationJobsAPI.CreateHostedRepositoryEvaluationJobsSchedulerStop``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostedRepositoryEvaluationJobsSchedulerStopRequest struct via the builder pattern


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


## CreateHostedRepositoryEvaluationJobsSchedulerUpdateSchedule

> CreateHostedRepositoryEvaluationJobsSchedulerUpdateSchedule(ctx).UpdateScheduleRequest(updateScheduleRequest).Execute()

Update the scheduler cron expression

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
	updateScheduleRequest := *sonatyperepo.NewUpdateScheduleRequest() // UpdateScheduleRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.HostedRepositoryEvaluationJobsAPI.CreateHostedRepositoryEvaluationJobsSchedulerUpdateSchedule(context.Background()).UpdateScheduleRequest(updateScheduleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationJobsAPI.CreateHostedRepositoryEvaluationJobsSchedulerUpdateSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHostedRepositoryEvaluationJobsSchedulerUpdateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateScheduleRequest** | [**UpdateScheduleRequest**](UpdateScheduleRequest.md) |  | 

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


## ListHostedRepositoryEvaluationJobsCount

> ListHostedRepositoryEvaluationJobsCount(ctx).Execute()

Get total count of evaluation jobs

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
	r, err := apiClient.HostedRepositoryEvaluationJobsAPI.ListHostedRepositoryEvaluationJobsCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationJobsAPI.ListHostedRepositoryEvaluationJobsCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHostedRepositoryEvaluationJobsCountRequest struct via the builder pattern


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


## ListHostedRepositoryEvaluationJobsQueued

> ListHostedRepositoryEvaluationJobsQueued(ctx).Execute()

Get all queued evaluation jobs

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
	r, err := apiClient.HostedRepositoryEvaluationJobsAPI.ListHostedRepositoryEvaluationJobsQueued(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationJobsAPI.ListHostedRepositoryEvaluationJobsQueued``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHostedRepositoryEvaluationJobsQueuedRequest struct via the builder pattern


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


## ListHostedRepositoryEvaluationJobsSchedulerStatus

> ListHostedRepositoryEvaluationJobsSchedulerStatus(ctx).Execute()

Get scheduler status and schedule

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
	r, err := apiClient.HostedRepositoryEvaluationJobsAPI.ListHostedRepositoryEvaluationJobsSchedulerStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationJobsAPI.ListHostedRepositoryEvaluationJobsSchedulerStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHostedRepositoryEvaluationJobsSchedulerStatusRequest struct via the builder pattern


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


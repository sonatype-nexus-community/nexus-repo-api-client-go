# \HostedRepositoryEvaluationAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetrepositorynameEvaluationSettingsRepository**](HostedRepositoryEvaluationAPI.md#GetrepositorynameEvaluationSettingsRepository) | **Get** /v1/repositories/{repositoryName}/evaluation-settings | Get evaluation settings for a repository
[**ListEvaluationSettings**](HostedRepositoryEvaluationAPI.md#ListEvaluationSettings) | **Get** /v1/evaluation/settings | Get global evaluation settings
[**PatchEvaluationSettingsWithRepos**](HostedRepositoryEvaluationAPI.md#PatchEvaluationSettingsWithRepos) | **Patch** /v1/evaluation/settings-with-repos | Partially update global evaluation settings and/or repository monitoring (PATCH)
[**UpdateEvaluationRepositories**](HostedRepositoryEvaluationAPI.md#UpdateEvaluationRepositories) | **Patch** /v1/evaluation/repositories | Independently update monitored repositories (add/remove without changing settings)
[**UpdateEvaluationSettings**](HostedRepositoryEvaluationAPI.md#UpdateEvaluationSettings) | **Patch** /v1/evaluation/settings | Update global evaluation settings
[**UpdateEvaluationSettingsWithRepos**](HostedRepositoryEvaluationAPI.md#UpdateEvaluationSettingsWithRepos) | **Put** /v1/evaluation/settings-with-repos | Atomically update global evaluation settings and enable repository monitoring (PUT)
[**UpdaterepositorynameEvaluationSettingsRepository**](HostedRepositoryEvaluationAPI.md#UpdaterepositorynameEvaluationSettingsRepository) | **Put** /v1/repositories/{repositoryName}/evaluation-settings | Update evaluation settings for a repository



## GetrepositorynameEvaluationSettingsRepository

> RepositoryEvaluationSettingsXO GetrepositorynameEvaluationSettingsRepository(ctx, repositoryName).Execute()

Get evaluation settings for a repository

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
	repositoryName := "repositoryName_example" // string | Repository name

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryEvaluationAPI.GetrepositorynameEvaluationSettingsRepository(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationAPI.GetrepositorynameEvaluationSettingsRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetrepositorynameEvaluationSettingsRepository`: RepositoryEvaluationSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryEvaluationAPI.GetrepositorynameEvaluationSettingsRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Repository name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetrepositorynameEvaluationSettingsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepositoryEvaluationSettingsXO**](RepositoryEvaluationSettingsXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEvaluationSettings

> GlobalEvaluationSettingsXO ListEvaluationSettings(ctx).Execute()

Get global evaluation settings

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
	resp, r, err := apiClient.HostedRepositoryEvaluationAPI.ListEvaluationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationAPI.ListEvaluationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEvaluationSettings`: GlobalEvaluationSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryEvaluationAPI.ListEvaluationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEvaluationSettingsRequest struct via the builder pattern


### Return type

[**GlobalEvaluationSettingsXO**](GlobalEvaluationSettingsXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchEvaluationSettingsWithRepos

> PatchGlobalSettingsAndRepositoriesResponseXO PatchEvaluationSettingsWithRepos(ctx).PatchGlobalSettingsAndRepositoriesRequestXO(patchGlobalSettingsAndRepositoriesRequestXO).Execute()

Partially update global evaluation settings and/or repository monitoring (PATCH)

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
	patchGlobalSettingsAndRepositoriesRequestXO := *sonatyperepo.NewPatchGlobalSettingsAndRepositoriesRequestXO() // PatchGlobalSettingsAndRepositoriesRequestXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryEvaluationAPI.PatchEvaluationSettingsWithRepos(context.Background()).PatchGlobalSettingsAndRepositoriesRequestXO(patchGlobalSettingsAndRepositoriesRequestXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationAPI.PatchEvaluationSettingsWithRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchEvaluationSettingsWithRepos`: PatchGlobalSettingsAndRepositoriesResponseXO
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryEvaluationAPI.PatchEvaluationSettingsWithRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPatchEvaluationSettingsWithReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchGlobalSettingsAndRepositoriesRequestXO** | [**PatchGlobalSettingsAndRepositoriesRequestXO**](PatchGlobalSettingsAndRepositoriesRequestXO.md) |  | 

### Return type

[**PatchGlobalSettingsAndRepositoriesResponseXO**](PatchGlobalSettingsAndRepositoriesResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEvaluationRepositories

> PatchGlobalSettingsAndRepositoriesResponseXO UpdateEvaluationRepositories(ctx).PatchGlobalSettingsAndRepositoriesRequestXO(patchGlobalSettingsAndRepositoriesRequestXO).Execute()

Independently update monitored repositories (add/remove without changing settings)

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
	patchGlobalSettingsAndRepositoriesRequestXO := *sonatyperepo.NewPatchGlobalSettingsAndRepositoriesRequestXO() // PatchGlobalSettingsAndRepositoriesRequestXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryEvaluationAPI.UpdateEvaluationRepositories(context.Background()).PatchGlobalSettingsAndRepositoriesRequestXO(patchGlobalSettingsAndRepositoriesRequestXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationAPI.UpdateEvaluationRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEvaluationRepositories`: PatchGlobalSettingsAndRepositoriesResponseXO
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryEvaluationAPI.UpdateEvaluationRepositories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEvaluationRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **patchGlobalSettingsAndRepositoriesRequestXO** | [**PatchGlobalSettingsAndRepositoriesRequestXO**](PatchGlobalSettingsAndRepositoriesRequestXO.md) |  | 

### Return type

[**PatchGlobalSettingsAndRepositoriesResponseXO**](PatchGlobalSettingsAndRepositoriesResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEvaluationSettings

> UpdateSettingsResponseXO UpdateEvaluationSettings(ctx).UpdateSettingsRequestXO(updateSettingsRequestXO).Execute()

Update global evaluation settings

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
	updateSettingsRequestXO := *sonatyperepo.NewUpdateSettingsRequestXO() // UpdateSettingsRequestXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryEvaluationAPI.UpdateEvaluationSettings(context.Background()).UpdateSettingsRequestXO(updateSettingsRequestXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationAPI.UpdateEvaluationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEvaluationSettings`: UpdateSettingsResponseXO
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryEvaluationAPI.UpdateEvaluationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEvaluationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSettingsRequestXO** | [**UpdateSettingsRequestXO**](UpdateSettingsRequestXO.md) |  | 

### Return type

[**UpdateSettingsResponseXO**](UpdateSettingsResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEvaluationSettingsWithRepos

> GlobalSettingsWithRepositoriesResponseXO UpdateEvaluationSettingsWithRepos(ctx).GlobalSettingsWithRepositoriesXO(globalSettingsWithRepositoriesXO).Execute()

Atomically update global evaluation settings and enable repository monitoring (PUT)

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
	globalSettingsWithRepositoriesXO := *sonatyperepo.NewGlobalSettingsWithRepositoriesXO(int32(30), int32(5), false, "RELEASE", []string{"["uuid1","uuid2","uuid3"]"}) // GlobalSettingsWithRepositoriesXO | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryEvaluationAPI.UpdateEvaluationSettingsWithRepos(context.Background()).GlobalSettingsWithRepositoriesXO(globalSettingsWithRepositoriesXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationAPI.UpdateEvaluationSettingsWithRepos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEvaluationSettingsWithRepos`: GlobalSettingsWithRepositoriesResponseXO
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryEvaluationAPI.UpdateEvaluationSettingsWithRepos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEvaluationSettingsWithReposRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalSettingsWithRepositoriesXO** | [**GlobalSettingsWithRepositoriesXO**](GlobalSettingsWithRepositoriesXO.md) |  | 

### Return type

[**GlobalSettingsWithRepositoriesResponseXO**](GlobalSettingsWithRepositoriesResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdaterepositorynameEvaluationSettingsRepository

> UpdaterepositorynameEvaluationSettingsRepository(ctx, repositoryName).RepositoryEvaluationSettingsXO(repositoryEvaluationSettingsXO).Execute()

Update evaluation settings for a repository

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
	repositoryName := "repositoryName_example" // string | Repository name
	repositoryEvaluationSettingsXO := *sonatyperepo.NewRepositoryEvaluationSettingsXO("INHERIT") // RepositoryEvaluationSettingsXO | Evaluation settings

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.HostedRepositoryEvaluationAPI.UpdaterepositorynameEvaluationSettingsRepository(context.Background(), repositoryName).RepositoryEvaluationSettingsXO(repositoryEvaluationSettingsXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryEvaluationAPI.UpdaterepositorynameEvaluationSettingsRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Repository name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdaterepositorynameEvaluationSettingsRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repositoryEvaluationSettingsXO** | [**RepositoryEvaluationSettingsXO**](RepositoryEvaluationSettingsXO.md) | Evaluation settings | 

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


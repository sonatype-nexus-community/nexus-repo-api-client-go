# \RepositoryManagementAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRepository**](RepositoryManagementAPI.md#CreateRepository) | **Post** /v1/repositories/maven/group | Create Maven group repository
[**CreateRepository1**](RepositoryManagementAPI.md#CreateRepository1) | **Post** /v1/repositories/maven/hosted | Create Maven hosted repository
[**CreateRepository10**](RepositoryManagementAPI.md#CreateRepository10) | **Post** /v1/repositories/npm/proxy | Create npm proxy repository
[**CreateRepository11**](RepositoryManagementAPI.md#CreateRepository11) | **Post** /v1/repositories/nuget/group | Create NuGet group repository
[**CreateRepository12**](RepositoryManagementAPI.md#CreateRepository12) | **Post** /v1/repositories/nuget/hosted | Create NuGet hosted repository
[**CreateRepository13**](RepositoryManagementAPI.md#CreateRepository13) | **Post** /v1/repositories/nuget/proxy | Create NuGet proxy repository
[**CreateRepository14**](RepositoryManagementAPI.md#CreateRepository14) | **Post** /v1/repositories/rubygems/group | Create RubyGems group repository
[**CreateRepository15**](RepositoryManagementAPI.md#CreateRepository15) | **Post** /v1/repositories/rubygems/hosted | Create RubyGems hosted repository
[**CreateRepository16**](RepositoryManagementAPI.md#CreateRepository16) | **Post** /v1/repositories/rubygems/proxy | Create RubyGems proxy repository
[**CreateRepository17**](RepositoryManagementAPI.md#CreateRepository17) | **Post** /v1/repositories/yum/group | Create Yum group repository
[**CreateRepository18**](RepositoryManagementAPI.md#CreateRepository18) | **Post** /v1/repositories/yum/hosted | Create Yum hosted repository
[**CreateRepository19**](RepositoryManagementAPI.md#CreateRepository19) | **Post** /v1/repositories/yum/proxy | Create Yum proxy repository
[**CreateRepository2**](RepositoryManagementAPI.md#CreateRepository2) | **Post** /v1/repositories/maven/proxy | Create Maven proxy repository
[**CreateRepository20**](RepositoryManagementAPI.md#CreateRepository20) | **Post** /v1/repositories/docker/group | Create Docker group repository
[**CreateRepository21**](RepositoryManagementAPI.md#CreateRepository21) | **Post** /v1/repositories/docker/hosted | Create Docker hosted repository
[**CreateRepository22**](RepositoryManagementAPI.md#CreateRepository22) | **Post** /v1/repositories/docker/proxy | Create Docker proxy repository
[**CreateRepository23**](RepositoryManagementAPI.md#CreateRepository23) | **Post** /v1/repositories/pypi/group | Create PyPI group repository
[**CreateRepository24**](RepositoryManagementAPI.md#CreateRepository24) | **Post** /v1/repositories/pypi/hosted | Create PyPI hosted repository
[**CreateRepository25**](RepositoryManagementAPI.md#CreateRepository25) | **Post** /v1/repositories/pypi/proxy | Create PyPI proxy repository
[**CreateRepository26**](RepositoryManagementAPI.md#CreateRepository26) | **Post** /v1/repositories/conda/proxy | Create conda proxy repository
[**CreateRepository27**](RepositoryManagementAPI.md#CreateRepository27) | **Post** /v1/repositories/conan/proxy | Create Conan proxy repository
[**CreateRepository28**](RepositoryManagementAPI.md#CreateRepository28) | **Post** /v1/repositories/gitlfs/hosted | Create Git LFS hosted repository
[**CreateRepository29**](RepositoryManagementAPI.md#CreateRepository29) | **Post** /v1/repositories/r/group | Create R group repository
[**CreateRepository3**](RepositoryManagementAPI.md#CreateRepository3) | **Post** /v1/repositories/apt/hosted | Create APT hosted repository
[**CreateRepository30**](RepositoryManagementAPI.md#CreateRepository30) | **Post** /v1/repositories/r/hosted | Create R hosted repository
[**CreateRepository31**](RepositoryManagementAPI.md#CreateRepository31) | **Post** /v1/repositories/r/proxy | Create R proxy repository
[**CreateRepository32**](RepositoryManagementAPI.md#CreateRepository32) | **Post** /v1/repositories/cocoapods/proxy | Create Cocoapods proxy repository
[**CreateRepository33**](RepositoryManagementAPI.md#CreateRepository33) | **Post** /v1/repositories/go/group | Create a Go group repository
[**CreateRepository34**](RepositoryManagementAPI.md#CreateRepository34) | **Post** /v1/repositories/go/proxy | Create a Go proxy repository
[**CreateRepository35**](RepositoryManagementAPI.md#CreateRepository35) | **Post** /v1/repositories/p2/proxy | Create p2 proxy repository
[**CreateRepository36**](RepositoryManagementAPI.md#CreateRepository36) | **Post** /v1/repositories/helm/hosted | Create Helm hosted repository
[**CreateRepository37**](RepositoryManagementAPI.md#CreateRepository37) | **Post** /v1/repositories/helm/proxy | Create Helm proxy repository
[**CreateRepository38**](RepositoryManagementAPI.md#CreateRepository38) | **Post** /v1/repositories/bower/group | Create Bower group repository
[**CreateRepository39**](RepositoryManagementAPI.md#CreateRepository39) | **Post** /v1/repositories/bower/hosted | Create Bower hosted repository
[**CreateRepository4**](RepositoryManagementAPI.md#CreateRepository4) | **Post** /v1/repositories/apt/proxy | Create APT proxy repository
[**CreateRepository40**](RepositoryManagementAPI.md#CreateRepository40) | **Post** /v1/repositories/bower/proxy | Create Bower proxy repository
[**CreateRepository5**](RepositoryManagementAPI.md#CreateRepository5) | **Post** /v1/repositories/raw/group | Create raw group repository
[**CreateRepository6**](RepositoryManagementAPI.md#CreateRepository6) | **Post** /v1/repositories/raw/hosted | Create raw hosted repository
[**CreateRepository7**](RepositoryManagementAPI.md#CreateRepository7) | **Post** /v1/repositories/raw/proxy | Create raw proxy repository
[**CreateRepository8**](RepositoryManagementAPI.md#CreateRepository8) | **Post** /v1/repositories/npm/group | Create npm group repository
[**CreateRepository9**](RepositoryManagementAPI.md#CreateRepository9) | **Post** /v1/repositories/npm/hosted | Create npm hosted repository
[**DeleteRepository**](RepositoryManagementAPI.md#DeleteRepository) | **Delete** /v1/repositories/{repositoryName} | Delete repository of any format
[**DisableRepositoryHealthCheck**](RepositoryManagementAPI.md#DisableRepositoryHealthCheck) | **Delete** /v1/repositories/{repositoryName}/health-check | Disable repository health check. Proxy repositories only.
[**EnableRepositoryHealthCheck**](RepositoryManagementAPI.md#EnableRepositoryHealthCheck) | **Post** /v1/repositories/{repositoryName}/health-check | Enable repository health check. Proxy repositories only.
[**GetRepositories**](RepositoryManagementAPI.md#GetRepositories) | **Get** /v1/repositorySettings | List repositories
[**GetRepositories1**](RepositoryManagementAPI.md#GetRepositories1) | **Get** /v1/repositories | List repositories
[**GetRepository**](RepositoryManagementAPI.md#GetRepository) | **Get** /v1/repositories/{repositoryName} | Get repository details
[**GetRepository1**](RepositoryManagementAPI.md#GetRepository1) | **Get** /v1/repositories/maven/group/{repositoryName} | Get repository
[**GetRepository10**](RepositoryManagementAPI.md#GetRepository10) | **Get** /v1/repositories/npm/hosted/{repositoryName} | Get repository
[**GetRepository11**](RepositoryManagementAPI.md#GetRepository11) | **Get** /v1/repositories/npm/proxy/{repositoryName} | Get repository
[**GetRepository12**](RepositoryManagementAPI.md#GetRepository12) | **Get** /v1/repositories/nuget/group/{repositoryName} | Get repository
[**GetRepository13**](RepositoryManagementAPI.md#GetRepository13) | **Get** /v1/repositories/nuget/hosted/{repositoryName} | Get repository
[**GetRepository14**](RepositoryManagementAPI.md#GetRepository14) | **Get** /v1/repositories/nuget/proxy/{repositoryName} | Get repository
[**GetRepository15**](RepositoryManagementAPI.md#GetRepository15) | **Get** /v1/repositories/rubygems/group/{repositoryName} | Get repository
[**GetRepository16**](RepositoryManagementAPI.md#GetRepository16) | **Get** /v1/repositories/rubygems/hosted/{repositoryName} | Get repository
[**GetRepository17**](RepositoryManagementAPI.md#GetRepository17) | **Get** /v1/repositories/rubygems/proxy/{repositoryName} | Get repository
[**GetRepository18**](RepositoryManagementAPI.md#GetRepository18) | **Get** /v1/repositories/yum/group/{repositoryName} | Get repository
[**GetRepository19**](RepositoryManagementAPI.md#GetRepository19) | **Get** /v1/repositories/yum/hosted/{repositoryName} | Get repository
[**GetRepository2**](RepositoryManagementAPI.md#GetRepository2) | **Get** /v1/repositories/maven/hosted/{repositoryName} | Get repository
[**GetRepository20**](RepositoryManagementAPI.md#GetRepository20) | **Get** /v1/repositories/yum/proxy/{repositoryName} | Get repository
[**GetRepository21**](RepositoryManagementAPI.md#GetRepository21) | **Get** /v1/repositories/docker/group/{repositoryName} | Get repository
[**GetRepository22**](RepositoryManagementAPI.md#GetRepository22) | **Get** /v1/repositories/docker/hosted/{repositoryName} | Get repository
[**GetRepository23**](RepositoryManagementAPI.md#GetRepository23) | **Get** /v1/repositories/docker/proxy/{repositoryName} | Get repository
[**GetRepository24**](RepositoryManagementAPI.md#GetRepository24) | **Get** /v1/repositories/pypi/group/{repositoryName} | Get repository
[**GetRepository25**](RepositoryManagementAPI.md#GetRepository25) | **Get** /v1/repositories/pypi/hosted/{repositoryName} | Get repository
[**GetRepository26**](RepositoryManagementAPI.md#GetRepository26) | **Get** /v1/repositories/pypi/proxy/{repositoryName} | Get repository
[**GetRepository27**](RepositoryManagementAPI.md#GetRepository27) | **Get** /v1/repositories/conda/proxy/{repositoryName} | Get repository
[**GetRepository28**](RepositoryManagementAPI.md#GetRepository28) | **Get** /v1/repositories/conan/proxy/{repositoryName} | Get repository
[**GetRepository29**](RepositoryManagementAPI.md#GetRepository29) | **Get** /v1/repositories/gitlfs/hosted/{repositoryName} | Get repository
[**GetRepository3**](RepositoryManagementAPI.md#GetRepository3) | **Get** /v1/repositories/maven/proxy/{repositoryName} | Get repository
[**GetRepository30**](RepositoryManagementAPI.md#GetRepository30) | **Get** /v1/repositories/r/group/{repositoryName} | Get repository
[**GetRepository31**](RepositoryManagementAPI.md#GetRepository31) | **Get** /v1/repositories/r/hosted/{repositoryName} | Get repository
[**GetRepository32**](RepositoryManagementAPI.md#GetRepository32) | **Get** /v1/repositories/r/proxy/{repositoryName} | Get repository
[**GetRepository33**](RepositoryManagementAPI.md#GetRepository33) | **Get** /v1/repositories/cocoapods/proxy/{repositoryName} | Get repository
[**GetRepository34**](RepositoryManagementAPI.md#GetRepository34) | **Get** /v1/repositories/go/group/{repositoryName} | Get repository
[**GetRepository35**](RepositoryManagementAPI.md#GetRepository35) | **Get** /v1/repositories/go/proxy/{repositoryName} | Get repository
[**GetRepository36**](RepositoryManagementAPI.md#GetRepository36) | **Get** /v1/repositories/p2/proxy/{repositoryName} | Get repository
[**GetRepository37**](RepositoryManagementAPI.md#GetRepository37) | **Get** /v1/repositories/helm/hosted/{repositoryName} | Get repository
[**GetRepository38**](RepositoryManagementAPI.md#GetRepository38) | **Get** /v1/repositories/helm/proxy/{repositoryName} | Get repository
[**GetRepository39**](RepositoryManagementAPI.md#GetRepository39) | **Get** /v1/repositories/bower/group/{repositoryName} | Get repository
[**GetRepository4**](RepositoryManagementAPI.md#GetRepository4) | **Get** /v1/repositories/apt/hosted/{repositoryName} | Get repository
[**GetRepository40**](RepositoryManagementAPI.md#GetRepository40) | **Get** /v1/repositories/bower/hosted/{repositoryName} | Get repository
[**GetRepository41**](RepositoryManagementAPI.md#GetRepository41) | **Get** /v1/repositories/bower/proxy/{repositoryName} | Get repository
[**GetRepository5**](RepositoryManagementAPI.md#GetRepository5) | **Get** /v1/repositories/apt/proxy/{repositoryName} | Get repository
[**GetRepository6**](RepositoryManagementAPI.md#GetRepository6) | **Get** /v1/repositories/raw/group/{repositoryName} | Get repository
[**GetRepository7**](RepositoryManagementAPI.md#GetRepository7) | **Get** /v1/repositories/raw/hosted/{repositoryName} | Get repository
[**GetRepository8**](RepositoryManagementAPI.md#GetRepository8) | **Get** /v1/repositories/raw/proxy/{repositoryName} | Get repository
[**GetRepository9**](RepositoryManagementAPI.md#GetRepository9) | **Get** /v1/repositories/npm/group/{repositoryName} | Get repository
[**InvalidateCache**](RepositoryManagementAPI.md#InvalidateCache) | **Post** /v1/repositories/{repositoryName}/invalidate-cache | Invalidate repository cache. Proxy or group repositories only.
[**RebuildIndex**](RepositoryManagementAPI.md#RebuildIndex) | **Post** /v1/repositories/{repositoryName}/rebuild-index | Schedule a &#39;Repair - Rebuild repository search&#39; Task. Hosted or proxy repositories only.
[**UpdateRepository**](RepositoryManagementAPI.md#UpdateRepository) | **Put** /v1/repositories/maven/group/{repositoryName} | Update Maven group repository
[**UpdateRepository1**](RepositoryManagementAPI.md#UpdateRepository1) | **Put** /v1/repositories/maven/hosted/{repositoryName} | Update Maven hosted repository
[**UpdateRepository10**](RepositoryManagementAPI.md#UpdateRepository10) | **Put** /v1/repositories/npm/proxy/{repositoryName} | Update npm proxy repository
[**UpdateRepository11**](RepositoryManagementAPI.md#UpdateRepository11) | **Put** /v1/repositories/nuget/group/{repositoryName} | Update NuGet group repository
[**UpdateRepository12**](RepositoryManagementAPI.md#UpdateRepository12) | **Put** /v1/repositories/nuget/hosted/{repositoryName} | Update NuGet hosted repository
[**UpdateRepository13**](RepositoryManagementAPI.md#UpdateRepository13) | **Put** /v1/repositories/nuget/proxy/{repositoryName} | Update NuGet proxy repository
[**UpdateRepository14**](RepositoryManagementAPI.md#UpdateRepository14) | **Put** /v1/repositories/rubygems/group/{repositoryName} | Update RubyGems group repository
[**UpdateRepository15**](RepositoryManagementAPI.md#UpdateRepository15) | **Put** /v1/repositories/rubygems/hosted/{repositoryName} | Update RubyGems hosted repository
[**UpdateRepository16**](RepositoryManagementAPI.md#UpdateRepository16) | **Put** /v1/repositories/rubygems/proxy/{repositoryName} | Update RubyGems proxy repository
[**UpdateRepository17**](RepositoryManagementAPI.md#UpdateRepository17) | **Put** /v1/repositories/yum/group/{repositoryName} | Update Yum group repository
[**UpdateRepository18**](RepositoryManagementAPI.md#UpdateRepository18) | **Put** /v1/repositories/yum/hosted/{repositoryName} | Update Yum hosted repository
[**UpdateRepository19**](RepositoryManagementAPI.md#UpdateRepository19) | **Put** /v1/repositories/yum/proxy/{repositoryName} | Update Yum proxy repository
[**UpdateRepository2**](RepositoryManagementAPI.md#UpdateRepository2) | **Put** /v1/repositories/maven/proxy/{repositoryName} | Update Maven proxy repository
[**UpdateRepository20**](RepositoryManagementAPI.md#UpdateRepository20) | **Put** /v1/repositories/docker/group/{repositoryName} | Update Docker group repository
[**UpdateRepository21**](RepositoryManagementAPI.md#UpdateRepository21) | **Put** /v1/repositories/docker/hosted/{repositoryName} | Update Docker hosted repository
[**UpdateRepository22**](RepositoryManagementAPI.md#UpdateRepository22) | **Put** /v1/repositories/docker/proxy/{repositoryName} | Update Docker proxy repository
[**UpdateRepository23**](RepositoryManagementAPI.md#UpdateRepository23) | **Put** /v1/repositories/pypi/group/{repositoryName} | Update PyPI group repository
[**UpdateRepository24**](RepositoryManagementAPI.md#UpdateRepository24) | **Put** /v1/repositories/pypi/hosted/{repositoryName} | Update PyPI hosted repository
[**UpdateRepository25**](RepositoryManagementAPI.md#UpdateRepository25) | **Put** /v1/repositories/pypi/proxy/{repositoryName} | Update PyPI proxy repository
[**UpdateRepository26**](RepositoryManagementAPI.md#UpdateRepository26) | **Put** /v1/repositories/conda/proxy/{repositoryName} | Update conda proxy repository
[**UpdateRepository27**](RepositoryManagementAPI.md#UpdateRepository27) | **Put** /v1/repositories/conan/proxy/{repositoryName} | Update Conan proxy repository
[**UpdateRepository28**](RepositoryManagementAPI.md#UpdateRepository28) | **Put** /v1/repositories/gitlfs/hosted/{repositoryName} | Update Git LFS hosted repository
[**UpdateRepository29**](RepositoryManagementAPI.md#UpdateRepository29) | **Put** /v1/repositories/r/group/{repositoryName} | Update R group repository
[**UpdateRepository3**](RepositoryManagementAPI.md#UpdateRepository3) | **Put** /v1/repositories/apt/hosted/{repositoryName} | Update APT hosted repository
[**UpdateRepository30**](RepositoryManagementAPI.md#UpdateRepository30) | **Put** /v1/repositories/r/hosted/{repositoryName} | Update R hosted repository
[**UpdateRepository31**](RepositoryManagementAPI.md#UpdateRepository31) | **Put** /v1/repositories/r/proxy/{repositoryName} | Update R proxy repository
[**UpdateRepository32**](RepositoryManagementAPI.md#UpdateRepository32) | **Put** /v1/repositories/cocoapods/proxy/{repositoryName} | Update Cocoapods proxy repository
[**UpdateRepository33**](RepositoryManagementAPI.md#UpdateRepository33) | **Put** /v1/repositories/go/group/{repositoryName} | Update a Go group repository
[**UpdateRepository34**](RepositoryManagementAPI.md#UpdateRepository34) | **Put** /v1/repositories/go/proxy/{repositoryName} | Update a Go proxy repository
[**UpdateRepository35**](RepositoryManagementAPI.md#UpdateRepository35) | **Put** /v1/repositories/p2/proxy/{repositoryName} | Update p2 proxy repository
[**UpdateRepository36**](RepositoryManagementAPI.md#UpdateRepository36) | **Put** /v1/repositories/helm/hosted/{repositoryName} | Update Helm hosted repository
[**UpdateRepository37**](RepositoryManagementAPI.md#UpdateRepository37) | **Put** /v1/repositories/helm/proxy/{repositoryName} | Update Helm proxy repository
[**UpdateRepository38**](RepositoryManagementAPI.md#UpdateRepository38) | **Put** /v1/repositories/bower/group/{repositoryName} | Update Bower group repository
[**UpdateRepository39**](RepositoryManagementAPI.md#UpdateRepository39) | **Put** /v1/repositories/bower/hosted/{repositoryName} | Update Bower hosted repository
[**UpdateRepository4**](RepositoryManagementAPI.md#UpdateRepository4) | **Put** /v1/repositories/apt/proxy/{repositoryName} | Update APT proxy repository
[**UpdateRepository40**](RepositoryManagementAPI.md#UpdateRepository40) | **Put** /v1/repositories/bower/proxy/{repositoryName} | Update Bower proxy repository
[**UpdateRepository5**](RepositoryManagementAPI.md#UpdateRepository5) | **Put** /v1/repositories/raw/group/{repositoryName} | Update raw group repository
[**UpdateRepository6**](RepositoryManagementAPI.md#UpdateRepository6) | **Put** /v1/repositories/raw/hosted/{repositoryName} | Update raw hosted repository
[**UpdateRepository7**](RepositoryManagementAPI.md#UpdateRepository7) | **Put** /v1/repositories/raw/proxy/{repositoryName} | Update raw proxy repository
[**UpdateRepository8**](RepositoryManagementAPI.md#UpdateRepository8) | **Put** /v1/repositories/npm/group/{repositoryName} | Update npm group repository
[**UpdateRepository9**](RepositoryManagementAPI.md#UpdateRepository9) | **Put** /v1/repositories/npm/hosted/{repositoryName} | Update npm hosted repository



## CreateRepository

> CreateRepository(ctx).Body(body).Execute()

Create Maven group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewMavenGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // MavenGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MavenGroupRepositoryApiRequest**](MavenGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository1

> CreateRepository1(ctx).Body(body).Execute()

Create Maven hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewMavenHostedRepositoryApiRequest(*sonatyperepo.NewMavenAttributes(), "internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // MavenHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository1(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MavenHostedRepositoryApiRequest**](MavenHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository10

> CreateRepository10(ctx).Body(body).Execute()

Create npm proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewNpmProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // NpmProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository10(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository10``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NpmProxyRepositoryApiRequest**](NpmProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository11

> CreateRepository11(ctx).Body(body).Execute()

Create NuGet group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewNugetGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // NugetGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository11(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository11``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NugetGroupRepositoryApiRequest**](NugetGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository12

> CreateRepository12(ctx).Body(body).Execute()

Create NuGet hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewNugetHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // NugetHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository12(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository12``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NugetHostedRepositoryApiRequest**](NugetHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository13

> CreateRepository13(ctx).Body(body).Execute()

Create NuGet proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewNugetProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), *sonatyperepo.NewNugetAttributes(), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // NugetProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository13(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository13``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NugetProxyRepositoryApiRequest**](NugetProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository14

> CreateRepository14(ctx).Body(body).Execute()

Create RubyGems group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRubyGemsGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // RubyGemsGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository14(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository14``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RubyGemsGroupRepositoryApiRequest**](RubyGemsGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository15

> CreateRepository15(ctx).Body(body).Execute()

Create RubyGems hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRubyGemsHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // RubyGemsHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository15(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository15``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RubyGemsHostedRepositoryApiRequest**](RubyGemsHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository16

> CreateRepository16(ctx).Body(body).Execute()

Create RubyGems proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRubyGemsProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // RubyGemsProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository16(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository16``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RubyGemsProxyRepositoryApiRequest**](RubyGemsProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository17

> CreateRepository17(ctx).Body(body).Execute()

Create Yum group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewYumGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // YumGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository17(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository17``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**YumGroupRepositoryApiRequest**](YumGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository18

> CreateRepository18(ctx).Body(body).Execute()

Create Yum hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewYumHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once"), *sonatyperepo.NewYumAttributes(int32(5))) // YumHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository18(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository18``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository18Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**YumHostedRepositoryApiRequest**](YumHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository19

> CreateRepository19(ctx).Body(body).Execute()

Create Yum proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewYumProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // YumProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository19(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository19``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository19Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**YumProxyRepositoryApiRequest**](YumProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository2

> CreateRepository2(ctx).Body(body).Execute()

Create Maven proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewMavenProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributesWithPreemptiveAuth(true, false), *sonatyperepo.NewMavenAttributes(), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // MavenProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MavenProxyRepositoryApiRequest**](MavenProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository20

> CreateRepository20(ctx).Body(body).Execute()

Create Docker group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewDockerGroupRepositoryApiRequest(*sonatyperepo.NewDockerAttributes(true, false), *sonatyperepo.NewGroupDeployAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // DockerGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository20(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository20``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DockerGroupRepositoryApiRequest**](DockerGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository21

> CreateRepository21(ctx).Body(body).Execute()

Create Docker hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewDockerHostedRepositoryApiRequest(*sonatyperepo.NewDockerAttributes(true, false), "internal", true, *sonatyperepo.NewDockerHostedStorageAttributes("default", true, "allow_once")) // DockerHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository21(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository21``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DockerHostedRepositoryApiRequest**](DockerHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository22

> CreateRepository22(ctx).Body(body).Execute()

Create Docker proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewDockerProxyRepositoryApiRequest(*sonatyperepo.NewDockerAttributes(true, false), *sonatyperepo.NewDockerProxyAttributes(), *sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // DockerProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository22(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository22``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository22Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DockerProxyRepositoryApiRequest**](DockerProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository23

> CreateRepository23(ctx).Body(body).Execute()

Create PyPI group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewPypiGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // PypiGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository23(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository23``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository23Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PypiGroupRepositoryApiRequest**](PypiGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository24

> CreateRepository24(ctx).Body(body).Execute()

Create PyPI hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewPypiHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // PypiHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository24(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository24``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository24Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PypiHostedRepositoryApiRequest**](PypiHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository25

> CreateRepository25(ctx).Body(body).Execute()

Create PyPI proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewPypiProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // PypiProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository25(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository25``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository25Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PypiProxyRepositoryApiRequest**](PypiProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository26

> CreateRepository26(ctx).Body(body).Execute()

Create conda proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewCondaProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // CondaProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository26(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository26``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository26Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CondaProxyRepositoryApiRequest**](CondaProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository27

> CreateRepository27(ctx).Body(body).Execute()

Create Conan proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewConanProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // ConanProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository27(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository27``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConanProxyRepositoryApiRequest**](ConanProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository28

> CreateRepository28(ctx).Body(body).Execute()

Create Git LFS hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewGitLfsHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // GitLfsHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository28(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository28``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GitLfsHostedRepositoryApiRequest**](GitLfsHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository29

> CreateRepository29(ctx).Body(body).Execute()

Create R group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // RGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository29(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository29``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository29Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RGroupRepositoryApiRequest**](RGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository3

> CreateRepository3(ctx).Body(body).Execute()

Create APT hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewAptHostedRepositoryApiRequest(*sonatyperepo.NewAptHostedRepositoriesAttributes(), *sonatyperepo.NewAptSigningRepositoriesAttributes(), "internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // AptHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository3(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AptHostedRepositoryApiRequest**](AptHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository30

> CreateRepository30(ctx).Body(body).Execute()

Create R hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // RHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository30(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository30``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository30Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RHostedRepositoryApiRequest**](RHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository31

> CreateRepository31(ctx).Body(body).Execute()

Create R proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // RProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository31(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository31``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository31Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RProxyRepositoryApiRequest**](RProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository32

> CreateRepository32(ctx).Body(body).Execute()

Create Cocoapods proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewCocoapodsProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // CocoapodsProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository32(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository32``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository32Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CocoapodsProxyRepositoryApiRequest**](CocoapodsProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository33

> CreateRepository33(ctx).Body(body).Execute()

Create a Go group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewGolangGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // GolangGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository33(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository33``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository33Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GolangGroupRepositoryApiRequest**](GolangGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository34

> CreateRepository34(ctx).Body(body).Execute()

Create a Go proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewGolangProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // GolangProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository34(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository34``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository34Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GolangProxyRepositoryApiRequest**](GolangProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository35

> CreateRepository35(ctx).Body(body).Execute()

Create p2 proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewP2ProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // P2ProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository35(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository35``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository35Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**P2ProxyRepositoryApiRequest**](P2ProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository36

> CreateRepository36(ctx).Body(body).Execute()

Create Helm hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewHelmHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // HelmHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository36(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository36``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository36Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HelmHostedRepositoryApiRequest**](HelmHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository37

> CreateRepository37(ctx).Body(body).Execute()

Create Helm proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewHelmProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // HelmProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository37(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository37``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository37Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HelmProxyRepositoryApiRequest**](HelmProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository38

> CreateRepository38(ctx).Body(body).Execute()

Create Bower group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewBowerGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // BowerGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository38(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository38``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository38Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BowerGroupRepositoryApiRequest**](BowerGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository39

> CreateRepository39(ctx).Body(body).Execute()

Create Bower hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewBowerHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // BowerHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository39(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository39``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository39Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BowerHostedRepositoryApiRequest**](BowerHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository4

> CreateRepository4(ctx).Body(body).Execute()

Create APT proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewAptProxyRepositoryApiRequest(*sonatyperepo.NewAptProxyRepositoriesAttributes(false), *sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // AptProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository4(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AptProxyRepositoryApiRequest**](AptProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository40

> CreateRepository40(ctx).Body(body).Execute()

Create Bower proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewBowerProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // BowerProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository40(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository40``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository40Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BowerProxyRepositoryApiRequest**](BowerProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository5

> CreateRepository5(ctx).Body(body).Execute()

Create raw group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRawGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // RawGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository5(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RawGroupRepositoryApiRequest**](RawGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository6

> CreateRepository6(ctx).Body(body).Execute()

Create raw hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRawHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // RawHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository6(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RawHostedRepositoryApiRequest**](RawHostedRepositoryApiRequest.md) |  | 

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


## CreateRepository7

> CreateRepository7(ctx).Body(body).Execute()

Create raw proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewRawProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // RawProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository7(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RawProxyRepositoryApiRequest**](RawProxyRepositoryApiRequest.md) |  | 

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


## CreateRepository8

> CreateRepository8(ctx).Body(body).Execute()

Create npm group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewNpmGroupRepositoryApiRequest(*sonatyperepo.NewGroupDeployAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // NpmGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository8(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NpmGroupRepositoryApiRequest**](NpmGroupRepositoryApiRequest.md) |  | 

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


## CreateRepository9

> CreateRepository9(ctx).Body(body).Execute()

Create npm hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	body := *sonatyperepo.NewNpmHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // NpmHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.CreateRepository9(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.CreateRepository9``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepository9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NpmHostedRepositoryApiRequest**](NpmHostedRepositoryApiRequest.md) |  | 

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


## DeleteRepository

> DeleteRepository(ctx, repositoryName).Execute()

Delete repository of any format

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.DeleteRepository(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.DeleteRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


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


## DisableRepositoryHealthCheck

> DisableRepositoryHealthCheck(ctx, repositoryName).Execute()

Disable repository health check. Proxy repositories only.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to disable Repository Health Check for

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.DisableRepositoryHealthCheck(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.DisableRepositoryHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to disable Repository Health Check for | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableRepositoryHealthCheckRequest struct via the builder pattern


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


## EnableRepositoryHealthCheck

> EnableRepositoryHealthCheck(ctx, repositoryName).Execute()

Enable repository health check. Proxy repositories only.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to enable Repository Health Check for

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.EnableRepositoryHealthCheck(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.EnableRepositoryHealthCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to enable Repository Health Check for | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableRepositoryHealthCheckRequest struct via the builder pattern


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


## GetRepositories

> []AbstractApiRepository GetRepositories(ctx).Execute()

List repositories

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepositories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories`: []AbstractApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


### Return type

[**[]AbstractApiRepository**](AbstractApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories1

> []RepositoryXO GetRepositories1(ctx).Execute()

List repositories

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepositories1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepositories1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories1`: []RepositoryXO
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepositories1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositories1Request struct via the builder pattern


### Return type

[**[]RepositoryXO**](RepositoryXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> RepositoryXO GetRepository(ctx, repositoryName).Execute()

Get repository details

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to get

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository`: RepositoryXO
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RepositoryXO**](RepositoryXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository1

> SimpleApiGroupRepository GetRepository1(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository1(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository1`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository10

> SimpleApiHostedRepository GetRepository10(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository10(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository10``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository10`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository10`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository11

> NpmProxyApiRepository GetRepository11(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository11(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository11``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository11`: NpmProxyApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository11`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NpmProxyApiRepository**](NpmProxyApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository12

> SimpleApiGroupRepository GetRepository12(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository12(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository12``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository12`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository12`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository13

> SimpleApiHostedRepository GetRepository13(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository13(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository13``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository13`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository13`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository14

> NugetProxyApiRepository GetRepository14(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository14(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository14``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository14`: NugetProxyApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository14`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NugetProxyApiRepository**](NugetProxyApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository15

> SimpleApiGroupRepository GetRepository15(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository15(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository15``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository15`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository15`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository16

> SimpleApiHostedRepository GetRepository16(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository16(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository16``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository16`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository16`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository17

> SimpleApiProxyRepository GetRepository17(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository17(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository17``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository17`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository17`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository18

> SimpleApiGroupRepository GetRepository18(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository18(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository18``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository18`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository18`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository18Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository19

> YumHostedApiRepository GetRepository19(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository19(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository19``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository19`: YumHostedApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository19`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository19Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**YumHostedApiRepository**](YumHostedApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository2

> MavenHostedApiRepository GetRepository2(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository2(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository2`: MavenHostedApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MavenHostedApiRepository**](MavenHostedApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository20

> SimpleApiProxyRepository GetRepository20(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository20(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository20``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository20`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository20`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository21

> DockerGroupApiRepository GetRepository21(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository21(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository21``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository21`: DockerGroupApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository21`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockerGroupApiRepository**](DockerGroupApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository22

> DockerHostedApiRepository GetRepository22(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository22(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository22``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository22`: DockerHostedApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository22`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository22Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockerHostedApiRepository**](DockerHostedApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository23

> DockerProxyApiRepository GetRepository23(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository23(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository23``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository23`: DockerProxyApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository23`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository23Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DockerProxyApiRepository**](DockerProxyApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository24

> SimpleApiGroupRepository GetRepository24(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository24(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository24``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository24`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository24`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository24Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository25

> SimpleApiHostedRepository GetRepository25(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository25(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository25``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository25`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository25`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository25Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository26

> SimpleApiProxyRepository GetRepository26(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository26(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository26``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository26`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository26`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository26Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository27

> SimpleApiProxyRepository GetRepository27(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository27(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository27``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository27`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository27`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository28

> SimpleApiProxyRepository GetRepository28(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository28(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository28``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository28`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository28`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository29

> SimpleApiHostedRepository GetRepository29(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository29(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository29``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository29`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository29`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository29Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository3

> MavenProxyApiRepository GetRepository3(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository3(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository3`: MavenProxyApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MavenProxyApiRepository**](MavenProxyApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository30

> SimpleApiGroupRepository GetRepository30(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository30(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository30``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository30`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository30`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository30Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository31

> SimpleApiHostedRepository GetRepository31(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository31(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository31``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository31`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository31`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository31Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository32

> SimpleApiProxyRepository GetRepository32(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository32(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository32``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository32`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository32`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository32Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository33

> SimpleApiProxyRepository GetRepository33(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository33(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository33``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository33`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository33`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository33Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository34

> SimpleApiGroupRepository GetRepository34(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository34(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository34``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository34`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository34`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository34Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository35

> SimpleApiProxyRepository GetRepository35(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository35(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository35``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository35`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository35`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository35Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository36

> SimpleApiProxyRepository GetRepository36(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository36(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository36``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository36`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository36`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository36Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository37

> SimpleApiHostedRepository GetRepository37(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository37(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository37``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository37`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository37`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository37Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository38

> SimpleApiProxyRepository GetRepository38(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository38(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository38``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository38`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository38`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository38Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository39

> SimpleApiGroupRepository GetRepository39(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository39(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository39``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository39`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository39`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository39Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository4

> AptHostedApiRepository GetRepository4(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository4(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository4`: AptHostedApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AptHostedApiRepository**](AptHostedApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository40

> SimpleApiHostedRepository GetRepository40(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository40(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository40``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository40`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository40`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository40Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository41

> BowerProxyApiRepository GetRepository41(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository41(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository41``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository41`: BowerProxyApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository41`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository41Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BowerProxyApiRepository**](BowerProxyApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository5

> AptProxyApiRepository GetRepository5(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository5(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository5`: AptProxyApiRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository5`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AptProxyApiRepository**](AptProxyApiRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository6

> SimpleApiGroupRepository GetRepository6(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository6(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository6`: SimpleApiGroupRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupRepository**](SimpleApiGroupRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository7

> SimpleApiHostedRepository GetRepository7(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository7(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository7`: SimpleApiHostedRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository7`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiHostedRepository**](SimpleApiHostedRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository8

> SimpleApiProxyRepository GetRepository8(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository8(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository8`: SimpleApiProxyRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository8`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiProxyRepository**](SimpleApiProxyRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository9

> SimpleApiGroupDeployRepository GetRepository9(ctx, repositoryName).Execute()

Get repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryManagementAPI.GetRepository9(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.GetRepository9``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository9`: SimpleApiGroupDeployRepository
	fmt.Fprintf(os.Stdout, "Response from `RepositoryManagementAPI.GetRepository9`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepository9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SimpleApiGroupDeployRepository**](SimpleApiGroupDeployRepository.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateCache

> InvalidateCache(ctx, repositoryName).Execute()

Invalidate repository cache. Proxy or group repositories only.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to invalidate cache

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.InvalidateCache(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.InvalidateCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to invalidate cache | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateCacheRequest struct via the builder pattern


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


## RebuildIndex

> RebuildIndex(ctx, repositoryName).Execute()

Schedule a 'Repair - Rebuild repository search' Task. Hosted or proxy repositories only.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to rebuild index

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.RebuildIndex(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.RebuildIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to rebuild index | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildIndexRequest struct via the builder pattern


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


## UpdateRepository

> UpdateRepository(ctx, repositoryName).Body(body).Execute()

Update Maven group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewMavenGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // MavenGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MavenGroupRepositoryApiRequest**](MavenGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository1

> UpdateRepository1(ctx, repositoryName).Body(body).Execute()

Update Maven hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewMavenHostedRepositoryApiRequest(*sonatyperepo.NewMavenAttributes(), "internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // MavenHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository1(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MavenHostedRepositoryApiRequest**](MavenHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository10

> UpdateRepository10(ctx, repositoryName).Body(body).Execute()

Update npm proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewNpmProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // NpmProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository10(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository10``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NpmProxyRepositoryApiRequest**](NpmProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository11

> UpdateRepository11(ctx, repositoryName).Body(body).Execute()

Update NuGet group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewNugetGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // NugetGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository11(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository11``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NugetGroupRepositoryApiRequest**](NugetGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository12

> UpdateRepository12(ctx, repositoryName).Body(body).Execute()

Update NuGet hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewNugetHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // NugetHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository12(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository12``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NugetHostedRepositoryApiRequest**](NugetHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository13

> UpdateRepository13(ctx, repositoryName).Body(body).Execute()

Update NuGet proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewNugetProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), *sonatyperepo.NewNugetAttributes(), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // NugetProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository13(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository13``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NugetProxyRepositoryApiRequest**](NugetProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository14

> UpdateRepository14(ctx, repositoryName).Body(body).Execute()

Update RubyGems group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRubyGemsGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // RubyGemsGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository14(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository14``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RubyGemsGroupRepositoryApiRequest**](RubyGemsGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository15

> UpdateRepository15(ctx, repositoryName).Body(body).Execute()

Update RubyGems hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRubyGemsHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // RubyGemsHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository15(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository15``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RubyGemsHostedRepositoryApiRequest**](RubyGemsHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository16

> UpdateRepository16(ctx, repositoryName).Body(body).Execute()

Update RubyGems proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRubyGemsProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // RubyGemsProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository16(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository16``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RubyGemsProxyRepositoryApiRequest**](RubyGemsProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository17

> UpdateRepository17(ctx, repositoryName).Body(body).Execute()

Update Yum group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewYumGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // YumGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository17(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository17``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**YumGroupRepositoryApiRequest**](YumGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository18

> UpdateRepository18(ctx, repositoryName).Body(body).Execute()

Update Yum hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewYumHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once"), *sonatyperepo.NewYumAttributes(int32(5))) // YumHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository18(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository18``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository18Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**YumHostedRepositoryApiRequest**](YumHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository19

> UpdateRepository19(ctx, repositoryName).Body(body).Execute()

Update Yum proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewYumProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // YumProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository19(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository19``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository19Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**YumProxyRepositoryApiRequest**](YumProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository2

> UpdateRepository2(ctx, repositoryName).Body(body).Execute()

Update Maven proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewMavenProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributesWithPreemptiveAuth(true, false), *sonatyperepo.NewMavenAttributes(), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // MavenProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository2(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MavenProxyRepositoryApiRequest**](MavenProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository20

> UpdateRepository20(ctx, repositoryName).Body(body).Execute()

Update Docker group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewDockerGroupRepositoryApiRequest(*sonatyperepo.NewDockerAttributes(true, false), *sonatyperepo.NewGroupDeployAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // DockerGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository20(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository20``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DockerGroupRepositoryApiRequest**](DockerGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository21

> UpdateRepository21(ctx, repositoryName).Body(body).Execute()

Update Docker hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewDockerHostedRepositoryApiRequest(*sonatyperepo.NewDockerAttributes(true, false), "internal", true, *sonatyperepo.NewDockerHostedStorageAttributes("default", true, "allow_once")) // DockerHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository21(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository21``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DockerHostedRepositoryApiRequest**](DockerHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository22

> UpdateRepository22(ctx, repositoryName).Body(body).Execute()

Update Docker proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewDockerProxyRepositoryApiRequest(*sonatyperepo.NewDockerAttributes(true, false), *sonatyperepo.NewDockerProxyAttributes(), *sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // DockerProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository22(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository22``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository22Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DockerProxyRepositoryApiRequest**](DockerProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository23

> UpdateRepository23(ctx, repositoryName).Body(body).Execute()

Update PyPI group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewPypiGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // PypiGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository23(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository23``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository23Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PypiGroupRepositoryApiRequest**](PypiGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository24

> UpdateRepository24(ctx, repositoryName).Body(body).Execute()

Update PyPI hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewPypiHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // PypiHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository24(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository24``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository24Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PypiHostedRepositoryApiRequest**](PypiHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository25

> UpdateRepository25(ctx, repositoryName).Body(body).Execute()

Update PyPI proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewPypiProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // PypiProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository25(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository25``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository25Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PypiProxyRepositoryApiRequest**](PypiProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository26

> UpdateRepository26(ctx, repositoryName).Body(body).Execute()

Update conda proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewCondaProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // CondaProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository26(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository26``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository26Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CondaProxyRepositoryApiRequest**](CondaProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository27

> UpdateRepository27(ctx, repositoryName).Body(body).Execute()

Update Conan proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewConanProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // ConanProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository27(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository27``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository27Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConanProxyRepositoryApiRequest**](ConanProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository28

> UpdateRepository28(ctx, repositoryName).Body(body).Execute()

Update Git LFS hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewGitLfsHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // GitLfsHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository28(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository28``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository28Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GitLfsHostedRepositoryApiRequest**](GitLfsHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository29

> UpdateRepository29(ctx, repositoryName).Body(body).Execute()

Update R group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // RGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository29(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository29``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository29Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RGroupRepositoryApiRequest**](RGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository3

> UpdateRepository3(ctx, repositoryName).Body(body).Execute()

Update APT hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewAptHostedRepositoryApiRequest(*sonatyperepo.NewAptHostedRepositoriesAttributes(), *sonatyperepo.NewAptSigningRepositoriesAttributes(), "internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // AptHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository3(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AptHostedRepositoryApiRequest**](AptHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository30

> UpdateRepository30(ctx, repositoryName).Body(body).Execute()

Update R hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // RHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository30(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository30``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository30Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RHostedRepositoryApiRequest**](RHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository31

> UpdateRepository31(ctx, repositoryName).Body(body).Execute()

Update R proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // RProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository31(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository31``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository31Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RProxyRepositoryApiRequest**](RProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository32

> UpdateRepository32(ctx, repositoryName).Body(body).Execute()

Update Cocoapods proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewCocoapodsProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // CocoapodsProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository32(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository32``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository32Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CocoapodsProxyRepositoryApiRequest**](CocoapodsProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository33

> UpdateRepository33(ctx, repositoryName).Body(body).Execute()

Update a Go group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewGolangGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // GolangGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository33(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository33``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository33Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GolangGroupRepositoryApiRequest**](GolangGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository34

> UpdateRepository34(ctx, repositoryName).Body(body).Execute()

Update a Go proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewGolangProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // GolangProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository34(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository34``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository34Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GolangProxyRepositoryApiRequest**](GolangProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository35

> UpdateRepository35(ctx, repositoryName).Body(body).Execute()

Update p2 proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewP2ProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // P2ProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository35(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository35``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository35Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**P2ProxyRepositoryApiRequest**](P2ProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository36

> UpdateRepository36(ctx, repositoryName).Body(body).Execute()

Update Helm hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewHelmHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // HelmHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository36(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository36``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository36Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HelmHostedRepositoryApiRequest**](HelmHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository37

> UpdateRepository37(ctx, repositoryName).Body(body).Execute()

Update Helm proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewHelmProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // HelmProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository37(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository37``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository37Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HelmProxyRepositoryApiRequest**](HelmProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository38

> UpdateRepository38(ctx, repositoryName).Body(body).Execute()

Update Bower group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewBowerGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // BowerGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository38(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository38``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository38Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BowerGroupRepositoryApiRequest**](BowerGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository39

> UpdateRepository39(ctx, repositoryName).Body(body).Execute()

Update Bower hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewBowerHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // BowerHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository39(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository39``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository39Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BowerHostedRepositoryApiRequest**](BowerHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository4

> UpdateRepository4(ctx, repositoryName).Body(body).Execute()

Update APT proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewAptProxyRepositoryApiRequest(*sonatyperepo.NewAptProxyRepositoriesAttributes(false), *sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // AptProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository4(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AptProxyRepositoryApiRequest**](AptProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository40

> UpdateRepository40(ctx, repositoryName).Body(body).Execute()

Update Bower proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewBowerProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // BowerProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository40(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository40``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository40Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**BowerProxyRepositoryApiRequest**](BowerProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository5

> UpdateRepository5(ctx, repositoryName).Body(body).Execute()

Update raw group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRawGroupRepositoryApiRequest(*sonatyperepo.NewGroupAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // RawGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository5(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RawGroupRepositoryApiRequest**](RawGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository6

> UpdateRepository6(ctx, repositoryName).Body(body).Execute()

Update raw hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRawHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // RawHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository6(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RawHostedRepositoryApiRequest**](RawHostedRepositoryApiRequest.md) |  | 

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


## UpdateRepository7

> UpdateRepository7(ctx, repositoryName).Body(body).Execute()

Update raw proxy repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewRawProxyRepositoryApiRequest(*sonatyperepo.NewHttpClientAttributes(true, false), "internal", *sonatyperepo.NewNegativeCacheAttributes(true, int32(1440)), true, *sonatyperepo.NewProxyAttributes(int32(1440), int32(1440)), *sonatyperepo.NewStorageAttributes("default", true)) // RawProxyRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository7(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RawProxyRepositoryApiRequest**](RawProxyRepositoryApiRequest.md) |  | 

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


## UpdateRepository8

> UpdateRepository8(ctx, repositoryName).Body(body).Execute()

Update npm group repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewNpmGroupRepositoryApiRequest(*sonatyperepo.NewGroupDeployAttributes(), "internal", true, *sonatyperepo.NewStorageAttributes("default", true)) // NpmGroupRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository8(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NpmGroupRepositoryApiRequest**](NpmGroupRepositoryApiRequest.md) |  | 

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


## UpdateRepository9

> UpdateRepository9(ctx, repositoryName).Body(body).Execute()

Update npm hosted repository

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
)

func main() {
	repositoryName := "repositoryName_example" // string | Name of the repository to update
	body := *sonatyperepo.NewNpmHostedRepositoryApiRequest("internal", true, *sonatyperepo.NewHostedStorageAttributes("default", true, "allow_once")) // NpmHostedRepositoryApiRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryManagementAPI.UpdateRepository9(context.Background(), repositoryName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryManagementAPI.UpdateRepository9``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Name of the repository to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepository9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**NpmHostedRepositoryApiRequest**](NpmHostedRepositoryApiRequest.md) |  | 

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


# \ComponentsAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponent**](ComponentsAPI.md#DeleteComponent) | **Delete** /v1/components/{id} | Delete a single component
[**GetComponentById**](ComponentsAPI.md#GetComponentById) | **Get** /v1/components/{id} | Get a single component
[**GetComponents**](ComponentsAPI.md#GetComponents) | **Get** /v1/components | List components
[**UploadComponent**](ComponentsAPI.md#UploadComponent) | **Post** /v1/components | Upload a single component



## DeleteComponent

> DeleteComponent(ctx, id).Execute()

Delete a single component

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
	id := "id_example" // string | ID of the component to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ComponentsAPI.DeleteComponent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.DeleteComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the component to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentRequest struct via the builder pattern


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


## GetComponentById

> ComponentXO GetComponentById(ctx, id).Execute()

Get a single component

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
	id := "id_example" // string | ID of the component to retrieve

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetComponentById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetComponentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentById`: ComponentXO
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetComponentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the component to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentXO**](ComponentXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponents

> PageComponentXO GetComponents(ctx).Repository(repository).ContinuationToken(continuationToken).Execute()

List components

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
	repository := "repository_example" // string | Repository from which you would like to retrieve components
	continuationToken := "continuationToken_example" // string | A token returned by a prior request. If present, the next page of results are returned (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetComponents(context.Background()).Repository(repository).ContinuationToken(continuationToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponents`: PageComponentXO
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetComponents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | **string** | Repository from which you would like to retrieve components | 
 **continuationToken** | **string** | A token returned by a prior request. If present, the next page of results are returned | 

### Return type

[**PageComponentXO**](PageComponentXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadComponent

> UploadComponent(ctx).Repository(repository).AptAsset(aptAsset).AptTag(aptTag).DockerAsset(dockerAsset).DockerTag(dockerTag).HelmAsset(helmAsset).HelmTag(helmTag).Maven2ArtifactId(maven2ArtifactId).Maven2Asset1(maven2Asset1).Maven2Asset1Classifier(maven2Asset1Classifier).Maven2Asset1Extension(maven2Asset1Extension).Maven2Asset2(maven2Asset2).Maven2Asset2Classifier(maven2Asset2Classifier).Maven2Asset2Extension(maven2Asset2Extension).Maven2Asset3(maven2Asset3).Maven2Asset3Classifier(maven2Asset3Classifier).Maven2Asset3Extension(maven2Asset3Extension).Maven2GeneratePom(maven2GeneratePom).Maven2GroupId(maven2GroupId).Maven2Packaging(maven2Packaging).Maven2Tag(maven2Tag).Maven2Version(maven2Version).NpmAsset(npmAsset).NpmTag(npmTag).NugetAsset(nugetAsset).NugetTag(nugetTag).PypiAsset(pypiAsset).PypiTag(pypiTag).RAsset(rAsset).RAssetPathId(rAssetPathId).RTag(rTag).RawAsset1(rawAsset1).RawAsset1Filename(rawAsset1Filename).RawAsset2(rawAsset2).RawAsset2Filename(rawAsset2Filename).RawAsset3(rawAsset3).RawAsset3Filename(rawAsset3Filename).RawDirectory(rawDirectory).RawTag(rawTag).RubygemsAsset(rubygemsAsset).RubygemsTag(rubygemsTag).YumAsset(yumAsset).YumAssetFilename(yumAssetFilename).YumDirectory(yumDirectory).YumTag(yumTag).Execute()

Upload a single component

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
	repository := "repository_example" // string | Name of the repository to which you would like to upload the component
	aptAsset := os.NewFile(1234, "some_file") // *os.File | apt Asset  (optional)
	aptTag := "aptTag_example" // string | apt Tag (optional)
	dockerAsset := os.NewFile(1234, "some_file") // *os.File | docker Asset  (optional)
	dockerTag := "dockerTag_example" // string | docker Tag (optional)
	helmAsset := os.NewFile(1234, "some_file") // *os.File | helm Asset  (optional)
	helmTag := "helmTag_example" // string | helm Tag (optional)
	maven2ArtifactId := "maven2ArtifactId_example" // string | maven2 Artifact ID (optional)
	maven2Asset1 := os.NewFile(1234, "some_file") // *os.File | maven2 Asset 1 (optional)
	maven2Asset1Classifier := "maven2Asset1Classifier_example" // string | maven2 Asset 1 Classifier (optional)
	maven2Asset1Extension := "maven2Asset1Extension_example" // string | maven2 Asset 1 Extension (optional)
	maven2Asset2 := os.NewFile(1234, "some_file") // *os.File | maven2 Asset 2 (optional)
	maven2Asset2Classifier := "maven2Asset2Classifier_example" // string | maven2 Asset 2 Classifier (optional)
	maven2Asset2Extension := "maven2Asset2Extension_example" // string | maven2 Asset 2 Extension (optional)
	maven2Asset3 := os.NewFile(1234, "some_file") // *os.File | maven2 Asset 3 (optional)
	maven2Asset3Classifier := "maven2Asset3Classifier_example" // string | maven2 Asset 3 Classifier (optional)
	maven2Asset3Extension := "maven2Asset3Extension_example" // string | maven2 Asset 3 Extension (optional)
	maven2GeneratePom := true // bool | maven2 Generate a POM file with these coordinates (optional)
	maven2GroupId := "maven2GroupId_example" // string | maven2 Group ID (optional)
	maven2Packaging := "maven2Packaging_example" // string | maven2 Packaging (optional)
	maven2Tag := "maven2Tag_example" // string | maven2 Tag (optional)
	maven2Version := "maven2Version_example" // string | maven2 Version (optional)
	npmAsset := os.NewFile(1234, "some_file") // *os.File | npm Asset  (optional)
	npmTag := "npmTag_example" // string | npm Tag (optional)
	nugetAsset := os.NewFile(1234, "some_file") // *os.File | nuget Asset  (optional)
	nugetTag := "nugetTag_example" // string | nuget Tag (optional)
	pypiAsset := os.NewFile(1234, "some_file") // *os.File | pypi Asset  (optional)
	pypiTag := "pypiTag_example" // string | pypi Tag (optional)
	rAsset := os.NewFile(1234, "some_file") // *os.File | r Asset  (optional)
	rAssetPathId := "rAssetPathId_example" // string | r Asset  Package Path (optional)
	rTag := "rTag_example" // string | r Tag (optional)
	rawAsset1 := os.NewFile(1234, "some_file") // *os.File | raw Asset 1 (optional)
	rawAsset1Filename := "rawAsset1Filename_example" // string | raw Asset 1 Filename (optional)
	rawAsset2 := os.NewFile(1234, "some_file") // *os.File | raw Asset 2 (optional)
	rawAsset2Filename := "rawAsset2Filename_example" // string | raw Asset 2 Filename (optional)
	rawAsset3 := os.NewFile(1234, "some_file") // *os.File | raw Asset 3 (optional)
	rawAsset3Filename := "rawAsset3Filename_example" // string | raw Asset 3 Filename (optional)
	rawDirectory := "rawDirectory_example" // string | raw Directory (optional)
	rawTag := "rawTag_example" // string | raw Tag (optional)
	rubygemsAsset := os.NewFile(1234, "some_file") // *os.File | rubygems Asset  (optional)
	rubygemsTag := "rubygemsTag_example" // string | rubygems Tag (optional)
	yumAsset := os.NewFile(1234, "some_file") // *os.File | yum Asset  (optional)
	yumAssetFilename := "yumAssetFilename_example" // string | yum Asset  Filename (optional)
	yumDirectory := "yumDirectory_example" // string | yum Directory (optional)
	yumTag := "yumTag_example" // string | yum Tag (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ComponentsAPI.UploadComponent(context.Background()).Repository(repository).AptAsset(aptAsset).AptTag(aptTag).DockerAsset(dockerAsset).DockerTag(dockerTag).HelmAsset(helmAsset).HelmTag(helmTag).Maven2ArtifactId(maven2ArtifactId).Maven2Asset1(maven2Asset1).Maven2Asset1Classifier(maven2Asset1Classifier).Maven2Asset1Extension(maven2Asset1Extension).Maven2Asset2(maven2Asset2).Maven2Asset2Classifier(maven2Asset2Classifier).Maven2Asset2Extension(maven2Asset2Extension).Maven2Asset3(maven2Asset3).Maven2Asset3Classifier(maven2Asset3Classifier).Maven2Asset3Extension(maven2Asset3Extension).Maven2GeneratePom(maven2GeneratePom).Maven2GroupId(maven2GroupId).Maven2Packaging(maven2Packaging).Maven2Tag(maven2Tag).Maven2Version(maven2Version).NpmAsset(npmAsset).NpmTag(npmTag).NugetAsset(nugetAsset).NugetTag(nugetTag).PypiAsset(pypiAsset).PypiTag(pypiTag).RAsset(rAsset).RAssetPathId(rAssetPathId).RTag(rTag).RawAsset1(rawAsset1).RawAsset1Filename(rawAsset1Filename).RawAsset2(rawAsset2).RawAsset2Filename(rawAsset2Filename).RawAsset3(rawAsset3).RawAsset3Filename(rawAsset3Filename).RawDirectory(rawDirectory).RawTag(rawTag).RubygemsAsset(rubygemsAsset).RubygemsTag(rubygemsTag).YumAsset(yumAsset).YumAssetFilename(yumAssetFilename).YumDirectory(yumDirectory).YumTag(yumTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.UploadComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **repository** | **string** | Name of the repository to which you would like to upload the component | 
 **aptAsset** | ***os.File** | apt Asset  | 
 **aptTag** | **string** | apt Tag | 
 **dockerAsset** | ***os.File** | docker Asset  | 
 **dockerTag** | **string** | docker Tag | 
 **helmAsset** | ***os.File** | helm Asset  | 
 **helmTag** | **string** | helm Tag | 
 **maven2ArtifactId** | **string** | maven2 Artifact ID | 
 **maven2Asset1** | ***os.File** | maven2 Asset 1 | 
 **maven2Asset1Classifier** | **string** | maven2 Asset 1 Classifier | 
 **maven2Asset1Extension** | **string** | maven2 Asset 1 Extension | 
 **maven2Asset2** | ***os.File** | maven2 Asset 2 | 
 **maven2Asset2Classifier** | **string** | maven2 Asset 2 Classifier | 
 **maven2Asset2Extension** | **string** | maven2 Asset 2 Extension | 
 **maven2Asset3** | ***os.File** | maven2 Asset 3 | 
 **maven2Asset3Classifier** | **string** | maven2 Asset 3 Classifier | 
 **maven2Asset3Extension** | **string** | maven2 Asset 3 Extension | 
 **maven2GeneratePom** | **bool** | maven2 Generate a POM file with these coordinates | 
 **maven2GroupId** | **string** | maven2 Group ID | 
 **maven2Packaging** | **string** | maven2 Packaging | 
 **maven2Tag** | **string** | maven2 Tag | 
 **maven2Version** | **string** | maven2 Version | 
 **npmAsset** | ***os.File** | npm Asset  | 
 **npmTag** | **string** | npm Tag | 
 **nugetAsset** | ***os.File** | nuget Asset  | 
 **nugetTag** | **string** | nuget Tag | 
 **pypiAsset** | ***os.File** | pypi Asset  | 
 **pypiTag** | **string** | pypi Tag | 
 **rAsset** | ***os.File** | r Asset  | 
 **rAssetPathId** | **string** | r Asset  Package Path | 
 **rTag** | **string** | r Tag | 
 **rawAsset1** | ***os.File** | raw Asset 1 | 
 **rawAsset1Filename** | **string** | raw Asset 1 Filename | 
 **rawAsset2** | ***os.File** | raw Asset 2 | 
 **rawAsset2Filename** | **string** | raw Asset 2 Filename | 
 **rawAsset3** | ***os.File** | raw Asset 3 | 
 **rawAsset3Filename** | **string** | raw Asset 3 Filename | 
 **rawDirectory** | **string** | raw Directory | 
 **rawTag** | **string** | raw Tag | 
 **rubygemsAsset** | ***os.File** | rubygems Asset  | 
 **rubygemsTag** | **string** | rubygems Tag | 
 **yumAsset** | ***os.File** | yum Asset  | 
 **yumAssetFilename** | **string** | yum Asset  Filename | 
 **yumDirectory** | **string** | yum Directory | 
 **yumTag** | **string** | yum Tag | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


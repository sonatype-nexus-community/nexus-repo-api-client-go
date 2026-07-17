# \BlobStoreAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlobstoresAzure**](BlobStoreAPI.md#CreateBlobstoresAzure) | **Post** /v1/blobstores/azure | Create an Azure blob store
[**CreateBlobstoresFile**](BlobStoreAPI.md#CreateBlobstoresFile) | **Post** /v1/blobstores/file | Create a file blob store
[**CreateBlobstoresGoogle**](BlobStoreAPI.md#CreateBlobstoresGoogle) | **Post** /v1/blobstores/google | Create a Google Cloud blob store
[**CreateBlobstoresGroup**](BlobStoreAPI.md#CreateBlobstoresGroup) | **Post** /v1/blobstores/group | Create a group blob store
[**CreateBlobstoresGroupConvert**](BlobStoreAPI.md#CreateBlobstoresGroupConvert) | **Post** /v1/blobstores/group/convert/{name}/{newNameForOriginal} | Convert a blob store to a group blob store
[**CreateS3BlobStore**](BlobStoreAPI.md#CreateS3BlobStore) | **Post** /v1/blobstores/s3 | Create an S3 blob store
[**DeleteBlobstores**](BlobStoreAPI.md#DeleteBlobstores) | **Delete** /v1/blobstores/{name} | Delete a blob store by name
[**GetBlobstoresAzure**](BlobStoreAPI.md#GetBlobstoresAzure) | **Get** /v1/blobstores/azure/{name} | Get an Azure blob store configuration by name
[**GetBlobstoresFile**](BlobStoreAPI.md#GetBlobstoresFile) | **Get** /v1/blobstores/file/{name} | Get a file blob store configuration by name
[**GetBlobstoresGoogle**](BlobStoreAPI.md#GetBlobstoresGoogle) | **Get** /v1/blobstores/google/{name} | Get the configuration for a Google Cloud blob store
[**GetBlobstoresGoogleRegions**](BlobStoreAPI.md#GetBlobstoresGoogleRegions) | **Get** /v1/blobstores/google/regions/{projectId} | Get the project regions by project&#39;s id
[**GetBlobstoresGroup**](BlobStoreAPI.md#GetBlobstoresGroup) | **Get** /v1/blobstores/group/{name} | Get a group blob store configuration by name
[**GetBlobstoresQuotaStatus**](BlobStoreAPI.md#GetBlobstoresQuotaStatus) | **Get** /v1/blobstores/{name}/quota-status | Get quota status for a given blob store
[**GetS3BlobStore**](BlobStoreAPI.md#GetS3BlobStore) | **Get** /v1/blobstores/s3/{name} | Get a S3 blob store configuration by name
[**ListBlobstores**](BlobStoreAPI.md#ListBlobstores) | **Get** /v1/blobstores | List the blob stores
[**UpdateBlobstoresAzure**](BlobStoreAPI.md#UpdateBlobstoresAzure) | **Put** /v1/blobstores/azure/{name} | Update an Azure blob store configuration by name
[**UpdateBlobstoresFile**](BlobStoreAPI.md#UpdateBlobstoresFile) | **Put** /v1/blobstores/file/{name} | Update a file blob store configuration by name
[**UpdateBlobstoresGoogle**](BlobStoreAPI.md#UpdateBlobstoresGoogle) | **Put** /v1/blobstores/google/{name} | Update a Google Cloud blob store
[**UpdateBlobstoresGroup**](BlobStoreAPI.md#UpdateBlobstoresGroup) | **Put** /v1/blobstores/group/{name} | Update a group blob store configuration by name
[**UpdateS3BlobStore**](BlobStoreAPI.md#UpdateS3BlobStore) | **Put** /v1/blobstores/s3/{name} | Update an S3 blob store configuration by name



## CreateBlobstoresAzure

> CreateBlobstoresAzure(ctx).AzureBlobStoreApiModel(azureBlobStoreApiModel).Execute()

Create an Azure blob store

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
	azureBlobStoreApiModel := *sonatyperepo.NewAzureBlobStoreApiModel(*sonatyperepo.NewAzureBlobStoreApiBucketConfiguration("AccountName_example", *sonatyperepo.NewAzureBlobStoreApiAuthentication("AuthenticationMethod_example"), "ContainerName_example"), "Name_example") // AzureBlobStoreApiModel |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateBlobstoresAzure(context.Background()).AzureBlobStoreApiModel(azureBlobStoreApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobstoresAzure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobstoresAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureBlobStoreApiModel** | [**AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md) |  | 

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


## CreateBlobstoresFile

> CreateBlobstoresFile(ctx).FileBlobStoreApiCreateRequest(fileBlobStoreApiCreateRequest).Execute()

Create a file blob store

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
	fileBlobStoreApiCreateRequest := *sonatyperepo.NewFileBlobStoreApiCreateRequest("Name_example", "Path_example") // FileBlobStoreApiCreateRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateBlobstoresFile(context.Background()).FileBlobStoreApiCreateRequest(fileBlobStoreApiCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobstoresFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobstoresFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileBlobStoreApiCreateRequest** | [**FileBlobStoreApiCreateRequest**](FileBlobStoreApiCreateRequest.md) |  | 

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


## CreateBlobstoresGoogle

> CreateBlobstoresGoogle(ctx).GoogleCloudBlobstoreApiModel(googleCloudBlobstoreApiModel).Execute()

Create a Google Cloud blob store

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
	googleCloudBlobstoreApiModel := *sonatyperepo.NewGoogleCloudBlobstoreApiModel(*sonatyperepo.NewGoogleCloudBlobStoreApiBucketConfiguration(*sonatyperepo.NewGoogleCloudBlobStoreApiBucket("Name_example")), "gc_storage") // GoogleCloudBlobstoreApiModel |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateBlobstoresGoogle(context.Background()).GoogleCloudBlobstoreApiModel(googleCloudBlobstoreApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobstoresGoogle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobstoresGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **googleCloudBlobstoreApiModel** | [**GoogleCloudBlobstoreApiModel**](GoogleCloudBlobstoreApiModel.md) |  | 

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


## CreateBlobstoresGroup

> CreateBlobstoresGroup(ctx).GroupBlobStoreApiCreateRequest(groupBlobStoreApiCreateRequest).Execute()

Create a group blob store

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
	groupBlobStoreApiCreateRequest := *sonatyperepo.NewGroupBlobStoreApiCreateRequest("Name_example") // GroupBlobStoreApiCreateRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateBlobstoresGroup(context.Background()).GroupBlobStoreApiCreateRequest(groupBlobStoreApiCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobstoresGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobstoresGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupBlobStoreApiCreateRequest** | [**GroupBlobStoreApiCreateRequest**](GroupBlobStoreApiCreateRequest.md) |  | 

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


## CreateBlobstoresGroupConvert

> GroupBlobStoreApiModel CreateBlobstoresGroupConvert(ctx, name, newNameForOriginal).Execute()

Convert a blob store to a group blob store

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
	name := "name_example" // string | The name of the group blob store
	newNameForOriginal := "newNameForOriginal_example" // string | A new name to the original blob store

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.CreateBlobstoresGroupConvert(context.Background(), name, newNameForOriginal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateBlobstoresGroupConvert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlobstoresGroupConvert`: GroupBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.CreateBlobstoresGroupConvert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the group blob store | 
**newNameForOriginal** | **string** | A new name to the original blob store | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlobstoresGroupConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupBlobStoreApiModel**](GroupBlobStoreApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateS3BlobStore

> CreateS3BlobStore(ctx).S3BlobStoreApiModel(s3BlobStoreApiModel).Execute()

Create an S3 blob store

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
	s3BlobStoreApiModel := *sonatyperepo.NewS3BlobStoreApiModel(*sonatyperepo.NewS3BlobStoreApiBucketConfiguration(*sonatyperepo.NewS3BlobStoreApiBucket("Name_example", "DEFAULT")), "s3") // S3BlobStoreApiModel |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.CreateS3BlobStore(context.Background()).S3BlobStoreApiModel(s3BlobStoreApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.CreateS3BlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateS3BlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **s3BlobStoreApiModel** | [**S3BlobStoreApiModel**](S3BlobStoreApiModel.md) |  | 

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


## DeleteBlobstores

> DeleteBlobstores(ctx, name).Execute()

Delete a blob store by name

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
	name := "name_example" // string | The name of the blob store to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.DeleteBlobstores(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.DeleteBlobstores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the blob store to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlobstoresRequest struct via the builder pattern


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


## GetBlobstoresAzure

> AzureBlobStoreApiModel GetBlobstoresAzure(ctx, name).Execute()

Get an Azure blob store configuration by name

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
	name := "name_example" // string | Name of the blob store configuration to fetch

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobstoresAzure(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobstoresAzure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobstoresAzure`: AzureBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobstoresAzure`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store configuration to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobstoresAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobstoresFile

> FileBlobStoreApiModel GetBlobstoresFile(ctx, name).Execute()

Get a file blob store configuration by name

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
	name := "default" // string | The name of the file blob store to read

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobstoresFile(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobstoresFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobstoresFile`: FileBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobstoresFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the file blob store to read | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobstoresFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileBlobStoreApiModel**](FileBlobStoreApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobstoresGoogle

> GoogleCloudBlobstoreApiModel GetBlobstoresGoogle(ctx, name).Execute()

Get the configuration for a Google Cloud blob store

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
	name := "name_example" // string | the name of the blob store

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobstoresGoogle(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobstoresGoogle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobstoresGoogle`: GoogleCloudBlobstoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobstoresGoogle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the name of the blob store | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobstoresGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GoogleCloudBlobstoreApiModel**](GoogleCloudBlobstoreApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobstoresGoogleRegions

> []string GetBlobstoresGoogleRegions(ctx, projectId).Execute()

Get the project regions by project's id

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
	projectId := "projectId_example" // string | projectId

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobstoresGoogleRegions(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobstoresGoogleRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobstoresGoogleRegions`: []string
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobstoresGoogleRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | projectId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobstoresGoogleRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobstoresGroup

> GroupBlobStoreApiModel GetBlobstoresGroup(ctx, name).Execute()

Get a group blob store configuration by name

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
	name := "name_example" // string | The name of the group blob store

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetBlobstoresGroup(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobstoresGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobstoresGroup`: GroupBlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobstoresGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the group blob store | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobstoresGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupBlobStoreApiModel**](GroupBlobStoreApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlobstoresQuotaStatus

> BlobStoreQuotaResultXO GetBlobstoresQuotaStatus(ctx, name).Execute()

Get quota status for a given blob store

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
	resp, r, err := apiClient.BlobStoreAPI.GetBlobstoresQuotaStatus(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetBlobstoresQuotaStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlobstoresQuotaStatus`: BlobStoreQuotaResultXO
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetBlobstoresQuotaStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlobstoresQuotaStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlobStoreQuotaResultXO**](BlobStoreQuotaResultXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetS3BlobStore

> S3BlobStoreApiModel GetS3BlobStore(ctx, name).Execute()

Get a S3 blob store configuration by name

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
	name := "name_example" // string | Name of the blob store configuration to fetch

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.BlobStoreAPI.GetS3BlobStore(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.GetS3BlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetS3BlobStore`: S3BlobStoreApiModel
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.GetS3BlobStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store configuration to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetS3BlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**S3BlobStoreApiModel**](S3BlobStoreApiModel.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBlobstores

> []GenericBlobStoreApiResponse ListBlobstores(ctx).Execute()

List the blob stores

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
	resp, r, err := apiClient.BlobStoreAPI.ListBlobstores(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.ListBlobstores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBlobstores`: []GenericBlobStoreApiResponse
	fmt.Fprintf(os.Stdout, "Response from `BlobStoreAPI.ListBlobstores`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBlobstoresRequest struct via the builder pattern


### Return type

[**[]GenericBlobStoreApiResponse**](GenericBlobStoreApiResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBlobstoresAzure

> UpdateBlobstoresAzure(ctx, name).AzureBlobStoreApiModel(azureBlobStoreApiModel).Execute()

Update an Azure blob store configuration by name

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
	name := "name_example" // string | Name of the blob store to update
	azureBlobStoreApiModel := *sonatyperepo.NewAzureBlobStoreApiModel(*sonatyperepo.NewAzureBlobStoreApiBucketConfiguration("AccountName_example", *sonatyperepo.NewAzureBlobStoreApiAuthentication("AuthenticationMethod_example"), "ContainerName_example"), "Name_example") // AzureBlobStoreApiModel |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateBlobstoresAzure(context.Background(), name).AzureBlobStoreApiModel(azureBlobStoreApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateBlobstoresAzure``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlobstoresAzureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureBlobStoreApiModel** | [**AzureBlobStoreApiModel**](AzureBlobStoreApiModel.md) |  | 

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


## UpdateBlobstoresFile

> UpdateBlobstoresFile(ctx, name).FileBlobStoreApiUpdateRequest(fileBlobStoreApiUpdateRequest).Execute()

Update a file blob store configuration by name

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
	name := "name_example" // string | The name of the file blob store to update
	fileBlobStoreApiUpdateRequest := *sonatyperepo.NewFileBlobStoreApiUpdateRequest("Path_example") // FileBlobStoreApiUpdateRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateBlobstoresFile(context.Background(), name).FileBlobStoreApiUpdateRequest(fileBlobStoreApiUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateBlobstoresFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the file blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlobstoresFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileBlobStoreApiUpdateRequest** | [**FileBlobStoreApiUpdateRequest**](FileBlobStoreApiUpdateRequest.md) |  | 

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


## UpdateBlobstoresGoogle

> UpdateBlobstoresGoogle(ctx, name).GoogleCloudBlobstoreApiModel(googleCloudBlobstoreApiModel).Execute()

Update a Google Cloud blob store

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
	name := "name_example" // string | the name of the blobstore
	googleCloudBlobstoreApiModel := *sonatyperepo.NewGoogleCloudBlobstoreApiModel(*sonatyperepo.NewGoogleCloudBlobStoreApiBucketConfiguration(*sonatyperepo.NewGoogleCloudBlobStoreApiBucket("Name_example")), "gc_storage") // GoogleCloudBlobstoreApiModel |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateBlobstoresGoogle(context.Background(), name).GoogleCloudBlobstoreApiModel(googleCloudBlobstoreApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateBlobstoresGoogle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | the name of the blobstore | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlobstoresGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **googleCloudBlobstoreApiModel** | [**GoogleCloudBlobstoreApiModel**](GoogleCloudBlobstoreApiModel.md) |  | 

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


## UpdateBlobstoresGroup

> UpdateBlobstoresGroup(ctx, name).GroupBlobStoreApiUpdateRequest(groupBlobStoreApiUpdateRequest).Execute()

Update a group blob store configuration by name

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
	name := "name_example" // string | The name of the blob store to update
	groupBlobStoreApiUpdateRequest := *sonatyperepo.NewGroupBlobStoreApiUpdateRequest() // GroupBlobStoreApiUpdateRequest |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateBlobstoresGroup(context.Background(), name).GroupBlobStoreApiUpdateRequest(groupBlobStoreApiUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateBlobstoresGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | The name of the blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBlobstoresGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBlobStoreApiUpdateRequest** | [**GroupBlobStoreApiUpdateRequest**](GroupBlobStoreApiUpdateRequest.md) |  | 

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


## UpdateS3BlobStore

> UpdateS3BlobStore(ctx, name).S3BlobStoreApiModel(s3BlobStoreApiModel).Execute()

Update an S3 blob store configuration by name

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
	name := "name_example" // string | Name of the blob store to update
	s3BlobStoreApiModel := *sonatyperepo.NewS3BlobStoreApiModel(*sonatyperepo.NewS3BlobStoreApiBucketConfiguration(*sonatyperepo.NewS3BlobStoreApiBucket("Name_example", "DEFAULT")), "s3") // S3BlobStoreApiModel |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.BlobStoreAPI.UpdateS3BlobStore(context.Background(), name).S3BlobStoreApiModel(s3BlobStoreApiModel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlobStoreAPI.UpdateS3BlobStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the blob store to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateS3BlobStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3BlobStoreApiModel** | [**S3BlobStoreApiModel**](S3BlobStoreApiModel.md) |  | 

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


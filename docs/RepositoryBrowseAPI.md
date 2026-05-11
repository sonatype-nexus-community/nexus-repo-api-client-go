# \RepositoryBrowseAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFolder**](RepositoryBrowseAPI.md#DeleteFolder) | **Delete** /v1/repositories/{repositoryName}/browse | Delete a folder and all its contents
[**GetrepositorynameBrowseRepository**](RepositoryBrowseAPI.md#GetrepositorynameBrowseRepository) | **Get** /v1/repositories/{repositoryName}/browse | List browse nodes for a repository path



## DeleteFolder

> DeleteFolder(ctx, repositoryName).Path(path).Execute()

Delete a folder and all its contents

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
	repositoryName := "repositoryName_example" // string | Repository name
	path := "path_example" // string | Folder path to delete (URL-encoded)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.RepositoryBrowseAPI.DeleteFolder(context.Background(), repositoryName).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryBrowseAPI.DeleteFolder``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Folder path to delete (URL-encoded) | 

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


## GetrepositorynameBrowseRepository

> []BrowseNodeXO GetrepositorynameBrowseRepository(ctx, repositoryName).Path(path).Execute()

List browse nodes for a repository path

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
	repositoryName := "repositoryName_example" // string | Repository name
	path := "path_example" // string | Path within the repository (URL-encoded, use '/' as separator). Use '/' for root. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoryBrowseAPI.GetrepositorynameBrowseRepository(context.Background(), repositoryName).Path(path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryBrowseAPI.GetrepositorynameBrowseRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetrepositorynameBrowseRepository`: []BrowseNodeXO
	fmt.Fprintf(os.Stdout, "Response from `RepositoryBrowseAPI.GetrepositorynameBrowseRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Repository name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetrepositorynameBrowseRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **path** | **string** | Path within the repository (URL-encoded, use &#39;/&#39; as separator). Use &#39;/&#39; for root. | 

### Return type

[**[]BrowseNodeXO**](BrowseNodeXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


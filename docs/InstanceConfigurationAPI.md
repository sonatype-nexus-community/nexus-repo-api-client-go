# \InstanceConfigurationAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyConfiguration**](InstanceConfigurationAPI.md#ApplyConfiguration) | **Put** /v1/configuration | Apply instance configuration
[**ExportConfiguration**](InstanceConfigurationAPI.md#ExportConfiguration) | **Get** /v1/configuration | Export instance configuration
[**GetAssets1**](InstanceConfigurationAPI.md#GetAssets1) | **Get** /v1/configuration/assets | List assets
[**ImportAsset**](InstanceConfigurationAPI.md#ImportAsset) | **Post** /v1/configuration/assets/{repositoryName}/import | Import assets to repository



## ApplyConfiguration

> ConfigurationApplyResponseXO ApplyConfiguration(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Body(body).Execute()

Apply instance configuration

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
	xNexusMigrationProtocolVersion := "xNexusMigrationProtocolVersion_example" // string | Migration protocol version (must be '1')
	body := *sonatyperepo.NewInstanceConfigurationXO() // InstanceConfigurationXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.ApplyConfiguration(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.ApplyConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApplyConfiguration`: ConfigurationApplyResponseXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.ApplyConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplyConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 
 **body** | [**InstanceConfigurationXO**](InstanceConfigurationXO.md) |  | 

### Return type

[**ConfigurationApplyResponseXO**](ConfigurationApplyResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportConfiguration

> InstanceConfigurationXO ExportConfiguration(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()

Export instance configuration

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
	xNexusMigrationProtocolVersion := "xNexusMigrationProtocolVersion_example" // string | Migration protocol version (must be '1')

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.ExportConfiguration(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.ExportConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportConfiguration`: InstanceConfigurationXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.ExportConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 

### Return type

[**InstanceConfigurationXO**](InstanceConfigurationXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssets1

> Page GetAssets1(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Repository(repository).ContinuationToken(continuationToken).NewerThan(newerThan).OlderThan(olderThan).Execute()

List assets

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
	xNexusMigrationProtocolVersion := "xNexusMigrationProtocolVersion_example" // string | Migration protocol version (must be '1')
	repository := "repository_example" // string | Repository from which you would like to retrieve assets.
	continuationToken := "continuationToken_example" // string | A token returned by a prior request. If present, the next page of results are returned (optional)
	newerThan := int64(789) // int64 | Unix timestamp (milliseconds since epoch). Return assets last updated after this timestamp (optional)
	olderThan := int64(789) // int64 | Unix timestamp (milliseconds since epoch). Return assets last updated before this timestamp (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.GetAssets1(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Repository(repository).ContinuationToken(continuationToken).NewerThan(newerThan).OlderThan(olderThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.GetAssets1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssets1`: Page
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.GetAssets1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAssets1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 
 **repository** | **string** | Repository from which you would like to retrieve assets. | 
 **continuationToken** | **string** | A token returned by a prior request. If present, the next page of results are returned | 
 **newerThan** | **int64** | Unix timestamp (milliseconds since epoch). Return assets last updated after this timestamp | 
 **olderThan** | **int64** | Unix timestamp (milliseconds since epoch). Return assets last updated before this timestamp | 

### Return type

[**Page**](Page.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportAsset

> []AssetImportResponseXO ImportAsset(ctx, repositoryName).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()

Import assets to repository

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
	xNexusMigrationProtocolVersion := "xNexusMigrationProtocolVersion_example" // string | Migration protocol version (must be '1')
	repositoryName := "repositoryName_example" // string | Repository name to import assets to

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.ImportAsset(context.Background(), repositoryName).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.ImportAsset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportAsset`: []AssetImportResponseXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.ImportAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Repository name to import assets to | 

### Other Parameters

Other parameters are passed through a pointer to a apiImportAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 


### Return type

[**[]AssetImportResponseXO**](AssetImportResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


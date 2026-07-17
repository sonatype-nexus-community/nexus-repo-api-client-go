# \InstanceConfigurationAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConfigurationAssetsImport**](InstanceConfigurationAPI.md#CreateConfigurationAssetsImport) | **Post** /v1/configuration/assets/{repositoryName}/import | Import assets to repository
[**CreateConfigurationCipher**](InstanceConfigurationAPI.md#CreateConfigurationCipher) | **Post** /v1/configuration/cipher | Set migration cipher password for decrypting imported secrets
[**CreateConfigurationTransferComplete**](InstanceConfigurationAPI.md#CreateConfigurationTransferComplete) | **Post** /v1/configuration/transfer-complete | Receive transfer complete notification from migrator
[**DeleteConfigurationCipher**](InstanceConfigurationAPI.md#DeleteConfigurationCipher) | **Delete** /v1/configuration/cipher | Clear migration cipher password
[**ListConfiguration**](InstanceConfigurationAPI.md#ListConfiguration) | **Get** /v1/configuration | Export instance configuration
[**ListConfigurationAssets**](InstanceConfigurationAPI.md#ListConfigurationAssets) | **Get** /v1/configuration/assets | List assets
[**UpdateConfiguration**](InstanceConfigurationAPI.md#UpdateConfiguration) | **Put** /v1/configuration | Apply instance configuration



## CreateConfigurationAssetsImport

> AssetImportResponseXO CreateConfigurationAssetsImport(ctx, repositoryName).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()

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
	resp, r, err := apiClient.InstanceConfigurationAPI.CreateConfigurationAssetsImport(context.Background(), repositoryName).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.CreateConfigurationAssetsImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfigurationAssetsImport`: AssetImportResponseXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.CreateConfigurationAssetsImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** | Repository name to import assets to | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationAssetsImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 


### Return type

[**AssetImportResponseXO**](AssetImportResponseXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfigurationCipher

> string CreateConfigurationCipher(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Body(body).Execute()

Set migration cipher password for decrypting imported secrets

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
	body := "body_example" // string | Cipher password as plain text

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.CreateConfigurationCipher(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.CreateConfigurationCipher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfigurationCipher`: string
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.CreateConfigurationCipher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationCipherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 
 **body** | **string** | Cipher password as plain text | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConfigurationTransferComplete

> string CreateConfigurationTransferComplete(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).TransferCompleteXO(transferCompleteXO).Execute()

Receive transfer complete notification from migrator

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
	transferCompleteXO := *sonatyperepo.NewTransferCompleteXO() // TransferCompleteXO | Transfer completion notification data

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.CreateConfigurationTransferComplete(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).TransferCompleteXO(transferCompleteXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.CreateConfigurationTransferComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConfigurationTransferComplete`: string
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.CreateConfigurationTransferComplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConfigurationTransferCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 
 **transferCompleteXO** | [**TransferCompleteXO**](TransferCompleteXO.md) | Transfer completion notification data | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConfigurationCipher

> string DeleteConfigurationCipher(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()

Clear migration cipher password

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
	resp, r, err := apiClient.InstanceConfigurationAPI.DeleteConfigurationCipher(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.DeleteConfigurationCipher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteConfigurationCipher`: string
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.DeleteConfigurationCipher`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationCipherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConfiguration

> InstanceConfigurationXO ListConfiguration(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()

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
	resp, r, err := apiClient.InstanceConfigurationAPI.ListConfiguration(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.ListConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfiguration`: InstanceConfigurationXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.ListConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConfigurationRequest struct via the builder pattern


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


## ListConfigurationAssets

> Page ListConfigurationAssets(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Repository(repository).ContinuationToken(continuationToken).NewerThan(newerThan).OlderThan(olderThan).Execute()

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
	resp, r, err := apiClient.InstanceConfigurationAPI.ListConfigurationAssets(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).Repository(repository).ContinuationToken(continuationToken).NewerThan(newerThan).OlderThan(olderThan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.ListConfigurationAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListConfigurationAssets`: Page
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.ListConfigurationAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConfigurationAssetsRequest struct via the builder pattern


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


## UpdateConfiguration

> ConfigurationApplyResponseXO UpdateConfiguration(ctx).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).InstanceConfigurationXO(instanceConfigurationXO).Execute()

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
	instanceConfigurationXO := *sonatyperepo.NewInstanceConfigurationXO() // InstanceConfigurationXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceConfigurationAPI.UpdateConfiguration(context.Background()).XNexusMigrationProtocolVersion(xNexusMigrationProtocolVersion).InstanceConfigurationXO(instanceConfigurationXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceConfigurationAPI.UpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfiguration`: ConfigurationApplyResponseXO
	fmt.Fprintf(os.Stdout, "Response from `InstanceConfigurationAPI.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xNexusMigrationProtocolVersion** | **string** | Migration protocol version (must be &#39;1&#39;) | 
 **instanceConfigurationXO** | [**InstanceConfigurationXO**](InstanceConfigurationXO.md) |  | 

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


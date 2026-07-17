# \SecurityIPAllowListAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityIpAllowlistEntries**](SecurityIPAllowListAPI.md#CreateSecurityIpAllowlistEntries) | **Post** /v1/security/ip-allowlist/entries | Add a new entry to the IP Allow List
[**CreateSecurityIpAllowlistEntriesBulk**](SecurityIPAllowListAPI.md#CreateSecurityIpAllowlistEntriesBulk) | **Post** /v1/security/ip-allowlist/entries/bulk | Bulk upload entries from CSV content
[**DeleteSecurityIpAllowlistEntries**](SecurityIPAllowListAPI.md#DeleteSecurityIpAllowlistEntries) | **Delete** /v1/security/ip-allowlist/entries | Clear all entries from the IP Allow List
[**DeleteSecurityIpAllowlistEntriesBulk**](SecurityIPAllowListAPI.md#DeleteSecurityIpAllowlistEntriesBulk) | **Delete** /v1/security/ip-allowlist/entries/bulk | Remove one or more entries from the IP Allow List by ID
[**ListSecurityIpAllowlist**](SecurityIPAllowListAPI.md#ListSecurityIpAllowlist) | **Get** /v1/security/ip-allowlist | Get IP Allow List settings (mode, entry counts, max entries)
[**ListSecurityIpAllowlistCurrentIp**](SecurityIPAllowListAPI.md#ListSecurityIpAllowlistCurrentIp) | **Get** /v1/security/ip-allowlist/current-ip | Get the current user&#39;s IP address and whether it is in the allow list
[**ListSecurityIpAllowlistEntries**](SecurityIPAllowListAPI.md#ListSecurityIpAllowlistEntries) | **Get** /v1/security/ip-allowlist/entries | Get paginated list of IP Allow List entries with optional search filter
[**UpdateSecurityIpAllowlistEntries**](SecurityIPAllowListAPI.md#UpdateSecurityIpAllowlistEntries) | **Put** /v1/security/ip-allowlist/entries/{id} | Update an existing entry in the IP Allow List
[**UpdateSecurityIpAllowlistMode**](SecurityIPAllowListAPI.md#UpdateSecurityIpAllowlistMode) | **Put** /v1/security/ip-allowlist/mode | Update the IP Allow List operational mode



## CreateSecurityIpAllowlistEntries

> CreateSecurityIpAllowlistEntries(ctx).EntryXO(entryXO).Execute()

Add a new entry to the IP Allow List

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
	entryXO := *sonatyperepo.NewEntryXO("192.168.1.0/24") // EntryXO | Entry to add

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityIPAllowListAPI.CreateSecurityIpAllowlistEntries(context.Background()).EntryXO(entryXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.CreateSecurityIpAllowlistEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityIpAllowlistEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entryXO** | [**EntryXO**](EntryXO.md) | Entry to add | 

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


## CreateSecurityIpAllowlistEntriesBulk

> BulkUploadResultXO CreateSecurityIpAllowlistEntriesBulk(ctx).BulkUploadRequest(bulkUploadRequest).Execute()

Bulk upload entries from CSV content

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
	bulkUploadRequest := *sonatyperepo.NewBulkUploadRequest("CsvContent_example") // BulkUploadRequest | CSV upload request

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIPAllowListAPI.CreateSecurityIpAllowlistEntriesBulk(context.Background()).BulkUploadRequest(bulkUploadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.CreateSecurityIpAllowlistEntriesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityIpAllowlistEntriesBulk`: BulkUploadResultXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.CreateSecurityIpAllowlistEntriesBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityIpAllowlistEntriesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUploadRequest** | [**BulkUploadRequest**](BulkUploadRequest.md) | CSV upload request | 

### Return type

[**BulkUploadResultXO**](BulkUploadResultXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityIpAllowlistEntries

> DeleteSecurityIpAllowlistEntries(ctx).Confirm(confirm).Execute()

Clear all entries from the IP Allow List

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
	confirm := true // bool | Must be 'true' to confirm the destructive operation

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityIPAllowListAPI.DeleteSecurityIpAllowlistEntries(context.Background()).Confirm(confirm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.DeleteSecurityIpAllowlistEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityIpAllowlistEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **confirm** | **bool** | Must be &#39;true&#39; to confirm the destructive operation | 

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


## DeleteSecurityIpAllowlistEntriesBulk

> DeleteSecurityIpAllowlistEntriesBulk(ctx).RequestBody(requestBody).Execute()

Remove one or more entries from the IP Allow List by ID

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
	requestBody := []string{"Property_example"} // []string | List of entry IDs to remove

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityIPAllowListAPI.DeleteSecurityIpAllowlistEntriesBulk(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.DeleteSecurityIpAllowlistEntriesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityIpAllowlistEntriesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | List of entry IDs to remove | 

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


## ListSecurityIpAllowlist

> IpAllowListSettingsXO ListSecurityIpAllowlist(ctx).Execute()

Get IP Allow List settings (mode, entry counts, max entries)

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
	resp, r, err := apiClient.SecurityIPAllowListAPI.ListSecurityIpAllowlist(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.ListSecurityIpAllowlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityIpAllowlist`: IpAllowListSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.ListSecurityIpAllowlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityIpAllowlistRequest struct via the builder pattern


### Return type

[**IpAllowListSettingsXO**](IpAllowListSettingsXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityIpAllowlistCurrentIp

> ListSecurityIpAllowlistCurrentIp(ctx).Execute()

Get the current user's IP address and whether it is in the allow list

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
	r, err := apiClient.SecurityIPAllowListAPI.ListSecurityIpAllowlistCurrentIp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.ListSecurityIpAllowlistCurrentIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityIpAllowlistCurrentIpRequest struct via the builder pattern


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


## ListSecurityIpAllowlistEntries

> IpAllowListEntriesPageXO ListSecurityIpAllowlistEntries(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

Get paginated list of IP Allow List entries with optional search filter

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
	page := int32(56) // int32 | Page number (0-based) (optional) (default to 0)
	pageSize := int32(56) // int32 | Number of entries per page (optional) (default to 20)
	search := "search_example" // string | Search query (matches against IP address, CIDR notation, and description) (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIPAllowListAPI.ListSecurityIpAllowlistEntries(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.ListSecurityIpAllowlistEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityIpAllowlistEntries`: IpAllowListEntriesPageXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.ListSecurityIpAllowlistEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityIpAllowlistEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number (0-based) | [default to 0]
 **pageSize** | **int32** | Number of entries per page | [default to 20]
 **search** | **string** | Search query (matches against IP address, CIDR notation, and description) | 

### Return type

[**IpAllowListEntriesPageXO**](IpAllowListEntriesPageXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityIpAllowlistEntries

> IpAllowListEntryXO UpdateSecurityIpAllowlistEntries(ctx, id).EntryXO(entryXO).Execute()

Update an existing entry in the IP Allow List

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
	id := "id_example" // string | Entry ID
	entryXO := *sonatyperepo.NewEntryXO("192.168.1.0/24") // EntryXO | Updated entry data

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIPAllowListAPI.UpdateSecurityIpAllowlistEntries(context.Background(), id).EntryXO(entryXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.UpdateSecurityIpAllowlistEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityIpAllowlistEntries`: IpAllowListEntryXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.UpdateSecurityIpAllowlistEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityIpAllowlistEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entryXO** | [**EntryXO**](EntryXO.md) | Updated entry data | 

### Return type

[**IpAllowListEntryXO**](IpAllowListEntryXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityIpAllowlistMode

> UpdateSecurityIpAllowlistMode(ctx).ModeChangeXO(modeChangeXO).Execute()

Update the IP Allow List operational mode

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
	modeChangeXO := *sonatyperepo.NewModeChangeXO("MONITOR") // ModeChangeXO | Mode change request

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityIPAllowListAPI.UpdateSecurityIpAllowlistMode(context.Background()).ModeChangeXO(modeChangeXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.UpdateSecurityIpAllowlistMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityIpAllowlistModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modeChangeXO** | [**ModeChangeXO**](ModeChangeXO.md) | Mode change request | 

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


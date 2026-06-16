# \SecurityIPAllowListAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEntry**](SecurityIPAllowListAPI.md#AddEntry) | **Post** /v1/security/ip-allowlist/entries | Add a new entry to the IP Allow List
[**BrowseEntries**](SecurityIPAllowListAPI.md#BrowseEntries) | **Get** /v1/security/ip-allowlist/entries | Get paginated list of IP Allow List entries with optional search filter
[**BulkDelete**](SecurityIPAllowListAPI.md#BulkDelete) | **Delete** /v1/security/ip-allowlist/entries/bulk | Remove one or more entries from the IP Allow List by ID
[**BulkUpload**](SecurityIPAllowListAPI.md#BulkUpload) | **Post** /v1/security/ip-allowlist/entries/bulk | Bulk upload entries from CSV content
[**ClearAllEntries**](SecurityIPAllowListAPI.md#ClearAllEntries) | **Delete** /v1/security/ip-allowlist/entries | Clear all entries from the IP Allow List
[**GetCurrentIp**](SecurityIPAllowListAPI.md#GetCurrentIp) | **Get** /v1/security/ip-allowlist/current-ip | Get the current user&#39;s IP address and whether it is in the allow list
[**GetSettings**](SecurityIPAllowListAPI.md#GetSettings) | **Get** /v1/security/ip-allowlist | Get IP Allow List settings (mode, entry counts, max entries)
[**UpdateEntry**](SecurityIPAllowListAPI.md#UpdateEntry) | **Put** /v1/security/ip-allowlist/entries/{id} | Update an existing entry in the IP Allow List
[**UpdateMode**](SecurityIPAllowListAPI.md#UpdateMode) | **Put** /v1/security/ip-allowlist/mode | Update the IP Allow List operational mode



## AddEntry

> AddEntry(ctx).Body(body).Execute()

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
	body := *sonatyperepo.NewEntryXO("192.168.1.0/24") // EntryXO | Entry to add

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityIPAllowListAPI.AddEntry(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.AddEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EntryXO**](EntryXO.md) | Entry to add | 

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


## BrowseEntries

> IpAllowListEntriesPageXO BrowseEntries(ctx).Page(page).PageSize(pageSize).Search(search).Execute()

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
	resp, r, err := apiClient.SecurityIPAllowListAPI.BrowseEntries(context.Background()).Page(page).PageSize(pageSize).Search(search).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.BrowseEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BrowseEntries`: IpAllowListEntriesPageXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.BrowseEntries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBrowseEntriesRequest struct via the builder pattern


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


## BulkDelete

> BulkDelete(ctx).Body(body).Execute()

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
	body := []string{"Property_example"} // []string | List of entry IDs to remove

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityIPAllowListAPI.BulkDelete(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.BulkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of entry IDs to remove | 

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


## BulkUpload

> BulkUploadResultXO BulkUpload(ctx).Body(body).Execute()

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
	body := *sonatyperepo.NewBulkUploadRequest("CsvContent_example") // BulkUploadRequest | CSV upload request

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIPAllowListAPI.BulkUpload(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.BulkUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkUpload`: BulkUploadResultXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.BulkUpload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**BulkUploadRequest**](BulkUploadRequest.md) | CSV upload request | 

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


## ClearAllEntries

> ClearAllEntries(ctx).Confirm(confirm).Execute()

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
	r, err := apiClient.SecurityIPAllowListAPI.ClearAllEntries(context.Background()).Confirm(confirm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.ClearAllEntries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearAllEntriesRequest struct via the builder pattern


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


## GetCurrentIp

> GetCurrentIp(ctx).Execute()

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
	r, err := apiClient.SecurityIPAllowListAPI.GetCurrentIp(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.GetCurrentIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentIpRequest struct via the builder pattern


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


## GetSettings

> IpAllowListSettingsXO GetSettings(ctx).Execute()

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
	resp, r, err := apiClient.SecurityIPAllowListAPI.GetSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: IpAllowListSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


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


## UpdateEntry

> IpAllowListEntryXO UpdateEntry(ctx, id).Body(body).Execute()

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
	body := *sonatyperepo.NewEntryXO("192.168.1.0/24") // EntryXO | Updated entry data

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityIPAllowListAPI.UpdateEntry(context.Background(), id).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.UpdateEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntry`: IpAllowListEntryXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityIPAllowListAPI.UpdateEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Entry ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**EntryXO**](EntryXO.md) | Updated entry data | 

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


## UpdateMode

> UpdateMode(ctx).Body(body).Execute()

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
	body := *sonatyperepo.NewModeChangeXO("MONITOR") // ModeChangeXO | Mode change request

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityIPAllowListAPI.UpdateMode(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityIPAllowListAPI.UpdateMode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ModeChangeXO**](ModeChangeXO.md) | Mode change request | 

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


# \SecurityManagementRealmsAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSecurityRealmsActive**](SecurityManagementRealmsAPI.md#ListSecurityRealmsActive) | **Get** /v1/security/realms/active | List the active realm IDs in order
[**ListSecurityRealmsAvailable**](SecurityManagementRealmsAPI.md#ListSecurityRealmsAvailable) | **Get** /v1/security/realms/available | List the available realms
[**UpdateSecurityRealmsActive**](SecurityManagementRealmsAPI.md#UpdateSecurityRealmsActive) | **Put** /v1/security/realms/active | Set the active security realms in the order they should be used



## ListSecurityRealmsActive

> []string ListSecurityRealmsActive(ctx).Execute()

List the active realm IDs in order

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
	resp, r, err := apiClient.SecurityManagementRealmsAPI.ListSecurityRealmsActive(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRealmsAPI.ListSecurityRealmsActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityRealmsActive`: []string
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementRealmsAPI.ListSecurityRealmsActive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityRealmsActiveRequest struct via the builder pattern


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


## ListSecurityRealmsAvailable

> []RealmApiXO ListSecurityRealmsAvailable(ctx).Execute()

List the available realms

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
	resp, r, err := apiClient.SecurityManagementRealmsAPI.ListSecurityRealmsAvailable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRealmsAPI.ListSecurityRealmsAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityRealmsAvailable`: []RealmApiXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementRealmsAPI.ListSecurityRealmsAvailable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityRealmsAvailableRequest struct via the builder pattern


### Return type

[**[]RealmApiXO**](RealmApiXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityRealmsActive

> UpdateSecurityRealmsActive(ctx).RequestBody(requestBody).Execute()

Set the active security realms in the order they should be used

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
	requestBody := []string{"Property_example"} // []string | The realm IDs (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementRealmsAPI.UpdateSecurityRealmsActive(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRealmsAPI.UpdateSecurityRealmsActive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRealmsActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | The realm IDs | 

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


# \SecurityManagementAnonymousAccessAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Read**](SecurityManagementAnonymousAccessAPI.md#Read) | **Get** /v1/security/anonymous | Get Anonymous Access settings
[**Update**](SecurityManagementAnonymousAccessAPI.md#Update) | **Put** /v1/security/anonymous | Update Anonymous Access settings



## Read

> AnonymousAccessSettingsXO Read(ctx).Execute()

Get Anonymous Access settings

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
	resp, r, err := apiClient.SecurityManagementAnonymousAccessAPI.Read(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAnonymousAccessAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: AnonymousAccessSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAnonymousAccessAPI.Read`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadRequest struct via the builder pattern


### Return type

[**AnonymousAccessSettingsXO**](AnonymousAccessSettingsXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> AnonymousAccessSettingsXO Update(ctx).Body(body).Execute()

Update Anonymous Access settings

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
	body := *sonatyperepo.NewAnonymousAccessSettingsXO() // AnonymousAccessSettingsXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementAnonymousAccessAPI.Update(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementAnonymousAccessAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: AnonymousAccessSettingsXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementAnonymousAccessAPI.Update`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AnonymousAccessSettingsXO**](AnonymousAccessSettingsXO.md) |  | 

### Return type

[**AnonymousAccessSettingsXO**](AnonymousAccessSettingsXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SecurityManagementApiKeysPrincipalsEncryptionAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReEncrypt**](SecurityManagementApiKeysPrincipalsEncryptionAPI.md#ReEncrypt) | **Put** /v1/apikeys/encryption/re-encrypt | Re-encrypt api keys principals using the specified configuration



## ReEncrypt

> ReEncrypt(ctx).Body(body).Execute()

Re-encrypt api keys principals using the specified configuration



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
	body := *sonatyperepo.NewApiKeysReEncryptionRequestApiXO() // ApiKeysReEncryptionRequestApiXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementApiKeysPrincipalsEncryptionAPI.ReEncrypt(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementApiKeysPrincipalsEncryptionAPI.ReEncrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiKeysReEncryptionRequestApiXO**](ApiKeysReEncryptionRequestApiXO.md) |  | 

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


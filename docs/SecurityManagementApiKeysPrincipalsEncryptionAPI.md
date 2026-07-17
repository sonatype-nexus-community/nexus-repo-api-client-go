# \SecurityManagementApiKeysPrincipalsEncryptionAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateApikeysEncryptionReEncrypt**](SecurityManagementApiKeysPrincipalsEncryptionAPI.md#UpdateApikeysEncryptionReEncrypt) | **Put** /v1/apikeys/encryption/re-encrypt | Re-encrypt api keys principals using the specified configuration



## UpdateApikeysEncryptionReEncrypt

> UpdateApikeysEncryptionReEncrypt(ctx).ApiKeysReEncryptionRequestApiXO(apiKeysReEncryptionRequestApiXO).Execute()

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
	apiKeysReEncryptionRequestApiXO := *sonatyperepo.NewApiKeysReEncryptionRequestApiXO() // ApiKeysReEncryptionRequestApiXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementApiKeysPrincipalsEncryptionAPI.UpdateApikeysEncryptionReEncrypt(context.Background()).ApiKeysReEncryptionRequestApiXO(apiKeysReEncryptionRequestApiXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementApiKeysPrincipalsEncryptionAPI.UpdateApikeysEncryptionReEncrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApikeysEncryptionReEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKeysReEncryptionRequestApiXO** | [**ApiKeysReEncryptionRequestApiXO**](ApiKeysReEncryptionRequestApiXO.md) |  | 

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


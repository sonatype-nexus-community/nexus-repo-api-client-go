# \SecurityManagementSecretsEncryptionAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSecretsEncryptionReEncrypt**](SecurityManagementSecretsEncryptionAPI.md#UpdateSecretsEncryptionReEncrypt) | **Put** /v1/secrets/encryption/re-encrypt | Re-encrypt secrets using the specified key



## UpdateSecretsEncryptionReEncrypt

> UpdateSecretsEncryptionReEncrypt(ctx).ReEncryptionRequestApiXO(reEncryptionRequestApiXO).Execute()

Re-encrypt secrets using the specified key



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
	reEncryptionRequestApiXO := *sonatyperepo.NewReEncryptionRequestApiXO("SecretKeyId_example") // ReEncryptionRequestApiXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSecretsEncryptionAPI.UpdateSecretsEncryptionReEncrypt(context.Background()).ReEncryptionRequestApiXO(reEncryptionRequestApiXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSecretsEncryptionAPI.UpdateSecretsEncryptionReEncrypt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretsEncryptionReEncryptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reEncryptionRequestApiXO** | [**ReEncryptionRequestApiXO**](ReEncryptionRequestApiXO.md) |  | 

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


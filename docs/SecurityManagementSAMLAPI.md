# \SecurityManagementSAMLAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSamlConfiguration**](SecurityManagementSAMLAPI.md#DeleteSamlConfiguration) | **Delete** /v1/security/saml | Delete SAML configuration
[**GetMetadata**](SecurityManagementSAMLAPI.md#GetMetadata) | **Get** /v1/security/saml/metadata | Get service provider metadata XML document
[**GetPublicCertificateInPemFormat**](SecurityManagementSAMLAPI.md#GetPublicCertificateInPemFormat) | **Get** /v1/security/saml/pem | Get service provider signing certificate in PEM format
[**GetSamlConfiguration**](SecurityManagementSAMLAPI.md#GetSamlConfiguration) | **Get** /v1/security/saml | Get SAML configuration
[**PutSamlConfiguration**](SecurityManagementSAMLAPI.md#PutSamlConfiguration) | **Put** /v1/security/saml | Create or update SAML configuration



## DeleteSamlConfiguration

> DeleteSamlConfiguration(ctx).Execute()

Delete SAML configuration

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
	r, err := apiClient.SecurityManagementSAMLAPI.DeleteSamlConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLAPI.DeleteSamlConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSamlConfigurationRequest struct via the builder pattern


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


## GetMetadata

> GetMetadata(ctx).Execute()

Get service provider metadata XML document

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
	r, err := apiClient.SecurityManagementSAMLAPI.GetMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLAPI.GetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRequest struct via the builder pattern


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


## GetPublicCertificateInPemFormat

> GetPublicCertificateInPemFormat(ctx).Execute()

Get service provider signing certificate in PEM format

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
	r, err := apiClient.SecurityManagementSAMLAPI.GetPublicCertificateInPemFormat(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLAPI.GetPublicCertificateInPemFormat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicCertificateInPemFormatRequest struct via the builder pattern


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


## GetSamlConfiguration

> GetSamlConfiguration(ctx).Execute()

Get SAML configuration

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
	r, err := apiClient.SecurityManagementSAMLAPI.GetSamlConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLAPI.GetSamlConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSamlConfigurationRequest struct via the builder pattern


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


## PutSamlConfiguration

> PutSamlConfiguration(ctx).Body(body).Execute()

Create or update SAML configuration

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
	body := *sonatyperepo.NewSamlConfigurationXO("IdpMetadata_example", "UsernameAttribute_example") // SamlConfigurationXO |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementSAMLAPI.PutSamlConfiguration(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementSAMLAPI.PutSamlConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutSamlConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SamlConfigurationXO**](SamlConfigurationXO.md) |  | 

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


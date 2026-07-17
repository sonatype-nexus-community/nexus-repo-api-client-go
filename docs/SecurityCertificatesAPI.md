# \SecurityCertificatesAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecuritySslTruststore**](SecurityCertificatesAPI.md#CreateSecuritySslTruststore) | **Post** /v1/security/ssl/truststore | Add a certificate to the trust store.
[**DeleteSecuritySslTruststore**](SecurityCertificatesAPI.md#DeleteSecuritySslTruststore) | **Delete** /v1/security/ssl/truststore/{id} | Remove a certificate in the trust store.
[**ListSecuritySsl**](SecurityCertificatesAPI.md#ListSecuritySsl) | **Get** /v1/security/ssl | Helper method to retrieve certificate details from a remote system.
[**ListSecuritySslTruststore**](SecurityCertificatesAPI.md#ListSecuritySslTruststore) | **Get** /v1/security/ssl/truststore | Retrieve a list of certificates added to the trust store.



## CreateSecuritySslTruststore

> ApiCertificate CreateSecuritySslTruststore(ctx).Body(body).Execute()

Add a certificate to the trust store.

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
	body := "body_example" // string | The certificate to add encoded in PEM format

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityCertificatesAPI.CreateSecuritySslTruststore(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.CreateSecuritySslTruststore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecuritySslTruststore`: ApiCertificate
	fmt.Fprintf(os.Stdout, "Response from `SecurityCertificatesAPI.CreateSecuritySslTruststore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecuritySslTruststoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | The certificate to add encoded in PEM format | 

### Return type

[**ApiCertificate**](ApiCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecuritySslTruststore

> DeleteSecuritySslTruststore(ctx, id).Execute()

Remove a certificate in the trust store.

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
	id := "id_example" // string | The id of the certificate that should be removed.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityCertificatesAPI.DeleteSecuritySslTruststore(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.DeleteSecuritySslTruststore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the certificate that should be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecuritySslTruststoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListSecuritySsl

> ListSecuritySsl(ctx).Host(host).Port(port).ProtocolHint(protocolHint).Execute()

Helper method to retrieve certificate details from a remote system.

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
	host := "host_example" // string | The remote system's host name
	port := int32(56) // int32 | The port on the remote system to connect to (optional) (default to 443)
	protocolHint := "protocolHint_example" // string | An optional hint of the protocol to try for the connection (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityCertificatesAPI.ListSecuritySsl(context.Background()).Host(host).Port(port).ProtocolHint(protocolHint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.ListSecuritySsl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecuritySslRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | The remote system&#39;s host name | 
 **port** | **int32** | The port on the remote system to connect to | [default to 443]
 **protocolHint** | **string** | An optional hint of the protocol to try for the connection | 

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


## ListSecuritySslTruststore

> ListSecuritySslTruststore(ctx).Execute()

Retrieve a list of certificates added to the trust store.

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
	r, err := apiClient.SecurityCertificatesAPI.ListSecuritySslTruststore(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.ListSecuritySslTruststore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecuritySslTruststoreRequest struct via the builder pattern


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


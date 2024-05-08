# \SecurityCertificatesAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCertificate**](SecurityCertificatesAPI.md#AddCertificate) | **Post** /v1/security/ssl/truststore | Add a certificate to the trust store.
[**GetTrustStoreCertificates**](SecurityCertificatesAPI.md#GetTrustStoreCertificates) | **Get** /v1/security/ssl/truststore | Retrieve a list of certificates added to the trust store.
[**RemoveCertificate**](SecurityCertificatesAPI.md#RemoveCertificate) | **Delete** /v1/security/ssl/truststore/{id} | Remove a certificate in the trust store.
[**RetrieveCertificate**](SecurityCertificatesAPI.md#RetrieveCertificate) | **Get** /v1/security/ssl | Helper method to retrieve certificate details from a remote system.



## AddCertificate

> ApiCertificate AddCertificate(ctx).Body(body).Execute()

Add a certificate to the trust store.

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
	body := "body_example" // string | The certificate to add encoded in PEM format (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityCertificatesAPI.AddCertificate(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.AddCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddCertificate`: ApiCertificate
	fmt.Fprintf(os.Stdout, "Response from `SecurityCertificatesAPI.AddCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddCertificateRequest struct via the builder pattern


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


## GetTrustStoreCertificates

> []ApiCertificate GetTrustStoreCertificates(ctx).Execute()

Retrieve a list of certificates added to the trust store.

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
	resp, r, err := apiClient.SecurityCertificatesAPI.GetTrustStoreCertificates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.GetTrustStoreCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrustStoreCertificates`: []ApiCertificate
	fmt.Fprintf(os.Stdout, "Response from `SecurityCertificatesAPI.GetTrustStoreCertificates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrustStoreCertificatesRequest struct via the builder pattern


### Return type

[**[]ApiCertificate**](ApiCertificate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCertificate

> RemoveCertificate(ctx, id).Execute()

Remove a certificate in the trust store.

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
	id := "id_example" // string | The id of the certificate that should be removed.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityCertificatesAPI.RemoveCertificate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.RemoveCertificate``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveCertificateRequest struct via the builder pattern


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


## RetrieveCertificate

> ApiCertificate RetrieveCertificate(ctx).Host(host).Port(port).ProtocolHint(protocolHint).Execute()

Helper method to retrieve certificate details from a remote system.

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
	host := "host_example" // string | The remote system's host name
	port := int32(56) // int32 | The port on the remote system to connect to (optional) (default to 443)
	protocolHint := "protocolHint_example" // string | An optional hint of the protocol to try for the connection (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityCertificatesAPI.RetrieveCertificate(context.Background()).Host(host).Port(port).ProtocolHint(protocolHint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityCertificatesAPI.RetrieveCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCertificate`: ApiCertificate
	fmt.Fprintf(os.Stdout, "Response from `SecurityCertificatesAPI.RetrieveCertificate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCertificateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **host** | **string** | The remote system&#39;s host name | 
 **port** | **int32** | The port on the remote system to connect to | [default to 443]
 **protocolHint** | **string** | An optional hint of the protocol to try for the connection | 

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


# \SecurityOAuth2OIDCAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSecurityOauth2**](SecurityOAuth2OIDCAPI.md#DeleteSecurityOauth2) | **Delete** /v1/security/oauth2 | Remove OAuth2/OIDC configuration
[**ListSecurityOauth2**](SecurityOAuth2OIDCAPI.md#ListSecurityOauth2) | **Get** /v1/security/oauth2 | Retrieve the current OAuth2/OIDC configuration
[**UpdateSecurityOauth2**](SecurityOAuth2OIDCAPI.md#UpdateSecurityOauth2) | **Put** /v1/security/oauth2 | Create or update OAuth2/OIDC configuration



## DeleteSecurityOauth2

> DeleteSecurityOauth2(ctx).Execute()

Remove OAuth2/OIDC configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityOAuth2OIDCAPI.DeleteSecurityOauth2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOAuth2OIDCAPI.DeleteSecurityOauth2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityOauth2Request struct via the builder pattern


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


## ListSecurityOauth2

> OAuth2OidcConfigurationXO ListSecurityOauth2(ctx).Execute()

Retrieve the current OAuth2/OIDC configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityOAuth2OIDCAPI.ListSecurityOauth2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOAuth2OIDCAPI.ListSecurityOauth2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityOauth2`: OAuth2OidcConfigurationXO
	fmt.Fprintf(os.Stdout, "Response from `SecurityOAuth2OIDCAPI.ListSecurityOauth2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityOauth2Request struct via the builder pattern


### Return type

[**OAuth2OidcConfigurationXO**](OAuth2OidcConfigurationXO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityOauth2

> UpdateSecurityOauth2(ctx).OAuth2OidcConfigurationXO(oAuth2OidcConfigurationXO).Execute()

Create or update OAuth2/OIDC configuration

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go/v395"
)

func main() {
	oAuth2OidcConfigurationXO := *sonatyperepo.NewOAuth2OidcConfigurationXO("IdpJwsAlgorithm_example", "UsernameClaim_example") // OAuth2OidcConfigurationXO | The OAuth2/OIDC configuration to create or update

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityOAuth2OIDCAPI.UpdateSecurityOauth2(context.Background()).OAuth2OidcConfigurationXO(oAuth2OidcConfigurationXO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityOAuth2OIDCAPI.UpdateSecurityOauth2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityOauth2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oAuth2OidcConfigurationXO** | [**OAuth2OidcConfigurationXO**](OAuth2OidcConfigurationXO.md) | The OAuth2/OIDC configuration to create or update | 

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


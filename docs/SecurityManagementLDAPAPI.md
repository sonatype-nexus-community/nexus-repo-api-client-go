# \SecurityManagementLDAPAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeOrder**](SecurityManagementLDAPAPI.md#ChangeOrder) | **Post** /v1/security/ldap/change-order | Change LDAP server order
[**ClearCache1**](SecurityManagementLDAPAPI.md#ClearCache1) | **Delete** /v1/security/ldap/cache | 
[**CreateLdapServer**](SecurityManagementLDAPAPI.md#CreateLdapServer) | **Post** /v1/security/ldap | Create LDAP server
[**DeleteLdapServer**](SecurityManagementLDAPAPI.md#DeleteLdapServer) | **Delete** /v1/security/ldap/{name} | Delete LDAP server
[**GetLdapServer**](SecurityManagementLDAPAPI.md#GetLdapServer) | **Get** /v1/security/ldap/{name} | Get LDAP server
[**GetLdapServers**](SecurityManagementLDAPAPI.md#GetLdapServers) | **Get** /v1/security/ldap | List LDAP servers
[**GetTemplates**](SecurityManagementLDAPAPI.md#GetTemplates) | **Get** /v1/security/ldap/templates | 
[**UpdateLdapServer**](SecurityManagementLDAPAPI.md#UpdateLdapServer) | **Put** /v1/security/ldap/{name} | Update LDAP server
[**VerifyConnection1**](SecurityManagementLDAPAPI.md#VerifyConnection1) | **Post** /v1/security/ldap/verify-connection | 
[**VerifyLogin**](SecurityManagementLDAPAPI.md#VerifyLogin) | **Post** /v1/security/ldap/verify-login | 
[**VerifyUserMapping**](SecurityManagementLDAPAPI.md#VerifyUserMapping) | **Post** /v1/security/ldap/verify-user-mapping | 



## ChangeOrder

> ChangeOrder(ctx).Body(body).Execute()

Change LDAP server order

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
	body := []string{"Property_example"} // []string | Ordered list of LDAP server names (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.ChangeOrder(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.ChangeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | Ordered list of LDAP server names | 

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


## ClearCache1

> ClearCache1(ctx).Execute()



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
	r, err := apiClient.SecurityManagementLDAPAPI.ClearCache1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.ClearCache1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiClearCache1Request struct via the builder pattern


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


## CreateLdapServer

> CreateLdapServer(ctx).Body(body).Execute()

Create LDAP server

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
	body := *sonatyperepo.NewCreateLdapServerXo("AuthPassword_example", "AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com") // CreateLdapServerXo |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.CreateLdapServer(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.CreateLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLdapServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 

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


## DeleteLdapServer

> DeleteLdapServer(ctx, name).Execute()

Delete LDAP server

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
	name := "name_example" // string | Name of the LDAP server to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.DeleteLdapServer(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.DeleteLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP server to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLdapServerRequest struct via the builder pattern


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


## GetLdapServer

> ReadLdapServerXo GetLdapServer(ctx, name).Execute()

Get LDAP server

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
	name := "name_example" // string | Name of the LDAP server to retrieve

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementLDAPAPI.GetLdapServer(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.GetLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapServer`: ReadLdapServerXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.GetLdapServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReadLdapServerXo**](ReadLdapServerXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLdapServers

> []ReadLdapServerXo GetLdapServers(ctx).Execute()

List LDAP servers

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
	resp, r, err := apiClient.SecurityManagementLDAPAPI.GetLdapServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.GetLdapServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLdapServers`: []ReadLdapServerXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.GetLdapServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLdapServersRequest struct via the builder pattern


### Return type

[**[]ReadLdapServerXo**](ReadLdapServerXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplates

> []LdapSchemaTemplateXo GetTemplates(ctx).Execute()



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
	resp, r, err := apiClient.SecurityManagementLDAPAPI.GetTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.GetTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplates`: []LdapSchemaTemplateXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.GetTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplatesRequest struct via the builder pattern


### Return type

[**[]LdapSchemaTemplateXo**](LdapSchemaTemplateXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLdapServer

> UpdateLdapServer(ctx, name).Body(body).Execute()

Update LDAP server

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
	name := "name_example" // string | Name of the LDAP server to update
	body := *sonatyperepo.NewUpdateLdapServerXo("AuthPassword_example", "AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com") // UpdateLdapServerXo | Updated values of LDAP server (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.UpdateLdapServer(context.Background(), name).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.UpdateLdapServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP server to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLdapServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateLdapServerXo**](UpdateLdapServerXo.md) | Updated values of LDAP server | 

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


## VerifyConnection1

> VerifyConnection1(ctx).ExistingServerName(existingServerName).Body(body).Execute()



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
	existingServerName := "existingServerName_example" // string |  (optional)
	body := *sonatyperepo.NewCreateLdapServerXo("AuthPassword_example", "AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com") // CreateLdapServerXo |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.VerifyConnection1(context.Background()).ExistingServerName(existingServerName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.VerifyConnection1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyConnection1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **existingServerName** | **string** |  | 
 **body** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 

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


## VerifyLogin

> VerifyLogin(ctx).Body(body).Execute()



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
	body := *sonatyperepo.NewLdapVerifyLoginXo("Base64Password_example", "Base64Username_example", *sonatyperepo.NewCreateLdapServerXo("AuthPassword_example", "AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com")) // LdapVerifyLoginXo |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.VerifyLogin(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.VerifyLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LdapVerifyLoginXo**](LdapVerifyLoginXo.md) |  | 

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


## VerifyUserMapping

> []LdapTestUserMappingXo VerifyUserMapping(ctx).ExistingServerName(existingServerName).Body(body).Execute()



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
	existingServerName := "existingServerName_example" // string |  (optional)
	body := *sonatyperepo.NewCreateLdapServerXo("AuthPassword_example", "AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com") // CreateLdapServerXo |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementLDAPAPI.VerifyUserMapping(context.Background()).ExistingServerName(existingServerName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.VerifyUserMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyUserMapping`: []LdapTestUserMappingXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.VerifyUserMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **existingServerName** | **string** |  | 
 **body** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 

### Return type

[**[]LdapTestUserMappingXo**](LdapTestUserMappingXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \SecurityManagementLDAPAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityLdap**](SecurityManagementLDAPAPI.md#CreateSecurityLdap) | **Post** /v1/security/ldap | Create LDAP server
[**CreateSecurityLdapChangeOrder**](SecurityManagementLDAPAPI.md#CreateSecurityLdapChangeOrder) | **Post** /v1/security/ldap/change-order | Change LDAP server order
[**CreateSecurityLdapVerifyConnection**](SecurityManagementLDAPAPI.md#CreateSecurityLdapVerifyConnection) | **Post** /v1/security/ldap/verify-connection | 
[**CreateSecurityLdapVerifyLogin**](SecurityManagementLDAPAPI.md#CreateSecurityLdapVerifyLogin) | **Post** /v1/security/ldap/verify-login | 
[**CreateSecurityLdapVerifyUserMapping**](SecurityManagementLDAPAPI.md#CreateSecurityLdapVerifyUserMapping) | **Post** /v1/security/ldap/verify-user-mapping | 
[**DeleteSecurityLdap**](SecurityManagementLDAPAPI.md#DeleteSecurityLdap) | **Delete** /v1/security/ldap/{name} | Delete LDAP server
[**DeleteSecurityLdapCache**](SecurityManagementLDAPAPI.md#DeleteSecurityLdapCache) | **Delete** /v1/security/ldap/cache | 
[**GetSecurityLdap**](SecurityManagementLDAPAPI.md#GetSecurityLdap) | **Get** /v1/security/ldap/{name} | Get LDAP server
[**ListSecurityLdap**](SecurityManagementLDAPAPI.md#ListSecurityLdap) | **Get** /v1/security/ldap | List LDAP servers
[**ListSecurityLdapTemplates**](SecurityManagementLDAPAPI.md#ListSecurityLdapTemplates) | **Get** /v1/security/ldap/templates | 
[**UpdateSecurityLdap**](SecurityManagementLDAPAPI.md#UpdateSecurityLdap) | **Put** /v1/security/ldap/{name} | Update LDAP server



## CreateSecurityLdap

> CreateSecurityLdap(ctx).CreateLdapServerXo(createLdapServerXo).Execute()

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
	createLdapServerXo := *sonatyperepo.NewCreateLdapServerXo("AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com", "mail", "uid", "inetOrgPerson", "cn") // CreateLdapServerXo | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.CreateSecurityLdap(context.Background()).CreateLdapServerXo(createLdapServerXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.CreateSecurityLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLdapServerXo** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 

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


## CreateSecurityLdapChangeOrder

> CreateSecurityLdapChangeOrder(ctx).RequestBody(requestBody).Execute()

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
	requestBody := []string{"Property_example"} // []string | Ordered list of LDAP server names

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.CreateSecurityLdapChangeOrder(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.CreateSecurityLdapChangeOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityLdapChangeOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]string** | Ordered list of LDAP server names | 

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


## CreateSecurityLdapVerifyConnection

> CreateSecurityLdapVerifyConnection(ctx).CreateLdapServerXo(createLdapServerXo).ExistingServerName(existingServerName).Execute()



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
	createLdapServerXo := *sonatyperepo.NewCreateLdapServerXo("AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com", "mail", "uid", "inetOrgPerson", "cn") // CreateLdapServerXo | 
	existingServerName := "existingServerName_example" // string |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.CreateSecurityLdapVerifyConnection(context.Background()).CreateLdapServerXo(createLdapServerXo).ExistingServerName(existingServerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.CreateSecurityLdapVerifyConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityLdapVerifyConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLdapServerXo** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 
 **existingServerName** | **string** |  | 

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


## CreateSecurityLdapVerifyLogin

> CreateSecurityLdapVerifyLogin(ctx).LdapVerifyLoginXo(ldapVerifyLoginXo).Execute()



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
	ldapVerifyLoginXo := *sonatyperepo.NewLdapVerifyLoginXo("Base64Password_example", "Base64Username_example", *sonatyperepo.NewCreateLdapServerXo("AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com", "mail", "uid", "inetOrgPerson", "cn")) // LdapVerifyLoginXo | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.CreateSecurityLdapVerifyLogin(context.Background()).LdapVerifyLoginXo(ldapVerifyLoginXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.CreateSecurityLdapVerifyLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityLdapVerifyLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapVerifyLoginXo** | [**LdapVerifyLoginXo**](LdapVerifyLoginXo.md) |  | 

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


## CreateSecurityLdapVerifyUserMapping

> []LdapTestUserMappingXo CreateSecurityLdapVerifyUserMapping(ctx).CreateLdapServerXo(createLdapServerXo).ExistingServerName(existingServerName).Execute()



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
	createLdapServerXo := *sonatyperepo.NewCreateLdapServerXo("AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com", "mail", "uid", "inetOrgPerson", "cn") // CreateLdapServerXo | 
	existingServerName := "existingServerName_example" // string |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementLDAPAPI.CreateSecurityLdapVerifyUserMapping(context.Background()).CreateLdapServerXo(createLdapServerXo).ExistingServerName(existingServerName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.CreateSecurityLdapVerifyUserMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityLdapVerifyUserMapping`: []LdapTestUserMappingXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.CreateSecurityLdapVerifyUserMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityLdapVerifyUserMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createLdapServerXo** | [**CreateLdapServerXo**](CreateLdapServerXo.md) |  | 
 **existingServerName** | **string** |  | 

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


## DeleteSecurityLdap

> DeleteSecurityLdap(ctx, name).Execute()

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
	r, err := apiClient.SecurityManagementLDAPAPI.DeleteSecurityLdap(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.DeleteSecurityLdap``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSecurityLdapRequest struct via the builder pattern


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


## DeleteSecurityLdapCache

> DeleteSecurityLdapCache(ctx).Execute()



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
	r, err := apiClient.SecurityManagementLDAPAPI.DeleteSecurityLdapCache(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.DeleteSecurityLdapCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityLdapCacheRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityLdap

> ReadLdapServerXo GetSecurityLdap(ctx, name).Execute()

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
	resp, r, err := apiClient.SecurityManagementLDAPAPI.GetSecurityLdap(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.GetSecurityLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityLdap`: ReadLdapServerXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.GetSecurityLdap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Name of the LDAP server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityLdapRequest struct via the builder pattern


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


## ListSecurityLdap

> []ReadLdapServerXo ListSecurityLdap(ctx).Execute()

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
	resp, r, err := apiClient.SecurityManagementLDAPAPI.ListSecurityLdap(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.ListSecurityLdap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityLdap`: []ReadLdapServerXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.ListSecurityLdap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityLdapRequest struct via the builder pattern


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


## ListSecurityLdapTemplates

> []LdapSchemaTemplateXo ListSecurityLdapTemplates(ctx).Execute()



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
	resp, r, err := apiClient.SecurityManagementLDAPAPI.ListSecurityLdapTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.ListSecurityLdapTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityLdapTemplates`: []LdapSchemaTemplateXo
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementLDAPAPI.ListSecurityLdapTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityLdapTemplatesRequest struct via the builder pattern


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


## UpdateSecurityLdap

> UpdateSecurityLdap(ctx, name).UpdateLdapServerXo(updateLdapServerXo).Execute()

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
	updateLdapServerXo := *sonatyperepo.NewUpdateLdapServerXo("AuthScheme_example", int32(123), int32(1), "Host_example", int32(123), "Name_example", int32(636), "Protocol_example", "dc=example,dc=com", "mail", "uid", "inetOrgPerson", "cn") // UpdateLdapServerXo | Updated values of LDAP server

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementLDAPAPI.UpdateSecurityLdap(context.Background(), name).UpdateLdapServerXo(updateLdapServerXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementLDAPAPI.UpdateSecurityLdap``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateSecurityLdapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLdapServerXo** | [**UpdateLdapServerXo**](UpdateLdapServerXo.md) | Updated values of LDAP server | 

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


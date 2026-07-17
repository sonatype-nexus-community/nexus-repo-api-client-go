# \SecurityManagementRolesAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityRoles**](SecurityManagementRolesAPI.md#CreateSecurityRoles) | **Post** /v1/security/roles | Create role
[**DeleteSecurityRoles**](SecurityManagementRolesAPI.md#DeleteSecurityRoles) | **Delete** /v1/security/roles/{id} | Delete role
[**GetSecurityRoles**](SecurityManagementRolesAPI.md#GetSecurityRoles) | **Get** /v1/security/roles/{id} | Get role
[**ListSecurityRoles**](SecurityManagementRolesAPI.md#ListSecurityRoles) | **Get** /v1/security/roles | List roles
[**UpdateSecurityRoles**](SecurityManagementRolesAPI.md#UpdateSecurityRoles) | **Put** /v1/security/roles/{id} | Update role



## CreateSecurityRoles

> CreateSecurityRoles(ctx).RoleXORequest(roleXORequest).Execute()

Create role

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
	roleXORequest := *sonatyperepo.NewRoleXORequest("Id_example", "Name_example") // RoleXORequest | A role configuration

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementRolesAPI.CreateSecurityRoles(context.Background()).RoleXORequest(roleXORequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRolesAPI.CreateSecurityRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleXORequest** | [**RoleXORequest**](RoleXORequest.md) | A role configuration | 

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


## DeleteSecurityRoles

> DeleteSecurityRoles(ctx, id).Execute()

Delete role

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
	id := "id_example" // string | The id of the role to delete

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementRolesAPI.DeleteSecurityRoles(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRolesAPI.DeleteSecurityRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the role to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityRolesRequest struct via the builder pattern


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


## GetSecurityRoles

> GetSecurityRoles(ctx, id).Source(source).Execute()

Get role

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
	id := "id_example" // string | The id of the role to get
	source := "source_example" // string | The id of the user source to filter the roles by. Available sources can be fetched using the 'User Sources' endpoint. (optional) (default to "default")

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementRolesAPI.GetSecurityRoles(context.Background(), id).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRolesAPI.GetSecurityRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the role to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **source** | **string** | The id of the user source to filter the roles by. Available sources can be fetched using the &#39;User Sources&#39; endpoint. | [default to &quot;default&quot;]

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


## ListSecurityRoles

> ListSecurityRoles(ctx).Source(source).Execute()

List roles

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
	source := "source_example" // string | The id of the user source to filter the roles by, if supplied. Otherwise roles from all user sources will be returned. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementRolesAPI.ListSecurityRoles(context.Background()).Source(source).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRolesAPI.ListSecurityRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | The id of the user source to filter the roles by, if supplied. Otherwise roles from all user sources will be returned. | 

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


## UpdateSecurityRoles

> UpdateSecurityRoles(ctx, id).RoleXORequest(roleXORequest).Execute()

Update role

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
	id := "id_example" // string | The id of the role to update
	roleXORequest := *sonatyperepo.NewRoleXORequest("Id_example", "Name_example") // RoleXORequest | A role configuration

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementRolesAPI.UpdateSecurityRoles(context.Background(), id).RoleXORequest(roleXORequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementRolesAPI.UpdateSecurityRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of the role to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleXORequest** | [**RoleXORequest**](RoleXORequest.md) | A role configuration | 

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


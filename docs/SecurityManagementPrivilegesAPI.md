# \SecurityManagementPrivilegesAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationPrivilege**](SecurityManagementPrivilegesAPI.md#CreateApplicationPrivilege) | **Post** /v1/security/privileges/application | Create an application type privilege.
[**CreateRepositoryAdminPrivilege**](SecurityManagementPrivilegesAPI.md#CreateRepositoryAdminPrivilege) | **Post** /v1/security/privileges/repository-admin | Create a repository admin type privilege.
[**CreateRepositoryContentSelectorPrivilege**](SecurityManagementPrivilegesAPI.md#CreateRepositoryContentSelectorPrivilege) | **Post** /v1/security/privileges/repository-content-selector | Create a repository content selector type privilege.
[**CreateRepositoryViewPrivilege**](SecurityManagementPrivilegesAPI.md#CreateRepositoryViewPrivilege) | **Post** /v1/security/privileges/repository-view | Create a repository view type privilege.
[**CreateScriptPrivilege**](SecurityManagementPrivilegesAPI.md#CreateScriptPrivilege) | **Post** /v1/security/privileges/script | Create a script type privilege.
[**CreateWildcardPrivilege**](SecurityManagementPrivilegesAPI.md#CreateWildcardPrivilege) | **Post** /v1/security/privileges/wildcard | Create a wildcard type privilege.
[**DeletePrivilege**](SecurityManagementPrivilegesAPI.md#DeletePrivilege) | **Delete** /v1/security/privileges/{privilegeName} | Delete a privilege by name.
[**GetPrivilege**](SecurityManagementPrivilegesAPI.md#GetPrivilege) | **Get** /v1/security/privileges/{privilegeName} | Retrieve a privilege by name.
[**GetPrivileges1**](SecurityManagementPrivilegesAPI.md#GetPrivileges1) | **Get** /v1/security/privileges | Retrieve a list of privileges.
[**UpdateApplicationPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateApplicationPrivilege) | **Put** /v1/security/privileges/application/{privilegeName} | Update an application type privilege.
[**UpdateRepositoryAdminPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateRepositoryAdminPrivilege) | **Put** /v1/security/privileges/repository-admin/{privilegeName} | Update a repository admin type privilege.
[**UpdateRepositoryContentSelectorPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateRepositoryContentSelectorPrivilege) | **Put** /v1/security/privileges/repository-content-selector/{privilegeName} | Update a repository content selector type privilege.
[**UpdateRepositoryViewPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateRepositoryViewPrivilege) | **Put** /v1/security/privileges/repository-view/{privilegeName} | Update a repository view type privilege.
[**UpdateScriptPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateScriptPrivilege) | **Put** /v1/security/privileges/script/{privilegeName} | Update a script type privilege.
[**UpdateWildcardPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateWildcardPrivilege) | **Put** /v1/security/privileges/wildcard/{privilegeName} | Update a wildcard type privilege.



## CreateApplicationPrivilege

> CreateApplicationPrivilege(ctx).Body(body).Execute()

Create an application type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeApplicationRequest() // ApiPrivilegeApplicationRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateApplicationPrivilege(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreateApplicationPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md) | The privilege to create. | 

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


## CreateRepositoryAdminPrivilege

> CreateRepositoryAdminPrivilege(ctx).Body(body).Execute()

Create a repository admin type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeRepositoryAdminRequest() // ApiPrivilegeRepositoryAdminRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateRepositoryAdminPrivilege(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreateRepositoryAdminPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryAdminPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md) | The privilege to create. | 

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


## CreateRepositoryContentSelectorPrivilege

> CreateRepositoryContentSelectorPrivilege(ctx).Body(body).Execute()

Create a repository content selector type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeRepositoryContentSelectorRequest() // ApiPrivilegeRepositoryContentSelectorRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateRepositoryContentSelectorPrivilege(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreateRepositoryContentSelectorPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryContentSelectorPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md) | The privilege to create. | 

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


## CreateRepositoryViewPrivilege

> CreateRepositoryViewPrivilege(ctx).Body(body).Execute()

Create a repository view type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeRepositoryViewRequest() // ApiPrivilegeRepositoryViewRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateRepositoryViewPrivilege(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreateRepositoryViewPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRepositoryViewPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md) | The privilege to create. | 

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


## CreateScriptPrivilege

> CreateScriptPrivilege(ctx).Body(body).Execute()

Create a script type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeScriptRequest() // ApiPrivilegeScriptRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateScriptPrivilege(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreateScriptPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScriptPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md) | The privilege to create. | 

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


## CreateWildcardPrivilege

> CreateWildcardPrivilege(ctx).Body(body).Execute()

Create a wildcard type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeWildcardRequest() // ApiPrivilegeWildcardRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateWildcardPrivilege(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreateWildcardPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWildcardPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md) | The privilege to create. | 

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


## DeletePrivilege

> DeletePrivilege(ctx, privilegeName).Execute()

Delete a privilege by name.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to delete.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.DeletePrivilege(context.Background(), privilegeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.DeletePrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivilegeRequest struct via the builder pattern


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


## GetPrivilege

> ApiPrivilege GetPrivilege(ctx, privilegeName).Execute()

Retrieve a privilege by name.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to retrieve.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityManagementPrivilegesAPI.GetPrivilege(context.Background(), privilegeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.GetPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivilege`: ApiPrivilege
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementPrivilegesAPI.GetPrivilege`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiPrivilege**](ApiPrivilege.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivileges1

> []ApiPrivilege GetPrivileges1(ctx).Execute()

Retrieve a list of privileges.

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
	resp, r, err := apiClient.SecurityManagementPrivilegesAPI.GetPrivileges1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.GetPrivileges1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivileges1`: []ApiPrivilege
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementPrivilegesAPI.GetPrivileges1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivileges1Request struct via the builder pattern


### Return type

[**[]ApiPrivilege**](ApiPrivilege.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationPrivilege

> UpdateApplicationPrivilege(ctx, privilegeName).Body(body).Execute()

Update an application type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeApplicationRequest() // ApiPrivilegeApplicationRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateApplicationPrivilege(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdateApplicationPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md) | The privilege to update. | 

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


## UpdateRepositoryAdminPrivilege

> UpdateRepositoryAdminPrivilege(ctx, privilegeName).Body(body).Execute()

Update a repository admin type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeRepositoryAdminRequest() // ApiPrivilegeRepositoryAdminRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateRepositoryAdminPrivilege(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdateRepositoryAdminPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryAdminPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md) | The privilege to update. | 

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


## UpdateRepositoryContentSelectorPrivilege

> UpdateRepositoryContentSelectorPrivilege(ctx, privilegeName).Body(body).Execute()

Update a repository content selector type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeRepositoryContentSelectorRequest() // ApiPrivilegeRepositoryContentSelectorRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateRepositoryContentSelectorPrivilege(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdateRepositoryContentSelectorPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryContentSelectorPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md) | The privilege to update. | 

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


## UpdateRepositoryViewPrivilege

> UpdateRepositoryViewPrivilege(ctx, privilegeName).Body(body).Execute()

Update a repository view type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeRepositoryViewRequest() // ApiPrivilegeRepositoryViewRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateRepositoryViewPrivilege(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdateRepositoryViewPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRepositoryViewPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md) | The privilege to update. | 

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


## UpdateScriptPrivilege

> UpdateScriptPrivilege(ctx, privilegeName).Body(body).Execute()

Update a script type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeScriptRequest() // ApiPrivilegeScriptRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateScriptPrivilege(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdateScriptPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScriptPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md) | The privilege to update. | 

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


## UpdateWildcardPrivilege

> UpdateWildcardPrivilege(ctx, privilegeName).Body(body).Execute()

Update a wildcard type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeWildcardRequest() // ApiPrivilegeWildcardRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateWildcardPrivilege(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdateWildcardPrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWildcardPrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md) | The privilege to update. | 

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


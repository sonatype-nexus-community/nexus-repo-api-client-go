# \SecurityManagementPrivilegesAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationPrivilege**](SecurityManagementPrivilegesAPI.md#CreateApplicationPrivilege) | **Post** /v1/security/privileges/application | Create an application type privilege.
[**CreateRepositoryAdminPrivilege**](SecurityManagementPrivilegesAPI.md#CreateRepositoryAdminPrivilege) | **Post** /v1/security/privileges/repository-admin | Create a repository admin type privilege.
[**CreateRepositoryContentSelectorPrivilege**](SecurityManagementPrivilegesAPI.md#CreateRepositoryContentSelectorPrivilege) | **Post** /v1/security/privileges/repository-content-selector | Create a repository content selector type privilege.
[**CreateRepositoryViewPrivilege**](SecurityManagementPrivilegesAPI.md#CreateRepositoryViewPrivilege) | **Post** /v1/security/privileges/repository-view | Create a repository view type privilege.
[**CreateScriptPrivilege**](SecurityManagementPrivilegesAPI.md#CreateScriptPrivilege) | **Post** /v1/security/privileges/script | Create a script type privilege.
[**CreateWildcardPrivilege**](SecurityManagementPrivilegesAPI.md#CreateWildcardPrivilege) | **Post** /v1/security/privileges/wildcard | Create a wildcard type privilege.
[**DeleteSecurityPrivileges**](SecurityManagementPrivilegesAPI.md#DeleteSecurityPrivileges) | **Delete** /v1/security/privileges/{privilegeName} | Delete a privilege by name.
[**GetAllPrivileges**](SecurityManagementPrivilegesAPI.md#GetAllPrivileges) | **Get** /v1/security/privileges | Retrieve a list of privileges.
[**GetSecurityPrivileges**](SecurityManagementPrivilegesAPI.md#GetSecurityPrivileges) | **Get** /v1/security/privileges/{privilegeName} | Retrieve a privilege by name.
[**UpdateApplicationPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateApplicationPrivilege) | **Put** /v1/security/privileges/application/{privilegeName} | Update an application type privilege.
[**UpdateRepositoryAdminPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateRepositoryAdminPrivilege) | **Put** /v1/security/privileges/repository-admin/{privilegeName} | Update a repository admin type privilege.
[**UpdateRepositoryContentSelectorPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateRepositoryContentSelectorPrivilege) | **Put** /v1/security/privileges/repository-content-selector/{privilegeName} | Update a repository content selector type privilege.
[**UpdateRepositoryViewPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateRepositoryViewPrivilege) | **Put** /v1/security/privileges/repository-view/{privilegeName} | Update a repository view type privilege.
[**UpdateScriptPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateScriptPrivilege) | **Put** /v1/security/privileges/script/{privilegeName} | Update a script type privilege.
[**UpdateWildcardPrivilege**](SecurityManagementPrivilegesAPI.md#UpdateWildcardPrivilege) | **Put** /v1/security/privileges/wildcard/{privilegeName} | Update a wildcard type privilege.



## CreateApplicationPrivilege

> CreateApplicationPrivilege(ctx).ApiPrivilegeApplicationRequest(apiPrivilegeApplicationRequest).Execute()

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
	apiPrivilegeApplicationRequest := *sonatyperepo.NewApiPrivilegeApplicationRequest([]string{"Actions_example"}, "Domain_example", "Name_example") // ApiPrivilegeApplicationRequest | The privilege to create.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateApplicationPrivilege(context.Background()).ApiPrivilegeApplicationRequest(apiPrivilegeApplicationRequest).Execute()
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
 **apiPrivilegeApplicationRequest** | [**ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md) | The privilege to create. | 

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

> CreateRepositoryAdminPrivilege(ctx).ApiPrivilegeRepositoryAdminRequest(apiPrivilegeRepositoryAdminRequest).Execute()

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
	apiPrivilegeRepositoryAdminRequest := *sonatyperepo.NewApiPrivilegeRepositoryAdminRequest([]string{"Actions_example"}, "Format_example", "Name_example", "Repository_example") // ApiPrivilegeRepositoryAdminRequest | The privilege to create.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateRepositoryAdminPrivilege(context.Background()).ApiPrivilegeRepositoryAdminRequest(apiPrivilegeRepositoryAdminRequest).Execute()
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
 **apiPrivilegeRepositoryAdminRequest** | [**ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md) | The privilege to create. | 

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

> CreateRepositoryContentSelectorPrivilege(ctx).ApiPrivilegeRepositoryContentSelectorRequest(apiPrivilegeRepositoryContentSelectorRequest).Execute()

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
	apiPrivilegeRepositoryContentSelectorRequest := *sonatyperepo.NewApiPrivilegeRepositoryContentSelectorRequest([]string{"Actions_example"}, "ContentSelector_example", "Format_example", "Name_example", "Repository_example") // ApiPrivilegeRepositoryContentSelectorRequest | The privilege to create.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateRepositoryContentSelectorPrivilege(context.Background()).ApiPrivilegeRepositoryContentSelectorRequest(apiPrivilegeRepositoryContentSelectorRequest).Execute()
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
 **apiPrivilegeRepositoryContentSelectorRequest** | [**ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md) | The privilege to create. | 

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

> CreateRepositoryViewPrivilege(ctx).ApiPrivilegeRepositoryViewRequest(apiPrivilegeRepositoryViewRequest).Execute()

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
	apiPrivilegeRepositoryViewRequest := *sonatyperepo.NewApiPrivilegeRepositoryViewRequest([]string{"Actions_example"}, "Format_example", "Name_example", "Repository_example") // ApiPrivilegeRepositoryViewRequest | The privilege to create.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateRepositoryViewPrivilege(context.Background()).ApiPrivilegeRepositoryViewRequest(apiPrivilegeRepositoryViewRequest).Execute()
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
 **apiPrivilegeRepositoryViewRequest** | [**ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md) | The privilege to create. | 

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

> CreateScriptPrivilege(ctx).ApiPrivilegeScriptRequest(apiPrivilegeScriptRequest).Execute()

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
	apiPrivilegeScriptRequest := *sonatyperepo.NewApiPrivilegeScriptRequest([]string{"Actions_example"}, "Name_example", "ScriptName_example") // ApiPrivilegeScriptRequest | The privilege to create.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateScriptPrivilege(context.Background()).ApiPrivilegeScriptRequest(apiPrivilegeScriptRequest).Execute()
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
 **apiPrivilegeScriptRequest** | [**ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md) | The privilege to create. | 

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

> CreateWildcardPrivilege(ctx).ApiPrivilegeWildcardRequest(apiPrivilegeWildcardRequest).Execute()

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
	apiPrivilegeWildcardRequest := *sonatyperepo.NewApiPrivilegeWildcardRequest("Name_example", "Pattern_example") // ApiPrivilegeWildcardRequest | The privilege to create.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreateWildcardPrivilege(context.Background()).ApiPrivilegeWildcardRequest(apiPrivilegeWildcardRequest).Execute()
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
 **apiPrivilegeWildcardRequest** | [**ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md) | The privilege to create. | 

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


## DeleteSecurityPrivileges

> DeleteSecurityPrivileges(ctx, privilegeName).Execute()

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
	r, err := apiClient.SecurityManagementPrivilegesAPI.DeleteSecurityPrivileges(context.Background(), privilegeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.DeleteSecurityPrivileges``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteSecurityPrivilegesRequest struct via the builder pattern


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


## GetAllPrivileges

> []ApiPrivilegeRequest GetAllPrivileges(ctx).Execute()

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
	resp, r, err := apiClient.SecurityManagementPrivilegesAPI.GetAllPrivileges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.GetAllPrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPrivileges`: []ApiPrivilegeRequest
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementPrivilegesAPI.GetAllPrivileges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPrivilegesRequest struct via the builder pattern


### Return type

[**[]ApiPrivilegeRequest**](ApiPrivilegeRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityPrivileges

> ApiPrivilegeRequest GetSecurityPrivileges(ctx, privilegeName).Execute()

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
	resp, r, err := apiClient.SecurityManagementPrivilegesAPI.GetSecurityPrivileges(context.Background(), privilegeName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.GetSecurityPrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityPrivileges`: ApiPrivilegeRequest
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementPrivilegesAPI.GetSecurityPrivileges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**privilegeName** | **string** | The name of the privilege to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityPrivilegesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiPrivilegeRequest**](ApiPrivilegeRequest.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationPrivilege

> UpdateApplicationPrivilege(ctx, privilegeName).ApiPrivilegeApplicationRequest(apiPrivilegeApplicationRequest).Execute()

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
	apiPrivilegeApplicationRequest := *sonatyperepo.NewApiPrivilegeApplicationRequest([]string{"Actions_example"}, "Domain_example", "Name_example") // ApiPrivilegeApplicationRequest | The privilege to update.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateApplicationPrivilege(context.Background(), privilegeName).ApiPrivilegeApplicationRequest(apiPrivilegeApplicationRequest).Execute()
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

 **apiPrivilegeApplicationRequest** | [**ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md) | The privilege to update. | 

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

> UpdateRepositoryAdminPrivilege(ctx, privilegeName).ApiPrivilegeRepositoryAdminRequest(apiPrivilegeRepositoryAdminRequest).Execute()

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
	apiPrivilegeRepositoryAdminRequest := *sonatyperepo.NewApiPrivilegeRepositoryAdminRequest([]string{"Actions_example"}, "Format_example", "Name_example", "Repository_example") // ApiPrivilegeRepositoryAdminRequest | The privilege to update.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateRepositoryAdminPrivilege(context.Background(), privilegeName).ApiPrivilegeRepositoryAdminRequest(apiPrivilegeRepositoryAdminRequest).Execute()
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

 **apiPrivilegeRepositoryAdminRequest** | [**ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md) | The privilege to update. | 

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

> UpdateRepositoryContentSelectorPrivilege(ctx, privilegeName).ApiPrivilegeRepositoryContentSelectorRequest(apiPrivilegeRepositoryContentSelectorRequest).Execute()

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
	apiPrivilegeRepositoryContentSelectorRequest := *sonatyperepo.NewApiPrivilegeRepositoryContentSelectorRequest([]string{"Actions_example"}, "ContentSelector_example", "Format_example", "Name_example", "Repository_example") // ApiPrivilegeRepositoryContentSelectorRequest | The privilege to update.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateRepositoryContentSelectorPrivilege(context.Background(), privilegeName).ApiPrivilegeRepositoryContentSelectorRequest(apiPrivilegeRepositoryContentSelectorRequest).Execute()
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

 **apiPrivilegeRepositoryContentSelectorRequest** | [**ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md) | The privilege to update. | 

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

> UpdateRepositoryViewPrivilege(ctx, privilegeName).ApiPrivilegeRepositoryViewRequest(apiPrivilegeRepositoryViewRequest).Execute()

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
	apiPrivilegeRepositoryViewRequest := *sonatyperepo.NewApiPrivilegeRepositoryViewRequest([]string{"Actions_example"}, "Format_example", "Name_example", "Repository_example") // ApiPrivilegeRepositoryViewRequest | The privilege to update.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateRepositoryViewPrivilege(context.Background(), privilegeName).ApiPrivilegeRepositoryViewRequest(apiPrivilegeRepositoryViewRequest).Execute()
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

 **apiPrivilegeRepositoryViewRequest** | [**ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md) | The privilege to update. | 

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

> UpdateScriptPrivilege(ctx, privilegeName).ApiPrivilegeScriptRequest(apiPrivilegeScriptRequest).Execute()

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
	apiPrivilegeScriptRequest := *sonatyperepo.NewApiPrivilegeScriptRequest([]string{"Actions_example"}, "Name_example", "ScriptName_example") // ApiPrivilegeScriptRequest | The privilege to update.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateScriptPrivilege(context.Background(), privilegeName).ApiPrivilegeScriptRequest(apiPrivilegeScriptRequest).Execute()
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

 **apiPrivilegeScriptRequest** | [**ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md) | The privilege to update. | 

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

> UpdateWildcardPrivilege(ctx, privilegeName).ApiPrivilegeWildcardRequest(apiPrivilegeWildcardRequest).Execute()

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
	apiPrivilegeWildcardRequest := *sonatyperepo.NewApiPrivilegeWildcardRequest("Name_example", "Pattern_example") // ApiPrivilegeWildcardRequest | The privilege to update.

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdateWildcardPrivilege(context.Background(), privilegeName).ApiPrivilegeWildcardRequest(apiPrivilegeWildcardRequest).Execute()
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

 **apiPrivilegeWildcardRequest** | [**ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md) | The privilege to update. | 

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


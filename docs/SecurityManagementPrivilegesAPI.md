# \SecurityManagementPrivilegesAPI

All URIs are relative to *http://localhost/service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivilege**](SecurityManagementPrivilegesAPI.md#CreatePrivilege) | **Post** /v1/security/privileges/application | Create an application type privilege.
[**CreatePrivilege1**](SecurityManagementPrivilegesAPI.md#CreatePrivilege1) | **Post** /v1/security/privileges/wildcard | Create a wildcard type privilege.
[**CreatePrivilege2**](SecurityManagementPrivilegesAPI.md#CreatePrivilege2) | **Post** /v1/security/privileges/repository-content-selector | Create a repository content selector type privilege.
[**CreatePrivilege3**](SecurityManagementPrivilegesAPI.md#CreatePrivilege3) | **Post** /v1/security/privileges/repository-admin | Create a repository admin type privilege.
[**CreatePrivilege4**](SecurityManagementPrivilegesAPI.md#CreatePrivilege4) | **Post** /v1/security/privileges/repository-view | Create a repository view type privilege.
[**CreatePrivilege5**](SecurityManagementPrivilegesAPI.md#CreatePrivilege5) | **Post** /v1/security/privileges/script | Create a script type privilege.
[**DeletePrivilege**](SecurityManagementPrivilegesAPI.md#DeletePrivilege) | **Delete** /v1/security/privileges/{privilegeName} | Delete a privilege by name.
[**GetPrivilege**](SecurityManagementPrivilegesAPI.md#GetPrivilege) | **Get** /v1/security/privileges/{privilegeName} | Retrieve a privilege by name.
[**GetPrivileges**](SecurityManagementPrivilegesAPI.md#GetPrivileges) | **Get** /v1/security/privileges | Retrieve a list of privileges.
[**UpdatePrivilege**](SecurityManagementPrivilegesAPI.md#UpdatePrivilege) | **Put** /v1/security/privileges/application/{privilegeName} | Update an application type privilege.
[**UpdatePrivilege1**](SecurityManagementPrivilegesAPI.md#UpdatePrivilege1) | **Put** /v1/security/privileges/wildcard/{privilegeName} | Update a wildcard type privilege.
[**UpdatePrivilege2**](SecurityManagementPrivilegesAPI.md#UpdatePrivilege2) | **Put** /v1/security/privileges/repository-view/{privilegeName} | Update a repository view type privilege.
[**UpdatePrivilege3**](SecurityManagementPrivilegesAPI.md#UpdatePrivilege3) | **Put** /v1/security/privileges/repository-content-selector/{privilegeName} | Update a repository content selector type privilege.
[**UpdatePrivilege4**](SecurityManagementPrivilegesAPI.md#UpdatePrivilege4) | **Put** /v1/security/privileges/repository-admin/{privilegeName} | Update a repository admin type privilege.
[**UpdatePrivilege5**](SecurityManagementPrivilegesAPI.md#UpdatePrivilege5) | **Put** /v1/security/privileges/script/{privilegeName} | Update a script type privilege.



## CreatePrivilege

> CreatePrivilege(ctx).Body(body).Execute()

Create an application type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeApplicationRequest() // ApiPrivilegeApplicationRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreatePrivilege``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md) | The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivilege1

> CreatePrivilege1(ctx).Body(body).Execute()

Create a wildcard type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeWildcardRequest() // ApiPrivilegeWildcardRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege1(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreatePrivilege1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilege1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md) | The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivilege2

> CreatePrivilege2(ctx).Body(body).Execute()

Create a repository content selector type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeRepositoryContentSelectorRequest() // ApiPrivilegeRepositoryContentSelectorRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege2(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreatePrivilege2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilege2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md) | The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivilege3

> CreatePrivilege3(ctx).Body(body).Execute()

Create a repository admin type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeRepositoryAdminRequest() // ApiPrivilegeRepositoryAdminRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege3(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreatePrivilege3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilege3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md) | The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivilege4

> CreatePrivilege4(ctx).Body(body).Execute()

Create a repository view type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeRepositoryViewRequest() // ApiPrivilegeRepositoryViewRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege4(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreatePrivilege4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilege4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md) | The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivilege5

> CreatePrivilege5(ctx).Body(body).Execute()

Create a script type privilege.

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
	body := *sonatyperepo.NewApiPrivilegeScriptRequest() // ApiPrivilegeScriptRequest | The privilege to create. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.CreatePrivilege5(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.CreatePrivilege5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivilege5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md) | The privilege to create. | 

### Return type

 (empty response body)

### Authorization

No authorization required

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
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
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

No authorization required

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
	sonatyperepo "github.com/sonatype-nexus-community/nexus-repo-api-client-go"
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

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivileges

> []ApiPrivilege GetPrivileges(ctx).Execute()

Retrieve a list of privileges.

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
	resp, r, err := apiClient.SecurityManagementPrivilegesAPI.GetPrivileges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.GetPrivileges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrivileges`: []ApiPrivilege
	fmt.Fprintf(os.Stdout, "Response from `SecurityManagementPrivilegesAPI.GetPrivileges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivilegesRequest struct via the builder pattern


### Return type

[**[]ApiPrivilege**](ApiPrivilege.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege

> UpdatePrivilege(ctx, privilegeName).Body(body).Execute()

Update an application type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeApplicationRequest() // ApiPrivilegeApplicationRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdatePrivilege``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePrivilegeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeApplicationRequest**](ApiPrivilegeApplicationRequest.md) | The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege1

> UpdatePrivilege1(ctx, privilegeName).Body(body).Execute()

Update a wildcard type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeWildcardRequest() // ApiPrivilegeWildcardRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege1(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdatePrivilege1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePrivilege1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeWildcardRequest**](ApiPrivilegeWildcardRequest.md) | The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege2

> UpdatePrivilege2(ctx, privilegeName).Body(body).Execute()

Update a repository view type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeRepositoryViewRequest() // ApiPrivilegeRepositoryViewRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege2(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdatePrivilege2``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePrivilege2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeRepositoryViewRequest**](ApiPrivilegeRepositoryViewRequest.md) | The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege3

> UpdatePrivilege3(ctx, privilegeName).Body(body).Execute()

Update a repository content selector type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeRepositoryContentSelectorRequest() // ApiPrivilegeRepositoryContentSelectorRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege3(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdatePrivilege3``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePrivilege3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeRepositoryContentSelectorRequest**](ApiPrivilegeRepositoryContentSelectorRequest.md) | The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege4

> UpdatePrivilege4(ctx, privilegeName).Body(body).Execute()

Update a repository admin type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeRepositoryAdminRequest() // ApiPrivilegeRepositoryAdminRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege4(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdatePrivilege4``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePrivilege4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeRepositoryAdminRequest**](ApiPrivilegeRepositoryAdminRequest.md) | The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrivilege5

> UpdatePrivilege5(ctx, privilegeName).Body(body).Execute()

Update a script type privilege.

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
	privilegeName := "privilegeName_example" // string | The name of the privilege to update.
	body := *sonatyperepo.NewApiPrivilegeScriptRequest() // ApiPrivilegeScriptRequest | The privilege to update. (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.SecurityManagementPrivilegesAPI.UpdatePrivilege5(context.Background(), privilegeName).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityManagementPrivilegesAPI.UpdatePrivilege5``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdatePrivilege5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApiPrivilegeScriptRequest**](ApiPrivilegeScriptRequest.md) | The privilege to update. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


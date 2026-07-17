# \ManageSonatypeRepositoryFirewallConfigurationAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIqCapabilitiesTest**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#CreateIqCapabilitiesTest) | **Post** /v1/iq/capabilities/test | Get IQ Server capabilities for a specific configuration
[**CreateIqDisable**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#CreateIqDisable) | **Post** /v1/iq/disable | Disable Sonatype Repository Firewall
[**CreateIqEnable**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#CreateIqEnable) | **Post** /v1/iq/enable | Enable Sonatype Repository Firewall
[**CreateIqTestNewConnection**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#CreateIqTestNewConnection) | **Post** /v1/iq/test-new-connection | Test new Sonatype Repository Firewall connection
[**GetIqAudit**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#GetIqAudit) | **Get** /v1/iq/audit/{repositoryName} | Get audit status for the repository
[**ListIq**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#ListIq) | **Get** /v1/iq | Get Sonatype Repository Firewall configuration
[**ListIqAudit**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#ListIqAudit) | **Get** /v1/iq/audit | List repositories audit statuses.
[**ListIqCapabilities**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#ListIqCapabilities) | **Get** /v1/iq/capabilities | Get IQ Server capabilities (Firewall and Lifecycle)
[**UpdateIq**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#UpdateIq) | **Put** /v1/iq | Update Sonatype Repository Firewall configuration
[**UpdateIqAudit**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#UpdateIqAudit) | **Put** /v1/iq/audit | Manage audit
[**VerifyIqConnection**](ManageSonatypeRepositoryFirewallConfigurationAPI.md#VerifyIqConnection) | **Post** /v1/iq/verify-connection | Verify Sonatype Repository Firewall connection



## CreateIqCapabilitiesTest

> IqCapabilitiesXo CreateIqCapabilitiesTest(ctx).IqConnectionXo(iqConnectionXo).Execute()

Get IQ Server capabilities for a specific configuration

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
	iqConnectionXo := *sonatyperepo.NewIqConnectionXo("AuthenticationType_example", "Url_example") // IqConnectionXo | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqCapabilitiesTest(context.Background()).IqConnectionXo(iqConnectionXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqCapabilitiesTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIqCapabilitiesTest`: IqCapabilitiesXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqCapabilitiesTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIqCapabilitiesTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iqConnectionXo** | [**IqConnectionXo**](IqConnectionXo.md) |  | 

### Return type

[**IqCapabilitiesXo**](IqCapabilitiesXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIqDisable

> CreateIqDisable(ctx).Execute()

Disable Sonatype Repository Firewall

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
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqDisable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqDisable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIqDisableRequest struct via the builder pattern


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


## CreateIqEnable

> CreateIqEnable(ctx).Execute()

Enable Sonatype Repository Firewall

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
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqEnable(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqEnable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIqEnableRequest struct via the builder pattern


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


## CreateIqTestNewConnection

> IqConnectionVerificationXo CreateIqTestNewConnection(ctx).IqConnectionXo(iqConnectionXo).Execute()

Test new Sonatype Repository Firewall connection

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
	iqConnectionXo := *sonatyperepo.NewIqConnectionXo("AuthenticationType_example", "Url_example") // IqConnectionXo | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqTestNewConnection(context.Background()).IqConnectionXo(iqConnectionXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqTestNewConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIqTestNewConnection`: IqConnectionVerificationXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.CreateIqTestNewConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIqTestNewConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iqConnectionXo** | [**IqConnectionXo**](IqConnectionXo.md) |  | 

### Return type

[**IqConnectionVerificationXo**](IqConnectionVerificationXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIqAudit

> GetIqAudit(ctx, repositoryName).Execute()

Get audit status for the repository

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
	repositoryName := "repositoryName_example" // string | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.GetIqAudit(context.Background(), repositoryName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.GetIqAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIqAuditRequest struct via the builder pattern


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


## ListIq

> IqConnectionXo ListIq(ctx).Execute()

Get Sonatype Repository Firewall configuration

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
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.ListIq(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.ListIq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIq`: IqConnectionXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.ListIq`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIqRequest struct via the builder pattern


### Return type

[**IqConnectionXo**](IqConnectionXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIqAudit

> []IqAuditXo ListIqAudit(ctx).Execute()

List repositories audit statuses.

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
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.ListIqAudit(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.ListIqAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIqAudit`: []IqAuditXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.ListIqAudit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIqAuditRequest struct via the builder pattern


### Return type

[**[]IqAuditXo**](IqAuditXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIqCapabilities

> IqCapabilitiesXo ListIqCapabilities(ctx).Execute()

Get IQ Server capabilities (Firewall and Lifecycle)

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
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.ListIqCapabilities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.ListIqCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIqCapabilities`: IqCapabilitiesXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.ListIqCapabilities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListIqCapabilitiesRequest struct via the builder pattern


### Return type

[**IqCapabilitiesXo**](IqCapabilitiesXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIq

> IqConnectionXo UpdateIq(ctx).IqConnectionXo(iqConnectionXo).Execute()

Update Sonatype Repository Firewall configuration

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
	iqConnectionXo := *sonatyperepo.NewIqConnectionXo("AuthenticationType_example", "Url_example") // IqConnectionXo | 

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateIq(context.Background()).IqConnectionXo(iqConnectionXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateIq``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIq`: IqConnectionXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateIq`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIqRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iqConnectionXo** | [**IqConnectionXo**](IqConnectionXo.md) |  | 

### Return type

[**IqConnectionXo**](IqConnectionXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIqAudit

> UpdateIqAudit(ctx).IqAuditXo(iqAuditXo).Execute()

Manage audit

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
	iqAuditXo := *sonatyperepo.NewIqAuditXo(false, "RepositoryName_example") // IqAuditXo |  (optional)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateIqAudit(context.Background()).IqAuditXo(iqAuditXo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.UpdateIqAudit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIqAuditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iqAuditXo** | [**IqAuditXo**](IqAuditXo.md) |  | 

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


## VerifyIqConnection

> IqConnectionVerificationXo VerifyIqConnection(ctx).Execute()

Verify Sonatype Repository Firewall connection

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
	resp, r, err := apiClient.ManageSonatypeRepositoryFirewallConfigurationAPI.VerifyIqConnection(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManageSonatypeRepositoryFirewallConfigurationAPI.VerifyIqConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyIqConnection`: IqConnectionVerificationXo
	fmt.Fprintf(os.Stdout, "Response from `ManageSonatypeRepositoryFirewallConfigurationAPI.VerifyIqConnection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyIqConnectionRequest struct via the builder pattern


### Return type

[**IqConnectionVerificationXo**](IqConnectionVerificationXo.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


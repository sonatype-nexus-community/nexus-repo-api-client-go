# \UsageHistoryAPI

All URIs are relative to */service/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsageHistory**](UsageHistoryAPI.md#ListUsageHistory) | **Get** /v1/usage-history | Get usage history for sparklines



## ListUsageHistory

> ListUsageHistory(ctx).Metric(metric).Period(period).Execute()

Get usage history for sparklines



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
	metric := "metric_example" // string | Metric type: 'requests' or 'components'
	period := "period_example" // string | Period: 'daily' (last 7 days) or 'monthly' (last 12 months)

	configuration := sonatyperepo.NewConfiguration()
	apiClient := sonatyperepo.NewAPIClient(configuration)
	r, err := apiClient.UsageHistoryAPI.ListUsageHistory(context.Background()).Metric(metric).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsageHistoryAPI.ListUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metric** | **string** | Metric type: &#39;requests&#39; or &#39;components&#39; | 
 **period** | **string** | Period: &#39;daily&#39; (last 7 days) or &#39;monthly&#39; (last 12 months) | 

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


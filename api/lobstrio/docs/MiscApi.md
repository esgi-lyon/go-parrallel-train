# \MiscApi

All URIs are relative to *https://api.lobstr.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTask**](MiscApi.md#CreateTask) | **Post** /tasks | CreateTask
[**GetCluster**](MiscApi.md#GetCluster) | **Get** /clusters/{clusterId} | GetCluster
[**ListCluster**](MiscApi.md#ListCluster) | **Get** /clusters | ListCluster
[**ListResult**](MiscApi.md#ListResult) | **Get** /results | ListResult
[**ListRun**](MiscApi.md#ListRun) | **Get** /runs | ListRun
[**RunCluster**](MiscApi.md#RunCluster) | **Post** /runs | RunCluster
[**UpdateCluster**](MiscApi.md#UpdateCluster) | **Post** /clusters/{clusterId} | UpdateCluster
[**ViewRun**](MiscApi.md#ViewRun) | **Get** /runs/{runId} | ViewRun



## CreateTask

> CreateTask(ctx).CreateTaskRequest(createTaskRequest).Execute()

CreateTask

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    createTaskRequest := *openapiclient.NewCreateTaskRequest("Cluster_example", []openapiclient.Task{*openapiclient.NewTask("Url_example")}) // CreateTaskRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.CreateTask(context.Background()).CreateTaskRequest(createTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.CreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTaskRequest** | [**CreateTaskRequest**](CreateTaskRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> GetCluster(ctx, clusterId).ContentType(contentType).Execute()

GetCluster

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    contentType := "application/json" // string | 
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.GetCluster(context.Background(), clusterId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCluster

> ListCluster(ctx).Id(id).ContentType(contentType).Execute()

ListCluster

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    id := "{{clusterId}}" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.ListCluster(context.Background()).Id(id).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.ListCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** |  | 
 **contentType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResult

> ListResult(ctx).Run(run).Cluster(cluster).ContentType(contentType).Execute()

ListResult

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    run := "{{runId}}" // string | 
    cluster := "{{clusterId}}" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.ListResult(context.Background()).Run(run).Cluster(cluster).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.ListResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **run** | **string** |  | 
 **cluster** | **string** |  | 
 **contentType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRun

> ListRun(ctx).Page(page).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()

ListRun

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    page := int32(1) // int32 | 
    limit := int32(10) // int32 | 
    authorization := "Token 85740612a05ece90fb7bbbf7b507947ffd530cf8" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.ListRun(context.Background()).Page(page).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.ListRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **limit** | **int32** |  | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 

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


## RunCluster

> RunCluster(ctx).RunClusterRequest(runClusterRequest).Execute()

RunCluster

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    runClusterRequest := *openapiclient.NewRunClusterRequest("Cluster_example") // RunClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.RunCluster(context.Background()).RunClusterRequest(runClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.RunCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **runClusterRequest** | [**RunClusterRequest**](RunClusterRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> UpdateCluster(ctx, clusterId).UpdateClusterRequest(updateClusterRequest).Execute()

UpdateCluster

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    clusterId := "clusterId_example" // string | 
    updateClusterRequest := *openapiclient.NewUpdateClusterRequest("Name_example", int32(123), false, false, false, *openapiclient.NewUpdateClusterRequestParams("MaxDate_example", int32(123), int32(123), "HoursBack_example", false), "Accounts_example", "RunNotify_example") // UpdateClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.UpdateCluster(context.Background(), clusterId).UpdateClusterRequest(updateClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClusterRequest** | [**UpdateClusterRequest**](UpdateClusterRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ViewRun

> ViewRun(ctx, runId).ContentType(contentType).Execute()

ViewRun

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/esgi-lyon/go-parrallel-train/api/lobstrio"
)

func main() {
    contentType := "application/json" // string | 
    runId := "runId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.ViewRun(context.Background(), runId).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.ViewRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string** |  | 


### Return type

 (empty response body)

### Authorization

[apiKey](../README.md#apiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


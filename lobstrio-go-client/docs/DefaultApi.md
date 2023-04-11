# \DefaultApi

All URIs are relative to *https://api.lobstr.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortRun**](DefaultApi.md#AbortRun) | **Post** /runs/:id/abort | Abort Run
[**ActivateTask**](DefaultApi.md#ActivateTask) | **Post** /tasks/:id/activate | Activate Task
[**CreateAccount**](DefaultApi.md#CreateAccount) | **Post** /accounts | Create Account
[**CreateCluster**](DefaultApi.md#CreateCluster) | **Post** /clusters | Create Cluster
[**CreateSchedule**](DefaultApi.md#CreateSchedule) | **Post** /schedules | Create Schedule
[**CreateTask**](DefaultApi.md#CreateTask) | **Post** /tasks | Create Task
[**DeactivateTask**](DefaultApi.md#DeactivateTask) | **Post** /tasks/:id/deactivate | Deactivate Task
[**DeleteAccount**](DefaultApi.md#DeleteAccount) | **Delete** /accounts/:id | Delete Account
[**DeleteCluster**](DefaultApi.md#DeleteCluster) | **Delete** /clusters/:id | Delete Cluster
[**DeleteRun**](DefaultApi.md#DeleteRun) | **Delete** /runs/:id | Delete Run
[**DeleteSchedule**](DefaultApi.md#DeleteSchedule) | **Delete** /schedules/:id | Delete Schedule
[**DeleteTask**](DefaultApi.md#DeleteTask) | **Delete** /tasks/:id | Delete Task
[**DownloadRun**](DefaultApi.md#DownloadRun) | **Get** /runs/:id/download | Download Run
[**EmptyCluster**](DefaultApi.md#EmptyCluster) | **Post** /clusters/:id/empty | Empty Cluster
[**GetAccount**](DefaultApi.md#GetAccount) | **Get** /accounts/:id | Get Account
[**GetCluster**](DefaultApi.md#GetCluster) | **Get** /clusters/:id | Get Cluster
[**GetCrawler**](DefaultApi.md#GetCrawler) | **Get** /crawlers/:id | Get Crawler
[**GetCrawlerParams**](DefaultApi.md#GetCrawlerParams) | **Get** /crawlers/:id/params | Get Crawler Params
[**GetResult**](DefaultApi.md#GetResult) | **Get** /results/:id | Get Result
[**GetRun**](DefaultApi.md#GetRun) | **Get** /runs/:id | Get Run
[**GetRunStats**](DefaultApi.md#GetRunStats) | **Get** /runs/:id/stats | Get Run Stats
[**GetRunTask**](DefaultApi.md#GetRunTask) | **Get** /runtasks/:id | Get RunTask
[**GetSchedule**](DefaultApi.md#GetSchedule) | **Get** /schedules/:id | Get Schedule
[**GetTask**](DefaultApi.md#GetTask) | **Get** /tasks/:id | Get Task
[**LaunchRun**](DefaultApi.md#LaunchRun) | **Post** /runs | Launch Run
[**ListAccounts**](DefaultApi.md#ListAccounts) | **Get** /accounts?page&#x3D;{page}&amp;limit&#x3D;{limit} | List Accounts
[**ListClusters**](DefaultApi.md#ListClusters) | **Get** /clusters?page&#x3D;{page}&amp;limit&#x3D;{limit} | List Clusters
[**ListCrawlers**](DefaultApi.md#ListCrawlers) | **Get** /crawlers?page&#x3D;{page}&amp;limit&#x3D;{limit} | List Crawlers
[**ListResults**](DefaultApi.md#ListResults) | **Get** /results | List Results
[**ListRunTasks**](DefaultApi.md#ListRunTasks) | **Get** /runtasks?page&#x3D;{page}&amp;limit&#x3D;{limit} | List RunTasks
[**ListRuns**](DefaultApi.md#ListRuns) | **Get** /runs?page&#x3D;{page}&amp;limit&#x3D;{limit} | List Runs
[**ListSchedules**](DefaultApi.md#ListSchedules) | **Get** /schedules?page&#x3D;{page}&amp;limit&#x3D;{limit} | List Schedules
[**ListTasks**](DefaultApi.md#ListTasks) | **Get** /tasks?page&#x3D;{page}&amp;limit&#x3D;{limit} | List Tasks
[**RefreshAccount**](DefaultApi.md#RefreshAccount) | **Post** /accounts/:id/refresh | Refresh Account
[**UpdateAccount**](DefaultApi.md#UpdateAccount) | **Post** /accounts/:id | Update Account
[**UpdateCluster**](DefaultApi.md#UpdateCluster) | **Post** /clusters/:id | Update Cluster
[**UpdateSchedule**](DefaultApi.md#UpdateSchedule) | **Post** /schedules/:id | Update Schedule



## AbortRun

> RunAborted AbortRun(ctx).Authorization(authorization).Execute()

Abort Run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AbortRun(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AbortRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortRun`: RunAborted
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AbortRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAbortRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**RunAborted**](RunAborted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ActivateTask

> TaskActivated ActivateTask(ctx).Authorization(authorization).Execute()

Activate Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ActivateTask(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ActivateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateTask`: TaskActivated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ActivateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**TaskActivated**](TaskActivated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAccount

> AccountCreated CreateAccount(ctx).Authorization(authorization).AccountCommonParams(accountCommonParams).Execute()

Create Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    accountCommonParams := *openapiclient.NewAccountCommonParams() // AccountCommonParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateAccount(context.Background()).Authorization(authorization).AccountCommonParams(accountCommonParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccount`: AccountCreated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **accountCommonParams** | [**AccountCommonParams**](AccountCommonParams.md) |  | 

### Return type

[**AccountCreated**](AccountCreated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCluster

> ClusterCreated CreateCluster(ctx).Authorization(authorization).ClusterCommonAttribsCreation(clusterCommonAttribsCreation).Execute()

Create Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    clusterCommonAttribsCreation := *openapiclient.NewClusterCommonAttribsCreation("6fbfd5e68d3306e51350bea0232f8fa5", "e9619cf7d9a9a2da1fed1f25d32ca3ae", "my-cool-cluster", *openapiclient.NewClusterCommonParams(), "3dbbddabf0a47a9c271ddf8c19a8dc09") // ClusterCommonAttribsCreation |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateCluster(context.Background()).Authorization(authorization).ClusterCommonAttribsCreation(clusterCommonAttribsCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: ClusterCreated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **clusterCommonAttribsCreation** | [**ClusterCommonAttribsCreation**](ClusterCommonAttribsCreation.md) |  | 

### Return type

[**ClusterCreated**](ClusterCreated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSchedule

> ScheduleCreated CreateSchedule(ctx).Authorization(authorization).ScheduleCommonParams(scheduleCommonParams).Execute()

Create Schedule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    scheduleCommonParams := *openapiclient.NewScheduleCommonParams("@daily", "my-cool-schedule") // ScheduleCommonParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateSchedule(context.Background()).Authorization(authorization).ScheduleCommonParams(scheduleCommonParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSchedule`: ScheduleCreated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **scheduleCommonParams** | [**ScheduleCommonParams**](ScheduleCommonParams.md) |  | 

### Return type

[**ScheduleCreated**](ScheduleCreated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTask

> TaskCreated CreateTask(ctx).Authorization(authorization).CreateTaskRequest(createTaskRequest).Execute()

Create Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    createTaskRequest := *openapiclient.NewCreateTaskRequest("6fbfd5e68d3306e51350bea0232f8fa5") // CreateTaskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateTask(context.Background()).Authorization(authorization).CreateTaskRequest(createTaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTask`: TaskCreated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **createTaskRequest** | [**CreateTaskRequest**](CreateTaskRequest.md) |  | 

### Return type

[**TaskCreated**](TaskCreated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateTask

> TaskDeactivated DeactivateTask(ctx).Authorization(authorization).Execute()

Deactivate Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeactivateTask(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeactivateTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateTask`: TaskDeactivated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeactivateTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**TaskDeactivated**](TaskDeactivated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccount

> AccountDeleted DeleteAccount(ctx).Authorization(authorization).Execute()

Delete Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteAccount(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAccount`: AccountDeleted
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**AccountDeleted**](AccountDeleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> ClusterDeleted DeleteCluster(ctx).Authorization(authorization).Execute()

Delete Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteCluster(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCluster`: ClusterDeleted
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**ClusterDeleted**](ClusterDeleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRun

> RunDeleted DeleteRun(ctx).Authorization(authorization).Execute()

Delete Run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteRun(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRun`: RunDeleted
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**RunDeleted**](RunDeleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSchedule

> ScheduleDeleted DeleteSchedule(ctx).Authorization(authorization).Execute()

Delete Schedule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteSchedule(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSchedule`: ScheduleDeleted
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**ScheduleDeleted**](ScheduleDeleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTask

> TaskDeleted DeleteTask(ctx).Authorization(authorization).Execute()

Delete Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DeleteTask(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTask`: TaskDeleted
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**TaskDeleted**](TaskDeleted.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadRun

> RunDownloaded DownloadRun(ctx).Authorization(authorization).Execute()

Download Run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.DownloadRun(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DownloadRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DownloadRun`: RunDownloaded
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DownloadRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDownloadRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**RunDownloaded**](RunDownloaded.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmptyCluster

> EmptiedCluster EmptyCluster(ctx).Authorization(authorization).Execute()

Empty Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EmptyCluster(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EmptyCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmptyCluster`: EmptiedCluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EmptyCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmptyClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**EmptiedCluster**](EmptiedCluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx).Authorization(authorization).Execute()

Get Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetAccount(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCluster

> Cluster GetCluster(ctx).Authorization(authorization).Execute()

Get Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCluster(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrawler

> Crawler GetCrawler(ctx).Authorization(authorization).Execute()

Get Crawler



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCrawler(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCrawler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrawler`: Crawler
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCrawler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrawlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**Crawler**](Crawler.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrawlerParams

> CrawlerParams GetCrawlerParams(ctx).Authorization(authorization).Execute()

Get Crawler Params



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCrawlerParams(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCrawlerParams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrawlerParams`: CrawlerParams
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCrawlerParams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrawlerParamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**CrawlerParams**](CrawlerParams.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResult

> Result GetResult(ctx).Authorization(authorization).Execute()

Get Result



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetResult(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResult`: Result
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetResult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**Result**](Result.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRun

> Run GetRun(ctx).Authorization(authorization).Execute()

Get Run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRun(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRun`: Run
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**Run**](Run.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunStats

> RunStats GetRunStats(ctx).Authorization(authorization).Execute()

Get Run Stats



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRunStats(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRunStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunStats`: RunStats
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRunStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**RunStats**](RunStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunTask

> RunTask GetRunTask(ctx).Authorization(authorization).Execute()

Get RunTask



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetRunTask(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetRunTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRunTask`: RunTask
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetRunTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRunTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**RunTask**](RunTask.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchedule

> Schedule GetSchedule(ctx).Authorization(authorization).Execute()

Get Schedule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetSchedule(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTask

> Task GetTask(ctx).Authorization(authorization).Execute()

Get Task



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTask(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: Task
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchRun

> RunCreated LaunchRun(ctx).Authorization(authorization).LaunchRunRequest(launchRunRequest).Execute()

Launch Run



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    launchRunRequest := *openapiclient.NewLaunchRunRequest("'65b9eea6e1cc6bb9f0cd2a47751a186f'") // LaunchRunRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.LaunchRun(context.Background()).Authorization(authorization).LaunchRunRequest(launchRunRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LaunchRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LaunchRun`: RunCreated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LaunchRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLaunchRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **launchRunRequest** | [**LaunchRunRequest**](LaunchRunRequest.md) |  | 

### Return type

[**RunCreated**](RunCreated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccounts

> ListClusters200Response ListAccounts(ctx, page, limit).Authorization(authorization).Execute()

List Accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(1) // float32 | 
    limit := float32(10) // float32 |  (optional) (default to 50)
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListAccounts(context.Background(), page, limit).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccounts`: ListClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **float32** |  | 
**limit** | **float32** |  | [default to 50]

### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> ListClusters200Response ListClusters(ctx, page, limit).Authorization(authorization).Execute()

List Clusters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(1) // float32 | 
    limit := float32(10) // float32 |  (optional) (default to 50)
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListClusters(context.Background(), page, limit).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusters`: ListClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **float32** |  | 
**limit** | **float32** |  | [default to 50]

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCrawlers

> ListClusters200Response ListCrawlers(ctx, page, limit).Authorization(authorization).Execute()

List Crawlers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(1) // float32 | 
    limit := float32(10) // float32 |  (optional) (default to 50)
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListCrawlers(context.Background(), page, limit).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCrawlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCrawlers`: ListClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCrawlers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **float32** |  | 
**limit** | **float32** |  | [default to 50]

### Other Parameters

Other parameters are passed through a pointer to a apiListCrawlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResults

> ListClusters200Response ListResults(ctx).Page(page).Limit(limit).Authorization(authorization).ListResultsRequest(listResultsRequest).Execute()

List Results



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(1) // float32 | 
    limit := float32(10) // float32 |  (optional) (default to 50)
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    listResultsRequest := *openapiclient.NewListResultsRequest("'65b9eea6e1cc6bb9f0cd2a47751a186f'") // ListResultsRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListResults(context.Background()).Page(page).Limit(limit).Authorization(authorization).ListResultsRequest(listResultsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListResults`: ListClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListResults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **float32** |  | 
 **limit** | **float32** |  | [default to 50]
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **listResultsRequest** | [**ListResultsRequest**](ListResultsRequest.md) |  | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRunTasks

> ListClusters200Response ListRunTasks(ctx, page, limit).Authorization(authorization).ListRunTasksRequest(listRunTasksRequest).Execute()

List RunTasks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(1) // float32 | 
    limit := float32(10) // float32 |  (optional) (default to 50)
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    listRunTasksRequest := *openapiclient.NewListRunTasksRequest() // ListRunTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRunTasks(context.Background(), page, limit).Authorization(authorization).ListRunTasksRequest(listRunTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRunTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRunTasks`: ListClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRunTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **float32** |  | 
**limit** | **float32** |  | [default to 50]

### Other Parameters

Other parameters are passed through a pointer to a apiListRunTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **listRunTasksRequest** | [**ListRunTasksRequest**](ListRunTasksRequest.md) |  | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRuns

> ListClusters200Response ListRuns(ctx, page, limit).Authorization(authorization).ListTasksRequest(listTasksRequest).Execute()

List Runs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(1) // float32 | 
    limit := float32(10) // float32 |  (optional) (default to 50)
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    listTasksRequest := *openapiclient.NewListTasksRequest() // ListTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListRuns(context.Background(), page, limit).Authorization(authorization).ListTasksRequest(listTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRuns`: ListClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **float32** |  | 
**limit** | **float32** |  | [default to 50]

### Other Parameters

Other parameters are passed through a pointer to a apiListRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **listTasksRequest** | [**ListTasksRequest**](ListTasksRequest.md) |  | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchedules

> ListSchedules200Response ListSchedules(ctx, page, limit).Authorization(authorization).Execute()

List Schedules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    page := float32(1) // float32 | 
    limit := float32(10) // float32 |  (optional) (default to 50)
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListSchedules(context.Background(), page, limit).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListSchedules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchedules`: ListSchedules200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListSchedules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**page** | **float32** |  | 
**limit** | **float32** |  | [default to 50]

### Other Parameters

Other parameters are passed through a pointer to a apiListSchedulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**ListSchedules200Response**](ListSchedules200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> ListClusters200Response ListTasks(ctx).Authorization(authorization).ListTasksRequest(listTasksRequest).Execute()

List Tasks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    listTasksRequest := *openapiclient.NewListTasksRequest() // ListTasksRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.ListTasks(context.Background()).Authorization(authorization).ListTasksRequest(listTasksRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTasks`: ListClusters200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **listTasksRequest** | [**ListTasksRequest**](ListTasksRequest.md) |  | 

### Return type

[**ListClusters200Response**](ListClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAccount

> AccountRefreshed RefreshAccount(ctx).Authorization(authorization).Execute()

Refresh Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RefreshAccount(context.Background()).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RefreshAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RefreshAccount`: AccountRefreshed
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RefreshAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 

### Return type

[**AccountRefreshed**](AccountRefreshed.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccount

> AccountUpdated UpdateAccount(ctx).Authorization(authorization).UpdateAccountCommonParams(updateAccountCommonParams).Execute()

Update Account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    updateAccountCommonParams := *openapiclient.NewUpdateAccountCommonParams() // UpdateAccountCommonParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateAccount(context.Background()).Authorization(authorization).UpdateAccountCommonParams(updateAccountCommonParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccount`: AccountUpdated
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **updateAccountCommonParams** | [**UpdateAccountCommonParams**](UpdateAccountCommonParams.md) |  | 

### Return type

[**AccountUpdated**](AccountUpdated.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> Cluster UpdateCluster(ctx).Authorization(authorization).ClusterCommonAttribsUpdate(clusterCommonAttribsUpdate).Execute()

Update Cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    clusterCommonAttribsUpdate := *openapiclient.NewClusterCommonAttribsUpdate("6fbfd5e68d3306e51350bea0232f8fa5", "e9619cf7d9a9a2da1fed1f25d32ca3ae", false, "my-cool-cluster", *openapiclient.NewClusterCommonParams(), "3dbbddabf0a47a9c271ddf8c19a8dc09") // ClusterCommonAttribsUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateCluster(context.Background()).Authorization(authorization).ClusterCommonAttribsUpdate(clusterCommonAttribsUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **clusterCommonAttribsUpdate** | [**ClusterCommonAttribsUpdate**](ClusterCommonAttribsUpdate.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSchedule

> Schedule UpdateSchedule(ctx).Authorization(authorization).ScheduleCommonParams(scheduleCommonParams).Execute()

Update Schedule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    authorization := "Token SECRET_TOKEN" // string | e.g. Token SECRET_TOKEN (optional)
    scheduleCommonParams := *openapiclient.NewScheduleCommonParams("@daily", "my-cool-schedule") // ScheduleCommonParams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateSchedule(context.Background()).Authorization(authorization).ScheduleCommonParams(scheduleCommonParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateSchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSchedule`: Schedule
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateSchedule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | e.g. Token SECRET_TOKEN | 
 **scheduleCommonParams** | [**ScheduleCommonParams**](ScheduleCommonParams.md) |  | 

### Return type

[**Schedule**](Schedule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


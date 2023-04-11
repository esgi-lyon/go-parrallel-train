# \MiscApi

All URIs are relative to *https://api.lobstr.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Activatetask**](MiscApi.md#Activatetask) | **Post** /tasks/292461f69940b3bfa9672092f6601acb/activate | Activate task
[**CreatetaskNotadded**](MiscApi.md#CreatetaskNotadded) | **Post** /tasks | Create task (not added)
[**Getcluster**](MiscApi.md#Getcluster) | **Get** /clusters | Get cluster
[**Getresult**](MiscApi.md#Getresult) | **Get** /results/16c0e52bd30340caa107c80f8a2e21c2 | Get result
[**Listruns**](MiscApi.md#Listruns) | **Get** /runs | List runs
[**Runcluster**](MiscApi.md#Runcluster) | **Post** /runs | Run cluster
[**Updatecluster**](MiscApi.md#Updatecluster) | **Post** /clusters/9fb56bc8d697ff45d47ad9ed332e262b | Update cluster
[**Viewrun**](MiscApi.md#Viewrun) | **Get** /runs/ce545f63b4fa4a6988c929dfefea243e | View run



## Activatetask

> Activatetask(ctx).Authorization(authorization).ActivatetaskRequest(activatetaskRequest).Execute()

Activate task

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
    authorization := "Token 85740612a05ece90fb7bbbf7b507947ffd530cf8" // string | 
    activatetaskRequest := *openapiclient.NewActivatetaskRequest("Id_example", "Object_example", false) // ActivatetaskRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.Activatetask(context.Background()).Authorization(authorization).ActivatetaskRequest(activatetaskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Activatetask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivatetaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **activatetaskRequest** | [**ActivatetaskRequest**](ActivatetaskRequest.md) |  | 

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


## CreatetaskNotadded

> CreatetaskNotadded(ctx).CreatetaskNotaddedRequest(createtaskNotaddedRequest).Execute()

Create task (not added)

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
    createtaskNotaddedRequest := *openapiclient.NewCreatetaskNotaddedRequest("Cluster_example", []openapiclient.Task{*openapiclient.NewTask("Url_example")}) // CreatetaskNotaddedRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.CreatetaskNotadded(context.Background()).CreatetaskNotaddedRequest(createtaskNotaddedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.CreatetaskNotadded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatetaskNotaddedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createtaskNotaddedRequest** | [**CreatetaskNotaddedRequest**](CreatetaskNotaddedRequest.md) |  | 

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


## Getcluster

> Getcluster(ctx).Id(id).ContentType(contentType).Execute()

Get cluster

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
    id := "9fb56bc8d697ff45d47ad9ed332e262b" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.Getcluster(context.Background()).Id(id).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Getcluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetclusterRequest struct via the builder pattern


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


## Getresult

> Getresult(ctx).Authorization(authorization).ContentType(contentType).Execute()

Get result

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
    authorization := "Token 85740612a05ece90fb7bbbf7b507947ffd530cf8" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.Getresult(context.Background()).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Getresult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetresultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## Listruns

> Listruns(ctx).Page(page).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()

List runs

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
    r, err := apiClient.MiscApi.Listruns(context.Background()).Page(page).Limit(limit).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Listruns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListrunsRequest struct via the builder pattern


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


## Runcluster

> Runcluster(ctx).Authorization(authorization).RunclusterRequest(runclusterRequest).Execute()

Run cluster

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
    authorization := "Token 85740612a05ece90fb7bbbf7b507947ffd530cf8" // string | 
    runclusterRequest := *openapiclient.NewRunclusterRequest("Cluster_example") // RunclusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.Runcluster(context.Background()).Authorization(authorization).RunclusterRequest(runclusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Runcluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunclusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **runclusterRequest** | [**RunclusterRequest**](RunclusterRequest.md) |  | 

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


## Updatecluster

> Updatecluster(ctx).UpdateclusterRequest(updateclusterRequest).Execute()

Update cluster

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
    updateclusterRequest := *openapiclient.NewUpdateclusterRequest("Name_example", int32(123), false, false, false, *openapiclient.NewUpdateclusterRequestParams("MaxDate_example", int32(123), int32(123), "HoursBack_example", false), "Accounts_example", "RunNotify_example") // UpdateclusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.Updatecluster(context.Background()).UpdateclusterRequest(updateclusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Updatecluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateclusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateclusterRequest** | [**UpdateclusterRequest**](UpdateclusterRequest.md) |  | 

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


## Viewrun

> Viewrun(ctx).Authorization(authorization).ContentType(contentType).Execute()

View run

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
    authorization := "Token 85740612a05ece90fb7bbbf7b507947ffd530cf8" // string | 
    contentType := "application/json" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MiscApi.Viewrun(context.Background()).Authorization(authorization).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MiscApi.Viewrun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiViewrunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


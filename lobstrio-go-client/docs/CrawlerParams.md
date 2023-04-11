# CrawlerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**ClusterCrawlerParams**](ClusterCrawlerParams.md) |  | 
**Task** | [**TaskCrawlerParams**](TaskCrawlerParams.md) |  | 

## Methods

### NewCrawlerParams

`func NewCrawlerParams(cluster ClusterCrawlerParams, task TaskCrawlerParams, ) *CrawlerParams`

NewCrawlerParams instantiates a new CrawlerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlerParamsWithDefaults

`func NewCrawlerParamsWithDefaults() *CrawlerParams`

NewCrawlerParamsWithDefaults instantiates a new CrawlerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *CrawlerParams) GetCluster() ClusterCrawlerParams`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CrawlerParams) GetClusterOk() (*ClusterCrawlerParams, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CrawlerParams) SetCluster(v ClusterCrawlerParams)`

SetCluster sets Cluster field to given value.


### GetTask

`func (o *CrawlerParams) GetTask() TaskCrawlerParams`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *CrawlerParams) GetTaskOk() (*TaskCrawlerParams, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *CrawlerParams) SetTask(v TaskCrawlerParams)`

SetTask sets Task field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



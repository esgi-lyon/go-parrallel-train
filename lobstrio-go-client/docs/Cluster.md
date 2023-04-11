# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] 
**Account** | **string** |  | 
**Concurrency** | **float32** |  | 
**CreatedAt** | **string** |  | 
**Crawler** | **string** |  | 
**IsActive** | **bool** |  | 
**Name** | **string** |  | 
**Params** | [**ClusterCommonParams**](ClusterCommonParams.md) |  | 
**Schedule** | **string** |  | 
**TotalRuns** | **float32** |  | 

## Methods

### NewCluster

`func NewCluster(id string, account string, concurrency float32, createdAt string, crawler string, isActive bool, name string, params ClusterCommonParams, schedule string, totalRuns float32, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Cluster) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Cluster) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Cluster) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Cluster) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetAccount

`func (o *Cluster) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Cluster) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Cluster) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetConcurrency

`func (o *Cluster) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *Cluster) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *Cluster) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.


### GetCreatedAt

`func (o *Cluster) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Cluster) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCrawler

`func (o *Cluster) GetCrawler() string`

GetCrawler returns the Crawler field if non-nil, zero value otherwise.

### GetCrawlerOk

`func (o *Cluster) GetCrawlerOk() (*string, bool)`

GetCrawlerOk returns a tuple with the Crawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawler

`func (o *Cluster) SetCrawler(v string)`

SetCrawler sets Crawler field to given value.


### GetIsActive

`func (o *Cluster) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Cluster) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Cluster) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *Cluster) GetParams() ClusterCommonParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Cluster) GetParamsOk() (*ClusterCommonParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Cluster) SetParams(v ClusterCommonParams)`

SetParams sets Params field to given value.


### GetSchedule

`func (o *Cluster) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Cluster) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Cluster) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.


### GetTotalRuns

`func (o *Cluster) GetTotalRuns() float32`

GetTotalRuns returns the TotalRuns field if non-nil, zero value otherwise.

### GetTotalRunsOk

`func (o *Cluster) GetTotalRunsOk() (*float32, bool)`

GetTotalRunsOk returns a tuple with the TotalRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRuns

`func (o *Cluster) SetTotalRuns(v float32)`

SetTotalRuns sets TotalRuns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



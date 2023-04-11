# ClusterCommonAttribsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** |  | 
**Concurrency** | Pointer to **float32** |  | [optional] 
**Crawler** | **string** |  | 
**IsActive** | **bool** |  | 
**Name** | **string** |  | 
**Params** | [**ClusterCommonParams**](ClusterCommonParams.md) |  | 
**Schedule** | **string** |  | 

## Methods

### NewClusterCommonAttribsUpdate

`func NewClusterCommonAttribsUpdate(account string, crawler string, isActive bool, name string, params ClusterCommonParams, schedule string, ) *ClusterCommonAttribsUpdate`

NewClusterCommonAttribsUpdate instantiates a new ClusterCommonAttribsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCommonAttribsUpdateWithDefaults

`func NewClusterCommonAttribsUpdateWithDefaults() *ClusterCommonAttribsUpdate`

NewClusterCommonAttribsUpdateWithDefaults instantiates a new ClusterCommonAttribsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ClusterCommonAttribsUpdate) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterCommonAttribsUpdate) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterCommonAttribsUpdate) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetConcurrency

`func (o *ClusterCommonAttribsUpdate) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *ClusterCommonAttribsUpdate) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *ClusterCommonAttribsUpdate) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *ClusterCommonAttribsUpdate) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetCrawler

`func (o *ClusterCommonAttribsUpdate) GetCrawler() string`

GetCrawler returns the Crawler field if non-nil, zero value otherwise.

### GetCrawlerOk

`func (o *ClusterCommonAttribsUpdate) GetCrawlerOk() (*string, bool)`

GetCrawlerOk returns a tuple with the Crawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawler

`func (o *ClusterCommonAttribsUpdate) SetCrawler(v string)`

SetCrawler sets Crawler field to given value.


### GetIsActive

`func (o *ClusterCommonAttribsUpdate) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ClusterCommonAttribsUpdate) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ClusterCommonAttribsUpdate) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetName

`func (o *ClusterCommonAttribsUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCommonAttribsUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCommonAttribsUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *ClusterCommonAttribsUpdate) GetParams() ClusterCommonParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ClusterCommonAttribsUpdate) GetParamsOk() (*ClusterCommonParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ClusterCommonAttribsUpdate) SetParams(v ClusterCommonParams)`

SetParams sets Params field to given value.


### GetSchedule

`func (o *ClusterCommonAttribsUpdate) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ClusterCommonAttribsUpdate) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ClusterCommonAttribsUpdate) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ClusterCommonAttribsCreation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | **string** |  | 
**Concurrency** | Pointer to **float32** |  | [optional] 
**Crawler** | **string** |  | 
**Name** | **string** |  | 
**Params** | [**ClusterCommonParams**](ClusterCommonParams.md) |  | 
**Schedule** | **string** |  | 

## Methods

### NewClusterCommonAttribsCreation

`func NewClusterCommonAttribsCreation(account string, crawler string, name string, params ClusterCommonParams, schedule string, ) *ClusterCommonAttribsCreation`

NewClusterCommonAttribsCreation instantiates a new ClusterCommonAttribsCreation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCommonAttribsCreationWithDefaults

`func NewClusterCommonAttribsCreationWithDefaults() *ClusterCommonAttribsCreation`

NewClusterCommonAttribsCreationWithDefaults instantiates a new ClusterCommonAttribsCreation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ClusterCommonAttribsCreation) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterCommonAttribsCreation) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterCommonAttribsCreation) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetConcurrency

`func (o *ClusterCommonAttribsCreation) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *ClusterCommonAttribsCreation) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *ClusterCommonAttribsCreation) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.

### HasConcurrency

`func (o *ClusterCommonAttribsCreation) HasConcurrency() bool`

HasConcurrency returns a boolean if a field has been set.

### GetCrawler

`func (o *ClusterCommonAttribsCreation) GetCrawler() string`

GetCrawler returns the Crawler field if non-nil, zero value otherwise.

### GetCrawlerOk

`func (o *ClusterCommonAttribsCreation) GetCrawlerOk() (*string, bool)`

GetCrawlerOk returns a tuple with the Crawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawler

`func (o *ClusterCommonAttribsCreation) SetCrawler(v string)`

SetCrawler sets Crawler field to given value.


### GetName

`func (o *ClusterCommonAttribsCreation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCommonAttribsCreation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCommonAttribsCreation) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *ClusterCommonAttribsCreation) GetParams() ClusterCommonParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ClusterCommonAttribsCreation) GetParamsOk() (*ClusterCommonParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ClusterCommonAttribsCreation) SetParams(v ClusterCommonParams)`

SetParams sets Params field to given value.


### GetSchedule

`func (o *ClusterCommonAttribsCreation) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ClusterCommonAttribsCreation) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ClusterCommonAttribsCreation) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



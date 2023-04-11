# ClusterCreated

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

## Methods

### NewClusterCreated

`func NewClusterCreated(id string, account string, concurrency float32, createdAt string, crawler string, isActive bool, name string, params ClusterCommonParams, schedule string, ) *ClusterCreated`

NewClusterCreated instantiates a new ClusterCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreatedWithDefaults

`func NewClusterCreatedWithDefaults() *ClusterCreated`

NewClusterCreatedWithDefaults instantiates a new ClusterCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterCreated) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ClusterCreated) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ClusterCreated) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ClusterCreated) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ClusterCreated) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetAccount

`func (o *ClusterCreated) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterCreated) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterCreated) SetAccount(v string)`

SetAccount sets Account field to given value.


### GetConcurrency

`func (o *ClusterCreated) GetConcurrency() float32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *ClusterCreated) GetConcurrencyOk() (*float32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *ClusterCreated) SetConcurrency(v float32)`

SetConcurrency sets Concurrency field to given value.


### GetCreatedAt

`func (o *ClusterCreated) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterCreated) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterCreated) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCrawler

`func (o *ClusterCreated) GetCrawler() string`

GetCrawler returns the Crawler field if non-nil, zero value otherwise.

### GetCrawlerOk

`func (o *ClusterCreated) GetCrawlerOk() (*string, bool)`

GetCrawlerOk returns a tuple with the Crawler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrawler

`func (o *ClusterCreated) SetCrawler(v string)`

SetCrawler sets Crawler field to given value.


### GetIsActive

`func (o *ClusterCreated) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ClusterCreated) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ClusterCreated) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetName

`func (o *ClusterCreated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCreated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCreated) SetName(v string)`

SetName sets Name field to given value.


### GetParams

`func (o *ClusterCreated) GetParams() ClusterCommonParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ClusterCreated) GetParamsOk() (*ClusterCommonParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ClusterCreated) SetParams(v ClusterCommonParams)`

SetParams sets Params field to given value.


### GetSchedule

`func (o *ClusterCreated) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ClusterCreated) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ClusterCreated) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



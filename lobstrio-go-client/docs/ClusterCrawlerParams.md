# ClusterCrawlerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AccountCrawlerParams**](AccountCrawlerParams.md) |  | 
**Dropcontact** | Pointer to [**ClusterCrawlerParamsDropcontact**](ClusterCrawlerParamsDropcontact.md) |  | [optional] 

## Methods

### NewClusterCrawlerParams

`func NewClusterCrawlerParams(account AccountCrawlerParams, ) *ClusterCrawlerParams`

NewClusterCrawlerParams instantiates a new ClusterCrawlerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCrawlerParamsWithDefaults

`func NewClusterCrawlerParamsWithDefaults() *ClusterCrawlerParams`

NewClusterCrawlerParamsWithDefaults instantiates a new ClusterCrawlerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ClusterCrawlerParams) GetAccount() AccountCrawlerParams`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ClusterCrawlerParams) GetAccountOk() (*AccountCrawlerParams, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ClusterCrawlerParams) SetAccount(v AccountCrawlerParams)`

SetAccount sets Account field to given value.


### GetDropcontact

`func (o *ClusterCrawlerParams) GetDropcontact() ClusterCrawlerParamsDropcontact`

GetDropcontact returns the Dropcontact field if non-nil, zero value otherwise.

### GetDropcontactOk

`func (o *ClusterCrawlerParams) GetDropcontactOk() (*ClusterCrawlerParamsDropcontact, bool)`

GetDropcontactOk returns a tuple with the Dropcontact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropcontact

`func (o *ClusterCrawlerParams) SetDropcontact(v ClusterCrawlerParamsDropcontact)`

SetDropcontact sets Dropcontact field to given value.

### HasDropcontact

`func (o *ClusterCrawlerParams) HasDropcontact() bool`

HasDropcontact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



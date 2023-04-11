# ListResultsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Run** | Pointer to **string** |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 

## Methods

### NewListResultsRequest

`func NewListResultsRequest(cluster string, ) *ListResultsRequest`

NewListResultsRequest instantiates a new ListResultsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListResultsRequestWithDefaults

`func NewListResultsRequestWithDefaults() *ListResultsRequest`

NewListResultsRequestWithDefaults instantiates a new ListResultsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *ListResultsRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ListResultsRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ListResultsRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetRun

`func (o *ListResultsRequest) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *ListResultsRequest) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *ListResultsRequest) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *ListResultsRequest) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetTask

`func (o *ListResultsRequest) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ListResultsRequest) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ListResultsRequest) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *ListResultsRequest) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



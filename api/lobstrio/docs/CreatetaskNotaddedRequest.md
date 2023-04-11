# CreatetasknotaddedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Tasks** | [**[]Task**](Task.md) |  | 

## Methods

### NewCreatetasknotaddedRequest

`func NewCreatetasknotaddedRequest(cluster string, tasks []Task, ) *CreatetasknotaddedRequest`

NewCreatetasknotaddedRequest instantiates a new CreatetasknotaddedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatetasknotaddedRequestWithDefaults

`func NewCreatetasknotaddedRequestWithDefaults() *CreatetasknotaddedRequest`

NewCreatetasknotaddedRequestWithDefaults instantiates a new CreatetasknotaddedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *CreatetasknotaddedRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CreatetasknotaddedRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CreatetasknotaddedRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetTasks

`func (o *CreatetasknotaddedRequest) GetTasks() []Task`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *CreatetasknotaddedRequest) GetTasksOk() (*[]Task, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *CreatetasknotaddedRequest) SetTasks(v []Task)`

SetTasks sets Tasks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateTaskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** |  | 
**Tasks** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewCreateTaskRequest

`func NewCreateTaskRequest(cluster string, ) *CreateTaskRequest`

NewCreateTaskRequest instantiates a new CreateTaskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTaskRequestWithDefaults

`func NewCreateTaskRequestWithDefaults() *CreateTaskRequest`

NewCreateTaskRequestWithDefaults instantiates a new CreateTaskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *CreateTaskRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *CreateTaskRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *CreateTaskRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetTasks

`func (o *CreateTaskRequest) GetTasks() []interface{}`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *CreateTaskRequest) GetTasksOk() (*[]interface{}, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *CreateTaskRequest) SetTasks(v []interface{})`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *CreateTaskRequest) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



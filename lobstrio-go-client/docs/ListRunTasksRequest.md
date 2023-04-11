# ListRunTasksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Run** | Pointer to **string** |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 

## Methods

### NewListRunTasksRequest

`func NewListRunTasksRequest() *ListRunTasksRequest`

NewListRunTasksRequest instantiates a new ListRunTasksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRunTasksRequestWithDefaults

`func NewListRunTasksRequestWithDefaults() *ListRunTasksRequest`

NewListRunTasksRequestWithDefaults instantiates a new ListRunTasksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRun

`func (o *ListRunTasksRequest) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *ListRunTasksRequest) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *ListRunTasksRequest) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *ListRunTasksRequest) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetTask

`func (o *ListRunTasksRequest) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ListRunTasksRequest) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ListRunTasksRequest) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *ListRunTasksRequest) HasTask() bool`

HasTask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



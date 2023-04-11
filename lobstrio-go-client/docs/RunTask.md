# RunTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**EndedAt** | Pointer to **string** | date of creation (ISO8601 format)&#x60; | [optional] 
**IsDone** | Pointer to **bool** |  | [optional] 
**IsError** | Pointer to **string** |  | [optional] 
**LastPage** | Pointer to **float32** |  | [optional] 
**LastResult** | Pointer to **float32** |  | [optional] 
**Task** | Pointer to **string** |  | [optional] 
**Run** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 
**TotalPages** | Pointer to **float32** |  | [optional] 
**TotalResults** | Pointer to **float32** |  | [optional] 

## Methods

### NewRunTask

`func NewRunTask() *RunTask`

NewRunTask instantiates a new RunTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunTaskWithDefaults

`func NewRunTaskWithDefaults() *RunTask`

NewRunTaskWithDefaults instantiates a new RunTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunTask) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunTask) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunTask) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RunTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *RunTask) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RunTask) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RunTask) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RunTask) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RunTask) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RunTask) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RunTask) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RunTask) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDuration

`func (o *RunTask) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RunTask) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RunTask) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RunTask) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndedAt

`func (o *RunTask) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *RunTask) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *RunTask) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *RunTask) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetIsDone

`func (o *RunTask) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *RunTask) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *RunTask) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.

### HasIsDone

`func (o *RunTask) HasIsDone() bool`

HasIsDone returns a boolean if a field has been set.

### GetIsError

`func (o *RunTask) GetIsError() string`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *RunTask) GetIsErrorOk() (*string, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *RunTask) SetIsError(v string)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *RunTask) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetLastPage

`func (o *RunTask) GetLastPage() float32`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *RunTask) GetLastPageOk() (*float32, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *RunTask) SetLastPage(v float32)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *RunTask) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetLastResult

`func (o *RunTask) GetLastResult() float32`

GetLastResult returns the LastResult field if non-nil, zero value otherwise.

### GetLastResultOk

`func (o *RunTask) GetLastResultOk() (*float32, bool)`

GetLastResultOk returns a tuple with the LastResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResult

`func (o *RunTask) SetLastResult(v float32)`

SetLastResult sets LastResult field to given value.

### HasLastResult

`func (o *RunTask) HasLastResult() bool`

HasLastResult returns a boolean if a field has been set.

### GetTask

`func (o *RunTask) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *RunTask) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *RunTask) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *RunTask) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetRun

`func (o *RunTask) GetRun() string`

GetRun returns the Run field if non-nil, zero value otherwise.

### GetRunOk

`func (o *RunTask) GetRunOk() (*string, bool)`

GetRunOk returns a tuple with the Run field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRun

`func (o *RunTask) SetRun(v string)`

SetRun sets Run field to given value.

### HasRun

`func (o *RunTask) HasRun() bool`

HasRun returns a boolean if a field has been set.

### GetStartedAt

`func (o *RunTask) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RunTask) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RunTask) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RunTask) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetTotalPages

`func (o *RunTask) GetTotalPages() float32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *RunTask) GetTotalPagesOk() (*float32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *RunTask) SetTotalPages(v float32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *RunTask) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalResults

`func (o *RunTask) GetTotalResults() float32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RunTask) GetTotalResultsOk() (*float32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RunTask) SetTotalResults(v float32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *RunTask) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



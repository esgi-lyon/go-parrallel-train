# RunStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id | 
**Object** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Eta** | **string** |  | 
**IsDone** | **bool** |  | 
**PercentDone** | **string** |  | 
**StartedAt** | **string** |  | 
**TotalRequests** | **float32** |  | 
**TotalResults** | **float32** |  | 
**TotalTasks** | **float32** |  | 
**TotalTasksDone** | **float32** |  | 
**TotalTasksLeft** | **float32** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewRunStats

`func NewRunStats(id string, eta string, isDone bool, percentDone string, startedAt string, totalRequests float32, totalResults float32, totalTasks float32, totalTasksDone float32, totalTasksLeft float32, updatedAt string, ) *RunStats`

NewRunStats instantiates a new RunStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunStatsWithDefaults

`func NewRunStatsWithDefaults() *RunStats`

NewRunStatsWithDefaults instantiates a new RunStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunStats) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *RunStats) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RunStats) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RunStats) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RunStats) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDuration

`func (o *RunStats) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RunStats) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RunStats) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RunStats) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndedAt

`func (o *RunStats) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *RunStats) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *RunStats) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *RunStats) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetEta

`func (o *RunStats) GetEta() string`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *RunStats) GetEtaOk() (*string, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *RunStats) SetEta(v string)`

SetEta sets Eta field to given value.


### GetIsDone

`func (o *RunStats) GetIsDone() bool`

GetIsDone returns the IsDone field if non-nil, zero value otherwise.

### GetIsDoneOk

`func (o *RunStats) GetIsDoneOk() (*bool, bool)`

GetIsDoneOk returns a tuple with the IsDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDone

`func (o *RunStats) SetIsDone(v bool)`

SetIsDone sets IsDone field to given value.


### GetPercentDone

`func (o *RunStats) GetPercentDone() string`

GetPercentDone returns the PercentDone field if non-nil, zero value otherwise.

### GetPercentDoneOk

`func (o *RunStats) GetPercentDoneOk() (*string, bool)`

GetPercentDoneOk returns a tuple with the PercentDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentDone

`func (o *RunStats) SetPercentDone(v string)`

SetPercentDone sets PercentDone field to given value.


### GetStartedAt

`func (o *RunStats) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RunStats) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RunStats) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetTotalRequests

`func (o *RunStats) GetTotalRequests() float32`

GetTotalRequests returns the TotalRequests field if non-nil, zero value otherwise.

### GetTotalRequestsOk

`func (o *RunStats) GetTotalRequestsOk() (*float32, bool)`

GetTotalRequestsOk returns a tuple with the TotalRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRequests

`func (o *RunStats) SetTotalRequests(v float32)`

SetTotalRequests sets TotalRequests field to given value.


### GetTotalResults

`func (o *RunStats) GetTotalResults() float32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *RunStats) GetTotalResultsOk() (*float32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *RunStats) SetTotalResults(v float32)`

SetTotalResults sets TotalResults field to given value.


### GetTotalTasks

`func (o *RunStats) GetTotalTasks() float32`

GetTotalTasks returns the TotalTasks field if non-nil, zero value otherwise.

### GetTotalTasksOk

`func (o *RunStats) GetTotalTasksOk() (*float32, bool)`

GetTotalTasksOk returns a tuple with the TotalTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTasks

`func (o *RunStats) SetTotalTasks(v float32)`

SetTotalTasks sets TotalTasks field to given value.


### GetTotalTasksDone

`func (o *RunStats) GetTotalTasksDone() float32`

GetTotalTasksDone returns the TotalTasksDone field if non-nil, zero value otherwise.

### GetTotalTasksDoneOk

`func (o *RunStats) GetTotalTasksDoneOk() (*float32, bool)`

GetTotalTasksDoneOk returns a tuple with the TotalTasksDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTasksDone

`func (o *RunStats) SetTotalTasksDone(v float32)`

SetTotalTasksDone sets TotalTasksDone field to given value.


### GetTotalTasksLeft

`func (o *RunStats) GetTotalTasksLeft() float32`

GetTotalTasksLeft returns the TotalTasksLeft field if non-nil, zero value otherwise.

### GetTotalTasksLeftOk

`func (o *RunStats) GetTotalTasksLeftOk() (*float32, bool)`

GetTotalTasksLeftOk returns a tuple with the TotalTasksLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTasksLeft

`func (o *RunStats) SetTotalTasksLeft(v float32)`

SetTotalTasksLeft sets TotalTasksLeft field to given value.


### GetUpdatedAt

`func (o *RunStats) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RunStats) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RunStats) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



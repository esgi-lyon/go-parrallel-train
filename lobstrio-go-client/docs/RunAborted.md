# RunAborted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id | 
**CreatedAt** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Origin** | **string** |  | 
**EndedAt** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**StatusReason** | Pointer to **string** |  | [optional] 

## Methods

### NewRunAborted

`func NewRunAborted(id string, origin string, status string, ) *RunAborted`

NewRunAborted instantiates a new RunAborted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunAbortedWithDefaults

`func NewRunAbortedWithDefaults() *RunAborted`

NewRunAbortedWithDefaults instantiates a new RunAborted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunAborted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunAborted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunAborted) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *RunAborted) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RunAborted) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RunAborted) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RunAborted) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDuration

`func (o *RunAborted) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RunAborted) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RunAborted) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *RunAborted) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetObject

`func (o *RunAborted) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RunAborted) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RunAborted) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RunAborted) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOrigin

`func (o *RunAborted) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *RunAborted) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *RunAborted) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetEndedAt

`func (o *RunAborted) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *RunAborted) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *RunAborted) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *RunAborted) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *RunAborted) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RunAborted) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RunAborted) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RunAborted) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *RunAborted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunAborted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunAborted) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusReason

`func (o *RunAborted) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *RunAborted) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *RunAborted) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *RunAborted) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



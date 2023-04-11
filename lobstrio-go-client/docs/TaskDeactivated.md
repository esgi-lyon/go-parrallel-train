# TaskDeactivated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** | date of creation (ISO8601 format) | 
**IsActive** | **bool** |  | 
**Params** | [**TaskCommonParams**](TaskCommonParams.md) |  | 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskDeactivated

`func NewTaskDeactivated(id string, createdAt string, isActive bool, params TaskCommonParams, ) *TaskDeactivated`

NewTaskDeactivated instantiates a new TaskDeactivated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskDeactivatedWithDefaults

`func NewTaskDeactivatedWithDefaults() *TaskDeactivated`

NewTaskDeactivatedWithDefaults instantiates a new TaskDeactivated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaskDeactivated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskDeactivated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskDeactivated) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *TaskDeactivated) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *TaskDeactivated) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *TaskDeactivated) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *TaskDeactivated) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskDeactivated) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskDeactivated) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskDeactivated) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsActive

`func (o *TaskDeactivated) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *TaskDeactivated) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *TaskDeactivated) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetParams

`func (o *TaskDeactivated) GetParams() TaskCommonParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TaskDeactivated) GetParamsOk() (*TaskCommonParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TaskDeactivated) SetParams(v TaskCommonParams)`

SetParams sets Params field to given value.


### GetStatus

`func (o *TaskDeactivated) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskDeactivated) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskDeactivated) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskDeactivated) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



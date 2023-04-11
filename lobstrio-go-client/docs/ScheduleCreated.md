# ScheduleCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id | 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 
**CronExpression** | **string** | frequency of execution | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewScheduleCreated

`func NewScheduleCreated(id string, cronExpression string, name string, ) *ScheduleCreated`

NewScheduleCreated instantiates a new ScheduleCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCreatedWithDefaults

`func NewScheduleCreatedWithDefaults() *ScheduleCreated`

NewScheduleCreatedWithDefaults instantiates a new ScheduleCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleCreated) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ScheduleCreated) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ScheduleCreated) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ScheduleCreated) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ScheduleCreated) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ScheduleCreated) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ScheduleCreated) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ScheduleCreated) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ScheduleCreated) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCronExpression

`func (o *ScheduleCreated) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ScheduleCreated) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ScheduleCreated) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetDescription

`func (o *ScheduleCreated) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleCreated) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleCreated) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleCreated) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ScheduleCreated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleCreated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleCreated) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



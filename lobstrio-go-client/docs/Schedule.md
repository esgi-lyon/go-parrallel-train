# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id | 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 
**CronExpression** | **string** | frequency of execution | 
**Description** | Pointer to **string** |  | [optional] 
**LastRunAt** | Pointer to **string** | date of last run (ISO8601 format) | [optional] 
**Name** | **string** |  | 
**NextRunAt** | Pointer to **string** | date of next run (ISO8601 format) | [optional] 

## Methods

### NewSchedule

`func NewSchedule(id string, cronExpression string, name string, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schedule) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Schedule) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Schedule) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Schedule) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Schedule) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Schedule) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Schedule) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Schedule) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Schedule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCronExpression

`func (o *Schedule) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *Schedule) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *Schedule) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetDescription

`func (o *Schedule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Schedule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Schedule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Schedule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLastRunAt

`func (o *Schedule) GetLastRunAt() string`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *Schedule) GetLastRunAtOk() (*string, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *Schedule) SetLastRunAt(v string)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *Schedule) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### GetName

`func (o *Schedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Schedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Schedule) SetName(v string)`

SetName sets Name field to given value.


### GetNextRunAt

`func (o *Schedule) GetNextRunAt() string`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *Schedule) GetNextRunAtOk() (*string, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *Schedule) SetNextRunAt(v string)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *Schedule) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



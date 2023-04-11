# ScheduleCommonParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewScheduleCommonParams

`func NewScheduleCommonParams(cronExpression string, name string, ) *ScheduleCommonParams`

NewScheduleCommonParams instantiates a new ScheduleCommonParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleCommonParamsWithDefaults

`func NewScheduleCommonParamsWithDefaults() *ScheduleCommonParams`

NewScheduleCommonParamsWithDefaults instantiates a new ScheduleCommonParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *ScheduleCommonParams) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *ScheduleCommonParams) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *ScheduleCommonParams) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.


### GetDescription

`func (o *ScheduleCommonParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleCommonParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleCommonParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleCommonParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ScheduleCommonParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleCommonParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleCommonParams) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



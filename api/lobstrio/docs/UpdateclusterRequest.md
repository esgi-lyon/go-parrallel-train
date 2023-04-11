# UpdateclusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Concurrency** | **int32** |  | 
**ExportUniqueResults** | **bool** |  | 
**NoLineBreaks** | **bool** |  | 
**ToComplete** | **bool** |  | 
**Params** | [**UpdateclusterRequestParams**](UpdateclusterRequestParams.md) |  | 
**Accounts** | **NullableString** |  | 
**RunNotify** | **string** |  | 

## Methods

### NewUpdateclusterRequest

`func NewUpdateclusterRequest(name string, concurrency int32, exportUniqueResults bool, noLineBreaks bool, toComplete bool, params UpdateclusterRequestParams, accounts NullableString, runNotify string, ) *UpdateclusterRequest`

NewUpdateclusterRequest instantiates a new UpdateclusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateclusterRequestWithDefaults

`func NewUpdateclusterRequestWithDefaults() *UpdateclusterRequest`

NewUpdateclusterRequestWithDefaults instantiates a new UpdateclusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateclusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateclusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateclusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConcurrency

`func (o *UpdateclusterRequest) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *UpdateclusterRequest) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *UpdateclusterRequest) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.


### GetExportUniqueResults

`func (o *UpdateclusterRequest) GetExportUniqueResults() bool`

GetExportUniqueResults returns the ExportUniqueResults field if non-nil, zero value otherwise.

### GetExportUniqueResultsOk

`func (o *UpdateclusterRequest) GetExportUniqueResultsOk() (*bool, bool)`

GetExportUniqueResultsOk returns a tuple with the ExportUniqueResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUniqueResults

`func (o *UpdateclusterRequest) SetExportUniqueResults(v bool)`

SetExportUniqueResults sets ExportUniqueResults field to given value.


### GetNoLineBreaks

`func (o *UpdateclusterRequest) GetNoLineBreaks() bool`

GetNoLineBreaks returns the NoLineBreaks field if non-nil, zero value otherwise.

### GetNoLineBreaksOk

`func (o *UpdateclusterRequest) GetNoLineBreaksOk() (*bool, bool)`

GetNoLineBreaksOk returns a tuple with the NoLineBreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLineBreaks

`func (o *UpdateclusterRequest) SetNoLineBreaks(v bool)`

SetNoLineBreaks sets NoLineBreaks field to given value.


### GetToComplete

`func (o *UpdateclusterRequest) GetToComplete() bool`

GetToComplete returns the ToComplete field if non-nil, zero value otherwise.

### GetToCompleteOk

`func (o *UpdateclusterRequest) GetToCompleteOk() (*bool, bool)`

GetToCompleteOk returns a tuple with the ToComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToComplete

`func (o *UpdateclusterRequest) SetToComplete(v bool)`

SetToComplete sets ToComplete field to given value.


### GetParams

`func (o *UpdateclusterRequest) GetParams() UpdateclusterRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *UpdateclusterRequest) GetParamsOk() (*UpdateclusterRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *UpdateclusterRequest) SetParams(v UpdateclusterRequestParams)`

SetParams sets Params field to given value.


### GetAccounts

`func (o *UpdateclusterRequest) GetAccounts() string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdateclusterRequest) GetAccountsOk() (*string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdateclusterRequest) SetAccounts(v string)`

SetAccounts sets Accounts field to given value.


### SetAccountsNil

`func (o *UpdateclusterRequest) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *UpdateclusterRequest) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil
### GetRunNotify

`func (o *UpdateclusterRequest) GetRunNotify() string`

GetRunNotify returns the RunNotify field if non-nil, zero value otherwise.

### GetRunNotifyOk

`func (o *UpdateclusterRequest) GetRunNotifyOk() (*string, bool)`

GetRunNotifyOk returns a tuple with the RunNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNotify

`func (o *UpdateclusterRequest) SetRunNotify(v string)`

SetRunNotify sets RunNotify field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Concurrency** | **int32** |  | 
**ExportUniqueResults** | **bool** |  | 
**NoLineBreaks** | **bool** |  | 
**ToComplete** | **bool** |  | 
**Params** | [**UpdateClusterRequestParams**](UpdateClusterRequestParams.md) |  | 
**Accounts** | **NullableString** |  | 
**RunNotify** | **string** |  | 

## Methods

### NewUpdateClusterRequest

`func NewUpdateClusterRequest(name string, concurrency int32, exportUniqueResults bool, noLineBreaks bool, toComplete bool, params UpdateClusterRequestParams, accounts NullableString, runNotify string, ) *UpdateClusterRequest`

NewUpdateClusterRequest instantiates a new UpdateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterRequestWithDefaults

`func NewUpdateClusterRequestWithDefaults() *UpdateClusterRequest`

NewUpdateClusterRequestWithDefaults instantiates a new UpdateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConcurrency

`func (o *UpdateClusterRequest) GetConcurrency() int32`

GetConcurrency returns the Concurrency field if non-nil, zero value otherwise.

### GetConcurrencyOk

`func (o *UpdateClusterRequest) GetConcurrencyOk() (*int32, bool)`

GetConcurrencyOk returns a tuple with the Concurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcurrency

`func (o *UpdateClusterRequest) SetConcurrency(v int32)`

SetConcurrency sets Concurrency field to given value.


### GetExportUniqueResults

`func (o *UpdateClusterRequest) GetExportUniqueResults() bool`

GetExportUniqueResults returns the ExportUniqueResults field if non-nil, zero value otherwise.

### GetExportUniqueResultsOk

`func (o *UpdateClusterRequest) GetExportUniqueResultsOk() (*bool, bool)`

GetExportUniqueResultsOk returns a tuple with the ExportUniqueResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUniqueResults

`func (o *UpdateClusterRequest) SetExportUniqueResults(v bool)`

SetExportUniqueResults sets ExportUniqueResults field to given value.


### GetNoLineBreaks

`func (o *UpdateClusterRequest) GetNoLineBreaks() bool`

GetNoLineBreaks returns the NoLineBreaks field if non-nil, zero value otherwise.

### GetNoLineBreaksOk

`func (o *UpdateClusterRequest) GetNoLineBreaksOk() (*bool, bool)`

GetNoLineBreaksOk returns a tuple with the NoLineBreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoLineBreaks

`func (o *UpdateClusterRequest) SetNoLineBreaks(v bool)`

SetNoLineBreaks sets NoLineBreaks field to given value.


### GetToComplete

`func (o *UpdateClusterRequest) GetToComplete() bool`

GetToComplete returns the ToComplete field if non-nil, zero value otherwise.

### GetToCompleteOk

`func (o *UpdateClusterRequest) GetToCompleteOk() (*bool, bool)`

GetToCompleteOk returns a tuple with the ToComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToComplete

`func (o *UpdateClusterRequest) SetToComplete(v bool)`

SetToComplete sets ToComplete field to given value.


### GetParams

`func (o *UpdateClusterRequest) GetParams() UpdateClusterRequestParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *UpdateClusterRequest) GetParamsOk() (*UpdateClusterRequestParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *UpdateClusterRequest) SetParams(v UpdateClusterRequestParams)`

SetParams sets Params field to given value.


### GetAccounts

`func (o *UpdateClusterRequest) GetAccounts() string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdateClusterRequest) GetAccountsOk() (*string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdateClusterRequest) SetAccounts(v string)`

SetAccounts sets Accounts field to given value.


### SetAccountsNil

`func (o *UpdateClusterRequest) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *UpdateClusterRequest) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil
### GetRunNotify

`func (o *UpdateClusterRequest) GetRunNotify() string`

GetRunNotify returns the RunNotify field if non-nil, zero value otherwise.

### GetRunNotifyOk

`func (o *UpdateClusterRequest) GetRunNotifyOk() (*string, bool)`

GetRunNotifyOk returns a tuple with the RunNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunNotify

`func (o *UpdateClusterRequest) SetRunNotify(v string)`

SetRunNotify sets RunNotify field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



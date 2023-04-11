# AccountCreated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 
**LastSynchronizationTime** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | **float32** |  | 
**StatusCodeInfo** | **string** |  | 
**StatusCodeDescription** | **string** |  | 
**TotalSynchronizations** | Pointer to **float32** | total synchronizations | [optional] 
**Type** | **string** |  | 
**UpdatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewAccountCreated

`func NewAccountCreated(id string, status float32, statusCodeInfo string, statusCodeDescription string, type_ string, ) *AccountCreated`

NewAccountCreated instantiates a new AccountCreated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCreatedWithDefaults

`func NewAccountCreatedWithDefaults() *AccountCreated`

NewAccountCreatedWithDefaults instantiates a new AccountCreated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountCreated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountCreated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountCreated) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *AccountCreated) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AccountCreated) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AccountCreated) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AccountCreated) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountCreated) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountCreated) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountCreated) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountCreated) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSynchronizationTime

`func (o *AccountCreated) GetLastSynchronizationTime() string`

GetLastSynchronizationTime returns the LastSynchronizationTime field if non-nil, zero value otherwise.

### GetLastSynchronizationTimeOk

`func (o *AccountCreated) GetLastSynchronizationTimeOk() (*string, bool)`

GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationTime

`func (o *AccountCreated) SetLastSynchronizationTime(v string)`

SetLastSynchronizationTime sets LastSynchronizationTime field to given value.

### HasLastSynchronizationTime

`func (o *AccountCreated) HasLastSynchronizationTime() bool`

HasLastSynchronizationTime returns a boolean if a field has been set.

### GetName

`func (o *AccountCreated) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountCreated) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountCreated) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountCreated) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AccountCreated) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountCreated) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountCreated) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetStatusCodeInfo

`func (o *AccountCreated) GetStatusCodeInfo() string`

GetStatusCodeInfo returns the StatusCodeInfo field if non-nil, zero value otherwise.

### GetStatusCodeInfoOk

`func (o *AccountCreated) GetStatusCodeInfoOk() (*string, bool)`

GetStatusCodeInfoOk returns a tuple with the StatusCodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeInfo

`func (o *AccountCreated) SetStatusCodeInfo(v string)`

SetStatusCodeInfo sets StatusCodeInfo field to given value.


### GetStatusCodeDescription

`func (o *AccountCreated) GetStatusCodeDescription() string`

GetStatusCodeDescription returns the StatusCodeDescription field if non-nil, zero value otherwise.

### GetStatusCodeDescriptionOk

`func (o *AccountCreated) GetStatusCodeDescriptionOk() (*string, bool)`

GetStatusCodeDescriptionOk returns a tuple with the StatusCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeDescription

`func (o *AccountCreated) SetStatusCodeDescription(v string)`

SetStatusCodeDescription sets StatusCodeDescription field to given value.


### GetTotalSynchronizations

`func (o *AccountCreated) GetTotalSynchronizations() float32`

GetTotalSynchronizations returns the TotalSynchronizations field if non-nil, zero value otherwise.

### GetTotalSynchronizationsOk

`func (o *AccountCreated) GetTotalSynchronizationsOk() (*float32, bool)`

GetTotalSynchronizationsOk returns a tuple with the TotalSynchronizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSynchronizations

`func (o *AccountCreated) SetTotalSynchronizations(v float32)`

SetTotalSynchronizations sets TotalSynchronizations field to given value.

### HasTotalSynchronizations

`func (o *AccountCreated) HasTotalSynchronizations() bool`

HasTotalSynchronizations returns a boolean if a field has been set.

### GetType

`func (o *AccountCreated) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountCreated) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountCreated) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *AccountCreated) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountCreated) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountCreated) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountCreated) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



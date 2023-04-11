# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 
**LastSynchronizationTime** | Pointer to **string** | date of last synchronization (ISO8601 format) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | **float32** |  | 
**StatusCodeInfo** | **string** |  | 
**StatusCodeDescription** | **string** |  | 
**TotalSynchronizations** | Pointer to **float32** | total synchronizations | [optional] 
**Type** | **string** |  | 
**UpdatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 

## Methods

### NewAccount

`func NewAccount(id string, status float32, statusCodeInfo string, statusCodeDescription string, type_ string, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Account) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Account) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Account) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Account) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Account) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Account) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Account) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Account) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSynchronizationTime

`func (o *Account) GetLastSynchronizationTime() string`

GetLastSynchronizationTime returns the LastSynchronizationTime field if non-nil, zero value otherwise.

### GetLastSynchronizationTimeOk

`func (o *Account) GetLastSynchronizationTimeOk() (*string, bool)`

GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationTime

`func (o *Account) SetLastSynchronizationTime(v string)`

SetLastSynchronizationTime sets LastSynchronizationTime field to given value.

### HasLastSynchronizationTime

`func (o *Account) HasLastSynchronizationTime() bool`

HasLastSynchronizationTime returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *Account) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetStatusCodeInfo

`func (o *Account) GetStatusCodeInfo() string`

GetStatusCodeInfo returns the StatusCodeInfo field if non-nil, zero value otherwise.

### GetStatusCodeInfoOk

`func (o *Account) GetStatusCodeInfoOk() (*string, bool)`

GetStatusCodeInfoOk returns a tuple with the StatusCodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeInfo

`func (o *Account) SetStatusCodeInfo(v string)`

SetStatusCodeInfo sets StatusCodeInfo field to given value.


### GetStatusCodeDescription

`func (o *Account) GetStatusCodeDescription() string`

GetStatusCodeDescription returns the StatusCodeDescription field if non-nil, zero value otherwise.

### GetStatusCodeDescriptionOk

`func (o *Account) GetStatusCodeDescriptionOk() (*string, bool)`

GetStatusCodeDescriptionOk returns a tuple with the StatusCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeDescription

`func (o *Account) SetStatusCodeDescription(v string)`

SetStatusCodeDescription sets StatusCodeDescription field to given value.


### GetTotalSynchronizations

`func (o *Account) GetTotalSynchronizations() float32`

GetTotalSynchronizations returns the TotalSynchronizations field if non-nil, zero value otherwise.

### GetTotalSynchronizationsOk

`func (o *Account) GetTotalSynchronizationsOk() (*float32, bool)`

GetTotalSynchronizationsOk returns a tuple with the TotalSynchronizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSynchronizations

`func (o *Account) SetTotalSynchronizations(v float32)`

SetTotalSynchronizations sets TotalSynchronizations field to given value.

### HasTotalSynchronizations

`func (o *Account) HasTotalSynchronizations() bool`

HasTotalSynchronizations returns a boolean if a field has been set.

### GetType

`func (o *Account) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Account) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Account) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *Account) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Account) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Account) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Account) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AccountRefreshed

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
**UpdatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 

## Methods

### NewAccountRefreshed

`func NewAccountRefreshed(id string, status float32, statusCodeInfo string, statusCodeDescription string, type_ string, ) *AccountRefreshed`

NewAccountRefreshed instantiates a new AccountRefreshed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountRefreshedWithDefaults

`func NewAccountRefreshedWithDefaults() *AccountRefreshed`

NewAccountRefreshedWithDefaults instantiates a new AccountRefreshed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountRefreshed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountRefreshed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountRefreshed) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *AccountRefreshed) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AccountRefreshed) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AccountRefreshed) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AccountRefreshed) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AccountRefreshed) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountRefreshed) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountRefreshed) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AccountRefreshed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastSynchronizationTime

`func (o *AccountRefreshed) GetLastSynchronizationTime() string`

GetLastSynchronizationTime returns the LastSynchronizationTime field if non-nil, zero value otherwise.

### GetLastSynchronizationTimeOk

`func (o *AccountRefreshed) GetLastSynchronizationTimeOk() (*string, bool)`

GetLastSynchronizationTimeOk returns a tuple with the LastSynchronizationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSynchronizationTime

`func (o *AccountRefreshed) SetLastSynchronizationTime(v string)`

SetLastSynchronizationTime sets LastSynchronizationTime field to given value.

### HasLastSynchronizationTime

`func (o *AccountRefreshed) HasLastSynchronizationTime() bool`

HasLastSynchronizationTime returns a boolean if a field has been set.

### GetName

`func (o *AccountRefreshed) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountRefreshed) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountRefreshed) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountRefreshed) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *AccountRefreshed) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountRefreshed) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountRefreshed) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetStatusCodeInfo

`func (o *AccountRefreshed) GetStatusCodeInfo() string`

GetStatusCodeInfo returns the StatusCodeInfo field if non-nil, zero value otherwise.

### GetStatusCodeInfoOk

`func (o *AccountRefreshed) GetStatusCodeInfoOk() (*string, bool)`

GetStatusCodeInfoOk returns a tuple with the StatusCodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeInfo

`func (o *AccountRefreshed) SetStatusCodeInfo(v string)`

SetStatusCodeInfo sets StatusCodeInfo field to given value.


### GetStatusCodeDescription

`func (o *AccountRefreshed) GetStatusCodeDescription() string`

GetStatusCodeDescription returns the StatusCodeDescription field if non-nil, zero value otherwise.

### GetStatusCodeDescriptionOk

`func (o *AccountRefreshed) GetStatusCodeDescriptionOk() (*string, bool)`

GetStatusCodeDescriptionOk returns a tuple with the StatusCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeDescription

`func (o *AccountRefreshed) SetStatusCodeDescription(v string)`

SetStatusCodeDescription sets StatusCodeDescription field to given value.


### GetTotalSynchronizations

`func (o *AccountRefreshed) GetTotalSynchronizations() float32`

GetTotalSynchronizations returns the TotalSynchronizations field if non-nil, zero value otherwise.

### GetTotalSynchronizationsOk

`func (o *AccountRefreshed) GetTotalSynchronizationsOk() (*float32, bool)`

GetTotalSynchronizationsOk returns a tuple with the TotalSynchronizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSynchronizations

`func (o *AccountRefreshed) SetTotalSynchronizations(v float32)`

SetTotalSynchronizations sets TotalSynchronizations field to given value.

### HasTotalSynchronizations

`func (o *AccountRefreshed) HasTotalSynchronizations() bool`

HasTotalSynchronizations returns a boolean if a field has been set.

### GetType

`func (o *AccountRefreshed) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountRefreshed) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountRefreshed) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *AccountRefreshed) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountRefreshed) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountRefreshed) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AccountRefreshed) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



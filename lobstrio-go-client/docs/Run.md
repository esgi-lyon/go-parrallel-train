# Run

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id | 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** | date of creation (ISO8601 format) | [optional] 
**DownloadLink** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** | date of next run (ISO8601 format) | [optional] 
**Origin** | **string** |  | 
**StartedAt** | Pointer to **string** | date of next run (ISO8601 format) | [optional] 
**Status** | **string** |  | 

## Methods

### NewRun

`func NewRun(id string, origin string, status string, ) *Run`

NewRun instantiates a new Run object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWithDefaults

`func NewRunWithDefaults() *Run`

NewRunWithDefaults instantiates a new Run object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Run) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Run) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Run) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Run) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Run) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Run) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Run) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Run) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Run) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Run) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Run) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDownloadLink

`func (o *Run) GetDownloadLink() string`

GetDownloadLink returns the DownloadLink field if non-nil, zero value otherwise.

### GetDownloadLinkOk

`func (o *Run) GetDownloadLinkOk() (*string, bool)`

GetDownloadLinkOk returns a tuple with the DownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLink

`func (o *Run) SetDownloadLink(v string)`

SetDownloadLink sets DownloadLink field to given value.

### HasDownloadLink

`func (o *Run) HasDownloadLink() bool`

HasDownloadLink returns a boolean if a field has been set.

### GetEndedAt

`func (o *Run) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *Run) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *Run) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *Run) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetOrigin

`func (o *Run) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *Run) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *Run) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetStartedAt

`func (o *Run) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Run) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Run) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Run) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Run) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Run) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Run) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



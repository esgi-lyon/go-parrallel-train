# RunDownloaded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id | 
**Object** | Pointer to **string** |  | [optional] 
**DownloadLink** | Pointer to **string** |  | [optional] 
**EndedAt** | Pointer to **string** |  | [optional] 
**Origin** | **string** |  | 
**StartedAt** | Pointer to **string** | date of next run (ISO8601 format) | [optional] 
**Status** | **string** |  | 

## Methods

### NewRunDownloaded

`func NewRunDownloaded(id string, origin string, status string, ) *RunDownloaded`

NewRunDownloaded instantiates a new RunDownloaded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunDownloadedWithDefaults

`func NewRunDownloadedWithDefaults() *RunDownloaded`

NewRunDownloadedWithDefaults instantiates a new RunDownloaded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunDownloaded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunDownloaded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunDownloaded) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *RunDownloaded) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RunDownloaded) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RunDownloaded) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *RunDownloaded) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetDownloadLink

`func (o *RunDownloaded) GetDownloadLink() string`

GetDownloadLink returns the DownloadLink field if non-nil, zero value otherwise.

### GetDownloadLinkOk

`func (o *RunDownloaded) GetDownloadLinkOk() (*string, bool)`

GetDownloadLinkOk returns a tuple with the DownloadLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLink

`func (o *RunDownloaded) SetDownloadLink(v string)`

SetDownloadLink sets DownloadLink field to given value.

### HasDownloadLink

`func (o *RunDownloaded) HasDownloadLink() bool`

HasDownloadLink returns a boolean if a field has been set.

### GetEndedAt

`func (o *RunDownloaded) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *RunDownloaded) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *RunDownloaded) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *RunDownloaded) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetOrigin

`func (o *RunDownloaded) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *RunDownloaded) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *RunDownloaded) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetStartedAt

`func (o *RunDownloaded) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *RunDownloaded) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *RunDownloaded) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *RunDownloaded) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *RunDownloaded) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RunDownloaded) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RunDownloaded) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



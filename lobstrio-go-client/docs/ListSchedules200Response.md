# ListSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** |  | [optional] 
**Page** | **float32** | page number | 
**Limit** | **float32** | number of items returned max | 
**Runs** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewListSchedules200Response

`func NewListSchedules200Response(page float32, limit float32, ) *ListSchedules200Response`

NewListSchedules200Response instantiates a new ListSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSchedules200ResponseWithDefaults

`func NewListSchedules200ResponseWithDefaults() *ListSchedules200Response`

NewListSchedules200ResponseWithDefaults instantiates a new ListSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListSchedules200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListSchedules200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListSchedules200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListSchedules200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *ListSchedules200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListSchedules200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListSchedules200Response) SetPage(v float32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *ListSchedules200Response) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListSchedules200Response) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListSchedules200Response) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetRuns

`func (o *ListSchedules200Response) GetRuns() []interface{}`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *ListSchedules200Response) GetRunsOk() (*[]interface{}, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *ListSchedules200Response) SetRuns(v []interface{})`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *ListSchedules200Response) HasRuns() bool`

HasRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



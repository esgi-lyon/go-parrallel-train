# ListClusters200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** |  | [optional] 
**Page** | **float32** | page number | 
**Limit** | **float32** | number of items returned max | 
**Data** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewListClusters200Response

`func NewListClusters200Response(page float32, limit float32, ) *ListClusters200Response`

NewListClusters200Response instantiates a new ListClusters200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusters200ResponseWithDefaults

`func NewListClusters200ResponseWithDefaults() *ListClusters200Response`

NewListClusters200ResponseWithDefaults instantiates a new ListClusters200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *ListClusters200Response) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListClusters200Response) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListClusters200Response) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListClusters200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetPage

`func (o *ListClusters200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ListClusters200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ListClusters200Response) SetPage(v float32)`

SetPage sets Page field to given value.


### GetLimit

`func (o *ListClusters200Response) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListClusters200Response) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListClusters200Response) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetData

`func (o *ListClusters200Response) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListClusters200Response) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListClusters200Response) SetData(v []interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ListClusters200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



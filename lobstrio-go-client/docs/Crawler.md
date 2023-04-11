# Crawler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | unique id | 
**Object** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** | date of creation (ISO8601 format) | 
**Description** | Pointer to **string** | description | [optional] 
**IsPublic** | **bool** | public status | 
**Labels** | Pointer to **[]interface{}** | labels | [optional] 
**MaxConcurrency** | **float32** | total number of concurrent threads allowed within one cluster | 
**Name** | **string** | public name | 
**Rank** | **float32** | position within the list of crawlers | 
**Slug** | **string** | unique slug | 
**UpdatedAt** | **string** | date of update (ISO8601 format) | 

## Methods

### NewCrawler

`func NewCrawler(id string, createdAt string, isPublic bool, maxConcurrency float32, name string, rank float32, slug string, updatedAt string, ) *Crawler`

NewCrawler instantiates a new Crawler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrawlerWithDefaults

`func NewCrawlerWithDefaults() *Crawler`

NewCrawlerWithDefaults instantiates a new Crawler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Crawler) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Crawler) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Crawler) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *Crawler) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Crawler) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Crawler) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Crawler) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Crawler) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Crawler) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Crawler) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *Crawler) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Crawler) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Crawler) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Crawler) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsPublic

`func (o *Crawler) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *Crawler) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *Crawler) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetLabels

`func (o *Crawler) GetLabels() []interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Crawler) GetLabelsOk() (*[]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Crawler) SetLabels(v []interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *Crawler) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMaxConcurrency

`func (o *Crawler) GetMaxConcurrency() float32`

GetMaxConcurrency returns the MaxConcurrency field if non-nil, zero value otherwise.

### GetMaxConcurrencyOk

`func (o *Crawler) GetMaxConcurrencyOk() (*float32, bool)`

GetMaxConcurrencyOk returns a tuple with the MaxConcurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrency

`func (o *Crawler) SetMaxConcurrency(v float32)`

SetMaxConcurrency sets MaxConcurrency field to given value.


### GetName

`func (o *Crawler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Crawler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Crawler) SetName(v string)`

SetName sets Name field to given value.


### GetRank

`func (o *Crawler) GetRank() float32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Crawler) GetRankOk() (*float32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Crawler) SetRank(v float32)`

SetRank sets Rank field to given value.


### GetSlug

`func (o *Crawler) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *Crawler) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *Crawler) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUpdatedAt

`func (o *Crawler) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Crawler) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Crawler) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



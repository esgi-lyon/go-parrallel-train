# AccountCrawlerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to [**AccountCrawlerParamsEmail**](AccountCrawlerParamsEmail.md) |  | [optional] 
**Password** | Pointer to [**AccountCrawlerParamsPassword**](AccountCrawlerParamsPassword.md) |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewAccountCrawlerParams

`func NewAccountCrawlerParams(type_ string, ) *AccountCrawlerParams`

NewAccountCrawlerParams instantiates a new AccountCrawlerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountCrawlerParamsWithDefaults

`func NewAccountCrawlerParamsWithDefaults() *AccountCrawlerParams`

NewAccountCrawlerParamsWithDefaults instantiates a new AccountCrawlerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AccountCrawlerParams) GetEmail() AccountCrawlerParamsEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AccountCrawlerParams) GetEmailOk() (*AccountCrawlerParamsEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AccountCrawlerParams) SetEmail(v AccountCrawlerParamsEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AccountCrawlerParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *AccountCrawlerParams) GetPassword() AccountCrawlerParamsPassword`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AccountCrawlerParams) GetPasswordOk() (*AccountCrawlerParamsPassword, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AccountCrawlerParams) SetPassword(v AccountCrawlerParamsPassword)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AccountCrawlerParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetType

`func (o *AccountCrawlerParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountCrawlerParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountCrawlerParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



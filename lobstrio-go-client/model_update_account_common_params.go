/*
lobstr API v1

Lobstr API is an easy-to-implement API to launch already-made data crawlers, handle accounts and schedules, and collect results.  ## Authentication  All requests must include the `Authorization` headers. You can obtain these from the settings menu.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateAccountCommonParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAccountCommonParams{}

// UpdateAccountCommonParams struct for UpdateAccountCommonParams
type UpdateAccountCommonParams struct {
	Name *string `json:"name,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewUpdateAccountCommonParams instantiates a new UpdateAccountCommonParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAccountCommonParams() *UpdateAccountCommonParams {
	this := UpdateAccountCommonParams{}
	return &this
}

// NewUpdateAccountCommonParamsWithDefaults instantiates a new UpdateAccountCommonParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAccountCommonParamsWithDefaults() *UpdateAccountCommonParams {
	this := UpdateAccountCommonParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAccountCommonParams) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCommonParams) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAccountCommonParams) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAccountCommonParams) SetName(v string) {
	o.Name = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateAccountCommonParams) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCommonParams) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateAccountCommonParams) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateAccountCommonParams) SetUsername(v string) {
	o.Username = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateAccountCommonParams) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAccountCommonParams) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateAccountCommonParams) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateAccountCommonParams) SetPassword(v string) {
	o.Password = &v
}

func (o UpdateAccountCommonParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAccountCommonParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableUpdateAccountCommonParams struct {
	value *UpdateAccountCommonParams
	isSet bool
}

func (v NullableUpdateAccountCommonParams) Get() *UpdateAccountCommonParams {
	return v.value
}

func (v *NullableUpdateAccountCommonParams) Set(val *UpdateAccountCommonParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAccountCommonParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAccountCommonParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAccountCommonParams(val *UpdateAccountCommonParams) *NullableUpdateAccountCommonParams {
	return &NullableUpdateAccountCommonParams{value: val, isSet: true}
}

func (v NullableUpdateAccountCommonParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAccountCommonParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


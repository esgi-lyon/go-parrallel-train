/*
Lobstr

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lobstrio

import (
	"encoding/json"
)

// checks if the ActivatetaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivatetaskRequest{}

// ActivatetaskRequest struct for ActivatetaskRequest
type ActivatetaskRequest struct {
	Id string `json:"id"`
	Object string `json:"object"`
	IsActive bool `json:"is_active"`
}

// NewActivatetaskRequest instantiates a new ActivatetaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivatetaskRequest(id string, object string, isActive bool) *ActivatetaskRequest {
	this := ActivatetaskRequest{}
	this.Id = id
	this.Object = object
	this.IsActive = isActive
	return &this
}

// NewActivatetaskRequestWithDefaults instantiates a new ActivatetaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivatetaskRequestWithDefaults() *ActivatetaskRequest {
	this := ActivatetaskRequest{}
	return &this
}

// GetId returns the Id field value
func (o *ActivatetaskRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActivatetaskRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActivatetaskRequest) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *ActivatetaskRequest) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ActivatetaskRequest) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ActivatetaskRequest) SetObject(v string) {
	o.Object = v
}

// GetIsActive returns the IsActive field value
func (o *ActivatetaskRequest) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *ActivatetaskRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *ActivatetaskRequest) SetIsActive(v bool) {
	o.IsActive = v
}

func (o ActivatetaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivatetaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["is_active"] = o.IsActive
	return toSerialize, nil
}

type NullableActivatetaskRequest struct {
	value *ActivatetaskRequest
	isSet bool
}

func (v NullableActivatetaskRequest) Get() *ActivatetaskRequest {
	return v.value
}

func (v *NullableActivatetaskRequest) Set(val *ActivatetaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActivatetaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActivatetaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivatetaskRequest(val *ActivatetaskRequest) *NullableActivatetaskRequest {
	return &NullableActivatetaskRequest{value: val, isSet: true}
}

func (v NullableActivatetaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivatetaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


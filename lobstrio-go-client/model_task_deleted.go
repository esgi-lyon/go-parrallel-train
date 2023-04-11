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

// checks if the TaskDeleted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskDeleted{}

// TaskDeleted struct for TaskDeleted
type TaskDeleted struct {
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
}

// NewTaskDeleted instantiates a new TaskDeleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskDeleted() *TaskDeleted {
	this := TaskDeleted{}
	return &this
}

// NewTaskDeletedWithDefaults instantiates a new TaskDeleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskDeletedWithDefaults() *TaskDeleted {
	this := TaskDeleted{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TaskDeleted) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDeleted) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TaskDeleted) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TaskDeleted) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *TaskDeleted) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDeleted) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *TaskDeleted) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *TaskDeleted) SetObject(v string) {
	o.Object = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *TaskDeleted) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskDeleted) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *TaskDeleted) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *TaskDeleted) SetDeleted(v bool) {
	o.Deleted = &v
}

func (o TaskDeleted) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskDeleted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	return toSerialize, nil
}

type NullableTaskDeleted struct {
	value *TaskDeleted
	isSet bool
}

func (v NullableTaskDeleted) Get() *TaskDeleted {
	return v.value
}

func (v *NullableTaskDeleted) Set(val *TaskDeleted) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskDeleted) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskDeleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskDeleted(val *TaskDeleted) *NullableTaskDeleted {
	return &NullableTaskDeleted{value: val, isSet: true}
}

func (v NullableTaskDeleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskDeleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



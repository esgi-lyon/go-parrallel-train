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

// checks if the TaskActivated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskActivated{}

// TaskActivated struct for TaskActivated
type TaskActivated struct {
	Id string `json:"id"`
	Object *string `json:"object,omitempty"`
	// date of creation (ISO8601 format)
	CreatedAt string `json:"created_at"`
	IsActive bool `json:"is_active"`
	Params TaskCommonParams `json:"params"`
	Status *string `json:"status,omitempty"`
}

// NewTaskActivated instantiates a new TaskActivated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskActivated(id string, createdAt string, isActive bool, params TaskCommonParams) *TaskActivated {
	this := TaskActivated{}
	this.Id = id
	this.CreatedAt = createdAt
	this.IsActive = isActive
	this.Params = params
	return &this
}

// NewTaskActivatedWithDefaults instantiates a new TaskActivated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskActivatedWithDefaults() *TaskActivated {
	this := TaskActivated{}
	return &this
}

// GetId returns the Id field value
func (o *TaskActivated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskActivated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskActivated) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *TaskActivated) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskActivated) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *TaskActivated) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *TaskActivated) SetObject(v string) {
	o.Object = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TaskActivated) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TaskActivated) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TaskActivated) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetIsActive returns the IsActive field value
func (o *TaskActivated) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *TaskActivated) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *TaskActivated) SetIsActive(v bool) {
	o.IsActive = v
}

// GetParams returns the Params field value
func (o *TaskActivated) GetParams() TaskCommonParams {
	if o == nil {
		var ret TaskCommonParams
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *TaskActivated) GetParamsOk() (*TaskCommonParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *TaskActivated) SetParams(v TaskCommonParams) {
	o.Params = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TaskActivated) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskActivated) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TaskActivated) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TaskActivated) SetStatus(v string) {
	o.Status = &v
}

func (o TaskActivated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskActivated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["is_active"] = o.IsActive
	toSerialize["params"] = o.Params
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableTaskActivated struct {
	value *TaskActivated
	isSet bool
}

func (v NullableTaskActivated) Get() *TaskActivated {
	return v.value
}

func (v *NullableTaskActivated) Set(val *TaskActivated) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskActivated) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskActivated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskActivated(val *TaskActivated) *NullableTaskActivated {
	return &NullableTaskActivated{value: val, isSet: true}
}

func (v NullableTaskActivated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskActivated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


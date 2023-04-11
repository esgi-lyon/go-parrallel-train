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

// checks if the CreatetasknotaddedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatetasknotaddedRequest{}

// CreatetasknotaddedRequest struct for CreatetasknotaddedRequest
type CreatetasknotaddedRequest struct {
	Cluster string `json:"cluster"`
	// 
	Tasks []Task `json:"tasks"`
}

// NewCreatetasknotaddedRequest instantiates a new CreatetasknotaddedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatetasknotaddedRequest(cluster string, tasks []Task) *CreatetasknotaddedRequest {
	this := CreatetasknotaddedRequest{}
	this.Cluster = cluster
	this.Tasks = tasks
	return &this
}

// NewCreatetasknotaddedRequestWithDefaults instantiates a new CreatetasknotaddedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatetasknotaddedRequestWithDefaults() *CreatetasknotaddedRequest {
	this := CreatetasknotaddedRequest{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *CreatetasknotaddedRequest) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *CreatetasknotaddedRequest) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *CreatetasknotaddedRequest) SetCluster(v string) {
	o.Cluster = v
}

// GetTasks returns the Tasks field value
func (o *CreatetasknotaddedRequest) GetTasks() []Task {
	if o == nil {
		var ret []Task
		return ret
	}

	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value
// and a boolean to check if the value has been set.
func (o *CreatetasknotaddedRequest) GetTasksOk() ([]Task, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tasks, true
}

// SetTasks sets field value
func (o *CreatetasknotaddedRequest) SetTasks(v []Task) {
	o.Tasks = v
}

func (o CreatetasknotaddedRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatetasknotaddedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["tasks"] = o.Tasks
	return toSerialize, nil
}

type NullableCreatetasknotaddedRequest struct {
	value *CreatetasknotaddedRequest
	isSet bool
}

func (v NullableCreatetasknotaddedRequest) Get() *CreatetasknotaddedRequest {
	return v.value
}

func (v *NullableCreatetasknotaddedRequest) Set(val *CreatetasknotaddedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatetasknotaddedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatetasknotaddedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatetasknotaddedRequest(val *CreatetasknotaddedRequest) *NullableCreatetasknotaddedRequest {
	return &NullableCreatetasknotaddedRequest{value: val, isSet: true}
}

func (v NullableCreatetasknotaddedRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatetasknotaddedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



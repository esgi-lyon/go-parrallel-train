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

// checks if the ListTasksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListTasksRequest{}

// ListTasksRequest struct for ListTasksRequest
type ListTasksRequest struct {
	Cluster *string `json:"cluster,omitempty"`
}

// NewListTasksRequest instantiates a new ListTasksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTasksRequest() *ListTasksRequest {
	this := ListTasksRequest{}
	return &this
}

// NewListTasksRequestWithDefaults instantiates a new ListTasksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTasksRequestWithDefaults() *ListTasksRequest {
	this := ListTasksRequest{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ListTasksRequest) GetCluster() string {
	if o == nil || IsNil(o.Cluster) {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTasksRequest) GetClusterOk() (*string, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ListTasksRequest) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *ListTasksRequest) SetCluster(v string) {
	o.Cluster = &v
}

func (o ListTasksRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListTasksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	return toSerialize, nil
}

type NullableListTasksRequest struct {
	value *ListTasksRequest
	isSet bool
}

func (v NullableListTasksRequest) Get() *ListTasksRequest {
	return v.value
}

func (v *NullableListTasksRequest) Set(val *ListTasksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableListTasksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableListTasksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTasksRequest(val *ListTasksRequest) *NullableListTasksRequest {
	return &NullableListTasksRequest{value: val, isSet: true}
}

func (v NullableListTasksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTasksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


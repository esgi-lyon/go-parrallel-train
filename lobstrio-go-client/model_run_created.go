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

// checks if the RunCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunCreated{}

// RunCreated struct for RunCreated
type RunCreated struct {
	Id *string `json:"id,omitempty"`
	Object *string `json:"object,omitempty"`
	// date of creation (ISO8601 format)
	CreatedAt *string `json:"created_at,omitempty"`
	DownloadLink *string `json:"download_link,omitempty"`
	EndedAt *string `json:"ended_at,omitempty"`
	Origin *string `json:"origin,omitempty"`
	StartedAt *string `json:"started_at,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewRunCreated instantiates a new RunCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunCreated() *RunCreated {
	this := RunCreated{}
	return &this
}

// NewRunCreatedWithDefaults instantiates a new RunCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunCreatedWithDefaults() *RunCreated {
	this := RunCreated{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RunCreated) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RunCreated) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RunCreated) SetId(v string) {
	o.Id = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *RunCreated) GetObject() string {
	if o == nil || IsNil(o.Object) {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetObjectOk() (*string, bool) {
	if o == nil || IsNil(o.Object) {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *RunCreated) HasObject() bool {
	if o != nil && !IsNil(o.Object) {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *RunCreated) SetObject(v string) {
	o.Object = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RunCreated) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RunCreated) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *RunCreated) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDownloadLink returns the DownloadLink field value if set, zero value otherwise.
func (o *RunCreated) GetDownloadLink() string {
	if o == nil || IsNil(o.DownloadLink) {
		var ret string
		return ret
	}
	return *o.DownloadLink
}

// GetDownloadLinkOk returns a tuple with the DownloadLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetDownloadLinkOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadLink) {
		return nil, false
	}
	return o.DownloadLink, true
}

// HasDownloadLink returns a boolean if a field has been set.
func (o *RunCreated) HasDownloadLink() bool {
	if o != nil && !IsNil(o.DownloadLink) {
		return true
	}

	return false
}

// SetDownloadLink gets a reference to the given string and assigns it to the DownloadLink field.
func (o *RunCreated) SetDownloadLink(v string) {
	o.DownloadLink = &v
}

// GetEndedAt returns the EndedAt field value if set, zero value otherwise.
func (o *RunCreated) GetEndedAt() string {
	if o == nil || IsNil(o.EndedAt) {
		var ret string
		return ret
	}
	return *o.EndedAt
}

// GetEndedAtOk returns a tuple with the EndedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetEndedAtOk() (*string, bool) {
	if o == nil || IsNil(o.EndedAt) {
		return nil, false
	}
	return o.EndedAt, true
}

// HasEndedAt returns a boolean if a field has been set.
func (o *RunCreated) HasEndedAt() bool {
	if o != nil && !IsNil(o.EndedAt) {
		return true
	}

	return false
}

// SetEndedAt gets a reference to the given string and assigns it to the EndedAt field.
func (o *RunCreated) SetEndedAt(v string) {
	o.EndedAt = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *RunCreated) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *RunCreated) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *RunCreated) SetOrigin(v string) {
	o.Origin = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *RunCreated) GetStartedAt() string {
	if o == nil || IsNil(o.StartedAt) {
		var ret string
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetStartedAtOk() (*string, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *RunCreated) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given string and assigns it to the StartedAt field.
func (o *RunCreated) SetStartedAt(v string) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RunCreated) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunCreated) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RunCreated) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RunCreated) SetStatus(v string) {
	o.Status = &v
}

func (o RunCreated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Object) {
		toSerialize["object"] = o.Object
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.DownloadLink) {
		toSerialize["download_link"] = o.DownloadLink
	}
	if !IsNil(o.EndedAt) {
		toSerialize["ended_at"] = o.EndedAt
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.StartedAt) {
		toSerialize["started_at"] = o.StartedAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableRunCreated struct {
	value *RunCreated
	isSet bool
}

func (v NullableRunCreated) Get() *RunCreated {
	return v.value
}

func (v *NullableRunCreated) Set(val *RunCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableRunCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableRunCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunCreated(val *RunCreated) *NullableRunCreated {
	return &NullableRunCreated{value: val, isSet: true}
}

func (v NullableRunCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



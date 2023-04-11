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

// checks if the TaskCrawlerParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskCrawlerParams{}

// TaskCrawlerParams struct for TaskCrawlerParams
type TaskCrawlerParams struct {
	MaxPages *TaskCrawlerParamsMaxPages `json:"max_pages,omitempty"`
	ResultsPerPage *TaskCrawlerParamsResultsPerPage `json:"results_per_page,omitempty"`
	Url *TaskCrawlerParamsUrl `json:"url,omitempty"`
}

// NewTaskCrawlerParams instantiates a new TaskCrawlerParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskCrawlerParams() *TaskCrawlerParams {
	this := TaskCrawlerParams{}
	return &this
}

// NewTaskCrawlerParamsWithDefaults instantiates a new TaskCrawlerParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskCrawlerParamsWithDefaults() *TaskCrawlerParams {
	this := TaskCrawlerParams{}
	return &this
}

// GetMaxPages returns the MaxPages field value if set, zero value otherwise.
func (o *TaskCrawlerParams) GetMaxPages() TaskCrawlerParamsMaxPages {
	if o == nil || IsNil(o.MaxPages) {
		var ret TaskCrawlerParamsMaxPages
		return ret
	}
	return *o.MaxPages
}

// GetMaxPagesOk returns a tuple with the MaxPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCrawlerParams) GetMaxPagesOk() (*TaskCrawlerParamsMaxPages, bool) {
	if o == nil || IsNil(o.MaxPages) {
		return nil, false
	}
	return o.MaxPages, true
}

// HasMaxPages returns a boolean if a field has been set.
func (o *TaskCrawlerParams) HasMaxPages() bool {
	if o != nil && !IsNil(o.MaxPages) {
		return true
	}

	return false
}

// SetMaxPages gets a reference to the given TaskCrawlerParamsMaxPages and assigns it to the MaxPages field.
func (o *TaskCrawlerParams) SetMaxPages(v TaskCrawlerParamsMaxPages) {
	o.MaxPages = &v
}

// GetResultsPerPage returns the ResultsPerPage field value if set, zero value otherwise.
func (o *TaskCrawlerParams) GetResultsPerPage() TaskCrawlerParamsResultsPerPage {
	if o == nil || IsNil(o.ResultsPerPage) {
		var ret TaskCrawlerParamsResultsPerPage
		return ret
	}
	return *o.ResultsPerPage
}

// GetResultsPerPageOk returns a tuple with the ResultsPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCrawlerParams) GetResultsPerPageOk() (*TaskCrawlerParamsResultsPerPage, bool) {
	if o == nil || IsNil(o.ResultsPerPage) {
		return nil, false
	}
	return o.ResultsPerPage, true
}

// HasResultsPerPage returns a boolean if a field has been set.
func (o *TaskCrawlerParams) HasResultsPerPage() bool {
	if o != nil && !IsNil(o.ResultsPerPage) {
		return true
	}

	return false
}

// SetResultsPerPage gets a reference to the given TaskCrawlerParamsResultsPerPage and assigns it to the ResultsPerPage field.
func (o *TaskCrawlerParams) SetResultsPerPage(v TaskCrawlerParamsResultsPerPage) {
	o.ResultsPerPage = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TaskCrawlerParams) GetUrl() TaskCrawlerParamsUrl {
	if o == nil || IsNil(o.Url) {
		var ret TaskCrawlerParamsUrl
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskCrawlerParams) GetUrlOk() (*TaskCrawlerParamsUrl, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TaskCrawlerParams) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given TaskCrawlerParamsUrl and assigns it to the Url field.
func (o *TaskCrawlerParams) SetUrl(v TaskCrawlerParamsUrl) {
	o.Url = &v
}

func (o TaskCrawlerParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskCrawlerParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxPages) {
		toSerialize["max_pages"] = o.MaxPages
	}
	if !IsNil(o.ResultsPerPage) {
		toSerialize["results_per_page"] = o.ResultsPerPage
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableTaskCrawlerParams struct {
	value *TaskCrawlerParams
	isSet bool
}

func (v NullableTaskCrawlerParams) Get() *TaskCrawlerParams {
	return v.value
}

func (v *NullableTaskCrawlerParams) Set(val *TaskCrawlerParams) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskCrawlerParams) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskCrawlerParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskCrawlerParams(val *TaskCrawlerParams) *NullableTaskCrawlerParams {
	return &NullableTaskCrawlerParams{value: val, isSet: true}
}

func (v NullableTaskCrawlerParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskCrawlerParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: v1alpha3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServingEnvironmentList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServingEnvironmentList{}

// ServingEnvironmentList List of ServingEnvironments.
type ServingEnvironmentList struct {
	// Token to use to retrieve next page of results.
	NextPageToken string `json:"nextPageToken"`
	// Maximum number of resources to return in the result.
	PageSize int32 `json:"pageSize"`
	// Number of items in result list.
	Size int32 `json:"size"`
	//
	Items []ServingEnvironment `json:"items,omitempty"`
}

// NewServingEnvironmentList instantiates a new ServingEnvironmentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServingEnvironmentList(nextPageToken string, pageSize int32, size int32) *ServingEnvironmentList {
	this := ServingEnvironmentList{}
	this.NextPageToken = nextPageToken
	this.PageSize = pageSize
	this.Size = size
	return &this
}

// NewServingEnvironmentListWithDefaults instantiates a new ServingEnvironmentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServingEnvironmentListWithDefaults() *ServingEnvironmentList {
	this := ServingEnvironmentList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value
func (o *ServingEnvironmentList) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentList) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *ServingEnvironmentList) SetNextPageToken(v string) {
	o.NextPageToken = v
}

// GetPageSize returns the PageSize field value
func (o *ServingEnvironmentList) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentList) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ServingEnvironmentList) SetPageSize(v int32) {
	o.PageSize = v
}

// GetSize returns the Size field value
func (o *ServingEnvironmentList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ServingEnvironmentList) SetSize(v int32) {
	o.Size = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ServingEnvironmentList) GetItems() []ServingEnvironment {
	if o == nil || IsNil(o.Items) {
		var ret []ServingEnvironment
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentList) GetItemsOk() ([]ServingEnvironment, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ServingEnvironmentList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ServingEnvironment and assigns it to the Items field.
func (o *ServingEnvironmentList) SetItems(v []ServingEnvironment) {
	o.Items = v
}

func (o ServingEnvironmentList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServingEnvironmentList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nextPageToken"] = o.NextPageToken
	toSerialize["pageSize"] = o.PageSize
	toSerialize["size"] = o.Size
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableServingEnvironmentList struct {
	value *ServingEnvironmentList
	isSet bool
}

func (v NullableServingEnvironmentList) Get() *ServingEnvironmentList {
	return v.value
}

func (v *NullableServingEnvironmentList) Set(val *ServingEnvironmentList) {
	v.value = val
	v.isSet = true
}

func (v NullableServingEnvironmentList) IsSet() bool {
	return v.isSet
}

func (v *NullableServingEnvironmentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServingEnvironmentList(val *ServingEnvironmentList) *NullableServingEnvironmentList {
	return &NullableServingEnvironmentList{value: val, isSet: true}
}

func (v NullableServingEnvironmentList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServingEnvironmentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

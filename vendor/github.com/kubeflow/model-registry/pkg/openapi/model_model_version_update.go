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

// checks if the ModelVersionUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelVersionUpdate{}

// ModelVersionUpdate Represents a ModelVersion belonging to a RegisteredModel.
type ModelVersionUpdate struct {
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalId *string            `json:"externalId,omitempty"`
	State      *ModelVersionState `json:"state,omitempty"`
	// Name of the author.
	Author *string `json:"author,omitempty"`
}

// NewModelVersionUpdate instantiates a new ModelVersionUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelVersionUpdate() *ModelVersionUpdate {
	this := ModelVersionUpdate{}
	var state ModelVersionState = MODELVERSIONSTATE_LIVE
	this.State = &state
	return &this
}

// NewModelVersionUpdateWithDefaults instantiates a new ModelVersionUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelVersionUpdateWithDefaults() *ModelVersionUpdate {
	this := ModelVersionUpdate{}
	var state ModelVersionState = MODELVERSIONSTATE_LIVE
	this.State = &state
	return &this
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *ModelVersionUpdate) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVersionUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *ModelVersionUpdate) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *ModelVersionUpdate) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelVersionUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVersionUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelVersionUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelVersionUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ModelVersionUpdate) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVersionUpdate) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ModelVersionUpdate) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ModelVersionUpdate) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ModelVersionUpdate) GetState() ModelVersionState {
	if o == nil || IsNil(o.State) {
		var ret ModelVersionState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVersionUpdate) GetStateOk() (*ModelVersionState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ModelVersionUpdate) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given ModelVersionState and assigns it to the State field.
func (o *ModelVersionUpdate) SetState(v ModelVersionState) {
	o.State = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *ModelVersionUpdate) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelVersionUpdate) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *ModelVersionUpdate) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *ModelVersionUpdate) SetAuthor(v string) {
	o.Author = &v
}

func (o ModelVersionUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelVersionUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	return toSerialize, nil
}

type NullableModelVersionUpdate struct {
	value *ModelVersionUpdate
	isSet bool
}

func (v NullableModelVersionUpdate) Get() *ModelVersionUpdate {
	return v.value
}

func (v *NullableModelVersionUpdate) Set(val *ModelVersionUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableModelVersionUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableModelVersionUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelVersionUpdate(val *ModelVersionUpdate) *NullableModelVersionUpdate {
	return &NullableModelVersionUpdate{value: val, isSet: true}
}

func (v NullableModelVersionUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelVersionUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

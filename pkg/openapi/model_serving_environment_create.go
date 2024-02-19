/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: v1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServingEnvironmentCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServingEnvironmentCreate{}

// ServingEnvironmentCreate A Model Serving environment for serving `RegisteredModels`.
type ServingEnvironmentCreate struct {
	// User provided custom properties which are not defined by its type.
	CustomProperties *map[string]MetadataValue `json:"customProperties,omitempty"`
	// An optional description about the resource.
	Description *string `json:"description,omitempty"`
	// The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance.
	ExternalID *string `json:"externalID,omitempty"`
	// The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set.
	Name *string `json:"name,omitempty"`
}

// NewServingEnvironmentCreate instantiates a new ServingEnvironmentCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServingEnvironmentCreate() *ServingEnvironmentCreate {
	this := ServingEnvironmentCreate{}
	return &this
}

// NewServingEnvironmentCreateWithDefaults instantiates a new ServingEnvironmentCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServingEnvironmentCreateWithDefaults() *ServingEnvironmentCreate {
	this := ServingEnvironmentCreate{}
	return &this
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *ServingEnvironmentCreate) GetCustomProperties() map[string]MetadataValue {
	if o == nil || IsNil(o.CustomProperties) {
		var ret map[string]MetadataValue
		return ret
	}
	return *o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool) {
	if o == nil || IsNil(o.CustomProperties) {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *ServingEnvironmentCreate) HasCustomProperties() bool {
	if o != nil && !IsNil(o.CustomProperties) {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given map[string]MetadataValue and assigns it to the CustomProperties field.
func (o *ServingEnvironmentCreate) SetCustomProperties(v map[string]MetadataValue) {
	o.CustomProperties = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServingEnvironmentCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServingEnvironmentCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServingEnvironmentCreate) SetDescription(v string) {
	o.Description = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *ServingEnvironmentCreate) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentCreate) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *ServingEnvironmentCreate) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *ServingEnvironmentCreate) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServingEnvironmentCreate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServingEnvironmentCreate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServingEnvironmentCreate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServingEnvironmentCreate) SetName(v string) {
	o.Name = &v
}

func (o ServingEnvironmentCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServingEnvironmentCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomProperties) {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableServingEnvironmentCreate struct {
	value *ServingEnvironmentCreate
	isSet bool
}

func (v NullableServingEnvironmentCreate) Get() *ServingEnvironmentCreate {
	return v.value
}

func (v *NullableServingEnvironmentCreate) Set(val *ServingEnvironmentCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableServingEnvironmentCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableServingEnvironmentCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServingEnvironmentCreate(val *ServingEnvironmentCreate) *NullableServingEnvironmentCreate {
	return &NullableServingEnvironmentCreate{value: val, isSet: true}
}

func (v NullableServingEnvironmentCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServingEnvironmentCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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

// checks if the MetadataStringValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataStringValue{}

// MetadataStringValue A string property value.
type MetadataStringValue struct {
	StringValue *string `json:"string_value,omitempty"`
}

// NewMetadataStringValue instantiates a new MetadataStringValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataStringValue() *MetadataStringValue {
	this := MetadataStringValue{}
	return &this
}

// NewMetadataStringValueWithDefaults instantiates a new MetadataStringValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataStringValueWithDefaults() *MetadataStringValue {
	this := MetadataStringValue{}
	return &this
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *MetadataStringValue) GetStringValue() string {
	if o == nil || IsNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataStringValue) GetStringValueOk() (*string, bool) {
	if o == nil || IsNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *MetadataStringValue) HasStringValue() bool {
	if o != nil && !IsNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *MetadataStringValue) SetStringValue(v string) {
	o.StringValue = &v
}

func (o MetadataStringValue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataStringValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StringValue) {
		toSerialize["string_value"] = o.StringValue
	}
	return toSerialize, nil
}

type NullableMetadataStringValue struct {
	value *MetadataStringValue
	isSet bool
}

func (v NullableMetadataStringValue) Get() *MetadataStringValue {
	return v.value
}

func (v *NullableMetadataStringValue) Set(val *MetadataStringValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataStringValue) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataStringValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataStringValue(val *MetadataStringValue) *NullableMetadataStringValue {
	return &NullableMetadataStringValue{value: val, isSet: true}
}

func (v NullableMetadataStringValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataStringValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
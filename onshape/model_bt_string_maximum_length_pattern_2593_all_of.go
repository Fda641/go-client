/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.113
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package onshape

import (
	"encoding/json"
)

// BTStringMaximumLengthPattern2593AllOf struct for BTStringMaximumLengthPattern2593AllOf
type BTStringMaximumLengthPattern2593AllOf struct {
	BtType *string `json:"btType,omitempty"`
	MaximumLength *int32 `json:"maximumLength,omitempty"`
}

// NewBTStringMaximumLengthPattern2593AllOf instantiates a new BTStringMaximumLengthPattern2593AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringMaximumLengthPattern2593AllOf() *BTStringMaximumLengthPattern2593AllOf {
	this := BTStringMaximumLengthPattern2593AllOf{}
	return &this
}

// NewBTStringMaximumLengthPattern2593AllOfWithDefaults instantiates a new BTStringMaximumLengthPattern2593AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringMaximumLengthPattern2593AllOfWithDefaults() *BTStringMaximumLengthPattern2593AllOf {
	this := BTStringMaximumLengthPattern2593AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringMaximumLengthPattern2593AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMaximumLengthPattern2593AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringMaximumLengthPattern2593AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringMaximumLengthPattern2593AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetMaximumLength returns the MaximumLength field value if set, zero value otherwise.
func (o *BTStringMaximumLengthPattern2593AllOf) GetMaximumLength() int32 {
	if o == nil || o.MaximumLength == nil {
		var ret int32
		return ret
	}
	return *o.MaximumLength
}

// GetMaximumLengthOk returns a tuple with the MaximumLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringMaximumLengthPattern2593AllOf) GetMaximumLengthOk() (*int32, bool) {
	if o == nil || o.MaximumLength == nil {
		return nil, false
	}
	return o.MaximumLength, true
}

// HasMaximumLength returns a boolean if a field has been set.
func (o *BTStringMaximumLengthPattern2593AllOf) HasMaximumLength() bool {
	if o != nil && o.MaximumLength != nil {
		return true
	}

	return false
}

// SetMaximumLength gets a reference to the given int32 and assigns it to the MaximumLength field.
func (o *BTStringMaximumLengthPattern2593AllOf) SetMaximumLength(v int32) {
	o.MaximumLength = &v
}

func (o BTStringMaximumLengthPattern2593AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.MaximumLength != nil {
		toSerialize["maximumLength"] = o.MaximumLength
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringMaximumLengthPattern2593AllOf struct {
	value *BTStringMaximumLengthPattern2593AllOf
	isSet bool
}

func (v NullableBTStringMaximumLengthPattern2593AllOf) Get() *BTStringMaximumLengthPattern2593AllOf {
	return v.value
}

func (v *NullableBTStringMaximumLengthPattern2593AllOf) Set(val *BTStringMaximumLengthPattern2593AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringMaximumLengthPattern2593AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringMaximumLengthPattern2593AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringMaximumLengthPattern2593AllOf(val *BTStringMaximumLengthPattern2593AllOf) *NullableBTStringMaximumLengthPattern2593AllOf {
	return &NullableBTStringMaximumLengthPattern2593AllOf{value: val, isSet: true}
}

func (v NullableBTStringMaximumLengthPattern2593AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringMaximumLengthPattern2593AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
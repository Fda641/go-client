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

// BTMParameterString149 struct for BTMParameterString149
type BTMParameterString149 struct {
	BTMParameter1
	BtType *string `json:"btType,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewBTMParameterString149 instantiates a new BTMParameterString149 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterString149() *BTMParameterString149 {
	this := BTMParameterString149{}
	return &this
}

// NewBTMParameterString149WithDefaults instantiates a new BTMParameterString149 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterString149WithDefaults() *BTMParameterString149 {
	this := BTMParameterString149{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterString149) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterString149) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterString149) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterString149) SetBtType(v string) {
	o.BtType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMParameterString149) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterString149) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMParameterString149) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTMParameterString149) SetValue(v string) {
	o.Value = &v
}

func (o BTMParameterString149) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMParameter1, errBTMParameter1 := json.Marshal(o.BTMParameter1)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	errBTMParameter1 = json.Unmarshal([]byte(serializedBTMParameter1), &toSerialize)
	if errBTMParameter1 != nil {
		return []byte{}, errBTMParameter1
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterString149 struct {
	value *BTMParameterString149
	isSet bool
}

func (v NullableBTMParameterString149) Get() *BTMParameterString149 {
	return v.value
}

func (v *NullableBTMParameterString149) Set(val *BTMParameterString149) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterString149) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterString149) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterString149(val *BTMParameterString149) *NullableBTMParameterString149 {
	return &NullableBTMParameterString149{value: val, isSet: true}
}

func (v NullableBTMParameterString149) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterString149) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
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

// BTParameterVisibilityOnEqual180 struct for BTParameterVisibilityOnEqual180
type BTParameterVisibilityOnEqual180 struct {
	BTParameterVisibilityCondition177
	BtType *string `json:"btType,omitempty"`
	ParameterId *string `json:"parameterId,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewBTParameterVisibilityOnEqual180 instantiates a new BTParameterVisibilityOnEqual180 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterVisibilityOnEqual180() *BTParameterVisibilityOnEqual180 {
	this := BTParameterVisibilityOnEqual180{}
	return &this
}

// NewBTParameterVisibilityOnEqual180WithDefaults instantiates a new BTParameterVisibilityOnEqual180 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterVisibilityOnEqual180WithDefaults() *BTParameterVisibilityOnEqual180 {
	this := BTParameterVisibilityOnEqual180{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnEqual180) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnEqual180) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnEqual180) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterVisibilityOnEqual180) SetBtType(v string) {
	o.BtType = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnEqual180) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnEqual180) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnEqual180) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTParameterVisibilityOnEqual180) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTParameterVisibilityOnEqual180) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterVisibilityOnEqual180) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTParameterVisibilityOnEqual180) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BTParameterVisibilityOnEqual180) SetValue(v string) {
	o.Value = &v
}

func (o BTParameterVisibilityOnEqual180) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterVisibilityCondition177, errBTParameterVisibilityCondition177 := json.Marshal(o.BTParameterVisibilityCondition177)
	if errBTParameterVisibilityCondition177 != nil {
		return []byte{}, errBTParameterVisibilityCondition177
	}
	errBTParameterVisibilityCondition177 = json.Unmarshal([]byte(serializedBTParameterVisibilityCondition177), &toSerialize)
	if errBTParameterVisibilityCondition177 != nil {
		return []byte{}, errBTParameterVisibilityCondition177
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterVisibilityOnEqual180 struct {
	value *BTParameterVisibilityOnEqual180
	isSet bool
}

func (v NullableBTParameterVisibilityOnEqual180) Get() *BTParameterVisibilityOnEqual180 {
	return v.value
}

func (v *NullableBTParameterVisibilityOnEqual180) Set(val *BTParameterVisibilityOnEqual180) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterVisibilityOnEqual180) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterVisibilityOnEqual180) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterVisibilityOnEqual180(val *BTParameterVisibilityOnEqual180) *NullableBTParameterVisibilityOnEqual180 {
	return &NullableBTParameterVisibilityOnEqual180{value: val, isSet: true}
}

func (v NullableBTParameterVisibilityOnEqual180) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterVisibilityOnEqual180) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
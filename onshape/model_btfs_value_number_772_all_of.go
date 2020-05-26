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

// BTFSValueNumber772AllOf struct for BTFSValueNumber772AllOf
type BTFSValueNumber772AllOf struct {
	BtType *string `json:"btType,omitempty"`
	Value *float64 `json:"value,omitempty"`
	ValueObject *float64 `json:"valueObject,omitempty"`
}

// NewBTFSValueNumber772AllOf instantiates a new BTFSValueNumber772AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSValueNumber772AllOf() *BTFSValueNumber772AllOf {
	this := BTFSValueNumber772AllOf{}
	return &this
}

// NewBTFSValueNumber772AllOfWithDefaults instantiates a new BTFSValueNumber772AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSValueNumber772AllOfWithDefaults() *BTFSValueNumber772AllOf {
	this := BTFSValueNumber772AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSValueNumber772AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueNumber772AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSValueNumber772AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSValueNumber772AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTFSValueNumber772AllOf) GetValue() float64 {
	if o == nil || o.Value == nil {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueNumber772AllOf) GetValueOk() (*float64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTFSValueNumber772AllOf) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *BTFSValueNumber772AllOf) SetValue(v float64) {
	o.Value = &v
}

// GetValueObject returns the ValueObject field value if set, zero value otherwise.
func (o *BTFSValueNumber772AllOf) GetValueObject() float64 {
	if o == nil || o.ValueObject == nil {
		var ret float64
		return ret
	}
	return *o.ValueObject
}

// GetValueObjectOk returns a tuple with the ValueObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSValueNumber772AllOf) GetValueObjectOk() (*float64, bool) {
	if o == nil || o.ValueObject == nil {
		return nil, false
	}
	return o.ValueObject, true
}

// HasValueObject returns a boolean if a field has been set.
func (o *BTFSValueNumber772AllOf) HasValueObject() bool {
	if o != nil && o.ValueObject != nil {
		return true
	}

	return false
}

// SetValueObject gets a reference to the given float64 and assigns it to the ValueObject field.
func (o *BTFSValueNumber772AllOf) SetValueObject(v float64) {
	o.ValueObject = &v
}

func (o BTFSValueNumber772AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueObject != nil {
		toSerialize["valueObject"] = o.ValueObject
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSValueNumber772AllOf struct {
	value *BTFSValueNumber772AllOf
	isSet bool
}

func (v NullableBTFSValueNumber772AllOf) Get() *BTFSValueNumber772AllOf {
	return v.value
}

func (v *NullableBTFSValueNumber772AllOf) Set(val *BTFSValueNumber772AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSValueNumber772AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSValueNumber772AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSValueNumber772AllOf(val *BTFSValueNumber772AllOf) *NullableBTFSValueNumber772AllOf {
	return &NullableBTFSValueNumber772AllOf{value: val, isSet: true}
}

func (v NullableBTFSValueNumber772AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSValueNumber772AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
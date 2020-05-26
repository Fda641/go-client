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

// BTPLiteralBoolean255 struct for BTPLiteralBoolean255
type BTPLiteralBoolean255 struct {
	BTPLiteral253
	BtType *string `json:"btType,omitempty"`
	Value *bool `json:"value,omitempty"`
}

// NewBTPLiteralBoolean255 instantiates a new BTPLiteralBoolean255 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLiteralBoolean255() *BTPLiteralBoolean255 {
	this := BTPLiteralBoolean255{}
	return &this
}

// NewBTPLiteralBoolean255WithDefaults instantiates a new BTPLiteralBoolean255 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLiteralBoolean255WithDefaults() *BTPLiteralBoolean255 {
	this := BTPLiteralBoolean255{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLiteralBoolean255) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralBoolean255) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLiteralBoolean255) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLiteralBoolean255) SetBtType(v string) {
	o.BtType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTPLiteralBoolean255) GetValue() bool {
	if o == nil || o.Value == nil {
		var ret bool
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLiteralBoolean255) GetValueOk() (*bool, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTPLiteralBoolean255) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given bool and assigns it to the Value field.
func (o *BTPLiteralBoolean255) SetValue(v bool) {
	o.Value = &v
}

func (o BTPLiteralBoolean255) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPLiteral253, errBTPLiteral253 := json.Marshal(o.BTPLiteral253)
	if errBTPLiteral253 != nil {
		return []byte{}, errBTPLiteral253
	}
	errBTPLiteral253 = json.Unmarshal([]byte(serializedBTPLiteral253), &toSerialize)
	if errBTPLiteral253 != nil {
		return []byte{}, errBTPLiteral253
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableBTPLiteralBoolean255 struct {
	value *BTPLiteralBoolean255
	isSet bool
}

func (v NullableBTPLiteralBoolean255) Get() *BTPLiteralBoolean255 {
	return v.value
}

func (v *NullableBTPLiteralBoolean255) Set(val *BTPLiteralBoolean255) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLiteralBoolean255) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLiteralBoolean255) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLiteralBoolean255(val *BTPLiteralBoolean255) *NullableBTPLiteralBoolean255 {
	return &NullableBTPLiteralBoolean255{value: val, isSet: true}
}

func (v NullableBTPLiteralBoolean255) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLiteralBoolean255) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
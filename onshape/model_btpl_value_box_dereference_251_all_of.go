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

// BTPLValueBoxDereference251AllOf struct for BTPLValueBoxDereference251AllOf
type BTPLValueBoxDereference251AllOf struct {
	Box *BTPExpression9 `json:"box,omitempty"`
	BtType *string `json:"btType,omitempty"`
	SpaceInside *BTPSpace10 `json:"spaceInside,omitempty"`
}

// NewBTPLValueBoxDereference251AllOf instantiates a new BTPLValueBoxDereference251AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLValueBoxDereference251AllOf() *BTPLValueBoxDereference251AllOf {
	this := BTPLValueBoxDereference251AllOf{}
	return &this
}

// NewBTPLValueBoxDereference251AllOfWithDefaults instantiates a new BTPLValueBoxDereference251AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLValueBoxDereference251AllOfWithDefaults() *BTPLValueBoxDereference251AllOf {
	this := BTPLValueBoxDereference251AllOf{}
	return &this
}

// GetBox returns the Box field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251AllOf) GetBox() BTPExpression9 {
	if o == nil || o.Box == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Box
}

// GetBoxOk returns a tuple with the Box field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251AllOf) GetBoxOk() (*BTPExpression9, bool) {
	if o == nil || o.Box == nil {
		return nil, false
	}
	return o.Box, true
}

// HasBox returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251AllOf) HasBox() bool {
	if o != nil && o.Box != nil {
		return true
	}

	return false
}

// SetBox gets a reference to the given BTPExpression9 and assigns it to the Box field.
func (o *BTPLValueBoxDereference251AllOf) SetBox(v BTPExpression9) {
	o.Box = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLValueBoxDereference251AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetSpaceInside returns the SpaceInside field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251AllOf) GetSpaceInside() BTPSpace10 {
	if o == nil || o.SpaceInside == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInside
}

// GetSpaceInsideOk returns a tuple with the SpaceInside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251AllOf) GetSpaceInsideOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInside == nil {
		return nil, false
	}
	return o.SpaceInside, true
}

// HasSpaceInside returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251AllOf) HasSpaceInside() bool {
	if o != nil && o.SpaceInside != nil {
		return true
	}

	return false
}

// SetSpaceInside gets a reference to the given BTPSpace10 and assigns it to the SpaceInside field.
func (o *BTPLValueBoxDereference251AllOf) SetSpaceInside(v BTPSpace10) {
	o.SpaceInside = &v
}

func (o BTPLValueBoxDereference251AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Box != nil {
		toSerialize["box"] = o.Box
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SpaceInside != nil {
		toSerialize["spaceInside"] = o.SpaceInside
	}
	return json.Marshal(toSerialize)
}

type NullableBTPLValueBoxDereference251AllOf struct {
	value *BTPLValueBoxDereference251AllOf
	isSet bool
}

func (v NullableBTPLValueBoxDereference251AllOf) Get() *BTPLValueBoxDereference251AllOf {
	return v.value
}

func (v *NullableBTPLValueBoxDereference251AllOf) Set(val *BTPLValueBoxDereference251AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLValueBoxDereference251AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLValueBoxDereference251AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLValueBoxDereference251AllOf(val *BTPLValueBoxDereference251AllOf) *NullableBTPLValueBoxDereference251AllOf {
	return &NullableBTPLValueBoxDereference251AllOf{value: val, isSet: true}
}

func (v NullableBTPLValueBoxDereference251AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLValueBoxDereference251AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
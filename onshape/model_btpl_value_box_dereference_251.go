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

// BTPLValueBoxDereference251 struct for BTPLValueBoxDereference251
type BTPLValueBoxDereference251 struct {
	BTPLValue249
	Box *BTPExpression9 `json:"box,omitempty"`
	BtType *string `json:"btType,omitempty"`
	SpaceInside *BTPSpace10 `json:"spaceInside,omitempty"`
}

// NewBTPLValueBoxDereference251 instantiates a new BTPLValueBoxDereference251 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPLValueBoxDereference251() *BTPLValueBoxDereference251 {
	this := BTPLValueBoxDereference251{}
	return &this
}

// NewBTPLValueBoxDereference251WithDefaults instantiates a new BTPLValueBoxDereference251 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPLValueBoxDereference251WithDefaults() *BTPLValueBoxDereference251 {
	this := BTPLValueBoxDereference251{}
	return &this
}

// GetBox returns the Box field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetBox() BTPExpression9 {
	if o == nil || o.Box == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Box
}

// GetBoxOk returns a tuple with the Box field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetBoxOk() (*BTPExpression9, bool) {
	if o == nil || o.Box == nil {
		return nil, false
	}
	return o.Box, true
}

// HasBox returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasBox() bool {
	if o != nil && o.Box != nil {
		return true
	}

	return false
}

// SetBox gets a reference to the given BTPExpression9 and assigns it to the Box field.
func (o *BTPLValueBoxDereference251) SetBox(v BTPExpression9) {
	o.Box = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPLValueBoxDereference251) SetBtType(v string) {
	o.BtType = &v
}

// GetSpaceInside returns the SpaceInside field value if set, zero value otherwise.
func (o *BTPLValueBoxDereference251) GetSpaceInside() BTPSpace10 {
	if o == nil || o.SpaceInside == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceInside
}

// GetSpaceInsideOk returns a tuple with the SpaceInside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPLValueBoxDereference251) GetSpaceInsideOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceInside == nil {
		return nil, false
	}
	return o.SpaceInside, true
}

// HasSpaceInside returns a boolean if a field has been set.
func (o *BTPLValueBoxDereference251) HasSpaceInside() bool {
	if o != nil && o.SpaceInside != nil {
		return true
	}

	return false
}

// SetSpaceInside gets a reference to the given BTPSpace10 and assigns it to the SpaceInside field.
func (o *BTPLValueBoxDereference251) SetSpaceInside(v BTPSpace10) {
	o.SpaceInside = &v
}

func (o BTPLValueBoxDereference251) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPLValue249, errBTPLValue249 := json.Marshal(o.BTPLValue249)
	if errBTPLValue249 != nil {
		return []byte{}, errBTPLValue249
	}
	errBTPLValue249 = json.Unmarshal([]byte(serializedBTPLValue249), &toSerialize)
	if errBTPLValue249 != nil {
		return []byte{}, errBTPLValue249
	}
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

type NullableBTPLValueBoxDereference251 struct {
	value *BTPLValueBoxDereference251
	isSet bool
}

func (v NullableBTPLValueBoxDereference251) Get() *BTPLValueBoxDereference251 {
	return v.value
}

func (v *NullableBTPLValueBoxDereference251) Set(val *BTPLValueBoxDereference251) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPLValueBoxDereference251) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPLValueBoxDereference251) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPLValueBoxDereference251(val *BTPLValueBoxDereference251) *NullableBTPLValueBoxDereference251 {
	return &NullableBTPLValueBoxDereference251{value: val, isSet: true}
}

func (v NullableBTPLValueBoxDereference251) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPLValueBoxDereference251) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

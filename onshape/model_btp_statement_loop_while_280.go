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

// BTPStatementLoopWhile280 struct for BTPStatementLoopWhile280
type BTPStatementLoopWhile280 struct {
	BTPStatementLoop277
	BtType *string `json:"btType,omitempty"`
	Condition *BTPExpression9 `json:"condition,omitempty"`
}

// NewBTPStatementLoopWhile280 instantiates a new BTPStatementLoopWhile280 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPStatementLoopWhile280() *BTPStatementLoopWhile280 {
	this := BTPStatementLoopWhile280{}
	return &this
}

// NewBTPStatementLoopWhile280WithDefaults instantiates a new BTPStatementLoopWhile280 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPStatementLoopWhile280WithDefaults() *BTPStatementLoopWhile280 {
	this := BTPStatementLoopWhile280{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPStatementLoopWhile280) SetBtType(v string) {
	o.BtType = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *BTPStatementLoopWhile280) GetCondition() BTPExpression9 {
	if o == nil || o.Condition == nil {
		var ret BTPExpression9
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPStatementLoopWhile280) GetConditionOk() (*BTPExpression9, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *BTPStatementLoopWhile280) HasCondition() bool {
	if o != nil && o.Condition != nil {
		return true
	}

	return false
}

// SetCondition gets a reference to the given BTPExpression9 and assigns it to the Condition field.
func (o *BTPStatementLoopWhile280) SetCondition(v BTPExpression9) {
	o.Condition = &v
}

func (o BTPStatementLoopWhile280) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPStatementLoop277, errBTPStatementLoop277 := json.Marshal(o.BTPStatementLoop277)
	if errBTPStatementLoop277 != nil {
		return []byte{}, errBTPStatementLoop277
	}
	errBTPStatementLoop277 = json.Unmarshal([]byte(serializedBTPStatementLoop277), &toSerialize)
	if errBTPStatementLoop277 != nil {
		return []byte{}, errBTPStatementLoop277
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	return json.Marshal(toSerialize)
}

type NullableBTPStatementLoopWhile280 struct {
	value *BTPStatementLoopWhile280
	isSet bool
}

func (v NullableBTPStatementLoopWhile280) Get() *BTPStatementLoopWhile280 {
	return v.value
}

func (v *NullableBTPStatementLoopWhile280) Set(val *BTPStatementLoopWhile280) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPStatementLoopWhile280) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPStatementLoopWhile280) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPStatementLoopWhile280(val *BTPStatementLoopWhile280) *NullableBTPStatementLoopWhile280 {
	return &NullableBTPStatementLoopWhile280{value: val, isSet: true}
}

func (v NullableBTPStatementLoopWhile280) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPStatementLoopWhile280) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
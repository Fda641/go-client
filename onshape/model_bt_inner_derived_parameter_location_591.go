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

// BTInnerDerivedParameterLocation591 struct for BTInnerDerivedParameterLocation591
type BTInnerDerivedParameterLocation591 struct {
	BTInnerParameterLocation1715
	BtType *string `json:"btType,omitempty"`
	OuterParameterId *string `json:"outerParameterId,omitempty"`
}

// NewBTInnerDerivedParameterLocation591 instantiates a new BTInnerDerivedParameterLocation591 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTInnerDerivedParameterLocation591() *BTInnerDerivedParameterLocation591 {
	this := BTInnerDerivedParameterLocation591{}
	return &this
}

// NewBTInnerDerivedParameterLocation591WithDefaults instantiates a new BTInnerDerivedParameterLocation591 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTInnerDerivedParameterLocation591WithDefaults() *BTInnerDerivedParameterLocation591 {
	this := BTInnerDerivedParameterLocation591{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTInnerDerivedParameterLocation591) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInnerDerivedParameterLocation591) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTInnerDerivedParameterLocation591) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTInnerDerivedParameterLocation591) SetBtType(v string) {
	o.BtType = &v
}

// GetOuterParameterId returns the OuterParameterId field value if set, zero value otherwise.
func (o *BTInnerDerivedParameterLocation591) GetOuterParameterId() string {
	if o == nil || o.OuterParameterId == nil {
		var ret string
		return ret
	}
	return *o.OuterParameterId
}

// GetOuterParameterIdOk returns a tuple with the OuterParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTInnerDerivedParameterLocation591) GetOuterParameterIdOk() (*string, bool) {
	if o == nil || o.OuterParameterId == nil {
		return nil, false
	}
	return o.OuterParameterId, true
}

// HasOuterParameterId returns a boolean if a field has been set.
func (o *BTInnerDerivedParameterLocation591) HasOuterParameterId() bool {
	if o != nil && o.OuterParameterId != nil {
		return true
	}

	return false
}

// SetOuterParameterId gets a reference to the given string and assigns it to the OuterParameterId field.
func (o *BTInnerDerivedParameterLocation591) SetOuterParameterId(v string) {
	o.OuterParameterId = &v
}

func (o BTInnerDerivedParameterLocation591) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTInnerParameterLocation1715, errBTInnerParameterLocation1715 := json.Marshal(o.BTInnerParameterLocation1715)
	if errBTInnerParameterLocation1715 != nil {
		return []byte{}, errBTInnerParameterLocation1715
	}
	errBTInnerParameterLocation1715 = json.Unmarshal([]byte(serializedBTInnerParameterLocation1715), &toSerialize)
	if errBTInnerParameterLocation1715 != nil {
		return []byte{}, errBTInnerParameterLocation1715
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.OuterParameterId != nil {
		toSerialize["outerParameterId"] = o.OuterParameterId
	}
	return json.Marshal(toSerialize)
}

type NullableBTInnerDerivedParameterLocation591 struct {
	value *BTInnerDerivedParameterLocation591
	isSet bool
}

func (v NullableBTInnerDerivedParameterLocation591) Get() *BTInnerDerivedParameterLocation591 {
	return v.value
}

func (v *NullableBTInnerDerivedParameterLocation591) Set(val *BTInnerDerivedParameterLocation591) {
	v.value = val
	v.isSet = true
}

func (v NullableBTInnerDerivedParameterLocation591) IsSet() bool {
	return v.isSet
}

func (v *NullableBTInnerDerivedParameterLocation591) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTInnerDerivedParameterLocation591(val *BTInnerDerivedParameterLocation591) *NullableBTInnerDerivedParameterLocation591 {
	return &NullableBTInnerDerivedParameterLocation591{value: val, isSet: true}
}

func (v NullableBTInnerDerivedParameterLocation591) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTInnerDerivedParameterLocation591) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
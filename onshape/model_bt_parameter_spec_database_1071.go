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

// BTParameterSpecDatabase1071 struct for BTParameterSpecDatabase1071
type BTParameterSpecDatabase1071 struct {
	BTParameterSpec6
	BtType *string `json:"btType,omitempty"`
}

// NewBTParameterSpecDatabase1071 instantiates a new BTParameterSpecDatabase1071 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTParameterSpecDatabase1071() *BTParameterSpecDatabase1071 {
	this := BTParameterSpecDatabase1071{}
	return &this
}

// NewBTParameterSpecDatabase1071WithDefaults instantiates a new BTParameterSpecDatabase1071 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTParameterSpecDatabase1071WithDefaults() *BTParameterSpecDatabase1071 {
	this := BTParameterSpecDatabase1071{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTParameterSpecDatabase1071) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTParameterSpecDatabase1071) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTParameterSpecDatabase1071) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTParameterSpecDatabase1071) SetBtType(v string) {
	o.BtType = &v
}

func (o BTParameterSpecDatabase1071) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTParameterSpec6, errBTParameterSpec6 := json.Marshal(o.BTParameterSpec6)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	errBTParameterSpec6 = json.Unmarshal([]byte(serializedBTParameterSpec6), &toSerialize)
	if errBTParameterSpec6 != nil {
		return []byte{}, errBTParameterSpec6
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTParameterSpecDatabase1071 struct {
	value *BTParameterSpecDatabase1071
	isSet bool
}

func (v NullableBTParameterSpecDatabase1071) Get() *BTParameterSpecDatabase1071 {
	return v.value
}

func (v *NullableBTParameterSpecDatabase1071) Set(val *BTParameterSpecDatabase1071) {
	v.value = val
	v.isSet = true
}

func (v NullableBTParameterSpecDatabase1071) IsSet() bool {
	return v.isSet
}

func (v *NullableBTParameterSpecDatabase1071) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTParameterSpecDatabase1071(val *BTParameterSpecDatabase1071) *NullableBTParameterSpecDatabase1071 {
	return &NullableBTParameterSpecDatabase1071{value: val, isSet: true}
}

func (v NullableBTParameterSpecDatabase1071) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTParameterSpecDatabase1071) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

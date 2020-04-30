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

// BTMMateGroup65 struct for BTMMateGroup65
type BTMMateGroup65 struct {
	BTMAssemblyFeature887
	BtType *string `json:"btType,omitempty"`
}

// NewBTMMateGroup65 instantiates a new BTMMateGroup65 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMMateGroup65() *BTMMateGroup65 {
	this := BTMMateGroup65{}
	return &this
}

// NewBTMMateGroup65WithDefaults instantiates a new BTMMateGroup65 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMMateGroup65WithDefaults() *BTMMateGroup65 {
	this := BTMMateGroup65{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMMateGroup65) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMateGroup65) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMMateGroup65) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMMateGroup65) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMMateGroup65) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTMAssemblyFeature887, errBTMAssemblyFeature887 := json.Marshal(o.BTMAssemblyFeature887)
	if errBTMAssemblyFeature887 != nil {
		return []byte{}, errBTMAssemblyFeature887
	}
	errBTMAssemblyFeature887 = json.Unmarshal([]byte(serializedBTMAssemblyFeature887), &toSerialize)
	if errBTMAssemblyFeature887 != nil {
		return []byte{}, errBTMAssemblyFeature887
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMMateGroup65 struct {
	value *BTMMateGroup65
	isSet bool
}

func (v NullableBTMMateGroup65) Get() *BTMMateGroup65 {
	return v.value
}

func (v *NullableBTMMateGroup65) Set(val *BTMMateGroup65) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMMateGroup65) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMMateGroup65) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMMateGroup65(val *BTMMateGroup65) *NullableBTMMateGroup65 {
	return &NullableBTMMateGroup65{value: val, isSet: true}
}

func (v NullableBTMMateGroup65) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMMateGroup65) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
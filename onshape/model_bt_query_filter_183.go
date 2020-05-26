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

// BTQueryFilter183 struct for BTQueryFilter183
type BTQueryFilter183 struct {
	BtType *string `json:"btType,omitempty"`
}

// NewBTQueryFilter183 instantiates a new BTQueryFilter183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTQueryFilter183() *BTQueryFilter183 {
	this := BTQueryFilter183{}
	return &this
}

// NewBTQueryFilter183WithDefaults instantiates a new BTQueryFilter183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTQueryFilter183WithDefaults() *BTQueryFilter183 {
	this := BTQueryFilter183{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTQueryFilter183) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTQueryFilter183) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTQueryFilter183) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTQueryFilter183) SetBtType(v string) {
	o.BtType = &v
}

func (o BTQueryFilter183) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTQueryFilter183 struct {
	value *BTQueryFilter183
	isSet bool
}

func (v NullableBTQueryFilter183) Get() *BTQueryFilter183 {
	return v.value
}

func (v *NullableBTQueryFilter183) Set(val *BTQueryFilter183) {
	v.value = val
	v.isSet = true
}

func (v NullableBTQueryFilter183) IsSet() bool {
	return v.isSet
}

func (v *NullableBTQueryFilter183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTQueryFilter183(val *BTQueryFilter183) *NullableBTQueryFilter183 {
	return &NullableBTQueryFilter183{value: val, isSet: true}
}

func (v NullableBTQueryFilter183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTQueryFilter183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

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

// BTMParameterAppearance627AllOf struct for BTMParameterAppearance627AllOf
type BTMParameterAppearance627AllOf struct {
	Appearance *BTGraphicsAppearance1152 `json:"appearance,omitempty"`
	BtType *string `json:"btType,omitempty"`
}

// NewBTMParameterAppearance627AllOf instantiates a new BTMParameterAppearance627AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterAppearance627AllOf() *BTMParameterAppearance627AllOf {
	this := BTMParameterAppearance627AllOf{}
	return &this
}

// NewBTMParameterAppearance627AllOfWithDefaults instantiates a new BTMParameterAppearance627AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterAppearance627AllOfWithDefaults() *BTMParameterAppearance627AllOf {
	this := BTMParameterAppearance627AllOf{}
	return &this
}

// GetAppearance returns the Appearance field value if set, zero value otherwise.
func (o *BTMParameterAppearance627AllOf) GetAppearance() BTGraphicsAppearance1152 {
	if o == nil || o.Appearance == nil {
		var ret BTGraphicsAppearance1152
		return ret
	}
	return *o.Appearance
}

// GetAppearanceOk returns a tuple with the Appearance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterAppearance627AllOf) GetAppearanceOk() (*BTGraphicsAppearance1152, bool) {
	if o == nil || o.Appearance == nil {
		return nil, false
	}
	return o.Appearance, true
}

// HasAppearance returns a boolean if a field has been set.
func (o *BTMParameterAppearance627AllOf) HasAppearance() bool {
	if o != nil && o.Appearance != nil {
		return true
	}

	return false
}

// SetAppearance gets a reference to the given BTGraphicsAppearance1152 and assigns it to the Appearance field.
func (o *BTMParameterAppearance627AllOf) SetAppearance(v BTGraphicsAppearance1152) {
	o.Appearance = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterAppearance627AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterAppearance627AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterAppearance627AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterAppearance627AllOf) SetBtType(v string) {
	o.BtType = &v
}

func (o BTMParameterAppearance627AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Appearance != nil {
		toSerialize["appearance"] = o.Appearance
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterAppearance627AllOf struct {
	value *BTMParameterAppearance627AllOf
	isSet bool
}

func (v NullableBTMParameterAppearance627AllOf) Get() *BTMParameterAppearance627AllOf {
	return v.value
}

func (v *NullableBTMParameterAppearance627AllOf) Set(val *BTMParameterAppearance627AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterAppearance627AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterAppearance627AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterAppearance627AllOf(val *BTMParameterAppearance627AllOf) *NullableBTMParameterAppearance627AllOf {
	return &NullableBTMParameterAppearance627AllOf{value: val, isSet: true}
}

func (v NullableBTMParameterAppearance627AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterAppearance627AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

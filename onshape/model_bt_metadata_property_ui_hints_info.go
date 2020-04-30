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

// BTMetadataPropertyUiHintsInfo struct for BTMetadataPropertyUiHintsInfo
type BTMetadataPropertyUiHintsInfo struct {
	Multiline *bool `json:"multiline,omitempty"`
}

// NewBTMetadataPropertyUiHintsInfo instantiates a new BTMetadataPropertyUiHintsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMetadataPropertyUiHintsInfo() *BTMetadataPropertyUiHintsInfo {
	this := BTMetadataPropertyUiHintsInfo{}
	return &this
}

// NewBTMetadataPropertyUiHintsInfoWithDefaults instantiates a new BTMetadataPropertyUiHintsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMetadataPropertyUiHintsInfoWithDefaults() *BTMetadataPropertyUiHintsInfo {
	this := BTMetadataPropertyUiHintsInfo{}
	return &this
}

// GetMultiline returns the Multiline field value if set, zero value otherwise.
func (o *BTMetadataPropertyUiHintsInfo) GetMultiline() bool {
	if o == nil || o.Multiline == nil {
		var ret bool
		return ret
	}
	return *o.Multiline
}

// GetMultilineOk returns a tuple with the Multiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMetadataPropertyUiHintsInfo) GetMultilineOk() (*bool, bool) {
	if o == nil || o.Multiline == nil {
		return nil, false
	}
	return o.Multiline, true
}

// HasMultiline returns a boolean if a field has been set.
func (o *BTMetadataPropertyUiHintsInfo) HasMultiline() bool {
	if o != nil && o.Multiline != nil {
		return true
	}

	return false
}

// SetMultiline gets a reference to the given bool and assigns it to the Multiline field.
func (o *BTMetadataPropertyUiHintsInfo) SetMultiline(v bool) {
	o.Multiline = &v
}

func (o BTMetadataPropertyUiHintsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Multiline != nil {
		toSerialize["multiline"] = o.Multiline
	}
	return json.Marshal(toSerialize)
}

type NullableBTMetadataPropertyUiHintsInfo struct {
	value *BTMetadataPropertyUiHintsInfo
	isSet bool
}

func (v NullableBTMetadataPropertyUiHintsInfo) Get() *BTMetadataPropertyUiHintsInfo {
	return v.value
}

func (v *NullableBTMetadataPropertyUiHintsInfo) Set(val *BTMetadataPropertyUiHintsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMetadataPropertyUiHintsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMetadataPropertyUiHintsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMetadataPropertyUiHintsInfo(val *BTMetadataPropertyUiHintsInfo) *NullableBTMetadataPropertyUiHintsInfo {
	return &NullableBTMetadataPropertyUiHintsInfo{value: val, isSet: true}
}

func (v NullableBTMetadataPropertyUiHintsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMetadataPropertyUiHintsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
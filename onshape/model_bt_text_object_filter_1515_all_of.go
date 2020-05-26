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

// BTTextObjectFilter1515AllOf struct for BTTextObjectFilter1515AllOf
type BTTextObjectFilter1515AllOf struct {
	BtType *string `json:"btType,omitempty"`
	IsText *bool `json:"isText,omitempty"`
}

// NewBTTextObjectFilter1515AllOf instantiates a new BTTextObjectFilter1515AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTTextObjectFilter1515AllOf() *BTTextObjectFilter1515AllOf {
	this := BTTextObjectFilter1515AllOf{}
	return &this
}

// NewBTTextObjectFilter1515AllOfWithDefaults instantiates a new BTTextObjectFilter1515AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTTextObjectFilter1515AllOfWithDefaults() *BTTextObjectFilter1515AllOf {
	this := BTTextObjectFilter1515AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTTextObjectFilter1515AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTextObjectFilter1515AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTTextObjectFilter1515AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTTextObjectFilter1515AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetIsText returns the IsText field value if set, zero value otherwise.
func (o *BTTextObjectFilter1515AllOf) GetIsText() bool {
	if o == nil || o.IsText == nil {
		var ret bool
		return ret
	}
	return *o.IsText
}

// GetIsTextOk returns a tuple with the IsText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTTextObjectFilter1515AllOf) GetIsTextOk() (*bool, bool) {
	if o == nil || o.IsText == nil {
		return nil, false
	}
	return o.IsText, true
}

// HasIsText returns a boolean if a field has been set.
func (o *BTTextObjectFilter1515AllOf) HasIsText() bool {
	if o != nil && o.IsText != nil {
		return true
	}

	return false
}

// SetIsText gets a reference to the given bool and assigns it to the IsText field.
func (o *BTTextObjectFilter1515AllOf) SetIsText(v bool) {
	o.IsText = &v
}

func (o BTTextObjectFilter1515AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsText != nil {
		toSerialize["isText"] = o.IsText
	}
	return json.Marshal(toSerialize)
}

type NullableBTTextObjectFilter1515AllOf struct {
	value *BTTextObjectFilter1515AllOf
	isSet bool
}

func (v NullableBTTextObjectFilter1515AllOf) Get() *BTTextObjectFilter1515AllOf {
	return v.value
}

func (v *NullableBTTextObjectFilter1515AllOf) Set(val *BTTextObjectFilter1515AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTTextObjectFilter1515AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTTextObjectFilter1515AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTTextObjectFilter1515AllOf(val *BTTextObjectFilter1515AllOf) *NullableBTTextObjectFilter1515AllOf {
	return &NullableBTTextObjectFilter1515AllOf{value: val, isSet: true}
}

func (v NullableBTTextObjectFilter1515AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTTextObjectFilter1515AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

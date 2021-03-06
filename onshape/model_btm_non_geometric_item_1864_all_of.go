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

// BTMNonGeometricItem1864AllOf struct for BTMNonGeometricItem1864AllOf
type BTMNonGeometricItem1864AllOf struct {
	BtType *string `json:"btType,omitempty"`
	ItemDefinitionId *string `json:"itemDefinitionId,omitempty"`
}

// NewBTMNonGeometricItem1864AllOf instantiates a new BTMNonGeometricItem1864AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMNonGeometricItem1864AllOf() *BTMNonGeometricItem1864AllOf {
	this := BTMNonGeometricItem1864AllOf{}
	return &this
}

// NewBTMNonGeometricItem1864AllOfWithDefaults instantiates a new BTMNonGeometricItem1864AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMNonGeometricItem1864AllOfWithDefaults() *BTMNonGeometricItem1864AllOf {
	this := BTMNonGeometricItem1864AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMNonGeometricItem1864AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetItemDefinitionId returns the ItemDefinitionId field value if set, zero value otherwise.
func (o *BTMNonGeometricItem1864AllOf) GetItemDefinitionId() string {
	if o == nil || o.ItemDefinitionId == nil {
		var ret string
		return ret
	}
	return *o.ItemDefinitionId
}

// GetItemDefinitionIdOk returns a tuple with the ItemDefinitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNonGeometricItem1864AllOf) GetItemDefinitionIdOk() (*string, bool) {
	if o == nil || o.ItemDefinitionId == nil {
		return nil, false
	}
	return o.ItemDefinitionId, true
}

// HasItemDefinitionId returns a boolean if a field has been set.
func (o *BTMNonGeometricItem1864AllOf) HasItemDefinitionId() bool {
	if o != nil && o.ItemDefinitionId != nil {
		return true
	}

	return false
}

// SetItemDefinitionId gets a reference to the given string and assigns it to the ItemDefinitionId field.
func (o *BTMNonGeometricItem1864AllOf) SetItemDefinitionId(v string) {
	o.ItemDefinitionId = &v
}

func (o BTMNonGeometricItem1864AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ItemDefinitionId != nil {
		toSerialize["itemDefinitionId"] = o.ItemDefinitionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMNonGeometricItem1864AllOf struct {
	value *BTMNonGeometricItem1864AllOf
	isSet bool
}

func (v NullableBTMNonGeometricItem1864AllOf) Get() *BTMNonGeometricItem1864AllOf {
	return v.value
}

func (v *NullableBTMNonGeometricItem1864AllOf) Set(val *BTMNonGeometricItem1864AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMNonGeometricItem1864AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMNonGeometricItem1864AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMNonGeometricItem1864AllOf(val *BTMNonGeometricItem1864AllOf) *NullableBTMNonGeometricItem1864AllOf {
	return &NullableBTMNonGeometricItem1864AllOf{value: val, isSet: true}
}

func (v NullableBTMNonGeometricItem1864AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMNonGeometricItem1864AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

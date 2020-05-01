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

// BTConstructionObjectFilter113 struct for BTConstructionObjectFilter113
type BTConstructionObjectFilter113 struct {
	BTQueryFilter183
	BtType *string `json:"btType,omitempty"`
	IsConstruction *bool `json:"isConstruction,omitempty"`
}

// NewBTConstructionObjectFilter113 instantiates a new BTConstructionObjectFilter113 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConstructionObjectFilter113() *BTConstructionObjectFilter113 {
	this := BTConstructionObjectFilter113{}
	return &this
}

// NewBTConstructionObjectFilter113WithDefaults instantiates a new BTConstructionObjectFilter113 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConstructionObjectFilter113WithDefaults() *BTConstructionObjectFilter113 {
	this := BTConstructionObjectFilter113{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConstructionObjectFilter113) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionObjectFilter113) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConstructionObjectFilter113) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConstructionObjectFilter113) SetBtType(v string) {
	o.BtType = &v
}

// GetIsConstruction returns the IsConstruction field value if set, zero value otherwise.
func (o *BTConstructionObjectFilter113) GetIsConstruction() bool {
	if o == nil || o.IsConstruction == nil {
		var ret bool
		return ret
	}
	return *o.IsConstruction
}

// GetIsConstructionOk returns a tuple with the IsConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConstructionObjectFilter113) GetIsConstructionOk() (*bool, bool) {
	if o == nil || o.IsConstruction == nil {
		return nil, false
	}
	return o.IsConstruction, true
}

// HasIsConstruction returns a boolean if a field has been set.
func (o *BTConstructionObjectFilter113) HasIsConstruction() bool {
	if o != nil && o.IsConstruction != nil {
		return true
	}

	return false
}

// SetIsConstruction gets a reference to the given bool and assigns it to the IsConstruction field.
func (o *BTConstructionObjectFilter113) SetIsConstruction(v bool) {
	o.IsConstruction = &v
}

func (o BTConstructionObjectFilter113) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsConstruction != nil {
		toSerialize["isConstruction"] = o.IsConstruction
	}
	return json.Marshal(toSerialize)
}

type NullableBTConstructionObjectFilter113 struct {
	value *BTConstructionObjectFilter113
	isSet bool
}

func (v NullableBTConstructionObjectFilter113) Get() *BTConstructionObjectFilter113 {
	return v.value
}

func (v *NullableBTConstructionObjectFilter113) Set(val *BTConstructionObjectFilter113) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConstructionObjectFilter113) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConstructionObjectFilter113) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConstructionObjectFilter113(val *BTConstructionObjectFilter113) *NullableBTConstructionObjectFilter113 {
	return &NullableBTConstructionObjectFilter113{value: val, isSet: true}
}

func (v NullableBTConstructionObjectFilter113) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConstructionObjectFilter113) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
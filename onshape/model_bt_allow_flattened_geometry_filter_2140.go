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

// BTAllowFlattenedGeometryFilter2140 struct for BTAllowFlattenedGeometryFilter2140
type BTAllowFlattenedGeometryFilter2140 struct {
	BTQueryFilter183
	AllowsFlattenedGeometry *bool `json:"allowsFlattenedGeometry,omitempty"`
	BtType *string `json:"btType,omitempty"`
}

// NewBTAllowFlattenedGeometryFilter2140 instantiates a new BTAllowFlattenedGeometryFilter2140 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAllowFlattenedGeometryFilter2140() *BTAllowFlattenedGeometryFilter2140 {
	this := BTAllowFlattenedGeometryFilter2140{}
	return &this
}

// NewBTAllowFlattenedGeometryFilter2140WithDefaults instantiates a new BTAllowFlattenedGeometryFilter2140 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAllowFlattenedGeometryFilter2140WithDefaults() *BTAllowFlattenedGeometryFilter2140 {
	this := BTAllowFlattenedGeometryFilter2140{}
	return &this
}

// GetAllowsFlattenedGeometry returns the AllowsFlattenedGeometry field value if set, zero value otherwise.
func (o *BTAllowFlattenedGeometryFilter2140) GetAllowsFlattenedGeometry() bool {
	if o == nil || o.AllowsFlattenedGeometry == nil {
		var ret bool
		return ret
	}
	return *o.AllowsFlattenedGeometry
}

// GetAllowsFlattenedGeometryOk returns a tuple with the AllowsFlattenedGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowFlattenedGeometryFilter2140) GetAllowsFlattenedGeometryOk() (*bool, bool) {
	if o == nil || o.AllowsFlattenedGeometry == nil {
		return nil, false
	}
	return o.AllowsFlattenedGeometry, true
}

// HasAllowsFlattenedGeometry returns a boolean if a field has been set.
func (o *BTAllowFlattenedGeometryFilter2140) HasAllowsFlattenedGeometry() bool {
	if o != nil && o.AllowsFlattenedGeometry != nil {
		return true
	}

	return false
}

// SetAllowsFlattenedGeometry gets a reference to the given bool and assigns it to the AllowsFlattenedGeometry field.
func (o *BTAllowFlattenedGeometryFilter2140) SetAllowsFlattenedGeometry(v bool) {
	o.AllowsFlattenedGeometry = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAllowFlattenedGeometryFilter2140) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAllowFlattenedGeometryFilter2140) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAllowFlattenedGeometryFilter2140) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAllowFlattenedGeometryFilter2140) SetBtType(v string) {
	o.BtType = &v
}

func (o BTAllowFlattenedGeometryFilter2140) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTQueryFilter183, errBTQueryFilter183 := json.Marshal(o.BTQueryFilter183)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	errBTQueryFilter183 = json.Unmarshal([]byte(serializedBTQueryFilter183), &toSerialize)
	if errBTQueryFilter183 != nil {
		return []byte{}, errBTQueryFilter183
	}
	if o.AllowsFlattenedGeometry != nil {
		toSerialize["allowsFlattenedGeometry"] = o.AllowsFlattenedGeometry
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	return json.Marshal(toSerialize)
}

type NullableBTAllowFlattenedGeometryFilter2140 struct {
	value *BTAllowFlattenedGeometryFilter2140
	isSet bool
}

func (v NullableBTAllowFlattenedGeometryFilter2140) Get() *BTAllowFlattenedGeometryFilter2140 {
	return v.value
}

func (v *NullableBTAllowFlattenedGeometryFilter2140) Set(val *BTAllowFlattenedGeometryFilter2140) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAllowFlattenedGeometryFilter2140) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAllowFlattenedGeometryFilter2140) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAllowFlattenedGeometryFilter2140(val *BTAllowFlattenedGeometryFilter2140) *NullableBTAllowFlattenedGeometryFilter2140 {
	return &NullableBTAllowFlattenedGeometryFilter2140{value: val, isSet: true}
}

func (v NullableBTAllowFlattenedGeometryFilter2140) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAllowFlattenedGeometryFilter2140) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

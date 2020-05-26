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

// BTGeometryFilter130 struct for BTGeometryFilter130
type BTGeometryFilter130 struct {
	BTQueryFilter183
	BtType *string `json:"btType,omitempty"`
	GeometryType *string `json:"geometryType,omitempty"`
}

// NewBTGeometryFilter130 instantiates a new BTGeometryFilter130 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGeometryFilter130() *BTGeometryFilter130 {
	this := BTGeometryFilter130{}
	return &this
}

// NewBTGeometryFilter130WithDefaults instantiates a new BTGeometryFilter130 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGeometryFilter130WithDefaults() *BTGeometryFilter130 {
	this := BTGeometryFilter130{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTGeometryFilter130) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryFilter130) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTGeometryFilter130) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTGeometryFilter130) SetBtType(v string) {
	o.BtType = &v
}

// GetGeometryType returns the GeometryType field value if set, zero value otherwise.
func (o *BTGeometryFilter130) GetGeometryType() string {
	if o == nil || o.GeometryType == nil {
		var ret string
		return ret
	}
	return *o.GeometryType
}

// GetGeometryTypeOk returns a tuple with the GeometryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGeometryFilter130) GetGeometryTypeOk() (*string, bool) {
	if o == nil || o.GeometryType == nil {
		return nil, false
	}
	return o.GeometryType, true
}

// HasGeometryType returns a boolean if a field has been set.
func (o *BTGeometryFilter130) HasGeometryType() bool {
	if o != nil && o.GeometryType != nil {
		return true
	}

	return false
}

// SetGeometryType gets a reference to the given string and assigns it to the GeometryType field.
func (o *BTGeometryFilter130) SetGeometryType(v string) {
	o.GeometryType = &v
}

func (o BTGeometryFilter130) MarshalJSON() ([]byte, error) {
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
	if o.GeometryType != nil {
		toSerialize["geometryType"] = o.GeometryType
	}
	return json.Marshal(toSerialize)
}

type NullableBTGeometryFilter130 struct {
	value *BTGeometryFilter130
	isSet bool
}

func (v NullableBTGeometryFilter130) Get() *BTGeometryFilter130 {
	return v.value
}

func (v *NullableBTGeometryFilter130) Set(val *BTGeometryFilter130) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGeometryFilter130) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGeometryFilter130) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGeometryFilter130(val *BTGeometryFilter130) *NullableBTGeometryFilter130 {
	return &NullableBTGeometryFilter130{value: val, isSet: true}
}

func (v NullableBTGeometryFilter130) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGeometryFilter130) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

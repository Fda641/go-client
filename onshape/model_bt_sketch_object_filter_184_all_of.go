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

// BTSketchObjectFilter184AllOf struct for BTSketchObjectFilter184AllOf
type BTSketchObjectFilter184AllOf struct {
	BtType *string `json:"btType,omitempty"`
	IsSketchObject *bool `json:"isSketchObject,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewBTSketchObjectFilter184AllOf instantiates a new BTSketchObjectFilter184AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSketchObjectFilter184AllOf() *BTSketchObjectFilter184AllOf {
	this := BTSketchObjectFilter184AllOf{}
	return &this
}

// NewBTSketchObjectFilter184AllOfWithDefaults instantiates a new BTSketchObjectFilter184AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSketchObjectFilter184AllOfWithDefaults() *BTSketchObjectFilter184AllOf {
	this := BTSketchObjectFilter184AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSketchObjectFilter184AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetIsSketchObject returns the IsSketchObject field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184AllOf) GetIsSketchObject() bool {
	if o == nil || o.IsSketchObject == nil {
		var ret bool
		return ret
	}
	return *o.IsSketchObject
}

// GetIsSketchObjectOk returns a tuple with the IsSketchObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184AllOf) GetIsSketchObjectOk() (*bool, bool) {
	if o == nil || o.IsSketchObject == nil {
		return nil, false
	}
	return o.IsSketchObject, true
}

// HasIsSketchObject returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184AllOf) HasIsSketchObject() bool {
	if o != nil && o.IsSketchObject != nil {
		return true
	}

	return false
}

// SetIsSketchObject gets a reference to the given bool and assigns it to the IsSketchObject field.
func (o *BTSketchObjectFilter184AllOf) SetIsSketchObject(v bool) {
	o.IsSketchObject = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *BTSketchObjectFilter184AllOf) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSketchObjectFilter184AllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *BTSketchObjectFilter184AllOf) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *BTSketchObjectFilter184AllOf) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o BTSketchObjectFilter184AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.IsSketchObject != nil {
		toSerialize["isSketchObject"] = o.IsSketchObject
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSketchObjectFilter184AllOf struct {
	value *BTSketchObjectFilter184AllOf
	isSet bool
}

func (v NullableBTSketchObjectFilter184AllOf) Get() *BTSketchObjectFilter184AllOf {
	return v.value
}

func (v *NullableBTSketchObjectFilter184AllOf) Set(val *BTSketchObjectFilter184AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSketchObjectFilter184AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSketchObjectFilter184AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSketchObjectFilter184AllOf(val *BTSketchObjectFilter184AllOf) *NullableBTSketchObjectFilter184AllOf {
	return &NullableBTSketchObjectFilter184AllOf{value: val, isSet: true}
}

func (v NullableBTSketchObjectFilter184AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSketchObjectFilter184AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
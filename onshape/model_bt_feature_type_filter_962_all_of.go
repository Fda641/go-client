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

// BTFeatureTypeFilter962AllOf struct for BTFeatureTypeFilter962AllOf
type BTFeatureTypeFilter962AllOf struct {
	BtType *string `json:"btType,omitempty"`
	FeatureType *string `json:"featureType,omitempty"`
}

// NewBTFeatureTypeFilter962AllOf instantiates a new BTFeatureTypeFilter962AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureTypeFilter962AllOf() *BTFeatureTypeFilter962AllOf {
	this := BTFeatureTypeFilter962AllOf{}
	return &this
}

// NewBTFeatureTypeFilter962AllOfWithDefaults instantiates a new BTFeatureTypeFilter962AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureTypeFilter962AllOfWithDefaults() *BTFeatureTypeFilter962AllOf {
	this := BTFeatureTypeFilter962AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureTypeFilter962AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureTypeFilter962AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureTypeFilter962AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureTypeFilter962AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTFeatureTypeFilter962AllOf) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureTypeFilter962AllOf) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTFeatureTypeFilter962AllOf) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTFeatureTypeFilter962AllOf) SetFeatureType(v string) {
	o.FeatureType = &v
}

func (o BTFeatureTypeFilter962AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureTypeFilter962AllOf struct {
	value *BTFeatureTypeFilter962AllOf
	isSet bool
}

func (v NullableBTFeatureTypeFilter962AllOf) Get() *BTFeatureTypeFilter962AllOf {
	return v.value
}

func (v *NullableBTFeatureTypeFilter962AllOf) Set(val *BTFeatureTypeFilter962AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureTypeFilter962AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureTypeFilter962AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureTypeFilter962AllOf(val *BTFeatureTypeFilter962AllOf) *NullableBTFeatureTypeFilter962AllOf {
	return &NullableBTFeatureTypeFilter962AllOf{value: val, isSet: true}
}

func (v NullableBTFeatureTypeFilter962AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureTypeFilter962AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
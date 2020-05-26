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

// BTAssemblyFeatureListResponse1174AllOf struct for BTAssemblyFeatureListResponse1174AllOf
type BTAssemblyFeatureListResponse1174AllOf struct {
	BtType *string `json:"btType,omitempty"`
	FeatureStates *map[string]BTFeatureState1688 `json:"featureStates,omitempty"`
	Features *[]BTMAssemblyFeature887 `json:"features,omitempty"`
	IsComplete *bool `json:"isComplete,omitempty"`
}

// NewBTAssemblyFeatureListResponse1174AllOf instantiates a new BTAssemblyFeatureListResponse1174AllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyFeatureListResponse1174AllOf() *BTAssemblyFeatureListResponse1174AllOf {
	this := BTAssemblyFeatureListResponse1174AllOf{}
	return &this
}

// NewBTAssemblyFeatureListResponse1174AllOfWithDefaults instantiates a new BTAssemblyFeatureListResponse1174AllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyFeatureListResponse1174AllOfWithDefaults() *BTAssemblyFeatureListResponse1174AllOf {
	this := BTAssemblyFeatureListResponse1174AllOf{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAssemblyFeatureListResponse1174AllOf) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureStates returns the FeatureStates field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetFeatureStates() map[string]BTFeatureState1688 {
	if o == nil || o.FeatureStates == nil {
		var ret map[string]BTFeatureState1688
		return ret
	}
	return *o.FeatureStates
}

// GetFeatureStatesOk returns a tuple with the FeatureStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool) {
	if o == nil || o.FeatureStates == nil {
		return nil, false
	}
	return o.FeatureStates, true
}

// HasFeatureStates returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) HasFeatureStates() bool {
	if o != nil && o.FeatureStates != nil {
		return true
	}

	return false
}

// SetFeatureStates gets a reference to the given map[string]BTFeatureState1688 and assigns it to the FeatureStates field.
func (o *BTAssemblyFeatureListResponse1174AllOf) SetFeatureStates(v map[string]BTFeatureState1688) {
	o.FeatureStates = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetFeatures() []BTMAssemblyFeature887 {
	if o == nil || o.Features == nil {
		var ret []BTMAssemblyFeature887
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetFeaturesOk() (*[]BTMAssemblyFeature887, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BTMAssemblyFeature887 and assigns it to the Features field.
func (o *BTAssemblyFeatureListResponse1174AllOf) SetFeatures(v []BTMAssemblyFeature887) {
	o.Features = &v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetIsComplete() bool {
	if o == nil || o.IsComplete == nil {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) GetIsCompleteOk() (*bool, bool) {
	if o == nil || o.IsComplete == nil {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174AllOf) HasIsComplete() bool {
	if o != nil && o.IsComplete != nil {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *BTAssemblyFeatureListResponse1174AllOf) SetIsComplete(v bool) {
	o.IsComplete = &v
}

func (o BTAssemblyFeatureListResponse1174AllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureStates != nil {
		toSerialize["featureStates"] = o.FeatureStates
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.IsComplete != nil {
		toSerialize["isComplete"] = o.IsComplete
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyFeatureListResponse1174AllOf struct {
	value *BTAssemblyFeatureListResponse1174AllOf
	isSet bool
}

func (v NullableBTAssemblyFeatureListResponse1174AllOf) Get() *BTAssemblyFeatureListResponse1174AllOf {
	return v.value
}

func (v *NullableBTAssemblyFeatureListResponse1174AllOf) Set(val *BTAssemblyFeatureListResponse1174AllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyFeatureListResponse1174AllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyFeatureListResponse1174AllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyFeatureListResponse1174AllOf(val *BTAssemblyFeatureListResponse1174AllOf) *NullableBTAssemblyFeatureListResponse1174AllOf {
	return &NullableBTAssemblyFeatureListResponse1174AllOf{value: val, isSet: true}
}

func (v NullableBTAssemblyFeatureListResponse1174AllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyFeatureListResponse1174AllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
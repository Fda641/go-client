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

// BTUpdateReferenceParams struct for BTUpdateReferenceParams
type BTUpdateReferenceParams struct {
	ReferenceUpdates *[]UpdateParams `json:"referenceUpdates,omitempty"`
}

// NewBTUpdateReferenceParams instantiates a new BTUpdateReferenceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUpdateReferenceParams() *BTUpdateReferenceParams {
	this := BTUpdateReferenceParams{}
	return &this
}

// NewBTUpdateReferenceParamsWithDefaults instantiates a new BTUpdateReferenceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUpdateReferenceParamsWithDefaults() *BTUpdateReferenceParams {
	this := BTUpdateReferenceParams{}
	return &this
}

// GetReferenceUpdates returns the ReferenceUpdates field value if set, zero value otherwise.
func (o *BTUpdateReferenceParams) GetReferenceUpdates() []UpdateParams {
	if o == nil || o.ReferenceUpdates == nil {
		var ret []UpdateParams
		return ret
	}
	return *o.ReferenceUpdates
}

// GetReferenceUpdatesOk returns a tuple with the ReferenceUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateReferenceParams) GetReferenceUpdatesOk() (*[]UpdateParams, bool) {
	if o == nil || o.ReferenceUpdates == nil {
		return nil, false
	}
	return o.ReferenceUpdates, true
}

// HasReferenceUpdates returns a boolean if a field has been set.
func (o *BTUpdateReferenceParams) HasReferenceUpdates() bool {
	if o != nil && o.ReferenceUpdates != nil {
		return true
	}

	return false
}

// SetReferenceUpdates gets a reference to the given []UpdateParams and assigns it to the ReferenceUpdates field.
func (o *BTUpdateReferenceParams) SetReferenceUpdates(v []UpdateParams) {
	o.ReferenceUpdates = &v
}

func (o BTUpdateReferenceParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ReferenceUpdates != nil {
		toSerialize["referenceUpdates"] = o.ReferenceUpdates
	}
	return json.Marshal(toSerialize)
}

type NullableBTUpdateReferenceParams struct {
	value *BTUpdateReferenceParams
	isSet bool
}

func (v NullableBTUpdateReferenceParams) Get() *BTUpdateReferenceParams {
	return v.value
}

func (v *NullableBTUpdateReferenceParams) Set(val *BTUpdateReferenceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUpdateReferenceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUpdateReferenceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUpdateReferenceParams(val *BTUpdateReferenceParams) *NullableBTUpdateReferenceParams {
	return &NullableBTUpdateReferenceParams{value: val, isSet: true}
}

func (v NullableBTUpdateReferenceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUpdateReferenceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

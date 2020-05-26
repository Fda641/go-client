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

// BTWebhookOptions struct for BTWebhookOptions
type BTWebhookOptions struct {
	CollapseEvents *bool `json:"collapseEvents,omitempty"`
}

// NewBTWebhookOptions instantiates a new BTWebhookOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWebhookOptions() *BTWebhookOptions {
	this := BTWebhookOptions{}
	return &this
}

// NewBTWebhookOptionsWithDefaults instantiates a new BTWebhookOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWebhookOptionsWithDefaults() *BTWebhookOptions {
	this := BTWebhookOptions{}
	return &this
}

// GetCollapseEvents returns the CollapseEvents field value if set, zero value otherwise.
func (o *BTWebhookOptions) GetCollapseEvents() bool {
	if o == nil || o.CollapseEvents == nil {
		var ret bool
		return ret
	}
	return *o.CollapseEvents
}

// GetCollapseEventsOk returns a tuple with the CollapseEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookOptions) GetCollapseEventsOk() (*bool, bool) {
	if o == nil || o.CollapseEvents == nil {
		return nil, false
	}
	return o.CollapseEvents, true
}

// HasCollapseEvents returns a boolean if a field has been set.
func (o *BTWebhookOptions) HasCollapseEvents() bool {
	if o != nil && o.CollapseEvents != nil {
		return true
	}

	return false
}

// SetCollapseEvents gets a reference to the given bool and assigns it to the CollapseEvents field.
func (o *BTWebhookOptions) SetCollapseEvents(v bool) {
	o.CollapseEvents = &v
}

func (o BTWebhookOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CollapseEvents != nil {
		toSerialize["collapseEvents"] = o.CollapseEvents
	}
	return json.Marshal(toSerialize)
}

type NullableBTWebhookOptions struct {
	value *BTWebhookOptions
	isSet bool
}

func (v NullableBTWebhookOptions) Get() *BTWebhookOptions {
	return v.value
}

func (v *NullableBTWebhookOptions) Set(val *BTWebhookOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWebhookOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWebhookOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWebhookOptions(val *BTWebhookOptions) *NullableBTWebhookOptions {
	return &NullableBTWebhookOptions{value: val, isSet: true}
}

func (v NullableBTWebhookOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWebhookOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

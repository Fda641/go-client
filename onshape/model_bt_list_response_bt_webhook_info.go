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

// BTListResponseBTWebhookInfo struct for BTListResponseBTWebhookInfo
type BTListResponseBTWebhookInfo struct {
	Href *string `json:"href,omitempty"`
	Items *[]BTWebhookInfo `json:"items,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}

// NewBTListResponseBTWebhookInfo instantiates a new BTListResponseBTWebhookInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTListResponseBTWebhookInfo() *BTListResponseBTWebhookInfo {
	this := BTListResponseBTWebhookInfo{}
	return &this
}

// NewBTListResponseBTWebhookInfoWithDefaults instantiates a new BTListResponseBTWebhookInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTListResponseBTWebhookInfoWithDefaults() *BTListResponseBTWebhookInfo {
	this := BTListResponseBTWebhookInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTListResponseBTWebhookInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTWebhookInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTListResponseBTWebhookInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTListResponseBTWebhookInfo) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTListResponseBTWebhookInfo) GetItems() []BTWebhookInfo {
	if o == nil || o.Items == nil {
		var ret []BTWebhookInfo
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTWebhookInfo) GetItemsOk() (*[]BTWebhookInfo, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTListResponseBTWebhookInfo) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTWebhookInfo and assigns it to the Items field.
func (o *BTListResponseBTWebhookInfo) SetItems(v []BTWebhookInfo) {
	o.Items = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BTListResponseBTWebhookInfo) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTWebhookInfo) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BTListResponseBTWebhookInfo) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *BTListResponseBTWebhookInfo) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *BTListResponseBTWebhookInfo) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTListResponseBTWebhookInfo) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *BTListResponseBTWebhookInfo) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *BTListResponseBTWebhookInfo) SetPrevious(v string) {
	o.Previous = &v
}

func (o BTListResponseBTWebhookInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	return json.Marshal(toSerialize)
}

type NullableBTListResponseBTWebhookInfo struct {
	value *BTListResponseBTWebhookInfo
	isSet bool
}

func (v NullableBTListResponseBTWebhookInfo) Get() *BTListResponseBTWebhookInfo {
	return v.value
}

func (v *NullableBTListResponseBTWebhookInfo) Set(val *BTListResponseBTWebhookInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTListResponseBTWebhookInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTListResponseBTWebhookInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTListResponseBTWebhookInfo(val *BTListResponseBTWebhookInfo) *NullableBTListResponseBTWebhookInfo {
	return &NullableBTListResponseBTWebhookInfo{value: val, isSet: true}
}

func (v NullableBTListResponseBTWebhookInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTListResponseBTWebhookInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

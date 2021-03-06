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

// BTGlobalTreeNodeListResponse struct for BTGlobalTreeNodeListResponse
type BTGlobalTreeNodeListResponse struct {
	Href *string `json:"href,omitempty"`
	Items *[]BTGlobalTreeNodeInfo `json:"items,omitempty"`
	Next *string `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
}

// NewBTGlobalTreeNodeListResponse instantiates a new BTGlobalTreeNodeListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTGlobalTreeNodeListResponse() *BTGlobalTreeNodeListResponse {
	this := BTGlobalTreeNodeListResponse{}
	return &this
}

// NewBTGlobalTreeNodeListResponseWithDefaults instantiates a new BTGlobalTreeNodeListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTGlobalTreeNodeListResponseWithDefaults() *BTGlobalTreeNodeListResponse {
	this := BTGlobalTreeNodeListResponse{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponse) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponse) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponse) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTGlobalTreeNodeListResponse) SetHref(v string) {
	o.Href = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponse) GetItems() []BTGlobalTreeNodeInfo {
	if o == nil || o.Items == nil {
		var ret []BTGlobalTreeNodeInfo
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponse) GetItemsOk() (*[]BTGlobalTreeNodeInfo, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BTGlobalTreeNodeInfo and assigns it to the Items field.
func (o *BTGlobalTreeNodeListResponse) SetItems(v []BTGlobalTreeNodeInfo) {
	o.Items = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponse) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponse) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponse) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *BTGlobalTreeNodeListResponse) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *BTGlobalTreeNodeListResponse) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTGlobalTreeNodeListResponse) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *BTGlobalTreeNodeListResponse) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *BTGlobalTreeNodeListResponse) SetPrevious(v string) {
	o.Previous = &v
}

func (o BTGlobalTreeNodeListResponse) MarshalJSON() ([]byte, error) {
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

type NullableBTGlobalTreeNodeListResponse struct {
	value *BTGlobalTreeNodeListResponse
	isSet bool
}

func (v NullableBTGlobalTreeNodeListResponse) Get() *BTGlobalTreeNodeListResponse {
	return v.value
}

func (v *NullableBTGlobalTreeNodeListResponse) Set(val *BTGlobalTreeNodeListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBTGlobalTreeNodeListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBTGlobalTreeNodeListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTGlobalTreeNodeListResponse(val *BTGlobalTreeNodeListResponse) *NullableBTGlobalTreeNodeListResponse {
	return &NullableBTGlobalTreeNodeListResponse{value: val, isSet: true}
}

func (v NullableBTGlobalTreeNodeListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTGlobalTreeNodeListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

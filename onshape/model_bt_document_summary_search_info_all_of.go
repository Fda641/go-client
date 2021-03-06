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

// BTDocumentSummarySearchInfoAllOf struct for BTDocumentSummarySearchInfoAllOf
type BTDocumentSummarySearchInfoAllOf struct {
	SearchHits *[]BTDocumentSearchHitInfo `json:"searchHits,omitempty"`
}

// NewBTDocumentSummarySearchInfoAllOf instantiates a new BTDocumentSummarySearchInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentSummarySearchInfoAllOf() *BTDocumentSummarySearchInfoAllOf {
	this := BTDocumentSummarySearchInfoAllOf{}
	return &this
}

// NewBTDocumentSummarySearchInfoAllOfWithDefaults instantiates a new BTDocumentSummarySearchInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentSummarySearchInfoAllOfWithDefaults() *BTDocumentSummarySearchInfoAllOf {
	this := BTDocumentSummarySearchInfoAllOf{}
	return &this
}

// GetSearchHits returns the SearchHits field value if set, zero value otherwise.
func (o *BTDocumentSummarySearchInfoAllOf) GetSearchHits() []BTDocumentSearchHitInfo {
	if o == nil || o.SearchHits == nil {
		var ret []BTDocumentSearchHitInfo
		return ret
	}
	return *o.SearchHits
}

// GetSearchHitsOk returns a tuple with the SearchHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentSummarySearchInfoAllOf) GetSearchHitsOk() (*[]BTDocumentSearchHitInfo, bool) {
	if o == nil || o.SearchHits == nil {
		return nil, false
	}
	return o.SearchHits, true
}

// HasSearchHits returns a boolean if a field has been set.
func (o *BTDocumentSummarySearchInfoAllOf) HasSearchHits() bool {
	if o != nil && o.SearchHits != nil {
		return true
	}

	return false
}

// SetSearchHits gets a reference to the given []BTDocumentSearchHitInfo and assigns it to the SearchHits field.
func (o *BTDocumentSummarySearchInfoAllOf) SetSearchHits(v []BTDocumentSearchHitInfo) {
	o.SearchHits = &v
}

func (o BTDocumentSummarySearchInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SearchHits != nil {
		toSerialize["searchHits"] = o.SearchHits
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentSummarySearchInfoAllOf struct {
	value *BTDocumentSummarySearchInfoAllOf
	isSet bool
}

func (v NullableBTDocumentSummarySearchInfoAllOf) Get() *BTDocumentSummarySearchInfoAllOf {
	return v.value
}

func (v *NullableBTDocumentSummarySearchInfoAllOf) Set(val *BTDocumentSummarySearchInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentSummarySearchInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentSummarySearchInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentSummarySearchInfoAllOf(val *BTDocumentSummarySearchInfoAllOf) *NullableBTDocumentSummarySearchInfoAllOf {
	return &NullableBTDocumentSummarySearchInfoAllOf{value: val, isSet: true}
}

func (v NullableBTDocumentSummarySearchInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentSummarySearchInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

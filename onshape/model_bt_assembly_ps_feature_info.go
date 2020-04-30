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

// BTAssemblyPSFeatureInfo struct for BTAssemblyPSFeatureInfo
type BTAssemblyPSFeatureInfo struct {
	Configuration *string `json:"configuration,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
	DocumentMicroversion *string `json:"documentMicroversion,omitempty"`
	DocumentVersion *string `json:"documentVersion,omitempty"`
	ElementId *string `json:"elementId,omitempty"`
	FeatureId *string `json:"featureId,omitempty"`
	FeatureType *string `json:"featureType,omitempty"`
	FullConfiguration *string `json:"fullConfiguration,omitempty"`
	Revision *string `json:"revision,omitempty"`
}

// NewBTAssemblyPSFeatureInfo instantiates a new BTAssemblyPSFeatureInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyPSFeatureInfo() *BTAssemblyPSFeatureInfo {
	this := BTAssemblyPSFeatureInfo{}
	return &this
}

// NewBTAssemblyPSFeatureInfoWithDefaults instantiates a new BTAssemblyPSFeatureInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyPSFeatureInfoWithDefaults() *BTAssemblyPSFeatureInfo {
	this := BTAssemblyPSFeatureInfo{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTAssemblyPSFeatureInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTAssemblyPSFeatureInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentMicroversion returns the DocumentMicroversion field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetDocumentMicroversion() string {
	if o == nil || o.DocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.DocumentMicroversion
}

// GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.DocumentMicroversion == nil {
		return nil, false
	}
	return o.DocumentMicroversion, true
}

// HasDocumentMicroversion returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasDocumentMicroversion() bool {
	if o != nil && o.DocumentMicroversion != nil {
		return true
	}

	return false
}

// SetDocumentMicroversion gets a reference to the given string and assigns it to the DocumentMicroversion field.
func (o *BTAssemblyPSFeatureInfo) SetDocumentMicroversion(v string) {
	o.DocumentMicroversion = &v
}

// GetDocumentVersion returns the DocumentVersion field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetDocumentVersion() string {
	if o == nil || o.DocumentVersion == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersion
}

// GetDocumentVersionOk returns a tuple with the DocumentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetDocumentVersionOk() (*string, bool) {
	if o == nil || o.DocumentVersion == nil {
		return nil, false
	}
	return o.DocumentVersion, true
}

// HasDocumentVersion returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasDocumentVersion() bool {
	if o != nil && o.DocumentVersion != nil {
		return true
	}

	return false
}

// SetDocumentVersion gets a reference to the given string and assigns it to the DocumentVersion field.
func (o *BTAssemblyPSFeatureInfo) SetDocumentVersion(v string) {
	o.DocumentVersion = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssemblyPSFeatureInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetFeatureId() string {
	if o == nil || o.FeatureId == nil {
		var ret string
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetFeatureIdOk() (*string, bool) {
	if o == nil || o.FeatureId == nil {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasFeatureId() bool {
	if o != nil && o.FeatureId != nil {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given string and assigns it to the FeatureId field.
func (o *BTAssemblyPSFeatureInfo) SetFeatureId(v string) {
	o.FeatureId = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTAssemblyPSFeatureInfo) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetFullConfiguration returns the FullConfiguration field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetFullConfiguration() string {
	if o == nil || o.FullConfiguration == nil {
		var ret string
		return ret
	}
	return *o.FullConfiguration
}

// GetFullConfigurationOk returns a tuple with the FullConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetFullConfigurationOk() (*string, bool) {
	if o == nil || o.FullConfiguration == nil {
		return nil, false
	}
	return o.FullConfiguration, true
}

// HasFullConfiguration returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasFullConfiguration() bool {
	if o != nil && o.FullConfiguration != nil {
		return true
	}

	return false
}

// SetFullConfiguration gets a reference to the given string and assigns it to the FullConfiguration field.
func (o *BTAssemblyPSFeatureInfo) SetFullConfiguration(v string) {
	o.FullConfiguration = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTAssemblyPSFeatureInfo) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyPSFeatureInfo) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTAssemblyPSFeatureInfo) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTAssemblyPSFeatureInfo) SetRevision(v string) {
	o.Revision = &v
}

func (o BTAssemblyPSFeatureInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentMicroversion != nil {
		toSerialize["documentMicroversion"] = o.DocumentMicroversion
	}
	if o.DocumentVersion != nil {
		toSerialize["documentVersion"] = o.DocumentVersion
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FeatureId != nil {
		toSerialize["featureId"] = o.FeatureId
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.FullConfiguration != nil {
		toSerialize["fullConfiguration"] = o.FullConfiguration
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyPSFeatureInfo struct {
	value *BTAssemblyPSFeatureInfo
	isSet bool
}

func (v NullableBTAssemblyPSFeatureInfo) Get() *BTAssemblyPSFeatureInfo {
	return v.value
}

func (v *NullableBTAssemblyPSFeatureInfo) Set(val *BTAssemblyPSFeatureInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyPSFeatureInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyPSFeatureInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyPSFeatureInfo(val *BTAssemblyPSFeatureInfo) *NullableBTAssemblyPSFeatureInfo {
	return &NullableBTAssemblyPSFeatureInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyPSFeatureInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyPSFeatureInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
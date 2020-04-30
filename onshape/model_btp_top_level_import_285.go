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

// BTPTopLevelImport285 struct for BTPTopLevelImport285
type BTPTopLevelImport285 struct {
	BTPTopLevelNode286
	BtType *string `json:"btType,omitempty"`
	CombinedNamespacePathAndVersion *string `json:"combinedNamespacePathAndVersion,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	ModuleId *BTPModuleId235 `json:"moduleId,omitempty"`
	Namespace *[]BTPIdentifier8 `json:"namespace,omitempty"`
	NamespaceString *string `json:"namespaceString,omitempty"`
	SpaceBeforeImport *BTPSpace10 `json:"spaceBeforeImport,omitempty"`
}

// NewBTPTopLevelImport285 instantiates a new BTPTopLevelImport285 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPTopLevelImport285() *BTPTopLevelImport285 {
	this := BTPTopLevelImport285{}
	return &this
}

// NewBTPTopLevelImport285WithDefaults instantiates a new BTPTopLevelImport285 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPTopLevelImport285WithDefaults() *BTPTopLevelImport285 {
	this := BTPTopLevelImport285{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTPTopLevelImport285) SetBtType(v string) {
	o.BtType = &v
}

// GetCombinedNamespacePathAndVersion returns the CombinedNamespacePathAndVersion field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetCombinedNamespacePathAndVersion() string {
	if o == nil || o.CombinedNamespacePathAndVersion == nil {
		var ret string
		return ret
	}
	return *o.CombinedNamespacePathAndVersion
}

// GetCombinedNamespacePathAndVersionOk returns a tuple with the CombinedNamespacePathAndVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetCombinedNamespacePathAndVersionOk() (*string, bool) {
	if o == nil || o.CombinedNamespacePathAndVersion == nil {
		return nil, false
	}
	return o.CombinedNamespacePathAndVersion, true
}

// HasCombinedNamespacePathAndVersion returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasCombinedNamespacePathAndVersion() bool {
	if o != nil && o.CombinedNamespacePathAndVersion != nil {
		return true
	}

	return false
}

// SetCombinedNamespacePathAndVersion gets a reference to the given string and assigns it to the CombinedNamespacePathAndVersion field.
func (o *BTPTopLevelImport285) SetCombinedNamespacePathAndVersion(v string) {
	o.CombinedNamespacePathAndVersion = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTPTopLevelImport285) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetModuleId() BTPModuleId235 {
	if o == nil || o.ModuleId == nil {
		var ret BTPModuleId235
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetModuleIdOk() (*BTPModuleId235, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given BTPModuleId235 and assigns it to the ModuleId field.
func (o *BTPTopLevelImport285) SetModuleId(v BTPModuleId235) {
	o.ModuleId = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetNamespace() []BTPIdentifier8 {
	if o == nil || o.Namespace == nil {
		var ret []BTPIdentifier8
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetNamespaceOk() (*[]BTPIdentifier8, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given []BTPIdentifier8 and assigns it to the Namespace field.
func (o *BTPTopLevelImport285) SetNamespace(v []BTPIdentifier8) {
	o.Namespace = &v
}

// GetNamespaceString returns the NamespaceString field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetNamespaceString() string {
	if o == nil || o.NamespaceString == nil {
		var ret string
		return ret
	}
	return *o.NamespaceString
}

// GetNamespaceStringOk returns a tuple with the NamespaceString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetNamespaceStringOk() (*string, bool) {
	if o == nil || o.NamespaceString == nil {
		return nil, false
	}
	return o.NamespaceString, true
}

// HasNamespaceString returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasNamespaceString() bool {
	if o != nil && o.NamespaceString != nil {
		return true
	}

	return false
}

// SetNamespaceString gets a reference to the given string and assigns it to the NamespaceString field.
func (o *BTPTopLevelImport285) SetNamespaceString(v string) {
	o.NamespaceString = &v
}

// GetSpaceBeforeImport returns the SpaceBeforeImport field value if set, zero value otherwise.
func (o *BTPTopLevelImport285) GetSpaceBeforeImport() BTPSpace10 {
	if o == nil || o.SpaceBeforeImport == nil {
		var ret BTPSpace10
		return ret
	}
	return *o.SpaceBeforeImport
}

// GetSpaceBeforeImportOk returns a tuple with the SpaceBeforeImport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPTopLevelImport285) GetSpaceBeforeImportOk() (*BTPSpace10, bool) {
	if o == nil || o.SpaceBeforeImport == nil {
		return nil, false
	}
	return o.SpaceBeforeImport, true
}

// HasSpaceBeforeImport returns a boolean if a field has been set.
func (o *BTPTopLevelImport285) HasSpaceBeforeImport() bool {
	if o != nil && o.SpaceBeforeImport != nil {
		return true
	}

	return false
}

// SetSpaceBeforeImport gets a reference to the given BTPSpace10 and assigns it to the SpaceBeforeImport field.
func (o *BTPTopLevelImport285) SetSpaceBeforeImport(v BTPSpace10) {
	o.SpaceBeforeImport = &v
}

func (o BTPTopLevelImport285) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBTPTopLevelNode286, errBTPTopLevelNode286 := json.Marshal(o.BTPTopLevelNode286)
	if errBTPTopLevelNode286 != nil {
		return []byte{}, errBTPTopLevelNode286
	}
	errBTPTopLevelNode286 = json.Unmarshal([]byte(serializedBTPTopLevelNode286), &toSerialize)
	if errBTPTopLevelNode286 != nil {
		return []byte{}, errBTPTopLevelNode286
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CombinedNamespacePathAndVersion != nil {
		toSerialize["combinedNamespacePathAndVersion"] = o.CombinedNamespacePathAndVersion
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.ModuleId != nil {
		toSerialize["moduleId"] = o.ModuleId
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NamespaceString != nil {
		toSerialize["namespaceString"] = o.NamespaceString
	}
	if o.SpaceBeforeImport != nil {
		toSerialize["spaceBeforeImport"] = o.SpaceBeforeImport
	}
	return json.Marshal(toSerialize)
}

type NullableBTPTopLevelImport285 struct {
	value *BTPTopLevelImport285
	isSet bool
}

func (v NullableBTPTopLevelImport285) Get() *BTPTopLevelImport285 {
	return v.value
}

func (v *NullableBTPTopLevelImport285) Set(val *BTPTopLevelImport285) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPTopLevelImport285) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPTopLevelImport285) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPTopLevelImport285(val *BTPTopLevelImport285) *NullableBTPTopLevelImport285 {
	return &NullableBTPTopLevelImport285{value: val, isSet: true}
}

func (v NullableBTPTopLevelImport285) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPTopLevelImport285) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
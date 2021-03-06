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

// BTFeatureSpec129 struct for BTFeatureSpec129
type BTFeatureSpec129 struct {
	AdditionalLocalizedStrings *int32 `json:"additionalLocalizedStrings,omitempty"`
	AllParameters *[]BTParameterSpec6 `json:"allParameters,omitempty"`
	BtType *string `json:"btType,omitempty"`
	EditingLogic *BTEditingLogic2350 `json:"editingLogic,omitempty"`
	FeatureNameTemplate *string `json:"featureNameTemplate,omitempty"`
	FeatureType *string `json:"featureType,omitempty"`
	FeatureTypeName *string `json:"featureTypeName,omitempty"`
	FilterSelectors *[]string `json:"filterSelectors,omitempty"`
	FullFeatureType *string `json:"fullFeatureType,omitempty"`
	Groups *[]BTParameterGroupSpec3469 `json:"groups,omitempty"`
	IconUri *string `json:"iconUri,omitempty"`
	LanguageVersion *int32 `json:"languageVersion,omitempty"`
	LinkedLocationName *string `json:"linkedLocationName,omitempty"`
	LocalizableName *string `json:"localizableName,omitempty"`
	LocalizedName *string `json:"localizedName,omitempty"`
	LocationInfos *[]BTLocationInfo226 `json:"locationInfos,omitempty"`
	ManipulatorChangeFunction *string `json:"manipulatorChangeFunction,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	NamespaceIncludingEnums *string `json:"namespaceIncludingEnums,omitempty"`
	NamespaceTheSource *bool `json:"namespaceTheSource,omitempty"`
	Parameters *[]BTParameterSpec6 `json:"parameters,omitempty"`
	Signature *string `json:"signature,omitempty"`
	SourceLocation *BTLocationInfo226 `json:"sourceLocation,omitempty"`
	SourceMicroversionId *string `json:"sourceMicroversionId,omitempty"`
	StringsToLocalize *[]string `json:"stringsToLocalize,omitempty"`
	TableSpec *bool `json:"tableSpec,omitempty"`
	UiHints *[]string `json:"uiHints,omitempty"`
}

// NewBTFeatureSpec129 instantiates a new BTFeatureSpec129 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureSpec129() *BTFeatureSpec129 {
	this := BTFeatureSpec129{}
	return &this
}

// NewBTFeatureSpec129WithDefaults instantiates a new BTFeatureSpec129 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureSpec129WithDefaults() *BTFeatureSpec129 {
	this := BTFeatureSpec129{}
	return &this
}

// GetAdditionalLocalizedStrings returns the AdditionalLocalizedStrings field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetAdditionalLocalizedStrings() int32 {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		var ret int32
		return ret
	}
	return *o.AdditionalLocalizedStrings
}

// GetAdditionalLocalizedStringsOk returns a tuple with the AdditionalLocalizedStrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetAdditionalLocalizedStringsOk() (*int32, bool) {
	if o == nil || o.AdditionalLocalizedStrings == nil {
		return nil, false
	}
	return o.AdditionalLocalizedStrings, true
}

// HasAdditionalLocalizedStrings returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasAdditionalLocalizedStrings() bool {
	if o != nil && o.AdditionalLocalizedStrings != nil {
		return true
	}

	return false
}

// SetAdditionalLocalizedStrings gets a reference to the given int32 and assigns it to the AdditionalLocalizedStrings field.
func (o *BTFeatureSpec129) SetAdditionalLocalizedStrings(v int32) {
	o.AdditionalLocalizedStrings = &v
}

// GetAllParameters returns the AllParameters field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetAllParameters() []BTParameterSpec6 {
	if o == nil || o.AllParameters == nil {
		var ret []BTParameterSpec6
		return ret
	}
	return *o.AllParameters
}

// GetAllParametersOk returns a tuple with the AllParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetAllParametersOk() (*[]BTParameterSpec6, bool) {
	if o == nil || o.AllParameters == nil {
		return nil, false
	}
	return o.AllParameters, true
}

// HasAllParameters returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasAllParameters() bool {
	if o != nil && o.AllParameters != nil {
		return true
	}

	return false
}

// SetAllParameters gets a reference to the given []BTParameterSpec6 and assigns it to the AllParameters field.
func (o *BTFeatureSpec129) SetAllParameters(v []BTParameterSpec6) {
	o.AllParameters = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFeatureSpec129) SetBtType(v string) {
	o.BtType = &v
}

// GetEditingLogic returns the EditingLogic field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetEditingLogic() BTEditingLogic2350 {
	if o == nil || o.EditingLogic == nil {
		var ret BTEditingLogic2350
		return ret
	}
	return *o.EditingLogic
}

// GetEditingLogicOk returns a tuple with the EditingLogic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetEditingLogicOk() (*BTEditingLogic2350, bool) {
	if o == nil || o.EditingLogic == nil {
		return nil, false
	}
	return o.EditingLogic, true
}

// HasEditingLogic returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasEditingLogic() bool {
	if o != nil && o.EditingLogic != nil {
		return true
	}

	return false
}

// SetEditingLogic gets a reference to the given BTEditingLogic2350 and assigns it to the EditingLogic field.
func (o *BTFeatureSpec129) SetEditingLogic(v BTEditingLogic2350) {
	o.EditingLogic = &v
}

// GetFeatureNameTemplate returns the FeatureNameTemplate field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetFeatureNameTemplate() string {
	if o == nil || o.FeatureNameTemplate == nil {
		var ret string
		return ret
	}
	return *o.FeatureNameTemplate
}

// GetFeatureNameTemplateOk returns a tuple with the FeatureNameTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetFeatureNameTemplateOk() (*string, bool) {
	if o == nil || o.FeatureNameTemplate == nil {
		return nil, false
	}
	return o.FeatureNameTemplate, true
}

// HasFeatureNameTemplate returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasFeatureNameTemplate() bool {
	if o != nil && o.FeatureNameTemplate != nil {
		return true
	}

	return false
}

// SetFeatureNameTemplate gets a reference to the given string and assigns it to the FeatureNameTemplate field.
func (o *BTFeatureSpec129) SetFeatureNameTemplate(v string) {
	o.FeatureNameTemplate = &v
}

// GetFeatureType returns the FeatureType field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetFeatureType() string {
	if o == nil || o.FeatureType == nil {
		var ret string
		return ret
	}
	return *o.FeatureType
}

// GetFeatureTypeOk returns a tuple with the FeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetFeatureTypeOk() (*string, bool) {
	if o == nil || o.FeatureType == nil {
		return nil, false
	}
	return o.FeatureType, true
}

// HasFeatureType returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasFeatureType() bool {
	if o != nil && o.FeatureType != nil {
		return true
	}

	return false
}

// SetFeatureType gets a reference to the given string and assigns it to the FeatureType field.
func (o *BTFeatureSpec129) SetFeatureType(v string) {
	o.FeatureType = &v
}

// GetFeatureTypeName returns the FeatureTypeName field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetFeatureTypeName() string {
	if o == nil || o.FeatureTypeName == nil {
		var ret string
		return ret
	}
	return *o.FeatureTypeName
}

// GetFeatureTypeNameOk returns a tuple with the FeatureTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetFeatureTypeNameOk() (*string, bool) {
	if o == nil || o.FeatureTypeName == nil {
		return nil, false
	}
	return o.FeatureTypeName, true
}

// HasFeatureTypeName returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasFeatureTypeName() bool {
	if o != nil && o.FeatureTypeName != nil {
		return true
	}

	return false
}

// SetFeatureTypeName gets a reference to the given string and assigns it to the FeatureTypeName field.
func (o *BTFeatureSpec129) SetFeatureTypeName(v string) {
	o.FeatureTypeName = &v
}

// GetFilterSelectors returns the FilterSelectors field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetFilterSelectors() []string {
	if o == nil || o.FilterSelectors == nil {
		var ret []string
		return ret
	}
	return *o.FilterSelectors
}

// GetFilterSelectorsOk returns a tuple with the FilterSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetFilterSelectorsOk() (*[]string, bool) {
	if o == nil || o.FilterSelectors == nil {
		return nil, false
	}
	return o.FilterSelectors, true
}

// HasFilterSelectors returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasFilterSelectors() bool {
	if o != nil && o.FilterSelectors != nil {
		return true
	}

	return false
}

// SetFilterSelectors gets a reference to the given []string and assigns it to the FilterSelectors field.
func (o *BTFeatureSpec129) SetFilterSelectors(v []string) {
	o.FilterSelectors = &v
}

// GetFullFeatureType returns the FullFeatureType field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetFullFeatureType() string {
	if o == nil || o.FullFeatureType == nil {
		var ret string
		return ret
	}
	return *o.FullFeatureType
}

// GetFullFeatureTypeOk returns a tuple with the FullFeatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetFullFeatureTypeOk() (*string, bool) {
	if o == nil || o.FullFeatureType == nil {
		return nil, false
	}
	return o.FullFeatureType, true
}

// HasFullFeatureType returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasFullFeatureType() bool {
	if o != nil && o.FullFeatureType != nil {
		return true
	}

	return false
}

// SetFullFeatureType gets a reference to the given string and assigns it to the FullFeatureType field.
func (o *BTFeatureSpec129) SetFullFeatureType(v string) {
	o.FullFeatureType = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetGroups() []BTParameterGroupSpec3469 {
	if o == nil || o.Groups == nil {
		var ret []BTParameterGroupSpec3469
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetGroupsOk() (*[]BTParameterGroupSpec3469, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []BTParameterGroupSpec3469 and assigns it to the Groups field.
func (o *BTFeatureSpec129) SetGroups(v []BTParameterGroupSpec3469) {
	o.Groups = &v
}

// GetIconUri returns the IconUri field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetIconUri() string {
	if o == nil || o.IconUri == nil {
		var ret string
		return ret
	}
	return *o.IconUri
}

// GetIconUriOk returns a tuple with the IconUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetIconUriOk() (*string, bool) {
	if o == nil || o.IconUri == nil {
		return nil, false
	}
	return o.IconUri, true
}

// HasIconUri returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasIconUri() bool {
	if o != nil && o.IconUri != nil {
		return true
	}

	return false
}

// SetIconUri gets a reference to the given string and assigns it to the IconUri field.
func (o *BTFeatureSpec129) SetIconUri(v string) {
	o.IconUri = &v
}

// GetLanguageVersion returns the LanguageVersion field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetLanguageVersion() int32 {
	if o == nil || o.LanguageVersion == nil {
		var ret int32
		return ret
	}
	return *o.LanguageVersion
}

// GetLanguageVersionOk returns a tuple with the LanguageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetLanguageVersionOk() (*int32, bool) {
	if o == nil || o.LanguageVersion == nil {
		return nil, false
	}
	return o.LanguageVersion, true
}

// HasLanguageVersion returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasLanguageVersion() bool {
	if o != nil && o.LanguageVersion != nil {
		return true
	}

	return false
}

// SetLanguageVersion gets a reference to the given int32 and assigns it to the LanguageVersion field.
func (o *BTFeatureSpec129) SetLanguageVersion(v int32) {
	o.LanguageVersion = &v
}

// GetLinkedLocationName returns the LinkedLocationName field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetLinkedLocationName() string {
	if o == nil || o.LinkedLocationName == nil {
		var ret string
		return ret
	}
	return *o.LinkedLocationName
}

// GetLinkedLocationNameOk returns a tuple with the LinkedLocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetLinkedLocationNameOk() (*string, bool) {
	if o == nil || o.LinkedLocationName == nil {
		return nil, false
	}
	return o.LinkedLocationName, true
}

// HasLinkedLocationName returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasLinkedLocationName() bool {
	if o != nil && o.LinkedLocationName != nil {
		return true
	}

	return false
}

// SetLinkedLocationName gets a reference to the given string and assigns it to the LinkedLocationName field.
func (o *BTFeatureSpec129) SetLinkedLocationName(v string) {
	o.LinkedLocationName = &v
}

// GetLocalizableName returns the LocalizableName field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetLocalizableName() string {
	if o == nil || o.LocalizableName == nil {
		var ret string
		return ret
	}
	return *o.LocalizableName
}

// GetLocalizableNameOk returns a tuple with the LocalizableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetLocalizableNameOk() (*string, bool) {
	if o == nil || o.LocalizableName == nil {
		return nil, false
	}
	return o.LocalizableName, true
}

// HasLocalizableName returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasLocalizableName() bool {
	if o != nil && o.LocalizableName != nil {
		return true
	}

	return false
}

// SetLocalizableName gets a reference to the given string and assigns it to the LocalizableName field.
func (o *BTFeatureSpec129) SetLocalizableName(v string) {
	o.LocalizableName = &v
}

// GetLocalizedName returns the LocalizedName field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetLocalizedName() string {
	if o == nil || o.LocalizedName == nil {
		var ret string
		return ret
	}
	return *o.LocalizedName
}

// GetLocalizedNameOk returns a tuple with the LocalizedName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetLocalizedNameOk() (*string, bool) {
	if o == nil || o.LocalizedName == nil {
		return nil, false
	}
	return o.LocalizedName, true
}

// HasLocalizedName returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasLocalizedName() bool {
	if o != nil && o.LocalizedName != nil {
		return true
	}

	return false
}

// SetLocalizedName gets a reference to the given string and assigns it to the LocalizedName field.
func (o *BTFeatureSpec129) SetLocalizedName(v string) {
	o.LocalizedName = &v
}

// GetLocationInfos returns the LocationInfos field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetLocationInfos() []BTLocationInfo226 {
	if o == nil || o.LocationInfos == nil {
		var ret []BTLocationInfo226
		return ret
	}
	return *o.LocationInfos
}

// GetLocationInfosOk returns a tuple with the LocationInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetLocationInfosOk() (*[]BTLocationInfo226, bool) {
	if o == nil || o.LocationInfos == nil {
		return nil, false
	}
	return o.LocationInfos, true
}

// HasLocationInfos returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasLocationInfos() bool {
	if o != nil && o.LocationInfos != nil {
		return true
	}

	return false
}

// SetLocationInfos gets a reference to the given []BTLocationInfo226 and assigns it to the LocationInfos field.
func (o *BTFeatureSpec129) SetLocationInfos(v []BTLocationInfo226) {
	o.LocationInfos = &v
}

// GetManipulatorChangeFunction returns the ManipulatorChangeFunction field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetManipulatorChangeFunction() string {
	if o == nil || o.ManipulatorChangeFunction == nil {
		var ret string
		return ret
	}
	return *o.ManipulatorChangeFunction
}

// GetManipulatorChangeFunctionOk returns a tuple with the ManipulatorChangeFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetManipulatorChangeFunctionOk() (*string, bool) {
	if o == nil || o.ManipulatorChangeFunction == nil {
		return nil, false
	}
	return o.ManipulatorChangeFunction, true
}

// HasManipulatorChangeFunction returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasManipulatorChangeFunction() bool {
	if o != nil && o.ManipulatorChangeFunction != nil {
		return true
	}

	return false
}

// SetManipulatorChangeFunction gets a reference to the given string and assigns it to the ManipulatorChangeFunction field.
func (o *BTFeatureSpec129) SetManipulatorChangeFunction(v string) {
	o.ManipulatorChangeFunction = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTFeatureSpec129) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNamespaceIncludingEnums returns the NamespaceIncludingEnums field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetNamespaceIncludingEnums() string {
	if o == nil || o.NamespaceIncludingEnums == nil {
		var ret string
		return ret
	}
	return *o.NamespaceIncludingEnums
}

// GetNamespaceIncludingEnumsOk returns a tuple with the NamespaceIncludingEnums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetNamespaceIncludingEnumsOk() (*string, bool) {
	if o == nil || o.NamespaceIncludingEnums == nil {
		return nil, false
	}
	return o.NamespaceIncludingEnums, true
}

// HasNamespaceIncludingEnums returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasNamespaceIncludingEnums() bool {
	if o != nil && o.NamespaceIncludingEnums != nil {
		return true
	}

	return false
}

// SetNamespaceIncludingEnums gets a reference to the given string and assigns it to the NamespaceIncludingEnums field.
func (o *BTFeatureSpec129) SetNamespaceIncludingEnums(v string) {
	o.NamespaceIncludingEnums = &v
}

// GetNamespaceTheSource returns the NamespaceTheSource field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetNamespaceTheSource() bool {
	if o == nil || o.NamespaceTheSource == nil {
		var ret bool
		return ret
	}
	return *o.NamespaceTheSource
}

// GetNamespaceTheSourceOk returns a tuple with the NamespaceTheSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetNamespaceTheSourceOk() (*bool, bool) {
	if o == nil || o.NamespaceTheSource == nil {
		return nil, false
	}
	return o.NamespaceTheSource, true
}

// HasNamespaceTheSource returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasNamespaceTheSource() bool {
	if o != nil && o.NamespaceTheSource != nil {
		return true
	}

	return false
}

// SetNamespaceTheSource gets a reference to the given bool and assigns it to the NamespaceTheSource field.
func (o *BTFeatureSpec129) SetNamespaceTheSource(v bool) {
	o.NamespaceTheSource = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetParameters() []BTParameterSpec6 {
	if o == nil || o.Parameters == nil {
		var ret []BTParameterSpec6
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetParametersOk() (*[]BTParameterSpec6, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTParameterSpec6 and assigns it to the Parameters field.
func (o *BTFeatureSpec129) SetParameters(v []BTParameterSpec6) {
	o.Parameters = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *BTFeatureSpec129) SetSignature(v string) {
	o.Signature = &v
}

// GetSourceLocation returns the SourceLocation field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetSourceLocation() BTLocationInfo226 {
	if o == nil || o.SourceLocation == nil {
		var ret BTLocationInfo226
		return ret
	}
	return *o.SourceLocation
}

// GetSourceLocationOk returns a tuple with the SourceLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetSourceLocationOk() (*BTLocationInfo226, bool) {
	if o == nil || o.SourceLocation == nil {
		return nil, false
	}
	return o.SourceLocation, true
}

// HasSourceLocation returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasSourceLocation() bool {
	if o != nil && o.SourceLocation != nil {
		return true
	}

	return false
}

// SetSourceLocation gets a reference to the given BTLocationInfo226 and assigns it to the SourceLocation field.
func (o *BTFeatureSpec129) SetSourceLocation(v BTLocationInfo226) {
	o.SourceLocation = &v
}

// GetSourceMicroversionId returns the SourceMicroversionId field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetSourceMicroversionId() string {
	if o == nil || o.SourceMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversionId
}

// GetSourceMicroversionIdOk returns a tuple with the SourceMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetSourceMicroversionIdOk() (*string, bool) {
	if o == nil || o.SourceMicroversionId == nil {
		return nil, false
	}
	return o.SourceMicroversionId, true
}

// HasSourceMicroversionId returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasSourceMicroversionId() bool {
	if o != nil && o.SourceMicroversionId != nil {
		return true
	}

	return false
}

// SetSourceMicroversionId gets a reference to the given string and assigns it to the SourceMicroversionId field.
func (o *BTFeatureSpec129) SetSourceMicroversionId(v string) {
	o.SourceMicroversionId = &v
}

// GetStringsToLocalize returns the StringsToLocalize field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetStringsToLocalize() []string {
	if o == nil || o.StringsToLocalize == nil {
		var ret []string
		return ret
	}
	return *o.StringsToLocalize
}

// GetStringsToLocalizeOk returns a tuple with the StringsToLocalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetStringsToLocalizeOk() (*[]string, bool) {
	if o == nil || o.StringsToLocalize == nil {
		return nil, false
	}
	return o.StringsToLocalize, true
}

// HasStringsToLocalize returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasStringsToLocalize() bool {
	if o != nil && o.StringsToLocalize != nil {
		return true
	}

	return false
}

// SetStringsToLocalize gets a reference to the given []string and assigns it to the StringsToLocalize field.
func (o *BTFeatureSpec129) SetStringsToLocalize(v []string) {
	o.StringsToLocalize = &v
}

// GetTableSpec returns the TableSpec field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetTableSpec() bool {
	if o == nil || o.TableSpec == nil {
		var ret bool
		return ret
	}
	return *o.TableSpec
}

// GetTableSpecOk returns a tuple with the TableSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetTableSpecOk() (*bool, bool) {
	if o == nil || o.TableSpec == nil {
		return nil, false
	}
	return o.TableSpec, true
}

// HasTableSpec returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasTableSpec() bool {
	if o != nil && o.TableSpec != nil {
		return true
	}

	return false
}

// SetTableSpec gets a reference to the given bool and assigns it to the TableSpec field.
func (o *BTFeatureSpec129) SetTableSpec(v bool) {
	o.TableSpec = &v
}

// GetUiHints returns the UiHints field value if set, zero value otherwise.
func (o *BTFeatureSpec129) GetUiHints() []string {
	if o == nil || o.UiHints == nil {
		var ret []string
		return ret
	}
	return *o.UiHints
}

// GetUiHintsOk returns a tuple with the UiHints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureSpec129) GetUiHintsOk() (*[]string, bool) {
	if o == nil || o.UiHints == nil {
		return nil, false
	}
	return o.UiHints, true
}

// HasUiHints returns a boolean if a field has been set.
func (o *BTFeatureSpec129) HasUiHints() bool {
	if o != nil && o.UiHints != nil {
		return true
	}

	return false
}

// SetUiHints gets a reference to the given []string and assigns it to the UiHints field.
func (o *BTFeatureSpec129) SetUiHints(v []string) {
	o.UiHints = &v
}

func (o BTFeatureSpec129) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalLocalizedStrings != nil {
		toSerialize["additionalLocalizedStrings"] = o.AdditionalLocalizedStrings
	}
	if o.AllParameters != nil {
		toSerialize["allParameters"] = o.AllParameters
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EditingLogic != nil {
		toSerialize["editingLogic"] = o.EditingLogic
	}
	if o.FeatureNameTemplate != nil {
		toSerialize["featureNameTemplate"] = o.FeatureNameTemplate
	}
	if o.FeatureType != nil {
		toSerialize["featureType"] = o.FeatureType
	}
	if o.FeatureTypeName != nil {
		toSerialize["featureTypeName"] = o.FeatureTypeName
	}
	if o.FilterSelectors != nil {
		toSerialize["filterSelectors"] = o.FilterSelectors
	}
	if o.FullFeatureType != nil {
		toSerialize["fullFeatureType"] = o.FullFeatureType
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.IconUri != nil {
		toSerialize["iconUri"] = o.IconUri
	}
	if o.LanguageVersion != nil {
		toSerialize["languageVersion"] = o.LanguageVersion
	}
	if o.LinkedLocationName != nil {
		toSerialize["linkedLocationName"] = o.LinkedLocationName
	}
	if o.LocalizableName != nil {
		toSerialize["localizableName"] = o.LocalizableName
	}
	if o.LocalizedName != nil {
		toSerialize["localizedName"] = o.LocalizedName
	}
	if o.LocationInfos != nil {
		toSerialize["locationInfos"] = o.LocationInfos
	}
	if o.ManipulatorChangeFunction != nil {
		toSerialize["manipulatorChangeFunction"] = o.ManipulatorChangeFunction
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NamespaceIncludingEnums != nil {
		toSerialize["namespaceIncludingEnums"] = o.NamespaceIncludingEnums
	}
	if o.NamespaceTheSource != nil {
		toSerialize["namespaceTheSource"] = o.NamespaceTheSource
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.SourceLocation != nil {
		toSerialize["sourceLocation"] = o.SourceLocation
	}
	if o.SourceMicroversionId != nil {
		toSerialize["sourceMicroversionId"] = o.SourceMicroversionId
	}
	if o.StringsToLocalize != nil {
		toSerialize["stringsToLocalize"] = o.StringsToLocalize
	}
	if o.TableSpec != nil {
		toSerialize["tableSpec"] = o.TableSpec
	}
	if o.UiHints != nil {
		toSerialize["uiHints"] = o.UiHints
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureSpec129 struct {
	value *BTFeatureSpec129
	isSet bool
}

func (v NullableBTFeatureSpec129) Get() *BTFeatureSpec129 {
	return v.value
}

func (v *NullableBTFeatureSpec129) Set(val *BTFeatureSpec129) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureSpec129) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureSpec129) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureSpec129(val *BTFeatureSpec129) *NullableBTFeatureSpec129 {
	return &NullableBTFeatureSpec129{value: val, isSet: true}
}

func (v NullableBTFeatureSpec129) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureSpec129) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

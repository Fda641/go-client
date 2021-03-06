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

// BTActiveWorkflowInfo struct for BTActiveWorkflowInfo
type BTActiveWorkflowInfo struct {
	CanCreateReleases *bool `json:"canCreateReleases,omitempty"`
	CanCurrentUserCreateReleases *bool `json:"canCurrentUserCreateReleases,omitempty"`
	CompanyId *string `json:"companyId,omitempty"`
	DocumentId *string `json:"documentId,omitempty"`
	DrawingCanDuplicatePartNumber *bool `json:"drawingCanDuplicatePartNumber,omitempty"`
	EnabledActiveMultipleWorkflows *bool `json:"enabledActiveMultipleWorkflows,omitempty"`
	ObsoletionWorkflow *BTPublishedWorkflowInfo `json:"obsoletionWorkflow,omitempty"`
	ObsoletionWorkflowId *string `json:"obsoletionWorkflowId,omitempty"`
	PickableWorkflows *[]BTPublishedWorkflowInfo `json:"pickableWorkflows,omitempty"`
	ReleaseWorkflow *BTPublishedWorkflowInfo `json:"releaseWorkflow,omitempty"`
	ReleaseWorkflowId *string `json:"releaseWorkflowId,omitempty"`
	UsingAutoPartNumbering *bool `json:"usingAutoPartNumbering,omitempty"`
	UsingAutoPartNumberingScheme *bool `json:"usingAutoPartNumberingScheme,omitempty"`
	UsingManagedWorkflow *bool `json:"usingManagedWorkflow,omitempty"`
}

// NewBTActiveWorkflowInfo instantiates a new BTActiveWorkflowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTActiveWorkflowInfo() *BTActiveWorkflowInfo {
	this := BTActiveWorkflowInfo{}
	return &this
}

// NewBTActiveWorkflowInfoWithDefaults instantiates a new BTActiveWorkflowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTActiveWorkflowInfoWithDefaults() *BTActiveWorkflowInfo {
	this := BTActiveWorkflowInfo{}
	return &this
}

// GetCanCreateReleases returns the CanCreateReleases field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCreateReleases() bool {
	if o == nil || o.CanCreateReleases == nil {
		var ret bool
		return ret
	}
	return *o.CanCreateReleases
}

// GetCanCreateReleasesOk returns a tuple with the CanCreateReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCreateReleasesOk() (*bool, bool) {
	if o == nil || o.CanCreateReleases == nil {
		return nil, false
	}
	return o.CanCreateReleases, true
}

// HasCanCreateReleases returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCreateReleases() bool {
	if o != nil && o.CanCreateReleases != nil {
		return true
	}

	return false
}

// SetCanCreateReleases gets a reference to the given bool and assigns it to the CanCreateReleases field.
func (o *BTActiveWorkflowInfo) SetCanCreateReleases(v bool) {
	o.CanCreateReleases = &v
}

// GetCanCurrentUserCreateReleases returns the CanCurrentUserCreateReleases field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleases() bool {
	if o == nil || o.CanCurrentUserCreateReleases == nil {
		var ret bool
		return ret
	}
	return *o.CanCurrentUserCreateReleases
}

// GetCanCurrentUserCreateReleasesOk returns a tuple with the CanCurrentUserCreateReleases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCanCurrentUserCreateReleasesOk() (*bool, bool) {
	if o == nil || o.CanCurrentUserCreateReleases == nil {
		return nil, false
	}
	return o.CanCurrentUserCreateReleases, true
}

// HasCanCurrentUserCreateReleases returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCanCurrentUserCreateReleases() bool {
	if o != nil && o.CanCurrentUserCreateReleases != nil {
		return true
	}

	return false
}

// SetCanCurrentUserCreateReleases gets a reference to the given bool and assigns it to the CanCurrentUserCreateReleases field.
func (o *BTActiveWorkflowInfo) SetCanCurrentUserCreateReleases(v bool) {
	o.CanCurrentUserCreateReleases = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTActiveWorkflowInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTActiveWorkflowInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDrawingCanDuplicatePartNumber returns the DrawingCanDuplicatePartNumber field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumber() bool {
	if o == nil || o.DrawingCanDuplicatePartNumber == nil {
		var ret bool
		return ret
	}
	return *o.DrawingCanDuplicatePartNumber
}

// GetDrawingCanDuplicatePartNumberOk returns a tuple with the DrawingCanDuplicatePartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetDrawingCanDuplicatePartNumberOk() (*bool, bool) {
	if o == nil || o.DrawingCanDuplicatePartNumber == nil {
		return nil, false
	}
	return o.DrawingCanDuplicatePartNumber, true
}

// HasDrawingCanDuplicatePartNumber returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasDrawingCanDuplicatePartNumber() bool {
	if o != nil && o.DrawingCanDuplicatePartNumber != nil {
		return true
	}

	return false
}

// SetDrawingCanDuplicatePartNumber gets a reference to the given bool and assigns it to the DrawingCanDuplicatePartNumber field.
func (o *BTActiveWorkflowInfo) SetDrawingCanDuplicatePartNumber(v bool) {
	o.DrawingCanDuplicatePartNumber = &v
}

// GetEnabledActiveMultipleWorkflows returns the EnabledActiveMultipleWorkflows field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflows() bool {
	if o == nil || o.EnabledActiveMultipleWorkflows == nil {
		var ret bool
		return ret
	}
	return *o.EnabledActiveMultipleWorkflows
}

// GetEnabledActiveMultipleWorkflowsOk returns a tuple with the EnabledActiveMultipleWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetEnabledActiveMultipleWorkflowsOk() (*bool, bool) {
	if o == nil || o.EnabledActiveMultipleWorkflows == nil {
		return nil, false
	}
	return o.EnabledActiveMultipleWorkflows, true
}

// HasEnabledActiveMultipleWorkflows returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasEnabledActiveMultipleWorkflows() bool {
	if o != nil && o.EnabledActiveMultipleWorkflows != nil {
		return true
	}

	return false
}

// SetEnabledActiveMultipleWorkflows gets a reference to the given bool and assigns it to the EnabledActiveMultipleWorkflows field.
func (o *BTActiveWorkflowInfo) SetEnabledActiveMultipleWorkflows(v bool) {
	o.EnabledActiveMultipleWorkflows = &v
}

// GetObsoletionWorkflow returns the ObsoletionWorkflow field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflow() BTPublishedWorkflowInfo {
	if o == nil || o.ObsoletionWorkflow == nil {
		var ret BTPublishedWorkflowInfo
		return ret
	}
	return *o.ObsoletionWorkflow
}

// GetObsoletionWorkflowOk returns a tuple with the ObsoletionWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowOk() (*BTPublishedWorkflowInfo, bool) {
	if o == nil || o.ObsoletionWorkflow == nil {
		return nil, false
	}
	return o.ObsoletionWorkflow, true
}

// HasObsoletionWorkflow returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasObsoletionWorkflow() bool {
	if o != nil && o.ObsoletionWorkflow != nil {
		return true
	}

	return false
}

// SetObsoletionWorkflow gets a reference to the given BTPublishedWorkflowInfo and assigns it to the ObsoletionWorkflow field.
func (o *BTActiveWorkflowInfo) SetObsoletionWorkflow(v BTPublishedWorkflowInfo) {
	o.ObsoletionWorkflow = &v
}

// GetObsoletionWorkflowId returns the ObsoletionWorkflowId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowId() string {
	if o == nil || o.ObsoletionWorkflowId == nil {
		var ret string
		return ret
	}
	return *o.ObsoletionWorkflowId
}

// GetObsoletionWorkflowIdOk returns a tuple with the ObsoletionWorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetObsoletionWorkflowIdOk() (*string, bool) {
	if o == nil || o.ObsoletionWorkflowId == nil {
		return nil, false
	}
	return o.ObsoletionWorkflowId, true
}

// HasObsoletionWorkflowId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasObsoletionWorkflowId() bool {
	if o != nil && o.ObsoletionWorkflowId != nil {
		return true
	}

	return false
}

// SetObsoletionWorkflowId gets a reference to the given string and assigns it to the ObsoletionWorkflowId field.
func (o *BTActiveWorkflowInfo) SetObsoletionWorkflowId(v string) {
	o.ObsoletionWorkflowId = &v
}

// GetPickableWorkflows returns the PickableWorkflows field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetPickableWorkflows() []BTPublishedWorkflowInfo {
	if o == nil || o.PickableWorkflows == nil {
		var ret []BTPublishedWorkflowInfo
		return ret
	}
	return *o.PickableWorkflows
}

// GetPickableWorkflowsOk returns a tuple with the PickableWorkflows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetPickableWorkflowsOk() (*[]BTPublishedWorkflowInfo, bool) {
	if o == nil || o.PickableWorkflows == nil {
		return nil, false
	}
	return o.PickableWorkflows, true
}

// HasPickableWorkflows returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasPickableWorkflows() bool {
	if o != nil && o.PickableWorkflows != nil {
		return true
	}

	return false
}

// SetPickableWorkflows gets a reference to the given []BTPublishedWorkflowInfo and assigns it to the PickableWorkflows field.
func (o *BTActiveWorkflowInfo) SetPickableWorkflows(v []BTPublishedWorkflowInfo) {
	o.PickableWorkflows = &v
}

// GetReleaseWorkflow returns the ReleaseWorkflow field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflow() BTPublishedWorkflowInfo {
	if o == nil || o.ReleaseWorkflow == nil {
		var ret BTPublishedWorkflowInfo
		return ret
	}
	return *o.ReleaseWorkflow
}

// GetReleaseWorkflowOk returns a tuple with the ReleaseWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflowOk() (*BTPublishedWorkflowInfo, bool) {
	if o == nil || o.ReleaseWorkflow == nil {
		return nil, false
	}
	return o.ReleaseWorkflow, true
}

// HasReleaseWorkflow returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasReleaseWorkflow() bool {
	if o != nil && o.ReleaseWorkflow != nil {
		return true
	}

	return false
}

// SetReleaseWorkflow gets a reference to the given BTPublishedWorkflowInfo and assigns it to the ReleaseWorkflow field.
func (o *BTActiveWorkflowInfo) SetReleaseWorkflow(v BTPublishedWorkflowInfo) {
	o.ReleaseWorkflow = &v
}

// GetReleaseWorkflowId returns the ReleaseWorkflowId field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflowId() string {
	if o == nil || o.ReleaseWorkflowId == nil {
		var ret string
		return ret
	}
	return *o.ReleaseWorkflowId
}

// GetReleaseWorkflowIdOk returns a tuple with the ReleaseWorkflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetReleaseWorkflowIdOk() (*string, bool) {
	if o == nil || o.ReleaseWorkflowId == nil {
		return nil, false
	}
	return o.ReleaseWorkflowId, true
}

// HasReleaseWorkflowId returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasReleaseWorkflowId() bool {
	if o != nil && o.ReleaseWorkflowId != nil {
		return true
	}

	return false
}

// SetReleaseWorkflowId gets a reference to the given string and assigns it to the ReleaseWorkflowId field.
func (o *BTActiveWorkflowInfo) SetReleaseWorkflowId(v string) {
	o.ReleaseWorkflowId = &v
}

// GetUsingAutoPartNumbering returns the UsingAutoPartNumbering field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumbering() bool {
	if o == nil || o.UsingAutoPartNumbering == nil {
		var ret bool
		return ret
	}
	return *o.UsingAutoPartNumbering
}

// GetUsingAutoPartNumberingOk returns a tuple with the UsingAutoPartNumbering field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingOk() (*bool, bool) {
	if o == nil || o.UsingAutoPartNumbering == nil {
		return nil, false
	}
	return o.UsingAutoPartNumbering, true
}

// HasUsingAutoPartNumbering returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasUsingAutoPartNumbering() bool {
	if o != nil && o.UsingAutoPartNumbering != nil {
		return true
	}

	return false
}

// SetUsingAutoPartNumbering gets a reference to the given bool and assigns it to the UsingAutoPartNumbering field.
func (o *BTActiveWorkflowInfo) SetUsingAutoPartNumbering(v bool) {
	o.UsingAutoPartNumbering = &v
}

// GetUsingAutoPartNumberingScheme returns the UsingAutoPartNumberingScheme field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingScheme() bool {
	if o == nil || o.UsingAutoPartNumberingScheme == nil {
		var ret bool
		return ret
	}
	return *o.UsingAutoPartNumberingScheme
}

// GetUsingAutoPartNumberingSchemeOk returns a tuple with the UsingAutoPartNumberingScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetUsingAutoPartNumberingSchemeOk() (*bool, bool) {
	if o == nil || o.UsingAutoPartNumberingScheme == nil {
		return nil, false
	}
	return o.UsingAutoPartNumberingScheme, true
}

// HasUsingAutoPartNumberingScheme returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasUsingAutoPartNumberingScheme() bool {
	if o != nil && o.UsingAutoPartNumberingScheme != nil {
		return true
	}

	return false
}

// SetUsingAutoPartNumberingScheme gets a reference to the given bool and assigns it to the UsingAutoPartNumberingScheme field.
func (o *BTActiveWorkflowInfo) SetUsingAutoPartNumberingScheme(v bool) {
	o.UsingAutoPartNumberingScheme = &v
}

// GetUsingManagedWorkflow returns the UsingManagedWorkflow field value if set, zero value otherwise.
func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflow() bool {
	if o == nil || o.UsingManagedWorkflow == nil {
		var ret bool
		return ret
	}
	return *o.UsingManagedWorkflow
}

// GetUsingManagedWorkflowOk returns a tuple with the UsingManagedWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTActiveWorkflowInfo) GetUsingManagedWorkflowOk() (*bool, bool) {
	if o == nil || o.UsingManagedWorkflow == nil {
		return nil, false
	}
	return o.UsingManagedWorkflow, true
}

// HasUsingManagedWorkflow returns a boolean if a field has been set.
func (o *BTActiveWorkflowInfo) HasUsingManagedWorkflow() bool {
	if o != nil && o.UsingManagedWorkflow != nil {
		return true
	}

	return false
}

// SetUsingManagedWorkflow gets a reference to the given bool and assigns it to the UsingManagedWorkflow field.
func (o *BTActiveWorkflowInfo) SetUsingManagedWorkflow(v bool) {
	o.UsingManagedWorkflow = &v
}

func (o BTActiveWorkflowInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanCreateReleases != nil {
		toSerialize["canCreateReleases"] = o.CanCreateReleases
	}
	if o.CanCurrentUserCreateReleases != nil {
		toSerialize["canCurrentUserCreateReleases"] = o.CanCurrentUserCreateReleases
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DrawingCanDuplicatePartNumber != nil {
		toSerialize["drawingCanDuplicatePartNumber"] = o.DrawingCanDuplicatePartNumber
	}
	if o.EnabledActiveMultipleWorkflows != nil {
		toSerialize["enabledActiveMultipleWorkflows"] = o.EnabledActiveMultipleWorkflows
	}
	if o.ObsoletionWorkflow != nil {
		toSerialize["obsoletionWorkflow"] = o.ObsoletionWorkflow
	}
	if o.ObsoletionWorkflowId != nil {
		toSerialize["obsoletionWorkflowId"] = o.ObsoletionWorkflowId
	}
	if o.PickableWorkflows != nil {
		toSerialize["pickableWorkflows"] = o.PickableWorkflows
	}
	if o.ReleaseWorkflow != nil {
		toSerialize["releaseWorkflow"] = o.ReleaseWorkflow
	}
	if o.ReleaseWorkflowId != nil {
		toSerialize["releaseWorkflowId"] = o.ReleaseWorkflowId
	}
	if o.UsingAutoPartNumbering != nil {
		toSerialize["usingAutoPartNumbering"] = o.UsingAutoPartNumbering
	}
	if o.UsingAutoPartNumberingScheme != nil {
		toSerialize["usingAutoPartNumberingScheme"] = o.UsingAutoPartNumberingScheme
	}
	if o.UsingManagedWorkflow != nil {
		toSerialize["usingManagedWorkflow"] = o.UsingManagedWorkflow
	}
	return json.Marshal(toSerialize)
}

type NullableBTActiveWorkflowInfo struct {
	value *BTActiveWorkflowInfo
	isSet bool
}

func (v NullableBTActiveWorkflowInfo) Get() *BTActiveWorkflowInfo {
	return v.value
}

func (v *NullableBTActiveWorkflowInfo) Set(val *BTActiveWorkflowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTActiveWorkflowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTActiveWorkflowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTActiveWorkflowInfo(val *BTActiveWorkflowInfo) *NullableBTActiveWorkflowInfo {
	return &NullableBTActiveWorkflowInfo{value: val, isSet: true}
}

func (v NullableBTActiveWorkflowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTActiveWorkflowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

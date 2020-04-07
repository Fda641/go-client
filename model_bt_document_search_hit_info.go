/*
 * Onshape REST API
 *
 * The Onshape REST API consumed by all clients.
 *
 * API version: 1.111
 * Contact: api-support@onshape.zendesk.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BtDocumentSearchHitInfo struct for BtDocumentSearchHitInfo
type BtDocumentSearchHitInfo struct {
	DocumentId string `json:"documentId,omitempty"`
	ElementName string `json:"elementName,omitempty"`
	HighlightedFields map[string][]string `json:"highlightedFields,omitempty"`
	Name string `json:"name,omitempty"`
	SourceMap map[string]map[string]interface{} `json:"sourceMap,omitempty"`
	Type string `json:"type,omitempty"`
	VersionOrWorkspaceName string `json:"versionOrWorkspaceName,omitempty"`
}
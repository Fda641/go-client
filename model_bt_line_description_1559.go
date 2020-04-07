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
// BtLineDescription1559 struct for BtLineDescription1559
type BtLineDescription1559 struct {
	BtType string `json:"btType,omitempty"`
	Type string `json:"type,omitempty"`
	Direction BtVector3d389 `json:"direction,omitempty"`
	Origin BtVector3d389 `json:"origin,omitempty"`
}
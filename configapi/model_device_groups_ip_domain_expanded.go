/*
 * Connectivity Service Configuration
 *
 * APIs to configure connectivity service in Aether Network
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package configapi

// DeviceGroupsIpDomainExpanded - This is APN for device
type DeviceGroupsIpDomainExpanded struct {

	Dnn string `json:"dnn,omitempty"`

	UeIpPool string `json:"ue-ip-pool,omitempty"`

	DnsPrimary string `json:"dns-primary,omitempty"`

	Mtu int32 `json:"mtu,omitempty"`
}

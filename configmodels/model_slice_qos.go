// SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Connectivity Service Configuration
 *
 * APIs to configure connectivity service in Aether Network
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package configmodels

type SliceQos struct {
	// uplink data rate in bps
	Uplink int32 `json:"uplink,omitempty"`

	// downlink data rate in bps
	Downlink int32 `json:"downlink,omitempty"`

	// data rate unit for uplink and downlink
	BitrateUnit string `json:"bitrate-unit,omitempty"`

	// QCI/QFI for the traffic
	TrafficClass string `json:"traffic-class,omitempty"`
}

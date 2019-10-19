// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EmbeddedSIMDeviceState Describes the embedded SIM activation code deployment state in relation to a device.
type EmbeddedSIMDeviceState struct {
	Entity
	// CreatedDateTime The time the embedded SIM device status was created. Generated service side.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedDateTime The time the embedded SIM device status was last modified. Updated service side.
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// LastSyncDateTime The time the embedded SIM device last checked in. Updated service side.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// UniversalIntegratedCircuitCardIdentifier The Universal Integrated Circuit Card Identifier (UICCID) identifying the hardware onto which a profile is to be deployed.
	UniversalIntegratedCircuitCardIdentifier *string `json:"universalIntegratedCircuitCardIdentifier,omitempty"`
	// DeviceName Device name to which the subscription was provisioned e.g. DESKTOP-JOE
	DeviceName *string `json:"deviceName,omitempty"`
	// UserName Username which the subscription was provisioned to e.g. joe@contoso.com
	UserName *string `json:"userName,omitempty"`
	// State The state of the profile operation applied to the device.
	State *EmbeddedSIMDeviceStateValue `json:"state,omitempty"`
	// StateDetails String description of the provisioning state.
	StateDetails *string `json:"stateDetails,omitempty"`
}

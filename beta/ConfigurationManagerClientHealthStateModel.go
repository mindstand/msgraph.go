// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ConfigurationManagerClientHealthState undocumented
type ConfigurationManagerClientHealthState struct {
	// State Current configuration manager client state.
	State *ConfigurationManagerClientState `json:"state,omitempty"`
	// ErrorCode Error code for failed state.
	ErrorCode *int `json:"errorCode,omitempty"`
	// LastSyncDateTime Datetime fo last sync with configuration manager management point.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
}

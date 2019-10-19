// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PrivilegedRoleSettings undocumented
type PrivilegedRoleSettings struct {
	Entity
	// ApproverIDs undocumented
	ApproverIDs []string `json:"approverIds,omitempty"`
	// MinElevationDuration undocumented
	MinElevationDuration *time.Duration `json:"minElevationDuration,omitempty"`
	// MaxElavationDuration undocumented
	MaxElavationDuration *time.Duration `json:"maxElavationDuration,omitempty"`
	// ElevationDuration undocumented
	ElevationDuration *time.Duration `json:"elevationDuration,omitempty"`
	// NotificationToUserOnElevation undocumented
	NotificationToUserOnElevation *bool `json:"notificationToUserOnElevation,omitempty"`
	// TicketingInfoOnElevation undocumented
	TicketingInfoOnElevation *bool `json:"ticketingInfoOnElevation,omitempty"`
	// MfaOnElevation undocumented
	MfaOnElevation *bool `json:"mfaOnElevation,omitempty"`
	// LastGlobalAdmin undocumented
	LastGlobalAdmin *bool `json:"lastGlobalAdmin,omitempty"`
	// IsMfaOnElevationConfigurable undocumented
	IsMfaOnElevationConfigurable *bool `json:"isMfaOnElevationConfigurable,omitempty"`
	// ApprovalOnElevation undocumented
	ApprovalOnElevation *bool `json:"approvalOnElevation,omitempty"`
}

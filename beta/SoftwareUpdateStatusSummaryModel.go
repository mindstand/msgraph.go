// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SoftwareUpdateStatusSummary undocumented
type SoftwareUpdateStatusSummary struct {
	Entity
	// DisplayName The name of the policy.
	DisplayName *string `json:"displayName,omitempty"`
	// CompliantDeviceCount Number of compliant devices.
	CompliantDeviceCount *int `json:"compliantDeviceCount,omitempty"`
	// NonCompliantDeviceCount Number of non compliant devices.
	NonCompliantDeviceCount *int `json:"nonCompliantDeviceCount,omitempty"`
	// RemediatedDeviceCount Number of remediated devices.
	RemediatedDeviceCount *int `json:"remediatedDeviceCount,omitempty"`
	// ErrorDeviceCount Number of devices had error.
	ErrorDeviceCount *int `json:"errorDeviceCount,omitempty"`
	// UnknownDeviceCount Number of unknown devices.
	UnknownDeviceCount *int `json:"unknownDeviceCount,omitempty"`
	// ConflictDeviceCount Number of conflict devices.
	ConflictDeviceCount *int `json:"conflictDeviceCount,omitempty"`
	// NotApplicableDeviceCount Number of not applicable devices.
	NotApplicableDeviceCount *int `json:"notApplicableDeviceCount,omitempty"`
	// CompliantUserCount Number of compliant users.
	CompliantUserCount *int `json:"compliantUserCount,omitempty"`
	// NonCompliantUserCount Number of non compliant users.
	NonCompliantUserCount *int `json:"nonCompliantUserCount,omitempty"`
	// RemediatedUserCount Number of remediated users.
	RemediatedUserCount *int `json:"remediatedUserCount,omitempty"`
	// ErrorUserCount Number of users had error.
	ErrorUserCount *int `json:"errorUserCount,omitempty"`
	// UnknownUserCount Number of unknown users.
	UnknownUserCount *int `json:"unknownUserCount,omitempty"`
	// ConflictUserCount Number of conflict users.
	ConflictUserCount *int `json:"conflictUserCount,omitempty"`
	// NotApplicableUserCount Number of not applicable users.
	NotApplicableUserCount *int `json:"notApplicableUserCount,omitempty"`
}

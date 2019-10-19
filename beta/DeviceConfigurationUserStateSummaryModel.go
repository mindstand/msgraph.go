// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationUserStateSummary undocumented
type DeviceConfigurationUserStateSummary struct {
	Entity
	// UnknownUserCount Number of unknown users
	UnknownUserCount *int `json:"unknownUserCount,omitempty"`
	// NotApplicableUserCount Number of not applicable users
	NotApplicableUserCount *int `json:"notApplicableUserCount,omitempty"`
	// CompliantUserCount Number of compliant users
	CompliantUserCount *int `json:"compliantUserCount,omitempty"`
	// RemediatedUserCount Number of remediated users
	RemediatedUserCount *int `json:"remediatedUserCount,omitempty"`
	// NonCompliantUserCount Number of NonCompliant users
	NonCompliantUserCount *int `json:"nonCompliantUserCount,omitempty"`
	// ErrorUserCount Number of error users
	ErrorUserCount *int `json:"errorUserCount,omitempty"`
	// ConflictUserCount Number of conflict users
	ConflictUserCount *int `json:"conflictUserCount,omitempty"`
}

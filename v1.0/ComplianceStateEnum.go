// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ComplianceState undocumented
type ComplianceState string

const (
	// ComplianceStateVUnknown undocumented
	ComplianceStateVUnknown ComplianceState = "Unknown"
	// ComplianceStateVCompliant undocumented
	ComplianceStateVCompliant ComplianceState = "Compliant"
	// ComplianceStateVNoncompliant undocumented
	ComplianceStateVNoncompliant ComplianceState = "Noncompliant"
	// ComplianceStateVConflict undocumented
	ComplianceStateVConflict ComplianceState = "Conflict"
	// ComplianceStateVError undocumented
	ComplianceStateVError ComplianceState = "Error"
	// ComplianceStateVInGracePeriod undocumented
	ComplianceStateVInGracePeriod ComplianceState = "InGracePeriod"
	// ComplianceStateVConfigManager undocumented
	ComplianceStateVConfigManager ComplianceState = "ConfigManager"
)

// ComplianceStatePUnknown returns a pointer to ComplianceStateVUnknown
func ComplianceStatePUnknown() *ComplianceState {
	v := ComplianceStateVUnknown
	return &v
}

// ComplianceStatePCompliant returns a pointer to ComplianceStateVCompliant
func ComplianceStatePCompliant() *ComplianceState {
	v := ComplianceStateVCompliant
	return &v
}

// ComplianceStatePNoncompliant returns a pointer to ComplianceStateVNoncompliant
func ComplianceStatePNoncompliant() *ComplianceState {
	v := ComplianceStateVNoncompliant
	return &v
}

// ComplianceStatePConflict returns a pointer to ComplianceStateVConflict
func ComplianceStatePConflict() *ComplianceState {
	v := ComplianceStateVConflict
	return &v
}

// ComplianceStatePError returns a pointer to ComplianceStateVError
func ComplianceStatePError() *ComplianceState {
	v := ComplianceStateVError
	return &v
}

// ComplianceStatePInGracePeriod returns a pointer to ComplianceStateVInGracePeriod
func ComplianceStatePInGracePeriod() *ComplianceState {
	v := ComplianceStateVInGracePeriod
	return &v
}

// ComplianceStatePConfigManager returns a pointer to ComplianceStateVConfigManager
func ComplianceStatePConfigManager() *ComplianceState {
	v := ComplianceStateVConfigManager
	return &v
}

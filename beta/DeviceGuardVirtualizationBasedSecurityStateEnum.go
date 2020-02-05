// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceGuardVirtualizationBasedSecurityState undocumented
type DeviceGuardVirtualizationBasedSecurityState string

const (
	// DeviceGuardVirtualizationBasedSecurityStateVRunning undocumented
	DeviceGuardVirtualizationBasedSecurityStateVRunning DeviceGuardVirtualizationBasedSecurityState = "Running"
	// DeviceGuardVirtualizationBasedSecurityStateVRebootRequired undocumented
	DeviceGuardVirtualizationBasedSecurityStateVRebootRequired DeviceGuardVirtualizationBasedSecurityState = "RebootRequired"
	// DeviceGuardVirtualizationBasedSecurityStateVRequire64BitArchitecture undocumented
	DeviceGuardVirtualizationBasedSecurityStateVRequire64BitArchitecture DeviceGuardVirtualizationBasedSecurityState = "Require64BitArchitecture"
	// DeviceGuardVirtualizationBasedSecurityStateVNotLicensed undocumented
	DeviceGuardVirtualizationBasedSecurityStateVNotLicensed DeviceGuardVirtualizationBasedSecurityState = "NotLicensed"
	// DeviceGuardVirtualizationBasedSecurityStateVNotConfigured undocumented
	DeviceGuardVirtualizationBasedSecurityStateVNotConfigured DeviceGuardVirtualizationBasedSecurityState = "NotConfigured"
	// DeviceGuardVirtualizationBasedSecurityStateVDoesNotMeetHardwareRequirements undocumented
	DeviceGuardVirtualizationBasedSecurityStateVDoesNotMeetHardwareRequirements DeviceGuardVirtualizationBasedSecurityState = "DoesNotMeetHardwareRequirements"
	// DeviceGuardVirtualizationBasedSecurityStateVOther undocumented
	DeviceGuardVirtualizationBasedSecurityStateVOther DeviceGuardVirtualizationBasedSecurityState = "Other"
)

// DeviceGuardVirtualizationBasedSecurityStatePRunning returns a pointer to DeviceGuardVirtualizationBasedSecurityStateVRunning
func DeviceGuardVirtualizationBasedSecurityStatePRunning() *DeviceGuardVirtualizationBasedSecurityState {
	v := DeviceGuardVirtualizationBasedSecurityStateVRunning
	return &v
}

// DeviceGuardVirtualizationBasedSecurityStatePRebootRequired returns a pointer to DeviceGuardVirtualizationBasedSecurityStateVRebootRequired
func DeviceGuardVirtualizationBasedSecurityStatePRebootRequired() *DeviceGuardVirtualizationBasedSecurityState {
	v := DeviceGuardVirtualizationBasedSecurityStateVRebootRequired
	return &v
}

// DeviceGuardVirtualizationBasedSecurityStatePRequire64BitArchitecture returns a pointer to DeviceGuardVirtualizationBasedSecurityStateVRequire64BitArchitecture
func DeviceGuardVirtualizationBasedSecurityStatePRequire64BitArchitecture() *DeviceGuardVirtualizationBasedSecurityState {
	v := DeviceGuardVirtualizationBasedSecurityStateVRequire64BitArchitecture
	return &v
}

// DeviceGuardVirtualizationBasedSecurityStatePNotLicensed returns a pointer to DeviceGuardVirtualizationBasedSecurityStateVNotLicensed
func DeviceGuardVirtualizationBasedSecurityStatePNotLicensed() *DeviceGuardVirtualizationBasedSecurityState {
	v := DeviceGuardVirtualizationBasedSecurityStateVNotLicensed
	return &v
}

// DeviceGuardVirtualizationBasedSecurityStatePNotConfigured returns a pointer to DeviceGuardVirtualizationBasedSecurityStateVNotConfigured
func DeviceGuardVirtualizationBasedSecurityStatePNotConfigured() *DeviceGuardVirtualizationBasedSecurityState {
	v := DeviceGuardVirtualizationBasedSecurityStateVNotConfigured
	return &v
}

// DeviceGuardVirtualizationBasedSecurityStatePDoesNotMeetHardwareRequirements returns a pointer to DeviceGuardVirtualizationBasedSecurityStateVDoesNotMeetHardwareRequirements
func DeviceGuardVirtualizationBasedSecurityStatePDoesNotMeetHardwareRequirements() *DeviceGuardVirtualizationBasedSecurityState {
	v := DeviceGuardVirtualizationBasedSecurityStateVDoesNotMeetHardwareRequirements
	return &v
}

// DeviceGuardVirtualizationBasedSecurityStatePOther returns a pointer to DeviceGuardVirtualizationBasedSecurityStateVOther
func DeviceGuardVirtualizationBasedSecurityStatePOther() *DeviceGuardVirtualizationBasedSecurityState {
	v := DeviceGuardVirtualizationBasedSecurityStateVOther
	return &v
}

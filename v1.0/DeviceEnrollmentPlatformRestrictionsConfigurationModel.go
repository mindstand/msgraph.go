// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceEnrollmentPlatformRestrictionsConfiguration undocumented
type DeviceEnrollmentPlatformRestrictionsConfiguration struct {
	DeviceEnrollmentConfiguration
	// IOSRestriction undocumented
	IOSRestriction *DeviceEnrollmentPlatformRestriction `json:"iosRestriction,omitempty"`
	// WindowsRestriction undocumented
	WindowsRestriction *DeviceEnrollmentPlatformRestriction `json:"windowsRestriction,omitempty"`
	// WindowsMobileRestriction undocumented
	WindowsMobileRestriction *DeviceEnrollmentPlatformRestriction `json:"windowsMobileRestriction,omitempty"`
	// AndroidRestriction undocumented
	AndroidRestriction *DeviceEnrollmentPlatformRestriction `json:"androidRestriction,omitempty"`
	// MacOSRestriction undocumented
	MacOSRestriction *DeviceEnrollmentPlatformRestriction `json:"macOSRestriction,omitempty"`
}

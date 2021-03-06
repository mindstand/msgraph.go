// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceEnrollmentWindowsHelloForBusinessConfiguration undocumented
type DeviceEnrollmentWindowsHelloForBusinessConfiguration struct {
	// DeviceEnrollmentConfiguration is the base model of DeviceEnrollmentWindowsHelloForBusinessConfiguration
	DeviceEnrollmentConfiguration
	// PinMinimumLength undocumented
	PinMinimumLength *int `json:"pinMinimumLength,omitempty"`
	// PinMaximumLength undocumented
	PinMaximumLength *int `json:"pinMaximumLength,omitempty"`
	// PinUppercaseCharactersUsage undocumented
	PinUppercaseCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinUppercaseCharactersUsage,omitempty"`
	// PinLowercaseCharactersUsage undocumented
	PinLowercaseCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinLowercaseCharactersUsage,omitempty"`
	// PinSpecialCharactersUsage undocumented
	PinSpecialCharactersUsage *WindowsHelloForBusinessPinUsage `json:"pinSpecialCharactersUsage,omitempty"`
	// State undocumented
	State *Enablement `json:"state,omitempty"`
	// SecurityDeviceRequired undocumented
	SecurityDeviceRequired *bool `json:"securityDeviceRequired,omitempty"`
	// UnlockWithBiometricsEnabled undocumented
	UnlockWithBiometricsEnabled *bool `json:"unlockWithBiometricsEnabled,omitempty"`
	// RemotePassportEnabled undocumented
	RemotePassportEnabled *bool `json:"remotePassportEnabled,omitempty"`
	// PinPreviousBlockCount undocumented
	PinPreviousBlockCount *int `json:"pinPreviousBlockCount,omitempty"`
	// PinExpirationInDays undocumented
	PinExpirationInDays *int `json:"pinExpirationInDays,omitempty"`
	// EnhancedBiometricsState undocumented
	EnhancedBiometricsState *Enablement `json:"enhancedBiometricsState,omitempty"`
}

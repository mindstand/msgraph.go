// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsDefenderAdvancedThreatProtectionConfiguration Windows Defender AdvancedThreatProtection Configuration.
type WindowsDefenderAdvancedThreatProtectionConfiguration struct {
	DeviceConfiguration
	// AdvancedThreatProtectionOnboardingBlob Windows Defender AdvancedThreatProtection Onboarding Blob.
	AdvancedThreatProtectionOnboardingBlob *string `json:"advancedThreatProtectionOnboardingBlob,omitempty"`
	// AdvancedThreatProtectionOnboardingFilename Name of the file from which AdvancedThreatProtectionOnboardingBlob was obtained.
	AdvancedThreatProtectionOnboardingFilename *string `json:"advancedThreatProtectionOnboardingFilename,omitempty"`
	// AdvancedThreatProtectionAutoPopulateOnboardingBlob Auto populate onboarding blob programmatically from Advanced Threat protection service
	AdvancedThreatProtectionAutoPopulateOnboardingBlob *bool `json:"advancedThreatProtectionAutoPopulateOnboardingBlob,omitempty"`
	// AllowSampleSharing Windows Defender AdvancedThreatProtection "Allow Sample Sharing" Rule
	AllowSampleSharing *bool `json:"allowSampleSharing,omitempty"`
	// EnableExpeditedTelemetryReporting Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
	EnableExpeditedTelemetryReporting *bool `json:"enableExpeditedTelemetryReporting,omitempty"`
	// AdvancedThreatProtectionOffboardingBlob Windows Defender AdvancedThreatProtection Offboarding Blob.
	AdvancedThreatProtectionOffboardingBlob *string `json:"advancedThreatProtectionOffboardingBlob,omitempty"`
	// AdvancedThreatProtectionOffboardingFilename Name of the file from which AdvancedThreatProtectionOffboardingBlob was obtained.
	AdvancedThreatProtectionOffboardingFilename *string `json:"advancedThreatProtectionOffboardingFilename,omitempty"`
}

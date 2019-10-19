// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OutOfBoxExperienceSettings undocumented
type OutOfBoxExperienceSettings struct {
	// HidePrivacySettings Show or hide privacy settings to user
	HidePrivacySettings *bool `json:"hidePrivacySettings,omitempty"`
	// HideEULA Show or hide EULA to user
	HideEULA *bool `json:"hideEULA,omitempty"`
	// UserType Type of user
	UserType *WindowsUserType `json:"userType,omitempty"`
	// DeviceUsageType AAD join authentication type
	DeviceUsageType *WindowsDeviceUsageType `json:"deviceUsageType,omitempty"`
	// SkipKeyboardSelectionPage If set, then skip the keyboard selection page if Language and Region are set
	SkipKeyboardSelectionPage *bool `json:"skipKeyboardSelectionPage,omitempty"`
	// HideEscapeLink If set to true, then the user can't start over with different account, on company sign-in
	HideEscapeLink *bool `json:"hideEscapeLink,omitempty"`
}

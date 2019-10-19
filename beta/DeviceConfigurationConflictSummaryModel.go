// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationConflictSummary Conflict summary for a set of device configuration policies.
type DeviceConfigurationConflictSummary struct {
	Entity
	// ConflictingDeviceConfigurations The set of policies in conflict with the given setting
	ConflictingDeviceConfigurations []SettingSource `json:"conflictingDeviceConfigurations,omitempty"`
	// ContributingSettings The set of settings in conflict with the given policies
	ContributingSettings []string `json:"contributingSettings,omitempty"`
	// DeviceCheckinsImpacted The count of checkins impacted by the conflicting policies and settings
	DeviceCheckinsImpacted *int `json:"deviceCheckinsImpacted,omitempty"`
}

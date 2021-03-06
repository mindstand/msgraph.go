// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SecurityBaselineTemplate The security baseline template of the account
type SecurityBaselineTemplate struct {
	// DeviceManagementTemplate is the base model of SecurityBaselineTemplate
	DeviceManagementTemplate
	// DeviceStateSummary undocumented
	DeviceStateSummary *SecurityBaselineStateSummary `json:"deviceStateSummary,omitempty"`
	// DeviceStates undocumented
	DeviceStates []SecurityBaselineDeviceState `json:"deviceStates,omitempty"`
	// CategoryDeviceStateSummaries undocumented
	CategoryDeviceStateSummaries []SecurityBaselineCategoryStateSummary `json:"categoryDeviceStateSummaries,omitempty"`
}

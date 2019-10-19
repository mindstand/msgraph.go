// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceComplianceScheduledActionForRule Scheduled Action for Rule
type DeviceComplianceScheduledActionForRule struct {
	Entity
	// RuleName Name of the rule which this scheduled action applies to.
	RuleName *string `json:"ruleName,omitempty"`
	// ScheduledActionConfigurations undocumented
	ScheduledActionConfigurations []DeviceComplianceActionItem `json:"scheduledActionConfigurations,omitempty"`
}

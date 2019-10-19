// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// IOSUpdateConfiguration IOS Update Configuration, allows you to configure time window within week to install iOS updates
type IOSUpdateConfiguration struct {
	DeviceConfiguration
	// ActiveHoursStart Active Hours Start (active hours mean the time window when updates install should not happen)
	ActiveHoursStart *time.Time `json:"activeHoursStart,omitempty"`
	// ActiveHoursEnd Active Hours End (active hours mean the time window when updates install should not happen)
	ActiveHoursEnd *time.Time `json:"activeHoursEnd,omitempty"`
	// ScheduledInstallDays Days in week for which active hours are configured. This collection can contain a maximum of 7 elements.
	ScheduledInstallDays []DayOfWeek `json:"scheduledInstallDays,omitempty"`
	// UtcTimeOffsetInMinutes UTC Time Offset indicated in minutes
	UtcTimeOffsetInMinutes *int `json:"utcTimeOffsetInMinutes,omitempty"`
}

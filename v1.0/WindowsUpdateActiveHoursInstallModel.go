// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// WindowsUpdateActiveHoursInstall undocumented
type WindowsUpdateActiveHoursInstall struct {
	WindowsUpdateInstallScheduleType
	// ActiveHoursStart Active Hours Start
	ActiveHoursStart *time.Time `json:"activeHoursStart,omitempty"`
	// ActiveHoursEnd Active Hours End
	ActiveHoursEnd *time.Time `json:"activeHoursEnd,omitempty"`
}

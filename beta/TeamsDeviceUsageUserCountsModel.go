// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// TeamsDeviceUsageUserCounts undocumented
type TeamsDeviceUsageUserCounts struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// Web undocumented
	Web *int `json:"web,omitempty"`
	// WindowsPhone undocumented
	WindowsPhone *int `json:"windowsPhone,omitempty"`
	// AndroidPhone undocumented
	AndroidPhone *int `json:"androidPhone,omitempty"`
	// IOS undocumented
	IOS *int `json:"ios,omitempty"`
	// Mac undocumented
	Mac *int `json:"mac,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}
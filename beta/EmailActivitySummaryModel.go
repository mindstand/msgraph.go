// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EmailActivitySummary undocumented
type EmailActivitySummary struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// Send undocumented
	Send *int `json:"send,omitempty"`
	// Receive undocumented
	Receive *int `json:"receive,omitempty"`
	// Read undocumented
	Read *int `json:"read,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

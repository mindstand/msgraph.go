// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// TeamsUserActivityUserCounts undocumented
type TeamsUserActivityUserCounts struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// TeamChatMessages undocumented
	TeamChatMessages *int `json:"teamChatMessages,omitempty"`
	// PrivateChatMessages undocumented
	PrivateChatMessages *int `json:"privateChatMessages,omitempty"`
	// Calls undocumented
	Calls *int `json:"calls,omitempty"`
	// Meetings undocumented
	Meetings *int `json:"meetings,omitempty"`
	// OtherActions undocumented
	OtherActions *int `json:"otherActions,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

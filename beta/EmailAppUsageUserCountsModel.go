// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EmailAppUsageUserCounts undocumented
type EmailAppUsageUserCounts struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// MailForMac undocumented
	MailForMac *int `json:"mailForMac,omitempty"`
	// OutlookForMac undocumented
	OutlookForMac *int `json:"outlookForMac,omitempty"`
	// OutlookForWindows undocumented
	OutlookForWindows *int `json:"outlookForWindows,omitempty"`
	// OutlookForMobile undocumented
	OutlookForMobile *int `json:"outlookForMobile,omitempty"`
	// OtherForMobile undocumented
	OtherForMobile *int `json:"otherForMobile,omitempty"`
	// OutlookForWeb undocumented
	OutlookForWeb *int `json:"outlookForWeb,omitempty"`
	// Pop3App undocumented
	Pop3App *int `json:"pop3App,omitempty"`
	// Imap4App undocumented
	Imap4App *int `json:"imap4App,omitempty"`
	// SMTPApp undocumented
	SMTPApp *int `json:"smtpApp,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

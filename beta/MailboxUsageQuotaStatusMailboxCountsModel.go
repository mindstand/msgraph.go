// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// MailboxUsageQuotaStatusMailboxCounts undocumented
type MailboxUsageQuotaStatusMailboxCounts struct {
	// Entity is the base model of MailboxUsageQuotaStatusMailboxCounts
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// UnderLimit undocumented
	UnderLimit *int `json:"underLimit,omitempty"`
	// WarningIssued undocumented
	WarningIssued *int `json:"warningIssued,omitempty"`
	// SendProhibited undocumented
	SendProhibited *int `json:"sendProhibited,omitempty"`
	// SendReceiveProhibited undocumented
	SendReceiveProhibited *int `json:"sendReceiveProhibited,omitempty"`
	// Indeterminate undocumented
	Indeterminate *int `json:"indeterminate,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

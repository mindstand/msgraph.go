// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// MailboxUsageDetail undocumented
type MailboxUsageDetail struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// DeletedDate undocumented
	DeletedDate *time.Time `json:"deletedDate,omitempty"`
	// CreatedDate undocumented
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *time.Time `json:"lastActivityDate,omitempty"`
	// ItemCount undocumented
	ItemCount *int `json:"itemCount,omitempty"`
	// StorageUsedInBytes undocumented
	StorageUsedInBytes *int `json:"storageUsedInBytes,omitempty"`
	// DeletedItemCount undocumented
	DeletedItemCount *int `json:"deletedItemCount,omitempty"`
	// DeletedItemSizeInBytes undocumented
	DeletedItemSizeInBytes *int `json:"deletedItemSizeInBytes,omitempty"`
	// IssueWarningQuotaInBytes undocumented
	IssueWarningQuotaInBytes *int `json:"issueWarningQuotaInBytes,omitempty"`
	// ProhibitSendQuotaInBytes undocumented
	ProhibitSendQuotaInBytes *int `json:"prohibitSendQuotaInBytes,omitempty"`
	// ProhibitSendReceiveQuotaInBytes undocumented
	ProhibitSendReceiveQuotaInBytes *int `json:"prohibitSendReceiveQuotaInBytes,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

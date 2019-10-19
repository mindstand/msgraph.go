// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// OneDriveUsageAccountDetail undocumented
type OneDriveUsageAccountDetail struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// SiteURL undocumented
	SiteURL *string `json:"siteUrl,omitempty"`
	// OwnerDisplayName undocumented
	OwnerDisplayName *string `json:"ownerDisplayName,omitempty"`
	// OwnerPrincipalName undocumented
	OwnerPrincipalName *string `json:"ownerPrincipalName,omitempty"`
	// IsDeleted undocumented
	IsDeleted *bool `json:"isDeleted,omitempty"`
	// LastActivityDate undocumented
	LastActivityDate *time.Time `json:"lastActivityDate,omitempty"`
	// FileCount undocumented
	FileCount *int `json:"fileCount,omitempty"`
	// ActiveFileCount undocumented
	ActiveFileCount *int `json:"activeFileCount,omitempty"`
	// StorageUsedInBytes undocumented
	StorageUsedInBytes *int `json:"storageUsedInBytes,omitempty"`
	// StorageAllocatedInBytes undocumented
	StorageAllocatedInBytes *int `json:"storageAllocatedInBytes,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

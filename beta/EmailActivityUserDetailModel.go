// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// EmailActivityUserDetail undocumented
type EmailActivityUserDetail struct {
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
	// LastActivityDate undocumented
	LastActivityDate *time.Time `json:"lastActivityDate,omitempty"`
	// SendCount undocumented
	SendCount *int `json:"sendCount,omitempty"`
	// ReceiveCount undocumented
	ReceiveCount *int `json:"receiveCount,omitempty"`
	// ReadCount undocumented
	ReadCount *int `json:"readCount,omitempty"`
	// AssignedProducts undocumented
	AssignedProducts []string `json:"assignedProducts,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

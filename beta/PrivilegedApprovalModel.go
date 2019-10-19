// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PrivilegedApproval undocumented
type PrivilegedApproval struct {
	Entity
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// RoleID undocumented
	RoleID *string `json:"roleId,omitempty"`
	// ApprovalType undocumented
	ApprovalType *string `json:"approvalType,omitempty"`
	// ApprovalState undocumented
	ApprovalState *ApprovalState `json:"approvalState,omitempty"`
	// ApprovalDuration undocumented
	ApprovalDuration *time.Duration `json:"approvalDuration,omitempty"`
	// RequestorReason undocumented
	RequestorReason *string `json:"requestorReason,omitempty"`
	// ApproverReason undocumented
	ApproverReason *string `json:"approverReason,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// RequestNavigation undocumented
	RequestNavigation *PrivilegedRoleAssignmentRequestObject `json:"request,omitempty"`
	// RoleInfo undocumented
	RoleInfo *PrivilegedRole `json:"roleInfo,omitempty"`
}

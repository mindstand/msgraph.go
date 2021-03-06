// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// WindowsDefenderApplicationControlSupplementalPolicy undocumented
type WindowsDefenderApplicationControlSupplementalPolicy struct {
	// Entity is the base model of WindowsDefenderApplicationControlSupplementalPolicy
	Entity
	// DisplayName The display name of WindowsDefenderApplicationControl supplemental policy.
	DisplayName *string `json:"displayName,omitempty"`
	// Description The description of WindowsDefenderApplicationControl supplemental policy.
	Description *string `json:"description,omitempty"`
	// Content The WindowsDefenderApplicationControl supplemental policy content in byte array format.
	Content *Binary `json:"content,omitempty"`
	// ContentFileName The WindowsDefenderApplicationControl supplemental policy content's file name.
	ContentFileName *string `json:"contentFileName,omitempty"`
	// Version The WindowsDefenderApplicationControl supplemental policy's version.
	Version *string `json:"version,omitempty"`
	// CreationDateTime The date and time when the WindowsDefenderApplicationControl supplemental policy was uploaded.
	CreationDateTime *time.Time `json:"creationDateTime,omitempty"`
	// LastModifiedDateTime The date and time when the WindowsDefenderApplicationControl supplemental policy was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// RoleScopeTagIDs List of Scope Tags for this WindowsDefenderApplicationControl supplemental policy entity.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// Assignments undocumented
	Assignments []WindowsDefenderApplicationControlSupplementalPolicyAssignment `json:"assignments,omitempty"`
	// DeploySummary undocumented
	DeploySummary *WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary `json:"deploySummary,omitempty"`
	// DeviceStatuses undocumented
	DeviceStatuses []WindowsDefenderApplicationControlSupplementalPolicyDeploymentStatus `json:"deviceStatuses,omitempty"`
}

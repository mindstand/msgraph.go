// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AccessPackageResource undocumented
type AccessPackageResource struct {
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
	// ResourceType undocumented
	ResourceType *string `json:"resourceType,omitempty"`
	// OriginID undocumented
	OriginID *string `json:"originId,omitempty"`
	// OriginSystem undocumented
	OriginSystem *string `json:"originSystem,omitempty"`
	// IsPendingOnboarding undocumented
	IsPendingOnboarding *bool `json:"isPendingOnboarding,omitempty"`
	// AddedBy undocumented
	AddedBy *string `json:"addedBy,omitempty"`
	// AddedOn undocumented
	AddedOn *time.Time `json:"addedOn,omitempty"`
	// AccessPackageResourceScopes undocumented
	AccessPackageResourceScopes []AccessPackageResourceScope `json:"accessPackageResourceScopes,omitempty"`
	// AccessPackageResourceRoles undocumented
	AccessPackageResourceRoles []AccessPackageResourceRole `json:"accessPackageResourceRoles,omitempty"`
}

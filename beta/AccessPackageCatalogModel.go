// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AccessPackageCatalog undocumented
type AccessPackageCatalog struct {
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// CatalogType undocumented
	CatalogType *string `json:"catalogType,omitempty"`
	// CatalogStatus undocumented
	CatalogStatus *string `json:"catalogStatus,omitempty"`
	// IsExternallyVisible undocumented
	IsExternallyVisible *bool `json:"isExternallyVisible,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// AccessPackageResources undocumented
	AccessPackageResources []AccessPackageResource `json:"accessPackageResources,omitempty"`
	// AccessPackageResourceRoles undocumented
	AccessPackageResourceRoles []AccessPackageResourceRole `json:"accessPackageResourceRoles,omitempty"`
	// AccessPackageResourceScopes undocumented
	AccessPackageResourceScopes []AccessPackageResourceScope `json:"accessPackageResourceScopes,omitempty"`
	// AccessPackages undocumented
	AccessPackages []AccessPackage `json:"accessPackages,omitempty"`
}
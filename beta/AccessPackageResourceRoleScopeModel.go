// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AccessPackageResourceRoleScope undocumented
type AccessPackageResourceRoleScope struct {
	Entity
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// AccessPackageResourceRole undocumented
	AccessPackageResourceRole *AccessPackageResourceRole `json:"accessPackageResourceRole,omitempty"`
	// AccessPackageResourceScope undocumented
	AccessPackageResourceScope *AccessPackageResourceScope `json:"accessPackageResourceScope,omitempty"`
}

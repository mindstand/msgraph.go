// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AccessPackageAssignmentResourceRole undocumented
type AccessPackageAssignmentResourceRole struct {
	Entity
	// OriginID undocumented
	OriginID *string `json:"originId,omitempty"`
	// OriginSystem undocumented
	OriginSystem *string `json:"originSystem,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// AccessPackageResourceScope undocumented
	AccessPackageResourceScope *AccessPackageResourceScope `json:"accessPackageResourceScope,omitempty"`
	// AccessPackageResourceRole undocumented
	AccessPackageResourceRole *AccessPackageResourceRole `json:"accessPackageResourceRole,omitempty"`
	// AccessPackageSubject undocumented
	AccessPackageSubject *AccessPackageSubject `json:"accessPackageSubject,omitempty"`
	// AccessPackageAssignments undocumented
	AccessPackageAssignments []AccessPackageAssignment `json:"accessPackageAssignments,omitempty"`
}

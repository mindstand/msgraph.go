// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ScopedRoleMembership undocumented
type ScopedRoleMembership struct {
	Entity
	// RoleID undocumented
	RoleID *string `json:"roleId,omitempty"`
	// AdministrativeUnitID undocumented
	AdministrativeUnitID *string `json:"administrativeUnitId,omitempty"`
	// RoleMemberInfo undocumented
	RoleMemberInfo *Identity `json:"roleMemberInfo,omitempty"`
}

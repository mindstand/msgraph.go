// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceAndAppManagementAssignedRoleDetails undocumented
type DeviceAndAppManagementAssignedRoleDetails struct {
	// RoleDefinitionIDs Role Definition IDs for the specifc Role Definitions assigned to a user.
	RoleDefinitionIDs []string `json:"roleDefinitionIds,omitempty"`
	// RoleAssignmentIDs Role Assignment IDs for the specifc Role Assignments assigned to a user.
	RoleAssignmentIDs []string `json:"roleAssignmentIds,omitempty"`
}
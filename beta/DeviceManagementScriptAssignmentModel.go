// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementScriptAssignment Contains properties used to assign a device management script to a group.
type DeviceManagementScriptAssignment struct {
	Entity
	// Target The Id of the Azure Active Directory group we are targeting the script to.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}

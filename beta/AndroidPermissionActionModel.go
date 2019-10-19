// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidPermissionAction undocumented
type AndroidPermissionAction struct {
	// Permission Android permission string, defined in the official Android documentation.  Example 'android.permission.READ_CONTACTS'.
	Permission *string `json:"permission,omitempty"`
	// Action Type of Android permission action.
	Action *AndroidPermissionActionType `json:"action,omitempty"`
}

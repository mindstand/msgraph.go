// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsKioskLocalGroup undocumented
type WindowsKioskLocalGroup struct {
	WindowsKioskUser
	// GroupName The name of the local group that will be locked to this kiosk configuration
	GroupName *string `json:"groupName,omitempty"`
}

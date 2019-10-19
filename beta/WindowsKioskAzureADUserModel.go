// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsKioskAzureADUser undocumented
type WindowsKioskAzureADUser struct {
	WindowsKioskUser
	// UserID The ID of the AzureAD user that will be locked to this kiosk configuration
	UserID *string `json:"userId,omitempty"`
	// UserPrincipalName The user accounts that will be locked to this kiosk configuration
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

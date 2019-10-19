// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementDerivedCredentialSettings Entity that describes tenant level settings for derived credentials
type DeviceManagementDerivedCredentialSettings struct {
	Entity
	// HelpURL The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
	HelpURL *string `json:"helpUrl,omitempty"`
	// DisplayName The display name for the profile.
	DisplayName *string `json:"displayName,omitempty"`
	// Issuer The derived credential provider to use.
	Issuer *DeviceManagementDerivedCredentialIssuer `json:"issuer,omitempty"`
	// NotificationType The methods used to inform the end user to open Company Portal to deliver Wi-Fi, VPN, or email profiles that use certificates to the device.
	NotificationType *DeviceManagementDerivedCredentialNotificationType `json:"notificationType,omitempty"`
}

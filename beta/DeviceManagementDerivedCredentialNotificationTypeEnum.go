// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementDerivedCredentialNotificationType undocumented
type DeviceManagementDerivedCredentialNotificationType string

const (
	// DeviceManagementDerivedCredentialNotificationTypeVNone undocumented
	DeviceManagementDerivedCredentialNotificationTypeVNone DeviceManagementDerivedCredentialNotificationType = "None"
	// DeviceManagementDerivedCredentialNotificationTypeVCompanyPortal undocumented
	DeviceManagementDerivedCredentialNotificationTypeVCompanyPortal DeviceManagementDerivedCredentialNotificationType = "CompanyPortal"
	// DeviceManagementDerivedCredentialNotificationTypeVEmail undocumented
	DeviceManagementDerivedCredentialNotificationTypeVEmail DeviceManagementDerivedCredentialNotificationType = "Email"
)

// DeviceManagementDerivedCredentialNotificationTypePNone returns a pointer to DeviceManagementDerivedCredentialNotificationTypeVNone
func DeviceManagementDerivedCredentialNotificationTypePNone() *DeviceManagementDerivedCredentialNotificationType {
	v := DeviceManagementDerivedCredentialNotificationTypeVNone
	return &v
}

// DeviceManagementDerivedCredentialNotificationTypePCompanyPortal returns a pointer to DeviceManagementDerivedCredentialNotificationTypeVCompanyPortal
func DeviceManagementDerivedCredentialNotificationTypePCompanyPortal() *DeviceManagementDerivedCredentialNotificationType {
	v := DeviceManagementDerivedCredentialNotificationTypeVCompanyPortal
	return &v
}

// DeviceManagementDerivedCredentialNotificationTypePEmail returns a pointer to DeviceManagementDerivedCredentialNotificationTypeVEmail
func DeviceManagementDerivedCredentialNotificationTypePEmail() *DeviceManagementDerivedCredentialNotificationType {
	v := DeviceManagementDerivedCredentialNotificationTypeVEmail
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DefenderSecurityCenterNotificationsFromAppType undocumented
type DefenderSecurityCenterNotificationsFromAppType string

const (
	// DefenderSecurityCenterNotificationsFromAppTypeVNotConfigured undocumented
	DefenderSecurityCenterNotificationsFromAppTypeVNotConfigured DefenderSecurityCenterNotificationsFromAppType = "NotConfigured"
	// DefenderSecurityCenterNotificationsFromAppTypeVBlockNoncriticalNotifications undocumented
	DefenderSecurityCenterNotificationsFromAppTypeVBlockNoncriticalNotifications DefenderSecurityCenterNotificationsFromAppType = "BlockNoncriticalNotifications"
	// DefenderSecurityCenterNotificationsFromAppTypeVBlockAllNotifications undocumented
	DefenderSecurityCenterNotificationsFromAppTypeVBlockAllNotifications DefenderSecurityCenterNotificationsFromAppType = "BlockAllNotifications"
)

// DefenderSecurityCenterNotificationsFromAppTypePNotConfigured returns a pointer to DefenderSecurityCenterNotificationsFromAppTypeVNotConfigured
func DefenderSecurityCenterNotificationsFromAppTypePNotConfigured() *DefenderSecurityCenterNotificationsFromAppType {
	v := DefenderSecurityCenterNotificationsFromAppTypeVNotConfigured
	return &v
}

// DefenderSecurityCenterNotificationsFromAppTypePBlockNoncriticalNotifications returns a pointer to DefenderSecurityCenterNotificationsFromAppTypeVBlockNoncriticalNotifications
func DefenderSecurityCenterNotificationsFromAppTypePBlockNoncriticalNotifications() *DefenderSecurityCenterNotificationsFromAppType {
	v := DefenderSecurityCenterNotificationsFromAppTypeVBlockNoncriticalNotifications
	return &v
}

// DefenderSecurityCenterNotificationsFromAppTypePBlockAllNotifications returns a pointer to DefenderSecurityCenterNotificationsFromAppTypeVBlockAllNotifications
func DefenderSecurityCenterNotificationsFromAppTypePBlockAllNotifications() *DefenderSecurityCenterNotificationsFromAppType {
	v := DefenderSecurityCenterNotificationsFromAppTypeVBlockAllNotifications
	return &v
}

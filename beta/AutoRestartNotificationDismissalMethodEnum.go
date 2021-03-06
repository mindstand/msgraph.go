// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AutoRestartNotificationDismissalMethod undocumented
type AutoRestartNotificationDismissalMethod string

const (
	// AutoRestartNotificationDismissalMethodVNotConfigured undocumented
	AutoRestartNotificationDismissalMethodVNotConfigured AutoRestartNotificationDismissalMethod = "NotConfigured"
	// AutoRestartNotificationDismissalMethodVAutomatic undocumented
	AutoRestartNotificationDismissalMethodVAutomatic AutoRestartNotificationDismissalMethod = "Automatic"
	// AutoRestartNotificationDismissalMethodVUser undocumented
	AutoRestartNotificationDismissalMethodVUser AutoRestartNotificationDismissalMethod = "User"
)

// AutoRestartNotificationDismissalMethodPNotConfigured returns a pointer to AutoRestartNotificationDismissalMethodVNotConfigured
func AutoRestartNotificationDismissalMethodPNotConfigured() *AutoRestartNotificationDismissalMethod {
	v := AutoRestartNotificationDismissalMethodVNotConfigured
	return &v
}

// AutoRestartNotificationDismissalMethodPAutomatic returns a pointer to AutoRestartNotificationDismissalMethodVAutomatic
func AutoRestartNotificationDismissalMethodPAutomatic() *AutoRestartNotificationDismissalMethod {
	v := AutoRestartNotificationDismissalMethodVAutomatic
	return &v
}

// AutoRestartNotificationDismissalMethodPUser returns a pointer to AutoRestartNotificationDismissalMethodVUser
func AutoRestartNotificationDismissalMethodPUser() *AutoRestartNotificationDismissalMethod {
	v := AutoRestartNotificationDismissalMethodVUser
	return &v
}

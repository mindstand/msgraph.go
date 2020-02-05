// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppActionType undocumented
type MobileAppActionType string

const (
	// MobileAppActionTypeVUnknown undocumented
	MobileAppActionTypeVUnknown MobileAppActionType = "Unknown"
	// MobileAppActionTypeVInstallCommandSent undocumented
	MobileAppActionTypeVInstallCommandSent MobileAppActionType = "InstallCommandSent"
	// MobileAppActionTypeVInstalled undocumented
	MobileAppActionTypeVInstalled MobileAppActionType = "Installed"
	// MobileAppActionTypeVUninstalled undocumented
	MobileAppActionTypeVUninstalled MobileAppActionType = "Uninstalled"
	// MobileAppActionTypeVUserRequestedInstall undocumented
	MobileAppActionTypeVUserRequestedInstall MobileAppActionType = "UserRequestedInstall"
)

// MobileAppActionTypePUnknown returns a pointer to MobileAppActionTypeVUnknown
func MobileAppActionTypePUnknown() *MobileAppActionType {
	v := MobileAppActionTypeVUnknown
	return &v
}

// MobileAppActionTypePInstallCommandSent returns a pointer to MobileAppActionTypeVInstallCommandSent
func MobileAppActionTypePInstallCommandSent() *MobileAppActionType {
	v := MobileAppActionTypeVInstallCommandSent
	return &v
}

// MobileAppActionTypePInstalled returns a pointer to MobileAppActionTypeVInstalled
func MobileAppActionTypePInstalled() *MobileAppActionType {
	v := MobileAppActionTypeVInstalled
	return &v
}

// MobileAppActionTypePUninstalled returns a pointer to MobileAppActionTypeVUninstalled
func MobileAppActionTypePUninstalled() *MobileAppActionType {
	v := MobileAppActionTypeVUninstalled
	return &v
}

// MobileAppActionTypePUserRequestedInstall returns a pointer to MobileAppActionTypeVUserRequestedInstall
func MobileAppActionTypePUserRequestedInstall() *MobileAppActionType {
	v := MobileAppActionTypeVUserRequestedInstall
	return &v
}

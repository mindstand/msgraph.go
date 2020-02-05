// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AutomaticUpdateMode undocumented
type AutomaticUpdateMode string

const (
	// AutomaticUpdateModeVUserDefined undocumented
	AutomaticUpdateModeVUserDefined AutomaticUpdateMode = "UserDefined"
	// AutomaticUpdateModeVNotifyDownload undocumented
	AutomaticUpdateModeVNotifyDownload AutomaticUpdateMode = "NotifyDownload"
	// AutomaticUpdateModeVAutoInstallAtMaintenanceTime undocumented
	AutomaticUpdateModeVAutoInstallAtMaintenanceTime AutomaticUpdateMode = "AutoInstallAtMaintenanceTime"
	// AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime undocumented
	AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime AutomaticUpdateMode = "AutoInstallAndRebootAtMaintenanceTime"
	// AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime undocumented
	AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime AutomaticUpdateMode = "AutoInstallAndRebootAtScheduledTime"
	// AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl undocumented
	AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl AutomaticUpdateMode = "AutoInstallAndRebootWithoutEndUserControl"
)

// AutomaticUpdateModePUserDefined returns a pointer to AutomaticUpdateModeVUserDefined
func AutomaticUpdateModePUserDefined() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVUserDefined
	return &v
}

// AutomaticUpdateModePNotifyDownload returns a pointer to AutomaticUpdateModeVNotifyDownload
func AutomaticUpdateModePNotifyDownload() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVNotifyDownload
	return &v
}

// AutomaticUpdateModePAutoInstallAtMaintenanceTime returns a pointer to AutomaticUpdateModeVAutoInstallAtMaintenanceTime
func AutomaticUpdateModePAutoInstallAtMaintenanceTime() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAtMaintenanceTime
	return &v
}

// AutomaticUpdateModePAutoInstallAndRebootAtMaintenanceTime returns a pointer to AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime
func AutomaticUpdateModePAutoInstallAndRebootAtMaintenanceTime() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAndRebootAtMaintenanceTime
	return &v
}

// AutomaticUpdateModePAutoInstallAndRebootAtScheduledTime returns a pointer to AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime
func AutomaticUpdateModePAutoInstallAndRebootAtScheduledTime() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAndRebootAtScheduledTime
	return &v
}

// AutomaticUpdateModePAutoInstallAndRebootWithoutEndUserControl returns a pointer to AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl
func AutomaticUpdateModePAutoInstallAndRebootWithoutEndUserControl() *AutomaticUpdateMode {
	v := AutomaticUpdateModeVAutoInstallAndRebootWithoutEndUserControl
	return &v
}

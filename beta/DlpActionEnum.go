// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DlpAction undocumented
type DlpAction string

const (
	// DlpActionVNotifyUser undocumented
	DlpActionVNotifyUser DlpAction = "NotifyUser"
	// DlpActionVBlockAccess undocumented
	DlpActionVBlockAccess DlpAction = "BlockAccess"
	// DlpActionVDeviceRestriction undocumented
	DlpActionVDeviceRestriction DlpAction = "DeviceRestriction"
)

// DlpActionPNotifyUser returns a pointer to DlpActionVNotifyUser
func DlpActionPNotifyUser() *DlpAction {
	v := DlpActionVNotifyUser
	return &v
}

// DlpActionPBlockAccess returns a pointer to DlpActionVBlockAccess
func DlpActionPBlockAccess() *DlpAction {
	v := DlpActionVBlockAccess
	return &v
}

// DlpActionPDeviceRestriction returns a pointer to DlpActionVDeviceRestriction
func DlpActionPDeviceRestriction() *DlpAction {
	v := DlpActionVDeviceRestriction
	return &v
}

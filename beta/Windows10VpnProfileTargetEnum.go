// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10VpnProfileTarget undocumented
type Windows10VpnProfileTarget string

const (
	// Windows10VpnProfileTargetVUser undocumented
	Windows10VpnProfileTargetVUser Windows10VpnProfileTarget = "User"
	// Windows10VpnProfileTargetVDevice undocumented
	Windows10VpnProfileTargetVDevice Windows10VpnProfileTarget = "Device"
	// Windows10VpnProfileTargetVAutoPilotDevice undocumented
	Windows10VpnProfileTargetVAutoPilotDevice Windows10VpnProfileTarget = "AutoPilotDevice"
)

// Windows10VpnProfileTargetPUser returns a pointer to Windows10VpnProfileTargetVUser
func Windows10VpnProfileTargetPUser() *Windows10VpnProfileTarget {
	v := Windows10VpnProfileTargetVUser
	return &v
}

// Windows10VpnProfileTargetPDevice returns a pointer to Windows10VpnProfileTargetVDevice
func Windows10VpnProfileTargetPDevice() *Windows10VpnProfileTarget {
	v := Windows10VpnProfileTargetVDevice
	return &v
}

// Windows10VpnProfileTargetPAutoPilotDevice returns a pointer to Windows10VpnProfileTargetVAutoPilotDevice
func Windows10VpnProfileTargetPAutoPilotDevice() *Windows10VpnProfileTarget {
	v := Windows10VpnProfileTargetVAutoPilotDevice
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAppFlaggedReason undocumented
type ManagedAppFlaggedReason string

const (
	// ManagedAppFlaggedReasonVNone undocumented
	ManagedAppFlaggedReasonVNone ManagedAppFlaggedReason = "None"
	// ManagedAppFlaggedReasonVRootedDevice undocumented
	ManagedAppFlaggedReasonVRootedDevice ManagedAppFlaggedReason = "RootedDevice"
)

// ManagedAppFlaggedReasonPNone returns a pointer to ManagedAppFlaggedReasonVNone
func ManagedAppFlaggedReasonPNone() *ManagedAppFlaggedReason {
	v := ManagedAppFlaggedReasonVNone
	return &v
}

// ManagedAppFlaggedReasonPRootedDevice returns a pointer to ManagedAppFlaggedReasonVRootedDevice
func ManagedAppFlaggedReasonPRootedDevice() *ManagedAppFlaggedReason {
	v := ManagedAppFlaggedReasonVRootedDevice
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementExchangeAccessLevel undocumented
type DeviceManagementExchangeAccessLevel int

const (
	// DeviceManagementExchangeAccessLevelVNone undocumented
	DeviceManagementExchangeAccessLevelVNone DeviceManagementExchangeAccessLevel = 0
	// DeviceManagementExchangeAccessLevelVAllow undocumented
	DeviceManagementExchangeAccessLevelVAllow DeviceManagementExchangeAccessLevel = 1
	// DeviceManagementExchangeAccessLevelVBlock undocumented
	DeviceManagementExchangeAccessLevelVBlock DeviceManagementExchangeAccessLevel = 2
	// DeviceManagementExchangeAccessLevelVQuarantine undocumented
	DeviceManagementExchangeAccessLevelVQuarantine DeviceManagementExchangeAccessLevel = 3
)

// DeviceManagementExchangeAccessLevelPNone returns a pointer to DeviceManagementExchangeAccessLevelVNone
func DeviceManagementExchangeAccessLevelPNone() *DeviceManagementExchangeAccessLevel {
	v := DeviceManagementExchangeAccessLevelVNone
	return &v
}

// DeviceManagementExchangeAccessLevelPAllow returns a pointer to DeviceManagementExchangeAccessLevelVAllow
func DeviceManagementExchangeAccessLevelPAllow() *DeviceManagementExchangeAccessLevel {
	v := DeviceManagementExchangeAccessLevelVAllow
	return &v
}

// DeviceManagementExchangeAccessLevelPBlock returns a pointer to DeviceManagementExchangeAccessLevelVBlock
func DeviceManagementExchangeAccessLevelPBlock() *DeviceManagementExchangeAccessLevel {
	v := DeviceManagementExchangeAccessLevelVBlock
	return &v
}

// DeviceManagementExchangeAccessLevelPQuarantine returns a pointer to DeviceManagementExchangeAccessLevelVQuarantine
func DeviceManagementExchangeAccessLevelPQuarantine() *DeviceManagementExchangeAccessLevel {
	v := DeviceManagementExchangeAccessLevelVQuarantine
	return &v
}

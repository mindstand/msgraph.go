// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceHealthScriptDetectionType undocumented
type DeviceHealthScriptDetectionType int

const (
	// DeviceHealthScriptDetectionTypeVNotConfigured undocumented
	DeviceHealthScriptDetectionTypeVNotConfigured DeviceHealthScriptDetectionType = 0
	// DeviceHealthScriptDetectionTypeVString undocumented
	DeviceHealthScriptDetectionTypeVString DeviceHealthScriptDetectionType = 1
)

// DeviceHealthScriptDetectionTypePNotConfigured returns a pointer to DeviceHealthScriptDetectionTypeVNotConfigured
func DeviceHealthScriptDetectionTypePNotConfigured() *DeviceHealthScriptDetectionType {
	v := DeviceHealthScriptDetectionTypeVNotConfigured
	return &v
}

// DeviceHealthScriptDetectionTypePString returns a pointer to DeviceHealthScriptDetectionTypeVString
func DeviceHealthScriptDetectionTypePString() *DeviceHealthScriptDetectionType {
	v := DeviceHealthScriptDetectionTypeVString
	return &v
}

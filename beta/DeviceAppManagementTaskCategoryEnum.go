// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceAppManagementTaskCategory undocumented
type DeviceAppManagementTaskCategory string

const (
	// DeviceAppManagementTaskCategoryVUnknown undocumented
	DeviceAppManagementTaskCategoryVUnknown DeviceAppManagementTaskCategory = "Unknown"
	// DeviceAppManagementTaskCategoryVAdvancedThreatProtection undocumented
	DeviceAppManagementTaskCategoryVAdvancedThreatProtection DeviceAppManagementTaskCategory = "AdvancedThreatProtection"
)

// DeviceAppManagementTaskCategoryPUnknown returns a pointer to DeviceAppManagementTaskCategoryVUnknown
func DeviceAppManagementTaskCategoryPUnknown() *DeviceAppManagementTaskCategory {
	v := DeviceAppManagementTaskCategoryVUnknown
	return &v
}

// DeviceAppManagementTaskCategoryPAdvancedThreatProtection returns a pointer to DeviceAppManagementTaskCategoryVAdvancedThreatProtection
func DeviceAppManagementTaskCategoryPAdvancedThreatProtection() *DeviceAppManagementTaskCategory {
	v := DeviceAppManagementTaskCategoryVAdvancedThreatProtection
	return &v
}

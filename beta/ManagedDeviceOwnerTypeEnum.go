// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedDeviceOwnerType undocumented
type ManagedDeviceOwnerType string

const (
	// ManagedDeviceOwnerTypeVUnknown undocumented
	ManagedDeviceOwnerTypeVUnknown ManagedDeviceOwnerType = "Unknown"
	// ManagedDeviceOwnerTypeVCompany undocumented
	ManagedDeviceOwnerTypeVCompany ManagedDeviceOwnerType = "Company"
	// ManagedDeviceOwnerTypeVPersonal undocumented
	ManagedDeviceOwnerTypeVPersonal ManagedDeviceOwnerType = "Personal"
)

// ManagedDeviceOwnerTypePUnknown returns a pointer to ManagedDeviceOwnerTypeVUnknown
func ManagedDeviceOwnerTypePUnknown() *ManagedDeviceOwnerType {
	v := ManagedDeviceOwnerTypeVUnknown
	return &v
}

// ManagedDeviceOwnerTypePCompany returns a pointer to ManagedDeviceOwnerTypeVCompany
func ManagedDeviceOwnerTypePCompany() *ManagedDeviceOwnerType {
	v := ManagedDeviceOwnerTypeVCompany
	return &v
}

// ManagedDeviceOwnerTypePPersonal returns a pointer to ManagedDeviceOwnerTypeVPersonal
func ManagedDeviceOwnerTypePPersonal() *ManagedDeviceOwnerType {
	v := ManagedDeviceOwnerTypeVPersonal
	return &v
}

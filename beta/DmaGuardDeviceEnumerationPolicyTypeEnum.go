// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DmaGuardDeviceEnumerationPolicyType undocumented
type DmaGuardDeviceEnumerationPolicyType string

const (
	// DmaGuardDeviceEnumerationPolicyTypeVDeviceDefault undocumented
	DmaGuardDeviceEnumerationPolicyTypeVDeviceDefault DmaGuardDeviceEnumerationPolicyType = "DeviceDefault"
	// DmaGuardDeviceEnumerationPolicyTypeVBlockAll undocumented
	DmaGuardDeviceEnumerationPolicyTypeVBlockAll DmaGuardDeviceEnumerationPolicyType = "BlockAll"
	// DmaGuardDeviceEnumerationPolicyTypeVAllowAll undocumented
	DmaGuardDeviceEnumerationPolicyTypeVAllowAll DmaGuardDeviceEnumerationPolicyType = "AllowAll"
)

// DmaGuardDeviceEnumerationPolicyTypePDeviceDefault returns a pointer to DmaGuardDeviceEnumerationPolicyTypeVDeviceDefault
func DmaGuardDeviceEnumerationPolicyTypePDeviceDefault() *DmaGuardDeviceEnumerationPolicyType {
	v := DmaGuardDeviceEnumerationPolicyTypeVDeviceDefault
	return &v
}

// DmaGuardDeviceEnumerationPolicyTypePBlockAll returns a pointer to DmaGuardDeviceEnumerationPolicyTypeVBlockAll
func DmaGuardDeviceEnumerationPolicyTypePBlockAll() *DmaGuardDeviceEnumerationPolicyType {
	v := DmaGuardDeviceEnumerationPolicyTypeVBlockAll
	return &v
}

// DmaGuardDeviceEnumerationPolicyTypePAllowAll returns a pointer to DmaGuardDeviceEnumerationPolicyTypeVAllowAll
func DmaGuardDeviceEnumerationPolicyTypePAllowAll() *DmaGuardDeviceEnumerationPolicyType {
	v := DmaGuardDeviceEnumerationPolicyTypeVAllowAll
	return &v
}

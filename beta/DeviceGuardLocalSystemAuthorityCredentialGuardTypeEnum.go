// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceGuardLocalSystemAuthorityCredentialGuardType undocumented
type DeviceGuardLocalSystemAuthorityCredentialGuardType string

const (
	// DeviceGuardLocalSystemAuthorityCredentialGuardTypeVNotConfigured undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardTypeVNotConfigured DeviceGuardLocalSystemAuthorityCredentialGuardType = "NotConfigured"
	// DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithUEFILock undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithUEFILock DeviceGuardLocalSystemAuthorityCredentialGuardType = "EnableWithUEFILock"
	// DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithoutUEFILock undocumented
	DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithoutUEFILock DeviceGuardLocalSystemAuthorityCredentialGuardType = "EnableWithoutUEFILock"
)

// DeviceGuardLocalSystemAuthorityCredentialGuardTypePNotConfigured returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardTypeVNotConfigured
func DeviceGuardLocalSystemAuthorityCredentialGuardTypePNotConfigured() *DeviceGuardLocalSystemAuthorityCredentialGuardType {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardTypeVNotConfigured
	return &v
}

// DeviceGuardLocalSystemAuthorityCredentialGuardTypePEnableWithUEFILock returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithUEFILock
func DeviceGuardLocalSystemAuthorityCredentialGuardTypePEnableWithUEFILock() *DeviceGuardLocalSystemAuthorityCredentialGuardType {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithUEFILock
	return &v
}

// DeviceGuardLocalSystemAuthorityCredentialGuardTypePEnableWithoutUEFILock returns a pointer to DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithoutUEFILock
func DeviceGuardLocalSystemAuthorityCredentialGuardTypePEnableWithoutUEFILock() *DeviceGuardLocalSystemAuthorityCredentialGuardType {
	v := DeviceGuardLocalSystemAuthorityCredentialGuardTypeVEnableWithoutUEFILock
	return &v
}

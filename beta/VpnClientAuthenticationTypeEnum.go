// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VpnClientAuthenticationType undocumented
type VpnClientAuthenticationType string

const (
	// VpnClientAuthenticationTypeVUserAuthentication undocumented
	VpnClientAuthenticationTypeVUserAuthentication VpnClientAuthenticationType = "UserAuthentication"
	// VpnClientAuthenticationTypeVDeviceAuthentication undocumented
	VpnClientAuthenticationTypeVDeviceAuthentication VpnClientAuthenticationType = "DeviceAuthentication"
)

// VpnClientAuthenticationTypePUserAuthentication returns a pointer to VpnClientAuthenticationTypeVUserAuthentication
func VpnClientAuthenticationTypePUserAuthentication() *VpnClientAuthenticationType {
	v := VpnClientAuthenticationTypeVUserAuthentication
	return &v
}

// VpnClientAuthenticationTypePDeviceAuthentication returns a pointer to VpnClientAuthenticationTypeVDeviceAuthentication
func VpnClientAuthenticationTypePDeviceAuthentication() *VpnClientAuthenticationType {
	v := VpnClientAuthenticationTypeVDeviceAuthentication
	return &v
}

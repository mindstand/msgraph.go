// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WiFiProxySetting undocumented
type WiFiProxySetting string

const (
	// WiFiProxySettingVNone undocumented
	WiFiProxySettingVNone WiFiProxySetting = "None"
	// WiFiProxySettingVManual undocumented
	WiFiProxySettingVManual WiFiProxySetting = "Manual"
	// WiFiProxySettingVAutomatic undocumented
	WiFiProxySettingVAutomatic WiFiProxySetting = "Automatic"
)

// WiFiProxySettingPNone returns a pointer to WiFiProxySettingVNone
func WiFiProxySettingPNone() *WiFiProxySetting {
	v := WiFiProxySettingVNone
	return &v
}

// WiFiProxySettingPManual returns a pointer to WiFiProxySettingVManual
func WiFiProxySettingPManual() *WiFiProxySetting {
	v := WiFiProxySettingVManual
	return &v
}

// WiFiProxySettingPAutomatic returns a pointer to WiFiProxySettingVAutomatic
func WiFiProxySettingPAutomatic() *WiFiProxySetting {
	v := WiFiProxySettingVAutomatic
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidOMACpConfiguration By providing a configuration in this profile you can configure Android devices that support OMA-CP.
type AndroidOMACpConfiguration struct {
	DeviceConfiguration
	// ConfigurationXML Configuration XML that will be applied to the device. When it is read, it only provides a placeholder string since the original data is encrypted and stored.
	ConfigurationXML *Binary `json:"configurationXml,omitempty"`
}

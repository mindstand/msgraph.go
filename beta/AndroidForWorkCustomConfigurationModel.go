// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkCustomConfiguration Android For Work custom configuration
type AndroidForWorkCustomConfiguration struct {
	DeviceConfiguration
	// OMASettings OMA settings. This collection can contain a maximum of 500 elements.
	OMASettings []OMASetting `json:"omaSettings,omitempty"`
}

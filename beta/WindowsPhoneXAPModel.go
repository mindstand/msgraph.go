// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhoneXAP Contains properties and inherited properties for Windows Phone XAP Line Of Business apps.
type WindowsPhoneXAP struct {
	MobileLobApp
	// MinimumSupportedOperatingSystem The value for the minimum applicable operating system.
	MinimumSupportedOperatingSystem *WindowsMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
	// ProductIdentifier The Product Identifier.
	ProductIdentifier *string `json:"productIdentifier,omitempty"`
	// IdentityVersion The identity version.
	IdentityVersion *string `json:"identityVersion,omitempty"`
}

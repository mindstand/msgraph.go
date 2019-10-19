// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsMobileMSI Contains properties and inherited properties for Windows Mobile MSI Line Of Business apps.
type WindowsMobileMSI struct {
	MobileLobApp
	// CommandLine The command line.
	CommandLine *string `json:"commandLine,omitempty"`
	// ProductCode The product code.
	ProductCode *string `json:"productCode,omitempty"`
	// ProductVersion The product version of Windows Mobile MSI Line of Business (LoB) app.
	ProductVersion *string `json:"productVersion,omitempty"`
	// IgnoreVersionDetection A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature.
	IgnoreVersionDetection *bool `json:"ignoreVersionDetection,omitempty"`
}

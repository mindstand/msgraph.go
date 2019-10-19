// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Win32LobAppPowerShellScriptDetection undocumented
type Win32LobAppPowerShellScriptDetection struct {
	Win32LobAppDetection
	// EnforceSignatureCheck A value indicating whether signature check is enforced
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`
	// RunAs32Bit A value indicating whether this script should run as 32-bit
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`
	// ScriptContent The base64 encoded script content to detect Win32 Line of Business (LoB) app
	ScriptContent *string `json:"scriptContent,omitempty"`
}

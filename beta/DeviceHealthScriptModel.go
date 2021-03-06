// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// DeviceHealthScript Intune will provide customer the ability to run their Powershell Health scripts (remediation + detection) on the enrolled windows 10 Azure Active Directory joined devices.
type DeviceHealthScript struct {
	// Entity is the base model of DeviceHealthScript
	Entity
	// Publisher Name of the device health script publisher
	Publisher *string `json:"publisher,omitempty"`
	// Version Version of the device health script
	Version *string `json:"version,omitempty"`
	// DisplayName Name of the device health script
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the device health script
	Description *string `json:"description,omitempty"`
	// DetectionScriptContent The entire content of the detection powershell script
	DetectionScriptContent *Binary `json:"detectionScriptContent,omitempty"`
	// RemediationScriptContent The entire content of the remediation powershell script
	RemediationScriptContent *Binary `json:"remediationScriptContent,omitempty"`
	// CreatedDateTime The timestamp of when the device health script was created. This property is read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime The timestamp of when the device health script was modified. This property is read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// RunAsAccount Indicates the type of execution context
	RunAsAccount *RunAsAccountType `json:"runAsAccount,omitempty"`
	// EnforceSignatureCheck Indicate whether the script signature needs be checked
	EnforceSignatureCheck *bool `json:"enforceSignatureCheck,omitempty"`
	// RunAs32Bit Indicate whether PowerShell script(s) should run as 32-bit
	RunAs32Bit *bool `json:"runAs32Bit,omitempty"`
	// RoleScopeTagIDs List of Scope Tag IDs for the device health script
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
	// Assignments undocumented
	Assignments []DeviceHealthScriptAssignment `json:"assignments,omitempty"`
	// RunSummary undocumented
	RunSummary *DeviceHealthScriptRunSummary `json:"runSummary,omitempty"`
	// DeviceRunStates undocumented
	DeviceRunStates []DeviceHealthScriptDeviceState `json:"deviceRunStates,omitempty"`
}

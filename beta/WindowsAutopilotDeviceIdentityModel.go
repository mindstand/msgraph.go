// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// WindowsAutopilotDeviceIdentity The windowsAutopilotDeviceIdentity resource represents a Windows Autopilot Device.
type WindowsAutopilotDeviceIdentity struct {
	Entity
	// DeploymentProfileAssignmentStatus Profile assignment status of the Windows autopilot device.
	DeploymentProfileAssignmentStatus *WindowsAutopilotProfileAssignmentStatus `json:"deploymentProfileAssignmentStatus,omitempty"`
	// DeploymentProfileAssignmentDetailedStatus Profile assignment detailed status of the Windows autopilot device.
	DeploymentProfileAssignmentDetailedStatus *WindowsAutopilotProfileAssignmentDetailedStatus `json:"deploymentProfileAssignmentDetailedStatus,omitempty"`
	// DeploymentProfileAssignedDateTime Profile set time of the Windows autopilot device.
	DeploymentProfileAssignedDateTime *time.Time `json:"deploymentProfileAssignedDateTime,omitempty"`
	// OrderIdentifier Order Identifier of the Windows autopilot device - Deprecated
	OrderIdentifier *string `json:"orderIdentifier,omitempty"`
	// GroupTag Group Tag of the Windows autopilot device.
	GroupTag *string `json:"groupTag,omitempty"`
	// PurchaseOrderIdentifier Purchase Order Identifier of the Windows autopilot device.
	PurchaseOrderIdentifier *string `json:"purchaseOrderIdentifier,omitempty"`
	// SerialNumber Serial number of the Windows autopilot device.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// ProductKey Product Key of the Windows autopilot device.
	ProductKey *string `json:"productKey,omitempty"`
	// Manufacturer Oem manufacturer of the Windows autopilot device.
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Model Model name of the Windows autopilot device.
	Model *string `json:"model,omitempty"`
	// EnrollmentState Intune enrollment state of the Windows autopilot device.
	EnrollmentState *EnrollmentState `json:"enrollmentState,omitempty"`
	// LastContactedDateTime Intune Last Contacted Date Time of the Windows autopilot device.
	LastContactedDateTime *time.Time `json:"lastContactedDateTime,omitempty"`
	// AddressableUserName Addressable user name.
	AddressableUserName *string `json:"addressableUserName,omitempty"`
	// UserPrincipalName User Principal Name.
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	// ResourceName Resource Name.
	ResourceName *string `json:"resourceName,omitempty"`
	// SKUNumber SKU Number
	SKUNumber *string `json:"skuNumber,omitempty"`
	// SystemFamily System Family
	SystemFamily *string `json:"systemFamily,omitempty"`
	// AzureActiveDirectoryDeviceID AAD Device ID
	AzureActiveDirectoryDeviceID *string `json:"azureActiveDirectoryDeviceId,omitempty"`
	// ManagedDeviceID Managed Device ID
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
	// DeploymentProfile undocumented
	DeploymentProfile *WindowsAutopilotDeploymentProfile `json:"deploymentProfile,omitempty"`
	// IntendedDeploymentProfile undocumented
	IntendedDeploymentProfile *WindowsAutopilotDeploymentProfile `json:"intendedDeploymentProfile,omitempty"`
}

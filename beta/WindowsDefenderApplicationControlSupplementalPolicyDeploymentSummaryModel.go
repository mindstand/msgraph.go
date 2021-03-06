// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary Contains properties for the deployment summary of a WindowsDefenderApplicationControl supplemental policy.
type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary struct {
	// Entity is the base model of WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary
	Entity
	// DeployedDeviceCount Number of Devices that have successfully deployed this WindowsDefenderApplicationControl supplemental policy.
	DeployedDeviceCount *int `json:"deployedDeviceCount,omitempty"`
	// FailedDeviceCount Number of Devices that have failed to deploy this WindowsDefenderApplicationControl supplemental policy.
	FailedDeviceCount *int `json:"failedDeviceCount,omitempty"`
}

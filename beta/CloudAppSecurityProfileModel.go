// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// CloudAppSecurityProfile undocumented
type CloudAppSecurityProfile struct {
	Entity
	// AzureSubscriptionID undocumented
	AzureSubscriptionID *string `json:"azureSubscriptionId,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DeploymentPackageURL undocumented
	DeploymentPackageURL *string `json:"deploymentPackageUrl,omitempty"`
	// DestinationServiceName undocumented
	DestinationServiceName *string `json:"destinationServiceName,omitempty"`
	// IsSigned undocumented
	IsSigned *bool `json:"isSigned,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Manifest undocumented
	Manifest *string `json:"manifest,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// PermissionsRequired undocumented
	PermissionsRequired *ApplicationPermissionsRequired `json:"permissionsRequired,omitempty"`
	// Platform undocumented
	Platform *string `json:"platform,omitempty"`
	// PolicyName undocumented
	PolicyName *string `json:"policyName,omitempty"`
	// Publisher undocumented
	Publisher *string `json:"publisher,omitempty"`
	// RiskScore undocumented
	RiskScore *string `json:"riskScore,omitempty"`
	// Tags undocumented
	Tags []string `json:"tags,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// VendorInformation undocumented
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}

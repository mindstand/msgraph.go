// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ProviderTenantSetting undocumented
type ProviderTenantSetting struct {
	Entity
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// Enabled undocumented
	Enabled *bool `json:"enabled,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Provider undocumented
	Provider *string `json:"provider,omitempty"`
	// Vendor undocumented
	Vendor *string `json:"vendor,omitempty"`
}

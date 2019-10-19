// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SecureScore undocumented
type SecureScore struct {
	Entity
	// ActiveUserCount undocumented
	ActiveUserCount *int `json:"activeUserCount,omitempty"`
	// AverageComparativeScores undocumented
	AverageComparativeScores []AverageComparativeScore `json:"averageComparativeScores,omitempty"`
	// AzureTenantID undocumented
	AzureTenantID *string `json:"azureTenantId,omitempty"`
	// ControlScores undocumented
	ControlScores []ControlScore `json:"controlScores,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// CurrentScore undocumented
	CurrentScore *float64 `json:"currentScore,omitempty"`
	// EnabledServices undocumented
	EnabledServices []string `json:"enabledServices,omitempty"`
	// LicensedUserCount undocumented
	LicensedUserCount *int `json:"licensedUserCount,omitempty"`
	// MaxScore undocumented
	MaxScore *float64 `json:"maxScore,omitempty"`
	// VendorInformation undocumented
	VendorInformation *SecurityVendorInformation `json:"vendorInformation,omitempty"`
}

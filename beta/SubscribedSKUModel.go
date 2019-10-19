// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SubscribedSKU undocumented
type SubscribedSKU struct {
	Entity
	// CapabilityStatus undocumented
	CapabilityStatus *string `json:"capabilityStatus,omitempty"`
	// ConsumedUnits undocumented
	ConsumedUnits *int `json:"consumedUnits,omitempty"`
	// PrepaidUnits undocumented
	PrepaidUnits *LicenseUnitsDetail `json:"prepaidUnits,omitempty"`
	// ServicePlans undocumented
	ServicePlans []ServicePlanInfo `json:"servicePlans,omitempty"`
	// SKUID undocumented
	SKUID *UUID `json:"skuId,omitempty"`
	// SKUPartNumber undocumented
	SKUPartNumber *string `json:"skuPartNumber,omitempty"`
	// AppliesTo undocumented
	AppliesTo *string `json:"appliesTo,omitempty"`
}

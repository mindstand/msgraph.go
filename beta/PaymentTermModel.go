// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PaymentTerm undocumented
type PaymentTerm struct {
	Entity
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DueDateCalculation undocumented
	DueDateCalculation *string `json:"dueDateCalculation,omitempty"`
	// DiscountDateCalculation undocumented
	DiscountDateCalculation *string `json:"discountDateCalculation,omitempty"`
	// DiscountPercent undocumented
	DiscountPercent *int `json:"discountPercent,omitempty"`
	// CalculateDiscountOnCreditMemos undocumented
	CalculateDiscountOnCreditMemos *bool `json:"calculateDiscountOnCreditMemos,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

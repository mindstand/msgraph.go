// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Account undocumented
type Account struct {
	Entity
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Category undocumented
	Category *string `json:"category,omitempty"`
	// SubCategory undocumented
	SubCategory *string `json:"subCategory,omitempty"`
	// Blocked undocumented
	Blocked *bool `json:"blocked,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

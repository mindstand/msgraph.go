// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PlannerExternalReference undocumented
type PlannerExternalReference struct {
	// Alias undocumented
	Alias *string `json:"alias,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// PreviewPriority undocumented
	PreviewPriority *string `json:"previewPriority,omitempty"`
	// LastModifiedBy undocumented
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ContentClassification undocumented
type ContentClassification struct {
	// SensitiveTypeID undocumented
	SensitiveTypeID *string `json:"sensitiveTypeId,omitempty"`
	// UniqueCount undocumented
	UniqueCount *int `json:"uniqueCount,omitempty"`
	// Confidence undocumented
	Confidence *int `json:"confidence,omitempty"`
	// Matches undocumented
	Matches []MatchLocation `json:"matches,omitempty"`
}

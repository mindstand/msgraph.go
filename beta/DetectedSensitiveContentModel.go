// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DetectedSensitiveContent undocumented
type DetectedSensitiveContent struct {
	// ID undocumented
	ID *UUID `json:"id,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// UniqueCount undocumented
	UniqueCount *int `json:"uniqueCount,omitempty"`
	// Confidence undocumented
	Confidence *int `json:"confidence,omitempty"`
	// RecommendedConfidence undocumented
	RecommendedConfidence *int `json:"recommendedConfidence,omitempty"`
	// Matches undocumented
	Matches []SensitiveContentLocation `json:"matches,omitempty"`
}

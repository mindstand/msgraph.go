// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ScopeOperatorMultiValuedComparisonType undocumented
type ScopeOperatorMultiValuedComparisonType string

const (
	// ScopeOperatorMultiValuedComparisonTypeVAll undocumented
	ScopeOperatorMultiValuedComparisonTypeVAll ScopeOperatorMultiValuedComparisonType = "All"
	// ScopeOperatorMultiValuedComparisonTypeVAny undocumented
	ScopeOperatorMultiValuedComparisonTypeVAny ScopeOperatorMultiValuedComparisonType = "Any"
)

// ScopeOperatorMultiValuedComparisonTypePAll returns a pointer to ScopeOperatorMultiValuedComparisonTypeVAll
func ScopeOperatorMultiValuedComparisonTypePAll() *ScopeOperatorMultiValuedComparisonType {
	v := ScopeOperatorMultiValuedComparisonTypeVAll
	return &v
}

// ScopeOperatorMultiValuedComparisonTypePAny returns a pointer to ScopeOperatorMultiValuedComparisonTypeVAny
func ScopeOperatorMultiValuedComparisonTypePAny() *ScopeOperatorMultiValuedComparisonType {
	v := ScopeOperatorMultiValuedComparisonTypeVAny
	return &v
}

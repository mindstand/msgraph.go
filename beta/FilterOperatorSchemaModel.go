// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FilterOperatorSchema undocumented
type FilterOperatorSchema struct {
	Entity
	// Arity undocumented
	Arity *ScopeOperatorType `json:"arity,omitempty"`
	// MultivaluedComparisonType undocumented
	MultivaluedComparisonType *ScopeOperatorMultiValuedComparisonType `json:"multivaluedComparisonType,omitempty"`
	// SupportedAttributeTypes undocumented
	SupportedAttributeTypes []AttributeType `json:"supportedAttributeTypes,omitempty"`
}

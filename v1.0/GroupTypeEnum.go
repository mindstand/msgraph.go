// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupType undocumented
type GroupType string

const (
	// GroupTypeVUnifiedGroups undocumented
	GroupTypeVUnifiedGroups GroupType = "UnifiedGroups"
	// GroupTypeVAzureAD undocumented
	GroupTypeVAzureAD GroupType = "AzureAD"
	// GroupTypeVUnknownFutureValue undocumented
	GroupTypeVUnknownFutureValue GroupType = "UnknownFutureValue"
)

// GroupTypePUnifiedGroups returns a pointer to GroupTypeVUnifiedGroups
func GroupTypePUnifiedGroups() *GroupType {
	v := GroupTypeVUnifiedGroups
	return &v
}

// GroupTypePAzureAD returns a pointer to GroupTypeVAzureAD
func GroupTypePAzureAD() *GroupType {
	v := GroupTypeVAzureAD
	return &v
}

// GroupTypePUnknownFutureValue returns a pointer to GroupTypeVUnknownFutureValue
func GroupTypePUnknownFutureValue() *GroupType {
	v := GroupTypeVUnknownFutureValue
	return &v
}

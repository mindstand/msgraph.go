// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SelectionLikelihoodInfo undocumented
type SelectionLikelihoodInfo string

const (
	// SelectionLikelihoodInfoVNotSpecified undocumented
	SelectionLikelihoodInfoVNotSpecified SelectionLikelihoodInfo = "NotSpecified"
	// SelectionLikelihoodInfoVHigh undocumented
	SelectionLikelihoodInfoVHigh SelectionLikelihoodInfo = "High"
)

// SelectionLikelihoodInfoPNotSpecified returns a pointer to SelectionLikelihoodInfoVNotSpecified
func SelectionLikelihoodInfoPNotSpecified() *SelectionLikelihoodInfo {
	v := SelectionLikelihoodInfoVNotSpecified
	return &v
}

// SelectionLikelihoodInfoPHigh returns a pointer to SelectionLikelihoodInfoVHigh
func SelectionLikelihoodInfoPHigh() *SelectionLikelihoodInfo {
	v := SelectionLikelihoodInfoVHigh
	return &v
}

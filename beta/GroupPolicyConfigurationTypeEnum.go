// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyConfigurationType undocumented
type GroupPolicyConfigurationType int

const (
	// GroupPolicyConfigurationTypeVPolicy undocumented
	GroupPolicyConfigurationTypeVPolicy GroupPolicyConfigurationType = 0
	// GroupPolicyConfigurationTypeVPreference undocumented
	GroupPolicyConfigurationTypeVPreference GroupPolicyConfigurationType = 1
)

// GroupPolicyConfigurationTypePPolicy returns a pointer to GroupPolicyConfigurationTypeVPolicy
func GroupPolicyConfigurationTypePPolicy() *GroupPolicyConfigurationType {
	v := GroupPolicyConfigurationTypeVPolicy
	return &v
}

// GroupPolicyConfigurationTypePPreference returns a pointer to GroupPolicyConfigurationTypeVPreference
func GroupPolicyConfigurationTypePPreference() *GroupPolicyConfigurationType {
	v := GroupPolicyConfigurationTypeVPreference
	return &v
}

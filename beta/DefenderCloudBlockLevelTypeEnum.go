// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DefenderCloudBlockLevelType undocumented
type DefenderCloudBlockLevelType string

const (
	// DefenderCloudBlockLevelTypeVNotConfigured undocumented
	DefenderCloudBlockLevelTypeVNotConfigured DefenderCloudBlockLevelType = "NotConfigured"
	// DefenderCloudBlockLevelTypeVHigh undocumented
	DefenderCloudBlockLevelTypeVHigh DefenderCloudBlockLevelType = "High"
	// DefenderCloudBlockLevelTypeVHighPlus undocumented
	DefenderCloudBlockLevelTypeVHighPlus DefenderCloudBlockLevelType = "HighPlus"
	// DefenderCloudBlockLevelTypeVZeroTolerance undocumented
	DefenderCloudBlockLevelTypeVZeroTolerance DefenderCloudBlockLevelType = "ZeroTolerance"
)

// DefenderCloudBlockLevelTypePNotConfigured returns a pointer to DefenderCloudBlockLevelTypeVNotConfigured
func DefenderCloudBlockLevelTypePNotConfigured() *DefenderCloudBlockLevelType {
	v := DefenderCloudBlockLevelTypeVNotConfigured
	return &v
}

// DefenderCloudBlockLevelTypePHigh returns a pointer to DefenderCloudBlockLevelTypeVHigh
func DefenderCloudBlockLevelTypePHigh() *DefenderCloudBlockLevelType {
	v := DefenderCloudBlockLevelTypeVHigh
	return &v
}

// DefenderCloudBlockLevelTypePHighPlus returns a pointer to DefenderCloudBlockLevelTypeVHighPlus
func DefenderCloudBlockLevelTypePHighPlus() *DefenderCloudBlockLevelType {
	v := DefenderCloudBlockLevelTypeVHighPlus
	return &v
}

// DefenderCloudBlockLevelTypePZeroTolerance returns a pointer to DefenderCloudBlockLevelTypeVZeroTolerance
func DefenderCloudBlockLevelTypePZeroTolerance() *DefenderCloudBlockLevelType {
	v := DefenderCloudBlockLevelTypeVZeroTolerance
	return &v
}

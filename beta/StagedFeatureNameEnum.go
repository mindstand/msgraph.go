// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// StagedFeatureName undocumented
type StagedFeatureName string

const (
	// StagedFeatureNameVPassthroughAuthentication undocumented
	StagedFeatureNameVPassthroughAuthentication StagedFeatureName = "PassthroughAuthentication"
	// StagedFeatureNameVSeamlessSso undocumented
	StagedFeatureNameVSeamlessSso StagedFeatureName = "SeamlessSso"
	// StagedFeatureNameVPasswordHashSync undocumented
	StagedFeatureNameVPasswordHashSync StagedFeatureName = "PasswordHashSync"
	// StagedFeatureNameVUnknownFutureValue undocumented
	StagedFeatureNameVUnknownFutureValue StagedFeatureName = "UnknownFutureValue"
)

// StagedFeatureNamePPassthroughAuthentication returns a pointer to StagedFeatureNameVPassthroughAuthentication
func StagedFeatureNamePPassthroughAuthentication() *StagedFeatureName {
	v := StagedFeatureNameVPassthroughAuthentication
	return &v
}

// StagedFeatureNamePSeamlessSso returns a pointer to StagedFeatureNameVSeamlessSso
func StagedFeatureNamePSeamlessSso() *StagedFeatureName {
	v := StagedFeatureNameVSeamlessSso
	return &v
}

// StagedFeatureNamePPasswordHashSync returns a pointer to StagedFeatureNameVPasswordHashSync
func StagedFeatureNamePPasswordHashSync() *StagedFeatureName {
	v := StagedFeatureNameVPasswordHashSync
	return &v
}

// StagedFeatureNamePUnknownFutureValue returns a pointer to StagedFeatureNameVUnknownFutureValue
func StagedFeatureNamePUnknownFutureValue() *StagedFeatureName {
	v := StagedFeatureNameVUnknownFutureValue
	return &v
}

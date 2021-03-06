// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// InternetSiteSecurityLevel undocumented
type InternetSiteSecurityLevel string

const (
	// InternetSiteSecurityLevelVUserDefined undocumented
	InternetSiteSecurityLevelVUserDefined InternetSiteSecurityLevel = "UserDefined"
	// InternetSiteSecurityLevelVMedium undocumented
	InternetSiteSecurityLevelVMedium InternetSiteSecurityLevel = "Medium"
	// InternetSiteSecurityLevelVMediumHigh undocumented
	InternetSiteSecurityLevelVMediumHigh InternetSiteSecurityLevel = "MediumHigh"
	// InternetSiteSecurityLevelVHigh undocumented
	InternetSiteSecurityLevelVHigh InternetSiteSecurityLevel = "High"
)

// InternetSiteSecurityLevelPUserDefined returns a pointer to InternetSiteSecurityLevelVUserDefined
func InternetSiteSecurityLevelPUserDefined() *InternetSiteSecurityLevel {
	v := InternetSiteSecurityLevelVUserDefined
	return &v
}

// InternetSiteSecurityLevelPMedium returns a pointer to InternetSiteSecurityLevelVMedium
func InternetSiteSecurityLevelPMedium() *InternetSiteSecurityLevel {
	v := InternetSiteSecurityLevelVMedium
	return &v
}

// InternetSiteSecurityLevelPMediumHigh returns a pointer to InternetSiteSecurityLevelVMediumHigh
func InternetSiteSecurityLevelPMediumHigh() *InternetSiteSecurityLevel {
	v := InternetSiteSecurityLevelVMediumHigh
	return &v
}

// InternetSiteSecurityLevelPHigh returns a pointer to InternetSiteSecurityLevelVHigh
func InternetSiteSecurityLevelPHigh() *InternetSiteSecurityLevel {
	v := InternetSiteSecurityLevelVHigh
	return &v
}

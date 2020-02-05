// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CloudAppSecuritySessionControlType undocumented
type CloudAppSecuritySessionControlType string

const (
	// CloudAppSecuritySessionControlTypeVMcasConfigured undocumented
	CloudAppSecuritySessionControlTypeVMcasConfigured CloudAppSecuritySessionControlType = "McasConfigured"
	// CloudAppSecuritySessionControlTypeVMonitorOnly undocumented
	CloudAppSecuritySessionControlTypeVMonitorOnly CloudAppSecuritySessionControlType = "MonitorOnly"
	// CloudAppSecuritySessionControlTypeVBlockDownloads undocumented
	CloudAppSecuritySessionControlTypeVBlockDownloads CloudAppSecuritySessionControlType = "BlockDownloads"
)

// CloudAppSecuritySessionControlTypePMcasConfigured returns a pointer to CloudAppSecuritySessionControlTypeVMcasConfigured
func CloudAppSecuritySessionControlTypePMcasConfigured() *CloudAppSecuritySessionControlType {
	v := CloudAppSecuritySessionControlTypeVMcasConfigured
	return &v
}

// CloudAppSecuritySessionControlTypePMonitorOnly returns a pointer to CloudAppSecuritySessionControlTypeVMonitorOnly
func CloudAppSecuritySessionControlTypePMonitorOnly() *CloudAppSecuritySessionControlType {
	v := CloudAppSecuritySessionControlTypeVMonitorOnly
	return &v
}

// CloudAppSecuritySessionControlTypePBlockDownloads returns a pointer to CloudAppSecuritySessionControlTypeVBlockDownloads
func CloudAppSecuritySessionControlTypePBlockDownloads() *CloudAppSecuritySessionControlType {
	v := CloudAppSecuritySessionControlTypeVBlockDownloads
	return &v
}
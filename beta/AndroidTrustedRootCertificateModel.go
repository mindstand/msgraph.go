// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidTrustedRootCertificate Android Trusted Root Certificate configuration profile
type AndroidTrustedRootCertificate struct {
	DeviceConfiguration
	// TrustedRootCertificate Trusted Root Certificate
	TrustedRootCertificate *Binary `json:"trustedRootCertificate,omitempty"`
	// CertFileName File name to display in UI.
	CertFileName *string `json:"certFileName,omitempty"`
}

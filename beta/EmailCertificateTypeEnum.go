// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EmailCertificateType undocumented
type EmailCertificateType string

const (
	// EmailCertificateTypeVNone undocumented
	EmailCertificateTypeVNone EmailCertificateType = "None"
	// EmailCertificateTypeVCertificate undocumented
	EmailCertificateTypeVCertificate EmailCertificateType = "Certificate"
	// EmailCertificateTypeVDerivedCredential undocumented
	EmailCertificateTypeVDerivedCredential EmailCertificateType = "DerivedCredential"
)

// EmailCertificateTypePNone returns a pointer to EmailCertificateTypeVNone
func EmailCertificateTypePNone() *EmailCertificateType {
	v := EmailCertificateTypeVNone
	return &v
}

// EmailCertificateTypePCertificate returns a pointer to EmailCertificateTypeVCertificate
func EmailCertificateTypePCertificate() *EmailCertificateType {
	v := EmailCertificateTypeVCertificate
	return &v
}

// EmailCertificateTypePDerivedCredential returns a pointer to EmailCertificateTypeVDerivedCredential
func EmailCertificateTypePDerivedCredential() *EmailCertificateType {
	v := EmailCertificateTypeVDerivedCredential
	return &v
}

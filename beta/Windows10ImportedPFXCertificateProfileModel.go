// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10ImportedPFXCertificateProfile Windows 10 Desktop and Mobile PFX Import certificate profile
type Windows10ImportedPFXCertificateProfile struct {
	WindowsCertificateProfileBase
	// IntendedPurpose undocumented
	IntendedPurpose *IntendedPurpose `json:"intendedPurpose,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}

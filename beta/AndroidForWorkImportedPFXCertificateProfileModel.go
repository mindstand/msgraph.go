// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkImportedPFXCertificateProfile Android For Work PFX Import certificate profile
type AndroidForWorkImportedPFXCertificateProfile struct {
	AndroidCertificateProfileBase
	// IntendedPurpose undocumented
	IntendedPurpose *IntendedPurpose `json:"intendedPurpose,omitempty"`
	// ManagedDeviceCertificateStates undocumented
	ManagedDeviceCertificateStates []ManagedDeviceCertificateState `json:"managedDeviceCertificateStates,omitempty"`
}

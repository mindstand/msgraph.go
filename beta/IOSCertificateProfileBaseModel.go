// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSCertificateProfileBase iOS certificate profile base.
type IOSCertificateProfileBase struct {
	IOSCertificateProfile
	// RenewalThresholdPercentage Certificate renewal threshold percentage. Valid values 1 to 99
	RenewalThresholdPercentage *int `json:"renewalThresholdPercentage,omitempty"`
	// SubjectNameFormat Certificate Subject Name Format.
	SubjectNameFormat *AppleSubjectNameFormat `json:"subjectNameFormat,omitempty"`
	// SubjectAlternativeNameType Certificate Subject Alternative Name type.
	SubjectAlternativeNameType *SubjectAlternativeNameType `json:"subjectAlternativeNameType,omitempty"`
	// CertificateValidityPeriodValue Value for the Certificate Validity Period.
	CertificateValidityPeriodValue *int `json:"certificateValidityPeriodValue,omitempty"`
	// CertificateValidityPeriodScale Scale for the Certificate Validity Period.
	CertificateValidityPeriodScale *CertificateValidityPeriodScale `json:"certificateValidityPeriodScale,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CertificateValidityPeriodScale undocumented
type CertificateValidityPeriodScale string

const (
	// CertificateValidityPeriodScaleVDays undocumented
	CertificateValidityPeriodScaleVDays CertificateValidityPeriodScale = "Days"
	// CertificateValidityPeriodScaleVMonths undocumented
	CertificateValidityPeriodScaleVMonths CertificateValidityPeriodScale = "Months"
	// CertificateValidityPeriodScaleVYears undocumented
	CertificateValidityPeriodScaleVYears CertificateValidityPeriodScale = "Years"
)

// CertificateValidityPeriodScalePDays returns a pointer to CertificateValidityPeriodScaleVDays
func CertificateValidityPeriodScalePDays() *CertificateValidityPeriodScale {
	v := CertificateValidityPeriodScaleVDays
	return &v
}

// CertificateValidityPeriodScalePMonths returns a pointer to CertificateValidityPeriodScaleVMonths
func CertificateValidityPeriodScalePMonths() *CertificateValidityPeriodScale {
	v := CertificateValidityPeriodScaleVMonths
	return &v
}

// CertificateValidityPeriodScalePYears returns a pointer to CertificateValidityPeriodScaleVYears
func CertificateValidityPeriodScalePYears() *CertificateValidityPeriodScale {
	v := CertificateValidityPeriodScaleVYears
	return &v
}

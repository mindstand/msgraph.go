// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Sensitivity undocumented
type Sensitivity string

const (
	// SensitivityVNormal undocumented
	SensitivityVNormal Sensitivity = "Normal"
	// SensitivityVPersonal undocumented
	SensitivityVPersonal Sensitivity = "Personal"
	// SensitivityVPrivate undocumented
	SensitivityVPrivate Sensitivity = "Private"
	// SensitivityVConfidential undocumented
	SensitivityVConfidential Sensitivity = "Confidential"
)

// SensitivityPNormal returns a pointer to SensitivityVNormal
func SensitivityPNormal() *Sensitivity {
	v := SensitivityVNormal
	return &v
}

// SensitivityPPersonal returns a pointer to SensitivityVPersonal
func SensitivityPPersonal() *Sensitivity {
	v := SensitivityVPersonal
	return &v
}

// SensitivityPPrivate returns a pointer to SensitivityVPrivate
func SensitivityPPrivate() *Sensitivity {
	v := SensitivityVPrivate
	return &v
}

// SensitivityPConfidential returns a pointer to SensitivityVConfidential
func SensitivityPConfidential() *Sensitivity {
	v := SensitivityVConfidential
	return &v
}

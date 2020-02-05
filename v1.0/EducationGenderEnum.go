// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationGender undocumented
type EducationGender string

const (
	// EducationGenderVFemale undocumented
	EducationGenderVFemale EducationGender = "Female"
	// EducationGenderVMale undocumented
	EducationGenderVMale EducationGender = "Male"
	// EducationGenderVOther undocumented
	EducationGenderVOther EducationGender = "Other"
	// EducationGenderVUnknownFutureValue undocumented
	EducationGenderVUnknownFutureValue EducationGender = "UnknownFutureValue"
)

// EducationGenderPFemale returns a pointer to EducationGenderVFemale
func EducationGenderPFemale() *EducationGender {
	v := EducationGenderVFemale
	return &v
}

// EducationGenderPMale returns a pointer to EducationGenderVMale
func EducationGenderPMale() *EducationGender {
	v := EducationGenderVMale
	return &v
}

// EducationGenderPOther returns a pointer to EducationGenderVOther
func EducationGenderPOther() *EducationGender {
	v := EducationGenderVOther
	return &v
}

// EducationGenderPUnknownFutureValue returns a pointer to EducationGenderVUnknownFutureValue
func EducationGenderPUnknownFutureValue() *EducationGender {
	v := EducationGenderVUnknownFutureValue
	return &v
}

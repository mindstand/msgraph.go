// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlannerPreviewType undocumented
type PlannerPreviewType string

const (
	// PlannerPreviewTypeVAutomatic undocumented
	PlannerPreviewTypeVAutomatic PlannerPreviewType = "Automatic"
	// PlannerPreviewTypeVNoPreview undocumented
	PlannerPreviewTypeVNoPreview PlannerPreviewType = "NoPreview"
	// PlannerPreviewTypeVChecklist undocumented
	PlannerPreviewTypeVChecklist PlannerPreviewType = "Checklist"
	// PlannerPreviewTypeVDescription undocumented
	PlannerPreviewTypeVDescription PlannerPreviewType = "Description"
	// PlannerPreviewTypeVReference undocumented
	PlannerPreviewTypeVReference PlannerPreviewType = "Reference"
)

// PlannerPreviewTypePAutomatic returns a pointer to PlannerPreviewTypeVAutomatic
func PlannerPreviewTypePAutomatic() *PlannerPreviewType {
	v := PlannerPreviewTypeVAutomatic
	return &v
}

// PlannerPreviewTypePNoPreview returns a pointer to PlannerPreviewTypeVNoPreview
func PlannerPreviewTypePNoPreview() *PlannerPreviewType {
	v := PlannerPreviewTypeVNoPreview
	return &v
}

// PlannerPreviewTypePChecklist returns a pointer to PlannerPreviewTypeVChecklist
func PlannerPreviewTypePChecklist() *PlannerPreviewType {
	v := PlannerPreviewTypeVChecklist
	return &v
}

// PlannerPreviewTypePDescription returns a pointer to PlannerPreviewTypeVDescription
func PlannerPreviewTypePDescription() *PlannerPreviewType {
	v := PlannerPreviewTypeVDescription
	return &v
}

// PlannerPreviewTypePReference returns a pointer to PlannerPreviewTypeVReference
func PlannerPreviewTypePReference() *PlannerPreviewType {
	v := PlannerPreviewTypeVReference
	return &v
}

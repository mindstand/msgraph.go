// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// InferenceClassificationType undocumented
type InferenceClassificationType int

const (
	// InferenceClassificationTypeVFocused undocumented
	InferenceClassificationTypeVFocused InferenceClassificationType = 0
	// InferenceClassificationTypeVOther undocumented
	InferenceClassificationTypeVOther InferenceClassificationType = 1
)

// InferenceClassificationTypePFocused returns a pointer to InferenceClassificationTypeVFocused
func InferenceClassificationTypePFocused() *InferenceClassificationType {
	v := InferenceClassificationTypeVFocused
	return &v
}

// InferenceClassificationTypePOther returns a pointer to InferenceClassificationTypeVOther
func InferenceClassificationTypePOther() *InferenceClassificationType {
	v := InferenceClassificationTypeVOther
	return &v
}

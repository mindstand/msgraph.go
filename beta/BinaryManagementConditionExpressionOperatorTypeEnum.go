// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BinaryManagementConditionExpressionOperatorType undocumented
type BinaryManagementConditionExpressionOperatorType string

const (
	// BinaryManagementConditionExpressionOperatorTypeVOr undocumented
	BinaryManagementConditionExpressionOperatorTypeVOr BinaryManagementConditionExpressionOperatorType = "Or"
	// BinaryManagementConditionExpressionOperatorTypeVAnd undocumented
	BinaryManagementConditionExpressionOperatorTypeVAnd BinaryManagementConditionExpressionOperatorType = "And"
)

// BinaryManagementConditionExpressionOperatorTypePOr returns a pointer to BinaryManagementConditionExpressionOperatorTypeVOr
func BinaryManagementConditionExpressionOperatorTypePOr() *BinaryManagementConditionExpressionOperatorType {
	v := BinaryManagementConditionExpressionOperatorTypeVOr
	return &v
}

// BinaryManagementConditionExpressionOperatorTypePAnd returns a pointer to BinaryManagementConditionExpressionOperatorTypeVAnd
func BinaryManagementConditionExpressionOperatorTypePAnd() *BinaryManagementConditionExpressionOperatorType {
	v := BinaryManagementConditionExpressionOperatorTypeVAnd
	return &v
}

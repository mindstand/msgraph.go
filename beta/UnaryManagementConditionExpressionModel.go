// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UnaryManagementConditionExpression undocumented
type UnaryManagementConditionExpression struct {
	ManagementConditionExpressionModel
	// Operator The operator used in the evaluation of the unary operation.
	Operator *UnaryManagementConditionExpressionOperatorType `json:"operator,omitempty"`
	// Operand The operand of the unary operation.
	Operand *ManagementConditionExpressionModel `json:"operand,omitempty"`
}

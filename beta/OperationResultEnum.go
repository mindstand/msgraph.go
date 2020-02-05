// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OperationResult undocumented
type OperationResult string

const (
	// OperationResultVSuccess undocumented
	OperationResultVSuccess OperationResult = "Success"
	// OperationResultVFailure undocumented
	OperationResultVFailure OperationResult = "Failure"
	// OperationResultVTimeout undocumented
	OperationResultVTimeout OperationResult = "Timeout"
	// OperationResultVUnknownFutureValue undocumented
	OperationResultVUnknownFutureValue OperationResult = "UnknownFutureValue"
)

// OperationResultPSuccess returns a pointer to OperationResultVSuccess
func OperationResultPSuccess() *OperationResult {
	v := OperationResultVSuccess
	return &v
}

// OperationResultPFailure returns a pointer to OperationResultVFailure
func OperationResultPFailure() *OperationResult {
	v := OperationResultVFailure
	return &v
}

// OperationResultPTimeout returns a pointer to OperationResultVTimeout
func OperationResultPTimeout() *OperationResult {
	v := OperationResultVTimeout
	return &v
}

// OperationResultPUnknownFutureValue returns a pointer to OperationResultVUnknownFutureValue
func OperationResultPUnknownFutureValue() *OperationResult {
	v := OperationResultVUnknownFutureValue
	return &v
}

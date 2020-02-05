// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConditionalAccessStatus undocumented
type ConditionalAccessStatus string

const (
	// ConditionalAccessStatusVSuccess undocumented
	ConditionalAccessStatusVSuccess ConditionalAccessStatus = "Success"
	// ConditionalAccessStatusVFailure undocumented
	ConditionalAccessStatusVFailure ConditionalAccessStatus = "Failure"
	// ConditionalAccessStatusVNotApplied undocumented
	ConditionalAccessStatusVNotApplied ConditionalAccessStatus = "NotApplied"
	// ConditionalAccessStatusVUnknownFutureValue undocumented
	ConditionalAccessStatusVUnknownFutureValue ConditionalAccessStatus = "UnknownFutureValue"
)

// ConditionalAccessStatusPSuccess returns a pointer to ConditionalAccessStatusVSuccess
func ConditionalAccessStatusPSuccess() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVSuccess
	return &v
}

// ConditionalAccessStatusPFailure returns a pointer to ConditionalAccessStatusVFailure
func ConditionalAccessStatusPFailure() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVFailure
	return &v
}

// ConditionalAccessStatusPNotApplied returns a pointer to ConditionalAccessStatusVNotApplied
func ConditionalAccessStatusPNotApplied() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVNotApplied
	return &v
}

// ConditionalAccessStatusPUnknownFutureValue returns a pointer to ConditionalAccessStatusVUnknownFutureValue
func ConditionalAccessStatusPUnknownFutureValue() *ConditionalAccessStatus {
	v := ConditionalAccessStatusVUnknownFutureValue
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RejectReason undocumented
type RejectReason string

const (
	// RejectReasonVNone undocumented
	RejectReasonVNone RejectReason = "None"
	// RejectReasonVBusy undocumented
	RejectReasonVBusy RejectReason = "Busy"
	// RejectReasonVForbidden undocumented
	RejectReasonVForbidden RejectReason = "Forbidden"
	// RejectReasonVUnknownFutureValue undocumented
	RejectReasonVUnknownFutureValue RejectReason = "UnknownFutureValue"
)

// RejectReasonPNone returns a pointer to RejectReasonVNone
func RejectReasonPNone() *RejectReason {
	v := RejectReasonVNone
	return &v
}

// RejectReasonPBusy returns a pointer to RejectReasonVBusy
func RejectReasonPBusy() *RejectReason {
	v := RejectReasonVBusy
	return &v
}

// RejectReasonPForbidden returns a pointer to RejectReasonVForbidden
func RejectReasonPForbidden() *RejectReason {
	v := RejectReasonVForbidden
	return &v
}

// RejectReasonPUnknownFutureValue returns a pointer to RejectReasonVUnknownFutureValue
func RejectReasonPUnknownFutureValue() *RejectReason {
	v := RejectReasonVUnknownFutureValue
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OperationStatus undocumented
type OperationStatus string

const (
	// OperationStatusVNotStarted undocumented
	OperationStatusVNotStarted OperationStatus = "NotStarted"
	// OperationStatusVRunning undocumented
	OperationStatusVRunning OperationStatus = "Running"
	// OperationStatusVCompleted undocumented
	OperationStatusVCompleted OperationStatus = "Completed"
	// OperationStatusVFailed undocumented
	OperationStatusVFailed OperationStatus = "Failed"
)

// OperationStatusPNotStarted returns a pointer to OperationStatusVNotStarted
func OperationStatusPNotStarted() *OperationStatus {
	v := OperationStatusVNotStarted
	return &v
}

// OperationStatusPRunning returns a pointer to OperationStatusVRunning
func OperationStatusPRunning() *OperationStatus {
	v := OperationStatusVRunning
	return &v
}

// OperationStatusPCompleted returns a pointer to OperationStatusVCompleted
func OperationStatusPCompleted() *OperationStatus {
	v := OperationStatusVCompleted
	return &v
}

// OperationStatusPFailed returns a pointer to OperationStatusVFailed
func OperationStatusPFailed() *OperationStatus {
	v := OperationStatusVFailed
	return &v
}

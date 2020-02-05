// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RegistryOperation undocumented
type RegistryOperation string

const (
	// RegistryOperationVUnknown undocumented
	RegistryOperationVUnknown RegistryOperation = "Unknown"
	// RegistryOperationVCreate undocumented
	RegistryOperationVCreate RegistryOperation = "Create"
	// RegistryOperationVModify undocumented
	RegistryOperationVModify RegistryOperation = "Modify"
	// RegistryOperationVDelete undocumented
	RegistryOperationVDelete RegistryOperation = "Delete"
	// RegistryOperationVUnknownFutureValue undocumented
	RegistryOperationVUnknownFutureValue RegistryOperation = "UnknownFutureValue"
)

// RegistryOperationPUnknown returns a pointer to RegistryOperationVUnknown
func RegistryOperationPUnknown() *RegistryOperation {
	v := RegistryOperationVUnknown
	return &v
}

// RegistryOperationPCreate returns a pointer to RegistryOperationVCreate
func RegistryOperationPCreate() *RegistryOperation {
	v := RegistryOperationVCreate
	return &v
}

// RegistryOperationPModify returns a pointer to RegistryOperationVModify
func RegistryOperationPModify() *RegistryOperation {
	v := RegistryOperationVModify
	return &v
}

// RegistryOperationPDelete returns a pointer to RegistryOperationVDelete
func RegistryOperationPDelete() *RegistryOperation {
	v := RegistryOperationVDelete
	return &v
}

// RegistryOperationPUnknownFutureValue returns a pointer to RegistryOperationVUnknownFutureValue
func RegistryOperationPUnknownFutureValue() *RegistryOperation {
	v := RegistryOperationVUnknownFutureValue
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TiAction undocumented
type TiAction string

const (
	// TiActionVUnknown undocumented
	TiActionVUnknown TiAction = "Unknown"
	// TiActionVAllow undocumented
	TiActionVAllow TiAction = "Allow"
	// TiActionVBlock undocumented
	TiActionVBlock TiAction = "Block"
	// TiActionVAlert undocumented
	TiActionVAlert TiAction = "Alert"
	// TiActionVUnknownFutureValue undocumented
	TiActionVUnknownFutureValue TiAction = "UnknownFutureValue"
)

// TiActionPUnknown returns a pointer to TiActionVUnknown
func TiActionPUnknown() *TiAction {
	v := TiActionVUnknown
	return &v
}

// TiActionPAllow returns a pointer to TiActionVAllow
func TiActionPAllow() *TiAction {
	v := TiActionVAllow
	return &v
}

// TiActionPBlock returns a pointer to TiActionVBlock
func TiActionPBlock() *TiAction {
	v := TiActionVBlock
	return &v
}

// TiActionPAlert returns a pointer to TiActionVAlert
func TiActionPAlert() *TiAction {
	v := TiActionVAlert
	return &v
}

// TiActionPUnknownFutureValue returns a pointer to TiActionVUnknownFutureValue
func TiActionPUnknownFutureValue() *TiAction {
	v := TiActionVUnknownFutureValue
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConnectionStatus undocumented
type ConnectionStatus int

const (
	// ConnectionStatusVUnknown undocumented
	ConnectionStatusVUnknown ConnectionStatus = 0
	// ConnectionStatusVAttempted undocumented
	ConnectionStatusVAttempted ConnectionStatus = 1
	// ConnectionStatusVSucceeded undocumented
	ConnectionStatusVSucceeded ConnectionStatus = 2
	// ConnectionStatusVBlocked undocumented
	ConnectionStatusVBlocked ConnectionStatus = 3
	// ConnectionStatusVFailed undocumented
	ConnectionStatusVFailed ConnectionStatus = 4
	// ConnectionStatusVUnknownFutureValue undocumented
	ConnectionStatusVUnknownFutureValue ConnectionStatus = 127
)

// ConnectionStatusPUnknown returns a pointer to ConnectionStatusVUnknown
func ConnectionStatusPUnknown() *ConnectionStatus {
	v := ConnectionStatusVUnknown
	return &v
}

// ConnectionStatusPAttempted returns a pointer to ConnectionStatusVAttempted
func ConnectionStatusPAttempted() *ConnectionStatus {
	v := ConnectionStatusVAttempted
	return &v
}

// ConnectionStatusPSucceeded returns a pointer to ConnectionStatusVSucceeded
func ConnectionStatusPSucceeded() *ConnectionStatus {
	v := ConnectionStatusVSucceeded
	return &v
}

// ConnectionStatusPBlocked returns a pointer to ConnectionStatusVBlocked
func ConnectionStatusPBlocked() *ConnectionStatus {
	v := ConnectionStatusVBlocked
	return &v
}

// ConnectionStatusPFailed returns a pointer to ConnectionStatusVFailed
func ConnectionStatusPFailed() *ConnectionStatus {
	v := ConnectionStatusVFailed
	return &v
}

// ConnectionStatusPUnknownFutureValue returns a pointer to ConnectionStatusVUnknownFutureValue
func ConnectionStatusPUnknownFutureValue() *ConnectionStatus {
	v := ConnectionStatusVUnknownFutureValue
	return &v
}

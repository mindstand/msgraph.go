// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SynchronizationJobRestartScope undocumented
type SynchronizationJobRestartScope string

const (
	// SynchronizationJobRestartScopeVForceDeletes undocumented
	SynchronizationJobRestartScopeVForceDeletes SynchronizationJobRestartScope = "ForceDeletes"
	// SynchronizationJobRestartScopeVFull undocumented
	SynchronizationJobRestartScopeVFull SynchronizationJobRestartScope = "Full"
	// SynchronizationJobRestartScopeVQuarantineState undocumented
	SynchronizationJobRestartScopeVQuarantineState SynchronizationJobRestartScope = "QuarantineState"
	// SynchronizationJobRestartScopeVWatermark undocumented
	SynchronizationJobRestartScopeVWatermark SynchronizationJobRestartScope = "Watermark"
	// SynchronizationJobRestartScopeVEscrows undocumented
	SynchronizationJobRestartScopeVEscrows SynchronizationJobRestartScope = "Escrows"
	// SynchronizationJobRestartScopeVConnectorDataStore undocumented
	SynchronizationJobRestartScopeVConnectorDataStore SynchronizationJobRestartScope = "ConnectorDataStore"
	// SynchronizationJobRestartScopeVNone undocumented
	SynchronizationJobRestartScopeVNone SynchronizationJobRestartScope = "None"
)

// SynchronizationJobRestartScopePForceDeletes returns a pointer to SynchronizationJobRestartScopeVForceDeletes
func SynchronizationJobRestartScopePForceDeletes() *SynchronizationJobRestartScope {
	v := SynchronizationJobRestartScopeVForceDeletes
	return &v
}

// SynchronizationJobRestartScopePFull returns a pointer to SynchronizationJobRestartScopeVFull
func SynchronizationJobRestartScopePFull() *SynchronizationJobRestartScope {
	v := SynchronizationJobRestartScopeVFull
	return &v
}

// SynchronizationJobRestartScopePQuarantineState returns a pointer to SynchronizationJobRestartScopeVQuarantineState
func SynchronizationJobRestartScopePQuarantineState() *SynchronizationJobRestartScope {
	v := SynchronizationJobRestartScopeVQuarantineState
	return &v
}

// SynchronizationJobRestartScopePWatermark returns a pointer to SynchronizationJobRestartScopeVWatermark
func SynchronizationJobRestartScopePWatermark() *SynchronizationJobRestartScope {
	v := SynchronizationJobRestartScopeVWatermark
	return &v
}

// SynchronizationJobRestartScopePEscrows returns a pointer to SynchronizationJobRestartScopeVEscrows
func SynchronizationJobRestartScopePEscrows() *SynchronizationJobRestartScope {
	v := SynchronizationJobRestartScopeVEscrows
	return &v
}

// SynchronizationJobRestartScopePConnectorDataStore returns a pointer to SynchronizationJobRestartScopeVConnectorDataStore
func SynchronizationJobRestartScopePConnectorDataStore() *SynchronizationJobRestartScope {
	v := SynchronizationJobRestartScopeVConnectorDataStore
	return &v
}

// SynchronizationJobRestartScopePNone returns a pointer to SynchronizationJobRestartScopeVNone
func SynchronizationJobRestartScopePNone() *SynchronizationJobRestartScope {
	v := SynchronizationJobRestartScopeVNone
	return &v
}

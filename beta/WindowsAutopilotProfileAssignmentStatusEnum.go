// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsAutopilotProfileAssignmentStatus undocumented
type WindowsAutopilotProfileAssignmentStatus string

const (
	// WindowsAutopilotProfileAssignmentStatusVUnknown undocumented
	WindowsAutopilotProfileAssignmentStatusVUnknown WindowsAutopilotProfileAssignmentStatus = "Unknown"
	// WindowsAutopilotProfileAssignmentStatusVAssignedInSync undocumented
	WindowsAutopilotProfileAssignmentStatusVAssignedInSync WindowsAutopilotProfileAssignmentStatus = "AssignedInSync"
	// WindowsAutopilotProfileAssignmentStatusVAssignedOutOfSync undocumented
	WindowsAutopilotProfileAssignmentStatusVAssignedOutOfSync WindowsAutopilotProfileAssignmentStatus = "AssignedOutOfSync"
	// WindowsAutopilotProfileAssignmentStatusVAssignedUnkownSyncState undocumented
	WindowsAutopilotProfileAssignmentStatusVAssignedUnkownSyncState WindowsAutopilotProfileAssignmentStatus = "AssignedUnkownSyncState"
	// WindowsAutopilotProfileAssignmentStatusVNotAssigned undocumented
	WindowsAutopilotProfileAssignmentStatusVNotAssigned WindowsAutopilotProfileAssignmentStatus = "NotAssigned"
	// WindowsAutopilotProfileAssignmentStatusVPending undocumented
	WindowsAutopilotProfileAssignmentStatusVPending WindowsAutopilotProfileAssignmentStatus = "Pending"
	// WindowsAutopilotProfileAssignmentStatusVFailed undocumented
	WindowsAutopilotProfileAssignmentStatusVFailed WindowsAutopilotProfileAssignmentStatus = "Failed"
)

// WindowsAutopilotProfileAssignmentStatusPUnknown returns a pointer to WindowsAutopilotProfileAssignmentStatusVUnknown
func WindowsAutopilotProfileAssignmentStatusPUnknown() *WindowsAutopilotProfileAssignmentStatus {
	v := WindowsAutopilotProfileAssignmentStatusVUnknown
	return &v
}

// WindowsAutopilotProfileAssignmentStatusPAssignedInSync returns a pointer to WindowsAutopilotProfileAssignmentStatusVAssignedInSync
func WindowsAutopilotProfileAssignmentStatusPAssignedInSync() *WindowsAutopilotProfileAssignmentStatus {
	v := WindowsAutopilotProfileAssignmentStatusVAssignedInSync
	return &v
}

// WindowsAutopilotProfileAssignmentStatusPAssignedOutOfSync returns a pointer to WindowsAutopilotProfileAssignmentStatusVAssignedOutOfSync
func WindowsAutopilotProfileAssignmentStatusPAssignedOutOfSync() *WindowsAutopilotProfileAssignmentStatus {
	v := WindowsAutopilotProfileAssignmentStatusVAssignedOutOfSync
	return &v
}

// WindowsAutopilotProfileAssignmentStatusPAssignedUnkownSyncState returns a pointer to WindowsAutopilotProfileAssignmentStatusVAssignedUnkownSyncState
func WindowsAutopilotProfileAssignmentStatusPAssignedUnkownSyncState() *WindowsAutopilotProfileAssignmentStatus {
	v := WindowsAutopilotProfileAssignmentStatusVAssignedUnkownSyncState
	return &v
}

// WindowsAutopilotProfileAssignmentStatusPNotAssigned returns a pointer to WindowsAutopilotProfileAssignmentStatusVNotAssigned
func WindowsAutopilotProfileAssignmentStatusPNotAssigned() *WindowsAutopilotProfileAssignmentStatus {
	v := WindowsAutopilotProfileAssignmentStatusVNotAssigned
	return &v
}

// WindowsAutopilotProfileAssignmentStatusPPending returns a pointer to WindowsAutopilotProfileAssignmentStatusVPending
func WindowsAutopilotProfileAssignmentStatusPPending() *WindowsAutopilotProfileAssignmentStatus {
	v := WindowsAutopilotProfileAssignmentStatusVPending
	return &v
}

// WindowsAutopilotProfileAssignmentStatusPFailed returns a pointer to WindowsAutopilotProfileAssignmentStatusVFailed
func WindowsAutopilotProfileAssignmentStatusPFailed() *WindowsAutopilotProfileAssignmentStatus {
	v := WindowsAutopilotProfileAssignmentStatusVFailed
	return &v
}

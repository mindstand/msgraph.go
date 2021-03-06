// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementReportStatus undocumented
type DeviceManagementReportStatus string

const (
	// DeviceManagementReportStatusVUnknown undocumented
	DeviceManagementReportStatusVUnknown DeviceManagementReportStatus = "Unknown"
	// DeviceManagementReportStatusVNotStarted undocumented
	DeviceManagementReportStatusVNotStarted DeviceManagementReportStatus = "NotStarted"
	// DeviceManagementReportStatusVInProgress undocumented
	DeviceManagementReportStatusVInProgress DeviceManagementReportStatus = "InProgress"
	// DeviceManagementReportStatusVCompleted undocumented
	DeviceManagementReportStatusVCompleted DeviceManagementReportStatus = "Completed"
	// DeviceManagementReportStatusVFailed undocumented
	DeviceManagementReportStatusVFailed DeviceManagementReportStatus = "Failed"
)

// DeviceManagementReportStatusPUnknown returns a pointer to DeviceManagementReportStatusVUnknown
func DeviceManagementReportStatusPUnknown() *DeviceManagementReportStatus {
	v := DeviceManagementReportStatusVUnknown
	return &v
}

// DeviceManagementReportStatusPNotStarted returns a pointer to DeviceManagementReportStatusVNotStarted
func DeviceManagementReportStatusPNotStarted() *DeviceManagementReportStatus {
	v := DeviceManagementReportStatusVNotStarted
	return &v
}

// DeviceManagementReportStatusPInProgress returns a pointer to DeviceManagementReportStatusVInProgress
func DeviceManagementReportStatusPInProgress() *DeviceManagementReportStatus {
	v := DeviceManagementReportStatusVInProgress
	return &v
}

// DeviceManagementReportStatusPCompleted returns a pointer to DeviceManagementReportStatusVCompleted
func DeviceManagementReportStatusPCompleted() *DeviceManagementReportStatus {
	v := DeviceManagementReportStatusVCompleted
	return &v
}

// DeviceManagementReportStatusPFailed returns a pointer to DeviceManagementReportStatusVFailed
func DeviceManagementReportStatusPFailed() *DeviceManagementReportStatus {
	v := DeviceManagementReportStatusVFailed
	return &v
}

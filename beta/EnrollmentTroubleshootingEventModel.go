// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EnrollmentTroubleshootingEvent Event representing an enrollment failure.
type EnrollmentTroubleshootingEvent struct {
	DeviceManagementTroubleshootingEvent
	// ManagedDeviceIdentifier Device identifier created or collected by Intune.
	ManagedDeviceIdentifier *string `json:"managedDeviceIdentifier,omitempty"`
	// OperatingSystem Operating System.
	OperatingSystem *string `json:"operatingSystem,omitempty"`
	// OsVersion OS Version.
	OsVersion *string `json:"osVersion,omitempty"`
	// UserID Identifier for the user that tried to enroll the device.
	UserID *string `json:"userId,omitempty"`
	// DeviceID Azure AD device identifier.
	DeviceID *string `json:"deviceId,omitempty"`
	// EnrollmentType Type of the enrollment.
	EnrollmentType *DeviceEnrollmentType `json:"enrollmentType,omitempty"`
	// FailureCategory Highlevel failure category.
	FailureCategory *DeviceEnrollmentFailureReason `json:"failureCategory,omitempty"`
	// FailureReason Detailed failure reason.
	FailureReason *string `json:"failureReason,omitempty"`
}

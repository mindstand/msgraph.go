// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MeetingMessageType undocumented
type MeetingMessageType string

const (
	// MeetingMessageTypeVNone undocumented
	MeetingMessageTypeVNone MeetingMessageType = "None"
	// MeetingMessageTypeVMeetingRequest undocumented
	MeetingMessageTypeVMeetingRequest MeetingMessageType = "MeetingRequest"
	// MeetingMessageTypeVMeetingCancelled undocumented
	MeetingMessageTypeVMeetingCancelled MeetingMessageType = "MeetingCancelled"
	// MeetingMessageTypeVMeetingAccepted undocumented
	MeetingMessageTypeVMeetingAccepted MeetingMessageType = "MeetingAccepted"
	// MeetingMessageTypeVMeetingTenativelyAccepted undocumented
	MeetingMessageTypeVMeetingTenativelyAccepted MeetingMessageType = "MeetingTenativelyAccepted"
	// MeetingMessageTypeVMeetingDeclined undocumented
	MeetingMessageTypeVMeetingDeclined MeetingMessageType = "MeetingDeclined"
)

// MeetingMessageTypePNone returns a pointer to MeetingMessageTypeVNone
func MeetingMessageTypePNone() *MeetingMessageType {
	v := MeetingMessageTypeVNone
	return &v
}

// MeetingMessageTypePMeetingRequest returns a pointer to MeetingMessageTypeVMeetingRequest
func MeetingMessageTypePMeetingRequest() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingRequest
	return &v
}

// MeetingMessageTypePMeetingCancelled returns a pointer to MeetingMessageTypeVMeetingCancelled
func MeetingMessageTypePMeetingCancelled() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingCancelled
	return &v
}

// MeetingMessageTypePMeetingAccepted returns a pointer to MeetingMessageTypeVMeetingAccepted
func MeetingMessageTypePMeetingAccepted() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingAccepted
	return &v
}

// MeetingMessageTypePMeetingTenativelyAccepted returns a pointer to MeetingMessageTypeVMeetingTenativelyAccepted
func MeetingMessageTypePMeetingTenativelyAccepted() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingTenativelyAccepted
	return &v
}

// MeetingMessageTypePMeetingDeclined returns a pointer to MeetingMessageTypeVMeetingDeclined
func MeetingMessageTypePMeetingDeclined() *MeetingMessageType {
	v := MeetingMessageTypeVMeetingDeclined
	return &v
}

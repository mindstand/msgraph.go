// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MeetingCapability undocumented
type MeetingCapability struct {
	// Object is the base model of MeetingCapability
	Object
	// AllowAnonymousUsersToDialOut undocumented
	AllowAnonymousUsersToDialOut *bool `json:"allowAnonymousUsersToDialOut,omitempty"`
	// AutoAdmittedUsers undocumented
	AutoAdmittedUsers *AutoAdmittedUsersType `json:"autoAdmittedUsers,omitempty"`
	// AllowAnonymousUsersToStartMeeting undocumented
	AllowAnonymousUsersToStartMeeting *bool `json:"allowAnonymousUsersToStartMeeting,omitempty"`
}

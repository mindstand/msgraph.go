// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ShiftItem undocumented
type ShiftItem struct {
	ScheduleEntity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Notes undocumented
	Notes *string `json:"notes,omitempty"`
	// Activities undocumented
	Activities []ShiftActivity `json:"activities,omitempty"`
}

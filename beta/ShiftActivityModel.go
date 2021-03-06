// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ShiftActivity undocumented
type ShiftActivity struct {
	// Object is the base model of ShiftActivity
	Object
	// IsPaid undocumented
	IsPaid *bool `json:"isPaid,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Theme undocumented
	Theme *ScheduleEntityTheme `json:"theme,omitempty"`
}

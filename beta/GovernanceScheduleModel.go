// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// GovernanceSchedule undocumented
type GovernanceSchedule struct {
	// Object is the base model of GovernanceSchedule
	Object
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Duration undocumented
	Duration *time.Duration `json:"duration,omitempty"`
}

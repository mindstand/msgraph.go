// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// CallActivityStatistics undocumented
type CallActivityStatistics struct {
	ActivityStatistics
	// AfterHours undocumented
	AfterHours *time.Duration `json:"afterHours,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ItemActivityStat undocumented
type ItemActivityStat struct {
	Entity
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Access undocumented
	Access *ItemActionStat `json:"access,omitempty"`
	// Create undocumented
	Create *ItemActionStat `json:"create,omitempty"`
	// Delete undocumented
	Delete *ItemActionStat `json:"delete,omitempty"`
	// Edit undocumented
	Edit *ItemActionStat `json:"edit,omitempty"`
	// Move undocumented
	Move *ItemActionStat `json:"move,omitempty"`
	// IsTrending undocumented
	IsTrending *bool `json:"isTrending,omitempty"`
	// IncompleteData undocumented
	IncompleteData *IncompleteData `json:"incompleteData,omitempty"`
	// Activities undocumented
	Activities []ItemActivity `json:"activities,omitempty"`
}

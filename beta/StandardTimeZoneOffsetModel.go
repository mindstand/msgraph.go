// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// StandardTimeZoneOffset undocumented
type StandardTimeZoneOffset struct {
	// Time undocumented
	Time *time.Time `json:"time,omitempty"`
	// DayOccurrence undocumented
	DayOccurrence *int `json:"dayOccurrence,omitempty"`
	// DayOfWeek undocumented
	DayOfWeek *DayOfWeek `json:"dayOfWeek,omitempty"`
	// Month undocumented
	Month *int `json:"month,omitempty"`
	// Year undocumented
	Year *int `json:"year,omitempty"`
}
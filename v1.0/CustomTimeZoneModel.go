// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CustomTimeZone undocumented
type CustomTimeZone struct {
	TimeZoneBase
	// Bias undocumented
	Bias *int `json:"bias,omitempty"`
	// StandardOffset undocumented
	StandardOffset *StandardTimeZoneOffset `json:"standardOffset,omitempty"`
	// DaylightOffset undocumented
	DaylightOffset *DaylightTimeZoneOffset `json:"daylightOffset,omitempty"`
}

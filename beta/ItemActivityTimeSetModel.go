// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ItemActivityTimeSet undocumented
type ItemActivityTimeSet struct {
	// Object is the base model of ItemActivityTimeSet
	Object
	// LastRecordedDateTime undocumented
	LastRecordedDateTime *time.Time `json:"lastRecordedDateTime,omitempty"`
	// ObservedDateTime undocumented
	ObservedDateTime *time.Time `json:"observedDateTime,omitempty"`
	// RecordedDateTime undocumented
	RecordedDateTime *time.Time `json:"recordedDateTime,omitempty"`
}

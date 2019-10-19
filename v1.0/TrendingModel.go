// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Trending undocumented
type Trending struct {
	Entity
	// Weight undocumented
	Weight *float64 `json:"weight,omitempty"`
	// ResourceVisualization undocumented
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
	// ResourceReference undocumented
	ResourceReference *ResourceReference `json:"resourceReference,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Resource undocumented
	Resource *Entity `json:"resource,omitempty"`
}

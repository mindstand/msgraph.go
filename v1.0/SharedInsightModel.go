// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SharedInsight undocumented
type SharedInsight struct {
	Entity
	// LastShared undocumented
	LastShared *SharingDetail `json:"lastShared,omitempty"`
	// SharingHistory undocumented
	SharingHistory []SharingDetail `json:"sharingHistory,omitempty"`
	// ResourceVisualization undocumented
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
	// ResourceReference undocumented
	ResourceReference *ResourceReference `json:"resourceReference,omitempty"`
	// LastSharedMethod undocumented
	LastSharedMethod *Entity `json:"lastSharedMethod,omitempty"`
	// Resource undocumented
	Resource *Entity `json:"resource,omitempty"`
}

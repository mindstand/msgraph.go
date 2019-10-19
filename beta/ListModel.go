// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// List undocumented
type List struct {
	BaseItem
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// List undocumented
	List *ListInfo `json:"list,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// System undocumented
	System *SystemFacet `json:"system,omitempty"`
	// Activities undocumented
	Activities []ItemActivityOLD `json:"activities,omitempty"`
	// Columns undocumented
	Columns []ColumnDefinition `json:"columns,omitempty"`
	// ContentTypes undocumented
	ContentTypes []ContentType `json:"contentTypes,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Items undocumented
	Items []ListItem `json:"items,omitempty"`
	// Subscriptions undocumented
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
}

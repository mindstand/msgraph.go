// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ListItem undocumented
type ListItem struct {
	// BaseItem is the base model of ListItem
	BaseItem
	// ContentType undocumented
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// Analytics undocumented
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
	// Fields undocumented
	Fields *FieldValueSet `json:"fields,omitempty"`
	// Versions undocumented
	Versions []ListItemVersion `json:"versions,omitempty"`
}

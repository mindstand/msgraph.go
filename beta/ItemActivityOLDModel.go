// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ItemActivityOLD undocumented
type ItemActivityOLD struct {
	Entity
	// Action undocumented
	Action *ItemActionSet `json:"action,omitempty"`
	// Actor undocumented
	Actor *IdentitySet `json:"actor,omitempty"`
	// Times undocumented
	Times *ItemActivityTimeSet `json:"times,omitempty"`
	// DriveItem undocumented
	DriveItem *DriveItem `json:"driveItem,omitempty"`
	// ListItem undocumented
	ListItem *ListItem `json:"listItem,omitempty"`
}

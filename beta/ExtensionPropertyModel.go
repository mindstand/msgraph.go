// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ExtensionProperty undocumented
type ExtensionProperty struct {
	DirectoryObject
	// AppDisplayName undocumented
	AppDisplayName *string `json:"appDisplayName,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// DataType undocumented
	DataType *string `json:"dataType,omitempty"`
	// IsSyncedFromOnPremises undocumented
	IsSyncedFromOnPremises *bool `json:"isSyncedFromOnPremises,omitempty"`
	// TargetObjects undocumented
	TargetObjects []string `json:"targetObjects,omitempty"`
}

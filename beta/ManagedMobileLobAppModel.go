// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedMobileLobApp An abstract base class containing properties for all managed mobile line of business apps.
type ManagedMobileLobApp struct {
	ManagedApp
	// CommittedContentVersion The internal committed content version.
	CommittedContentVersion *string `json:"committedContentVersion,omitempty"`
	// FileName The name of the main Lob application file.
	FileName *string `json:"fileName,omitempty"`
	// Size The total size, including all uploaded files.
	Size *int `json:"size,omitempty"`
	// ContentVersions undocumented
	ContentVersions []MobileAppContent `json:"contentVersions,omitempty"`
}

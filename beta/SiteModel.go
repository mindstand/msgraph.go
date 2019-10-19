// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Site undocumented
type Site struct {
	BaseItem
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Root undocumented
	Root *Root `json:"root,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// SiteCollection undocumented
	SiteCollection *SiteCollection `json:"siteCollection,omitempty"`
	// Analytics undocumented
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// Columns undocumented
	Columns []ColumnDefinition `json:"columns,omitempty"`
	// ContentTypes undocumented
	ContentTypes []ContentType `json:"contentTypes,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Drives undocumented
	Drives []Drive `json:"drives,omitempty"`
	// Items undocumented
	Items []BaseItem `json:"items,omitempty"`
	// Lists undocumented
	Lists []List `json:"lists,omitempty"`
	// Pages undocumented
	Pages []SitePage `json:"pages,omitempty"`
	// Sites undocumented
	Sites []Site `json:"sites,omitempty"`
	// Onenote undocumented
	Onenote *Onenote `json:"onenote,omitempty"`
}

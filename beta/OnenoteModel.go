// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Onenote undocumented
type Onenote struct {
	Entity
	// Notebooks undocumented
	Notebooks []Notebook `json:"notebooks,omitempty"`
	// Sections undocumented
	Sections []OnenoteSection `json:"sections,omitempty"`
	// SectionGroups undocumented
	SectionGroups []SectionGroup `json:"sectionGroups,omitempty"`
	// Pages undocumented
	Pages []OnenotePage `json:"pages,omitempty"`
	// Resources undocumented
	Resources []OnenoteResource `json:"resources,omitempty"`
	// Operations undocumented
	Operations []OnenoteOperation `json:"operations,omitempty"`
}
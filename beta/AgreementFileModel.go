// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AgreementFile undocumented
type AgreementFile struct {
	Entity
	// Language undocumented
	Language *string `json:"language,omitempty"`
	// FileName undocumented
	FileName *string `json:"fileName,omitempty"`
	// FileData undocumented
	FileData *AgreementFileData `json:"fileData,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
}

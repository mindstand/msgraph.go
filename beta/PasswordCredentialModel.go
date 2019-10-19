// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// PasswordCredential undocumented
type PasswordCredential struct {
	// CustomKeyIdentifier undocumented
	CustomKeyIdentifier *Binary `json:"customKeyIdentifier,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// KeyID undocumented
	KeyID *UUID `json:"keyId,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// SecretText undocumented
	SecretText *string `json:"secretText,omitempty"`
	// Hint undocumented
	Hint *string `json:"hint,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

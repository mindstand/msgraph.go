// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PublicError undocumented
type PublicError struct {
	// Code undocumented
	Code *string `json:"code,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Target undocumented
	Target *string `json:"target,omitempty"`
	// Details undocumented
	Details []PublicErrorDetail `json:"details,omitempty"`
	// InnerError undocumented
	InnerError *PublicInnerError `json:"innerError,omitempty"`
}

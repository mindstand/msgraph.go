// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// NonEapAuthenticationMethodForPeap undocumented
type NonEapAuthenticationMethodForPeap string

const (
	// NonEapAuthenticationMethodForPeapVNone undocumented
	NonEapAuthenticationMethodForPeapVNone NonEapAuthenticationMethodForPeap = "None"
	// NonEapAuthenticationMethodForPeapVMicrosoftChapVersionTwo undocumented
	NonEapAuthenticationMethodForPeapVMicrosoftChapVersionTwo NonEapAuthenticationMethodForPeap = "MicrosoftChapVersionTwo"
)

// NonEapAuthenticationMethodForPeapPNone returns a pointer to NonEapAuthenticationMethodForPeapVNone
func NonEapAuthenticationMethodForPeapPNone() *NonEapAuthenticationMethodForPeap {
	v := NonEapAuthenticationMethodForPeapVNone
	return &v
}

// NonEapAuthenticationMethodForPeapPMicrosoftChapVersionTwo returns a pointer to NonEapAuthenticationMethodForPeapVMicrosoftChapVersionTwo
func NonEapAuthenticationMethodForPeapPMicrosoftChapVersionTwo() *NonEapAuthenticationMethodForPeap {
	v := NonEapAuthenticationMethodForPeapVMicrosoftChapVersionTwo
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10VpnAuthenticationMethod undocumented
type Windows10VpnAuthenticationMethod string

const (
	// Windows10VpnAuthenticationMethodVCertificate undocumented
	Windows10VpnAuthenticationMethodVCertificate Windows10VpnAuthenticationMethod = "Certificate"
	// Windows10VpnAuthenticationMethodVUsernameAndPassword undocumented
	Windows10VpnAuthenticationMethodVUsernameAndPassword Windows10VpnAuthenticationMethod = "UsernameAndPassword"
	// Windows10VpnAuthenticationMethodVCustomEapXML undocumented
	Windows10VpnAuthenticationMethodVCustomEapXML Windows10VpnAuthenticationMethod = "CustomEapXML"
)

// Windows10VpnAuthenticationMethodPCertificate returns a pointer to Windows10VpnAuthenticationMethodVCertificate
func Windows10VpnAuthenticationMethodPCertificate() *Windows10VpnAuthenticationMethod {
	v := Windows10VpnAuthenticationMethodVCertificate
	return &v
}

// Windows10VpnAuthenticationMethodPUsernameAndPassword returns a pointer to Windows10VpnAuthenticationMethodVUsernameAndPassword
func Windows10VpnAuthenticationMethodPUsernameAndPassword() *Windows10VpnAuthenticationMethod {
	v := Windows10VpnAuthenticationMethodVUsernameAndPassword
	return &v
}

// Windows10VpnAuthenticationMethodPCustomEapXML returns a pointer to Windows10VpnAuthenticationMethodVCustomEapXML
func Windows10VpnAuthenticationMethodPCustomEapXML() *Windows10VpnAuthenticationMethod {
	v := Windows10VpnAuthenticationMethodVCustomEapXML
	return &v
}

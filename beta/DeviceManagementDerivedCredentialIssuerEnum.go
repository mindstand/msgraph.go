// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementDerivedCredentialIssuer undocumented
type DeviceManagementDerivedCredentialIssuer string

const (
	// DeviceManagementDerivedCredentialIssuerVIntercede undocumented
	DeviceManagementDerivedCredentialIssuerVIntercede DeviceManagementDerivedCredentialIssuer = "Intercede"
	// DeviceManagementDerivedCredentialIssuerVEntrustDatacard undocumented
	DeviceManagementDerivedCredentialIssuerVEntrustDatacard DeviceManagementDerivedCredentialIssuer = "EntrustDatacard"
	// DeviceManagementDerivedCredentialIssuerVPurebred undocumented
	DeviceManagementDerivedCredentialIssuerVPurebred DeviceManagementDerivedCredentialIssuer = "Purebred"
)

// DeviceManagementDerivedCredentialIssuerPIntercede returns a pointer to DeviceManagementDerivedCredentialIssuerVIntercede
func DeviceManagementDerivedCredentialIssuerPIntercede() *DeviceManagementDerivedCredentialIssuer {
	v := DeviceManagementDerivedCredentialIssuerVIntercede
	return &v
}

// DeviceManagementDerivedCredentialIssuerPEntrustDatacard returns a pointer to DeviceManagementDerivedCredentialIssuerVEntrustDatacard
func DeviceManagementDerivedCredentialIssuerPEntrustDatacard() *DeviceManagementDerivedCredentialIssuer {
	v := DeviceManagementDerivedCredentialIssuerVEntrustDatacard
	return &v
}

// DeviceManagementDerivedCredentialIssuerPPurebred returns a pointer to DeviceManagementDerivedCredentialIssuerVPurebred
func DeviceManagementDerivedCredentialIssuerPPurebred() *DeviceManagementDerivedCredentialIssuer {
	v := DeviceManagementDerivedCredentialIssuerVPurebred
	return &v
}

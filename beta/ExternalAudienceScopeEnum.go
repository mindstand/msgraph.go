// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ExternalAudienceScope undocumented
type ExternalAudienceScope string

const (
	// ExternalAudienceScopeVNone undocumented
	ExternalAudienceScopeVNone ExternalAudienceScope = "None"
	// ExternalAudienceScopeVContactsOnly undocumented
	ExternalAudienceScopeVContactsOnly ExternalAudienceScope = "ContactsOnly"
	// ExternalAudienceScopeVAll undocumented
	ExternalAudienceScopeVAll ExternalAudienceScope = "All"
)

// ExternalAudienceScopePNone returns a pointer to ExternalAudienceScopeVNone
func ExternalAudienceScopePNone() *ExternalAudienceScope {
	v := ExternalAudienceScopeVNone
	return &v
}

// ExternalAudienceScopePContactsOnly returns a pointer to ExternalAudienceScopeVContactsOnly
func ExternalAudienceScopePContactsOnly() *ExternalAudienceScope {
	v := ExternalAudienceScopeVContactsOnly
	return &v
}

// ExternalAudienceScopePAll returns a pointer to ExternalAudienceScopeVAll
func ExternalAudienceScopePAll() *ExternalAudienceScope {
	v := ExternalAudienceScopeVAll
	return &v
}

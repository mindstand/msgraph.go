// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ApplicationType undocumented
type ApplicationType string

const (
	// ApplicationTypeVUniversal undocumented
	ApplicationTypeVUniversal ApplicationType = "Universal"
	// ApplicationTypeVDesktop undocumented
	ApplicationTypeVDesktop ApplicationType = "Desktop"
)

// ApplicationTypePUniversal returns a pointer to ApplicationTypeVUniversal
func ApplicationTypePUniversal() *ApplicationType {
	v := ApplicationTypeVUniversal
	return &v
}

// ApplicationTypePDesktop returns a pointer to ApplicationTypeVDesktop
func ApplicationTypePDesktop() *ApplicationType {
	v := ApplicationTypeVDesktop
	return &v
}

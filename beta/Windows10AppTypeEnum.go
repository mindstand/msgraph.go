// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10AppType undocumented
type Windows10AppType int

const (
	// Windows10AppTypeVDesktop undocumented
	Windows10AppTypeVDesktop Windows10AppType = 0
	// Windows10AppTypeVUniversal undocumented
	Windows10AppTypeVUniversal Windows10AppType = 1
)

// Windows10AppTypePDesktop returns a pointer to Windows10AppTypeVDesktop
func Windows10AppTypePDesktop() *Windows10AppType {
	v := Windows10AppTypeVDesktop
	return &v
}

// Windows10AppTypePUniversal returns a pointer to Windows10AppTypeVUniversal
func Windows10AppTypePUniversal() *Windows10AppType {
	v := Windows10AppTypeVUniversal
	return &v
}

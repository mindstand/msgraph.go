// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PersistentBrowserSessionMode undocumented
type PersistentBrowserSessionMode string

const (
	// PersistentBrowserSessionModeVAlways undocumented
	PersistentBrowserSessionModeVAlways PersistentBrowserSessionMode = "Always"
	// PersistentBrowserSessionModeVNever undocumented
	PersistentBrowserSessionModeVNever PersistentBrowserSessionMode = "Never"
)

// PersistentBrowserSessionModePAlways returns a pointer to PersistentBrowserSessionModeVAlways
func PersistentBrowserSessionModePAlways() *PersistentBrowserSessionMode {
	v := PersistentBrowserSessionModeVAlways
	return &v
}

// PersistentBrowserSessionModePNever returns a pointer to PersistentBrowserSessionModeVNever
func PersistentBrowserSessionModePNever() *PersistentBrowserSessionMode {
	v := PersistentBrowserSessionModeVNever
	return &v
}

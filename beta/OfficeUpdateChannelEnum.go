// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OfficeUpdateChannel undocumented
type OfficeUpdateChannel string

const (
	// OfficeUpdateChannelVNone undocumented
	OfficeUpdateChannelVNone OfficeUpdateChannel = "None"
	// OfficeUpdateChannelVCurrent undocumented
	OfficeUpdateChannelVCurrent OfficeUpdateChannel = "Current"
	// OfficeUpdateChannelVDeferred undocumented
	OfficeUpdateChannelVDeferred OfficeUpdateChannel = "Deferred"
	// OfficeUpdateChannelVFirstReleaseCurrent undocumented
	OfficeUpdateChannelVFirstReleaseCurrent OfficeUpdateChannel = "FirstReleaseCurrent"
	// OfficeUpdateChannelVFirstReleaseDeferred undocumented
	OfficeUpdateChannelVFirstReleaseDeferred OfficeUpdateChannel = "FirstReleaseDeferred"
)

// OfficeUpdateChannelPNone returns a pointer to OfficeUpdateChannelVNone
func OfficeUpdateChannelPNone() *OfficeUpdateChannel {
	v := OfficeUpdateChannelVNone
	return &v
}

// OfficeUpdateChannelPCurrent returns a pointer to OfficeUpdateChannelVCurrent
func OfficeUpdateChannelPCurrent() *OfficeUpdateChannel {
	v := OfficeUpdateChannelVCurrent
	return &v
}

// OfficeUpdateChannelPDeferred returns a pointer to OfficeUpdateChannelVDeferred
func OfficeUpdateChannelPDeferred() *OfficeUpdateChannel {
	v := OfficeUpdateChannelVDeferred
	return &v
}

// OfficeUpdateChannelPFirstReleaseCurrent returns a pointer to OfficeUpdateChannelVFirstReleaseCurrent
func OfficeUpdateChannelPFirstReleaseCurrent() *OfficeUpdateChannel {
	v := OfficeUpdateChannelVFirstReleaseCurrent
	return &v
}

// OfficeUpdateChannelPFirstReleaseDeferred returns a pointer to OfficeUpdateChannelVFirstReleaseDeferred
func OfficeUpdateChannelPFirstReleaseDeferred() *OfficeUpdateChannel {
	v := OfficeUpdateChannelVFirstReleaseDeferred
	return &v
}

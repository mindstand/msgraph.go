// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MicrosoftEdgeChannel undocumented
type MicrosoftEdgeChannel string

const (
	// MicrosoftEdgeChannelVDev undocumented
	MicrosoftEdgeChannelVDev MicrosoftEdgeChannel = "Dev"
	// MicrosoftEdgeChannelVBeta undocumented
	MicrosoftEdgeChannelVBeta MicrosoftEdgeChannel = "Beta"
	// MicrosoftEdgeChannelVStable undocumented
	MicrosoftEdgeChannelVStable MicrosoftEdgeChannel = "Stable"
)

// MicrosoftEdgeChannelPDev returns a pointer to MicrosoftEdgeChannelVDev
func MicrosoftEdgeChannelPDev() *MicrosoftEdgeChannel {
	v := MicrosoftEdgeChannelVDev
	return &v
}

// MicrosoftEdgeChannelPBeta returns a pointer to MicrosoftEdgeChannelVBeta
func MicrosoftEdgeChannelPBeta() *MicrosoftEdgeChannel {
	v := MicrosoftEdgeChannelVBeta
	return &v
}

// MicrosoftEdgeChannelPStable returns a pointer to MicrosoftEdgeChannelVStable
func MicrosoftEdgeChannelPStable() *MicrosoftEdgeChannel {
	v := MicrosoftEdgeChannelVStable
	return &v
}
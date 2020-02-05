// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ChannelMembershipType undocumented
type ChannelMembershipType string

const (
	// ChannelMembershipTypeVStandard undocumented
	ChannelMembershipTypeVStandard ChannelMembershipType = "Standard"
	// ChannelMembershipTypeVPrivate undocumented
	ChannelMembershipTypeVPrivate ChannelMembershipType = "Private"
	// ChannelMembershipTypeVUnknownFutureValue undocumented
	ChannelMembershipTypeVUnknownFutureValue ChannelMembershipType = "UnknownFutureValue"
)

// ChannelMembershipTypePStandard returns a pointer to ChannelMembershipTypeVStandard
func ChannelMembershipTypePStandard() *ChannelMembershipType {
	v := ChannelMembershipTypeVStandard
	return &v
}

// ChannelMembershipTypePPrivate returns a pointer to ChannelMembershipTypeVPrivate
func ChannelMembershipTypePPrivate() *ChannelMembershipType {
	v := ChannelMembershipTypeVPrivate
	return &v
}

// ChannelMembershipTypePUnknownFutureValue returns a pointer to ChannelMembershipTypeVUnknownFutureValue
func ChannelMembershipTypePUnknownFutureValue() *ChannelMembershipType {
	v := ChannelMembershipTypeVUnknownFutureValue
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FirewallPacketQueueingMethodType undocumented
type FirewallPacketQueueingMethodType int

const (
	// FirewallPacketQueueingMethodTypeVDeviceDefault undocumented
	FirewallPacketQueueingMethodTypeVDeviceDefault FirewallPacketQueueingMethodType = 0
	// FirewallPacketQueueingMethodTypeVDisabled undocumented
	FirewallPacketQueueingMethodTypeVDisabled FirewallPacketQueueingMethodType = 1
	// FirewallPacketQueueingMethodTypeVQueueInbound undocumented
	FirewallPacketQueueingMethodTypeVQueueInbound FirewallPacketQueueingMethodType = 2
	// FirewallPacketQueueingMethodTypeVQueueOutbound undocumented
	FirewallPacketQueueingMethodTypeVQueueOutbound FirewallPacketQueueingMethodType = 3
	// FirewallPacketQueueingMethodTypeVQueueBoth undocumented
	FirewallPacketQueueingMethodTypeVQueueBoth FirewallPacketQueueingMethodType = 4
)

// FirewallPacketQueueingMethodTypePDeviceDefault returns a pointer to FirewallPacketQueueingMethodTypeVDeviceDefault
func FirewallPacketQueueingMethodTypePDeviceDefault() *FirewallPacketQueueingMethodType {
	v := FirewallPacketQueueingMethodTypeVDeviceDefault
	return &v
}

// FirewallPacketQueueingMethodTypePDisabled returns a pointer to FirewallPacketQueueingMethodTypeVDisabled
func FirewallPacketQueueingMethodTypePDisabled() *FirewallPacketQueueingMethodType {
	v := FirewallPacketQueueingMethodTypeVDisabled
	return &v
}

// FirewallPacketQueueingMethodTypePQueueInbound returns a pointer to FirewallPacketQueueingMethodTypeVQueueInbound
func FirewallPacketQueueingMethodTypePQueueInbound() *FirewallPacketQueueingMethodType {
	v := FirewallPacketQueueingMethodTypeVQueueInbound
	return &v
}

// FirewallPacketQueueingMethodTypePQueueOutbound returns a pointer to FirewallPacketQueueingMethodTypeVQueueOutbound
func FirewallPacketQueueingMethodTypePQueueOutbound() *FirewallPacketQueueingMethodType {
	v := FirewallPacketQueueingMethodTypeVQueueOutbound
	return &v
}

// FirewallPacketQueueingMethodTypePQueueBoth returns a pointer to FirewallPacketQueueingMethodTypeVQueueBoth
func FirewallPacketQueueingMethodTypePQueueBoth() *FirewallPacketQueueingMethodType {
	v := FirewallPacketQueueingMethodTypeVQueueBoth
	return &v
}

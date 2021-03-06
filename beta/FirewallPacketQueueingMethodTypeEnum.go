// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// FirewallPacketQueueingMethodType undocumented
type FirewallPacketQueueingMethodType string

const (
	// FirewallPacketQueueingMethodTypeVDeviceDefault undocumented
	FirewallPacketQueueingMethodTypeVDeviceDefault FirewallPacketQueueingMethodType = "DeviceDefault"
	// FirewallPacketQueueingMethodTypeVDisabled undocumented
	FirewallPacketQueueingMethodTypeVDisabled FirewallPacketQueueingMethodType = "Disabled"
	// FirewallPacketQueueingMethodTypeVQueueInbound undocumented
	FirewallPacketQueueingMethodTypeVQueueInbound FirewallPacketQueueingMethodType = "QueueInbound"
	// FirewallPacketQueueingMethodTypeVQueueOutbound undocumented
	FirewallPacketQueueingMethodTypeVQueueOutbound FirewallPacketQueueingMethodType = "QueueOutbound"
	// FirewallPacketQueueingMethodTypeVQueueBoth undocumented
	FirewallPacketQueueingMethodTypeVQueueBoth FirewallPacketQueueingMethodType = "QueueBoth"
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

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RoutingMode undocumented
type RoutingMode string

const (
	// RoutingModeVOneToOne undocumented
	RoutingModeVOneToOne RoutingMode = "OneToOne"
	// RoutingModeVMulticast undocumented
	RoutingModeVMulticast RoutingMode = "Multicast"
)

// RoutingModePOneToOne returns a pointer to RoutingModeVOneToOne
func RoutingModePOneToOne() *RoutingMode {
	v := RoutingModeVOneToOne
	return &v
}

// RoutingModePMulticast returns a pointer to RoutingModeVMulticast
func RoutingModePMulticast() *RoutingMode {
	v := RoutingModeVMulticast
	return &v
}

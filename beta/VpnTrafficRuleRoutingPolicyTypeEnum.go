// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VpnTrafficRuleRoutingPolicyType undocumented
type VpnTrafficRuleRoutingPolicyType string

const (
	// VpnTrafficRuleRoutingPolicyTypeVNone undocumented
	VpnTrafficRuleRoutingPolicyTypeVNone VpnTrafficRuleRoutingPolicyType = "None"
	// VpnTrafficRuleRoutingPolicyTypeVSplitTunnel undocumented
	VpnTrafficRuleRoutingPolicyTypeVSplitTunnel VpnTrafficRuleRoutingPolicyType = "SplitTunnel"
	// VpnTrafficRuleRoutingPolicyTypeVForceTunnel undocumented
	VpnTrafficRuleRoutingPolicyTypeVForceTunnel VpnTrafficRuleRoutingPolicyType = "ForceTunnel"
)

// VpnTrafficRuleRoutingPolicyTypePNone returns a pointer to VpnTrafficRuleRoutingPolicyTypeVNone
func VpnTrafficRuleRoutingPolicyTypePNone() *VpnTrafficRuleRoutingPolicyType {
	v := VpnTrafficRuleRoutingPolicyTypeVNone
	return &v
}

// VpnTrafficRuleRoutingPolicyTypePSplitTunnel returns a pointer to VpnTrafficRuleRoutingPolicyTypeVSplitTunnel
func VpnTrafficRuleRoutingPolicyTypePSplitTunnel() *VpnTrafficRuleRoutingPolicyType {
	v := VpnTrafficRuleRoutingPolicyTypeVSplitTunnel
	return &v
}

// VpnTrafficRuleRoutingPolicyTypePForceTunnel returns a pointer to VpnTrafficRuleRoutingPolicyTypeVForceTunnel
func VpnTrafficRuleRoutingPolicyTypePForceTunnel() *VpnTrafficRuleRoutingPolicyType {
	v := VpnTrafficRuleRoutingPolicyTypeVForceTunnel
	return &v
}

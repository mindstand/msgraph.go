// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// VpnDeadPeerDetectionRate undocumented
type VpnDeadPeerDetectionRate string

const (
	// VpnDeadPeerDetectionRateVMedium undocumented
	VpnDeadPeerDetectionRateVMedium VpnDeadPeerDetectionRate = "Medium"
	// VpnDeadPeerDetectionRateVNone undocumented
	VpnDeadPeerDetectionRateVNone VpnDeadPeerDetectionRate = "None"
	// VpnDeadPeerDetectionRateVLow undocumented
	VpnDeadPeerDetectionRateVLow VpnDeadPeerDetectionRate = "Low"
	// VpnDeadPeerDetectionRateVHigh undocumented
	VpnDeadPeerDetectionRateVHigh VpnDeadPeerDetectionRate = "High"
)

// VpnDeadPeerDetectionRatePMedium returns a pointer to VpnDeadPeerDetectionRateVMedium
func VpnDeadPeerDetectionRatePMedium() *VpnDeadPeerDetectionRate {
	v := VpnDeadPeerDetectionRateVMedium
	return &v
}

// VpnDeadPeerDetectionRatePNone returns a pointer to VpnDeadPeerDetectionRateVNone
func VpnDeadPeerDetectionRatePNone() *VpnDeadPeerDetectionRate {
	v := VpnDeadPeerDetectionRateVNone
	return &v
}

// VpnDeadPeerDetectionRatePLow returns a pointer to VpnDeadPeerDetectionRateVLow
func VpnDeadPeerDetectionRatePLow() *VpnDeadPeerDetectionRate {
	v := VpnDeadPeerDetectionRateVLow
	return &v
}

// VpnDeadPeerDetectionRatePHigh returns a pointer to VpnDeadPeerDetectionRateVHigh
func VpnDeadPeerDetectionRatePHigh() *VpnDeadPeerDetectionRate {
	v := VpnDeadPeerDetectionRateVHigh
	return &v
}

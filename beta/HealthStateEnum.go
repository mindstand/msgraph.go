// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// HealthState undocumented
type HealthState int

const (
	// HealthStateVUnknown undocumented
	HealthStateVUnknown HealthState = 0
	// HealthStateVHealthy undocumented
	HealthStateVHealthy HealthState = 1
	// HealthStateVUnhealthy undocumented
	HealthStateVUnhealthy HealthState = 2
)

// HealthStatePUnknown returns a pointer to HealthStateVUnknown
func HealthStatePUnknown() *HealthState {
	v := HealthStateVUnknown
	return &v
}

// HealthStatePHealthy returns a pointer to HealthStateVHealthy
func HealthStatePHealthy() *HealthState {
	v := HealthStateVHealthy
	return &v
}

// HealthStatePUnhealthy returns a pointer to HealthStateVUnhealthy
func HealthStatePUnhealthy() *HealthState {
	v := HealthStateVUnhealthy
	return &v
}

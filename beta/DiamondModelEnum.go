// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DiamondModel undocumented
type DiamondModel string

const (
	// DiamondModelVUnknown undocumented
	DiamondModelVUnknown DiamondModel = "Unknown"
	// DiamondModelVAdversary undocumented
	DiamondModelVAdversary DiamondModel = "Adversary"
	// DiamondModelVCapability undocumented
	DiamondModelVCapability DiamondModel = "Capability"
	// DiamondModelVInfrastructure undocumented
	DiamondModelVInfrastructure DiamondModel = "Infrastructure"
	// DiamondModelVVictim undocumented
	DiamondModelVVictim DiamondModel = "Victim"
	// DiamondModelVUnknownFutureValue undocumented
	DiamondModelVUnknownFutureValue DiamondModel = "UnknownFutureValue"
)

// DiamondModelPUnknown returns a pointer to DiamondModelVUnknown
func DiamondModelPUnknown() *DiamondModel {
	v := DiamondModelVUnknown
	return &v
}

// DiamondModelPAdversary returns a pointer to DiamondModelVAdversary
func DiamondModelPAdversary() *DiamondModel {
	v := DiamondModelVAdversary
	return &v
}

// DiamondModelPCapability returns a pointer to DiamondModelVCapability
func DiamondModelPCapability() *DiamondModel {
	v := DiamondModelVCapability
	return &v
}

// DiamondModelPInfrastructure returns a pointer to DiamondModelVInfrastructure
func DiamondModelPInfrastructure() *DiamondModel {
	v := DiamondModelVInfrastructure
	return &v
}

// DiamondModelPVictim returns a pointer to DiamondModelVVictim
func DiamondModelPVictim() *DiamondModel {
	v := DiamondModelVVictim
	return &v
}

// DiamondModelPUnknownFutureValue returns a pointer to DiamondModelVUnknownFutureValue
func DiamondModelPUnknownFutureValue() *DiamondModel {
	v := DiamondModelVUnknownFutureValue
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AdministratorConfiguredDeviceComplianceState undocumented
type AdministratorConfiguredDeviceComplianceState string

const (
	// AdministratorConfiguredDeviceComplianceStateVBasedOnDeviceCompliancePolicy undocumented
	AdministratorConfiguredDeviceComplianceStateVBasedOnDeviceCompliancePolicy AdministratorConfiguredDeviceComplianceState = "BasedOnDeviceCompliancePolicy"
	// AdministratorConfiguredDeviceComplianceStateVNonCompliant undocumented
	AdministratorConfiguredDeviceComplianceStateVNonCompliant AdministratorConfiguredDeviceComplianceState = "NonCompliant"
)

// AdministratorConfiguredDeviceComplianceStatePBasedOnDeviceCompliancePolicy returns a pointer to AdministratorConfiguredDeviceComplianceStateVBasedOnDeviceCompliancePolicy
func AdministratorConfiguredDeviceComplianceStatePBasedOnDeviceCompliancePolicy() *AdministratorConfiguredDeviceComplianceState {
	v := AdministratorConfiguredDeviceComplianceStateVBasedOnDeviceCompliancePolicy
	return &v
}

// AdministratorConfiguredDeviceComplianceStatePNonCompliant returns a pointer to AdministratorConfiguredDeviceComplianceStateVNonCompliant
func AdministratorConfiguredDeviceComplianceStatePNonCompliant() *AdministratorConfiguredDeviceComplianceState {
	v := AdministratorConfiguredDeviceComplianceStateVNonCompliant
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BitLockerRecoveryInformationType undocumented
type BitLockerRecoveryInformationType string

const (
	// BitLockerRecoveryInformationTypeVPasswordAndKey undocumented
	BitLockerRecoveryInformationTypeVPasswordAndKey BitLockerRecoveryInformationType = "PasswordAndKey"
	// BitLockerRecoveryInformationTypeVPasswordOnly undocumented
	BitLockerRecoveryInformationTypeVPasswordOnly BitLockerRecoveryInformationType = "PasswordOnly"
)

// BitLockerRecoveryInformationTypePPasswordAndKey returns a pointer to BitLockerRecoveryInformationTypeVPasswordAndKey
func BitLockerRecoveryInformationTypePPasswordAndKey() *BitLockerRecoveryInformationType {
	v := BitLockerRecoveryInformationTypeVPasswordAndKey
	return &v
}

// BitLockerRecoveryInformationTypePPasswordOnly returns a pointer to BitLockerRecoveryInformationTypeVPasswordOnly
func BitLockerRecoveryInformationTypePPasswordOnly() *BitLockerRecoveryInformationType {
	v := BitLockerRecoveryInformationTypeVPasswordOnly
	return &v
}

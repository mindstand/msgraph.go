// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RiskDetail undocumented
type RiskDetail string

const (
	// RiskDetailVNone undocumented
	RiskDetailVNone RiskDetail = "None"
	// RiskDetailVAdminGeneratedTemporaryPassword undocumented
	RiskDetailVAdminGeneratedTemporaryPassword RiskDetail = "AdminGeneratedTemporaryPassword"
	// RiskDetailVUserPerformedSecuredPasswordChange undocumented
	RiskDetailVUserPerformedSecuredPasswordChange RiskDetail = "UserPerformedSecuredPasswordChange"
	// RiskDetailVUserPerformedSecuredPasswordReset undocumented
	RiskDetailVUserPerformedSecuredPasswordReset RiskDetail = "UserPerformedSecuredPasswordReset"
	// RiskDetailVAdminConfirmedSigninSafe undocumented
	RiskDetailVAdminConfirmedSigninSafe RiskDetail = "AdminConfirmedSigninSafe"
	// RiskDetailVAiConfirmedSigninSafe undocumented
	RiskDetailVAiConfirmedSigninSafe RiskDetail = "AiConfirmedSigninSafe"
	// RiskDetailVUserPassedMFADrivenByRiskBasedPolicy undocumented
	RiskDetailVUserPassedMFADrivenByRiskBasedPolicy RiskDetail = "UserPassedMFADrivenByRiskBasedPolicy"
	// RiskDetailVAdminDismissedAllRiskForUser undocumented
	RiskDetailVAdminDismissedAllRiskForUser RiskDetail = "AdminDismissedAllRiskForUser"
	// RiskDetailVAdminConfirmedSigninCompromised undocumented
	RiskDetailVAdminConfirmedSigninCompromised RiskDetail = "AdminConfirmedSigninCompromised"
	// RiskDetailVHidden undocumented
	RiskDetailVHidden RiskDetail = "Hidden"
	// RiskDetailVAdminConfirmedUserCompromised undocumented
	RiskDetailVAdminConfirmedUserCompromised RiskDetail = "AdminConfirmedUserCompromised"
	// RiskDetailVUnknownFutureValue undocumented
	RiskDetailVUnknownFutureValue RiskDetail = "UnknownFutureValue"
)

// RiskDetailPNone returns a pointer to RiskDetailVNone
func RiskDetailPNone() *RiskDetail {
	v := RiskDetailVNone
	return &v
}

// RiskDetailPAdminGeneratedTemporaryPassword returns a pointer to RiskDetailVAdminGeneratedTemporaryPassword
func RiskDetailPAdminGeneratedTemporaryPassword() *RiskDetail {
	v := RiskDetailVAdminGeneratedTemporaryPassword
	return &v
}

// RiskDetailPUserPerformedSecuredPasswordChange returns a pointer to RiskDetailVUserPerformedSecuredPasswordChange
func RiskDetailPUserPerformedSecuredPasswordChange() *RiskDetail {
	v := RiskDetailVUserPerformedSecuredPasswordChange
	return &v
}

// RiskDetailPUserPerformedSecuredPasswordReset returns a pointer to RiskDetailVUserPerformedSecuredPasswordReset
func RiskDetailPUserPerformedSecuredPasswordReset() *RiskDetail {
	v := RiskDetailVUserPerformedSecuredPasswordReset
	return &v
}

// RiskDetailPAdminConfirmedSigninSafe returns a pointer to RiskDetailVAdminConfirmedSigninSafe
func RiskDetailPAdminConfirmedSigninSafe() *RiskDetail {
	v := RiskDetailVAdminConfirmedSigninSafe
	return &v
}

// RiskDetailPAiConfirmedSigninSafe returns a pointer to RiskDetailVAiConfirmedSigninSafe
func RiskDetailPAiConfirmedSigninSafe() *RiskDetail {
	v := RiskDetailVAiConfirmedSigninSafe
	return &v
}

// RiskDetailPUserPassedMFADrivenByRiskBasedPolicy returns a pointer to RiskDetailVUserPassedMFADrivenByRiskBasedPolicy
func RiskDetailPUserPassedMFADrivenByRiskBasedPolicy() *RiskDetail {
	v := RiskDetailVUserPassedMFADrivenByRiskBasedPolicy
	return &v
}

// RiskDetailPAdminDismissedAllRiskForUser returns a pointer to RiskDetailVAdminDismissedAllRiskForUser
func RiskDetailPAdminDismissedAllRiskForUser() *RiskDetail {
	v := RiskDetailVAdminDismissedAllRiskForUser
	return &v
}

// RiskDetailPAdminConfirmedSigninCompromised returns a pointer to RiskDetailVAdminConfirmedSigninCompromised
func RiskDetailPAdminConfirmedSigninCompromised() *RiskDetail {
	v := RiskDetailVAdminConfirmedSigninCompromised
	return &v
}

// RiskDetailPHidden returns a pointer to RiskDetailVHidden
func RiskDetailPHidden() *RiskDetail {
	v := RiskDetailVHidden
	return &v
}

// RiskDetailPAdminConfirmedUserCompromised returns a pointer to RiskDetailVAdminConfirmedUserCompromised
func RiskDetailPAdminConfirmedUserCompromised() *RiskDetail {
	v := RiskDetailVAdminConfirmedUserCompromised
	return &v
}

// RiskDetailPUnknownFutureValue returns a pointer to RiskDetailVUnknownFutureValue
func RiskDetailPUnknownFutureValue() *RiskDetail {
	v := RiskDetailVUnknownFutureValue
	return &v
}

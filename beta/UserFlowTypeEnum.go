// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserFlowType undocumented
type UserFlowType string

const (
	// UserFlowTypeVSignUp undocumented
	UserFlowTypeVSignUp UserFlowType = "SignUp"
	// UserFlowTypeVSignIn undocumented
	UserFlowTypeVSignIn UserFlowType = "SignIn"
	// UserFlowTypeVSignUpOrSignIn undocumented
	UserFlowTypeVSignUpOrSignIn UserFlowType = "SignUpOrSignIn"
	// UserFlowTypeVPasswordReset undocumented
	UserFlowTypeVPasswordReset UserFlowType = "PasswordReset"
	// UserFlowTypeVProfileUpdate undocumented
	UserFlowTypeVProfileUpdate UserFlowType = "ProfileUpdate"
	// UserFlowTypeVResourceOwner undocumented
	UserFlowTypeVResourceOwner UserFlowType = "ResourceOwner"
	// UserFlowTypeVUnknownFutureValue undocumented
	UserFlowTypeVUnknownFutureValue UserFlowType = "UnknownFutureValue"
)

// UserFlowTypePSignUp returns a pointer to UserFlowTypeVSignUp
func UserFlowTypePSignUp() *UserFlowType {
	v := UserFlowTypeVSignUp
	return &v
}

// UserFlowTypePSignIn returns a pointer to UserFlowTypeVSignIn
func UserFlowTypePSignIn() *UserFlowType {
	v := UserFlowTypeVSignIn
	return &v
}

// UserFlowTypePSignUpOrSignIn returns a pointer to UserFlowTypeVSignUpOrSignIn
func UserFlowTypePSignUpOrSignIn() *UserFlowType {
	v := UserFlowTypeVSignUpOrSignIn
	return &v
}

// UserFlowTypePPasswordReset returns a pointer to UserFlowTypeVPasswordReset
func UserFlowTypePPasswordReset() *UserFlowType {
	v := UserFlowTypeVPasswordReset
	return &v
}

// UserFlowTypePProfileUpdate returns a pointer to UserFlowTypeVProfileUpdate
func UserFlowTypePProfileUpdate() *UserFlowType {
	v := UserFlowTypeVProfileUpdate
	return &v
}

// UserFlowTypePResourceOwner returns a pointer to UserFlowTypeVResourceOwner
func UserFlowTypePResourceOwner() *UserFlowType {
	v := UserFlowTypeVResourceOwner
	return &v
}

// UserFlowTypePUnknownFutureValue returns a pointer to UserFlowTypeVUnknownFutureValue
func UserFlowTypePUnknownFutureValue() *UserFlowType {
	v := UserFlowTypeVUnknownFutureValue
	return &v
}

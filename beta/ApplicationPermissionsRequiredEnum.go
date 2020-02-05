// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ApplicationPermissionsRequired undocumented
type ApplicationPermissionsRequired string

const (
	// ApplicationPermissionsRequiredVUnknown undocumented
	ApplicationPermissionsRequiredVUnknown ApplicationPermissionsRequired = "Unknown"
	// ApplicationPermissionsRequiredVAnonymous undocumented
	ApplicationPermissionsRequiredVAnonymous ApplicationPermissionsRequired = "Anonymous"
	// ApplicationPermissionsRequiredVGuest undocumented
	ApplicationPermissionsRequiredVGuest ApplicationPermissionsRequired = "Guest"
	// ApplicationPermissionsRequiredVUser undocumented
	ApplicationPermissionsRequiredVUser ApplicationPermissionsRequired = "User"
	// ApplicationPermissionsRequiredVAdministrator undocumented
	ApplicationPermissionsRequiredVAdministrator ApplicationPermissionsRequired = "Administrator"
	// ApplicationPermissionsRequiredVSystem undocumented
	ApplicationPermissionsRequiredVSystem ApplicationPermissionsRequired = "System"
	// ApplicationPermissionsRequiredVUnknownFutureValue undocumented
	ApplicationPermissionsRequiredVUnknownFutureValue ApplicationPermissionsRequired = "UnknownFutureValue"
)

// ApplicationPermissionsRequiredPUnknown returns a pointer to ApplicationPermissionsRequiredVUnknown
func ApplicationPermissionsRequiredPUnknown() *ApplicationPermissionsRequired {
	v := ApplicationPermissionsRequiredVUnknown
	return &v
}

// ApplicationPermissionsRequiredPAnonymous returns a pointer to ApplicationPermissionsRequiredVAnonymous
func ApplicationPermissionsRequiredPAnonymous() *ApplicationPermissionsRequired {
	v := ApplicationPermissionsRequiredVAnonymous
	return &v
}

// ApplicationPermissionsRequiredPGuest returns a pointer to ApplicationPermissionsRequiredVGuest
func ApplicationPermissionsRequiredPGuest() *ApplicationPermissionsRequired {
	v := ApplicationPermissionsRequiredVGuest
	return &v
}

// ApplicationPermissionsRequiredPUser returns a pointer to ApplicationPermissionsRequiredVUser
func ApplicationPermissionsRequiredPUser() *ApplicationPermissionsRequired {
	v := ApplicationPermissionsRequiredVUser
	return &v
}

// ApplicationPermissionsRequiredPAdministrator returns a pointer to ApplicationPermissionsRequiredVAdministrator
func ApplicationPermissionsRequiredPAdministrator() *ApplicationPermissionsRequired {
	v := ApplicationPermissionsRequiredVAdministrator
	return &v
}

// ApplicationPermissionsRequiredPSystem returns a pointer to ApplicationPermissionsRequiredVSystem
func ApplicationPermissionsRequiredPSystem() *ApplicationPermissionsRequired {
	v := ApplicationPermissionsRequiredVSystem
	return &v
}

// ApplicationPermissionsRequiredPUnknownFutureValue returns a pointer to ApplicationPermissionsRequiredVUnknownFutureValue
func ApplicationPermissionsRequiredPUnknownFutureValue() *ApplicationPermissionsRequired {
	v := ApplicationPermissionsRequiredVUnknownFutureValue
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ContactRelationship undocumented
type ContactRelationship string

const (
	// ContactRelationshipVParent undocumented
	ContactRelationshipVParent ContactRelationship = "Parent"
	// ContactRelationshipVRelative undocumented
	ContactRelationshipVRelative ContactRelationship = "Relative"
	// ContactRelationshipVAide undocumented
	ContactRelationshipVAide ContactRelationship = "Aide"
	// ContactRelationshipVDoctor undocumented
	ContactRelationshipVDoctor ContactRelationship = "Doctor"
	// ContactRelationshipVGuardian undocumented
	ContactRelationshipVGuardian ContactRelationship = "Guardian"
	// ContactRelationshipVChild undocumented
	ContactRelationshipVChild ContactRelationship = "Child"
	// ContactRelationshipVOther undocumented
	ContactRelationshipVOther ContactRelationship = "Other"
	// ContactRelationshipVUnknownFutureValue undocumented
	ContactRelationshipVUnknownFutureValue ContactRelationship = "UnknownFutureValue"
)

// ContactRelationshipPParent returns a pointer to ContactRelationshipVParent
func ContactRelationshipPParent() *ContactRelationship {
	v := ContactRelationshipVParent
	return &v
}

// ContactRelationshipPRelative returns a pointer to ContactRelationshipVRelative
func ContactRelationshipPRelative() *ContactRelationship {
	v := ContactRelationshipVRelative
	return &v
}

// ContactRelationshipPAide returns a pointer to ContactRelationshipVAide
func ContactRelationshipPAide() *ContactRelationship {
	v := ContactRelationshipVAide
	return &v
}

// ContactRelationshipPDoctor returns a pointer to ContactRelationshipVDoctor
func ContactRelationshipPDoctor() *ContactRelationship {
	v := ContactRelationshipVDoctor
	return &v
}

// ContactRelationshipPGuardian returns a pointer to ContactRelationshipVGuardian
func ContactRelationshipPGuardian() *ContactRelationship {
	v := ContactRelationshipVGuardian
	return &v
}

// ContactRelationshipPChild returns a pointer to ContactRelationshipVChild
func ContactRelationshipPChild() *ContactRelationship {
	v := ContactRelationshipVChild
	return &v
}

// ContactRelationshipPOther returns a pointer to ContactRelationshipVOther
func ContactRelationshipPOther() *ContactRelationship {
	v := ContactRelationshipVOther
	return &v
}

// ContactRelationshipPUnknownFutureValue returns a pointer to ContactRelationshipVUnknownFutureValue
func ContactRelationshipPUnknownFutureValue() *ContactRelationship {
	v := ContactRelationshipVUnknownFutureValue
	return &v
}

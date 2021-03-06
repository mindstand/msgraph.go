// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ACLType undocumented
type ACLType string

const (
	// ACLTypeVUser undocumented
	ACLTypeVUser ACLType = "User"
	// ACLTypeVGroup undocumented
	ACLTypeVGroup ACLType = "Group"
	// ACLTypeVEveryone undocumented
	ACLTypeVEveryone ACLType = "Everyone"
	// ACLTypeVEveryoneExceptGuests undocumented
	ACLTypeVEveryoneExceptGuests ACLType = "EveryoneExceptGuests"
)

// ACLTypePUser returns a pointer to ACLTypeVUser
func ACLTypePUser() *ACLType {
	v := ACLTypeVUser
	return &v
}

// ACLTypePGroup returns a pointer to ACLTypeVGroup
func ACLTypePGroup() *ACLType {
	v := ACLTypeVGroup
	return &v
}

// ACLTypePEveryone returns a pointer to ACLTypeVEveryone
func ACLTypePEveryone() *ACLType {
	v := ACLTypeVEveryone
	return &v
}

// ACLTypePEveryoneExceptGuests returns a pointer to ACLTypeVEveryoneExceptGuests
func ACLTypePEveryoneExceptGuests() *ACLType {
	v := ACLTypeVEveryoneExceptGuests
	return &v
}

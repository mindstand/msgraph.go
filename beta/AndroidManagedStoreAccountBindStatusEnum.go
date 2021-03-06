// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidManagedStoreAccountBindStatus undocumented
type AndroidManagedStoreAccountBindStatus string

const (
	// AndroidManagedStoreAccountBindStatusVNotBound undocumented
	AndroidManagedStoreAccountBindStatusVNotBound AndroidManagedStoreAccountBindStatus = "NotBound"
	// AndroidManagedStoreAccountBindStatusVBound undocumented
	AndroidManagedStoreAccountBindStatusVBound AndroidManagedStoreAccountBindStatus = "Bound"
	// AndroidManagedStoreAccountBindStatusVBoundAndValidated undocumented
	AndroidManagedStoreAccountBindStatusVBoundAndValidated AndroidManagedStoreAccountBindStatus = "BoundAndValidated"
	// AndroidManagedStoreAccountBindStatusVUnbinding undocumented
	AndroidManagedStoreAccountBindStatusVUnbinding AndroidManagedStoreAccountBindStatus = "Unbinding"
)

// AndroidManagedStoreAccountBindStatusPNotBound returns a pointer to AndroidManagedStoreAccountBindStatusVNotBound
func AndroidManagedStoreAccountBindStatusPNotBound() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVNotBound
	return &v
}

// AndroidManagedStoreAccountBindStatusPBound returns a pointer to AndroidManagedStoreAccountBindStatusVBound
func AndroidManagedStoreAccountBindStatusPBound() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVBound
	return &v
}

// AndroidManagedStoreAccountBindStatusPBoundAndValidated returns a pointer to AndroidManagedStoreAccountBindStatusVBoundAndValidated
func AndroidManagedStoreAccountBindStatusPBoundAndValidated() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVBoundAndValidated
	return &v
}

// AndroidManagedStoreAccountBindStatusPUnbinding returns a pointer to AndroidManagedStoreAccountBindStatusVUnbinding
func AndroidManagedStoreAccountBindStatusPUnbinding() *AndroidManagedStoreAccountBindStatus {
	v := AndroidManagedStoreAccountBindStatusVUnbinding
	return &v
}

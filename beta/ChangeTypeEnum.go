// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ChangeType undocumented
type ChangeType string

const (
	// ChangeTypeVCreated undocumented
	ChangeTypeVCreated ChangeType = "Created"
	// ChangeTypeVUpdated undocumented
	ChangeTypeVUpdated ChangeType = "Updated"
	// ChangeTypeVDeleted undocumented
	ChangeTypeVDeleted ChangeType = "Deleted"
)

// ChangeTypePCreated returns a pointer to ChangeTypeVCreated
func ChangeTypePCreated() *ChangeType {
	v := ChangeTypeVCreated
	return &v
}

// ChangeTypePUpdated returns a pointer to ChangeTypeVUpdated
func ChangeTypePUpdated() *ChangeType {
	v := ChangeTypeVUpdated
	return &v
}

// ChangeTypePDeleted returns a pointer to ChangeTypeVDeleted
func ChangeTypePDeleted() *ChangeType {
	v := ChangeTypeVDeleted
	return &v
}

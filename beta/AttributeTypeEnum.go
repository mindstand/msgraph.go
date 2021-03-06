// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AttributeType undocumented
type AttributeType string

const (
	// AttributeTypeVDateTime undocumented
	AttributeTypeVDateTime AttributeType = "DateTime"
	// AttributeTypeVBoolean undocumented
	AttributeTypeVBoolean AttributeType = "Boolean"
	// AttributeTypeVBinary undocumented
	AttributeTypeVBinary AttributeType = "Binary"
	// AttributeTypeVReference undocumented
	AttributeTypeVReference AttributeType = "Reference"
	// AttributeTypeVInteger undocumented
	AttributeTypeVInteger AttributeType = "Integer"
	// AttributeTypeVString undocumented
	AttributeTypeVString AttributeType = "String"
)

// AttributeTypePDateTime returns a pointer to AttributeTypeVDateTime
func AttributeTypePDateTime() *AttributeType {
	v := AttributeTypeVDateTime
	return &v
}

// AttributeTypePBoolean returns a pointer to AttributeTypeVBoolean
func AttributeTypePBoolean() *AttributeType {
	v := AttributeTypeVBoolean
	return &v
}

// AttributeTypePBinary returns a pointer to AttributeTypeVBinary
func AttributeTypePBinary() *AttributeType {
	v := AttributeTypeVBinary
	return &v
}

// AttributeTypePReference returns a pointer to AttributeTypeVReference
func AttributeTypePReference() *AttributeType {
	v := AttributeTypeVReference
	return &v
}

// AttributeTypePInteger returns a pointer to AttributeTypeVInteger
func AttributeTypePInteger() *AttributeType {
	v := AttributeTypeVInteger
	return &v
}

// AttributeTypePString returns a pointer to AttributeTypeVString
func AttributeTypePString() *AttributeType {
	v := AttributeTypeVString
	return &v
}

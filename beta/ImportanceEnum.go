// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Importance undocumented
type Importance string

const (
	// ImportanceVLow undocumented
	ImportanceVLow Importance = "Low"
	// ImportanceVNormal undocumented
	ImportanceVNormal Importance = "Normal"
	// ImportanceVHigh undocumented
	ImportanceVHigh Importance = "High"
)

// ImportancePLow returns a pointer to ImportanceVLow
func ImportancePLow() *Importance {
	v := ImportanceVLow
	return &v
}

// ImportancePNormal returns a pointer to ImportanceVNormal
func ImportancePNormal() *Importance {
	v := ImportanceVNormal
	return &v
}

// ImportancePHigh returns a pointer to ImportanceVHigh
func ImportancePHigh() *Importance {
	v := ImportanceVHigh
	return &v
}

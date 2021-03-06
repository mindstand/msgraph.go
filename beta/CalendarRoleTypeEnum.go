// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CalendarRoleType undocumented
type CalendarRoleType string

const (
	// CalendarRoleTypeVNone undocumented
	CalendarRoleTypeVNone CalendarRoleType = "None"
	// CalendarRoleTypeVFreeBusyRead undocumented
	CalendarRoleTypeVFreeBusyRead CalendarRoleType = "FreeBusyRead"
	// CalendarRoleTypeVLimitedRead undocumented
	CalendarRoleTypeVLimitedRead CalendarRoleType = "LimitedRead"
	// CalendarRoleTypeVRead undocumented
	CalendarRoleTypeVRead CalendarRoleType = "Read"
	// CalendarRoleTypeVWrite undocumented
	CalendarRoleTypeVWrite CalendarRoleType = "Write"
	// CalendarRoleTypeVDelegateWithoutPrivateEventAccess undocumented
	CalendarRoleTypeVDelegateWithoutPrivateEventAccess CalendarRoleType = "DelegateWithoutPrivateEventAccess"
	// CalendarRoleTypeVDelegateWithPrivateEventAccess undocumented
	CalendarRoleTypeVDelegateWithPrivateEventAccess CalendarRoleType = "DelegateWithPrivateEventAccess"
	// CalendarRoleTypeVCustom undocumented
	CalendarRoleTypeVCustom CalendarRoleType = "Custom"
)

// CalendarRoleTypePNone returns a pointer to CalendarRoleTypeVNone
func CalendarRoleTypePNone() *CalendarRoleType {
	v := CalendarRoleTypeVNone
	return &v
}

// CalendarRoleTypePFreeBusyRead returns a pointer to CalendarRoleTypeVFreeBusyRead
func CalendarRoleTypePFreeBusyRead() *CalendarRoleType {
	v := CalendarRoleTypeVFreeBusyRead
	return &v
}

// CalendarRoleTypePLimitedRead returns a pointer to CalendarRoleTypeVLimitedRead
func CalendarRoleTypePLimitedRead() *CalendarRoleType {
	v := CalendarRoleTypeVLimitedRead
	return &v
}

// CalendarRoleTypePRead returns a pointer to CalendarRoleTypeVRead
func CalendarRoleTypePRead() *CalendarRoleType {
	v := CalendarRoleTypeVRead
	return &v
}

// CalendarRoleTypePWrite returns a pointer to CalendarRoleTypeVWrite
func CalendarRoleTypePWrite() *CalendarRoleType {
	v := CalendarRoleTypeVWrite
	return &v
}

// CalendarRoleTypePDelegateWithoutPrivateEventAccess returns a pointer to CalendarRoleTypeVDelegateWithoutPrivateEventAccess
func CalendarRoleTypePDelegateWithoutPrivateEventAccess() *CalendarRoleType {
	v := CalendarRoleTypeVDelegateWithoutPrivateEventAccess
	return &v
}

// CalendarRoleTypePDelegateWithPrivateEventAccess returns a pointer to CalendarRoleTypeVDelegateWithPrivateEventAccess
func CalendarRoleTypePDelegateWithPrivateEventAccess() *CalendarRoleType {
	v := CalendarRoleTypeVDelegateWithPrivateEventAccess
	return &v
}

// CalendarRoleTypePCustom returns a pointer to CalendarRoleTypeVCustom
func CalendarRoleTypePCustom() *CalendarRoleType {
	v := CalendarRoleTypeVCustom
	return &v
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CalendarSharingAction undocumented
type CalendarSharingAction string

const (
	// CalendarSharingActionVAccept undocumented
	CalendarSharingActionVAccept CalendarSharingAction = "Accept"
	// CalendarSharingActionVAcceptAndViewCalendar undocumented
	CalendarSharingActionVAcceptAndViewCalendar CalendarSharingAction = "AcceptAndViewCalendar"
	// CalendarSharingActionVViewCalendar undocumented
	CalendarSharingActionVViewCalendar CalendarSharingAction = "ViewCalendar"
	// CalendarSharingActionVAddThisCalendar undocumented
	CalendarSharingActionVAddThisCalendar CalendarSharingAction = "AddThisCalendar"
)

// CalendarSharingActionPAccept returns a pointer to CalendarSharingActionVAccept
func CalendarSharingActionPAccept() *CalendarSharingAction {
	v := CalendarSharingActionVAccept
	return &v
}

// CalendarSharingActionPAcceptAndViewCalendar returns a pointer to CalendarSharingActionVAcceptAndViewCalendar
func CalendarSharingActionPAcceptAndViewCalendar() *CalendarSharingAction {
	v := CalendarSharingActionVAcceptAndViewCalendar
	return &v
}

// CalendarSharingActionPViewCalendar returns a pointer to CalendarSharingActionVViewCalendar
func CalendarSharingActionPViewCalendar() *CalendarSharingAction {
	v := CalendarSharingActionVViewCalendar
	return &v
}

// CalendarSharingActionPAddThisCalendar returns a pointer to CalendarSharingActionVAddThisCalendar
func CalendarSharingActionPAddThisCalendar() *CalendarSharingAction {
	v := CalendarSharingActionVAddThisCalendar
	return &v
}
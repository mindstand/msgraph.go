// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserExperienceAnalyticsHealthState undocumented
type UserExperienceAnalyticsHealthState string

const (
	// UserExperienceAnalyticsHealthStateVUnknown undocumented
	UserExperienceAnalyticsHealthStateVUnknown UserExperienceAnalyticsHealthState = "Unknown"
	// UserExperienceAnalyticsHealthStateVInsufficientData undocumented
	UserExperienceAnalyticsHealthStateVInsufficientData UserExperienceAnalyticsHealthState = "InsufficientData"
	// UserExperienceAnalyticsHealthStateVNeedsAttention undocumented
	UserExperienceAnalyticsHealthStateVNeedsAttention UserExperienceAnalyticsHealthState = "NeedsAttention"
	// UserExperienceAnalyticsHealthStateVMeetingGoals undocumented
	UserExperienceAnalyticsHealthStateVMeetingGoals UserExperienceAnalyticsHealthState = "MeetingGoals"
)

// UserExperienceAnalyticsHealthStatePUnknown returns a pointer to UserExperienceAnalyticsHealthStateVUnknown
func UserExperienceAnalyticsHealthStatePUnknown() *UserExperienceAnalyticsHealthState {
	v := UserExperienceAnalyticsHealthStateVUnknown
	return &v
}

// UserExperienceAnalyticsHealthStatePInsufficientData returns a pointer to UserExperienceAnalyticsHealthStateVInsufficientData
func UserExperienceAnalyticsHealthStatePInsufficientData() *UserExperienceAnalyticsHealthState {
	v := UserExperienceAnalyticsHealthStateVInsufficientData
	return &v
}

// UserExperienceAnalyticsHealthStatePNeedsAttention returns a pointer to UserExperienceAnalyticsHealthStateVNeedsAttention
func UserExperienceAnalyticsHealthStatePNeedsAttention() *UserExperienceAnalyticsHealthState {
	v := UserExperienceAnalyticsHealthStateVNeedsAttention
	return &v
}

// UserExperienceAnalyticsHealthStatePMeetingGoals returns a pointer to UserExperienceAnalyticsHealthStateVMeetingGoals
func UserExperienceAnalyticsHealthStatePMeetingGoals() *UserExperienceAnalyticsHealthState {
	v := UserExperienceAnalyticsHealthStateVMeetingGoals
	return &v
}

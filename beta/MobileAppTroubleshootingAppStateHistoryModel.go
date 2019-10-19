// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppTroubleshootingAppStateHistory undocumented
type MobileAppTroubleshootingAppStateHistory struct {
	MobileAppTroubleshootingHistoryItem
	// ActionType AAD security group id to which it was targeted.
	ActionType *MobileAppActionType `json:"actionType,omitempty"`
	// RunState Status of the item.
	RunState *RunState `json:"runState,omitempty"`
	// ErrorCode Error code for the failure, empty if no failure.
	ErrorCode *string `json:"errorCode,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlannerAssignedToTaskBoardTaskFormat undocumented
type PlannerAssignedToTaskBoardTaskFormat struct {
	Entity
	// UnassignedOrderHint undocumented
	UnassignedOrderHint *string `json:"unassignedOrderHint,omitempty"`
	// OrderHintsByAssignee undocumented
	OrderHintsByAssignee *PlannerOrderHintsByAssignee `json:"orderHintsByAssignee,omitempty"`
}

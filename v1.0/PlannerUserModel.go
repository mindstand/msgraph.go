// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlannerUser undocumented
type PlannerUser struct {
	Entity
	// Tasks undocumented
	Tasks []PlannerTask `json:"tasks,omitempty"`
	// Plans undocumented
	Plans []PlannerPlan `json:"plans,omitempty"`
}

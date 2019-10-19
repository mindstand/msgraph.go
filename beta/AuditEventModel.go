// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// AuditEvent A class containing the properties for Audit Event.
type AuditEvent struct {
	Entity
	// DisplayName Event display name.
	DisplayName *string `json:"displayName,omitempty"`
	// ComponentName Component name.
	ComponentName *string `json:"componentName,omitempty"`
	// Actor AAD user and application that are associated with the audit event.
	Actor *AuditActor `json:"actor,omitempty"`
	// Activity Friendly name of the activity.
	Activity *string `json:"activity,omitempty"`
	// ActivityDateTime The date time in UTC when the activity was performed.
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// ActivityType The type of activity that was being performed.
	ActivityType *string `json:"activityType,omitempty"`
	// ActivityOperationType The HTTP operation type of the activity.
	ActivityOperationType *string `json:"activityOperationType,omitempty"`
	// ActivityResult The result of the activity.
	ActivityResult *string `json:"activityResult,omitempty"`
	// CorrelationID The client request Id that is used to correlate activity within the system.
	CorrelationID *UUID `json:"correlationId,omitempty"`
	// Resources Resources being modified.
	Resources []AuditResource `json:"resources,omitempty"`
	// Category Audit category.
	Category *string `json:"category,omitempty"`
}

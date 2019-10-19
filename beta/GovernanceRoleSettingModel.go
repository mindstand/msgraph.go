// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// GovernanceRoleSetting undocumented
type GovernanceRoleSetting struct {
	Entity
	// ResourceID undocumented
	ResourceID *string `json:"resourceId,omitempty"`
	// RoleDefinitionID undocumented
	RoleDefinitionID *string `json:"roleDefinitionId,omitempty"`
	// IsDefault undocumented
	IsDefault *bool `json:"isDefault,omitempty"`
	// LastUpdatedDateTime undocumented
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
	// LastUpdatedBy undocumented
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	// AdminEligibleSettings undocumented
	AdminEligibleSettings []GovernanceRuleSetting `json:"adminEligibleSettings,omitempty"`
	// AdminMemberSettings undocumented
	AdminMemberSettings []GovernanceRuleSetting `json:"adminMemberSettings,omitempty"`
	// UserEligibleSettings undocumented
	UserEligibleSettings []GovernanceRuleSetting `json:"userEligibleSettings,omitempty"`
	// UserMemberSettings undocumented
	UserMemberSettings []GovernanceRuleSetting `json:"userMemberSettings,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *GovernanceRoleDefinition `json:"roleDefinition,omitempty"`
	// Resource undocumented
	Resource *GovernanceResource `json:"resource,omitempty"`
}

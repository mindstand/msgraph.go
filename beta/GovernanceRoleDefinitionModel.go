// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GovernanceRoleDefinition undocumented
type GovernanceRoleDefinition struct {
	Entity
	// ResourceID undocumented
	ResourceID *string `json:"resourceId,omitempty"`
	// ExternalID undocumented
	ExternalID *string `json:"externalId,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Resource undocumented
	Resource *GovernanceResource `json:"resource,omitempty"`
	// RoleSetting undocumented
	RoleSetting *GovernanceRoleSetting `json:"roleSetting,omitempty"`
}

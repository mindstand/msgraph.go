// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UnifiedRoleDefinition undocumented
type UnifiedRoleDefinition struct {
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsBuiltIn undocumented
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// ResourceScopes undocumented
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// RolePermissions undocumented
	RolePermissions []UnifiedRolePermission `json:"rolePermissions,omitempty"`
	// TemplateID undocumented
	TemplateID *string `json:"templateId,omitempty"`
	// Version undocumented
	Version *string `json:"version,omitempty"`
}

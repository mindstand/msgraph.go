// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PublishedResource undocumented
type PublishedResource struct {
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ResourceName undocumented
	ResourceName *string `json:"resourceName,omitempty"`
	// PublishingType undocumented
	PublishingType *OnPremisesPublishingType `json:"publishingType,omitempty"`
	// AgentGroups undocumented
	AgentGroups []OnPremisesAgentGroup `json:"agentGroups,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnPremisesPublishingProfile undocumented
type OnPremisesPublishingProfile struct {
	Entity
	// HybridAgentUpdaterConfiguration undocumented
	HybridAgentUpdaterConfiguration *HybridAgentUpdaterConfiguration `json:"hybridAgentUpdaterConfiguration,omitempty"`
	// Agents undocumented
	Agents []OnPremisesAgent `json:"agents,omitempty"`
	// AgentGroups undocumented
	AgentGroups []OnPremisesAgentGroup `json:"agentGroups,omitempty"`
	// PublishedResources undocumented
	PublishedResources []PublishedResource `json:"publishedResources,omitempty"`
}

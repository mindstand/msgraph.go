// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Team undocumented
type Team struct {
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// InternalID undocumented
	InternalID *string `json:"internalId,omitempty"`
	// Classification undocumented
	Classification *string `json:"classification,omitempty"`
	// Specialization undocumented
	Specialization *TeamSpecialization `json:"specialization,omitempty"`
	// Visibility undocumented
	Visibility *TeamVisibilityType `json:"visibility,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
	// MemberSettings undocumented
	MemberSettings *TeamMemberSettings `json:"memberSettings,omitempty"`
	// GuestSettings undocumented
	GuestSettings *TeamGuestSettings `json:"guestSettings,omitempty"`
	// MessagingSettings undocumented
	MessagingSettings *TeamMessagingSettings `json:"messagingSettings,omitempty"`
	// FunSettings undocumented
	FunSettings *TeamFunSettings `json:"funSettings,omitempty"`
	// DiscoverySettings undocumented
	DiscoverySettings *TeamDiscoverySettings `json:"discoverySettings,omitempty"`
	// IsArchived undocumented
	IsArchived *bool `json:"isArchived,omitempty"`
	// Schedule undocumented
	Schedule *Schedule `json:"schedule,omitempty"`
	// Photo undocumented
	Photo *ProfilePhoto `json:"photo,omitempty"`
	// Template undocumented
	Template *TeamsTemplate `json:"template,omitempty"`
	// Channels undocumented
	Channels []Channel `json:"channels,omitempty"`
	// Apps undocumented
	Apps []TeamsCatalogApp `json:"apps,omitempty"`
	// InstalledApps undocumented
	InstalledApps []TeamsAppInstallation `json:"installedApps,omitempty"`
	// Operations undocumented
	Operations []TeamsAsyncOperation `json:"operations,omitempty"`
	// Owners undocumented
	Owners []User `json:"owners,omitempty"`
	// PrimaryChannel undocumented
	PrimaryChannel *Channel `json:"primaryChannel,omitempty"`
	// Group undocumented
	Group *Group `json:"group,omitempty"`
}

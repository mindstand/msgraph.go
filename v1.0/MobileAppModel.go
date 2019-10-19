// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// MobileApp An abstract class containing the base properties for Intune mobile apps.
type MobileApp struct {
	Entity
	// DisplayName The admin provided or imported title of the app.
	DisplayName *string `json:"displayName,omitempty"`
	// Description The description of the app.
	Description *string `json:"description,omitempty"`
	// Publisher The publisher of the app.
	Publisher *string `json:"publisher,omitempty"`
	// LargeIcon The large icon, to be displayed in the app details and used for upload of the icon.
	LargeIcon *MimeContent `json:"largeIcon,omitempty"`
	// CreatedDateTime The date and time the app was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime The date and time the app was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// IsFeatured The value indicating whether the app is marked as featured by the admin.
	IsFeatured *bool `json:"isFeatured,omitempty"`
	// PrivacyInformationURL The privacy statement Url.
	PrivacyInformationURL *string `json:"privacyInformationUrl,omitempty"`
	// InformationURL The more information Url.
	InformationURL *string `json:"informationUrl,omitempty"`
	// Owner The owner of the app.
	Owner *string `json:"owner,omitempty"`
	// Developer The developer of the app.
	Developer *string `json:"developer,omitempty"`
	// Notes Notes for the app.
	Notes *string `json:"notes,omitempty"`
	// PublishingState The publishing state for the app. The app cannot be assigned unless the app is published.
	PublishingState *MobileAppPublishingState `json:"publishingState,omitempty"`
	// Categories undocumented
	Categories []MobileAppCategory `json:"categories,omitempty"`
	// Assignments undocumented
	Assignments []MobileAppAssignment `json:"assignments,omitempty"`
}

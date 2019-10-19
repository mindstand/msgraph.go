// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// ManagedEBook An abstract class containing the base properties for Managed eBook.
type ManagedEBook struct {
	Entity
	// DisplayName Name of the eBook.
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description.
	Description *string `json:"description,omitempty"`
	// Publisher Publisher.
	Publisher *string `json:"publisher,omitempty"`
	// PublishedDateTime The date and time when the eBook was published.
	PublishedDateTime *time.Time `json:"publishedDateTime,omitempty"`
	// LargeCover Book cover.
	LargeCover *MimeContent `json:"largeCover,omitempty"`
	// CreatedDateTime The date and time when the eBook file was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// LastModifiedDateTime The date and time when the eBook was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// InformationURL The more information Url.
	InformationURL *string `json:"informationUrl,omitempty"`
	// PrivacyInformationURL The privacy statement Url.
	PrivacyInformationURL *string `json:"privacyInformationUrl,omitempty"`
	// Assignments undocumented
	Assignments []ManagedEBookAssignment `json:"assignments,omitempty"`
	// InstallSummary undocumented
	InstallSummary *EBookInstallSummary `json:"installSummary,omitempty"`
	// DeviceStates undocumented
	DeviceStates []DeviceInstallState `json:"deviceStates,omitempty"`
	// UserStateSummary undocumented
	UserStateSummary []UserInstallStateSummary `json:"userStateSummary,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// Subscription undocumented
type Subscription struct {
	Entity
	// Resource undocumented
	Resource *string `json:"resource,omitempty"`
	// ChangeType undocumented
	ChangeType *string `json:"changeType,omitempty"`
	// ClientState undocumented
	ClientState *string `json:"clientState,omitempty"`
	// NotificationURL undocumented
	NotificationURL *string `json:"notificationUrl,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// ApplicationID undocumented
	ApplicationID *string `json:"applicationId,omitempty"`
	// CreatorID undocumented
	CreatorID *string `json:"creatorId,omitempty"`
	// IncludeProperties undocumented
	IncludeProperties *bool `json:"includeProperties,omitempty"`
	// IncludeResourceData undocumented
	IncludeResourceData *bool `json:"includeResourceData,omitempty"`
	// LifecycleNotificationURL undocumented
	LifecycleNotificationURL *string `json:"lifecycleNotificationUrl,omitempty"`
	// EncryptionCertificate undocumented
	EncryptionCertificate *string `json:"encryptionCertificate,omitempty"`
	// EncryptionCertificateID undocumented
	EncryptionCertificateID *string `json:"encryptionCertificateId,omitempty"`
}

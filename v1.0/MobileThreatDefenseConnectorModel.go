// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// MobileThreatDefenseConnector Entity which represents a connection to Mobile threat defense partner.
type MobileThreatDefenseConnector struct {
	Entity
	// LastHeartbeatDateTime DateTime of last Heartbeat recieved from the Data Sync Partner
	LastHeartbeatDateTime *time.Time `json:"lastHeartbeatDateTime,omitempty"`
	// PartnerState Data Sync Partner state for this account
	PartnerState *MobileThreatPartnerTenantState `json:"partnerState,omitempty"`
	// AndroidEnabled For Android, set whether data from the data sync partner should be used during compliance evaluations
	AndroidEnabled *bool `json:"androidEnabled,omitempty"`
	// IOSEnabled For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
	IOSEnabled *bool `json:"iosEnabled,omitempty"`
	// AndroidDeviceBlockedOnMissingPartnerData For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	AndroidDeviceBlockedOnMissingPartnerData *bool `json:"androidDeviceBlockedOnMissingPartnerData,omitempty"`
	// IOSDeviceBlockedOnMissingPartnerData For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
	IOSDeviceBlockedOnMissingPartnerData *bool `json:"iosDeviceBlockedOnMissingPartnerData,omitempty"`
	// PartnerUnsupportedOsVersionBlocked Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
	PartnerUnsupportedOsVersionBlocked *bool `json:"partnerUnsupportedOsVersionBlocked,omitempty"`
	// PartnerUnresponsivenessThresholdInDays Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
	PartnerUnresponsivenessThresholdInDays *int `json:"partnerUnresponsivenessThresholdInDays,omitempty"`
}

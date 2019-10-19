// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// WindowsAutopilotSettings The windowsAutopilotSettings resource represents a Windows Autopilot Account to sync data with Windows device data sync service.
type WindowsAutopilotSettings struct {
	Entity
	// LastSyncDateTime Last data sync date time with DDS service.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// LastManualSyncTriggerDateTime Last data sync date time with DDS service.
	LastManualSyncTriggerDateTime *time.Time `json:"lastManualSyncTriggerDateTime,omitempty"`
	// SyncStatus Indicates the status of sync with Device data sync (DDS) service.
	SyncStatus *WindowsAutopilotSyncStatus `json:"syncStatus,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementExchangeConnectorSyncType undocumented
type DeviceManagementExchangeConnectorSyncType string

const (
	// DeviceManagementExchangeConnectorSyncTypeVFullSync undocumented
	DeviceManagementExchangeConnectorSyncTypeVFullSync DeviceManagementExchangeConnectorSyncType = "FullSync"
	// DeviceManagementExchangeConnectorSyncTypeVDeltaSync undocumented
	DeviceManagementExchangeConnectorSyncTypeVDeltaSync DeviceManagementExchangeConnectorSyncType = "DeltaSync"
)

// DeviceManagementExchangeConnectorSyncTypePFullSync returns a pointer to DeviceManagementExchangeConnectorSyncTypeVFullSync
func DeviceManagementExchangeConnectorSyncTypePFullSync() *DeviceManagementExchangeConnectorSyncType {
	v := DeviceManagementExchangeConnectorSyncTypeVFullSync
	return &v
}

// DeviceManagementExchangeConnectorSyncTypePDeltaSync returns a pointer to DeviceManagementExchangeConnectorSyncTypeVDeltaSync
func DeviceManagementExchangeConnectorSyncTypePDeltaSync() *DeviceManagementExchangeConnectorSyncType {
	v := DeviceManagementExchangeConnectorSyncTypeVDeltaSync
	return &v
}

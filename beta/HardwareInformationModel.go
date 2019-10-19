// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// HardwareInformation undocumented
type HardwareInformation struct {
	// SerialNumber Serial number.
	SerialNumber *string `json:"serialNumber,omitempty"`
	// TotalStorageSpace Total storage space of the device.
	TotalStorageSpace *int `json:"totalStorageSpace,omitempty"`
	// FreeStorageSpace Free storage space of the device.
	FreeStorageSpace *int `json:"freeStorageSpace,omitempty"`
	// Imei IMEI
	Imei *string `json:"imei,omitempty"`
	// Meid MEID
	Meid *string `json:"meid,omitempty"`
	// Manufacturer Manufacturer of the device
	Manufacturer *string `json:"manufacturer,omitempty"`
	// Model Model of the device
	Model *string `json:"model,omitempty"`
	// PhoneNumber Phone number of the device
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// SubscriberCarrier Subscriber carrier of the device
	SubscriberCarrier *string `json:"subscriberCarrier,omitempty"`
	// CellularTechnology Cellular technology of the device
	CellularTechnology *string `json:"cellularTechnology,omitempty"`
	// WifiMac WiFi MAC address of the device
	WifiMac *string `json:"wifiMac,omitempty"`
	// OperatingSystemLanguage Operating system language of the device
	OperatingSystemLanguage *string `json:"operatingSystemLanguage,omitempty"`
	// IsSupervised Supervised mode of the device
	IsSupervised *bool `json:"isSupervised,omitempty"`
	// IsEncrypted Encryption status of the device
	IsEncrypted *bool `json:"isEncrypted,omitempty"`
	// IsSharedDevice Shared iPad
	IsSharedDevice *bool `json:"isSharedDevice,omitempty"`
	// SharedDeviceCachedUsers All users on the shared Apple device
	SharedDeviceCachedUsers []SharedAppleDeviceUser `json:"sharedDeviceCachedUsers,omitempty"`
	// TpmSpecificationVersion String that specifies the specification version.
	TpmSpecificationVersion *string `json:"tpmSpecificationVersion,omitempty"`
	// OperatingSystemEdition String that specifies the OS edition.
	OperatingSystemEdition *string `json:"operatingSystemEdition,omitempty"`
	// DeviceFullQualifiedDomainName Returns the fully qualified domain name of the device (if any). If the device is not domain-joined, it returns an empty string.
	DeviceFullQualifiedDomainName *string `json:"deviceFullQualifiedDomainName,omitempty"`
	// DeviceGuardVirtualizationBasedSecurityHardwareRequirementState Virtualization-based security hardware requirement status.
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState `json:"deviceGuardVirtualizationBasedSecurityHardwareRequirementState,omitempty"`
	// DeviceGuardVirtualizationBasedSecurityState Virtualization-based security status.
	DeviceGuardVirtualizationBasedSecurityState *DeviceGuardVirtualizationBasedSecurityState `json:"deviceGuardVirtualizationBasedSecurityState,omitempty"`
	// DeviceGuardLocalSystemAuthorityCredentialGuardState Local System Authority (LSA) credential guard status.
	DeviceGuardLocalSystemAuthorityCredentialGuardState *DeviceGuardLocalSystemAuthorityCredentialGuardState `json:"deviceGuardLocalSystemAuthorityCredentialGuardState,omitempty"`
	// OsBuildNumber Operating System Build Number on Android device
	OsBuildNumber *string `json:"osBuildNumber,omitempty"`
}

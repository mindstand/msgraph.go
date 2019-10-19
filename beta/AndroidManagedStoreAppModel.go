// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidManagedStoreApp Contains properties and inherited properties for Android Managed Store Apps.
type AndroidManagedStoreApp struct {
	MobileApp
	// PackageID The package identifier.
	PackageID *string `json:"packageId,omitempty"`
	// AppIdentifier The Identity Name.
	AppIdentifier *string `json:"appIdentifier,omitempty"`
	// UsedLicenseCount The number of VPP licenses in use.
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// TotalLicenseCount The total number of VPP licenses.
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// AppStoreURL The Play for Work Store app URL.
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// IsPrivate Indicates whether the app is only available to a given enterprise's users.
	IsPrivate *bool `json:"isPrivate,omitempty"`
	// IsSystemApp Indicates whether the app is a preinstalled system app.
	IsSystemApp *bool `json:"isSystemApp,omitempty"`
	// SupportsOemConfig Whether this app supports OEMConfig policy.
	SupportsOemConfig *bool `json:"supportsOemConfig,omitempty"`
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// IOSVppApp Contains properties and inherited properties for iOS Volume-Purchased Program (VPP) Apps.
type IOSVppApp struct {
	MobileApp
	// UsedLicenseCount The number of VPP licenses in use.
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
	// TotalLicenseCount The total number of VPP licenses.
	TotalLicenseCount *int `json:"totalLicenseCount,omitempty"`
	// ReleaseDateTime The VPP application release date and time.
	ReleaseDateTime *time.Time `json:"releaseDateTime,omitempty"`
	// AppStoreURL The store URL.
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// LicensingType The supported License Type.
	LicensingType *VppLicensingType `json:"licensingType,omitempty"`
	// ApplicableDeviceType The applicable iOS Device Type.
	ApplicableDeviceType *IOSDeviceType `json:"applicableDeviceType,omitempty"`
	// VppTokenOrganizationName The organization associated with the Apple Volume Purchase Program Token
	VppTokenOrganizationName *string `json:"vppTokenOrganizationName,omitempty"`
	// VppTokenAccountType The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: `business`, `education`.
	VppTokenAccountType *VppTokenAccountType `json:"vppTokenAccountType,omitempty"`
	// VppTokenAppleID The Apple Id associated with the given Apple Volume Purchase Program Token.
	VppTokenAppleID *string `json:"vppTokenAppleId,omitempty"`
	// BundleID The Identity Name.
	BundleID *string `json:"bundleId,omitempty"`
}

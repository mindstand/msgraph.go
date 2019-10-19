// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAndroidStoreApp Contains properties and inherited properties for Android store apps that you can manage with an Intune app protection policy.
type ManagedAndroidStoreApp struct {
	ManagedApp
	// PackageID The app's package ID.
	PackageID *string `json:"packageId,omitempty"`
	// AppStoreURL The Android AppStoreUrl.
	AppStoreURL *string `json:"appStoreUrl,omitempty"`
	// MinimumSupportedOperatingSystem The value for the minimum supported operating system.
	MinimumSupportedOperatingSystem *AndroidMinimumOperatingSystem `json:"minimumSupportedOperatingSystem,omitempty"`
}

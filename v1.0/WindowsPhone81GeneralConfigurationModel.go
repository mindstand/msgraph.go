// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPhone81GeneralConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the windowsPhone81GeneralConfiguration resource.
type WindowsPhone81GeneralConfiguration struct {
	DeviceConfiguration
	// ApplyOnlyToWindowsPhone81 Value indicating whether this policy only applies to Windows Phone 8.1. This property is read-only.
	ApplyOnlyToWindowsPhone81 *bool `json:"applyOnlyToWindowsPhone81,omitempty"`
	// AppsBlockCopyPaste Indicates whether or not to block copy paste.
	AppsBlockCopyPaste *bool `json:"appsBlockCopyPaste,omitempty"`
	// BluetoothBlocked Indicates whether or not to block bluetooth.
	BluetoothBlocked *bool `json:"bluetoothBlocked,omitempty"`
	// CameraBlocked Indicates whether or not to block camera.
	CameraBlocked *bool `json:"cameraBlocked,omitempty"`
	// CellularBlockWifiTethering Indicates whether or not to block Wi-Fi tethering. Has no impact if Wi-Fi is blocked.
	CellularBlockWifiTethering *bool `json:"cellularBlockWifiTethering,omitempty"`
	// CompliantAppsList List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
	CompliantAppsList []AppListItem `json:"compliantAppsList,omitempty"`
	// CompliantAppListType List that is in the AppComplianceList.
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`
	// DiagnosticDataBlockSubmission Indicates whether or not to block diagnostic data submission.
	DiagnosticDataBlockSubmission *bool `json:"diagnosticDataBlockSubmission,omitempty"`
	// EmailBlockAddingAccounts Indicates whether or not to block custom email accounts.
	EmailBlockAddingAccounts *bool `json:"emailBlockAddingAccounts,omitempty"`
	// LocationServicesBlocked Indicates whether or not to block location services.
	LocationServicesBlocked *bool `json:"locationServicesBlocked,omitempty"`
	// MicrosoftAccountBlocked Indicates whether or not to block using a Microsoft Account.
	MicrosoftAccountBlocked *bool `json:"microsoftAccountBlocked,omitempty"`
	// NfcBlocked Indicates whether or not to block Near-Field Communication.
	NfcBlocked *bool `json:"nfcBlocked,omitempty"`
	// PasswordBlockSimple Indicates whether or not to block syncing the calendar.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`
	// PasswordExpirationDays Number of days before the password expires.
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumLength Minimum length of passwords.
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout Minutes of inactivity before screen timeout.
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordMinimumCharacterSetCount Number of character sets a password must contain.
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordPreviousPasswordBlockCount Number of previous passwords to block. Valid values 0 to 24
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordSignInFailureCountBeforeFactoryReset Number of sign in failures allowed before factory reset.
	PasswordSignInFailureCountBeforeFactoryReset *int `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	// PasswordRequiredType Password type that is required.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordRequired Indicates whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
	// ScreenCaptureBlocked Indicates whether or not to block screenshots.
	ScreenCaptureBlocked *bool `json:"screenCaptureBlocked,omitempty"`
	// StorageBlockRemovableStorage Indicates whether or not to block removable storage.
	StorageBlockRemovableStorage *bool `json:"storageBlockRemovableStorage,omitempty"`
	// StorageRequireEncryption Indicates whether or not to require encryption.
	StorageRequireEncryption *bool `json:"storageRequireEncryption,omitempty"`
	// WebBrowserBlocked Indicates whether or not to block the web browser.
	WebBrowserBlocked *bool `json:"webBrowserBlocked,omitempty"`
	// WifiBlocked Indicates whether or not to block Wi-Fi.
	WifiBlocked *bool `json:"wifiBlocked,omitempty"`
	// WifiBlockAutomaticConnectHotspots Indicates whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
	WifiBlockAutomaticConnectHotspots *bool `json:"wifiBlockAutomaticConnectHotspots,omitempty"`
	// WifiBlockHotspotReporting Indicates whether or not to block Wi-Fi hotspot reporting. Has no impact if Wi-Fi is blocked.
	WifiBlockHotspotReporting *bool `json:"wifiBlockHotspotReporting,omitempty"`
	// WindowsStoreBlocked Indicates whether or not to block the Windows Store.
	WindowsStoreBlocked *bool `json:"windowsStoreBlocked,omitempty"`
}

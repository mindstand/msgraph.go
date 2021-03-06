// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SharedPCConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the sharedPCConfiguration resource.
type SharedPCConfiguration struct {
	// DeviceConfiguration is the base model of SharedPCConfiguration
	DeviceConfiguration
	// AccountManagerPolicy Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
	AccountManagerPolicy *SharedPCAccountManagerPolicy `json:"accountManagerPolicy,omitempty"`
	// AllowedAccounts Indicates which type of accounts are allowed to use on a shared PC.
	AllowedAccounts *SharedPCAllowedAccountType `json:"allowedAccounts,omitempty"`
	// AllowLocalStorage Specifies whether local storage is allowed on a shared PC.
	AllowLocalStorage *bool `json:"allowLocalStorage,omitempty"`
	// DisableAccountManager Disables the account manager for shared PC mode.
	DisableAccountManager *bool `json:"disableAccountManager,omitempty"`
	// DisableEduPolicies Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true.
	DisableEduPolicies *bool `json:"disableEduPolicies,omitempty"`
	// DisablePowerPolicies Specifies whether the default shared PC power policies should be disabled.
	DisablePowerPolicies *bool `json:"disablePowerPolicies,omitempty"`
	// DisableSignInOnResume Disables the requirement to sign in whenever the device wakes up from sleep mode.
	DisableSignInOnResume *bool `json:"disableSignInOnResume,omitempty"`
	// Enabled Enables shared PC mode and applies the shared pc policies.
	Enabled *bool `json:"enabled,omitempty"`
	// IdleTimeBeforeSleepInSeconds Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring.
	IdleTimeBeforeSleepInSeconds *int `json:"idleTimeBeforeSleepInSeconds,omitempty"`
	// KioskAppDisplayName Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
	KioskAppDisplayName *string `json:"kioskAppDisplayName,omitempty"`
	// KioskAppUserModelID Specifies the application user model ID of the app to use with assigned access.
	KioskAppUserModelID *string `json:"kioskAppUserModelId,omitempty"`
	// MaintenanceStartTime Specifies the daily start time of maintenance hour.
	MaintenanceStartTime *time.Time `json:"maintenanceStartTime,omitempty"`
}

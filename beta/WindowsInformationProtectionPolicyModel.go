// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsInformationProtectionPolicy Policy for Windows information protection without MDM
type WindowsInformationProtectionPolicy struct {
	WindowsInformationProtection
	// RevokeOnMdmHandoffDisabled New property in RS2, pending documentation
	RevokeOnMdmHandoffDisabled *bool `json:"revokeOnMdmHandoffDisabled,omitempty"`
	// MdmEnrollmentURL Enrollment url for the MDM
	MdmEnrollmentURL *string `json:"mdmEnrollmentUrl,omitempty"`
	// WindowsHelloForBusinessBlocked Boolean value that sets Windows Hello for Business as a method for signing into Windows.
	WindowsHelloForBusinessBlocked *bool `json:"windowsHelloForBusinessBlocked,omitempty"`
	// PinMinimumLength Integer value that sets the minimum number of characters required for the PIN. Default value is 4. The lowest number you can configure for this policy setting is 4. The largest number you can configure must be less than the number configured in the Maximum PIN length policy setting or the number 127, whichever is the lowest.
	PinMinimumLength *int `json:"pinMinimumLength,omitempty"`
	// PinUppercaseLetters Integer value that configures the use of uppercase letters in the Windows Hello for Business PIN. Default is NotAllow.
	PinUppercaseLetters *WindowsInformationProtectionPinCharacterRequirements `json:"pinUppercaseLetters,omitempty"`
	// PinLowercaseLetters Integer value that configures the use of lowercase letters in the Windows Hello for Business PIN. Default is NotAllow.
	PinLowercaseLetters *WindowsInformationProtectionPinCharacterRequirements `json:"pinLowercaseLetters,omitempty"`
	// PinSpecialCharacters Integer value that configures the use of special characters in the Windows Hello for Business PIN. Valid special characters for Windows Hello for Business PIN gestures include: ! " # $ % & ' ( ) * + , - . / : ; < = > ? @ [ \ ] ^ _ ` { | } ~. Default is NotAllow.
	PinSpecialCharacters *WindowsInformationProtectionPinCharacterRequirements `json:"pinSpecialCharacters,omitempty"`
	// PinExpirationDays Integer value specifies the period of time (in days) that a PIN can be used before the system requires the user to change it. The largest number you can configure for this policy setting is 730. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then the user's PIN will never expire. This node was added in Windows 10, version 1511. Default is 0.
	PinExpirationDays *int `json:"pinExpirationDays,omitempty"`
	// NumberOfPastPinsRemembered Integer value that specifies the number of past PINs that can be associated to a user account that can't be reused. The largest number you can configure for this policy setting is 50. The lowest number you can configure for this policy setting is 0. If this policy is set to 0, then storage of previous PINs is not required. This node was added in Windows 10, version 1511. Default is 0.
	NumberOfPastPinsRemembered *int `json:"numberOfPastPinsRemembered,omitempty"`
	// PasswordMaximumAttemptCount The number of authentication failures allowed before the device will be wiped. A value of 0 disables device wipe functionality. Range is an integer X where 4 <= X <= 16 for desktop and 0 <= X <= 999 for mobile devices.
	PasswordMaximumAttemptCount *int `json:"passwordMaximumAttemptCount,omitempty"`
	// MinutesOfInactivityBeforeDeviceLock Specifies the maximum amount of time (in minutes) allowed after the device is idle that will cause the device to become PIN or password locked.   Range is an integer X where 0 <= X <= 999.
	MinutesOfInactivityBeforeDeviceLock *int `json:"minutesOfInactivityBeforeDeviceLock,omitempty"`
	// DaysWithoutContactBeforeUnenroll Offline interval before app data is wiped (days)
	DaysWithoutContactBeforeUnenroll *int `json:"daysWithoutContactBeforeUnenroll,omitempty"`
}

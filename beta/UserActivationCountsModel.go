// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// UserActivationCounts undocumented
type UserActivationCounts struct {
	// ProductType undocumented
	ProductType *string `json:"productType,omitempty"`
	// LastActivatedDate undocumented
	LastActivatedDate *time.Time `json:"lastActivatedDate,omitempty"`
	// Windows undocumented
	Windows *int `json:"windows,omitempty"`
	// Mac undocumented
	Mac *int `json:"mac,omitempty"`
	// Windows10Mobile undocumented
	Windows10Mobile *int `json:"windows10Mobile,omitempty"`
	// IOS undocumented
	IOS *int `json:"ios,omitempty"`
	// Android undocumented
	Android *int `json:"android,omitempty"`
	// ActivatedOnSharedComputer undocumented
	ActivatedOnSharedComputer *bool `json:"activatedOnSharedComputer,omitempty"`
}

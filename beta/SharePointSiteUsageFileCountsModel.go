// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SharePointSiteUsageFileCounts undocumented
type SharePointSiteUsageFileCounts struct {
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// SiteType undocumented
	SiteType *string `json:"siteType,omitempty"`
	// Total undocumented
	Total *int `json:"total,omitempty"`
	// Active undocumented
	Active *int `json:"active,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}

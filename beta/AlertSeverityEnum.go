// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AlertSeverity undocumented
type AlertSeverity string

const (
	// AlertSeverityVUnknown undocumented
	AlertSeverityVUnknown AlertSeverity = "Unknown"
	// AlertSeverityVInformational undocumented
	AlertSeverityVInformational AlertSeverity = "Informational"
	// AlertSeverityVLow undocumented
	AlertSeverityVLow AlertSeverity = "Low"
	// AlertSeverityVMedium undocumented
	AlertSeverityVMedium AlertSeverity = "Medium"
	// AlertSeverityVHigh undocumented
	AlertSeverityVHigh AlertSeverity = "High"
	// AlertSeverityVUnknownFutureValue undocumented
	AlertSeverityVUnknownFutureValue AlertSeverity = "UnknownFutureValue"
)

// AlertSeverityPUnknown returns a pointer to AlertSeverityVUnknown
func AlertSeverityPUnknown() *AlertSeverity {
	v := AlertSeverityVUnknown
	return &v
}

// AlertSeverityPInformational returns a pointer to AlertSeverityVInformational
func AlertSeverityPInformational() *AlertSeverity {
	v := AlertSeverityVInformational
	return &v
}

// AlertSeverityPLow returns a pointer to AlertSeverityVLow
func AlertSeverityPLow() *AlertSeverity {
	v := AlertSeverityVLow
	return &v
}

// AlertSeverityPMedium returns a pointer to AlertSeverityVMedium
func AlertSeverityPMedium() *AlertSeverity {
	v := AlertSeverityVMedium
	return &v
}

// AlertSeverityPHigh returns a pointer to AlertSeverityVHigh
func AlertSeverityPHigh() *AlertSeverity {
	v := AlertSeverityVHigh
	return &v
}

// AlertSeverityPUnknownFutureValue returns a pointer to AlertSeverityVUnknownFutureValue
func AlertSeverityPUnknownFutureValue() *AlertSeverity {
	v := AlertSeverityVUnknownFutureValue
	return &v
}

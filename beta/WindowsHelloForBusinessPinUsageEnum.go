// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsHelloForBusinessPinUsage undocumented
type WindowsHelloForBusinessPinUsage int

const (
	// WindowsHelloForBusinessPinUsageVAllowed undocumented
	WindowsHelloForBusinessPinUsageVAllowed WindowsHelloForBusinessPinUsage = 0
	// WindowsHelloForBusinessPinUsageVRequired undocumented
	WindowsHelloForBusinessPinUsageVRequired WindowsHelloForBusinessPinUsage = 1
	// WindowsHelloForBusinessPinUsageVDisallowed undocumented
	WindowsHelloForBusinessPinUsageVDisallowed WindowsHelloForBusinessPinUsage = 2
)

// WindowsHelloForBusinessPinUsagePAllowed returns a pointer to WindowsHelloForBusinessPinUsageVAllowed
func WindowsHelloForBusinessPinUsagePAllowed() *WindowsHelloForBusinessPinUsage {
	v := WindowsHelloForBusinessPinUsageVAllowed
	return &v
}

// WindowsHelloForBusinessPinUsagePRequired returns a pointer to WindowsHelloForBusinessPinUsageVRequired
func WindowsHelloForBusinessPinUsagePRequired() *WindowsHelloForBusinessPinUsage {
	v := WindowsHelloForBusinessPinUsageVRequired
	return &v
}

// WindowsHelloForBusinessPinUsagePDisallowed returns a pointer to WindowsHelloForBusinessPinUsageVDisallowed
func WindowsHelloForBusinessPinUsagePDisallowed() *WindowsHelloForBusinessPinUsage {
	v := WindowsHelloForBusinessPinUsageVDisallowed
	return &v
}

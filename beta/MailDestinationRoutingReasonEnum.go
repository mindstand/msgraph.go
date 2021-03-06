// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MailDestinationRoutingReason undocumented
type MailDestinationRoutingReason string

const (
	// MailDestinationRoutingReasonVNone undocumented
	MailDestinationRoutingReasonVNone MailDestinationRoutingReason = "None"
	// MailDestinationRoutingReasonVMailFlowRule undocumented
	MailDestinationRoutingReasonVMailFlowRule MailDestinationRoutingReason = "MailFlowRule"
	// MailDestinationRoutingReasonVSafeSender undocumented
	MailDestinationRoutingReasonVSafeSender MailDestinationRoutingReason = "SafeSender"
	// MailDestinationRoutingReasonVBlockedSender undocumented
	MailDestinationRoutingReasonVBlockedSender MailDestinationRoutingReason = "BlockedSender"
	// MailDestinationRoutingReasonVAdvancedSpamFiltering undocumented
	MailDestinationRoutingReasonVAdvancedSpamFiltering MailDestinationRoutingReason = "AdvancedSpamFiltering"
	// MailDestinationRoutingReasonVDomainAllowList undocumented
	MailDestinationRoutingReasonVDomainAllowList MailDestinationRoutingReason = "DomainAllowList"
	// MailDestinationRoutingReasonVDomainBlockList undocumented
	MailDestinationRoutingReasonVDomainBlockList MailDestinationRoutingReason = "DomainBlockList"
	// MailDestinationRoutingReasonVNotInAddressBook undocumented
	MailDestinationRoutingReasonVNotInAddressBook MailDestinationRoutingReason = "NotInAddressBook"
	// MailDestinationRoutingReasonVFirstTimeSender undocumented
	MailDestinationRoutingReasonVFirstTimeSender MailDestinationRoutingReason = "FirstTimeSender"
	// MailDestinationRoutingReasonVAutoPurgeToInbox undocumented
	MailDestinationRoutingReasonVAutoPurgeToInbox MailDestinationRoutingReason = "AutoPurgeToInbox"
	// MailDestinationRoutingReasonVAutoPurgeToJunk undocumented
	MailDestinationRoutingReasonVAutoPurgeToJunk MailDestinationRoutingReason = "AutoPurgeToJunk"
	// MailDestinationRoutingReasonVAutoPurgeToDeleted undocumented
	MailDestinationRoutingReasonVAutoPurgeToDeleted MailDestinationRoutingReason = "AutoPurgeToDeleted"
	// MailDestinationRoutingReasonVOutbound undocumented
	MailDestinationRoutingReasonVOutbound MailDestinationRoutingReason = "Outbound"
	// MailDestinationRoutingReasonVNotJunk undocumented
	MailDestinationRoutingReasonVNotJunk MailDestinationRoutingReason = "NotJunk"
	// MailDestinationRoutingReasonVJunk undocumented
	MailDestinationRoutingReasonVJunk MailDestinationRoutingReason = "Junk"
	// MailDestinationRoutingReasonVUnknownFutureValue undocumented
	MailDestinationRoutingReasonVUnknownFutureValue MailDestinationRoutingReason = "UnknownFutureValue"
)

// MailDestinationRoutingReasonPNone returns a pointer to MailDestinationRoutingReasonVNone
func MailDestinationRoutingReasonPNone() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVNone
	return &v
}

// MailDestinationRoutingReasonPMailFlowRule returns a pointer to MailDestinationRoutingReasonVMailFlowRule
func MailDestinationRoutingReasonPMailFlowRule() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVMailFlowRule
	return &v
}

// MailDestinationRoutingReasonPSafeSender returns a pointer to MailDestinationRoutingReasonVSafeSender
func MailDestinationRoutingReasonPSafeSender() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVSafeSender
	return &v
}

// MailDestinationRoutingReasonPBlockedSender returns a pointer to MailDestinationRoutingReasonVBlockedSender
func MailDestinationRoutingReasonPBlockedSender() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVBlockedSender
	return &v
}

// MailDestinationRoutingReasonPAdvancedSpamFiltering returns a pointer to MailDestinationRoutingReasonVAdvancedSpamFiltering
func MailDestinationRoutingReasonPAdvancedSpamFiltering() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVAdvancedSpamFiltering
	return &v
}

// MailDestinationRoutingReasonPDomainAllowList returns a pointer to MailDestinationRoutingReasonVDomainAllowList
func MailDestinationRoutingReasonPDomainAllowList() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVDomainAllowList
	return &v
}

// MailDestinationRoutingReasonPDomainBlockList returns a pointer to MailDestinationRoutingReasonVDomainBlockList
func MailDestinationRoutingReasonPDomainBlockList() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVDomainBlockList
	return &v
}

// MailDestinationRoutingReasonPNotInAddressBook returns a pointer to MailDestinationRoutingReasonVNotInAddressBook
func MailDestinationRoutingReasonPNotInAddressBook() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVNotInAddressBook
	return &v
}

// MailDestinationRoutingReasonPFirstTimeSender returns a pointer to MailDestinationRoutingReasonVFirstTimeSender
func MailDestinationRoutingReasonPFirstTimeSender() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVFirstTimeSender
	return &v
}

// MailDestinationRoutingReasonPAutoPurgeToInbox returns a pointer to MailDestinationRoutingReasonVAutoPurgeToInbox
func MailDestinationRoutingReasonPAutoPurgeToInbox() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVAutoPurgeToInbox
	return &v
}

// MailDestinationRoutingReasonPAutoPurgeToJunk returns a pointer to MailDestinationRoutingReasonVAutoPurgeToJunk
func MailDestinationRoutingReasonPAutoPurgeToJunk() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVAutoPurgeToJunk
	return &v
}

// MailDestinationRoutingReasonPAutoPurgeToDeleted returns a pointer to MailDestinationRoutingReasonVAutoPurgeToDeleted
func MailDestinationRoutingReasonPAutoPurgeToDeleted() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVAutoPurgeToDeleted
	return &v
}

// MailDestinationRoutingReasonPOutbound returns a pointer to MailDestinationRoutingReasonVOutbound
func MailDestinationRoutingReasonPOutbound() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVOutbound
	return &v
}

// MailDestinationRoutingReasonPNotJunk returns a pointer to MailDestinationRoutingReasonVNotJunk
func MailDestinationRoutingReasonPNotJunk() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVNotJunk
	return &v
}

// MailDestinationRoutingReasonPJunk returns a pointer to MailDestinationRoutingReasonVJunk
func MailDestinationRoutingReasonPJunk() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVJunk
	return &v
}

// MailDestinationRoutingReasonPUnknownFutureValue returns a pointer to MailDestinationRoutingReasonVUnknownFutureValue
func MailDestinationRoutingReasonPUnknownFutureValue() *MailDestinationRoutingReason {
	v := MailDestinationRoutingReasonVUnknownFutureValue
	return &v
}

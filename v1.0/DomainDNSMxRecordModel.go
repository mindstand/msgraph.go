// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DomainDNSMxRecord undocumented
type DomainDNSMxRecord struct {
	// DomainDNSRecord is the base model of DomainDNSMxRecord
	DomainDNSRecord
	// MailExchange undocumented
	MailExchange *string `json:"mailExchange,omitempty"`
	// Preference undocumented
	Preference *int `json:"preference,omitempty"`
}

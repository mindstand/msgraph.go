// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DomainDNSRecord undocumented
type DomainDNSRecord struct {
	Entity
	// IsOptional undocumented
	IsOptional *bool `json:"isOptional,omitempty"`
	// Label undocumented
	Label *string `json:"label,omitempty"`
	// RecordType undocumented
	RecordType *string `json:"recordType,omitempty"`
	// SupportedService undocumented
	SupportedService *string `json:"supportedService,omitempty"`
	// TTL undocumented
	TTL *int `json:"ttl,omitempty"`
}

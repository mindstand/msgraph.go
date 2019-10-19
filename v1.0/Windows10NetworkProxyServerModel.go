// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// Windows10NetworkProxyServer undocumented
type Windows10NetworkProxyServer struct {
	// Address Address to the proxy server. Specify an address in the format <server>[“:”<port>]
	Address *string `json:"address,omitempty"`
	// Exceptions Addresses that should not use the proxy server. The system will not use the proxy server for addresses beginning with what is specified in this node.
	Exceptions []string `json:"exceptions,omitempty"`
	// UseForLocalAddresses Specifies whether the proxy server should be used for local (intranet) addresses.
	UseForLocalAddresses *bool `json:"useForLocalAddresses,omitempty"`
}

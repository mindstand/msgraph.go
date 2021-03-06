// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsWifiEnterpriseEAPConfiguration This entity provides descriptions of the declared methods, properties and relationships exposed by the Wifi CSP.
type WindowsWifiEnterpriseEAPConfiguration struct {
	// WindowsWifiConfiguration is the base model of WindowsWifiEnterpriseEAPConfiguration
	WindowsWifiConfiguration
	// NetworkSingleSignOn Specify the network single sign on type.
	NetworkSingleSignOn *NetworkSingleSignOnType `json:"networkSingleSignOn,omitempty"`
	// MaximumAuthenticationTimeoutInSeconds Specify maximum authentication timeout (in seconds).  Valid range: 1-120
	MaximumAuthenticationTimeoutInSeconds *int `json:"maximumAuthenticationTimeoutInSeconds,omitempty"`
	// PromptForAdditionalAuthenticationCredentials Specify whether the wifi connection should prompt for additional authentication credentials.
	PromptForAdditionalAuthenticationCredentials *bool `json:"promptForAdditionalAuthenticationCredentials,omitempty"`
	// EnablePairwiseMasterKeyCaching Specify whether the wifi connection should enable pairwise master key caching.
	EnablePairwiseMasterKeyCaching *bool `json:"enablePairwiseMasterKeyCaching,omitempty"`
	// MaximumPairwiseMasterKeyCacheTimeInMinutes Specify maximum pairwise master key cache time (in minutes).  Valid range: 5-1440
	MaximumPairwiseMasterKeyCacheTimeInMinutes *int `json:"maximumPairwiseMasterKeyCacheTimeInMinutes,omitempty"`
	// MaximumNumberOfPairwiseMasterKeysInCache Specify maximum number of pairwise master keys in cache.  Valid range: 1-255
	MaximumNumberOfPairwiseMasterKeysInCache *int `json:"maximumNumberOfPairwiseMasterKeysInCache,omitempty"`
	// EnablePreAuthentication Specify whether pre-authentication should be enabled.
	EnablePreAuthentication *bool `json:"enablePreAuthentication,omitempty"`
	// MaximumPreAuthenticationAttempts Specify maximum pre-authentication attempts.  Valid range: 1-16
	MaximumPreAuthenticationAttempts *int `json:"maximumPreAuthenticationAttempts,omitempty"`
	// EapType Extensible Authentication Protocol (EAP). Indicates the type of EAP protocol set on the Wi-Fi endpoint (router).
	EapType *EapType `json:"eapType,omitempty"`
	// TrustedServerCertificateNames Specify trusted server certificate names.
	TrustedServerCertificateNames []string `json:"trustedServerCertificateNames,omitempty"`
	// AuthenticationMethod Specify the authentication method.
	AuthenticationMethod *WiFiAuthenticationMethod `json:"authenticationMethod,omitempty"`
	// InnerAuthenticationProtocolForEAPTTLS Specify inner authentication protocol for EAP TTLS.
	InnerAuthenticationProtocolForEAPTTLS *NonEapAuthenticationMethodForEapTtlsType `json:"innerAuthenticationProtocolForEAPTTLS,omitempty"`
	// OuterIdentityPrivacyTemporaryValue Specify the string to replace usernames for privacy when using EAP TTLS or PEAP.
	OuterIdentityPrivacyTemporaryValue *string `json:"outerIdentityPrivacyTemporaryValue,omitempty"`
	// RootCertificatesForServerValidation undocumented
	RootCertificatesForServerValidation []Windows81TrustedRootCertificate `json:"rootCertificatesForServerValidation,omitempty"`
	// IdentityCertificateForClientAuthentication undocumented
	IdentityCertificateForClientAuthentication *WindowsCertificateProfileBase `json:"identityCertificateForClientAuthentication,omitempty"`
}

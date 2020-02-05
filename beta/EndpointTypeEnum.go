// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EndpointType undocumented
type EndpointType string

const (
	// EndpointTypeVDefault undocumented
	EndpointTypeVDefault EndpointType = "Default"
	// EndpointTypeVVoicemail undocumented
	EndpointTypeVVoicemail EndpointType = "Voicemail"
	// EndpointTypeVSkypeForBusiness undocumented
	EndpointTypeVSkypeForBusiness EndpointType = "SkypeForBusiness"
	// EndpointTypeVSkypeForBusinessVoipPhone undocumented
	EndpointTypeVSkypeForBusinessVoipPhone EndpointType = "SkypeForBusinessVoipPhone"
	// EndpointTypeVUnknownFutureValue undocumented
	EndpointTypeVUnknownFutureValue EndpointType = "UnknownFutureValue"
)

// EndpointTypePDefault returns a pointer to EndpointTypeVDefault
func EndpointTypePDefault() *EndpointType {
	v := EndpointTypeVDefault
	return &v
}

// EndpointTypePVoicemail returns a pointer to EndpointTypeVVoicemail
func EndpointTypePVoicemail() *EndpointType {
	v := EndpointTypeVVoicemail
	return &v
}

// EndpointTypePSkypeForBusiness returns a pointer to EndpointTypeVSkypeForBusiness
func EndpointTypePSkypeForBusiness() *EndpointType {
	v := EndpointTypeVSkypeForBusiness
	return &v
}

// EndpointTypePSkypeForBusinessVoipPhone returns a pointer to EndpointTypeVSkypeForBusinessVoipPhone
func EndpointTypePSkypeForBusinessVoipPhone() *EndpointType {
	v := EndpointTypeVSkypeForBusinessVoipPhone
	return &v
}

// EndpointTypePUnknownFutureValue returns a pointer to EndpointTypeVUnknownFutureValue
func EndpointTypePUnknownFutureValue() *EndpointType {
	v := EndpointTypeVUnknownFutureValue
	return &v
}

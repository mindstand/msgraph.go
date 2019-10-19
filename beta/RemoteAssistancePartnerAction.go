// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RemoteAssistancePartnerBeginOnboardingRequestParameter undocumented
type RemoteAssistancePartnerBeginOnboardingRequestParameter struct {
}

// RemoteAssistancePartnerDisconnectRequestParameter undocumented
type RemoteAssistancePartnerDisconnectRequestParameter struct {
}

//
type RemoteAssistancePartnerBeginOnboardingRequestBuilder struct{ BaseRequestBuilder }

// BeginOnboarding action undocumented
func (b *RemoteAssistancePartnerRequestBuilder) BeginOnboarding(reqObj *RemoteAssistancePartnerBeginOnboardingRequestParameter) *RemoteAssistancePartnerBeginOnboardingRequestBuilder {
	bb := &RemoteAssistancePartnerBeginOnboardingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/beginOnboarding"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RemoteAssistancePartnerBeginOnboardingRequest struct{ BaseRequest }

//
func (b *RemoteAssistancePartnerBeginOnboardingRequestBuilder) Request() *RemoteAssistancePartnerBeginOnboardingRequest {
	return &RemoteAssistancePartnerBeginOnboardingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RemoteAssistancePartnerBeginOnboardingRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *RemoteAssistancePartnerBeginOnboardingRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

//
type RemoteAssistancePartnerDisconnectRequestBuilder struct{ BaseRequestBuilder }

// Disconnect action undocumented
func (b *RemoteAssistancePartnerRequestBuilder) Disconnect(reqObj *RemoteAssistancePartnerDisconnectRequestParameter) *RemoteAssistancePartnerDisconnectRequestBuilder {
	bb := &RemoteAssistancePartnerDisconnectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/disconnect"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type RemoteAssistancePartnerDisconnectRequest struct{ BaseRequest }

//
func (b *RemoteAssistancePartnerDisconnectRequestBuilder) Request() *RemoteAssistancePartnerDisconnectRequest {
	return &RemoteAssistancePartnerDisconnectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *RemoteAssistancePartnerDisconnectRequest) Do(method, path string, reqObj interface{}) error {
	return r.JSONRequestWithPath(method, path, reqObj, nil)
}

//
func (r *RemoteAssistancePartnerDisconnectRequest) Post() error {
	return r.Do("POST", "", r.requestObject)
}

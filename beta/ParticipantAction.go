// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ParticipantMuteRequestParameter undocumented
type ParticipantMuteRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// ParticipantCollectionInviteRequestParameter undocumented
type ParticipantCollectionInviteRequestParameter struct {
	// Participants undocumented
	Participants []InvitationParticipantInfo `json:"participants,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// ParticipantCollectionMuteAllRequestParameter undocumented
type ParticipantCollectionMuteAllRequestParameter struct {
	// Participants undocumented
	Participants []string `json:"participants,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

//
type ParticipantMuteRequestBuilder struct{ BaseRequestBuilder }

// Mute action undocumented
func (b *ParticipantRequestBuilder) Mute(reqObj *ParticipantMuteRequestParameter) *ParticipantMuteRequestBuilder {
	bb := &ParticipantMuteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/mute"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantMuteRequest struct{ BaseRequest }

//
func (b *ParticipantMuteRequestBuilder) Request() *ParticipantMuteRequest {
	return &ParticipantMuteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantMuteRequest) Do(method, path string, reqObj interface{}) (resObj *MuteParticipantOperation, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *ParticipantMuteRequest) Post() (*MuteParticipantOperation, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type ParticipantCollectionInviteRequestBuilder struct{ BaseRequestBuilder }

// Invite action undocumented
func (b *CallParticipantsCollectionRequestBuilder) Invite(reqObj *ParticipantCollectionInviteRequestParameter) *ParticipantCollectionInviteRequestBuilder {
	bb := &ParticipantCollectionInviteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/invite"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantCollectionInviteRequest struct{ BaseRequest }

//
func (b *ParticipantCollectionInviteRequestBuilder) Request() *ParticipantCollectionInviteRequest {
	return &ParticipantCollectionInviteRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantCollectionInviteRequest) Do(method, path string, reqObj interface{}) (resObj *InviteParticipantsOperation, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *ParticipantCollectionInviteRequest) Post() (*InviteParticipantsOperation, error) {
	return r.Do("POST", "", r.requestObject)
}

//
type ParticipantCollectionMuteAllRequestBuilder struct{ BaseRequestBuilder }

// MuteAll action undocumented
func (b *CallParticipantsCollectionRequestBuilder) MuteAll(reqObj *ParticipantCollectionMuteAllRequestParameter) *ParticipantCollectionMuteAllRequestBuilder {
	bb := &ParticipantCollectionMuteAllRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/muteAll"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ParticipantCollectionMuteAllRequest struct{ BaseRequest }

//
func (b *ParticipantCollectionMuteAllRequestBuilder) Request() *ParticipantCollectionMuteAllRequest {
	return &ParticipantCollectionMuteAllRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ParticipantCollectionMuteAllRequest) Do(method, path string, reqObj interface{}) (resObj *MuteParticipantsOperation, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

//
func (r *ParticipantCollectionMuteAllRequest) Post() (*MuteParticipantsOperation, error) {
	return r.Do("POST", "", r.requestObject)
}

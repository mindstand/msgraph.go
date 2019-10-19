// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// InvitationRequestBuilder is request builder for Invitation
type InvitationRequestBuilder struct{ BaseRequestBuilder }

// Request returns InvitationRequest
func (b *InvitationRequestBuilder) Request() *InvitationRequest {
	return &InvitationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// InvitationRequest is request for Invitation
type InvitationRequest struct{ BaseRequest }

// Do performs HTTP request for Invitation
func (r *InvitationRequest) Do(method, path string, reqObj interface{}) (resObj *Invitation, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Invitation
func (r *InvitationRequest) Get() (*Invitation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Invitation
func (r *InvitationRequest) Update(reqObj *Invitation) (*Invitation, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Invitation
func (r *InvitationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// InvitedUser is navigation property
func (b *InvitationRequestBuilder) InvitedUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/invitedUser"
	return bb
}

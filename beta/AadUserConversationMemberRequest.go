// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AadUserConversationMemberRequestBuilder is request builder for AadUserConversationMember
type AadUserConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns AadUserConversationMemberRequest
func (b *AadUserConversationMemberRequestBuilder) Request() *AadUserConversationMemberRequest {
	return &AadUserConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AadUserConversationMemberRequest is request for AadUserConversationMember
type AadUserConversationMemberRequest struct{ BaseRequest }

// Do performs HTTP request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Do(method, path string, reqObj interface{}) (resObj *AadUserConversationMember, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Get() (*AadUserConversationMember, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Update(reqObj *AadUserConversationMember) (*AadUserConversationMember, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AadUserConversationMember
func (r *AadUserConversationMemberRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// User is navigation property
func (b *AadUserConversationMemberRequestBuilder) User() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/user"
	return bb
}

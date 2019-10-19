// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RemoteAssistancePartnerRequestBuilder is request builder for RemoteAssistancePartner
type RemoteAssistancePartnerRequestBuilder struct{ BaseRequestBuilder }

// Request returns RemoteAssistancePartnerRequest
func (b *RemoteAssistancePartnerRequestBuilder) Request() *RemoteAssistancePartnerRequest {
	return &RemoteAssistancePartnerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RemoteAssistancePartnerRequest is request for RemoteAssistancePartner
type RemoteAssistancePartnerRequest struct{ BaseRequest }

// Do performs HTTP request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Do(method, path string, reqObj interface{}) (resObj *RemoteAssistancePartner, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Get() (*RemoteAssistancePartner, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Update(reqObj *RemoteAssistancePartner) (*RemoteAssistancePartner, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for RemoteAssistancePartner
func (r *RemoteAssistancePartnerRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

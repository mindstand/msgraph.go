// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IdentityRiskEventRequestBuilder is request builder for IdentityRiskEvent
type IdentityRiskEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns IdentityRiskEventRequest
func (b *IdentityRiskEventRequestBuilder) Request() *IdentityRiskEventRequest {
	return &IdentityRiskEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IdentityRiskEventRequest is request for IdentityRiskEvent
type IdentityRiskEventRequest struct{ BaseRequest }

// Do performs HTTP request for IdentityRiskEvent
func (r *IdentityRiskEventRequest) Do(method, path string, reqObj interface{}) (resObj *IdentityRiskEvent, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for IdentityRiskEvent
func (r *IdentityRiskEventRequest) Get() (*IdentityRiskEvent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for IdentityRiskEvent
func (r *IdentityRiskEventRequest) Update(reqObj *IdentityRiskEvent) (*IdentityRiskEvent, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for IdentityRiskEvent
func (r *IdentityRiskEventRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// ImpactedUser is navigation property
func (b *IdentityRiskEventRequestBuilder) ImpactedUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/impactedUser"
	return bb
}

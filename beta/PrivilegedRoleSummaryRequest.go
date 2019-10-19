// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PrivilegedRoleSummaryRequestBuilder is request builder for PrivilegedRoleSummary
type PrivilegedRoleSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleSummaryRequest
func (b *PrivilegedRoleSummaryRequestBuilder) Request() *PrivilegedRoleSummaryRequest {
	return &PrivilegedRoleSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleSummaryRequest is request for PrivilegedRoleSummary
type PrivilegedRoleSummaryRequest struct{ BaseRequest }

// Do performs HTTP request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Do(method, path string, reqObj interface{}) (resObj *PrivilegedRoleSummary, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Get() (*PrivilegedRoleSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Update(reqObj *PrivilegedRoleSummary) (*PrivilegedRoleSummary, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

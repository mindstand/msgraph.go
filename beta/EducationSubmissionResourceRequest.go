// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationSubmissionResourceRequestBuilder is request builder for EducationSubmissionResource
type EducationSubmissionResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationSubmissionResourceRequest
func (b *EducationSubmissionResourceRequestBuilder) Request() *EducationSubmissionResourceRequest {
	return &EducationSubmissionResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationSubmissionResourceRequest is request for EducationSubmissionResource
type EducationSubmissionResourceRequest struct{ BaseRequest }

// Do performs HTTP request for EducationSubmissionResource
func (r *EducationSubmissionResourceRequest) Do(method, path string, reqObj interface{}) (resObj *EducationSubmissionResource, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for EducationSubmissionResource
func (r *EducationSubmissionResourceRequest) Get() (*EducationSubmissionResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for EducationSubmissionResource
func (r *EducationSubmissionResourceRequest) Update(reqObj *EducationSubmissionResource) (*EducationSubmissionResource, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for EducationSubmissionResource
func (r *EducationSubmissionResourceRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

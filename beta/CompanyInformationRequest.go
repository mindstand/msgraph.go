// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CompanyInformationRequestBuilder is request builder for CompanyInformation
type CompanyInformationRequestBuilder struct{ BaseRequestBuilder }

// Request returns CompanyInformationRequest
func (b *CompanyInformationRequestBuilder) Request() *CompanyInformationRequest {
	return &CompanyInformationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CompanyInformationRequest is request for CompanyInformation
type CompanyInformationRequest struct{ BaseRequest }

// Do performs HTTP request for CompanyInformation
func (r *CompanyInformationRequest) Do(method, path string, reqObj interface{}) (resObj *CompanyInformation, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for CompanyInformation
func (r *CompanyInformationRequest) Get() (*CompanyInformation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for CompanyInformation
func (r *CompanyInformationRequest) Update(reqObj *CompanyInformation) (*CompanyInformation, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for CompanyInformation
func (r *CompanyInformationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

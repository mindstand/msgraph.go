// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OnenoteOperationRequestBuilder is request builder for OnenoteOperation
type OnenoteOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteOperationRequest
func (b *OnenoteOperationRequestBuilder) Request() *OnenoteOperationRequest {
	return &OnenoteOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenoteOperationRequest is request for OnenoteOperation
type OnenoteOperationRequest struct{ BaseRequest }

// Do performs HTTP request for OnenoteOperation
func (r *OnenoteOperationRequest) Do(method, path string, reqObj interface{}) (resObj *OnenoteOperation, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OnenoteOperation
func (r *OnenoteOperationRequest) Get() (*OnenoteOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OnenoteOperation
func (r *OnenoteOperationRequest) Update(reqObj *OnenoteOperation) (*OnenoteOperation, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OnenoteOperation
func (r *OnenoteOperationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ShiftRequestBuilder is request builder for Shift
type ShiftRequestBuilder struct{ BaseRequestBuilder }

// Request returns ShiftRequest
func (b *ShiftRequestBuilder) Request() *ShiftRequest {
	return &ShiftRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ShiftRequest is request for Shift
type ShiftRequest struct{ BaseRequest }

// Do performs HTTP request for Shift
func (r *ShiftRequest) Do(method, path string, reqObj interface{}) (resObj *Shift, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Shift
func (r *ShiftRequest) Get() (*Shift, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Shift
func (r *ShiftRequest) Update(reqObj *Shift) (*Shift, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Shift
func (r *ShiftRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

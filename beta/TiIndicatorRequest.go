// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TiIndicatorRequestBuilder is request builder for TiIndicator
type TiIndicatorRequestBuilder struct{ BaseRequestBuilder }

// Request returns TiIndicatorRequest
func (b *TiIndicatorRequestBuilder) Request() *TiIndicatorRequest {
	return &TiIndicatorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TiIndicatorRequest is request for TiIndicator
type TiIndicatorRequest struct{ BaseRequest }

// Do performs HTTP request for TiIndicator
func (r *TiIndicatorRequest) Do(method, path string, reqObj interface{}) (resObj *TiIndicator, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TiIndicator
func (r *TiIndicatorRequest) Get() (*TiIndicator, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TiIndicator
func (r *TiIndicatorRequest) Update(reqObj *TiIndicator) (*TiIndicator, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TiIndicator
func (r *TiIndicatorRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

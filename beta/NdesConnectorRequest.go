// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// NdesConnectorRequestBuilder is request builder for NdesConnector
type NdesConnectorRequestBuilder struct{ BaseRequestBuilder }

// Request returns NdesConnectorRequest
func (b *NdesConnectorRequestBuilder) Request() *NdesConnectorRequest {
	return &NdesConnectorRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NdesConnectorRequest is request for NdesConnector
type NdesConnectorRequest struct{ BaseRequest }

// Do performs HTTP request for NdesConnector
func (r *NdesConnectorRequest) Do(method, path string, reqObj interface{}) (resObj *NdesConnector, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for NdesConnector
func (r *NdesConnectorRequest) Get() (*NdesConnector, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for NdesConnector
func (r *NdesConnectorRequest) Update(reqObj *NdesConnector) (*NdesConnector, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for NdesConnector
func (r *NdesConnectorRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

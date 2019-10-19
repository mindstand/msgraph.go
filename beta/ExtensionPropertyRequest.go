// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ExtensionPropertyRequestBuilder is request builder for ExtensionProperty
type ExtensionPropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExtensionPropertyRequest
func (b *ExtensionPropertyRequestBuilder) Request() *ExtensionPropertyRequest {
	return &ExtensionPropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ExtensionPropertyRequest is request for ExtensionProperty
type ExtensionPropertyRequest struct{ BaseRequest }

// Do performs HTTP request for ExtensionProperty
func (r *ExtensionPropertyRequest) Do(method, path string, reqObj interface{}) (resObj *ExtensionProperty, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ExtensionProperty
func (r *ExtensionPropertyRequest) Get() (*ExtensionProperty, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ExtensionProperty
func (r *ExtensionPropertyRequest) Update(reqObj *ExtensionProperty) (*ExtensionProperty, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ExtensionProperty
func (r *ExtensionPropertyRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AgedAccountsPayableRequestBuilder is request builder for AgedAccountsPayable
type AgedAccountsPayableRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgedAccountsPayableRequest
func (b *AgedAccountsPayableRequestBuilder) Request() *AgedAccountsPayableRequest {
	return &AgedAccountsPayableRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgedAccountsPayableRequest is request for AgedAccountsPayable
type AgedAccountsPayableRequest struct{ BaseRequest }

// Do performs HTTP request for AgedAccountsPayable
func (r *AgedAccountsPayableRequest) Do(method, path string, reqObj interface{}) (resObj *AgedAccountsPayable, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AgedAccountsPayable
func (r *AgedAccountsPayableRequest) Get() (*AgedAccountsPayable, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AgedAccountsPayable
func (r *AgedAccountsPayableRequest) Update(reqObj *AgedAccountsPayable) (*AgedAccountsPayable, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AgedAccountsPayable
func (r *AgedAccountsPayableRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

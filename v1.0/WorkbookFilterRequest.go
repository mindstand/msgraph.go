// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookFilterRequestBuilder is request builder for WorkbookFilter
type WorkbookFilterRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookFilterRequest
func (b *WorkbookFilterRequestBuilder) Request() *WorkbookFilterRequest {
	return &WorkbookFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookFilterRequest is request for WorkbookFilter
type WorkbookFilterRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookFilter
func (r *WorkbookFilterRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookFilter, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookFilter
func (r *WorkbookFilterRequest) Get() (*WorkbookFilter, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookFilter
func (r *WorkbookFilterRequest) Update(reqObj *WorkbookFilter) (*WorkbookFilter, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookFilter
func (r *WorkbookFilterRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

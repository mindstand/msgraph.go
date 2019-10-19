// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookChartFillRequestBuilder is request builder for WorkbookChartFill
type WorkbookChartFillRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartFillRequest
func (b *WorkbookChartFillRequestBuilder) Request() *WorkbookChartFillRequest {
	return &WorkbookChartFillRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartFillRequest is request for WorkbookChartFill
type WorkbookChartFillRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookChartFill
func (r *WorkbookChartFillRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChartFill, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookChartFill
func (r *WorkbookChartFillRequest) Get() (*WorkbookChartFill, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookChartFill
func (r *WorkbookChartFillRequest) Update(reqObj *WorkbookChartFill) (*WorkbookChartFill, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookChartFill
func (r *WorkbookChartFillRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

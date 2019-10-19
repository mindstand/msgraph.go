// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookChartGridlinesRequestBuilder is request builder for WorkbookChartGridlines
type WorkbookChartGridlinesRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartGridlinesRequest
func (b *WorkbookChartGridlinesRequestBuilder) Request() *WorkbookChartGridlinesRequest {
	return &WorkbookChartGridlinesRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartGridlinesRequest is request for WorkbookChartGridlines
type WorkbookChartGridlinesRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookChartGridlines
func (r *WorkbookChartGridlinesRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChartGridlines, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookChartGridlines
func (r *WorkbookChartGridlinesRequest) Get() (*WorkbookChartGridlines, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookChartGridlines
func (r *WorkbookChartGridlinesRequest) Update(reqObj *WorkbookChartGridlines) (*WorkbookChartGridlines, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookChartGridlines
func (r *WorkbookChartGridlinesRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Format is navigation property
func (b *WorkbookChartGridlinesRequestBuilder) Format() *WorkbookChartGridlinesFormatRequestBuilder {
	bb := &WorkbookChartGridlinesFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/format"
	return bb
}

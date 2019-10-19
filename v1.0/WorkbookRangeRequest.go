// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookRangeRequestBuilder is request builder for WorkbookRange
type WorkbookRangeRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookRangeRequest
func (b *WorkbookRangeRequestBuilder) Request() *WorkbookRangeRequest {
	return &WorkbookRangeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookRangeRequest is request for WorkbookRange
type WorkbookRangeRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookRange
func (r *WorkbookRangeRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookRange, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookRange
func (r *WorkbookRangeRequest) Get() (*WorkbookRange, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookRange
func (r *WorkbookRangeRequest) Update(reqObj *WorkbookRange) (*WorkbookRange, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookRange
func (r *WorkbookRangeRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Format is navigation property
func (b *WorkbookRangeRequestBuilder) Format() *WorkbookRangeFormatRequestBuilder {
	bb := &WorkbookRangeFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/format"
	return bb
}

// Sort is navigation property
func (b *WorkbookRangeRequestBuilder) Sort() *WorkbookRangeSortRequestBuilder {
	bb := &WorkbookRangeSortRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sort"
	return bb
}

// Worksheet is navigation property
func (b *WorkbookRangeRequestBuilder) Worksheet() *WorkbookWorksheetRequestBuilder {
	bb := &WorkbookWorksheetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/worksheet"
	return bb
}

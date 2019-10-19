// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookChartAxisFormatRequestBuilder is request builder for WorkbookChartAxisFormat
type WorkbookChartAxisFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartAxisFormatRequest
func (b *WorkbookChartAxisFormatRequestBuilder) Request() *WorkbookChartAxisFormatRequest {
	return &WorkbookChartAxisFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartAxisFormatRequest is request for WorkbookChartAxisFormat
type WorkbookChartAxisFormatRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookChartAxisFormat
func (r *WorkbookChartAxisFormatRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChartAxisFormat, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookChartAxisFormat
func (r *WorkbookChartAxisFormatRequest) Get() (*WorkbookChartAxisFormat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookChartAxisFormat
func (r *WorkbookChartAxisFormatRequest) Update(reqObj *WorkbookChartAxisFormat) (*WorkbookChartAxisFormat, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookChartAxisFormat
func (r *WorkbookChartAxisFormatRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Font is navigation property
func (b *WorkbookChartAxisFormatRequestBuilder) Font() *WorkbookChartFontRequestBuilder {
	bb := &WorkbookChartFontRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/font"
	return bb
}

// Line is navigation property
func (b *WorkbookChartAxisFormatRequestBuilder) Line() *WorkbookChartLineFormatRequestBuilder {
	bb := &WorkbookChartLineFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/line"
	return bb
}

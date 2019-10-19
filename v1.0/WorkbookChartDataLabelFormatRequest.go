// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookChartDataLabelFormatRequestBuilder is request builder for WorkbookChartDataLabelFormat
type WorkbookChartDataLabelFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartDataLabelFormatRequest
func (b *WorkbookChartDataLabelFormatRequestBuilder) Request() *WorkbookChartDataLabelFormatRequest {
	return &WorkbookChartDataLabelFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookChartDataLabelFormatRequest is request for WorkbookChartDataLabelFormat
type WorkbookChartDataLabelFormatRequest struct{ BaseRequest }

// Do performs HTTP request for WorkbookChartDataLabelFormat
func (r *WorkbookChartDataLabelFormatRequest) Do(method, path string, reqObj interface{}) (resObj *WorkbookChartDataLabelFormat, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WorkbookChartDataLabelFormat
func (r *WorkbookChartDataLabelFormatRequest) Get() (*WorkbookChartDataLabelFormat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WorkbookChartDataLabelFormat
func (r *WorkbookChartDataLabelFormatRequest) Update(reqObj *WorkbookChartDataLabelFormat) (*WorkbookChartDataLabelFormat, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WorkbookChartDataLabelFormat
func (r *WorkbookChartDataLabelFormatRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Fill is navigation property
func (b *WorkbookChartDataLabelFormatRequestBuilder) Fill() *WorkbookChartFillRequestBuilder {
	bb := &WorkbookChartFillRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fill"
	return bb
}

// Font is navigation property
func (b *WorkbookChartDataLabelFormatRequestBuilder) Font() *WorkbookChartFontRequestBuilder {
	bb := &WorkbookChartFontRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/font"
	return bb
}

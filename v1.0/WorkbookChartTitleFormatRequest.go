// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookChartTitleFormatRequestBuilder is request builder for WorkbookChartTitleFormat
type WorkbookChartTitleFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartTitleFormatRequest
func (b *WorkbookChartTitleFormatRequestBuilder) Request() *WorkbookChartTitleFormatRequest {
	return &WorkbookChartTitleFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *WorkbookChartTitleFormatRequestBuilder) Delta() *WorkbookChartTitleFormatRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// WorkbookChartTitleFormatRequest is request for WorkbookChartTitleFormat
type WorkbookChartTitleFormatRequest struct{ BaseRequest }

// Get performs GET request for WorkbookChartTitleFormat
func (r *WorkbookChartTitleFormatRequest) Get(ctx context.Context) (resObj *WorkbookChartTitleFormat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookChartTitleFormat
func (r *WorkbookChartTitleFormatRequest) Update(ctx context.Context, reqObj *WorkbookChartTitleFormat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookChartTitleFormat
func (r *WorkbookChartTitleFormatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Fill is navigation property
func (b *WorkbookChartTitleFormatRequestBuilder) Fill() *WorkbookChartFillRequestBuilder {
	bb := &WorkbookChartFillRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fill"
	return bb
}

// Font is navigation property
func (b *WorkbookChartTitleFormatRequestBuilder) Font() *WorkbookChartFontRequestBuilder {
	bb := &WorkbookChartFontRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/font"
	return bb
}

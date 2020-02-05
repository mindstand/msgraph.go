// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookChartAreaFormatRequestBuilder is request builder for WorkbookChartAreaFormat
type WorkbookChartAreaFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookChartAreaFormatRequest
func (b *WorkbookChartAreaFormatRequestBuilder) Request() *WorkbookChartAreaFormatRequest {
	return &WorkbookChartAreaFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *WorkbookChartAreaFormatRequestBuilder) Delta() *WorkbookChartAreaFormatRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// WorkbookChartAreaFormatRequest is request for WorkbookChartAreaFormat
type WorkbookChartAreaFormatRequest struct{ BaseRequest }

// Get performs GET request for WorkbookChartAreaFormat
func (r *WorkbookChartAreaFormatRequest) Get(ctx context.Context) (resObj *WorkbookChartAreaFormat, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookChartAreaFormat
func (r *WorkbookChartAreaFormatRequest) Update(ctx context.Context, reqObj *WorkbookChartAreaFormat) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookChartAreaFormat
func (r *WorkbookChartAreaFormatRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Fill is navigation property
func (b *WorkbookChartAreaFormatRequestBuilder) Fill() *WorkbookChartFillRequestBuilder {
	bb := &WorkbookChartFillRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/fill"
	return bb
}

// Font is navigation property
func (b *WorkbookChartAreaFormatRequestBuilder) Font() *WorkbookChartFontRequestBuilder {
	bb := &WorkbookChartFontRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/font"
	return bb
}

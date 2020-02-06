// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

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

// Get performs GET request for WorkbookChartFill
func (r *WorkbookChartFillRequest) Get(ctx context.Context) (resObj *WorkbookChartFill, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookChartFill
func (r *WorkbookChartFillRequest) Update(ctx context.Context, reqObj *WorkbookChartFill) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookChartFill
func (r *WorkbookChartFillRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

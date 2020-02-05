// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookRangeSortRequestBuilder is request builder for WorkbookRangeSort
type WorkbookRangeSortRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookRangeSortRequest
func (b *WorkbookRangeSortRequestBuilder) Request() *WorkbookRangeSortRequest {
	return &WorkbookRangeSortRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *WorkbookRangeSortRequestBuilder) Delta() *WorkbookRangeSortRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// WorkbookRangeSortRequest is request for WorkbookRangeSort
type WorkbookRangeSortRequest struct{ BaseRequest }

// Get performs GET request for WorkbookRangeSort
func (r *WorkbookRangeSortRequest) Get(ctx context.Context) (resObj *WorkbookRangeSort, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookRangeSort
func (r *WorkbookRangeSortRequest) Update(ctx context.Context, reqObj *WorkbookRangeSort) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookRangeSort
func (r *WorkbookRangeSortRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookApplicationRequestBuilder is request builder for WorkbookApplication
type WorkbookApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookApplicationRequest
func (b *WorkbookApplicationRequestBuilder) Request() *WorkbookApplicationRequest {
	return &WorkbookApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *WorkbookApplicationRequestBuilder) Delta() *WorkbookApplicationRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// WorkbookApplicationRequest is request for WorkbookApplication
type WorkbookApplicationRequest struct{ BaseRequest }

// Get performs GET request for WorkbookApplication
func (r *WorkbookApplicationRequest) Get(ctx context.Context) (resObj *WorkbookApplication, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookApplication
func (r *WorkbookApplicationRequest) Update(ctx context.Context, reqObj *WorkbookApplication) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookApplication
func (r *WorkbookApplicationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

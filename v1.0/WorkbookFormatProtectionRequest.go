// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookFormatProtectionRequestBuilder is request builder for WorkbookFormatProtection
type WorkbookFormatProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookFormatProtectionRequest
func (b *WorkbookFormatProtectionRequestBuilder) Request() *WorkbookFormatProtectionRequest {
	return &WorkbookFormatProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookFormatProtectionRequest is request for WorkbookFormatProtection
type WorkbookFormatProtectionRequest struct{ BaseRequest }

// Get performs GET request for WorkbookFormatProtection
func (r *WorkbookFormatProtectionRequest) Get(ctx context.Context) (resObj *WorkbookFormatProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookFormatProtection
func (r *WorkbookFormatProtectionRequest) Update(ctx context.Context, reqObj *WorkbookFormatProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookFormatProtection
func (r *WorkbookFormatProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

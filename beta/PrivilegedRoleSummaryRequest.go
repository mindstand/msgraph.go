// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PrivilegedRoleSummaryRequestBuilder is request builder for PrivilegedRoleSummary
type PrivilegedRoleSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleSummaryRequest
func (b *PrivilegedRoleSummaryRequestBuilder) Request() *PrivilegedRoleSummaryRequest {
	return &PrivilegedRoleSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *PrivilegedRoleSummaryRequestBuilder) Delta() *PrivilegedRoleSummaryRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// PrivilegedRoleSummaryRequest is request for PrivilegedRoleSummary
type PrivilegedRoleSummaryRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Get(ctx context.Context) (resObj *PrivilegedRoleSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Update(ctx context.Context, reqObj *PrivilegedRoleSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagedAppStatusRequestBuilder is request builder for ManagedAppStatus
type ManagedAppStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppStatusRequest
func (b *ManagedAppStatusRequestBuilder) Request() *ManagedAppStatusRequest {
	return &ManagedAppStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ManagedAppStatusRequestBuilder) Delta() *ManagedAppStatusRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ManagedAppStatusRequest is request for ManagedAppStatus
type ManagedAppStatusRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Get(ctx context.Context) (resObj *ManagedAppStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Update(ctx context.Context, reqObj *ManagedAppStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

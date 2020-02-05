// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagedAppOperationRequestBuilder is request builder for ManagedAppOperation
type ManagedAppOperationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppOperationRequest
func (b *ManagedAppOperationRequestBuilder) Request() *ManagedAppOperationRequest {
	return &ManagedAppOperationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ManagedAppOperationRequestBuilder) Delta() *ManagedAppOperationRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ManagedAppOperationRequest is request for ManagedAppOperation
type ManagedAppOperationRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppOperation
func (r *ManagedAppOperationRequest) Get(ctx context.Context) (resObj *ManagedAppOperation, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppOperation
func (r *ManagedAppOperationRequest) Update(ctx context.Context, reqObj *ManagedAppOperation) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppOperation
func (r *ManagedAppOperationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

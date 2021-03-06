// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// OpenShiftRequestBuilder is request builder for OpenShift
type OpenShiftRequestBuilder struct{ BaseRequestBuilder }

// Request returns OpenShiftRequest
func (b *OpenShiftRequestBuilder) Request() *OpenShiftRequest {
	return &OpenShiftRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OpenShiftRequest is request for OpenShift
type OpenShiftRequest struct{ BaseRequest }

// Get performs GET request for OpenShift
func (r *OpenShiftRequest) Get(ctx context.Context) (resObj *OpenShift, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OpenShift
func (r *OpenShiftRequest) Update(ctx context.Context, reqObj *OpenShift) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OpenShift
func (r *OpenShiftRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

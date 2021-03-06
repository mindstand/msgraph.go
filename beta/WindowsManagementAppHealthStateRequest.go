// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsManagementAppHealthStateRequestBuilder is request builder for WindowsManagementAppHealthState
type WindowsManagementAppHealthStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsManagementAppHealthStateRequest
func (b *WindowsManagementAppHealthStateRequestBuilder) Request() *WindowsManagementAppHealthStateRequest {
	return &WindowsManagementAppHealthStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsManagementAppHealthStateRequest is request for WindowsManagementAppHealthState
type WindowsManagementAppHealthStateRequest struct{ BaseRequest }

// Get performs GET request for WindowsManagementAppHealthState
func (r *WindowsManagementAppHealthStateRequest) Get(ctx context.Context) (resObj *WindowsManagementAppHealthState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsManagementAppHealthState
func (r *WindowsManagementAppHealthStateRequest) Update(ctx context.Context, reqObj *WindowsManagementAppHealthState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsManagementAppHealthState
func (r *WindowsManagementAppHealthStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

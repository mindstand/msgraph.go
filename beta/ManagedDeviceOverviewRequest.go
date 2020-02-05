// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagedDeviceOverviewRequestBuilder is request builder for ManagedDeviceOverview
type ManagedDeviceOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceOverviewRequest
func (b *ManagedDeviceOverviewRequestBuilder) Request() *ManagedDeviceOverviewRequest {
	return &ManagedDeviceOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ManagedDeviceOverviewRequestBuilder) Delta() *ManagedDeviceOverviewRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ManagedDeviceOverviewRequest is request for ManagedDeviceOverview
type ManagedDeviceOverviewRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceOverview
func (r *ManagedDeviceOverviewRequest) Get(ctx context.Context) (resObj *ManagedDeviceOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceOverview
func (r *ManagedDeviceOverviewRequest) Update(ctx context.Context, reqObj *ManagedDeviceOverview) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceOverview
func (r *ManagedDeviceOverviewRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

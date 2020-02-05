// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceConfigurationUserStatusRequestBuilder is request builder for DeviceConfigurationUserStatus
type DeviceConfigurationUserStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationUserStatusRequest
func (b *DeviceConfigurationUserStatusRequestBuilder) Request() *DeviceConfigurationUserStatusRequest {
	return &DeviceConfigurationUserStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *DeviceConfigurationUserStatusRequestBuilder) Delta() *DeviceConfigurationUserStatusRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// DeviceConfigurationUserStatusRequest is request for DeviceConfigurationUserStatus
type DeviceConfigurationUserStatusRequest struct{ BaseRequest }

// Get performs GET request for DeviceConfigurationUserStatus
func (r *DeviceConfigurationUserStatusRequest) Get(ctx context.Context) (resObj *DeviceConfigurationUserStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceConfigurationUserStatus
func (r *DeviceConfigurationUserStatusRequest) Update(ctx context.Context, reqObj *DeviceConfigurationUserStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceConfigurationUserStatus
func (r *DeviceConfigurationUserStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

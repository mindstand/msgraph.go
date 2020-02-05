// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DeviceConfigurationStateRequestBuilder is request builder for DeviceConfigurationState
type DeviceConfigurationStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationStateRequest
func (b *DeviceConfigurationStateRequestBuilder) Request() *DeviceConfigurationStateRequest {
	return &DeviceConfigurationStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *DeviceConfigurationStateRequestBuilder) Delta() *DeviceConfigurationStateRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// DeviceConfigurationStateRequest is request for DeviceConfigurationState
type DeviceConfigurationStateRequest struct{ BaseRequest }

// Get performs GET request for DeviceConfigurationState
func (r *DeviceConfigurationStateRequest) Get(ctx context.Context) (resObj *DeviceConfigurationState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceConfigurationState
func (r *DeviceConfigurationStateRequest) Update(ctx context.Context, reqObj *DeviceConfigurationState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceConfigurationState
func (r *DeviceConfigurationStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

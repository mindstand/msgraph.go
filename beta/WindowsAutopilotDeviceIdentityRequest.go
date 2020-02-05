// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WindowsAutopilotDeviceIdentityRequestBuilder is request builder for WindowsAutopilotDeviceIdentity
type WindowsAutopilotDeviceIdentityRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsAutopilotDeviceIdentityRequest
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) Request() *WindowsAutopilotDeviceIdentityRequest {
	return &WindowsAutopilotDeviceIdentityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) Delta() *WindowsAutopilotDeviceIdentityRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// WindowsAutopilotDeviceIdentityRequest is request for WindowsAutopilotDeviceIdentity
type WindowsAutopilotDeviceIdentityRequest struct{ BaseRequest }

// Get performs GET request for WindowsAutopilotDeviceIdentity
func (r *WindowsAutopilotDeviceIdentityRequest) Get(ctx context.Context) (resObj *WindowsAutopilotDeviceIdentity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsAutopilotDeviceIdentity
func (r *WindowsAutopilotDeviceIdentityRequest) Update(ctx context.Context, reqObj *WindowsAutopilotDeviceIdentity) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsAutopilotDeviceIdentity
func (r *WindowsAutopilotDeviceIdentityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DeploymentProfile is navigation property
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) DeploymentProfile() *WindowsAutopilotDeploymentProfileRequestBuilder {
	bb := &WindowsAutopilotDeploymentProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deploymentProfile"
	return bb
}

// IntendedDeploymentProfile is navigation property
func (b *WindowsAutopilotDeviceIdentityRequestBuilder) IntendedDeploymentProfile() *WindowsAutopilotDeploymentProfileRequestBuilder {
	bb := &WindowsAutopilotDeploymentProfileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/intendedDeploymentProfile"
	return bb
}

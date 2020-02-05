// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder is request builder for AdvancedThreatProtectionOnboardingDeviceSettingState
type AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdvancedThreatProtectionOnboardingDeviceSettingStateRequest
func (b *AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder) Request() *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest {
	return &AdvancedThreatProtectionOnboardingDeviceSettingStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder) Delta() *AdvancedThreatProtectionOnboardingDeviceSettingStateRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// AdvancedThreatProtectionOnboardingDeviceSettingStateRequest is request for AdvancedThreatProtectionOnboardingDeviceSettingState
type AdvancedThreatProtectionOnboardingDeviceSettingStateRequest struct{ BaseRequest }

// Get performs GET request for AdvancedThreatProtectionOnboardingDeviceSettingState
func (r *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest) Get(ctx context.Context) (resObj *AdvancedThreatProtectionOnboardingDeviceSettingState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AdvancedThreatProtectionOnboardingDeviceSettingState
func (r *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest) Update(ctx context.Context, reqObj *AdvancedThreatProtectionOnboardingDeviceSettingState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AdvancedThreatProtectionOnboardingDeviceSettingState
func (r *AdvancedThreatProtectionOnboardingDeviceSettingStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ManagedAppProtectionTargetAppsRequestParameter undocumented
type ManagedAppProtectionTargetAppsRequestParameter struct {
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
}

//
type ManagedAppProtectionTargetAppsRequestBuilder struct{ BaseRequestBuilder }

// TargetApps action undocumented
func (b *ManagedAppProtectionRequestBuilder) TargetApps(reqObj *ManagedAppProtectionTargetAppsRequestParameter) *ManagedAppProtectionTargetAppsRequestBuilder {
	bb := &ManagedAppProtectionTargetAppsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/targetApps"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type ManagedAppProtectionTargetAppsRequest struct{ BaseRequest }

//
func (b *ManagedAppProtectionTargetAppsRequestBuilder) Request() *ManagedAppProtectionTargetAppsRequest {
	return &ManagedAppProtectionTargetAppsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *ManagedAppProtectionTargetAppsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

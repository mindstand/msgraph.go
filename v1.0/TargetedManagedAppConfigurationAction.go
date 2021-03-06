// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// TargetedManagedAppConfigurationAssignRequestParameter undocumented
type TargetedManagedAppConfigurationAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []TargetedManagedAppPolicyAssignment `json:"assignments,omitempty"`
}

// TargetedManagedAppConfigurationTargetAppsRequestParameter undocumented
type TargetedManagedAppConfigurationTargetAppsRequestParameter struct {
	// Apps undocumented
	Apps []ManagedMobileApp `json:"apps,omitempty"`
}

//
type TargetedManagedAppConfigurationAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *TargetedManagedAppConfigurationRequestBuilder) Assign(reqObj *TargetedManagedAppConfigurationAssignRequestParameter) *TargetedManagedAppConfigurationAssignRequestBuilder {
	bb := &TargetedManagedAppConfigurationAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppConfigurationAssignRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppConfigurationAssignRequestBuilder) Request() *TargetedManagedAppConfigurationAssignRequest {
	return &TargetedManagedAppConfigurationAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppConfigurationAssignRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

//
type TargetedManagedAppConfigurationTargetAppsRequestBuilder struct{ BaseRequestBuilder }

// TargetApps action undocumented
func (b *TargetedManagedAppConfigurationRequestBuilder) TargetApps(reqObj *TargetedManagedAppConfigurationTargetAppsRequestParameter) *TargetedManagedAppConfigurationTargetAppsRequestBuilder {
	bb := &TargetedManagedAppConfigurationTargetAppsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/targetApps"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type TargetedManagedAppConfigurationTargetAppsRequest struct{ BaseRequest }

//
func (b *TargetedManagedAppConfigurationTargetAppsRequestBuilder) Request() *TargetedManagedAppConfigurationTargetAppsRequest {
	return &TargetedManagedAppConfigurationTargetAppsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *TargetedManagedAppConfigurationTargetAppsRequest) Post(ctx context.Context) error {
	return r.JSONRequest(ctx, "POST", "", r.requestObject, nil)
}

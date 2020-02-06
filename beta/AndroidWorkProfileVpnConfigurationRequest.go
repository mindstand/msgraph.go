// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidWorkProfileVpnConfigurationRequestBuilder is request builder for AndroidWorkProfileVpnConfiguration
type AndroidWorkProfileVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidWorkProfileVpnConfigurationRequest
func (b *AndroidWorkProfileVpnConfigurationRequestBuilder) Request() *AndroidWorkProfileVpnConfigurationRequest {
	return &AndroidWorkProfileVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidWorkProfileVpnConfigurationRequest is request for AndroidWorkProfileVpnConfiguration
type AndroidWorkProfileVpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AndroidWorkProfileVpnConfiguration
func (r *AndroidWorkProfileVpnConfigurationRequest) Get(ctx context.Context) (resObj *AndroidWorkProfileVpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidWorkProfileVpnConfiguration
func (r *AndroidWorkProfileVpnConfigurationRequest) Update(ctx context.Context, reqObj *AndroidWorkProfileVpnConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidWorkProfileVpnConfiguration
func (r *AndroidWorkProfileVpnConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificate is navigation property
func (b *AndroidWorkProfileVpnConfigurationRequestBuilder) IdentityCertificate() *AndroidWorkProfileCertificateProfileBaseRequestBuilder {
	bb := &AndroidWorkProfileCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}

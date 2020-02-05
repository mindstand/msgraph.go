// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// MacOSVpnConfigurationRequestBuilder is request builder for MacOSVpnConfiguration
type MacOSVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSVpnConfigurationRequest
func (b *MacOSVpnConfigurationRequestBuilder) Request() *MacOSVpnConfigurationRequest {
	return &MacOSVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *MacOSVpnConfigurationRequestBuilder) Delta() *MacOSVpnConfigurationRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// MacOSVpnConfigurationRequest is request for MacOSVpnConfiguration
type MacOSVpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Get(ctx context.Context) (resObj *MacOSVpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Update(ctx context.Context, reqObj *MacOSVpnConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificate is navigation property
func (b *MacOSVpnConfigurationRequestBuilder) IdentityCertificate() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}

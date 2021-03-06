// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// IOSDeviceFeaturesConfigurationRequestBuilder is request builder for IOSDeviceFeaturesConfiguration
type IOSDeviceFeaturesConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSDeviceFeaturesConfigurationRequest
func (b *IOSDeviceFeaturesConfigurationRequestBuilder) Request() *IOSDeviceFeaturesConfigurationRequest {
	return &IOSDeviceFeaturesConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSDeviceFeaturesConfigurationRequest is request for IOSDeviceFeaturesConfiguration
type IOSDeviceFeaturesConfigurationRequest struct{ BaseRequest }

// Get performs GET request for IOSDeviceFeaturesConfiguration
func (r *IOSDeviceFeaturesConfigurationRequest) Get(ctx context.Context) (resObj *IOSDeviceFeaturesConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSDeviceFeaturesConfiguration
func (r *IOSDeviceFeaturesConfigurationRequest) Update(ctx context.Context, reqObj *IOSDeviceFeaturesConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSDeviceFeaturesConfiguration
func (r *IOSDeviceFeaturesConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *IOSDeviceFeaturesConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *IOSCertificateProfileBaseRequestBuilder {
	bb := &IOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// SingleSignOnExtensionPkinitCertificate is navigation property
func (b *IOSDeviceFeaturesConfigurationRequestBuilder) SingleSignOnExtensionPkinitCertificate() *IOSCertificateProfileBaseRequestBuilder {
	bb := &IOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/singleSignOnExtensionPkinitCertificate"
	return bb
}

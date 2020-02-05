// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AndroidForWorkEnterpriseWiFiConfigurationRequestBuilder is request builder for AndroidForWorkEnterpriseWiFiConfiguration
type AndroidForWorkEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkEnterpriseWiFiConfigurationRequest
func (b *AndroidForWorkEnterpriseWiFiConfigurationRequestBuilder) Request() *AndroidForWorkEnterpriseWiFiConfigurationRequest {
	return &AndroidForWorkEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *AndroidForWorkEnterpriseWiFiConfigurationRequestBuilder) Delta() *AndroidForWorkEnterpriseWiFiConfigurationRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// AndroidForWorkEnterpriseWiFiConfigurationRequest is request for AndroidForWorkEnterpriseWiFiConfiguration
type AndroidForWorkEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Get performs GET request for AndroidForWorkEnterpriseWiFiConfiguration
func (r *AndroidForWorkEnterpriseWiFiConfigurationRequest) Get(ctx context.Context) (resObj *AndroidForWorkEnterpriseWiFiConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidForWorkEnterpriseWiFiConfiguration
func (r *AndroidForWorkEnterpriseWiFiConfigurationRequest) Update(ctx context.Context, reqObj *AndroidForWorkEnterpriseWiFiConfiguration) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidForWorkEnterpriseWiFiConfiguration
func (r *AndroidForWorkEnterpriseWiFiConfigurationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *AndroidForWorkEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *AndroidForWorkCertificateProfileBaseRequestBuilder {
	bb := &AndroidForWorkCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *AndroidForWorkEnterpriseWiFiConfigurationRequestBuilder) RootCertificateForServerValidation() *AndroidForWorkTrustedRootCertificateRequestBuilder {
	bb := &AndroidForWorkTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}

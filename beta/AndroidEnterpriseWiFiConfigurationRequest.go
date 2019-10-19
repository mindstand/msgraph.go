// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidEnterpriseWiFiConfigurationRequestBuilder is request builder for AndroidEnterpriseWiFiConfiguration
type AndroidEnterpriseWiFiConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidEnterpriseWiFiConfigurationRequest
func (b *AndroidEnterpriseWiFiConfigurationRequestBuilder) Request() *AndroidEnterpriseWiFiConfigurationRequest {
	return &AndroidEnterpriseWiFiConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidEnterpriseWiFiConfigurationRequest is request for AndroidEnterpriseWiFiConfiguration
type AndroidEnterpriseWiFiConfigurationRequest struct{ BaseRequest }

// Do performs HTTP request for AndroidEnterpriseWiFiConfiguration
func (r *AndroidEnterpriseWiFiConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidEnterpriseWiFiConfiguration, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidEnterpriseWiFiConfiguration
func (r *AndroidEnterpriseWiFiConfigurationRequest) Get() (*AndroidEnterpriseWiFiConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidEnterpriseWiFiConfiguration
func (r *AndroidEnterpriseWiFiConfigurationRequest) Update(reqObj *AndroidEnterpriseWiFiConfiguration) (*AndroidEnterpriseWiFiConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidEnterpriseWiFiConfiguration
func (r *AndroidEnterpriseWiFiConfigurationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// IdentityCertificateForClientAuthentication is navigation property
func (b *AndroidEnterpriseWiFiConfigurationRequestBuilder) IdentityCertificateForClientAuthentication() *AndroidCertificateProfileBaseRequestBuilder {
	bb := &AndroidCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificateForClientAuthentication"
	return bb
}

// RootCertificateForServerValidation is navigation property
func (b *AndroidEnterpriseWiFiConfigurationRequestBuilder) RootCertificateForServerValidation() *AndroidTrustedRootCertificateRequestBuilder {
	bb := &AndroidTrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificateForServerValidation"
	return bb
}

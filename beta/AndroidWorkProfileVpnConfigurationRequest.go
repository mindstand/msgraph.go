// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

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

// Do performs HTTP request for AndroidWorkProfileVpnConfiguration
func (r *AndroidWorkProfileVpnConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *AndroidWorkProfileVpnConfiguration, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AndroidWorkProfileVpnConfiguration
func (r *AndroidWorkProfileVpnConfigurationRequest) Get() (*AndroidWorkProfileVpnConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AndroidWorkProfileVpnConfiguration
func (r *AndroidWorkProfileVpnConfigurationRequest) Update(reqObj *AndroidWorkProfileVpnConfiguration) (*AndroidWorkProfileVpnConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AndroidWorkProfileVpnConfiguration
func (r *AndroidWorkProfileVpnConfigurationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// IdentityCertificate is navigation property
func (b *AndroidWorkProfileVpnConfigurationRequestBuilder) IdentityCertificate() *AndroidWorkProfileCertificateProfileBaseRequestBuilder {
	bb := &AndroidWorkProfileCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}

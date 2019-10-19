// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSDerivedCredentialAuthenticationConfigurationRequestBuilder is request builder for IOSDerivedCredentialAuthenticationConfiguration
type IOSDerivedCredentialAuthenticationConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSDerivedCredentialAuthenticationConfigurationRequest
func (b *IOSDerivedCredentialAuthenticationConfigurationRequestBuilder) Request() *IOSDerivedCredentialAuthenticationConfigurationRequest {
	return &IOSDerivedCredentialAuthenticationConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSDerivedCredentialAuthenticationConfigurationRequest is request for IOSDerivedCredentialAuthenticationConfiguration
type IOSDerivedCredentialAuthenticationConfigurationRequest struct{ BaseRequest }

// Do performs HTTP request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Do(method, path string, reqObj interface{}) (resObj *IOSDerivedCredentialAuthenticationConfiguration, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Get() (*IOSDerivedCredentialAuthenticationConfiguration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Update(reqObj *IOSDerivedCredentialAuthenticationConfiguration) (*IOSDerivedCredentialAuthenticationConfiguration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for IOSDerivedCredentialAuthenticationConfiguration
func (r *IOSDerivedCredentialAuthenticationConfigurationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DerivedCredentialSettings is navigation property
func (b *IOSDerivedCredentialAuthenticationConfigurationRequestBuilder) DerivedCredentialSettings() *DeviceManagementDerivedCredentialSettingsRequestBuilder {
	bb := &DeviceManagementDerivedCredentialSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/derivedCredentialSettings"
	return bb
}
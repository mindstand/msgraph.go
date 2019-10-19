// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsInformationProtectionDeviceRegistrationRequestBuilder is request builder for WindowsInformationProtectionDeviceRegistration
type WindowsInformationProtectionDeviceRegistrationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsInformationProtectionDeviceRegistrationRequest
func (b *WindowsInformationProtectionDeviceRegistrationRequestBuilder) Request() *WindowsInformationProtectionDeviceRegistrationRequest {
	return &WindowsInformationProtectionDeviceRegistrationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsInformationProtectionDeviceRegistrationRequest is request for WindowsInformationProtectionDeviceRegistration
type WindowsInformationProtectionDeviceRegistrationRequest struct{ BaseRequest }

// Do performs HTTP request for WindowsInformationProtectionDeviceRegistration
func (r *WindowsInformationProtectionDeviceRegistrationRequest) Do(method, path string, reqObj interface{}) (resObj *WindowsInformationProtectionDeviceRegistration, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for WindowsInformationProtectionDeviceRegistration
func (r *WindowsInformationProtectionDeviceRegistrationRequest) Get() (*WindowsInformationProtectionDeviceRegistration, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for WindowsInformationProtectionDeviceRegistration
func (r *WindowsInformationProtectionDeviceRegistrationRequest) Update(reqObj *WindowsInformationProtectionDeviceRegistration) (*WindowsInformationProtectionDeviceRegistration, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for WindowsInformationProtectionDeviceRegistration
func (r *WindowsInformationProtectionDeviceRegistrationRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

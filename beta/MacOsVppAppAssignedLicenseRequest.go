// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOsVppAppAssignedLicenseRequestBuilder is request builder for MacOsVppAppAssignedLicense
type MacOsVppAppAssignedLicenseRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOsVppAppAssignedLicenseRequest
func (b *MacOsVppAppAssignedLicenseRequestBuilder) Request() *MacOsVppAppAssignedLicenseRequest {
	return &MacOsVppAppAssignedLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOsVppAppAssignedLicenseRequest is request for MacOsVppAppAssignedLicense
type MacOsVppAppAssignedLicenseRequest struct{ BaseRequest }

// Do performs HTTP request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Do(method, path string, reqObj interface{}) (resObj *MacOsVppAppAssignedLicense, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Get() (*MacOsVppAppAssignedLicense, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Update(reqObj *MacOsVppAppAssignedLicense) (*MacOsVppAppAssignedLicense, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for MacOsVppAppAssignedLicense
func (r *MacOsVppAppAssignedLicenseRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

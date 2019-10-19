// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationAssignmentRequestBuilder is request builder for DeviceConfigurationAssignment
type DeviceConfigurationAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationAssignmentRequest
func (b *DeviceConfigurationAssignmentRequestBuilder) Request() *DeviceConfigurationAssignmentRequest {
	return &DeviceConfigurationAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceConfigurationAssignmentRequest is request for DeviceConfigurationAssignment
type DeviceConfigurationAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceConfigurationAssignment
func (r *DeviceConfigurationAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceConfigurationAssignment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceConfigurationAssignment
func (r *DeviceConfigurationAssignmentRequest) Get() (*DeviceConfigurationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceConfigurationAssignment
func (r *DeviceConfigurationAssignmentRequest) Update(reqObj *DeviceConfigurationAssignment) (*DeviceConfigurationAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceConfigurationAssignment
func (r *DeviceConfigurationAssignmentRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

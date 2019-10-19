// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedDeviceEncryptionStateRequestBuilder is request builder for ManagedDeviceEncryptionState
type ManagedDeviceEncryptionStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceEncryptionStateRequest
func (b *ManagedDeviceEncryptionStateRequestBuilder) Request() *ManagedDeviceEncryptionStateRequest {
	return &ManagedDeviceEncryptionStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceEncryptionStateRequest is request for ManagedDeviceEncryptionState
type ManagedDeviceEncryptionStateRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedDeviceEncryptionState
func (r *ManagedDeviceEncryptionStateRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedDeviceEncryptionState, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ManagedDeviceEncryptionState
func (r *ManagedDeviceEncryptionStateRequest) Get() (*ManagedDeviceEncryptionState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ManagedDeviceEncryptionState
func (r *ManagedDeviceEncryptionStateRequest) Update(reqObj *ManagedDeviceEncryptionState) (*ManagedDeviceEncryptionState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ManagedDeviceEncryptionState
func (r *ManagedDeviceEncryptionStateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
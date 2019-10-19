// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementPartnerRequestBuilder is request builder for DeviceManagementPartner
type DeviceManagementPartnerRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementPartnerRequest
func (b *DeviceManagementPartnerRequestBuilder) Request() *DeviceManagementPartnerRequest {
	return &DeviceManagementPartnerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementPartnerRequest is request for DeviceManagementPartner
type DeviceManagementPartnerRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementPartner
func (r *DeviceManagementPartnerRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementPartner, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementPartner
func (r *DeviceManagementPartnerRequest) Get() (*DeviceManagementPartner, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementPartner
func (r *DeviceManagementPartnerRequest) Update(reqObj *DeviceManagementPartner) (*DeviceManagementPartner, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementPartner
func (r *DeviceManagementPartnerRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementTroubleshootingEventRequestBuilder is request builder for DeviceManagementTroubleshootingEvent
type DeviceManagementTroubleshootingEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementTroubleshootingEventRequest
func (b *DeviceManagementTroubleshootingEventRequestBuilder) Request() *DeviceManagementTroubleshootingEventRequest {
	return &DeviceManagementTroubleshootingEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementTroubleshootingEventRequest is request for DeviceManagementTroubleshootingEvent
type DeviceManagementTroubleshootingEventRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementTroubleshootingEvent
func (r *DeviceManagementTroubleshootingEventRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementTroubleshootingEvent, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementTroubleshootingEvent
func (r *DeviceManagementTroubleshootingEventRequest) Get() (*DeviceManagementTroubleshootingEvent, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementTroubleshootingEvent
func (r *DeviceManagementTroubleshootingEventRequest) Update(reqObj *DeviceManagementTroubleshootingEvent) (*DeviceManagementTroubleshootingEvent, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementTroubleshootingEvent
func (r *DeviceManagementTroubleshootingEventRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
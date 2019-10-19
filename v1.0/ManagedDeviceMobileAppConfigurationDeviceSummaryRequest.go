// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationDeviceSummary
type ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationDeviceSummaryRequest
func (b *ManagedDeviceMobileAppConfigurationDeviceSummaryRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest {
	return &ManagedDeviceMobileAppConfigurationDeviceSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationDeviceSummaryRequest is request for ManagedDeviceMobileAppConfigurationDeviceSummary
type ManagedDeviceMobileAppConfigurationDeviceSummaryRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedDeviceMobileAppConfigurationDeviceSummary, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Get() (*ManagedDeviceMobileAppConfigurationDeviceSummary, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Update(reqObj *ManagedDeviceMobileAppConfigurationDeviceSummary) (*ManagedDeviceMobileAppConfigurationDeviceSummary, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationDeviceSummary
func (r *ManagedDeviceMobileAppConfigurationDeviceSummaryRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

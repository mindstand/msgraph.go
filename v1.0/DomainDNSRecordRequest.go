// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DomainDNSRecordRequestBuilder is request builder for DomainDNSRecord
type DomainDNSRecordRequestBuilder struct{ BaseRequestBuilder }

// Request returns DomainDNSRecordRequest
func (b *DomainDNSRecordRequestBuilder) Request() *DomainDNSRecordRequest {
	return &DomainDNSRecordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DomainDNSRecordRequest is request for DomainDNSRecord
type DomainDNSRecordRequest struct{ BaseRequest }

// Do performs HTTP request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Do(method, path string, reqObj interface{}) (resObj *DomainDNSRecord, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Get() (*DomainDNSRecord, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Update(reqObj *DomainDNSRecord) (*DomainDNSRecord, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

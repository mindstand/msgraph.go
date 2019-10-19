// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ScheduleChangeRequestObjectRequestBuilder is request builder for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns ScheduleChangeRequestObjectRequest
func (b *ScheduleChangeRequestObjectRequestBuilder) Request() *ScheduleChangeRequestObjectRequest {
	return &ScheduleChangeRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ScheduleChangeRequestObjectRequest is request for ScheduleChangeRequestObject
type ScheduleChangeRequestObjectRequest struct{ BaseRequest }

// Do performs HTTP request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Do(method, path string, reqObj interface{}) (resObj *ScheduleChangeRequestObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Get() (*ScheduleChangeRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Update(reqObj *ScheduleChangeRequestObject) (*ScheduleChangeRequestObject, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ScheduleChangeRequestObject
func (r *ScheduleChangeRequestObjectRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

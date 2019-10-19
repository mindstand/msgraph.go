// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupLifecyclePolicyRequestBuilder is request builder for GroupLifecyclePolicy
type GroupLifecyclePolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupLifecyclePolicyRequest
func (b *GroupLifecyclePolicyRequestBuilder) Request() *GroupLifecyclePolicyRequest {
	return &GroupLifecyclePolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupLifecyclePolicyRequest is request for GroupLifecyclePolicy
type GroupLifecyclePolicyRequest struct{ BaseRequest }

// Do performs HTTP request for GroupLifecyclePolicy
func (r *GroupLifecyclePolicyRequest) Do(method, path string, reqObj interface{}) (resObj *GroupLifecyclePolicy, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for GroupLifecyclePolicy
func (r *GroupLifecyclePolicyRequest) Get() (*GroupLifecyclePolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for GroupLifecyclePolicy
func (r *GroupLifecyclePolicyRequest) Update(reqObj *GroupLifecyclePolicy) (*GroupLifecyclePolicy, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for GroupLifecyclePolicy
func (r *GroupLifecyclePolicyRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

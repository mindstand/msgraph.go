// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TargetedManagedAppPolicyAssignmentRequestBuilder is request builder for TargetedManagedAppPolicyAssignment
type TargetedManagedAppPolicyAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppPolicyAssignmentRequest
func (b *TargetedManagedAppPolicyAssignmentRequestBuilder) Request() *TargetedManagedAppPolicyAssignmentRequest {
	return &TargetedManagedAppPolicyAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppPolicyAssignmentRequest is request for TargetedManagedAppPolicyAssignment
type TargetedManagedAppPolicyAssignmentRequest struct{ BaseRequest }

// Do performs HTTP request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Do(method, path string, reqObj interface{}) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Get() (*TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Update(reqObj *TargetedManagedAppPolicyAssignment) (*TargetedManagedAppPolicyAssignment, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TargetedManagedAppPolicyAssignment
func (r *TargetedManagedAppPolicyAssignmentRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GovernanceRoleAssignmentRequestObjectRequestBuilder is request builder for GovernanceRoleAssignmentRequestObject
type GovernanceRoleAssignmentRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns GovernanceRoleAssignmentRequestObjectRequest
func (b *GovernanceRoleAssignmentRequestObjectRequestBuilder) Request() *GovernanceRoleAssignmentRequestObjectRequest {
	return &GovernanceRoleAssignmentRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GovernanceRoleAssignmentRequestObjectRequest is request for GovernanceRoleAssignmentRequestObject
type GovernanceRoleAssignmentRequestObjectRequest struct{ BaseRequest }

// Do performs HTTP request for GovernanceRoleAssignmentRequestObject
func (r *GovernanceRoleAssignmentRequestObjectRequest) Do(method, path string, reqObj interface{}) (resObj *GovernanceRoleAssignmentRequestObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for GovernanceRoleAssignmentRequestObject
func (r *GovernanceRoleAssignmentRequestObjectRequest) Get() (*GovernanceRoleAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for GovernanceRoleAssignmentRequestObject
func (r *GovernanceRoleAssignmentRequestObjectRequest) Update(reqObj *GovernanceRoleAssignmentRequestObject) (*GovernanceRoleAssignmentRequestObject, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for GovernanceRoleAssignmentRequestObject
func (r *GovernanceRoleAssignmentRequestObjectRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Resource is navigation property
func (b *GovernanceRoleAssignmentRequestObjectRequestBuilder) Resource() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// RoleDefinition is navigation property
func (b *GovernanceRoleAssignmentRequestObjectRequestBuilder) RoleDefinition() *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}

// Subject is navigation property
func (b *GovernanceRoleAssignmentRequestObjectRequestBuilder) Subject() *GovernanceSubjectRequestBuilder {
	bb := &GovernanceSubjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subject"
	return bb
}

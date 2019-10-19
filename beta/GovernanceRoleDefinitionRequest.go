// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GovernanceRoleDefinitionRequestBuilder is request builder for GovernanceRoleDefinition
type GovernanceRoleDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns GovernanceRoleDefinitionRequest
func (b *GovernanceRoleDefinitionRequestBuilder) Request() *GovernanceRoleDefinitionRequest {
	return &GovernanceRoleDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GovernanceRoleDefinitionRequest is request for GovernanceRoleDefinition
type GovernanceRoleDefinitionRequest struct{ BaseRequest }

// Do performs HTTP request for GovernanceRoleDefinition
func (r *GovernanceRoleDefinitionRequest) Do(method, path string, reqObj interface{}) (resObj *GovernanceRoleDefinition, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for GovernanceRoleDefinition
func (r *GovernanceRoleDefinitionRequest) Get() (*GovernanceRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for GovernanceRoleDefinition
func (r *GovernanceRoleDefinitionRequest) Update(reqObj *GovernanceRoleDefinition) (*GovernanceRoleDefinition, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for GovernanceRoleDefinition
func (r *GovernanceRoleDefinitionRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Resource is navigation property
func (b *GovernanceRoleDefinitionRequestBuilder) Resource() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// RoleSetting is navigation property
func (b *GovernanceRoleDefinitionRequestBuilder) RoleSetting() *GovernanceRoleSettingRequestBuilder {
	bb := &GovernanceRoleSettingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleSetting"
	return bb
}

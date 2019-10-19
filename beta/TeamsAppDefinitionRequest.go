// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// TeamsAppDefinitionRequestBuilder is request builder for TeamsAppDefinition
type TeamsAppDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns TeamsAppDefinitionRequest
func (b *TeamsAppDefinitionRequestBuilder) Request() *TeamsAppDefinitionRequest {
	return &TeamsAppDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TeamsAppDefinitionRequest is request for TeamsAppDefinition
type TeamsAppDefinitionRequest struct{ BaseRequest }

// Do performs HTTP request for TeamsAppDefinition
func (r *TeamsAppDefinitionRequest) Do(method, path string, reqObj interface{}) (resObj *TeamsAppDefinition, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TeamsAppDefinition
func (r *TeamsAppDefinitionRequest) Get() (*TeamsAppDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TeamsAppDefinition
func (r *TeamsAppDefinitionRequest) Update(reqObj *TeamsAppDefinition) (*TeamsAppDefinition, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TeamsAppDefinition
func (r *TeamsAppDefinitionRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

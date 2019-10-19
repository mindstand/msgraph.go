// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedEBookCategoryRequestBuilder is request builder for ManagedEBookCategory
type ManagedEBookCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedEBookCategoryRequest
func (b *ManagedEBookCategoryRequestBuilder) Request() *ManagedEBookCategoryRequest {
	return &ManagedEBookCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedEBookCategoryRequest is request for ManagedEBookCategory
type ManagedEBookCategoryRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedEBookCategory
func (r *ManagedEBookCategoryRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedEBookCategory, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ManagedEBookCategory
func (r *ManagedEBookCategoryRequest) Get() (*ManagedEBookCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ManagedEBookCategory
func (r *ManagedEBookCategoryRequest) Update(reqObj *ManagedEBookCategory) (*ManagedEBookCategory, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ManagedEBookCategory
func (r *ManagedEBookCategoryRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

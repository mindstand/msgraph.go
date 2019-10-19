// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ProfilePhotoRequestBuilder is request builder for ProfilePhoto
type ProfilePhotoRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProfilePhotoRequest
func (b *ProfilePhotoRequestBuilder) Request() *ProfilePhotoRequest {
	return &ProfilePhotoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProfilePhotoRequest is request for ProfilePhoto
type ProfilePhotoRequest struct{ BaseRequest }

// Do performs HTTP request for ProfilePhoto
func (r *ProfilePhotoRequest) Do(method, path string, reqObj interface{}) (resObj *ProfilePhoto, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ProfilePhoto
func (r *ProfilePhotoRequest) Get() (*ProfilePhoto, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ProfilePhoto
func (r *ProfilePhotoRequest) Update(reqObj *ProfilePhoto) (*ProfilePhoto, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ProfilePhoto
func (r *ProfilePhotoRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
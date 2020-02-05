// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DirectorySettingRequestBuilder is request builder for DirectorySetting
type DirectorySettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns DirectorySettingRequest
func (b *DirectorySettingRequestBuilder) Request() *DirectorySettingRequest {
	return &DirectorySettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *DirectorySettingRequestBuilder) Delta() *DirectorySettingRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// DirectorySettingRequest is request for DirectorySetting
type DirectorySettingRequest struct{ BaseRequest }

// Get performs GET request for DirectorySetting
func (r *DirectorySettingRequest) Get(ctx context.Context) (resObj *DirectorySetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DirectorySetting
func (r *DirectorySettingRequest) Update(ctx context.Context, reqObj *DirectorySetting) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DirectorySetting
func (r *DirectorySettingRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// UserSettingsRequestBuilder is request builder for UserSettings
type UserSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserSettingsRequest
func (b *UserSettingsRequestBuilder) Request() *UserSettingsRequest {
	return &UserSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *UserSettingsRequestBuilder) Delta() *UserSettingsRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// UserSettingsRequest is request for UserSettings
type UserSettingsRequest struct{ BaseRequest }

// Get performs GET request for UserSettings
func (r *UserSettingsRequest) Get(ctx context.Context) (resObj *UserSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserSettings
func (r *UserSettingsRequest) Update(ctx context.Context, reqObj *UserSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserSettings
func (r *UserSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

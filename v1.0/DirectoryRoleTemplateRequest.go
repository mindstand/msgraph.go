// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// DirectoryRoleTemplateRequestBuilder is request builder for DirectoryRoleTemplate
type DirectoryRoleTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns DirectoryRoleTemplateRequest
func (b *DirectoryRoleTemplateRequestBuilder) Request() *DirectoryRoleTemplateRequest {
	return &DirectoryRoleTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *DirectoryRoleTemplateRequestBuilder) Delta() *DirectoryRoleTemplateRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// DirectoryRoleTemplateRequest is request for DirectoryRoleTemplate
type DirectoryRoleTemplateRequest struct{ BaseRequest }

// Get performs GET request for DirectoryRoleTemplate
func (r *DirectoryRoleTemplateRequest) Get(ctx context.Context) (resObj *DirectoryRoleTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DirectoryRoleTemplate
func (r *DirectoryRoleTemplateRequest) Update(ctx context.Context, reqObj *DirectoryRoleTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DirectoryRoleTemplate
func (r *DirectoryRoleTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

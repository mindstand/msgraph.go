// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// OnenoteResourceRequestBuilder is request builder for OnenoteResource
type OnenoteResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteResourceRequest
func (b *OnenoteResourceRequestBuilder) Request() *OnenoteResourceRequest {
	return &OnenoteResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *OnenoteResourceRequestBuilder) Delta() *OnenoteResourceRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// OnenoteResourceRequest is request for OnenoteResource
type OnenoteResourceRequest struct{ BaseRequest }

// Get performs GET request for OnenoteResource
func (r *OnenoteResourceRequest) Get(ctx context.Context) (resObj *OnenoteResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnenoteResource
func (r *OnenoteResourceRequest) Update(ctx context.Context, reqObj *OnenoteResource) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnenoteResource
func (r *OnenoteResourceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

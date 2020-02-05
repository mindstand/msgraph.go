// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// ThumbnailSetRequestBuilder is request builder for ThumbnailSet
type ThumbnailSetRequestBuilder struct{ BaseRequestBuilder }

// Request returns ThumbnailSetRequest
func (b *ThumbnailSetRequestBuilder) Request() *ThumbnailSetRequest {
	return &ThumbnailSetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ThumbnailSetRequestBuilder) Delta() *ThumbnailSetRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ThumbnailSetRequest is request for ThumbnailSet
type ThumbnailSetRequest struct{ BaseRequest }

// Get performs GET request for ThumbnailSet
func (r *ThumbnailSetRequest) Get(ctx context.Context) (resObj *ThumbnailSet, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ThumbnailSet
func (r *ThumbnailSetRequest) Update(ctx context.Context, reqObj *ThumbnailSet) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ThumbnailSet
func (r *ThumbnailSetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

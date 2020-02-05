// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PayloadResponseRequestBuilder is request builder for PayloadResponse
type PayloadResponseRequestBuilder struct{ BaseRequestBuilder }

// Request returns PayloadResponseRequest
func (b *PayloadResponseRequestBuilder) Request() *PayloadResponseRequest {
	return &PayloadResponseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *PayloadResponseRequestBuilder) Delta() *PayloadResponseRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// PayloadResponseRequest is request for PayloadResponse
type PayloadResponseRequest struct{ BaseRequest }

// Get performs GET request for PayloadResponse
func (r *PayloadResponseRequest) Get(ctx context.Context) (resObj *PayloadResponse, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PayloadResponse
func (r *PayloadResponseRequest) Update(ctx context.Context, reqObj *PayloadResponse) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PayloadResponse
func (r *PayloadResponseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

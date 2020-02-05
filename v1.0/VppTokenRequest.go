// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// VppTokenRequestBuilder is request builder for VppToken
type VppTokenRequestBuilder struct{ BaseRequestBuilder }

// Request returns VppTokenRequest
func (b *VppTokenRequestBuilder) Request() *VppTokenRequest {
	return &VppTokenRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *VppTokenRequestBuilder) Delta() *VppTokenRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// VppTokenRequest is request for VppToken
type VppTokenRequest struct{ BaseRequest }

// Get performs GET request for VppToken
func (r *VppTokenRequest) Get(ctx context.Context) (resObj *VppToken, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for VppToken
func (r *VppTokenRequest) Update(ctx context.Context, reqObj *VppToken) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for VppToken
func (r *VppTokenRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// PolicySetItemRequestBuilder is request builder for PolicySetItem
type PolicySetItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns PolicySetItemRequest
func (b *PolicySetItemRequestBuilder) Request() *PolicySetItemRequest {
	return &PolicySetItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *PolicySetItemRequestBuilder) Delta() *PolicySetItemRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// PolicySetItemRequest is request for PolicySetItem
type PolicySetItemRequest struct{ BaseRequest }

// Get performs GET request for PolicySetItem
func (r *PolicySetItemRequest) Get(ctx context.Context) (resObj *PolicySetItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PolicySetItem
func (r *PolicySetItemRequest) Update(ctx context.Context, reqObj *PolicySetItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PolicySetItem
func (r *PolicySetItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

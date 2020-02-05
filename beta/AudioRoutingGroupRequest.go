// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// AudioRoutingGroupRequestBuilder is request builder for AudioRoutingGroup
type AudioRoutingGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns AudioRoutingGroupRequest
func (b *AudioRoutingGroupRequestBuilder) Request() *AudioRoutingGroupRequest {
	return &AudioRoutingGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *AudioRoutingGroupRequestBuilder) Delta() *AudioRoutingGroupRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// AudioRoutingGroupRequest is request for AudioRoutingGroup
type AudioRoutingGroupRequest struct{ BaseRequest }

// Get performs GET request for AudioRoutingGroup
func (r *AudioRoutingGroupRequest) Get(ctx context.Context) (resObj *AudioRoutingGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AudioRoutingGroup
func (r *AudioRoutingGroupRequest) Update(ctx context.Context, reqObj *AudioRoutingGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AudioRoutingGroup
func (r *AudioRoutingGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// RoomRequestBuilder is request builder for Room
type RoomRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoomRequest
func (b *RoomRequestBuilder) Request() *RoomRequest {
	return &RoomRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *RoomRequestBuilder) Delta() *RoomRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// RoomRequest is request for Room
type RoomRequest struct{ BaseRequest }

// Get performs GET request for Room
func (r *RoomRequest) Get(ctx context.Context) (resObj *Room, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Room
func (r *RoomRequest) Update(ctx context.Context, reqObj *Room) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Room
func (r *RoomRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// RoomListRequestBuilder is request builder for RoomList
type RoomListRequestBuilder struct{ BaseRequestBuilder }

// Request returns RoomListRequest
func (b *RoomListRequestBuilder) Request() *RoomListRequest {
	return &RoomListRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *RoomListRequestBuilder) Delta() *RoomListRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// RoomListRequest is request for RoomList
type RoomListRequest struct{ BaseRequest }

// Get performs GET request for RoomList
func (r *RoomListRequest) Get(ctx context.Context) (resObj *RoomList, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for RoomList
func (r *RoomListRequest) Update(ctx context.Context, reqObj *RoomList) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for RoomList
func (r *RoomListRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Rooms returns request builder for Room collection
func (b *RoomListRequestBuilder) Rooms() *RoomListRoomsCollectionRequestBuilder {
	bb := &RoomListRoomsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rooms"
	return bb
}

// RoomListRoomsCollectionRequestBuilder is request builder for Room collection
type RoomListRoomsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Room collection
func (b *RoomListRoomsCollectionRequestBuilder) Request() *RoomListRoomsCollectionRequest {
	return &RoomListRoomsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Room item
func (b *RoomListRoomsCollectionRequestBuilder) ID(id string) *RoomRequestBuilder {
	bb := &RoomRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RoomListRoomsCollectionRequest is request for Room collection
type RoomListRoomsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Room collection
func (r *RoomListRoomsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Room, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Room
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []Room
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for Room collection
func (r *RoomListRoomsCollectionRequest) Get(ctx context.Context) ([]Room, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Room collection
func (r *RoomListRoomsCollectionRequest) Add(ctx context.Context, reqObj *Room) (resObj *Room, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

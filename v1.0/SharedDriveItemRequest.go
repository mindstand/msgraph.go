// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SharedDriveItemRequestBuilder is request builder for SharedDriveItem
type SharedDriveItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns SharedDriveItemRequest
func (b *SharedDriveItemRequestBuilder) Request() *SharedDriveItemRequest {
	return &SharedDriveItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *SharedDriveItemRequestBuilder) Delta() *SharedDriveItemRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// SharedDriveItemRequest is request for SharedDriveItem
type SharedDriveItemRequest struct{ BaseRequest }

// Get performs GET request for SharedDriveItem
func (r *SharedDriveItemRequest) Get(ctx context.Context) (resObj *SharedDriveItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SharedDriveItem
func (r *SharedDriveItemRequest) Update(ctx context.Context, reqObj *SharedDriveItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SharedDriveItem
func (r *SharedDriveItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DriveItem is navigation property
func (b *SharedDriveItemRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}

// Items returns request builder for DriveItem collection
func (b *SharedDriveItemRequestBuilder) Items() *SharedDriveItemItemsCollectionRequestBuilder {
	bb := &SharedDriveItemItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// SharedDriveItemItemsCollectionRequestBuilder is request builder for DriveItem collection
type SharedDriveItemItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *SharedDriveItemItemsCollectionRequestBuilder) Request() *SharedDriveItemItemsCollectionRequest {
	return &SharedDriveItemItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *SharedDriveItemItemsCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SharedDriveItemItemsCollectionRequest is request for DriveItem collection
type SharedDriveItemItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *SharedDriveItemItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DriveItem, error) {
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
	var values []DriveItem
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
			value  []DriveItem
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

// Get performs GET request for DriveItem collection
func (r *SharedDriveItemItemsCollectionRequest) Get(ctx context.Context) ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DriveItem collection
func (r *SharedDriveItemItemsCollectionRequest) Add(ctx context.Context, reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// List is navigation property
func (b *SharedDriveItemRequestBuilder) List() *ListRequestBuilder {
	bb := &ListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/list"
	return bb
}

// ListItem is navigation property
func (b *SharedDriveItemRequestBuilder) ListItem() *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/listItem"
	return bb
}

// Root is navigation property
func (b *SharedDriveItemRequestBuilder) Root() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/root"
	return bb
}

// Site is navigation property
func (b *SharedDriveItemRequestBuilder) Site() *SiteRequestBuilder {
	bb := &SiteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/site"
	return bb
}

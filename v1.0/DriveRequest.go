// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DriveRequestBuilder is request builder for Drive
type DriveRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveRequest
func (b *DriveRequestBuilder) Request() *DriveRequest {
	return &DriveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *DriveRequestBuilder) Delta() *DriveRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// DriveRequest is request for Drive
type DriveRequest struct{ BaseRequest }

// Get performs GET request for Drive
func (r *DriveRequest) Get(ctx context.Context) (resObj *Drive, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Drive
func (r *DriveRequest) Update(ctx context.Context, reqObj *Drive) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Drive
func (r *DriveRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Items returns request builder for DriveItem collection
func (b *DriveRequestBuilder) Items() *DriveItemsCollectionRequestBuilder {
	bb := &DriveItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// DriveItemsCollectionRequestBuilder is request builder for DriveItem collection
type DriveItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *DriveItemsCollectionRequestBuilder) Request() *DriveItemsCollectionRequest {
	return &DriveItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *DriveItemsCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemsCollectionRequest is request for DriveItem collection
type DriveItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *DriveItemsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DriveItem, error) {
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
func (r *DriveItemsCollectionRequest) Get(ctx context.Context) ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DriveItem collection
func (r *DriveItemsCollectionRequest) Add(ctx context.Context, reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// List is navigation property
func (b *DriveRequestBuilder) List() *ListRequestBuilder {
	bb := &ListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/list"
	return bb
}

// Root is navigation property
func (b *DriveRequestBuilder) Root() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/root"
	return bb
}

// Special returns request builder for DriveItem collection
func (b *DriveRequestBuilder) Special() *DriveSpecialCollectionRequestBuilder {
	bb := &DriveSpecialCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/special"
	return bb
}

// DriveSpecialCollectionRequestBuilder is request builder for DriveItem collection
type DriveSpecialCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *DriveSpecialCollectionRequestBuilder) Request() *DriveSpecialCollectionRequest {
	return &DriveSpecialCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *DriveSpecialCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveSpecialCollectionRequest is request for DriveItem collection
type DriveSpecialCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *DriveSpecialCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DriveItem, error) {
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
func (r *DriveSpecialCollectionRequest) Get(ctx context.Context) ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DriveItem collection
func (r *DriveSpecialCollectionRequest) Add(ctx context.Context, reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

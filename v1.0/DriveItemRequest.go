// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DriveItemRequestBuilder is request builder for DriveItem
type DriveItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns DriveItemRequest
func (b *DriveItemRequestBuilder) Request() *DriveItemRequest {
	return &DriveItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DriveItemRequest is request for DriveItem
type DriveItemRequest struct{ BaseRequest }

// Get performs GET request for DriveItem
func (r *DriveItemRequest) Get(ctx context.Context) (resObj *DriveItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DriveItem
func (r *DriveItemRequest) Update(ctx context.Context, reqObj *DriveItem) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DriveItem
func (r *DriveItemRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Analytics is navigation property
func (b *DriveItemRequestBuilder) Analytics() *ItemAnalyticsRequestBuilder {
	bb := &ItemAnalyticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/analytics"
	return bb
}

// Children returns request builder for DriveItem collection
func (b *DriveItemRequestBuilder) Children() *DriveItemChildrenCollectionRequestBuilder {
	bb := &DriveItemChildrenCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/children"
	return bb
}

// DriveItemChildrenCollectionRequestBuilder is request builder for DriveItem collection
type DriveItemChildrenCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *DriveItemChildrenCollectionRequestBuilder) Request() *DriveItemChildrenCollectionRequest {
	return &DriveItemChildrenCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *DriveItemChildrenCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemChildrenCollectionRequest is request for DriveItem collection
type DriveItemChildrenCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DriveItem, error) {
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
func (r *DriveItemChildrenCollectionRequest) Get(ctx context.Context) ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DriveItem collection
func (r *DriveItemChildrenCollectionRequest) Add(ctx context.Context, reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ListItem is navigation property
func (b *DriveItemRequestBuilder) ListItem() *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/listItem"
	return bb
}

// Permissions returns request builder for Permission collection
func (b *DriveItemRequestBuilder) Permissions() *DriveItemPermissionsCollectionRequestBuilder {
	bb := &DriveItemPermissionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/permissions"
	return bb
}

// DriveItemPermissionsCollectionRequestBuilder is request builder for Permission collection
type DriveItemPermissionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Permission collection
func (b *DriveItemPermissionsCollectionRequestBuilder) Request() *DriveItemPermissionsCollectionRequest {
	return &DriveItemPermissionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Permission item
func (b *DriveItemPermissionsCollectionRequestBuilder) ID(id string) *PermissionRequestBuilder {
	bb := &PermissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemPermissionsCollectionRequest is request for Permission collection
type DriveItemPermissionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Permission, error) {
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
	var values []Permission
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
			value  []Permission
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

// Get performs GET request for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Get(ctx context.Context) ([]Permission, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Permission collection
func (r *DriveItemPermissionsCollectionRequest) Add(ctx context.Context, reqObj *Permission) (resObj *Permission, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Subscriptions returns request builder for Subscription collection
func (b *DriveItemRequestBuilder) Subscriptions() *DriveItemSubscriptionsCollectionRequestBuilder {
	bb := &DriveItemSubscriptionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/subscriptions"
	return bb
}

// DriveItemSubscriptionsCollectionRequestBuilder is request builder for Subscription collection
type DriveItemSubscriptionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Subscription collection
func (b *DriveItemSubscriptionsCollectionRequestBuilder) Request() *DriveItemSubscriptionsCollectionRequest {
	return &DriveItemSubscriptionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Subscription item
func (b *DriveItemSubscriptionsCollectionRequestBuilder) ID(id string) *SubscriptionRequestBuilder {
	bb := &SubscriptionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemSubscriptionsCollectionRequest is request for Subscription collection
type DriveItemSubscriptionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]Subscription, error) {
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
	var values []Subscription
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
			value  []Subscription
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

// Get performs GET request for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Get(ctx context.Context) ([]Subscription, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for Subscription collection
func (r *DriveItemSubscriptionsCollectionRequest) Add(ctx context.Context, reqObj *Subscription) (resObj *Subscription, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Thumbnails returns request builder for ThumbnailSet collection
func (b *DriveItemRequestBuilder) Thumbnails() *DriveItemThumbnailsCollectionRequestBuilder {
	bb := &DriveItemThumbnailsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/thumbnails"
	return bb
}

// DriveItemThumbnailsCollectionRequestBuilder is request builder for ThumbnailSet collection
type DriveItemThumbnailsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ThumbnailSet collection
func (b *DriveItemThumbnailsCollectionRequestBuilder) Request() *DriveItemThumbnailsCollectionRequest {
	return &DriveItemThumbnailsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ThumbnailSet item
func (b *DriveItemThumbnailsCollectionRequestBuilder) ID(id string) *ThumbnailSetRequestBuilder {
	bb := &ThumbnailSetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemThumbnailsCollectionRequest is request for ThumbnailSet collection
type DriveItemThumbnailsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ThumbnailSet, error) {
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
	var values []ThumbnailSet
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
			value  []ThumbnailSet
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

// Get performs GET request for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Get(ctx context.Context) ([]ThumbnailSet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ThumbnailSet collection
func (r *DriveItemThumbnailsCollectionRequest) Add(ctx context.Context, reqObj *ThumbnailSet) (resObj *ThumbnailSet, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Versions returns request builder for DriveItemVersion collection
func (b *DriveItemRequestBuilder) Versions() *DriveItemVersionsCollectionRequestBuilder {
	bb := &DriveItemVersionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/versions"
	return bb
}

// DriveItemVersionsCollectionRequestBuilder is request builder for DriveItemVersion collection
type DriveItemVersionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItemVersion collection
func (b *DriveItemVersionsCollectionRequestBuilder) Request() *DriveItemVersionsCollectionRequest {
	return &DriveItemVersionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItemVersion item
func (b *DriveItemVersionsCollectionRequestBuilder) ID(id string) *DriveItemVersionRequestBuilder {
	bb := &DriveItemVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DriveItemVersionsCollectionRequest is request for DriveItemVersion collection
type DriveItemVersionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DriveItemVersion, error) {
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
	var values []DriveItemVersion
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
			value  []DriveItemVersion
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

// Get performs GET request for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Get(ctx context.Context) ([]DriveItemVersion, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DriveItemVersion collection
func (r *DriveItemVersionsCollectionRequest) Add(ctx context.Context, reqObj *DriveItemVersion) (resObj *DriveItemVersion, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Workbook is navigation property
func (b *DriveItemRequestBuilder) Workbook() *WorkbookRequestBuilder {
	bb := &WorkbookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/workbook"
	return bb
}

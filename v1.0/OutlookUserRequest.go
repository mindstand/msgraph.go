// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OutlookUserRequestBuilder is request builder for OutlookUser
type OutlookUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookUserRequest
func (b *OutlookUserRequestBuilder) Request() *OutlookUserRequest {
	return &OutlookUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *OutlookUserRequestBuilder) Delta() *OutlookUserRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// OutlookUserRequest is request for OutlookUser
type OutlookUserRequest struct{ BaseRequest }

// Get performs GET request for OutlookUser
func (r *OutlookUserRequest) Get(ctx context.Context) (resObj *OutlookUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OutlookUser
func (r *OutlookUserRequest) Update(ctx context.Context, reqObj *OutlookUser) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OutlookUser
func (r *OutlookUserRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// MasterCategories returns request builder for OutlookCategory collection
func (b *OutlookUserRequestBuilder) MasterCategories() *OutlookUserMasterCategoriesCollectionRequestBuilder {
	bb := &OutlookUserMasterCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/masterCategories"
	return bb
}

// OutlookUserMasterCategoriesCollectionRequestBuilder is request builder for OutlookCategory collection
type OutlookUserMasterCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OutlookCategory collection
func (b *OutlookUserMasterCategoriesCollectionRequestBuilder) Request() *OutlookUserMasterCategoriesCollectionRequest {
	return &OutlookUserMasterCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OutlookCategory item
func (b *OutlookUserMasterCategoriesCollectionRequestBuilder) ID(id string) *OutlookCategoryRequestBuilder {
	bb := &OutlookCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OutlookUserMasterCategoriesCollectionRequest is request for OutlookCategory collection
type OutlookUserMasterCategoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]OutlookCategory, error) {
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
	var values []OutlookCategory
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
			value  []OutlookCategory
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

// Get performs GET request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Get(ctx context.Context) ([]OutlookCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Add(ctx context.Context, reqObj *OutlookCategory) (resObj *OutlookCategory, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OutlookUserRequestBuilder is request builder for OutlookUser
type OutlookUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookUserRequest
func (b *OutlookUserRequestBuilder) Request() *OutlookUserRequest {
	return &OutlookUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookUserRequest is request for OutlookUser
type OutlookUserRequest struct{ BaseRequest }

// Do performs HTTP request for OutlookUser
func (r *OutlookUserRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookUser, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OutlookUser
func (r *OutlookUserRequest) Get() (*OutlookUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OutlookUser
func (r *OutlookUserRequest) Update(reqObj *OutlookUser) (*OutlookUser, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OutlookUser
func (r *OutlookUserRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
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

// Do performs HTTP request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookCategory, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Paging(method, path string, obj interface{}) ([]OutlookCategory, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
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
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OutlookCategory
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Get() ([]OutlookCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Add(reqObj *OutlookCategory) (*OutlookCategory, error) {
	return r.Do("POST", "", reqObj)
}

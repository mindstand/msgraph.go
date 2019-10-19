// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// ItemActivityStatRequestBuilder is request builder for ItemActivityStat
type ItemActivityStatRequestBuilder struct{ BaseRequestBuilder }

// Request returns ItemActivityStatRequest
func (b *ItemActivityStatRequestBuilder) Request() *ItemActivityStatRequest {
	return &ItemActivityStatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ItemActivityStatRequest is request for ItemActivityStat
type ItemActivityStatRequest struct{ BaseRequest }

// Do performs HTTP request for ItemActivityStat
func (r *ItemActivityStatRequest) Do(method, path string, reqObj interface{}) (resObj *ItemActivityStat, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for ItemActivityStat
func (r *ItemActivityStatRequest) Get() (*ItemActivityStat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for ItemActivityStat
func (r *ItemActivityStatRequest) Update(reqObj *ItemActivityStat) (*ItemActivityStat, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for ItemActivityStat
func (r *ItemActivityStatRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Activities returns request builder for ItemActivity collection
func (b *ItemActivityStatRequestBuilder) Activities() *ItemActivityStatActivitiesCollectionRequestBuilder {
	bb := &ItemActivityStatActivitiesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activities"
	return bb
}

// ItemActivityStatActivitiesCollectionRequestBuilder is request builder for ItemActivity collection
type ItemActivityStatActivitiesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ItemActivity collection
func (b *ItemActivityStatActivitiesCollectionRequestBuilder) Request() *ItemActivityStatActivitiesCollectionRequest {
	return &ItemActivityStatActivitiesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ItemActivity item
func (b *ItemActivityStatActivitiesCollectionRequestBuilder) ID(id string) *ItemActivityRequestBuilder {
	bb := &ItemActivityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ItemActivityStatActivitiesCollectionRequest is request for ItemActivity collection
type ItemActivityStatActivitiesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ItemActivity collection
func (r *ItemActivityStatActivitiesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ItemActivity, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ItemActivity collection
func (r *ItemActivityStatActivitiesCollectionRequest) Paging(method, path string, obj interface{}) ([]ItemActivity, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ItemActivity
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ItemActivity
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

// Get performs GET request for ItemActivity collection
func (r *ItemActivityStatActivitiesCollectionRequest) Get() ([]ItemActivity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ItemActivity collection
func (r *ItemActivityStatActivitiesCollectionRequest) Add(reqObj *ItemActivity) (*ItemActivity, error) {
	return r.Do("POST", "", reqObj)
}

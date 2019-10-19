// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PlannerGroupRequestBuilder is request builder for PlannerGroup
type PlannerGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerGroupRequest
func (b *PlannerGroupRequestBuilder) Request() *PlannerGroupRequest {
	return &PlannerGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerGroupRequest is request for PlannerGroup
type PlannerGroupRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerGroup
func (r *PlannerGroupRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerGroup, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PlannerGroup
func (r *PlannerGroupRequest) Get() (*PlannerGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PlannerGroup
func (r *PlannerGroupRequest) Update(reqObj *PlannerGroup) (*PlannerGroup, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PlannerGroup
func (r *PlannerGroupRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Plans returns request builder for PlannerPlan collection
func (b *PlannerGroupRequestBuilder) Plans() *PlannerGroupPlansCollectionRequestBuilder {
	bb := &PlannerGroupPlansCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/plans"
	return bb
}

// PlannerGroupPlansCollectionRequestBuilder is request builder for PlannerPlan collection
type PlannerGroupPlansCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PlannerPlan collection
func (b *PlannerGroupPlansCollectionRequestBuilder) Request() *PlannerGroupPlansCollectionRequest {
	return &PlannerGroupPlansCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PlannerPlan item
func (b *PlannerGroupPlansCollectionRequestBuilder) ID(id string) *PlannerPlanRequestBuilder {
	bb := &PlannerPlanRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PlannerGroupPlansCollectionRequest is request for PlannerPlan collection
type PlannerGroupPlansCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerPlan, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Paging(method, path string, obj interface{}) ([]PlannerPlan, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PlannerPlan
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []PlannerPlan
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

// Get performs GET request for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Get() ([]PlannerPlan, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for PlannerPlan collection
func (r *PlannerGroupPlansCollectionRequest) Add(reqObj *PlannerPlan) (*PlannerPlan, error) {
	return r.Do("POST", "", reqObj)
}

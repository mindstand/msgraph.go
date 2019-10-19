// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// RequestObjectRequestBuilder is request builder for RequestObject
type RequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns RequestObjectRequest
func (b *RequestObjectRequestBuilder) Request() *RequestObjectRequest {
	return &RequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// RequestObjectRequest is request for RequestObject
type RequestObjectRequest struct{ BaseRequest }

// Do performs HTTP request for RequestObject
func (r *RequestObjectRequest) Do(method, path string, reqObj interface{}) (resObj *RequestObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for RequestObject
func (r *RequestObjectRequest) Get() (*RequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for RequestObject
func (r *RequestObjectRequest) Update(reqObj *RequestObject) (*RequestObject, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for RequestObject
func (r *RequestObjectRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Decisions returns request builder for AccessReviewDecision collection
func (b *RequestObjectRequestBuilder) Decisions() *RequestObjectDecisionsCollectionRequestBuilder {
	bb := &RequestObjectDecisionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/decisions"
	return bb
}

// RequestObjectDecisionsCollectionRequestBuilder is request builder for AccessReviewDecision collection
type RequestObjectDecisionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessReviewDecision collection
func (b *RequestObjectDecisionsCollectionRequestBuilder) Request() *RequestObjectDecisionsCollectionRequest {
	return &RequestObjectDecisionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessReviewDecision item
func (b *RequestObjectDecisionsCollectionRequestBuilder) ID(id string) *AccessReviewDecisionRequestBuilder {
	bb := &AccessReviewDecisionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RequestObjectDecisionsCollectionRequest is request for AccessReviewDecision collection
type RequestObjectDecisionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for AccessReviewDecision collection
func (r *RequestObjectDecisionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *AccessReviewDecision, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for AccessReviewDecision collection
func (r *RequestObjectDecisionsCollectionRequest) Paging(method, path string, obj interface{}) ([]AccessReviewDecision, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AccessReviewDecision
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AccessReviewDecision
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

// Get performs GET request for AccessReviewDecision collection
func (r *RequestObjectDecisionsCollectionRequest) Get() ([]AccessReviewDecision, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AccessReviewDecision collection
func (r *RequestObjectDecisionsCollectionRequest) Add(reqObj *AccessReviewDecision) (*AccessReviewDecision, error) {
	return r.Do("POST", "", reqObj)
}

// MyDecisions returns request builder for AccessReviewDecision collection
func (b *RequestObjectRequestBuilder) MyDecisions() *RequestObjectMyDecisionsCollectionRequestBuilder {
	bb := &RequestObjectMyDecisionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/myDecisions"
	return bb
}

// RequestObjectMyDecisionsCollectionRequestBuilder is request builder for AccessReviewDecision collection
type RequestObjectMyDecisionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessReviewDecision collection
func (b *RequestObjectMyDecisionsCollectionRequestBuilder) Request() *RequestObjectMyDecisionsCollectionRequest {
	return &RequestObjectMyDecisionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessReviewDecision item
func (b *RequestObjectMyDecisionsCollectionRequestBuilder) ID(id string) *AccessReviewDecisionRequestBuilder {
	bb := &AccessReviewDecisionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// RequestObjectMyDecisionsCollectionRequest is request for AccessReviewDecision collection
type RequestObjectMyDecisionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for AccessReviewDecision collection
func (r *RequestObjectMyDecisionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *AccessReviewDecision, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for AccessReviewDecision collection
func (r *RequestObjectMyDecisionsCollectionRequest) Paging(method, path string, obj interface{}) ([]AccessReviewDecision, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AccessReviewDecision
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AccessReviewDecision
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

// Get performs GET request for AccessReviewDecision collection
func (r *RequestObjectMyDecisionsCollectionRequest) Get() ([]AccessReviewDecision, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AccessReviewDecision collection
func (r *RequestObjectMyDecisionsCollectionRequest) Add(reqObj *AccessReviewDecision) (*AccessReviewDecision, error) {
	return r.Do("POST", "", reqObj)
}

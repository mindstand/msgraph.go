// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ExternalRequestBuilder is request builder for External
type ExternalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ExternalRequest
func (b *ExternalRequestBuilder) Request() *ExternalRequest {
	return &ExternalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ExternalRequestBuilder) Delta() *ExternalRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ExternalRequest is request for External
type ExternalRequest struct{ BaseRequest }

// Get performs GET request for External
func (r *ExternalRequest) Get(ctx context.Context) (resObj *External, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for External
func (r *ExternalRequest) Update(ctx context.Context, reqObj *External) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for External
func (r *ExternalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Connections returns request builder for ExternalConnection collection
func (b *ExternalRequestBuilder) Connections() *ExternalConnectionsCollectionRequestBuilder {
	bb := &ExternalConnectionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/connections"
	return bb
}

// ExternalConnectionsCollectionRequestBuilder is request builder for ExternalConnection collection
type ExternalConnectionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ExternalConnection collection
func (b *ExternalConnectionsCollectionRequestBuilder) Request() *ExternalConnectionsCollectionRequest {
	return &ExternalConnectionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ExternalConnection item
func (b *ExternalConnectionsCollectionRequestBuilder) ID(id string) *ExternalConnectionRequestBuilder {
	bb := &ExternalConnectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ExternalConnectionsCollectionRequest is request for ExternalConnection collection
type ExternalConnectionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ExternalConnection collection
func (r *ExternalConnectionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ExternalConnection, error) {
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
	var values []ExternalConnection
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
			value  []ExternalConnection
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

// Get performs GET request for ExternalConnection collection
func (r *ExternalConnectionsCollectionRequest) Get(ctx context.Context) ([]ExternalConnection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ExternalConnection collection
func (r *ExternalConnectionsCollectionRequest) Add(ctx context.Context, reqObj *ExternalConnection) (resObj *ExternalConnection, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

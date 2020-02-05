// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// UserAnalyticsRequestBuilder is request builder for UserAnalytics
type UserAnalyticsRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserAnalyticsRequest
func (b *UserAnalyticsRequestBuilder) Request() *UserAnalyticsRequest {
	return &UserAnalyticsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *UserAnalyticsRequestBuilder) Delta() *UserAnalyticsRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// UserAnalyticsRequest is request for UserAnalytics
type UserAnalyticsRequest struct{ BaseRequest }

// Get performs GET request for UserAnalytics
func (r *UserAnalyticsRequest) Get(ctx context.Context) (resObj *UserAnalytics, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserAnalytics
func (r *UserAnalyticsRequest) Update(ctx context.Context, reqObj *UserAnalytics) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserAnalytics
func (r *UserAnalyticsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ActivityStatistics returns request builder for ActivityStatistics collection
func (b *UserAnalyticsRequestBuilder) ActivityStatistics() *UserAnalyticsActivityStatisticsCollectionRequestBuilder {
	bb := &UserAnalyticsActivityStatisticsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activityStatistics"
	return bb
}

// UserAnalyticsActivityStatisticsCollectionRequestBuilder is request builder for ActivityStatistics collection
type UserAnalyticsActivityStatisticsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ActivityStatistics collection
func (b *UserAnalyticsActivityStatisticsCollectionRequestBuilder) Request() *UserAnalyticsActivityStatisticsCollectionRequest {
	return &UserAnalyticsActivityStatisticsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ActivityStatistics item
func (b *UserAnalyticsActivityStatisticsCollectionRequestBuilder) ID(id string) *ActivityStatisticsRequestBuilder {
	bb := &ActivityStatisticsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserAnalyticsActivityStatisticsCollectionRequest is request for ActivityStatistics collection
type UserAnalyticsActivityStatisticsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ActivityStatistics, error) {
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
	var values []ActivityStatistics
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
			value  []ActivityStatistics
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

// Get performs GET request for ActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionRequest) Get(ctx context.Context) ([]ActivityStatistics, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ActivityStatistics collection
func (r *UserAnalyticsActivityStatisticsCollectionRequest) Add(ctx context.Context, reqObj *ActivityStatistics) (resObj *ActivityStatistics, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

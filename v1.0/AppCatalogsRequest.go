// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AppCatalogsRequestBuilder is request builder for AppCatalogs
type AppCatalogsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppCatalogsRequest
func (b *AppCatalogsRequestBuilder) Request() *AppCatalogsRequest {
	return &AppCatalogsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *AppCatalogsRequestBuilder) Delta() *AppCatalogsRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// AppCatalogsRequest is request for AppCatalogs
type AppCatalogsRequest struct{ BaseRequest }

// Get performs GET request for AppCatalogs
func (r *AppCatalogsRequest) Get(ctx context.Context) (resObj *AppCatalogs, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppCatalogs
func (r *AppCatalogsRequest) Update(ctx context.Context, reqObj *AppCatalogs) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppCatalogs
func (r *AppCatalogsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// TeamsApps returns request builder for TeamsApp collection
func (b *AppCatalogsRequestBuilder) TeamsApps() *AppCatalogsTeamsAppsCollectionRequestBuilder {
	bb := &AppCatalogsTeamsAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teamsApps"
	return bb
}

// AppCatalogsTeamsAppsCollectionRequestBuilder is request builder for TeamsApp collection
type AppCatalogsTeamsAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsApp collection
func (b *AppCatalogsTeamsAppsCollectionRequestBuilder) Request() *AppCatalogsTeamsAppsCollectionRequest {
	return &AppCatalogsTeamsAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsApp item
func (b *AppCatalogsTeamsAppsCollectionRequestBuilder) ID(id string) *TeamsAppRequestBuilder {
	bb := &TeamsAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppCatalogsTeamsAppsCollectionRequest is request for TeamsApp collection
type AppCatalogsTeamsAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TeamsApp, error) {
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
	var values []TeamsApp
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
			value  []TeamsApp
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

// Get performs GET request for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Get(ctx context.Context) ([]TeamsApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TeamsApp collection
func (r *AppCatalogsTeamsAppsCollectionRequest) Add(ctx context.Context, reqObj *TeamsApp) (resObj *TeamsApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

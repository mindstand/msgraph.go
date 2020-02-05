// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// IOSVppAppRequestBuilder is request builder for IOSVppApp
type IOSVppAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSVppAppRequest
func (b *IOSVppAppRequestBuilder) Request() *IOSVppAppRequest {
	return &IOSVppAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *IOSVppAppRequestBuilder) Delta() *IOSVppAppRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// IOSVppAppRequest is request for IOSVppApp
type IOSVppAppRequest struct{ BaseRequest }

// Get performs GET request for IOSVppApp
func (r *IOSVppAppRequest) Get(ctx context.Context) (resObj *IOSVppApp, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSVppApp
func (r *IOSVppAppRequest) Update(ctx context.Context, reqObj *IOSVppApp) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSVppApp
func (r *IOSVppAppRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AssignedLicenses returns request builder for IOSVppAppAssignedLicense collection
func (b *IOSVppAppRequestBuilder) AssignedLicenses() *IOSVppAppAssignedLicensesCollectionRequestBuilder {
	bb := &IOSVppAppAssignedLicensesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignedLicenses"
	return bb
}

// IOSVppAppAssignedLicensesCollectionRequestBuilder is request builder for IOSVppAppAssignedLicense collection
type IOSVppAppAssignedLicensesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for IOSVppAppAssignedLicense collection
func (b *IOSVppAppAssignedLicensesCollectionRequestBuilder) Request() *IOSVppAppAssignedLicensesCollectionRequest {
	return &IOSVppAppAssignedLicensesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for IOSVppAppAssignedLicense item
func (b *IOSVppAppAssignedLicensesCollectionRequestBuilder) ID(id string) *IOSVppAppAssignedLicenseRequestBuilder {
	bb := &IOSVppAppAssignedLicenseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IOSVppAppAssignedLicensesCollectionRequest is request for IOSVppAppAssignedLicense collection
type IOSVppAppAssignedLicensesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for IOSVppAppAssignedLicense collection
func (r *IOSVppAppAssignedLicensesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]IOSVppAppAssignedLicense, error) {
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
	var values []IOSVppAppAssignedLicense
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
			value  []IOSVppAppAssignedLicense
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

// Get performs GET request for IOSVppAppAssignedLicense collection
func (r *IOSVppAppAssignedLicensesCollectionRequest) Get(ctx context.Context) ([]IOSVppAppAssignedLicense, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for IOSVppAppAssignedLicense collection
func (r *IOSVppAppAssignedLicensesCollectionRequest) Add(ctx context.Context, reqObj *IOSVppAppAssignedLicense) (resObj *IOSVppAppAssignedLicense, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IOSVppAppRequestBuilder is request builder for IOSVppApp
type IOSVppAppRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSVppAppRequest
func (b *IOSVppAppRequestBuilder) Request() *IOSVppAppRequest {
	return &IOSVppAppRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSVppAppRequest is request for IOSVppApp
type IOSVppAppRequest struct{ BaseRequest }

// Do performs HTTP request for IOSVppApp
func (r *IOSVppAppRequest) Do(method, path string, reqObj interface{}) (resObj *IOSVppApp, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for IOSVppApp
func (r *IOSVppAppRequest) Get() (*IOSVppApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for IOSVppApp
func (r *IOSVppAppRequest) Update(reqObj *IOSVppApp) (*IOSVppApp, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for IOSVppApp
func (r *IOSVppAppRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
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

// Do performs HTTP request for IOSVppAppAssignedLicense collection
func (r *IOSVppAppAssignedLicensesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *IOSVppAppAssignedLicense, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for IOSVppAppAssignedLicense collection
func (r *IOSVppAppAssignedLicensesCollectionRequest) Paging(method, path string, obj interface{}) ([]IOSVppAppAssignedLicense, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
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
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []IOSVppAppAssignedLicense
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

// Get performs GET request for IOSVppAppAssignedLicense collection
func (r *IOSVppAppAssignedLicensesCollectionRequest) Get() ([]IOSVppAppAssignedLicense, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for IOSVppAppAssignedLicense collection
func (r *IOSVppAppAssignedLicensesCollectionRequest) Add(reqObj *IOSVppAppAssignedLicense) (*IOSVppAppAssignedLicense, error) {
	return r.Do("POST", "", reqObj)
}

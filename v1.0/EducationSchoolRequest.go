// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// EducationSchoolRequestBuilder is request builder for EducationSchool
type EducationSchoolRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationSchoolRequest
func (b *EducationSchoolRequestBuilder) Request() *EducationSchoolRequest {
	return &EducationSchoolRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationSchoolRequest is request for EducationSchool
type EducationSchoolRequest struct{ BaseRequest }

// Do performs HTTP request for EducationSchool
func (r *EducationSchoolRequest) Do(method, path string, reqObj interface{}) (resObj *EducationSchool, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for EducationSchool
func (r *EducationSchoolRequest) Get() (*EducationSchool, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for EducationSchool
func (r *EducationSchoolRequest) Update(reqObj *EducationSchool) (*EducationSchool, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for EducationSchool
func (r *EducationSchoolRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Classes returns request builder for EducationClass collection
func (b *EducationSchoolRequestBuilder) Classes() *EducationSchoolClassesCollectionRequestBuilder {
	bb := &EducationSchoolClassesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classes"
	return bb
}

// EducationSchoolClassesCollectionRequestBuilder is request builder for EducationClass collection
type EducationSchoolClassesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationClass collection
func (b *EducationSchoolClassesCollectionRequestBuilder) Request() *EducationSchoolClassesCollectionRequest {
	return &EducationSchoolClassesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationClass item
func (b *EducationSchoolClassesCollectionRequestBuilder) ID(id string) *EducationClassRequestBuilder {
	bb := &EducationClassRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationSchoolClassesCollectionRequest is request for EducationClass collection
type EducationSchoolClassesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *EducationClass, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationClass, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationClass
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationClass
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

// Get performs GET request for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) Get() ([]EducationClass, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationClass collection
func (r *EducationSchoolClassesCollectionRequest) Add(reqObj *EducationClass) (*EducationClass, error) {
	return r.Do("POST", "", reqObj)
}

// Users returns request builder for EducationUser collection
func (b *EducationSchoolRequestBuilder) Users() *EducationSchoolUsersCollectionRequestBuilder {
	bb := &EducationSchoolUsersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/users"
	return bb
}

// EducationSchoolUsersCollectionRequestBuilder is request builder for EducationUser collection
type EducationSchoolUsersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationUser collection
func (b *EducationSchoolUsersCollectionRequestBuilder) Request() *EducationSchoolUsersCollectionRequest {
	return &EducationSchoolUsersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationUser item
func (b *EducationSchoolUsersCollectionRequestBuilder) ID(id string) *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationSchoolUsersCollectionRequest is request for EducationUser collection
type EducationSchoolUsersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *EducationUser, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationUser, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationUser
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationUser
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

// Get performs GET request for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) Get() ([]EducationUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationUser collection
func (r *EducationSchoolUsersCollectionRequest) Add(reqObj *EducationUser) (*EducationUser, error) {
	return r.Do("POST", "", reqObj)
}

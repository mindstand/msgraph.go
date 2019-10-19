// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AdministrativeUnitRequestBuilder is request builder for AdministrativeUnit
type AdministrativeUnitRequestBuilder struct{ BaseRequestBuilder }

// Request returns AdministrativeUnitRequest
func (b *AdministrativeUnitRequestBuilder) Request() *AdministrativeUnitRequest {
	return &AdministrativeUnitRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AdministrativeUnitRequest is request for AdministrativeUnit
type AdministrativeUnitRequest struct{ BaseRequest }

// Do performs HTTP request for AdministrativeUnit
func (r *AdministrativeUnitRequest) Do(method, path string, reqObj interface{}) (resObj *AdministrativeUnit, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AdministrativeUnit
func (r *AdministrativeUnitRequest) Get() (*AdministrativeUnit, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AdministrativeUnit
func (r *AdministrativeUnitRequest) Update(reqObj *AdministrativeUnit) (*AdministrativeUnit, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AdministrativeUnit
func (r *AdministrativeUnitRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// Extensions returns request builder for Extension collection
func (b *AdministrativeUnitRequestBuilder) Extensions() *AdministrativeUnitExtensionsCollectionRequestBuilder {
	bb := &AdministrativeUnitExtensionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/extensions"
	return bb
}

// AdministrativeUnitExtensionsCollectionRequestBuilder is request builder for Extension collection
type AdministrativeUnitExtensionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Extension collection
func (b *AdministrativeUnitExtensionsCollectionRequestBuilder) Request() *AdministrativeUnitExtensionsCollectionRequest {
	return &AdministrativeUnitExtensionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Extension item
func (b *AdministrativeUnitExtensionsCollectionRequestBuilder) ID(id string) *ExtensionRequestBuilder {
	bb := &ExtensionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitExtensionsCollectionRequest is request for Extension collection
type AdministrativeUnitExtensionsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for Extension collection
func (r *AdministrativeUnitExtensionsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *Extension, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for Extension collection
func (r *AdministrativeUnitExtensionsCollectionRequest) Paging(method, path string, obj interface{}) ([]Extension, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Extension
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Extension
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

// Get performs GET request for Extension collection
func (r *AdministrativeUnitExtensionsCollectionRequest) Get() ([]Extension, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Extension collection
func (r *AdministrativeUnitExtensionsCollectionRequest) Add(reqObj *Extension) (*Extension, error) {
	return r.Do("POST", "", reqObj)
}

// Members returns request builder for DirectoryObject collection
func (b *AdministrativeUnitRequestBuilder) Members() *AdministrativeUnitMembersCollectionRequestBuilder {
	bb := &AdministrativeUnitMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// AdministrativeUnitMembersCollectionRequestBuilder is request builder for DirectoryObject collection
type AdministrativeUnitMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *AdministrativeUnitMembersCollectionRequestBuilder) Request() *AdministrativeUnitMembersCollectionRequest {
	return &AdministrativeUnitMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *AdministrativeUnitMembersCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitMembersCollectionRequest is request for DirectoryObject collection
type AdministrativeUnitMembersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryObject collection
func (r *AdministrativeUnitMembersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryObject, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryObject collection
func (r *AdministrativeUnitMembersCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *AdministrativeUnitMembersCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *AdministrativeUnitMembersCollectionRequest) Add(reqObj *DirectoryObject) (*DirectoryObject, error) {
	return r.Do("POST", "", reqObj)
}

// ScopedRoleMembers returns request builder for ScopedRoleMembership collection
func (b *AdministrativeUnitRequestBuilder) ScopedRoleMembers() *AdministrativeUnitScopedRoleMembersCollectionRequestBuilder {
	bb := &AdministrativeUnitScopedRoleMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/scopedRoleMembers"
	return bb
}

// AdministrativeUnitScopedRoleMembersCollectionRequestBuilder is request builder for ScopedRoleMembership collection
type AdministrativeUnitScopedRoleMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ScopedRoleMembership collection
func (b *AdministrativeUnitScopedRoleMembersCollectionRequestBuilder) Request() *AdministrativeUnitScopedRoleMembersCollectionRequest {
	return &AdministrativeUnitScopedRoleMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ScopedRoleMembership item
func (b *AdministrativeUnitScopedRoleMembersCollectionRequestBuilder) ID(id string) *ScopedRoleMembershipRequestBuilder {
	bb := &ScopedRoleMembershipRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AdministrativeUnitScopedRoleMembersCollectionRequest is request for ScopedRoleMembership collection
type AdministrativeUnitScopedRoleMembersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ScopedRoleMembership collection
func (r *AdministrativeUnitScopedRoleMembersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ScopedRoleMembership, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ScopedRoleMembership collection
func (r *AdministrativeUnitScopedRoleMembersCollectionRequest) Paging(method, path string, obj interface{}) ([]ScopedRoleMembership, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ScopedRoleMembership
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ScopedRoleMembership
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

// Get performs GET request for ScopedRoleMembership collection
func (r *AdministrativeUnitScopedRoleMembersCollectionRequest) Get() ([]ScopedRoleMembership, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ScopedRoleMembership collection
func (r *AdministrativeUnitScopedRoleMembersCollectionRequest) Add(reqObj *ScopedRoleMembership) (*ScopedRoleMembership, error) {
	return r.Do("POST", "", reqObj)
}

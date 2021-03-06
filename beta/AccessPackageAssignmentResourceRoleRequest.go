// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AccessPackageAssignmentResourceRoleRequestBuilder is request builder for AccessPackageAssignmentResourceRole
type AccessPackageAssignmentResourceRoleRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessPackageAssignmentResourceRoleRequest
func (b *AccessPackageAssignmentResourceRoleRequestBuilder) Request() *AccessPackageAssignmentResourceRoleRequest {
	return &AccessPackageAssignmentResourceRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessPackageAssignmentResourceRoleRequest is request for AccessPackageAssignmentResourceRole
type AccessPackageAssignmentResourceRoleRequest struct{ BaseRequest }

// Get performs GET request for AccessPackageAssignmentResourceRole
func (r *AccessPackageAssignmentResourceRoleRequest) Get(ctx context.Context) (resObj *AccessPackageAssignmentResourceRole, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessPackageAssignmentResourceRole
func (r *AccessPackageAssignmentResourceRoleRequest) Update(ctx context.Context, reqObj *AccessPackageAssignmentResourceRole) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessPackageAssignmentResourceRole
func (r *AccessPackageAssignmentResourceRoleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// AccessPackageAssignments returns request builder for AccessPackageAssignment collection
func (b *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageAssignments() *AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequestBuilder {
	bb := &AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageAssignments"
	return bb
}

// AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequestBuilder is request builder for AccessPackageAssignment collection
type AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AccessPackageAssignment collection
func (b *AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequestBuilder) Request() *AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequest {
	return &AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AccessPackageAssignment item
func (b *AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequestBuilder) ID(id string) *AccessPackageAssignmentRequestBuilder {
	bb := &AccessPackageAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequest is request for AccessPackageAssignment collection
type AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AccessPackageAssignment collection
func (r *AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]AccessPackageAssignment, error) {
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
	var values []AccessPackageAssignment
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
			value  []AccessPackageAssignment
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

// Get performs GET request for AccessPackageAssignment collection
func (r *AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequest) Get(ctx context.Context) ([]AccessPackageAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for AccessPackageAssignment collection
func (r *AccessPackageAssignmentResourceRoleAccessPackageAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *AccessPackageAssignment) (resObj *AccessPackageAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// AccessPackageResourceRole is navigation property
func (b *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageResourceRole() *AccessPackageResourceRoleRequestBuilder {
	bb := &AccessPackageResourceRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceRole"
	return bb
}

// AccessPackageResourceScope is navigation property
func (b *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageResourceScope() *AccessPackageResourceScopeRequestBuilder {
	bb := &AccessPackageResourceScopeRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageResourceScope"
	return bb
}

// AccessPackageSubject is navigation property
func (b *AccessPackageAssignmentResourceRoleRequestBuilder) AccessPackageSubject() *AccessPackageSubjectRequestBuilder {
	bb := &AccessPackageSubjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/accessPackageSubject"
	return bb
}

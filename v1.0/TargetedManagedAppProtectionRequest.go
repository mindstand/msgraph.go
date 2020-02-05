// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// TargetedManagedAppProtectionRequestBuilder is request builder for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppProtectionRequest
func (b *TargetedManagedAppProtectionRequestBuilder) Request() *TargetedManagedAppProtectionRequest {
	return &TargetedManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *TargetedManagedAppProtectionRequestBuilder) Delta() *TargetedManagedAppProtectionRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// TargetedManagedAppProtectionRequest is request for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequest struct{ BaseRequest }

// Get performs GET request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Get(ctx context.Context) (resObj *TargetedManagedAppProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Update(ctx context.Context, reqObj *TargetedManagedAppProtection) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Assignments returns request builder for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppProtectionRequestBuilder) Assignments() *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder {
	bb := &TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder is request builder for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder) Request() *TargetedManagedAppProtectionAssignmentsCollectionRequest {
	return &TargetedManagedAppProtectionAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TargetedManagedAppPolicyAssignment item
func (b *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder) ID(id string) *TargetedManagedAppPolicyAssignmentRequestBuilder {
	bb := &TargetedManagedAppPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TargetedManagedAppProtectionAssignmentsCollectionRequest is request for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppProtectionAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]TargetedManagedAppPolicyAssignment, error) {
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
	var values []TargetedManagedAppPolicyAssignment
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
			value  []TargetedManagedAppPolicyAssignment
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

// Get performs GET request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Get(ctx context.Context) ([]TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *TargetedManagedAppPolicyAssignment) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

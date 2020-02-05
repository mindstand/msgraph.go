// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AndroidForWorkScepCertificateProfileRequestBuilder is request builder for AndroidForWorkScepCertificateProfile
type AndroidForWorkScepCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidForWorkScepCertificateProfileRequest
func (b *AndroidForWorkScepCertificateProfileRequestBuilder) Request() *AndroidForWorkScepCertificateProfileRequest {
	return &AndroidForWorkScepCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *AndroidForWorkScepCertificateProfileRequestBuilder) Delta() *AndroidForWorkScepCertificateProfileRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// AndroidForWorkScepCertificateProfileRequest is request for AndroidForWorkScepCertificateProfile
type AndroidForWorkScepCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for AndroidForWorkScepCertificateProfile
func (r *AndroidForWorkScepCertificateProfileRequest) Get(ctx context.Context) (resObj *AndroidForWorkScepCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidForWorkScepCertificateProfile
func (r *AndroidForWorkScepCertificateProfileRequest) Update(ctx context.Context, reqObj *AndroidForWorkScepCertificateProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidForWorkScepCertificateProfile
func (r *AndroidForWorkScepCertificateProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *AndroidForWorkScepCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
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
	var values []ManagedDeviceCertificateState
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
			value  []ManagedDeviceCertificateState
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

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *AndroidForWorkScepCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

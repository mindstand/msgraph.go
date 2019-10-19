// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IOSImportedPFXCertificateProfileRequestBuilder is request builder for IOSImportedPFXCertificateProfile
type IOSImportedPFXCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSImportedPFXCertificateProfileRequest
func (b *IOSImportedPFXCertificateProfileRequestBuilder) Request() *IOSImportedPFXCertificateProfileRequest {
	return &IOSImportedPFXCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSImportedPFXCertificateProfileRequest is request for IOSImportedPFXCertificateProfile
type IOSImportedPFXCertificateProfileRequest struct{ BaseRequest }

// Do performs HTTP request for IOSImportedPFXCertificateProfile
func (r *IOSImportedPFXCertificateProfileRequest) Do(method, path string, reqObj interface{}) (resObj *IOSImportedPFXCertificateProfile, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for IOSImportedPFXCertificateProfile
func (r *IOSImportedPFXCertificateProfileRequest) Get() (*IOSImportedPFXCertificateProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for IOSImportedPFXCertificateProfile
func (r *IOSImportedPFXCertificateProfileRequest) Update(reqObj *IOSImportedPFXCertificateProfile) (*IOSImportedPFXCertificateProfile, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for IOSImportedPFXCertificateProfile
func (r *IOSImportedPFXCertificateProfileRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *IOSImportedPFXCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedDeviceCertificateState collection
func (r *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
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
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ManagedDeviceCertificateState
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

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get() ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *IOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(reqObj *ManagedDeviceCertificateState) (*ManagedDeviceCertificateState, error) {
	return r.Do("POST", "", reqObj)
}

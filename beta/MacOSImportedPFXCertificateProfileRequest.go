// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// MacOSImportedPFXCertificateProfileRequestBuilder is request builder for MacOSImportedPFXCertificateProfile
type MacOSImportedPFXCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSImportedPFXCertificateProfileRequest
func (b *MacOSImportedPFXCertificateProfileRequestBuilder) Request() *MacOSImportedPFXCertificateProfileRequest {
	return &MacOSImportedPFXCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSImportedPFXCertificateProfileRequest is request for MacOSImportedPFXCertificateProfile
type MacOSImportedPFXCertificateProfileRequest struct{ BaseRequest }

// Do performs HTTP request for MacOSImportedPFXCertificateProfile
func (r *MacOSImportedPFXCertificateProfileRequest) Do(method, path string, reqObj interface{}) (resObj *MacOSImportedPFXCertificateProfile, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for MacOSImportedPFXCertificateProfile
func (r *MacOSImportedPFXCertificateProfileRequest) Get() (*MacOSImportedPFXCertificateProfile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for MacOSImportedPFXCertificateProfile
func (r *MacOSImportedPFXCertificateProfileRequest) Update(reqObj *MacOSImportedPFXCertificateProfile) (*MacOSImportedPFXCertificateProfile, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for MacOSImportedPFXCertificateProfile
func (r *MacOSImportedPFXCertificateProfileRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *MacOSImportedPFXCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for ManagedDeviceCertificateState collection
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
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
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get() ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *MacOSImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(reqObj *ManagedDeviceCertificateState) (*ManagedDeviceCertificateState, error) {
	return r.Do("POST", "", reqObj)
}

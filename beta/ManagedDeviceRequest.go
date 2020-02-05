// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ManagedDeviceRequestBuilder is request builder for ManagedDevice
type ManagedDeviceRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceRequest
func (b *ManagedDeviceRequestBuilder) Request() *ManagedDeviceRequest {
	return &ManagedDeviceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Delta appends "/delta" onto the builder
func (b *ManagedDeviceRequestBuilder) Delta() *ManagedDeviceRequestBuilder {
	b.baseUrl = b.baseUrl + "/delta"
	return b
}

// ManagedDeviceRequest is request for ManagedDevice
type ManagedDeviceRequest struct{ BaseRequest }

// Get performs GET request for ManagedDevice
func (r *ManagedDeviceRequest) Get(ctx context.Context) (resObj *ManagedDevice, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDevice
func (r *ManagedDeviceRequest) Update(ctx context.Context, reqObj *ManagedDevice) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDevice
func (r *ManagedDeviceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// DetectedApps returns request builder for DetectedApp collection
func (b *ManagedDeviceRequestBuilder) DetectedApps() *ManagedDeviceDetectedAppsCollectionRequestBuilder {
	bb := &ManagedDeviceDetectedAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/detectedApps"
	return bb
}

// ManagedDeviceDetectedAppsCollectionRequestBuilder is request builder for DetectedApp collection
type ManagedDeviceDetectedAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DetectedApp collection
func (b *ManagedDeviceDetectedAppsCollectionRequestBuilder) Request() *ManagedDeviceDetectedAppsCollectionRequest {
	return &ManagedDeviceDetectedAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DetectedApp item
func (b *ManagedDeviceDetectedAppsCollectionRequestBuilder) ID(id string) *DetectedAppRequestBuilder {
	bb := &DetectedAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedDeviceDetectedAppsCollectionRequest is request for DetectedApp collection
type ManagedDeviceDetectedAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DetectedApp collection
func (r *ManagedDeviceDetectedAppsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DetectedApp, error) {
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
	var values []DetectedApp
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
			value  []DetectedApp
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

// Get performs GET request for DetectedApp collection
func (r *ManagedDeviceDetectedAppsCollectionRequest) Get(ctx context.Context) ([]DetectedApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DetectedApp collection
func (r *ManagedDeviceDetectedAppsCollectionRequest) Add(ctx context.Context, reqObj *DetectedApp) (resObj *DetectedApp, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceCategory is navigation property
func (b *ManagedDeviceRequestBuilder) DeviceCategory() *DeviceCategoryRequestBuilder {
	bb := &DeviceCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceCategory"
	return bb
}

// DeviceCompliancePolicyStates returns request builder for DeviceCompliancePolicyState collection
func (b *ManagedDeviceRequestBuilder) DeviceCompliancePolicyStates() *ManagedDeviceDeviceCompliancePolicyStatesCollectionRequestBuilder {
	bb := &ManagedDeviceDeviceCompliancePolicyStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceCompliancePolicyStates"
	return bb
}

// ManagedDeviceDeviceCompliancePolicyStatesCollectionRequestBuilder is request builder for DeviceCompliancePolicyState collection
type ManagedDeviceDeviceCompliancePolicyStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceCompliancePolicyState collection
func (b *ManagedDeviceDeviceCompliancePolicyStatesCollectionRequestBuilder) Request() *ManagedDeviceDeviceCompliancePolicyStatesCollectionRequest {
	return &ManagedDeviceDeviceCompliancePolicyStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceCompliancePolicyState item
func (b *ManagedDeviceDeviceCompliancePolicyStatesCollectionRequestBuilder) ID(id string) *DeviceCompliancePolicyStateRequestBuilder {
	bb := &DeviceCompliancePolicyStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedDeviceDeviceCompliancePolicyStatesCollectionRequest is request for DeviceCompliancePolicyState collection
type ManagedDeviceDeviceCompliancePolicyStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceCompliancePolicyState collection
func (r *ManagedDeviceDeviceCompliancePolicyStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceCompliancePolicyState, error) {
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
	var values []DeviceCompliancePolicyState
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
			value  []DeviceCompliancePolicyState
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

// Get performs GET request for DeviceCompliancePolicyState collection
func (r *ManagedDeviceDeviceCompliancePolicyStatesCollectionRequest) Get(ctx context.Context) ([]DeviceCompliancePolicyState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceCompliancePolicyState collection
func (r *ManagedDeviceDeviceCompliancePolicyStatesCollectionRequest) Add(ctx context.Context, reqObj *DeviceCompliancePolicyState) (resObj *DeviceCompliancePolicyState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// DeviceConfigurationStates returns request builder for DeviceConfigurationState collection
func (b *ManagedDeviceRequestBuilder) DeviceConfigurationStates() *ManagedDeviceDeviceConfigurationStatesCollectionRequestBuilder {
	bb := &ManagedDeviceDeviceConfigurationStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceConfigurationStates"
	return bb
}

// ManagedDeviceDeviceConfigurationStatesCollectionRequestBuilder is request builder for DeviceConfigurationState collection
type ManagedDeviceDeviceConfigurationStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceConfigurationState collection
func (b *ManagedDeviceDeviceConfigurationStatesCollectionRequestBuilder) Request() *ManagedDeviceDeviceConfigurationStatesCollectionRequest {
	return &ManagedDeviceDeviceConfigurationStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceConfigurationState item
func (b *ManagedDeviceDeviceConfigurationStatesCollectionRequestBuilder) ID(id string) *DeviceConfigurationStateRequestBuilder {
	bb := &DeviceConfigurationStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedDeviceDeviceConfigurationStatesCollectionRequest is request for DeviceConfigurationState collection
type ManagedDeviceDeviceConfigurationStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DeviceConfigurationState collection
func (r *ManagedDeviceDeviceConfigurationStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceConfigurationState, error) {
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
	var values []DeviceConfigurationState
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
			value  []DeviceConfigurationState
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

// Get performs GET request for DeviceConfigurationState collection
func (r *ManagedDeviceDeviceConfigurationStatesCollectionRequest) Get(ctx context.Context) ([]DeviceConfigurationState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceConfigurationState collection
func (r *ManagedDeviceDeviceConfigurationStatesCollectionRequest) Add(ctx context.Context, reqObj *DeviceConfigurationState) (resObj *DeviceConfigurationState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// ManagedDeviceMobileAppConfigurationStates returns request builder for ManagedDeviceMobileAppConfigurationState collection
func (b *ManagedDeviceRequestBuilder) ManagedDeviceMobileAppConfigurationStates() *ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequestBuilder {
	bb := &ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceMobileAppConfigurationStates"
	return bb
}

// ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationState collection
type ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceMobileAppConfigurationState collection
func (b *ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequestBuilder) Request() *ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequest {
	return &ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceMobileAppConfigurationState item
func (b *ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceMobileAppConfigurationStateRequestBuilder {
	bb := &ManagedDeviceMobileAppConfigurationStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequest is request for ManagedDeviceMobileAppConfigurationState collection
type ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceMobileAppConfigurationState collection
func (r *ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]ManagedDeviceMobileAppConfigurationState, error) {
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
	var values []ManagedDeviceMobileAppConfigurationState
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
			value  []ManagedDeviceMobileAppConfigurationState
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

// Get performs GET request for ManagedDeviceMobileAppConfigurationState collection
func (r *ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequest) Get(ctx context.Context) ([]ManagedDeviceMobileAppConfigurationState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for ManagedDeviceMobileAppConfigurationState collection
func (r *ManagedDeviceManagedDeviceMobileAppConfigurationStatesCollectionRequest) Add(ctx context.Context, reqObj *ManagedDeviceMobileAppConfigurationState) (resObj *ManagedDeviceMobileAppConfigurationState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// SecurityBaselineStates returns request builder for SecurityBaselineState collection
func (b *ManagedDeviceRequestBuilder) SecurityBaselineStates() *ManagedDeviceSecurityBaselineStatesCollectionRequestBuilder {
	bb := &ManagedDeviceSecurityBaselineStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/securityBaselineStates"
	return bb
}

// ManagedDeviceSecurityBaselineStatesCollectionRequestBuilder is request builder for SecurityBaselineState collection
type ManagedDeviceSecurityBaselineStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SecurityBaselineState collection
func (b *ManagedDeviceSecurityBaselineStatesCollectionRequestBuilder) Request() *ManagedDeviceSecurityBaselineStatesCollectionRequest {
	return &ManagedDeviceSecurityBaselineStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SecurityBaselineState item
func (b *ManagedDeviceSecurityBaselineStatesCollectionRequestBuilder) ID(id string) *SecurityBaselineStateRequestBuilder {
	bb := &SecurityBaselineStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedDeviceSecurityBaselineStatesCollectionRequest is request for SecurityBaselineState collection
type ManagedDeviceSecurityBaselineStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SecurityBaselineState collection
func (r *ManagedDeviceSecurityBaselineStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]SecurityBaselineState, error) {
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
	var values []SecurityBaselineState
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
			value  []SecurityBaselineState
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

// Get performs GET request for SecurityBaselineState collection
func (r *ManagedDeviceSecurityBaselineStatesCollectionRequest) Get(ctx context.Context) ([]SecurityBaselineState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for SecurityBaselineState collection
func (r *ManagedDeviceSecurityBaselineStatesCollectionRequest) Add(ctx context.Context, reqObj *SecurityBaselineState) (resObj *SecurityBaselineState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Users returns request builder for User collection
func (b *ManagedDeviceRequestBuilder) Users() *ManagedDeviceUsersCollectionRequestBuilder {
	bb := &ManagedDeviceUsersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/users"
	return bb
}

// ManagedDeviceUsersCollectionRequestBuilder is request builder for User collection
type ManagedDeviceUsersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for User collection
func (b *ManagedDeviceUsersCollectionRequestBuilder) Request() *ManagedDeviceUsersCollectionRequest {
	return &ManagedDeviceUsersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for User item
func (b *ManagedDeviceUsersCollectionRequestBuilder) ID(id string) *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ManagedDeviceUsersCollectionRequest is request for User collection
type ManagedDeviceUsersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for User collection
func (r *ManagedDeviceUsersCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]User, error) {
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
	var values []User
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
			value  []User
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

// Get performs GET request for User collection
func (r *ManagedDeviceUsersCollectionRequest) Get(ctx context.Context) ([]User, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for User collection
func (r *ManagedDeviceUsersCollectionRequest) Add(ctx context.Context, reqObj *User) (resObj *User, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// WindowsProtectionState is navigation property
func (b *ManagedDeviceRequestBuilder) WindowsProtectionState() *WindowsProtectionStateRequestBuilder {
	bb := &WindowsProtectionStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/windowsProtectionState"
	return bb
}

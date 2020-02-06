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

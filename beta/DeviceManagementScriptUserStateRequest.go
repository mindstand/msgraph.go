// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeviceManagementScriptUserStateRequestBuilder is request builder for DeviceManagementScriptUserState
type DeviceManagementScriptUserStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementScriptUserStateRequest
func (b *DeviceManagementScriptUserStateRequestBuilder) Request() *DeviceManagementScriptUserStateRequest {
	return &DeviceManagementScriptUserStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementScriptUserStateRequest is request for DeviceManagementScriptUserState
type DeviceManagementScriptUserStateRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScriptUserState
func (r *DeviceManagementScriptUserStateRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScriptUserState, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for DeviceManagementScriptUserState
func (r *DeviceManagementScriptUserStateRequest) Get() (*DeviceManagementScriptUserState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for DeviceManagementScriptUserState
func (r *DeviceManagementScriptUserStateRequest) Update(reqObj *DeviceManagementScriptUserState) (*DeviceManagementScriptUserState, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for DeviceManagementScriptUserState
func (r *DeviceManagementScriptUserStateRequest) Delete() error {
	return r.JSONRequestWithPath("DELETE", "", nil, nil)
}

// DeviceRunStates returns request builder for DeviceManagementScriptDeviceState collection
func (b *DeviceManagementScriptUserStateRequestBuilder) DeviceRunStates() *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequestBuilder {
	bb := &DeviceManagementScriptUserStateDeviceRunStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceRunStates"
	return bb
}

// DeviceManagementScriptUserStateDeviceRunStatesCollectionRequestBuilder is request builder for DeviceManagementScriptDeviceState collection
type DeviceManagementScriptUserStateDeviceRunStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DeviceManagementScriptDeviceState collection
func (b *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequestBuilder) Request() *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest {
	return &DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DeviceManagementScriptDeviceState item
func (b *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequestBuilder) ID(id string) *DeviceManagementScriptDeviceStateRequestBuilder {
	bb := &DeviceManagementScriptDeviceStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest is request for DeviceManagementScriptDeviceState collection
type DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DeviceManagementScriptDeviceState, err error) {
	err = r.JSONRequestWithPath(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]DeviceManagementScriptDeviceState, error) {
	req, err := r.NewJSONRequestWithPath(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DeviceManagementScriptDeviceState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DeviceManagementScriptDeviceState
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

// Get performs GET request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest) Get() ([]DeviceManagementScriptDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest) Add(reqObj *DeviceManagementScriptDeviceState) (*DeviceManagementScriptDeviceState, error) {
	return r.Do("POST", "", reqObj)
}

// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
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

// Get performs GET request for DeviceManagementScriptUserState
func (r *DeviceManagementScriptUserStateRequest) Get(ctx context.Context) (resObj *DeviceManagementScriptUserState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementScriptUserState
func (r *DeviceManagementScriptUserStateRequest) Update(ctx context.Context, reqObj *DeviceManagementScriptUserState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementScriptUserState
func (r *DeviceManagementScriptUserStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
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

// Paging perfoms paging operation for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}) ([]DeviceManagementScriptDeviceState, error) {
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
	var values []DeviceManagementScriptDeviceState
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
			value  []DeviceManagementScriptDeviceState
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

// Get performs GET request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest) Get(ctx context.Context) ([]DeviceManagementScriptDeviceState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil)
}

// Add performs POST request for DeviceManagementScriptDeviceState collection
func (r *DeviceManagementScriptUserStateDeviceRunStatesCollectionRequest) Add(ctx context.Context, reqObj *DeviceManagementScriptDeviceState) (resObj *DeviceManagementScriptDeviceState, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
